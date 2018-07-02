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
	"fmt"
	"encoding/json"
	"time"
)

type AppLink struct {
	AppAssignmentId string `json:"appAssignmentId, omitempty"`
	AppInstanceId string `json:"appInstanceId, omitempty"`
	AppName string `json:"appName, omitempty"`
	CredentialsSetup bool `json:"credentialsSetup, omitempty"`
	Hidden bool `json:"hidden, omitempty"`
	Id string `json:"id, omitempty"`
	Label string `json:"label, omitempty"`
	LinkUrl string `json:"linkUrl, omitempty"`
	LogoUrl string `json:"logoUrl, omitempty"`
	SortOrder int64 `json:"sortOrder, omitempty"`
}



type AuthenticationProvider struct {
	Name string `json:"name, omitempty"`
	Type string `json:"type, omitempty"`
}

func (m *AuthenticationProvider) WithName(v string) *AuthenticationProvider {
	m.Name = v
	return m
}
func (m *AuthenticationProvider) WithType(v string) *AuthenticationProvider {
	m.Type = v
	return m
}


type AuthenticationProviderType struct {
}



type ChangePasswordRequest struct {
	NewPassword PasswordCredential `json:"newPassword, omitempty"`
	OldPassword PasswordCredential `json:"oldPassword, omitempty"`
}

func (m *ChangePasswordRequest) WithNewPassword(v PasswordCredential) *ChangePasswordRequest {
	m.NewPassword = v
	return m
}
func (m *ChangePasswordRequest) WithOldPassword(v PasswordCredential) *ChangePasswordRequest {
	m.OldPassword = v
	return m
}


type ForgotPasswordResponse struct {
	ResetPasswordUrl string `json:"resetPasswordUrl, omitempty"`
}



type PasswordCredential struct {
	Value string `json:"value, omitempty"`
}

func (m *PasswordCredential) WithValue(v string) *PasswordCredential {
	m.Value = v
	return m
}


type RecoveryQuestionCredential struct {
	Answer string `json:"answer, omitempty"`
	Question string `json:"question, omitempty"`
}

func (m *RecoveryQuestionCredential) WithAnswer(v string) *RecoveryQuestionCredential {
	m.Answer = v
	return m
}
func (m *RecoveryQuestionCredential) WithQuestion(v string) *RecoveryQuestionCredential {
	m.Question = v
	return m
}


type ResetPasswordToken struct {
	ResetPasswordUrl string `json:"resetPasswordUrl, omitempty"`
}



type Role struct {
	Embedded map[string]interface{} `json:"_embedded, omitempty"`
	Created time.Time `json:"created, omitempty"`
	Description string `json:"description, omitempty"`
	Id string `json:"id, omitempty"`
	Label string `json:"label, omitempty"`
	LastUpdated time.Time `json:"lastUpdated, omitempty"`
	Status string `json:"status, omitempty"`
	Type string `json:"type, omitempty"`
}

func (m *Role) WithDescription(v string) *Role {
	m.Description = v
	return m
}


type RoleStatus struct {
}



type TempPassword struct {
	TempPassword string `json:"tempPassword, omitempty"`
}


type UserResource resource

