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
	"fmt"
	"time"
)

type UserResource resource

type User struct {
	Embedded []string `json:"_embedded, omitempty"`
	Links []string `json:"_links, omitempty"`
	Activated time.Time `json:"activated, omitempty"`
	Created time.Time `json:"created, omitempty"`
	Credentials UserCredentials `json:"credentials, omitempty"`
	Id string `json:"id, omitempty"`
	LastLogin time.Time `json:"lastLogin, omitempty"`
	LastUpdated time.Time `json:"lastUpdated, omitempty"`
	PasswordChanged time.Time `json:"passwordChanged, omitempty"`
	Profile UserProfile `json:"profile, omitempty"`
	Status string `json:"status, omitempty"`
	StatusChanged time.Time `json:"statusChanged, omitempty"`
	TransitioningToStatus string `json:"transitioningToStatus, omitempty"`
}

func (m *User) WithCredentials(v UserCredentials) *User {
	m.Credentials = v
	return m
}

func (m *User) WithProfile(v UserProfile) *User {
	m.Profile = v
	return m
}

func (m *UserResource) CreateUser(body User)  *User {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/users", bytes.NewReader(iobytes))
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := User{}
	
	json.Unmarshal(resp, &r)
	
	return r
}

func (m *UserResource) GetUser(userId string)  *User {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := User{}
	
	json.Unmarshal(resp, &r)
	
	return r
}

func (m *UserResource) UpdateUser(userId string, body User)  *User {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	resp, err := m.client.requestExecutor.Put("/api/v1/users/"+userId+"", bytes.NewReader(iobytes))
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := User{}
	
	json.Unmarshal(resp, &r)
	
	return r
}

func (m *UserResource) DeactivateOrDeleteUser(userId string)  {
	_, err := m.client.requestExecutor.Delete("/api/v1/users/"+userId+"")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}


func (m *UserResource) EndAllUserSessions(userId string)  {
	_, err := m.client.requestExecutor.Delete("/api/v1/users/"+userId+"/sessions")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}
func (m *UserResource) ListAppLinks(userId string)  *AppLink {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/appLinks")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := AppLink{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) ChangePassword(userId string, body ChangePasswordRequest)  *UserCredentials {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/credentials/change_password", bytes.NewReader(iobytes))
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := UserCredentials{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) ChangeRecoveryQuestion(userId string, body UserCredentials)  *UserCredentials {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/credentials/change_recovery_question", bytes.NewReader(iobytes))
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := UserCredentials{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) ForgotPassword(userId string, body UserCredentials)  *ForgotPasswordResponse {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/credentials/forgot_password", bytes.NewReader(iobytes))
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := ForgotPasswordResponse{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) ListAssignedRoles(userId string)  *Role {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/roles")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := Role{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) AddRoleToUser(userId string, body Role)  *Role {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/roles", bytes.NewReader(iobytes))
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := Role{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) RemoveRoleFromUser(userId string, roleId string)  {
	_, err := m.client.requestExecutor.Delete("/api/v1/users/"+userId+"/roles/"+roleId+"")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}
func (m *UserResource) ListGroupTargetsForRole(userId string, roleId string)  *Group {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/roles/"+roleId+"/targets/groups")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := Group{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) RemoveGroupTargetFromRole(userId string, roleId string, groupId string)  {
	_, err := m.client.requestExecutor.Delete("/api/v1/users/"+userId+"/roles/"+roleId+"/targets/groups/"+groupId+"")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}
func (m *UserResource) AddGroupTargetToRole(userId string, roleId string, groupId string)  {
	_, err := m.client.requestExecutor.Put("/api/v1/users/"+userId+"/roles/"+roleId+"/targets/groups/"+groupId+"", nil)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}
func (m *UserResource) ListUserGroups(userId string)  *Group {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/groups")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := Group{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) ActivateUser(userId string)  *UserActivationToken {
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/activate", nil)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := UserActivationToken{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) DeactivateUser(userId string)  {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/deactivate", nil)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}
func (m *UserResource) SuspendUser(userId string)  {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/suspend", nil)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}
func (m *UserResource) UnsuspendUser(userId string)  {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/unsuspend", nil)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}
func (m *UserResource) ResetPassword(userId string)  *ResetPasswordToken {
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/reset_password", nil)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := ResetPasswordToken{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) ExpirePassword(userId string)  *TempPassword {
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/expire_password", nil)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := TempPassword{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) UnlockUser(userId string)  {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/unlock", nil)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}
func (m *UserResource) ResetAllFactors(userId string)  {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/reset_factors", nil)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}
func (m *UserResource) AddUserToGroup(groupId string, userId string)  {
	_, err := m.client.requestExecutor.Put("/api/v1/groups/"+groupId+"/users/"+userId+"", nil)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}
func (m *UserResource) AddFactor(userId string, body Factor)  *Factor {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/factors", bytes.NewReader(iobytes))
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := Factor{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) ListSupportedFactors(userId string)  *Factor {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/factors/catalog")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := Factor{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) ListFactors(userId string)  *Factor {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/factors")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := Factor{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) ListSupportedSecurityQuestions(userId string)  *SecurityQuestion {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/factors/questions")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := SecurityQuestion{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
func (m *UserResource) GetFactor(userId string, factorId string)  *Factor {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/factors/"+factorId+"")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	
	r := Factor{}
	
	json.Unmarshal(resp, &r)
	
	return r
}
