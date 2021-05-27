package mycache

import (
	"fmt"
	"time"
)

const (
	CacheExpired = 1 + iota
	CacheNonExisted
)

type CacheError struct {
	Msg        string
	ErrCode    int
	LastAccess time.Time
}

func (e *CacheError) Error() string {
	return fmt.Sprintf("%d: %v", e.ErrCode, e.Msg)
}

func NewErrorNonExisted(errMsg string) *CacheError {
	return &CacheError{
		Msg:     errMsg,
		ErrCode: CacheNonExisted,
	}
}

func NewErrorExpired(lastAccess time.Time, errMsg string) *CacheError {
	return &CacheError{
		Msg:        errMsg,
		ErrCode:    CacheExpired,
		LastAccess: lastAccess,
	}
}

func IsExpired(err error) bool {
	if e, ok := err.(*CacheError); ok {
		return e.ErrCode == CacheExpired
	}
	return false
}

func IsNonExisted(err error) bool {
	if e, ok := err.(*CacheError); ok {
		return e.ErrCode == CacheNonExisted
	}
	return false
}
