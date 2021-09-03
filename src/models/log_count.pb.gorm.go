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
	Count uint64
	Id    uint64
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
	to.Id = m.Id
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
	to.Id = m.Id
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

// DefaultReadLogCount executes a basic gorm read call
func DefaultReadLogCount(ctx context.Context, in *LogCount, db *gorm1.DB) (*LogCount, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(LogCountORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &LogCountORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogCountORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := LogCountORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(LogCountORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type LogCountORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteLogCount(ctx context.Context, in *LogCount, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(LogCountORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&LogCountORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(LogCountORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type LogCountORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteLogCountSet(ctx context.Context, in []*LogCount, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	var err error
	keys := []uint64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors1.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&LogCountORM{})).(LogCountORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&LogCountORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&LogCountORM{})).(LogCountORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type LogCountORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*LogCount, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*LogCount, *gorm1.DB) error
}

// DefaultStrictUpdateLogCount clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateLogCount(ctx context.Context, in *LogCount, db *gorm1.DB) (*LogCount, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateLogCount")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &LogCountORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(LogCountORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(LogCountORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(LogCountORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type LogCountORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchLogCount executes a basic gorm update call with patch behavior
func DefaultPatchLogCount(ctx context.Context, in *LogCount, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*LogCount, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	var pbObj LogCount
	var err error
	if hook, ok := interface{}(&pbObj).(LogCountWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadLogCount(ctx, &LogCount{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(LogCountWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskLogCount(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(LogCountWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateLogCount(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(LogCountWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type LogCountWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *LogCount, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *LogCount, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *LogCount, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type LogCountWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *LogCount, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultPatchSetLogCount executes a bulk gorm update call with patch behavior
func DefaultPatchSetLogCount(ctx context.Context, objects []*LogCount, updateMasks []*field_mask1.FieldMask, db *gorm1.DB) ([]*LogCount, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors1.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*LogCount, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchLogCount(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
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
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
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
	db = db.Order("id")
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