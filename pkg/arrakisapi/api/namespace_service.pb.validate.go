// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: REDACTEDapi/api/namespace_service.proto

package REDACTEDapi

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

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
)

// Validate checks the field values on ReadConfigRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ReadConfigRequest) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetNamespace()) > 128 {
		return ReadConfigRequestValidationError{
			field:  "Namespace",
			reason: "value length must be at most 128 bytes",
		}
	}

	if !_ReadConfigRequest_Namespace_Pattern.MatchString(m.GetNamespace()) {
		return ReadConfigRequestValidationError{
			field:  "Namespace",
			reason: "value does not match regex pattern \"^([a-z][a-z0-9_]{2,62}[a-z0-9]/)?[a-z][a-z0-9_]{2,62}[a-z0-9]$\"",
		}
	}

	if v, ok := interface{}(m.GetAtRevision()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadConfigRequestValidationError{
				field:  "AtRevision",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ReadConfigRequestValidationError is the validation error returned by
// ReadConfigRequest.Validate if the designated constraints aren't met.
type ReadConfigRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadConfigRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadConfigRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadConfigRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadConfigRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadConfigRequestValidationError) ErrorName() string {
	return "ReadConfigRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReadConfigRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadConfigRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadConfigRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadConfigRequestValidationError{}

var _ReadConfigRequest_Namespace_Pattern = regexp.MustCompile("^([a-z][a-z0-9_]{2,62}[a-z0-9]/)?[a-z][a-z0-9_]{2,62}[a-z0-9]$")

// Validate checks the field values on ReadConfigResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ReadConfigResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Namespace

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadConfigResponseValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRevision()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadConfigResponseValidationError{
				field:  "Revision",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ReadConfigResponseValidationError is the validation error returned by
// ReadConfigResponse.Validate if the designated constraints aren't met.
type ReadConfigResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadConfigResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadConfigResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadConfigResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadConfigResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadConfigResponseValidationError) ErrorName() string {
	return "ReadConfigResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ReadConfigResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadConfigResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadConfigResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadConfigResponseValidationError{}

// Validate checks the field values on WriteConfigRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *WriteConfigRequest) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetConfigs()) < 1 {
		return WriteConfigRequestValidationError{
			field:  "Configs",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetConfigs() {
		_, _ = idx, item

		if item == nil {
			return WriteConfigRequestValidationError{
				field:  fmt.Sprintf("Configs[%v]", idx),
				reason: "value is required",
			}
		}

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WriteConfigRequestValidationError{
					field:  fmt.Sprintf("Configs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// WriteConfigRequestValidationError is the validation error returned by
// WriteConfigRequest.Validate if the designated constraints aren't met.
type WriteConfigRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WriteConfigRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WriteConfigRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WriteConfigRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WriteConfigRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WriteConfigRequestValidationError) ErrorName() string {
	return "WriteConfigRequestValidationError"
}

// Error satisfies the builtin error interface
func (e WriteConfigRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWriteConfigRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WriteConfigRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WriteConfigRequestValidationError{}

// Validate checks the field values on WriteConfigResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *WriteConfigResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRevision()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WriteConfigResponseValidationError{
				field:  "Revision",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WriteConfigResponseValidationError is the validation error returned by
// WriteConfigResponse.Validate if the designated constraints aren't met.
type WriteConfigResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WriteConfigResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WriteConfigResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WriteConfigResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WriteConfigResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WriteConfigResponseValidationError) ErrorName() string {
	return "WriteConfigResponseValidationError"
}

// Error satisfies the builtin error interface
func (e WriteConfigResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWriteConfigResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WriteConfigResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WriteConfigResponseValidationError{}