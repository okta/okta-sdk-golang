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

type CallFactor struct {
	Profile CallFactorProfile `json:"profile, omitempty"`
}

func (m *CallFactor) WithProfile(v CallFactorProfile) *CallFactor {
	m.Profile = v
	return m
}


type CallFactorProfile struct {
	PhoneExtension string `json:"phoneExtension, omitempty"`
	PhoneNumber string `json:"phoneNumber, omitempty"`
}

func (m *CallFactorProfile) WithPhoneExtension(v string) *CallFactorProfile {
	m.PhoneExtension = v
	return m
}
func (m *CallFactorProfile) WithPhoneNumber(v string) *CallFactorProfile {
	m.PhoneNumber = v
	return m
}


type EmailFactor struct {
	Profile EmailFactorProfile `json:"profile, omitempty"`
}

func (m *EmailFactor) WithProfile(v EmailFactorProfile) *EmailFactor {
	m.Profile = v
	return m
}


type EmailFactorProfile struct {
	Email string `json:"email, omitempty"`
}

func (m *EmailFactorProfile) WithEmail(v string) *EmailFactorProfile {
	m.Email = v
	return m
}

type FactorResource resource

type Factor struct {
	Embedded map[string]interface{} `json:"_embedded, omitempty"`
	Links map[string]interface{} `json:"_links, omitempty"`
	Device string `json:"device, omitempty"`
	DeviceType string `json:"deviceType, omitempty"`
	FactorType string `json:"factorType, omitempty"`
	Id string `json:"id, omitempty"`
	MfaStateTokenId string `json:"mfaStateTokenId, omitempty"`
	Profile FactorProfile `json:"profile, omitempty"`
	Provider string `json:"provider, omitempty"`
	RechallengeExistingFactor bool `json:"rechallengeExistingFactor, omitempty"`
	SessionId string `json:"sessionId, omitempty"`
	Status string `json:"status, omitempty"`
	UserId string `json:"userId, omitempty"`
	Verify VerifyFactorRequest `json:"verify, omitempty"`
}

func (m *Factor) WithDevice(v string) *Factor {
	m.Device = v
	return m
}
func (m *Factor) WithFactorType(v string) *Factor {
	m.FactorType = v
	return m
}
func (m *Factor) WithMfaStateTokenId(v string) *Factor {
	m.MfaStateTokenId = v
	return m
}
func (m *Factor) WithProfile(v FactorProfile) *Factor {
	m.Profile = v
	return m
}
func (m *Factor) WithProvider(v string) *Factor {
	m.Provider = v
	return m
}
func (m *Factor) WithRechallengeExistingFactor(v bool) *Factor {
	m.RechallengeExistingFactor = v
	return m
}
func (m *Factor) WithSessionId(v string) *Factor {
	m.SessionId = v
	return m
}
func (m *Factor) WithUserId(v string) *Factor {
	m.UserId = v
	return m
}
func (m *Factor) WithVerify(v VerifyFactorRequest) *Factor {
	m.Verify = v
	return m
}

