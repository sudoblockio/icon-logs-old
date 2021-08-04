package crud

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
	"github.com/cenkalti/backoff/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/geometry-labs/icon-logs/config"
	"github.com/geometry-labs/icon-logs/global"
)

type MongoConn struct {
	client *mongo.Client
}

var mongoInstance *MongoConn
var mongoConnOnce sync.Once

func GetMongoConn() *MongoConn {
	mongoConnOnce.Do(func() {
		uri := fmt.Sprintf("%s://%s:%s", config.Config.DbDriver, config.Config.DbHost, config.Config.DbPort) //"mongodb://localhost:27017"
		// Getting client (or session)
		client, err := retryMongoConn(uri)
		if err != nil {
			zap.S().Fatal("MONGO: Connection cannot be established")
		} else {
			zap.S().Info("MONGO: Connection established")
		}

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			defer cancel()
			global.WaitShutdownSig()
			zap.S().Info("Closing Mongodb client context")
		}()

		err = client.Connect(ctx)
		if err != nil {
			zap.S().Fatal("Cannot connect to context for mongodb", err)
		}
		go clientClose(client)

		mongoInstance = &MongoConn{
			client: client,
		}

		err = mongoInstance.retryPing(ctx)
		if err != nil {
			zap.S().Fatal("MONGO: Finally cannot ping mongodb")
		} else {
			zap.S().Info("MONGO: Finally pinged mongodb")
		}
	})

	return mongoInstance
}

func (m *MongoConn) GetClient() *mongo.Client {
	return m.client
}

func clientClose(client *mongo.Client) error {

	global.WaitShutdownSig()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := client.Disconnect(ctx)
	if err != nil {
		zap.S().Fatal("Cannot disconnect from mongodb", err)
	}
	return err
}

func (m *MongoConn) Ping(ctx context.Context) error {
	err := m.client.Ping(ctx, readpref.Primary())
	if err != nil {
		zap.S().Info("Cannot ping mongodb", err)
	}
	return err
}

func (m *MongoConn) retryPing(ctx context.Context) error {
	operation := func() error {
		return m.Ping(ctx)
	}
	neb := backoff.NewExponentialBackOff()
	err := backoff.Retry(operation, neb)

	return err
}

func (m *MongoConn) ListAllDatabases() []string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	databases, err := m.client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		zap.S().Fatal("Cannot List databases", err)
	}
	return databases
}

func (m *MongoConn) DatabaseHandle(database string) *mongo.Database {
	return m.client.Database(database)
}

func retryMongoConn(uri string) (*mongo.Client, error) {
	var client *mongo.Client
	operation := func() error {
		cli, err := mongo.NewClient(options.Client().ApplyURI(uri).SetAuth(options.Credential{
			AuthMechanism:           "",
			AuthMechanismProperties: nil,
			AuthSource:              "",
			Username:                config.Config.DbUser,     //"mongo",
			Password:                config.Config.DbPassword, //"changethis",
			PasswordSet:             true,
		}))
		if err != nil {
			zap.S().Fatal("Cannot create a connection to mongodb", err)
		} else {
			client = cli
		}
		return err
	}
	neb := backoff.NewExponentialBackOff()
	err := backoff.Retry(operation, neb)
	return client, err
}
