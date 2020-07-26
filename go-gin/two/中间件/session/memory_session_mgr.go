/**
* @Author: redgr
* @Description:
* @File:  memory_session_mgr
* @Version: 1.0.0
* @Date: 2020/7/25 17:29
 */

package session

import (
	uuid "github.com/satori/go.uuid"
	"sync"
)

type MemorySessionMgr struct {
	rwlock     sync.RWMutex
	sessionMap map[string]Session
}

// 构造函数
func NewMemorySessionMgr() SessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return sr
}
func (m *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}
func (m *MemorySessionMgr) CreateSession() (session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	// 用uuid 做session id
	id := uuid.NewV4()
	// 转string
	sessionID := id.String()
	// 创建session
	session = NewMemorySession(sessionID)
	return MemorySession

}
func (m *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	return
}
