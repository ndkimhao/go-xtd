package atomic

import (
	"time"
	"unsafe"

	"go.uber.org/atomic"
)

type Bool = atomic.Bool
type Duration = atomic.Duration
type Error = atomic.Error
type Float64 = atomic.Float64
type Int32 = atomic.Int32
type Int64 = atomic.Int64
type String = atomic.String
type Time = atomic.Time
type Uint32 = atomic.Uint32
type Uint64 = atomic.Uint64
type Uintptr = atomic.Uintptr
type UnsafePointer = atomic.UnsafePointer

func NewBool(val bool) *Bool {
	return atomic.NewBool(val)
}

func NewDuration(val time.Duration) *Duration {
	return atomic.NewDuration(val)
}

func NewError(val error) *Error {
	return atomic.NewError(val)
}

func NewFloat64(val float64) *Float64 {
	return atomic.NewFloat64(val)
}

func NewInt32(val int32) *Int32 {
	return atomic.NewInt32(val)
}

func NewInt64(val int64) *Int64 {
	return atomic.NewInt64(val)
}

func NewString(val string) *String {
	return atomic.NewString(val)
}

func NewTime(val time.Time) *Time {
	return atomic.NewTime(val)
}

func NewUint32(val uint32) *Uint32 {
	return atomic.NewUint32(val)
}

func NewUint64(val uint64) *Uint64 {
	return atomic.NewUint64(val)
}

func NewUintptr(val uintptr) *Uintptr {
	return atomic.NewUintptr(val)
}

func NewUnsafePointer(val unsafe.Pointer) *UnsafePointer {
	return atomic.NewUnsafePointer(val)
}