func (m *FactorResource) Activate(userId string, factorId string, body VerifyFactorRequest) Factor {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/factors/"+factorId+"/lifecycle/activate", bytes.NewReader(iobytes))
		if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := Factor{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *FactorResource) Verify(userId string, factorId string, body VerifyFactorRequest) VerifyFactorResponse {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/users/"+userId+"/factors/"+factorId+"/verify", bytes.NewReader(iobytes))
		if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := VerifyFactorResponse{}

	json.Unmarshal(resp, &r)

	return r
}


type FactorProfile struct {
}



type FactorProvider struct {
}



type FactorResultType struct {
}



type FactorStatus struct {
}



type FactorType struct {
}



type HardwareFactor struct {
	Profile HardwareFactorProfile `json:"profile, omitempty"`
}

func (m *HardwareFactor) WithProfile(v HardwareFactorProfile) *HardwareFactor {
	m.Profile = v
	return m
}


type HardwareFactorProfile struct {
	CredentialId string `json:"credentialId, omitempty"`
}

func (m *HardwareFactorProfile) WithCredentialId(v string) *HardwareFactorProfile {
	m.CredentialId = v
	return m
}


type PushFactor struct {
	Profile PushFactorProfile `json:"profile, omitempty"`
}

func (m *PushFactor) WithProfile(v PushFactorProfile) *PushFactor {
	m.Profile = v
	return m
}


type PushFactorProfile struct {
	CredentialId string `json:"credentialId, omitempty"`
	DeviceType string `json:"deviceType, omitempty"`
	Name string `json:"name, omitempty"`
	Platform string `json:"platform, omitempty"`
	Version string `json:"version, omitempty"`
}

func (m *PushFactorProfile) WithCredentialId(v string) *PushFactorProfile {
	m.CredentialId = v
	return m
}
func (m *PushFactorProfile) WithName(v string) *PushFactorProfile {
	m.Name = v
	return m
}
func (m *PushFactorProfile) WithVersion(v string) *PushFactorProfile {
	m.Version = v
	return m
}


type SecurityQuestion struct {
	Answer string `json:"answer, omitempty"`
	Question string `json:"question, omitempty"`
	QuestionText string `json:"questionText, omitempty"`
}

func (m *SecurityQuestion) WithAnswer(v string) *SecurityQuestion {
	m.Answer = v
	return m
}
func (m *SecurityQuestion) WithQuestion(v string) *SecurityQuestion {
	m.Question = v
	return m
}
func (m *SecurityQuestion) WithQuestionText(v string) *SecurityQuestion {
	m.QuestionText = v
	return m
}


type SecurityQuestionFactor struct {
	Profile SecurityQuestionFactorProfile `json:"profile, omitempty"`
}

func (m *SecurityQuestionFactor) WithProfile(v SecurityQuestionFactorProfile) *SecurityQuestionFactor {
	m.Profile = v
	return m
}


type SecurityQuestionFactorProfile struct {
	Answer string `json:"answer, omitempty"`
	Question string `json:"question, omitempty"`
	QuestionText string `json:"questionText, omitempty"`
}

func (m *SecurityQuestionFactorProfile) WithAnswer(v string) *SecurityQuestionFactorProfile {
	m.Answer = v
	return m
}
func (m *SecurityQuestionFactorProfile) WithQuestion(v string) *SecurityQuestionFactorProfile {
	m.Question = v
	return m
}
func (m *SecurityQuestionFactorProfile) WithQuestionText(v string) *SecurityQuestionFactorProfile {
	m.QuestionText = v
	return m
}


type SmsFactor struct {
	Profile SmsFactorProfile `json:"profile, omitempty"`
}

func (m *SmsFactor) WithProfile(v SmsFactorProfile) *SmsFactor {
	m.Profile = v
	return m
}


type SmsFactorProfile struct {
	PhoneNumber string `json:"phoneNumber, omitempty"`
}

func (m *SmsFactorProfile) WithPhoneNumber(v string) *SmsFactorProfile {
	m.PhoneNumber = v
	return m
}


type TokenFactor struct {
	Profile TokenFactorProfile `json:"profile, omitempty"`
}

func (m *TokenFactor) WithProfile(v TokenFactorProfile) *TokenFactor {
	m.Profile = v
	return m
}


type TokenFactorProfile struct {
	CredentialId string `json:"credentialId, omitempty"`
}

func (m *TokenFactorProfile) WithCredentialId(v string) *TokenFactorProfile {
	m.CredentialId = v
	return m
}


type TotpFactor struct {
	Profile TotpFactorProfile `json:"profile, omitempty"`
}

func (m *TotpFactor) WithProfile(v TotpFactorProfile) *TotpFactor {
	m.Profile = v
	return m
}


type TotpFactorProfile struct {
	CredentialId string `json:"credentialId, omitempty"`
}

func (m *TotpFactorProfile) WithCredentialId(v string) *TotpFactorProfile {
	m.CredentialId = v
	return m
}


type VerifyFactorRequest struct {
	ActivationToken string `json:"activationToken, omitempty"`
	Answer string `json:"answer, omitempty"`
	NextPassCode string `json:"nextPassCode, omitempty"`
	PassCode string `json:"passCode, omitempty"`
}

func (m *VerifyFactorRequest) WithActivationToken(v string) *VerifyFactorRequest {
	m.ActivationToken = v
	return m
}
func (m *VerifyFactorRequest) WithAnswer(v string) *VerifyFactorRequest {
	m.Answer = v
	return m
}
func (m *VerifyFactorRequest) WithNextPassCode(v string) *VerifyFactorRequest {
	m.NextPassCode = v
	return m
}
func (m *VerifyFactorRequest) WithPassCode(v string) *VerifyFactorRequest {
	m.PassCode = v
	return m
}


type VerifyFactorResponse struct {
	Embedded map[string]interface{} `json:"_embedded, omitempty"`
	Links map[string]interface{} `json:"_links, omitempty"`
	ExpiresAt time.Time `json:"expiresAt, omitempty"`
	FactorResult string `json:"factorResult, omitempty"`
	FactorResultMessage string `json:"factorResultMessage, omitempty"`
}



type WebFactor struct {
	Profile WebFactorProfile `json:"profile, omitempty"`
}

func (m *WebFactor) WithProfile(v WebFactorProfile) *WebFactor {
	m.Profile = v
	return m
}


type WebFactorProfile struct {
	CredentialId string `json:"credentialId, omitempty"`
}

func (m *WebFactorProfile) WithCredentialId(v string) *WebFactorProfile {
	m.CredentialId = v
	return m
}




