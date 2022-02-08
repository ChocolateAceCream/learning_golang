package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gin/services"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type Session struct {
	Cookie      string                 `json: "cookie"`
	ExpireTime  int64                  `json: "expire_time"`
	SessionList map[string]interface{} `json: "session_list"`
	Key         string                 `json:"-"`
	Lock        *sync.Mutex
}

var cookieName = "my_gin_cookie"
var expireTime = 3600
var secure = false
var httpOnly = true
var path = "/"

// middleware executed before actual logic of route handler function
func SessionHandler(c *gin.Context) {
	// c.Cookie() return the cookie value fetched with cookieName, then use that value to search redis
	cookie, err := c.Cookie(cookieName)

	// if session not expired in redis, continue, otherwise create new session
	if err == nil {
		sessionString, err := services.GetRedisClient().Get(context.TODO(), cookie).Result()

		// if session not expired in redis, store session in current context, otherwise create new session
		if err == nil {
			var session Session
			json.Unmarshal([]byte(sessionString), &session)
			c.Set("_session", session)
			return
		}
	}
	sessionKey := uuid.NewV4().String()
	domain := c.Request.Host[:strings.Index(c.Request.Host, ":")]
	c.SetCookie(cookieName, sessionKey, expireTime, path, domain, secure, httpOnly)
	session := Session{
		Cookie:      cookieName,
		ExpireTime:  time.Now().Unix() + int64(expireTime),
		SessionList: make(map[string]interface{}),
	}

	// store session in gin.Context
	fmt.Println("----SessionHandler-----set session.Cookie------------")
	c.Set("_session", session)
	jsonString, _ := json.Marshal(session)
	services.GetRedisClient().Set(c, sessionKey, jsonString, time.Duration(expireTime)*time.Second)
}

// first get cookie name from context, then use cookie name to get cookie uuid from context's cookie
func GetSession(c *gin.Context) *Session {
	cookie, ok := c.Get("_session")
	if !ok {
		// cannot retrive cookie from current context
		return nil
	}
	session, ok := cookie.(Session)
	if !ok {
		// if cookie is not of type Session
		return nil
	}
	sessionKey, err := c.Cookie(session.Cookie)
	if err != nil {
		// if cookie is not of type Session
		return nil
	}
	session.Key = sessionKey
	session.Lock = &sync.Mutex{}
	return &session
}

func (s *Session) Set(key string, val interface{}) error {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	sessionString, err := services.GetRedisClient().Get(context.TODO(), s.Key).Result()
	if err != nil {
		fmt.Println("----GetRedisClient error----", err)
		return err
	}
	var session Session
	err = json.Unmarshal([]byte(sessionString), &session)
	if err != nil {
		fmt.Println("----json.Unmarshal error----", err)
		return err
	}
	session.SessionList[key] = val

	sessionStringNew, err := json.Marshal(session)
	if err != nil {
		fmt.Println("----json.Marshal error----", err)
		return err
	}

	e := s.ExpireTime - time.Now().Unix()

	if e < 0 {
		return errors.New("session has expired")
	}

	services.GetRedisClient().Set(context.TODO(), s.Key, sessionStringNew, time.Duration(e)*time.Second)
	return nil
}

func (s *Session) Get(key string) (interface{}, error) {
	sessionString, err := services.GetRedisClient().Get(context.TODO(), s.Key).Result()
	if err != nil {
		fmt.Println("----GetRedisClient error----", err)
		return nil, err
	}
	var session Session
	err = json.Unmarshal([]byte(sessionString), &session)
	if err != nil {
		fmt.Println("----json.Unmarshal error----", err)
		return nil, err
	}
	if val, ok := session.SessionList[key]; ok {
		return val, nil
	}

	return nil, errors.New("not found key :" + key)
}

func (s *Session) Remove(key string) error {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	sessionString, err := services.GetRedisClient().Get(context.TODO(), s.Key).Result()
	if err != nil {
		fmt.Println("----GetRedisClient error----", err)
		return err
	}

	var session Session
	err = json.Unmarshal([]byte(sessionString), &session)
	if err != nil {
		fmt.Println("----json.Unmarshal error----", err)
		return err
	}
	delete(session.SessionList, key)
	sessionStringNew, err := json.Marshal(session)
	if err != nil {
		fmt.Println("----json.Marshal error----", err)
		return err
	}

	e := s.ExpireTime - time.Now().Unix()

	if e < 0 {
		return errors.New("session has expired")
	}

	services.GetRedisClient().Set(context.TODO(), s.Key, sessionStringNew, time.Duration(e)*time.Second)
	return nil

}
