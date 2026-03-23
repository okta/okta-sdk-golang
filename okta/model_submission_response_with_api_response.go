/*
Okta Admin Management API

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the SubmissionResponseWithApiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmissionResponseWithApiResponse{}

// SubmissionResponseWithApiResponse struct for SubmissionResponseWithApiResponse
type SubmissionResponseWithApiResponse struct {
	ApiService *ApiService `json:"apiService,omitempty"`
	// List of contact details for the app integration
	AppContactDetails []SubmissionResponseAppContactDetailsInner `json:"appContactDetails,omitempty"`
	AuthSettings      *AuthSettings                              `json:"authSettings,omitempty"`
	// List of capabilities supported by this integration with embedded protocol configurations
	Capabilities []SubmissionCapabilityEnhanced `json:"capabilities,omitempty"`
	// The most recent date and time when a Salesforce case associated with the OIN integration was updated
	CaseLastUpdated *string `json:"caseLastUpdated,omitempty"`
	// List of org-level variables for the customer per-tenant configuration. For example, a `subdomain` variable can be used in the ACS URL: `https://${org.subdomain}.example.com/saml/login`
	Config []SubmissionResponseConfigInner `json:"config,omitempty"`
	// Indicates whether the app submission uses a default logo or a custom-uploaded logo: * If `true`: Uses the default Okta-provided placeholder logo. * If `false`: Uses a custom logo type other than the default logo.
	DefaultLogo *bool `json:"defaultLogo,omitempty"`
	// A general description of your application and the benefits provided to your customers
	Description           *string                `json:"description,omitempty"`
	GlobalTokenRevocation *GlobalTokenRevocation `json:"globalTokenRevocation,omitempty"`
	// OIN Integration ID
	Id *string `json:"id,omitempty"`
	// Timestamp when the OIN Integration was last published
	LastPublished *string `json:"lastPublished,omitempty"`
	// Timestamp when the OIN Integration instance was last updated
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// ID of the user who made the last update
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	// URL to an uploaded application logo. This logo appears next to your app integration name in the OIN catalog. You must first [Upload an OIN Integration logo](/openapi/okta-management/management/tag/YourOinIntegrations/#tag/YourOinIntegrations/operation/uploadSubmissionLogo) to obtain the logo URL before you can specify this value.
	Logo *string `json:"logo,omitempty"`
	// The app integration name. This is the main title used for your integration in the OIN catalog.
	Name *string `json:"name,omitempty"`
	// Type of feature supported by the OIN integration
	OinFeatures  *string              `json:"oinFeatures,omitempty"`
	Provisioning *ProvisioningDetails `json:"provisioning,omitempty"`
	Sso          *Sso                 `json:"sso,omitempty"`
	// Status of the OIN Integration submission
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubmissionResponseWithApiResponse SubmissionResponseWithApiResponse

// NewSubmissionResponseWithApiResponse instantiates a new SubmissionResponseWithApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmissionResponseWithApiResponse() *SubmissionResponseWithApiResponse {
	this := SubmissionResponseWithApiResponse{}
	var defaultLogo bool = false
	this.DefaultLogo = &defaultLogo
	return &this
}

// NewSubmissionResponseWithApiResponseWithDefaults instantiates a new SubmissionResponseWithApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmissionResponseWithApiResponseWithDefaults() *SubmissionResponseWithApiResponse {
	this := SubmissionResponseWithApiResponse{}
	var defaultLogo bool = false
	this.DefaultLogo = &defaultLogo
	return &this
}

// GetApiService returns the ApiService field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetApiService() ApiService {
	if o == nil || IsNil(o.ApiService) {
		var ret ApiService
		return ret
	}
	return *o.ApiService
}

// GetApiServiceOk returns a tuple with the ApiService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetApiServiceOk() (*ApiService, bool) {
	if o == nil || IsNil(o.ApiService) {
		return nil, false
	}
	return o.ApiService, true
}

// HasApiService returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasApiService() bool {
	if o != nil && !IsNil(o.ApiService) {
		return true
	}

	return false
}

// SetApiService gets a reference to the given ApiService and assigns it to the ApiService field.
func (o *SubmissionResponseWithApiResponse) SetApiService(v ApiService) {
	o.ApiService = &v
}

// GetAppContactDetails returns the AppContactDetails field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetAppContactDetails() []SubmissionResponseAppContactDetailsInner {
	if o == nil || IsNil(o.AppContactDetails) {
		var ret []SubmissionResponseAppContactDetailsInner
		return ret
	}
	return o.AppContactDetails
}

// GetAppContactDetailsOk returns a tuple with the AppContactDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetAppContactDetailsOk() ([]SubmissionResponseAppContactDetailsInner, bool) {
	if o == nil || IsNil(o.AppContactDetails) {
		return nil, false
	}
	return o.AppContactDetails, true
}

// HasAppContactDetails returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasAppContactDetails() bool {
	if o != nil && !IsNil(o.AppContactDetails) {
		return true
	}

	return false
}

// SetAppContactDetails gets a reference to the given []SubmissionResponseAppContactDetailsInner and assigns it to the AppContactDetails field.
func (o *SubmissionResponseWithApiResponse) SetAppContactDetails(v []SubmissionResponseAppContactDetailsInner) {
	o.AppContactDetails = v
}

// GetAuthSettings returns the AuthSettings field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetAuthSettings() AuthSettings {
	if o == nil || IsNil(o.AuthSettings) {
		var ret AuthSettings
		return ret
	}
	return *o.AuthSettings
}

// GetAuthSettingsOk returns a tuple with the AuthSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetAuthSettingsOk() (*AuthSettings, bool) {
	if o == nil || IsNil(o.AuthSettings) {
		return nil, false
	}
	return o.AuthSettings, true
}

// HasAuthSettings returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasAuthSettings() bool {
	if o != nil && !IsNil(o.AuthSettings) {
		return true
	}

	return false
}

// SetAuthSettings gets a reference to the given AuthSettings and assigns it to the AuthSettings field.
func (o *SubmissionResponseWithApiResponse) SetAuthSettings(v AuthSettings) {
	o.AuthSettings = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetCapabilities() []SubmissionCapabilityEnhanced {
	if o == nil || IsNil(o.Capabilities) {
		var ret []SubmissionCapabilityEnhanced
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetCapabilitiesOk() ([]SubmissionCapabilityEnhanced, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []SubmissionCapabilityEnhanced and assigns it to the Capabilities field.
func (o *SubmissionResponseWithApiResponse) SetCapabilities(v []SubmissionCapabilityEnhanced) {
	o.Capabilities = v
}

// GetCaseLastUpdated returns the CaseLastUpdated field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetCaseLastUpdated() string {
	if o == nil || IsNil(o.CaseLastUpdated) {
		var ret string
		return ret
	}
	return *o.CaseLastUpdated
}

// GetCaseLastUpdatedOk returns a tuple with the CaseLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetCaseLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.CaseLastUpdated) {
		return nil, false
	}
	return o.CaseLastUpdated, true
}

// HasCaseLastUpdated returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasCaseLastUpdated() bool {
	if o != nil && !IsNil(o.CaseLastUpdated) {
		return true
	}

	return false
}

// SetCaseLastUpdated gets a reference to the given string and assigns it to the CaseLastUpdated field.
func (o *SubmissionResponseWithApiResponse) SetCaseLastUpdated(v string) {
	o.CaseLastUpdated = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetConfig() []SubmissionResponseConfigInner {
	if o == nil || IsNil(o.Config) {
		var ret []SubmissionResponseConfigInner
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetConfigOk() ([]SubmissionResponseConfigInner, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given []SubmissionResponseConfigInner and assigns it to the Config field.
func (o *SubmissionResponseWithApiResponse) SetConfig(v []SubmissionResponseConfigInner) {
	o.Config = v
}

// GetDefaultLogo returns the DefaultLogo field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetDefaultLogo() bool {
	if o == nil || IsNil(o.DefaultLogo) {
		var ret bool
		return ret
	}
	return *o.DefaultLogo
}

// GetDefaultLogoOk returns a tuple with the DefaultLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetDefaultLogoOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultLogo) {
		return nil, false
	}
	return o.DefaultLogo, true
}

// HasDefaultLogo returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasDefaultLogo() bool {
	if o != nil && !IsNil(o.DefaultLogo) {
		return true
	}

	return false
}

// SetDefaultLogo gets a reference to the given bool and assigns it to the DefaultLogo field.
func (o *SubmissionResponseWithApiResponse) SetDefaultLogo(v bool) {
	o.DefaultLogo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SubmissionResponseWithApiResponse) SetDescription(v string) {
	o.Description = &v
}

// GetGlobalTokenRevocation returns the GlobalTokenRevocation field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetGlobalTokenRevocation() GlobalTokenRevocation {
	if o == nil || IsNil(o.GlobalTokenRevocation) {
		var ret GlobalTokenRevocation
		return ret
	}
	return *o.GlobalTokenRevocation
}

// GetGlobalTokenRevocationOk returns a tuple with the GlobalTokenRevocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetGlobalTokenRevocationOk() (*GlobalTokenRevocation, bool) {
	if o == nil || IsNil(o.GlobalTokenRevocation) {
		return nil, false
	}
	return o.GlobalTokenRevocation, true
}

// HasGlobalTokenRevocation returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasGlobalTokenRevocation() bool {
	if o != nil && !IsNil(o.GlobalTokenRevocation) {
		return true
	}

	return false
}

// SetGlobalTokenRevocation gets a reference to the given GlobalTokenRevocation and assigns it to the GlobalTokenRevocation field.
func (o *SubmissionResponseWithApiResponse) SetGlobalTokenRevocation(v GlobalTokenRevocation) {
	o.GlobalTokenRevocation = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubmissionResponseWithApiResponse) SetId(v string) {
	o.Id = &v
}

// GetLastPublished returns the LastPublished field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetLastPublished() string {
	if o == nil || IsNil(o.LastPublished) {
		var ret string
		return ret
	}
	return *o.LastPublished
}

// GetLastPublishedOk returns a tuple with the LastPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetLastPublishedOk() (*string, bool) {
	if o == nil || IsNil(o.LastPublished) {
		return nil, false
	}
	return o.LastPublished, true
}

// HasLastPublished returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasLastPublished() bool {
	if o != nil && !IsNil(o.LastPublished) {
		return true
	}

	return false
}

// SetLastPublished gets a reference to the given string and assigns it to the LastPublished field.
func (o *SubmissionResponseWithApiResponse) SetLastPublished(v string) {
	o.LastPublished = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *SubmissionResponseWithApiResponse) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetLastUpdatedBy() string {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *SubmissionResponseWithApiResponse) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *SubmissionResponseWithApiResponse) SetLogo(v string) {
	o.Logo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubmissionResponseWithApiResponse) SetName(v string) {
	o.Name = &v
}

// GetOinFeatures returns the OinFeatures field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetOinFeatures() string {
	if o == nil || IsNil(o.OinFeatures) {
		var ret string
		return ret
	}
	return *o.OinFeatures
}

// GetOinFeaturesOk returns a tuple with the OinFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetOinFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.OinFeatures) {
		return nil, false
	}
	return o.OinFeatures, true
}

// HasOinFeatures returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasOinFeatures() bool {
	if o != nil && !IsNil(o.OinFeatures) {
		return true
	}

	return false
}

// SetOinFeatures gets a reference to the given string and assigns it to the OinFeatures field.
func (o *SubmissionResponseWithApiResponse) SetOinFeatures(v string) {
	o.OinFeatures = &v
}

// GetProvisioning returns the Provisioning field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetProvisioning() ProvisioningDetails {
	if o == nil || IsNil(o.Provisioning) {
		var ret ProvisioningDetails
		return ret
	}
	return *o.Provisioning
}

// GetProvisioningOk returns a tuple with the Provisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetProvisioningOk() (*ProvisioningDetails, bool) {
	if o == nil || IsNil(o.Provisioning) {
		return nil, false
	}
	return o.Provisioning, true
}

// HasProvisioning returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasProvisioning() bool {
	if o != nil && !IsNil(o.Provisioning) {
		return true
	}

	return false
}

// SetProvisioning gets a reference to the given ProvisioningDetails and assigns it to the Provisioning field.
func (o *SubmissionResponseWithApiResponse) SetProvisioning(v ProvisioningDetails) {
	o.Provisioning = &v
}

// GetSso returns the Sso field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetSso() Sso {
	if o == nil || IsNil(o.Sso) {
		var ret Sso
		return ret
	}
	return *o.Sso
}

// GetSsoOk returns a tuple with the Sso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetSsoOk() (*Sso, bool) {
	if o == nil || IsNil(o.Sso) {
		return nil, false
	}
	return o.Sso, true
}

// HasSso returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasSso() bool {
	if o != nil && !IsNil(o.Sso) {
		return true
	}

	return false
}

// SetSso gets a reference to the given Sso and assigns it to the Sso field.
func (o *SubmissionResponseWithApiResponse) SetSso(v Sso) {
	o.Sso = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SubmissionResponseWithApiResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseWithApiResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SubmissionResponseWithApiResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SubmissionResponseWithApiResponse) SetStatus(v string) {
	o.Status = &v
}

func (o SubmissionResponseWithApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmissionResponseWithApiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiService) {
		toSerialize["apiService"] = o.ApiService
	}
	if !IsNil(o.AppContactDetails) {
		toSerialize["appContactDetails"] = o.AppContactDetails
	}
	if !IsNil(o.AuthSettings) {
		toSerialize["authSettings"] = o.AuthSettings
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.CaseLastUpdated) {
		toSerialize["caseLastUpdated"] = o.CaseLastUpdated
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.DefaultLogo) {
		toSerialize["defaultLogo"] = o.DefaultLogo
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.GlobalTokenRevocation) {
		toSerialize["globalTokenRevocation"] = o.GlobalTokenRevocation
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastPublished) {
		toSerialize["lastPublished"] = o.LastPublished
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.LastUpdatedBy) {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OinFeatures) {
		toSerialize["oinFeatures"] = o.OinFeatures
	}
	if !IsNil(o.Provisioning) {
		toSerialize["provisioning"] = o.Provisioning
	}
	if !IsNil(o.Sso) {
		toSerialize["sso"] = o.Sso
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubmissionResponseWithApiResponse) UnmarshalJSON(data []byte) (err error) {
	varSubmissionResponseWithApiResponse := _SubmissionResponseWithApiResponse{}

	err = json.Unmarshal(data, &varSubmissionResponseWithApiResponse)

	if err != nil {
		return err
	}

	*o = SubmissionResponseWithApiResponse(varSubmissionResponseWithApiResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apiService")
		delete(additionalProperties, "appContactDetails")
		delete(additionalProperties, "authSettings")
		delete(additionalProperties, "capabilities")
		delete(additionalProperties, "caseLastUpdated")
		delete(additionalProperties, "config")
		delete(additionalProperties, "defaultLogo")
		delete(additionalProperties, "description")
		delete(additionalProperties, "globalTokenRevocation")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastPublished")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "name")
		delete(additionalProperties, "oinFeatures")
		delete(additionalProperties, "provisioning")
		delete(additionalProperties, "sso")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubmissionResponseWithApiResponse struct {
	value *SubmissionResponseWithApiResponse
	isSet bool
}

func (v NullableSubmissionResponseWithApiResponse) Get() *SubmissionResponseWithApiResponse {
	return v.value
}

func (v *NullableSubmissionResponseWithApiResponse) Set(val *SubmissionResponseWithApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmissionResponseWithApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmissionResponseWithApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmissionResponseWithApiResponse(val *SubmissionResponseWithApiResponse) *NullableSubmissionResponseWithApiResponse {
	return &NullableSubmissionResponseWithApiResponse{value: val, isSet: true}
}

func (v NullableSubmissionResponseWithApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmissionResponseWithApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
