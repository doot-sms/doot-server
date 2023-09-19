// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type UssrRequestor string

const (
	UssrRequestorUser   UssrRequestor = "user"
	UssrRequestorSender UssrRequestor = "sender"
)

func (e *UssrRequestor) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UssrRequestor(s)
	case string:
		*e = UssrRequestor(s)
	default:
		return fmt.Errorf("unsupported scan type for UssrRequestor: %T", src)
	}
	return nil
}

type NullUssrRequestor struct {
	UssrRequestor UssrRequestor
	Valid         bool // Valid is true if UssrRequestor is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUssrRequestor) Scan(value interface{}) error {
	if value == nil {
		ns.UssrRequestor, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UssrRequestor.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUssrRequestor) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UssrRequestor), nil
}

type UssrStatus string

const (
	UssrStatusRequested UssrStatus = "requested"
	UssrStatusAccepted  UssrStatus = "accepted"
	UssrStatusRejected  UssrStatus = "rejected"
)

func (e *UssrStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UssrStatus(s)
	case string:
		*e = UssrStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UssrStatus: %T", src)
	}
	return nil
}

type NullUssrStatus struct {
	UssrStatus UssrStatus
	Valid      bool // Valid is true if UssrStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUssrStatus) Scan(value interface{}) error {
	if value == nil {
		ns.UssrStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UssrStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUssrStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UssrStatus), nil
}

type Batch struct {
	ID        int32
	QueuedAt  sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Message struct {
	ID        int32
	To        string
	Content   string
	BatchID   sql.NullInt32
	SentAt    sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int32
	SenderID  int32
}

type RefreshToken struct {
	ID        int32
	UserID    int32
	UserAgent string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Sender struct {
	ID        int32
	UserID    int32
	DeviceID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID        int32
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserApiKey struct {
	ApiKey       string
	UserID       int32
	ApiSecret    string
	ExpiresAfter sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserSender struct {
	UserID    int32
	SenderID  int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserSenderReq struct {
	ID        int32
	UserID    int32
	SenderID  int32
	Requestor UssrRequestor
	Status    UssrStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
