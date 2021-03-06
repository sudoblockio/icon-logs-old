// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: log_count.proto

package models

import (
	context "context"
	fmt "fmt"
	
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	math "math"

	gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
	errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
	gorm1 "github.com/jinzhu/gorm"
	field_mask1 "google.golang.org/genproto/protobuf/field_mask"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type LogCountORM struct {
	Count           uint64
	LogIndex        uint64
	TransactionHash string
	Type            string `gorm:"primary_key"`
}

// TableName overrides the default tablename generated by GORM
func (LogCountORM) TableName() string {
	return "log_counts"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *LogCount) ToORM(ctx context.Context) (LogCountORM, error) {
	to := LogCountORM{}
	var err error
	if prehook, ok := interface{}(m).(LogCountWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Count = m.Count
	to.Type = m.Type
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	if posthook, ok := interface{}(m).(LogCountWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *LogCountORM) ToPB(ctx context.Context) (LogCount, error) {
	to := LogCount{}
	var err error
	if prehook, ok := interface{}(m).(LogCountWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Count = m.Count
	to.Type = m.Type
	to.TransactionHash = m.TransactionHash
	to.LogIndex = m.LogIndex
	if posthook, ok := interface{}(m).(LogCountWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type LogCount the arg will be the target, the caller the one being converted from

// LogCountBeforeToORM called before default ToORM code
type LogCountWithBeforeToORM interface {
	BeforeToORM(context.Context, *LogCountORM) error
}

// LogCountAfterToORM called after default ToORM code
type LogCountWithAfterToORM interface {
	AfterToORM(context.Context, *LogCountORM) error
}

// LogCountBeforeToPB called before default ToPB code
type LogCountWithBeforeToPB interface {
	BeforeToPB(context.Context, *LogCount) error
}

// LogCountAfterToPB called after default ToPB code
type LogCountWithAfterToPB interface {
	AfterToPB(context.Context, *LogCount) error
}

// DefaultCreateLogCount executes a basic gorm create call
func DefaultCreateLogCount(ctx context.Context, in *LogCount, db *gorm1.DB) (*LogCount, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogCountORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogCountORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type LogCountORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultApplyFieldMaskLogCount patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskLogCount(ctx context.Context, patchee *LogCount, patcher *LogCount, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*LogCount, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Count" {
			patchee.Count = patcher.Count
			continue
		}
		if f == prefix+"Type" {
			patchee.Type = patcher.Type
			continue
		}
		if f == prefix+"TransactionHash" {
			patchee.TransactionHash = patcher.TransactionHash
			continue
		}
		if f == prefix+"LogIndex" {
			patchee.LogIndex = patcher.LogIndex
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListLogCount executes a gorm list call
func DefaultListLogCount(ctx context.Context, db *gorm1.DB) ([]*LogCount, error) {
	in := LogCount{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogCountORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &LogCountORM{}, &LogCount{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogCountORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("type")
	ormResponse := []LogCountORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogCountORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*LogCount{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type LogCountORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]LogCountORM) error
}
