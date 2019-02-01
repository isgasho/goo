//session.go
package goo

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

//set cookie
func SetCookie(w http.ResponseWriter, cookie *http.Cookie) {
	http.SetCookie(w, cookie)
}

//get cookie
func GetCookie(r *http.Request, cookieName string) (*http.Cookie, error) {
	return r.Cookie(cookieName)
}

// use mysql to local session object.

// Session interface
type Session interface {
	Get(key interface{}) interface{}
	Set(key, value interface{}) error
	Delete(key interface{}) error
	SessionID() string
}

//the main manager about the session system
type SessionStore struct {
	value    map[string]interface{} `json:"value"`
	sid      string                 `json:"sid"`
	lastTime time.Time              `json:"last_time"`
}

// the content about the system.
func (s *SessionStore) Get(key interface{}) interface{} {
	keyString := key.(string)
	if value, ok := s.value[keyString]; ok {
		return value
	} else {
		return ""
	}
}
func (s *SessionStore) Set(key, value interface{}) error {
	if keyString, ok := key.(string); ok {
		s.value[keyString] = value
		return nil
	} else {
		return fmt.Errorf("error")
	}

}
func (s *SessionStore) Delete(key interface{}) error {
	if keyString, ok := key.(string); ok {
		delete(s.value, keyString)
		return nil
	} else {
		return fmt.Errorf("error")
	}

}
func (s *SessionStore) SessionID() string {
	return s.sid
}

// newSession new a session object.
func NewSession() Session {
	st := new(SessionStore)
	st.lastTime = time.Now()
	a := make([]byte, 64)
	rand.New(rand.NewSource(time.Now().UnixNano())).Read(a)
	st.sid = base64.URLEncoding.EncodeToString(a)
	st.value = make(map[string]interface{})
	fmt.Println(st)
	return st
}
