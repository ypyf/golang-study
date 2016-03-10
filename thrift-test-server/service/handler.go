package service

import "wifi/session/rpc" //导入Thrift生成的接口包

const (
	TEST_SESSION = "s-test-session"
)

type SessionHandler struct {
}

func NewSessionHandler() *SessionHandler {
	return new(SessionHandler)
}

func (s *SessionHandler) CreateSession(ctx *rpc.SessionContext) (sid string, err error) {
	sid = TEST_SESSION
	return
}

func (s *SessionHandler) GetSession(sid string) (ctx *rpc.SessionContext, err error) {
	if sid == TEST_SESSION {
		ctx = rpc.NewSessionContext()
	} else {
		ouch := rpc.NewInvalidSession()
		err = ouch
	}
	return
}
