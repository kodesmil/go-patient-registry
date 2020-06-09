// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/kodesmil/ks-model/profile.proto

// Generated with protoc-gen-gorm version: master
// Anticipating compatibility with atlas-app-toolkit version: master

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	github.com/kodesmil/ks-model/profile.proto
	github.com/kodesmil/ks-model/log.proto
	github.com/kodesmil/ks-model/chat.proto
	github.com/kodesmil/ks-model/group.proto
	github.com/kodesmil/ks-model/feed.proto
	github.com/kodesmil/ks-model/journal.proto
	github.com/kodesmil/ks-model/period.proto
	github.com/kodesmil/ks-model/notification.proto

It has these top-level messages:
	Profile
	CreateProfileRequest
	CreateProfileResponse
	ReadProfileRequest
	ReadProfileResponse
	UpdateProfileRequest
	UpdateProfileResponse
	DeleteProfileRequest
	DeleteProfileResponse
	ListProfileRequest
	ListProfileResponse
	LogActivity
	ChatMessage
	ChatRoom
	ChatRoomParticipant
	StreamChatEvent
	EventNone
	EventLoadRoom
	EventLoadRooms
	EventLeaveRoom
	EventLeaveRooms
	EventSendMessage
	EventSendMessages
	EventSendRooms
	EventInviteProfile
	EventForceClose
	ListChatMessageRequest
	ListChatMessageResponse
	ListChatRoomRequest
	ListChatRoomResponse
	Group
	CreateGroupRequest
	CreateGroupResponse
	ReadGroupRequest
	ReadGroupResponse
	UpdateGroupRequest
	UpdateGroupResponse
	DeleteGroupRequest
	DeleteGroupResponse
	ListGroupRequest
	ListGroupResponse
	FeedTag
	FeedAuthor
	FeedArticleDetail
	ReadFeedArticleDetailsRequest
	ReadFeedArticleDetailsResponse
	FeedArticle
	ListFeedArticleRequest
	ListFeedArticleResponse
	JournalSubjectType
	JournalSubject
	JournalEntry
	CreateJournalEntryRequest
	CreateJournalEntryResponse
	ReadJournalEntryRequest
	ReadJournalEntryResponse
	UpdateJournalEntryRequest
	UpdateJournalEntryResponse
	DeleteJournalEntryRequest
	DeleteJournalEntryResponse
	ListJournalEntryRequest
	ListJournalEntryResponse
	PeriodInfo
	PeriodDailyEntry
	CreatePeriodDailyEntryRequest
	CreatePeriodDailyEntryResponse
	ReadPeriodDailyEntryRequest
	ReadPeriodDailyEntryResponse
	UpdatePeriodDailyEntryRequest
	UpdatePeriodDailyEntryResponse
	DeletePeriodDailyEntryRequest
	DeletePeriodDailyEntryResponse
	ListPeriodDailyEntryRequest
	ListPeriodDailyEntryResponse
	CreatePeriodInfoRequest
	CreatePeriodInfoResponse
	ReadPeriodInfoRequest
	ReadPeriodInfoResponse
	UpdatePeriodInfoRequest
	UpdatePeriodInfoResponse
	DeletePeriodInfoRequest
	DeletePeriodInfoResponse
	ListPeriodInfoRequest
	ListPeriodInfoResponse
	NotificationSetting
	CreateNotificationSettingRequest
	CreateNotificationSettingResponse
	ReadNotificationSettingRequest
	ReadNotificationSettingResponse
	UpdateNotificationSettingRequest
	UpdateNotificationSettingResponse
	DeleteNotificationSettingRequest
	DeleteNotificationSettingResponse
	ListNotificationSettingRequest
	ListNotificationSettingResponse
	NotificationDevice
	CreateNotificationDeviceRequest
	CreateNotificationDeviceResponse
*/
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

type ProfileORM struct {
	AccountID         string
	FirstName         string
	Groups            []*GroupORM `gorm:"foreignkey:ProfileId;association_foreignkey:Id"`
	Id                string      `gorm:"type:text;primary_key;not null"`
	LastName          string
	Notes             string
	PrimaryEmail      string `gorm:"unique"`
	ProfilePictureUrl string
}