type User struct {
	Embedded map[string]interface{} `json:"_embedded, omitempty"`
	Links map[string]interface{} `json:"_links, omitempty"`
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

func (m *UserResource) Endallsessions(userId string)  {
	_, err := m.client.requestExecutor.Delete("/api/v1/users/"+userId+"/sessions")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}

func (m *UserResource) Listapplinks(userId string) AppLink {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/appLinks")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := AppLink{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *UserResource) Changepassword(userId string, body ChangePasswordRequest) UserCredentials {
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

func (m *UserResource) Changerecoveryquestion(userId string, body UserCredentials) UserCredentials {
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

func (m *UserResource) Forgotpassword(userId string, body UserCredentials) ForgotPasswordResponse {
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

func (m *UserResource) Listroles(userId string) Role {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/roles")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := Role{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *UserResource) Addrole(userId string, body Role) Role {
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

func (m *UserResource) Removerole(userId string, roleId string)  {
	_, err := m.client.requestExecutor.Delete("/api/v1/users/"+userId+"/roles/"+roleId+"")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}

func (m *UserResource) Listgrouptargetsforrole(userId string, roleId string) Group {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/roles/"+roleId+"/targets/groups")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := Group{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *UserResource) Removegrouptargetfromrole(userId string, roleId string, groupId string)  {
	_, err := m.client.requestExecutor.Delete("/api/v1/users/"+userId+"/roles/"+roleId+"/targets/groups/"+groupId+"")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
}

func (m *UserResource) Addgrouptargettorole(userId string, roleId string, groupId string)  {
	_, err := m.client.requestExecutor.Put("/api/v1/users/"+userId+"/roles/"+roleId+"/targets/groups/"+groupId+"", nil)
		if err != nil  {
		fmt.Printf("error: %v", err)
	}
}

func (m *UserResource) Listgroups(userId string) Group {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/groups")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := Group{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *UserResource) Activate(userId string) UserActivationToken {
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/activate", nil)
		if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := UserActivationToken{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *UserResource) Deactivate(userId string)  {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/deactivate", nil)
		if err != nil  {
		fmt.Printf("error: %v", err)
	}
}

func (m *UserResource) Suspend(userId string)  {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/suspend", nil)
		if err != nil  {
		fmt.Printf("error: %v", err)
	}
}

func (m *UserResource) Unsuspend(userId string)  {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/unsuspend", nil)
		if err != nil  {
		fmt.Printf("error: %v", err)
	}
}

func (m *UserResource) Resetpassword(userId string) ResetPasswordToken {
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/reset_password", nil)
		if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := ResetPasswordToken{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *UserResource) Expirepassword(userId string) TempPassword {
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/expire_password", nil)
		if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := TempPassword{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *UserResource) Unlock(userId string)  {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/unlock", nil)
		if err != nil  {
		fmt.Printf("error: %v", err)
	}
}

func (m *UserResource) Resetfactors(userId string)  {
	_, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/lifecycle/reset_factors", nil)
		if err != nil  {
		fmt.Printf("error: %v", err)
	}
}

func (m *UserResource) Addtogroup(groupId string, userId string)  {
	_, err := m.client.requestExecutor.Put("/api/v1/groups/"+groupId+"/users/"+userId+"", nil)
		if err != nil  {
		fmt.Printf("error: %v", err)
	}
}

func (m *UserResource) Addfactor(userId string, body Factor) Factor {
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

func (m *UserResource) Listsupportedfactors(userId string) Factor {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/factors/catalog")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := Factor{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *UserResource) Listfactors(userId string) Factor {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/factors")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := Factor{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *UserResource) Listsupportedsecurityquestions(userId string) SecurityQuestion {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/factors/questions")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := SecurityQuestion{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *UserResource) Getfactor(userId string, factorId string) Factor {
	resp, err := m.client.requestExecutor.Get("/api/v1/users/"+userId+"/factors/"+factorId+"")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := Factor{}

	json.Unmarshal(resp, &r)

	return r
}


type UserActivationToken struct {
	ActivationToken string `json:"activationToken, omitempty"`
	ActivationUrl string `json:"activationUrl, omitempty"`
}



type UserCredentials struct {
	Password PasswordCredential `json:"password, omitempty"`
	Provider AuthenticationProvider `json:"provider, omitempty"`
	RecoveryQuestion RecoveryQuestionCredential `json:"recovery_question, omitempty"`
}

func (m *UserCredentials) WithPassword(v PasswordCredential) *UserCredentials {
	m.Password = v
	return m
}
func (m *UserCredentials) WithProvider(v AuthenticationProvider) *UserCredentials {
	m.Provider = v
	return m
}
func (m *UserCredentials) WithRecoveryQuestion(v RecoveryQuestionCredential) *UserCredentials {
	m.RecoveryQuestion = v
	return m
}


type UserProfile struct {
	Email string `json:"email, omitempty"`
	FirstName string `json:"firstName, omitempty"`
	LastName string `json:"lastName, omitempty"`
	Login string `json:"login, omitempty"`
	MobilePhone string `json:"mobilePhone, omitempty"`
	SecondEmail string `json:"secondEmail, omitempty"`
}

func (m *UserProfile) WithEmail(v string) *UserProfile {
	m.Email = v
	return m
}
func (m *UserProfile) WithFirstName(v string) *UserProfile {
	m.FirstName = v
	return m
}
func (m *UserProfile) WithLastName(v string) *UserProfile {
	m.LastName = v
	return m
}
func (m *UserProfile) WithLogin(v string) *UserProfile {
	m.Login = v
	return m
}
func (m *UserProfile) WithMobilePhone(v string) *UserProfile {
	m.MobilePhone = v
	return m
}
func (m *UserProfile) WithSecondEmail(v string) *UserProfile {
	m.SecondEmail = v
	return m
}


type UserStatus struct {
}





