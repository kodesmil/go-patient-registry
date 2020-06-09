// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/kodesmil/ks-model/group.proto

// Generated with protoc-gen-gorm version: master
// Anticipating compatibility with atlas-app-toolkit version: master

package pb

import context "context"
import fmt "fmt"

import auth1 "github.com/infobloxopen/atlas-app-toolkit/auth"
import errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
import field_mask1 "google.golang.org/genproto/protobuf/field_mask"
import gorm1 "github.com/jinzhu/gorm"
import gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
import query1 "github.com/infobloxopen/atlas-app-toolkit/query"
import resource1 "github.com/infobloxopen/atlas-app-toolkit/gorm/resource"

import math "math"
import _ "google.golang.org/genproto/protobuf/field_mask"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"
import _ "github.com/infobloxopen/atlas-app-toolkit/query"
import _ "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type GroupORM struct {
	AccountID string
	Id        int64 `gorm:"type:serial;primary_key"`
	Name      string
	Notes     string
	ProfileId *string
}

// TableName overrides the default tablename generated by GORM
func (GroupORM) TableName() string {
	return "groups"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Group) ToORM(ctx context.Context) (GroupORM, error) {
	to := GroupORM{}
	var err error
	if prehook, ok := interface{}(m).(GroupWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource1.DecodeInt64(&Group{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	to.Name = m.Name
	to.Notes = m.Notes
	if m.ProfileId != nil {
		if v, err := resource1.Decode(&Profile{}, m.ProfileId); err != nil {
			return to, err
		} else if v != nil {
			vv := v.(string)
			to.ProfileId = &vv
		}
	}
	accountID, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return to, err
	}
	to.AccountID = accountID
	if posthook, ok := interface{}(m).(GroupWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *GroupORM) ToPB(ctx context.Context) (Group, error) {
	to := Group{}
	var err error
	if prehook, ok := interface{}(m).(GroupWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource1.Encode(&Group{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	to.Name = m.Name
	to.Notes = m.Notes
	if m.ProfileId != nil {
		if v, err := resource1.Encode(&Profile{}, *m.ProfileId); err != nil {
			return to, err
		} else {
			to.ProfileId = v
		}
	}
	if posthook, ok := interface{}(m).(GroupWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Group the arg will be the target, the caller the one being converted from

// GroupBeforeToORM called before default ToORM code
type GroupWithBeforeToORM interface {
	BeforeToORM(context.Context, *GroupORM) error
}

// GroupAfterToORM called after default ToORM code
type GroupWithAfterToORM interface {
	AfterToORM(context.Context, *GroupORM) error
}

// GroupBeforeToPB called before default ToPB code
type GroupWithBeforeToPB interface {
	BeforeToPB(context.Context, *Group) error
}

// GroupAfterToPB called after default ToPB code
type GroupWithAfterToPB interface {
	AfterToPB(context.Context, *Group) error
}

// DefaultCreateGroup executes a basic gorm create call
func DefaultCreateGroup(ctx context.Context, in *Group, db *gorm1.DB) (*Group, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(GroupORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(GroupORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type GroupORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type GroupORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultReadGroup executes a basic gorm read call
func DefaultReadGroup(ctx context.Context, in *Group, db *gorm1.DB) (*Group, error) {
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
	if hook, ok := interface{}(&ormObj).(GroupORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &GroupORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(GroupORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := GroupORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(GroupORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type GroupORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type GroupORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type GroupORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteGroup(ctx context.Context, in *Group, db *gorm1.DB) error {
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
	if hook, ok := interface{}(&ormObj).(GroupORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&GroupORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(GroupORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type GroupORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type GroupORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteGroupSet(ctx context.Context, in []*Group, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	var err error
	keys := []int64{}
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
	if hook, ok := (interface{}(&GroupORM{})).(GroupORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	acctId, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return err
	}
	err = db.Where("account_id = ? AND id in (?)", acctId, keys).Delete(&GroupORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&GroupORM{})).(GroupORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type GroupORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Group, *gorm1.DB) (*gorm1.DB, error)
}
type GroupORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Group, *gorm1.DB) error
}

// DefaultStrictUpdateGroup clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateGroup(ctx context.Context, in *Group, db *gorm1.DB) (*Group, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateGroup")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	accountID, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return nil, err
	}
	db = db.Where(map[string]interface{}{"account_id": accountID})
	lockedRow := &GroupORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(GroupORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(GroupORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(GroupORMWithAfterStrictUpdateSave); ok {
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

type GroupORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type GroupORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type GroupORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchGroup executes a basic gorm update call with patch behavior
func DefaultPatchGroup(ctx context.Context, in *Group, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*Group, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	var pbObj Group
	var err error
	if hook, ok := interface{}(&pbObj).(GroupWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadGroup(ctx, &Group{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(GroupWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskGroup(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(GroupWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateGroup(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(GroupWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type GroupWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Group, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type GroupWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Group, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type GroupWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Group, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type GroupWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Group, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultPatchSetGroup executes a bulk gorm update call with patch behavior
func DefaultPatchSetGroup(ctx context.Context, objects []*Group, updateMasks []*field_mask1.FieldMask, db *gorm1.DB) ([]*Group, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors1.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Group, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchGroup(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskGroup patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskGroup(ctx context.Context, patchee *Group, patcher *Group, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*Group, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"Notes" {
			patchee.Notes = patcher.Notes
			continue
		}
		if f == prefix+"ProfileId" {
			patchee.ProfileId = patcher.ProfileId
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListGroup executes a gorm list call
func DefaultListGroup(ctx context.Context, db *gorm1.DB, f *query1.Filtering, s *query1.Sorting, p *query1.Pagination, fs *query1.FieldSelection) ([]*Group, error) {
	in := Group{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(GroupORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &GroupORM{}, &Group{}, f, s, p, fs)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(GroupORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []GroupORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(GroupORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Group{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type GroupORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) (*gorm1.DB, error)
}
type GroupORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) (*gorm1.DB, error)
}
type GroupORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]GroupORM, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) error
}
type GroupsDefaultServer struct {
	DB *gorm1.DB
}

// Create ...
func (m *GroupsDefaultServer) Create(ctx context.Context, in *CreateGroupRequest) (*CreateGroupResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(GroupsGroupWithBeforeCreate); ok {
		var err error
		if db, err = custom.BeforeCreate(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultCreateGroup(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	out := &CreateGroupResponse{Result: res}
	if custom, ok := interface{}(in).(GroupsGroupWithAfterCreate); ok {
		var err error
		if err = custom.AfterCreate(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// GroupsGroupWithBeforeCreate called before DefaultCreateGroup in the default Create handler
type GroupsGroupWithBeforeCreate interface {
	BeforeCreate(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// GroupsGroupWithAfterCreate called before DefaultCreateGroup in the default Create handler
type GroupsGroupWithAfterCreate interface {
	AfterCreate(context.Context, *CreateGroupResponse, *gorm1.DB) error
}

// Read ...
func (m *GroupsDefaultServer) Read(ctx context.Context, in *ReadGroupRequest) (*ReadGroupResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(GroupsGroupWithBeforeRead); ok {
		var err error
		if db, err = custom.BeforeRead(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultReadGroup(ctx, &Group{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &ReadGroupResponse{Result: res}
	if custom, ok := interface{}(in).(GroupsGroupWithAfterRead); ok {
		var err error
		if err = custom.AfterRead(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// GroupsGroupWithBeforeRead called before DefaultReadGroup in the default Read handler
type GroupsGroupWithBeforeRead interface {
	BeforeRead(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// GroupsGroupWithAfterRead called before DefaultReadGroup in the default Read handler
type GroupsGroupWithAfterRead interface {
	AfterRead(context.Context, *ReadGroupResponse, *gorm1.DB) error
}

// Update ...
func (m *GroupsDefaultServer) Update(ctx context.Context, in *UpdateGroupRequest) (*UpdateGroupResponse, error) {
	var err error
	var res *Group
	db := m.DB
	if custom, ok := interface{}(in).(GroupsGroupWithBeforeUpdate); ok {
		var err error
		if db, err = custom.BeforeUpdate(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err = DefaultStrictUpdateGroup(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	out := &UpdateGroupResponse{Result: res}
	if custom, ok := interface{}(in).(GroupsGroupWithAfterUpdate); ok {
		var err error
		if err = custom.AfterUpdate(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// GroupsGroupWithBeforeUpdate called before DefaultUpdateGroup in the default Update handler
type GroupsGroupWithBeforeUpdate interface {
	BeforeUpdate(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// GroupsGroupWithAfterUpdate called before DefaultUpdateGroup in the default Update handler
type GroupsGroupWithAfterUpdate interface {
	AfterUpdate(context.Context, *UpdateGroupResponse, *gorm1.DB) error
}

// Delete ...
func (m *GroupsDefaultServer) Delete(ctx context.Context, in *DeleteGroupRequest) (*DeleteGroupResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(GroupsGroupWithBeforeDelete); ok {
		var err error
		if db, err = custom.BeforeDelete(ctx, db); err != nil {
			return nil, err
		}
	}
	err := DefaultDeleteGroup(ctx, &Group{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &DeleteGroupResponse{}
	if custom, ok := interface{}(in).(GroupsGroupWithAfterDelete); ok {
		var err error
		if err = custom.AfterDelete(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// GroupsGroupWithBeforeDelete called before DefaultDeleteGroup in the default Delete handler
type GroupsGroupWithBeforeDelete interface {
	BeforeDelete(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// GroupsGroupWithAfterDelete called before DefaultDeleteGroup in the default Delete handler
type GroupsGroupWithAfterDelete interface {
	AfterDelete(context.Context, *DeleteGroupResponse, *gorm1.DB) error
}

// List ...
func (m *GroupsDefaultServer) List(ctx context.Context, in *ListGroupRequest) (*ListGroupResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(GroupsGroupWithBeforeList); ok {
		var err error
		if db, err = custom.BeforeList(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultListGroup(ctx, db, in.Filter, in.OrderBy, in.Paging, in.Fields)
	if err != nil {
		return nil, err
	}
	out := &ListGroupResponse{Results: res}
	if custom, ok := interface{}(in).(GroupsGroupWithAfterList); ok {
		var err error
		if err = custom.AfterList(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// GroupsGroupWithBeforeList called before DefaultListGroup in the default List handler
type GroupsGroupWithBeforeList interface {
	BeforeList(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// GroupsGroupWithAfterList called before DefaultListGroup in the default List handler
type GroupsGroupWithAfterList interface {
	AfterList(context.Context, *ListGroupResponse, *gorm1.DB) error
}