// TableName overrides the default tablename generated by GORM
func (ProfileORM) TableName() string {
	return "profiles"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Profile) ToORM(ctx context.Context) (ProfileORM, error) {
	to := ProfileORM{}
	var err error
	if prehook, ok := interface{}(m).(ProfileWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource1.Decode(&Profile{}, m.Id); err != nil {
		return to, err
	} else if v != nil {
		to.Id = v.(string)
	}
	to.Notes = m.Notes
	to.FirstName = m.FirstName
	to.LastName = m.LastName
	to.PrimaryEmail = m.PrimaryEmail
	for _, v := range m.Groups {
		if v != nil {
			if tempGroups, cErr := v.ToORM(ctx); cErr == nil {
				to.Groups = append(to.Groups, &tempGroups)
			} else {
				return to, cErr
			}
		} else {
			to.Groups = append(to.Groups, nil)
		}
	}
	to.ProfilePictureUrl = m.ProfilePictureUrl
	accountID, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return to, err
	}
	to.AccountID = accountID
	if posthook, ok := interface{}(m).(ProfileWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ProfileORM) ToPB(ctx context.Context) (Profile, error) {
	to := Profile{}
	var err error
	if prehook, ok := interface{}(m).(ProfileWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource1.Encode(&Profile{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	to.Notes = m.Notes
	to.FirstName = m.FirstName
	to.LastName = m.LastName
	to.PrimaryEmail = m.PrimaryEmail
	for _, v := range m.Groups {
		if v != nil {
			if tempGroups, cErr := v.ToPB(ctx); cErr == nil {
				to.Groups = append(to.Groups, &tempGroups)
			} else {
				return to, cErr
			}
		} else {
			to.Groups = append(to.Groups, nil)
		}
	}
	to.ProfilePictureUrl = m.ProfilePictureUrl
	if posthook, ok := interface{}(m).(ProfileWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Profile the arg will be the target, the caller the one being converted from

// ProfileBeforeToORM called before default ToORM code
type ProfileWithBeforeToORM interface {
	BeforeToORM(context.Context, *ProfileORM) error
}

// ProfileAfterToORM called after default ToORM code
type ProfileWithAfterToORM interface {
	AfterToORM(context.Context, *ProfileORM) error
}

// ProfileBeforeToPB called before default ToPB code
type ProfileWithBeforeToPB interface {
	BeforeToPB(context.Context, *Profile) error
}

// ProfileAfterToPB called after default ToPB code
type ProfileWithAfterToPB interface {
	AfterToPB(context.Context, *Profile) error
}

// DefaultCreateProfile executes a basic gorm create call
func DefaultCreateProfile(ctx context.Context, in *Profile, db *gorm1.DB) (*Profile, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type ProfileORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultReadProfile executes a basic gorm read call
func DefaultReadProfile(ctx context.Context, in *Profile, db *gorm1.DB) (*Profile, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == "" {
		return nil, errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &ProfileORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := ProfileORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(ProfileORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type ProfileORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteProfile(ctx context.Context, in *Profile, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&ProfileORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type ProfileORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteProfileSet(ctx context.Context, in []*Profile, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	var err error
	keys := []string{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == "" {
			return errors1.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&ProfileORM{})).(ProfileORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	acctId, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return err
	}
	err = db.Where("account_id = ? AND id in (?)", acctId, keys).Delete(&ProfileORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&ProfileORM{})).(ProfileORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type ProfileORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*Profile, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*Profile, *gorm1.DB) error
}

// DefaultStrictUpdateProfile clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateProfile(ctx context.Context, in *Profile, db *gorm1.DB) (*Profile, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateProfile")
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
	lockedRow := &ProfileORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	filterGroups := GroupORM{}
	if ormObj.Id == "" {
		return nil, errors1.EmptyIdError
	}
	filterGroups.ProfileId = new(string)
	*filterGroups.ProfileId = ormObj.Id
	if err = db.Where(filterGroups).Delete(GroupORM{}).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithAfterStrictUpdateSave); ok {
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

type ProfileORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchProfile executes a basic gorm update call with patch behavior
func DefaultPatchProfile(ctx context.Context, in *Profile, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*Profile, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	var pbObj Profile
	var err error
	if hook, ok := interface{}(&pbObj).(ProfileWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadProfile(ctx, &Profile{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(ProfileWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskProfile(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(ProfileWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateProfile(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(ProfileWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type ProfileWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *Profile, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *Profile, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *Profile, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type ProfileWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *Profile, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultPatchSetProfile executes a bulk gorm update call with patch behavior
func DefaultPatchSetProfile(ctx context.Context, objects []*Profile, updateMasks []*field_mask1.FieldMask, db *gorm1.DB) ([]*Profile, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors1.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*Profile, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchProfile(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskProfile patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskProfile(ctx context.Context, patchee *Profile, patcher *Profile, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*Profile, error) {
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
		if f == prefix+"Notes" {
			patchee.Notes = patcher.Notes
			continue
		}
		if f == prefix+"FirstName" {
			patchee.FirstName = patcher.FirstName
			continue
		}
		if f == prefix+"LastName" {
			patchee.LastName = patcher.LastName
			continue
		}
		if f == prefix+"PrimaryEmail" {
			patchee.PrimaryEmail = patcher.PrimaryEmail
			continue
		}
		if f == prefix+"Groups" {
			patchee.Groups = patcher.Groups
			continue
		}
		if f == prefix+"ProfilePictureUrl" {
			patchee.ProfilePictureUrl = patcher.ProfilePictureUrl
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListProfile executes a gorm list call
func DefaultListProfile(ctx context.Context, db *gorm1.DB, f *query1.Filtering, s *query1.Sorting, p *query1.Pagination, fs *query1.FieldSelection) ([]*Profile, error) {
	in := Profile{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &ProfileORM{}, &Profile{}, f, s, p, fs)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []ProfileORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ProfileORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	pbResponse := []*Profile{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type ProfileORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) (*gorm1.DB, error)
}
type ProfileORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) (*gorm1.DB, error)
}
type ProfileORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]ProfileORM, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) error
}
type ProfilesDefaultServer struct {
	DB *gorm1.DB
}

// Create ...
func (m *ProfilesDefaultServer) Create(ctx context.Context, in *CreateProfileRequest) (*CreateProfileResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(ProfilesProfileWithBeforeCreate); ok {
		var err error
		if db, err = custom.BeforeCreate(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultCreateProfile(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	out := &CreateProfileResponse{Result: res}
	if custom, ok := interface{}(in).(ProfilesProfileWithAfterCreate); ok {
		var err error
		if err = custom.AfterCreate(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ProfilesProfileWithBeforeCreate called before DefaultCreateProfile in the default Create handler
type ProfilesProfileWithBeforeCreate interface {
	BeforeCreate(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// ProfilesProfileWithAfterCreate called before DefaultCreateProfile in the default Create handler
type ProfilesProfileWithAfterCreate interface {
	AfterCreate(context.Context, *CreateProfileResponse, *gorm1.DB) error
}

// Read ...
func (m *ProfilesDefaultServer) Read(ctx context.Context, in *ReadProfileRequest) (*ReadProfileResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(ProfilesProfileWithBeforeRead); ok {
		var err error
		if db, err = custom.BeforeRead(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultReadProfile(ctx, &Profile{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &ReadProfileResponse{Result: res}
	if custom, ok := interface{}(in).(ProfilesProfileWithAfterRead); ok {
		var err error
		if err = custom.AfterRead(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ProfilesProfileWithBeforeRead called before DefaultReadProfile in the default Read handler
type ProfilesProfileWithBeforeRead interface {
	BeforeRead(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// ProfilesProfileWithAfterRead called before DefaultReadProfile in the default Read handler
type ProfilesProfileWithAfterRead interface {
	AfterRead(context.Context, *ReadProfileResponse, *gorm1.DB) error
}

// Update ...
func (m *ProfilesDefaultServer) Update(ctx context.Context, in *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	var err error
	var res *Profile
	db := m.DB
	if custom, ok := interface{}(in).(ProfilesProfileWithBeforeUpdate); ok {
		var err error
		if db, err = custom.BeforeUpdate(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err = DefaultStrictUpdateProfile(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	out := &UpdateProfileResponse{Result: res}
	if custom, ok := interface{}(in).(ProfilesProfileWithAfterUpdate); ok {
		var err error
		if err = custom.AfterUpdate(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ProfilesProfileWithBeforeUpdate called before DefaultUpdateProfile in the default Update handler
type ProfilesProfileWithBeforeUpdate interface {
	BeforeUpdate(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// ProfilesProfileWithAfterUpdate called before DefaultUpdateProfile in the default Update handler
type ProfilesProfileWithAfterUpdate interface {
	AfterUpdate(context.Context, *UpdateProfileResponse, *gorm1.DB) error
}

// Delete ...
func (m *ProfilesDefaultServer) Delete(ctx context.Context, in *DeleteProfileRequest) (*DeleteProfileResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(ProfilesProfileWithBeforeDelete); ok {
		var err error
		if db, err = custom.BeforeDelete(ctx, db); err != nil {
			return nil, err
		}
	}
	err := DefaultDeleteProfile(ctx, &Profile{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &DeleteProfileResponse{}
	if custom, ok := interface{}(in).(ProfilesProfileWithAfterDelete); ok {
		var err error
		if err = custom.AfterDelete(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ProfilesProfileWithBeforeDelete called before DefaultDeleteProfile in the default Delete handler
type ProfilesProfileWithBeforeDelete interface {
	BeforeDelete(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// ProfilesProfileWithAfterDelete called before DefaultDeleteProfile in the default Delete handler
type ProfilesProfileWithAfterDelete interface {
	AfterDelete(context.Context, *DeleteProfileResponse, *gorm1.DB) error
}

// List ...
func (m *ProfilesDefaultServer) List(ctx context.Context, in *ListProfileRequest) (*ListProfileResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(ProfilesProfileWithBeforeList); ok {
		var err error
		if db, err = custom.BeforeList(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultListProfile(ctx, db, in.Filter, in.OrderBy, in.Paging, in.Fields)
	if err != nil {
		return nil, err
	}
	out := &ListProfileResponse{Results: res}
	if custom, ok := interface{}(in).(ProfilesProfileWithAfterList); ok {
		var err error
		if err = custom.AfterList(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ProfilesProfileWithBeforeList called before DefaultListProfile in the default List handler
type ProfilesProfileWithBeforeList interface {
	BeforeList(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// ProfilesProfileWithAfterList called before DefaultListProfile in the default List handler
type ProfilesProfileWithAfterList interface {
	AfterList(context.Context, *ListProfileResponse, *gorm1.DB) error
}
