/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
	"fmt"
	"time"
)

type SessionResource resource

type Session struct {
	Links                    interface{}                    `json:"_links,omitempty"`
	Amr                      []*SessionAuthenticationMethod `json:"amr,omitempty"`
	CreatedAt                *time.Time                     `json:"createdAt,omitempty"`
	ExpiresAt                *time.Time                     `json:"expiresAt,omitempty"`
	Id                       string                         `json:"id,omitempty"`
	Idp                      *SessionIdentityProvider       `json:"idp,omitempty"`
	LastFactorVerification   *time.Time                     `json:"lastFactorVerification,omitempty"`
	LastPasswordVerification *time.Time                     `json:"lastPasswordVerification,omitempty"`
	Login                    string                         `json:"login,omitempty"`
	Status                   string                         `json:"status,omitempty"`
	UserId                   string                         `json:"userId,omitempty"`
}

func (m *SessionResource) GetSession(sessionId string) (*Session, *Response, error) {
	url := fmt.Sprintf("/api/v1/sessions/%v", sessionId)
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var session *Session
	resp, err := m.client.requestExecutor.Do(req, &session)
	if err != nil {
		return nil, resp, err
	}
	return session, resp, nil
}
func (m *SessionResource) EndSession(sessionId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/sessions/%v", sessionId)
	req, err := m.client.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (m *SessionResource) CreateSession(body CreateSessionRequest) (*Session, *Response, error) {
	url := fmt.Sprintf("/api/v1/sessions")
	req, err := m.client.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var session *Session
	resp, err := m.client.requestExecutor.Do(req, &session)
	if err != nil {
		return nil, resp, err
	}
	return session, resp, nil
}
func (m *SessionResource) RefreshSession(sessionId string) (*Session, *Response, error) {
	url := fmt.Sprintf("/api/v1/sessions/%v/lifecycle/refresh", sessionId)
	req, err := m.client.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var session *Session
	resp, err := m.client.requestExecutor.Do(req, &session)
	if err != nil {
		return nil, resp, err
	}
	return session, resp, nil
}
