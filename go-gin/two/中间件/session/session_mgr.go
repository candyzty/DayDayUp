/**
* @Author: redgr
* @Description:
* @File:  session_mgr
* @Version: 1.0.0
* @Date: 2020/7/25 17:08
 */

package session

type SessionMgr interface {
	//	 初始化
	Init(addr string, options ...string) (err error)
	CreateSession() (Session Session, err error)
	Get(sessionId string) (Session Session, err error)
}
