// Package types provides data types for persisting objects in the databases.
package types

import (
	"database/sql/driver"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"sort"
	"strings"
	"time"
)

// StoreError satisfies Error interface but allows constant values for
// direct comparison.
type StoreError string

// Error is required by error interface.
func (s StoreError) Error() string {
	return string(s)
}

const (
	// ErrInternal means DB or other internal failure.
	ErrInternal = StoreError("internal")
)

// Uid is a database-specific record id, suitable to be used as a primary key.
type Uid uint64

// ZeroUid is a constant representing uninitialized Uid.
const ZeroUid Uid = 0

// NullValue is a Unicode DEL character which indicated that the value is being deleted.
const NullValue = "\u2421"

