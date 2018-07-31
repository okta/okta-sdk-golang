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
	"bytes"
	"encoding/json"
	"time"

	"github.com/okta/okta-sdk-golang/okta/query"
)

type UserResource resource

type User struct {
	Embedded              []string         `json:"_embedded,omitempty"`
	Links                 []string         `json:"_links,omitempty"`
	Activated             *time.Time       `json:"activated,omitempty"`
	Created               *time.Time       `json:"created,omitempty"`
	Credentials           *UserCredentials `json:"credentials,omitempty"`
	Id                    string           `json:"id,omitempty"`
	LastLogin             *time.Time       `json:"lastLogin,omitempty"`
	LastUpdated           *time.Time       `json:"lastUpdated,omitempty"`
	PasswordChanged       *time.Time       `json:"passwordChanged,omitempty"`
	Profile               *UserProfile     `json:"profile,omitempty"`
	Status                string           `json:"status,omitempty"`
	StatusChanged         *time.Time       `json:"statusChanged,omitempty"`
	TransitioningToStatus string           `json:"transitioningToStatus,omitempty"`
}

func (m *User) WithCredentials(v *UserCredentials) *User {
	m.Credentials = v
	return m
}

func (m *User) WithProfile(v *UserProfile) *User {
	m.Profile = v
	return m
}

func (m *UserResource) CreateUser(body User, qp *query.Params) (*User, error) {
	iobytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/users", bytes.NewReader(iobytes), qp)
	if err != nil {
		return nil, err
	}

	r := User{}

	json.Unmarshal(resp, &r)

	return &r, nil
}

func (m *UserResource) GetUser(userId string, qp *query.Params) (*User, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"", qp)
	if err != nil {
		return nil, err
	}

	r := User{}

	json.Unmarshal(resp, &r)

	return &r, nil
}

func (m *UserResource) UpdateUser(userId string, body User, qp *query.Params) (*User, error) {
	iobytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.requestExecutor.Put("/api/v1/users/"+userId+"", bytes.NewReader(iobytes), qp)
	if err != nil {
		return nil, err
	}

	r := User{}

	json.Unmarshal(resp, &r)

	return &r, nil
}

func (m *UserResource) DeactivateOrDeleteUser(userId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Delete("/api/v1/users/"+userId+"", qp)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserResource) EndAllUserSessions(userId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Delete("/api/v1/users/"+userId+"/sessions", qp)
	if err != nil {
		return err
	}
	return nil
}
func (m *UserResource) ListAppLinks(userId string, qp *query.Params) (*AppLink, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/appLinks", qp)
	if err != nil {
		return nil, err
	}

	r := AppLink{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) ChangePassword(userId string, body ChangePasswordRequest, qp *query.Params) (*UserCredentials, error) {
	iobytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/credentials/change_password", bytes.NewReader(iobytes), qp)
	if err != nil {
		return nil, err
	}

	r := UserCredentials{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) ChangeRecoveryQuestion(userId string, body UserCredentials, qp *query.Params) (*UserCredentials, error) {
	iobytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/credentials/change_recovery_question", bytes.NewReader(iobytes), qp)
	if err != nil {
		return nil, err
	}

	r := UserCredentials{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) ForgotPassword(userId string, body UserCredentials, qp *query.Params) (*ForgotPasswordResponse, error) {
	iobytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/credentials/forgot_password", bytes.NewReader(iobytes), qp)
	if err != nil {
		return nil, err
	}

	r := ForgotPasswordResponse{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) ListAssignedRoles(userId string, qp *query.Params) (*Role, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/roles", qp)
	if err != nil {
		return nil, err
	}

	r := Role{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) AddRoleToUser(userId string, body Role, qp *query.Params) (*Role, error) {
	iobytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/roles", bytes.NewReader(iobytes), qp)
	if err != nil {
		return nil, err
	}

	r := Role{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) RemoveRoleFromUser(userId string, roleId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Delete("/api/v1/users/"+userId+"/roles/"+roleId+"", qp)
	if err != nil {
		return err
	}
	return nil
}
func (m *UserResource) ListGroupTargetsForRole(userId string, roleId string, qp *query.Params) (*Group, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/roles/"+roleId+"/targets/groups", qp)
	if err != nil {
		return nil, err
	}

	r := Group{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) RemoveGroupTargetFromRole(userId string, roleId string, groupId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Delete("/api/v1/users/"+userId+"/roles/"+roleId+"/targets/groups/"+groupId+"", qp)
	if err != nil {
		return err
	}
	return nil
}
func (m *UserResource) AddGroupTargetToRole(userId string, roleId string, groupId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Put("/api/v1/users/"+userId+"/roles/"+roleId+"/targets/groups/"+groupId+"", nil, qp)
	if err != nil {
		return err
	}
	return nil
}
func (m *UserResource) ListUserGroups(userId string, qp *query.Params) (*Group, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/groups", qp)
	if err != nil {
		return nil, err
	}

	r := Group{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) ActivateUser(userId string, qp *query.Params) (*UserActivationToken, error) {
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/activate", nil, qp)
	if err != nil {
		return nil, err
	}

	r := UserActivationToken{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) DeactivateUser(userId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/deactivate", nil, qp)
	if err != nil {
		return err
	}
	return nil
}
func (m *UserResource) SuspendUser(userId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/suspend", nil, qp)
	if err != nil {
		return err
	}
	return nil
}
func (m *UserResource) UnsuspendUser(userId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/unsuspend", nil, qp)
	if err != nil {
		return err
	}
	return nil
}
func (m *UserResource) ResetPassword(userId string, qp *query.Params) (*ResetPasswordToken, error) {
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/reset_password", nil, qp)
	if err != nil {
		return nil, err
	}

	r := ResetPasswordToken{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) ExpirePassword(userId string, qp *query.Params) (*TempPassword, error) {
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/expire_password", nil, qp)
	if err != nil {
		return nil, err
	}

	r := TempPassword{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) UnlockUser(userId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/unlock", nil, qp)
	if err != nil {
		return err
	}
	return nil
}
func (m *UserResource) ResetAllFactors(userId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/reset_factors", nil, qp)
	if err != nil {
		return err
	}
	return nil
}
func (m *UserResource) AddUserToGroup(groupId string, userId string, qp *query.Params) error {
	_, err := m.client.requestExecutor.Put("/api/v1/groups/"+groupId+"/users/"+userId+"", nil, qp)
	if err != nil {
		return err
	}
	return nil
}
func (m *UserResource) AddFactor(userId string, body Factor, qp *query.Params) (*Factor, error) {
	iobytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/factors", bytes.NewReader(iobytes), qp)
	if err != nil {
		return nil, err
	}

	r := Factor{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) ListSupportedFactors(userId string, qp *query.Params) (*Factor, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/factors/catalog", qp)
	if err != nil {
		return nil, err
	}

	r := Factor{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) ListFactors(userId string, qp *query.Params) (*Factor, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/factors", qp)
	if err != nil {
		return nil, err
	}

	r := Factor{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) ListSupportedSecurityQuestions(userId string, qp *query.Params) (*SecurityQuestion, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/factors/questions", qp)
	if err != nil {
		return nil, err
	}

	r := SecurityQuestion{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
func (m *UserResource) GetFactor(userId string, factorId string, qp *query.Params) (*Factor, error) {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/factors/"+factorId+"", qp)
	if err != nil {
		return nil, err
	}

	r := Factor{}

	json.Unmarshal(resp, &r)

	return &r, nil
}
