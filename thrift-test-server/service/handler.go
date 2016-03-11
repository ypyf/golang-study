package service

import (
	"crypto/md5"
	"fmt"
	"wifi/session/rpc"

	"encoding/json"
	"time"

	"github.com/Sirupsen/logrus"
	"gopkg.in/redis.v3"
)

const (
	TEST_SESSION = "s-test-session"
)

type SessionData struct {
	*rpc.SessionContext
	Timestamp time.Time `json:"timestamp"`
}

type SessionHandler struct {
	*redis.Client
}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func NewSessionHandler() *SessionHandler {
	return &SessionHandler{redis.NewClient(&redis.Options{
		Addr:     "10.10.149.62:6379",
		Password: "",
		DB:       0, // use default DB
	})}
}

func (s *SessionHandler) CreateSession(ctx *rpc.SessionContext) (sid string, err error) {
	logrus.Debug("调用 CreateSession()")

	// Algorithm: 's-' . substr($token, 0,5) . ':' . md5($token . $uuid)
	md5sum := md5.Sum([]byte(ctx.Token + ctx.UUID))
	sid = fmt.Sprintf("s-%s:%x", ctx.Token[:5], md5sum)
	logrus.Infof("生成会话ID %s", sid)
	sd, err := json.Marshal(&SessionData{ctx, time.Now()})
	if err != nil {
		logrus.Error(err)
		return "", rpc.NewSerializeError()
	} else {
		if s.Set(sid, sd, time.Hour*24).Err() != nil {
			logrus.Error(err)
			return "", rpc.NewIOError()
		}
	}
	// TODO 持久化到MYSQL
	return
}

func (s *SessionHandler) GetSessionContext(sid string) (ctx *rpc.SessionContext, err error) {
	logrus.Debug("调用 GetSessionContext()")
	value, err := s.Get(sid).Result()
	if err != nil {
		logrus.Error(err)
		return nil, rpc.NewInvalidSession()
	}
	var sessionData SessionData
	if err = json.Unmarshal([]byte(value), &sessionData); err != nil {
		logrus.Error(err)
		return nil, rpc.NewUnserializeError()
	}

	if time.Since(sessionData.Timestamp) > time.Hour*24 {
		logrus.Infof("删除过期会话 %s", sid)
		if _, err = s.Del(sid).Result(); err != nil {
			logrus.Error(err)
		}
		// TODO 持久化到MYSQL
		return nil, rpc.NewInvalidSession()
	}
	return sessionData.SessionContext, nil
}
