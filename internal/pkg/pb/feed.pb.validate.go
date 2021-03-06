// Code generated by protoc-gen-validate
// source: github.com/kodesmil/ks-model/feed.proto
// DO NOT EDIT!!!

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on FeedTag with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FeedTag) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FeedTagValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Key

	// no validation rules for Name

	for idx, item := range m.GetFeedArticles() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FeedTagValidationError{
					field:  fmt.Sprintf("FeedArticles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// FeedTagValidationError is the validation error returned by FeedTag.Validate
// if the designated constraints aren't met.
type FeedTagValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeedTagValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeedTagValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeedTagValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeedTagValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeedTagValidationError) ErrorName() string { return "FeedTagValidationError" }

// Error satisfies the builtin error interface
func (e FeedTagValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeedTag.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeedTagValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeedTagValidationError{}

// Validate checks the field values on FeedAuthor with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FeedAuthor) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FeedAuthorValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Name

	// no validation rules for ProfilePhotoUrl

	// no validation rules for Bio

	for idx, item := range m.GetFeedArticles() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FeedAuthorValidationError{
					field:  fmt.Sprintf("FeedArticles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// FeedAuthorValidationError is the validation error returned by
// FeedAuthor.Validate if the designated constraints aren't met.
type FeedAuthorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeedAuthorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeedAuthorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeedAuthorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeedAuthorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeedAuthorValidationError) ErrorName() string { return "FeedAuthorValidationError" }

// Error satisfies the builtin error interface
func (e FeedAuthorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeedAuthor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeedAuthorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeedAuthorValidationError{}

// Validate checks the field values on FeedArticleDetail with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *FeedArticleDetail) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FeedArticleDetailValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Content

	if v, ok := interface{}(m.GetFeedArticle()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FeedArticleDetailValidationError{
				field:  "FeedArticle",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FeedArticleDetailValidationError is the validation error returned by
// FeedArticleDetail.Validate if the designated constraints aren't met.
type FeedArticleDetailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeedArticleDetailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeedArticleDetailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeedArticleDetailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeedArticleDetailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeedArticleDetailValidationError) ErrorName() string {
	return "FeedArticleDetailValidationError"
}

// Error satisfies the builtin error interface
func (e FeedArticleDetailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeedArticleDetail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeedArticleDetailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeedArticleDetailValidationError{}

// Validate checks the field values on ReadFeedArticleDetailsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ReadFeedArticleDetailsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadFeedArticleDetailsRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ReadFeedArticleDetailsRequestValidationError is the validation error
// returned by ReadFeedArticleDetailsRequest.Validate if the designated
// constraints aren't met.
type ReadFeedArticleDetailsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadFeedArticleDetailsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadFeedArticleDetailsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadFeedArticleDetailsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadFeedArticleDetailsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadFeedArticleDetailsRequestValidationError) ErrorName() string {
	return "ReadFeedArticleDetailsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReadFeedArticleDetailsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadFeedArticleDetailsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadFeedArticleDetailsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadFeedArticleDetailsRequestValidationError{}

// Validate checks the field values on ReadFeedArticleDetailsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ReadFeedArticleDetailsResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadFeedArticleDetailsResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ReadFeedArticleDetailsResponseValidationError is the validation error
// returned by ReadFeedArticleDetailsResponse.Validate if the designated
// constraints aren't met.
type ReadFeedArticleDetailsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadFeedArticleDetailsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadFeedArticleDetailsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadFeedArticleDetailsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadFeedArticleDetailsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadFeedArticleDetailsResponseValidationError) ErrorName() string {
	return "ReadFeedArticleDetailsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ReadFeedArticleDetailsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadFeedArticleDetailsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadFeedArticleDetailsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadFeedArticleDetailsResponseValidationError{}

// Validate checks the field values on FeedArticle with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FeedArticle) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FeedArticleValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FeedArticleValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Title

	// no validation rules for Subtitle

	// no validation rules for CoverPictureUrl

	// no validation rules for Content

	for idx, item := range m.GetFeedTags() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FeedArticleValidationError{
					field:  fmt.Sprintf("FeedTags[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetFeedAuthor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FeedArticleValidationError{
				field:  "FeedAuthor",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FeedArticleValidationError is the validation error returned by
// FeedArticle.Validate if the designated constraints aren't met.
type FeedArticleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeedArticleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeedArticleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeedArticleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeedArticleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeedArticleValidationError) ErrorName() string { return "FeedArticleValidationError" }

// Error satisfies the builtin error interface
func (e FeedArticleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeedArticle.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeedArticleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeedArticleValidationError{}

// Validate checks the field values on ListFeedArticleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListFeedArticleRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListFeedArticleRequestValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetOrderBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListFeedArticleRequestValidationError{
				field:  "OrderBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFields()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListFeedArticleRequestValidationError{
				field:  "Fields",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPaging()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListFeedArticleRequestValidationError{
				field:  "Paging",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListFeedArticleRequestValidationError is the validation error returned by
// ListFeedArticleRequest.Validate if the designated constraints aren't met.
type ListFeedArticleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFeedArticleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFeedArticleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFeedArticleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFeedArticleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFeedArticleRequestValidationError) ErrorName() string {
	return "ListFeedArticleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListFeedArticleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFeedArticleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFeedArticleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFeedArticleRequestValidationError{}

// Validate checks the field values on ListFeedArticleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListFeedArticleResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListFeedArticleResponseValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListFeedArticleResponseValidationError is the validation error returned by
// ListFeedArticleResponse.Validate if the designated constraints aren't met.
type ListFeedArticleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFeedArticleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFeedArticleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFeedArticleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFeedArticleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFeedArticleResponseValidationError) ErrorName() string {
	return "ListFeedArticleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListFeedArticleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFeedArticleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFeedArticleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFeedArticleResponseValidationError{}
