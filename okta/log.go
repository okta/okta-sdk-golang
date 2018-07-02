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
	"time"
)

type LogActor struct {
	AlternateId string `json:"alternateId, omitempty"`
	Detail map[string]interface{} `json:"detail, omitempty"`
	DisplayName string `json:"displayName, omitempty"`
	Id string `json:"id, omitempty"`
	Type string `json:"type, omitempty"`
}



type LogAuthenticationContext struct {
	AuthenticationProvider string `json:"authenticationProvider, omitempty"`
	AuthenticationStep int64 `json:"authenticationStep, omitempty"`
	CredentialProvider map[string]interface{} `json:"credentialProvider, omitempty"`
	CredentialType map[string]interface{} `json:"credentialType, omitempty"`
	ExternalSessionId string `json:"externalSessionId, omitempty"`
	Interface string `json:"interface, omitempty"`
	Issuer LogIssuer `json:"issuer, omitempty"`
}



type LogAuthenticationProvider struct {
}



type LogClient struct {
	Device string `json:"device, omitempty"`
	GeographicalContext LogGeographicalContext `json:"geographicalContext, omitempty"`
	Id string `json:"id, omitempty"`
	IpAddress string `json:"ipAddress, omitempty"`
	UserAgent LogUserAgent `json:"userAgent, omitempty"`
	Zone string `json:"zone, omitempty"`
}



type LogCredentialProvider struct {
}



type LogCredentialType struct {
}



type LogDebugContext struct {
	DebugData map[string]interface{} `json:"debugData, omitempty"`
}



type LogEvent struct {
	Actor LogActor `json:"actor, omitempty"`
	AuthenticationContext LogAuthenticationContext `json:"authenticationContext, omitempty"`
	Client LogClient `json:"client, omitempty"`
	DebugContext LogDebugContext `json:"debugContext, omitempty"`
	DisplayMessage string `json:"displayMessage, omitempty"`
	EventType string `json:"eventType, omitempty"`
	LegacyEventType string `json:"legacyEventType, omitempty"`
	Outcome LogOutcome `json:"outcome, omitempty"`
	Published time.Time `json:"published, omitempty"`
	Request LogRequest `json:"request, omitempty"`
	SecurityContext LogSecurityContext `json:"securityContext, omitempty"`
	Severity string `json:"severity, omitempty"`
	Target map[string]interface{} `json:"target, omitempty"`
	Transaction LogTransaction `json:"transaction, omitempty"`
	Uuid string `json:"uuid, omitempty"`
	Version string `json:"version, omitempty"`
}



type LogGeographicalContext struct {
	City string `json:"city, omitempty"`
	Country string `json:"country, omitempty"`
	Geolocation LogGeolocation `json:"geolocation, omitempty"`
	PostalCode string `json:"postalCode, omitempty"`
	State string `json:"state, omitempty"`
}



type LogGeolocation struct {
	Lat float64 `json:"lat, omitempty"`
	Lon float64 `json:"lon, omitempty"`
}



type LogIpAddress struct {
	GeographicalContext LogGeographicalContext `json:"geographicalContext, omitempty"`
	Ip string `json:"ip, omitempty"`
	Source string `json:"source, omitempty"`
	Version string `json:"version, omitempty"`
}



type LogIssuer struct {
	Id string `json:"id, omitempty"`
	Type string `json:"type, omitempty"`
}



type LogOutcome struct {
	Reason string `json:"reason, omitempty"`
	Result string `json:"result, omitempty"`
}



type LogRequest struct {
	IpChain map[string]interface{} `json:"ipChain, omitempty"`
}



type LogSecurityContext struct {
	AsNumber int64 `json:"asNumber, omitempty"`
	AsOrg string `json:"asOrg, omitempty"`
	Domain string `json:"domain, omitempty"`
	IsProxy bool `json:"isProxy, omitempty"`
	Isp string `json:"isp, omitempty"`
}



type LogSeverity struct {
}



type LogTarget struct {
	AlternateId string `json:"alternateId, omitempty"`
	DetailEntry map[string]interface{} `json:"detailEntry, omitempty"`
	DisplayName string `json:"displayName, omitempty"`
	Id string `json:"id, omitempty"`
	Type string `json:"type, omitempty"`
}



type LogTransaction struct {
	Detail map[string]interface{} `json:"detail, omitempty"`
	Id string `json:"id, omitempty"`
	Type string `json:"type, omitempty"`
}



type LogUserAgent struct {
	Browser string `json:"browser, omitempty"`
	Os string `json:"os, omitempty"`
	RawUserAgent string `json:"rawUserAgent, omitempty"`
}





