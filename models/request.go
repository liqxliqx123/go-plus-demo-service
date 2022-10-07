package models

import (
	"github.com/segmentio/ksuid"
)

// RequestInfo api request info
type RequestInfo struct {
	// CustomerID gid
	CustomerID string
	UserID     string
	RequestID  string
	Vendor     string
	Lang       string
	Entity     string // siren 服务 entity 设置
}

// SetRequestID SetRequestID
func (r *RequestInfo) SetRequestID() {
	if r.RequestID == "" {
		r.RequestID = ksuid.New().String()
	}
}
