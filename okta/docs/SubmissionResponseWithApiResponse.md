# SubmissionResponseWithApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiService** | Pointer to [**ApiService**](ApiService.md) |  | [optional] 
**AppContactDetails** | Pointer to [**[]SubmissionResponseAppContactDetailsInner**](SubmissionResponseAppContactDetailsInner.md) | List of contact details for the app integration | [optional] 
**AuthSettings** | Pointer to [**AuthSettings**](AuthSettings.md) |  | [optional] 
**Capabilities** | Pointer to [**[]SubmissionCapabilityEnhanced**](SubmissionCapabilityEnhanced.md) | List of capabilities supported by this integration with embedded protocol configurations | [optional] 
**CaseLastUpdated** | Pointer to **string** | The most recent date and time when a Salesforce case associated with the OIN integration was updated | [optional] [readonly] 
**Config** | Pointer to [**[]SubmissionResponseConfigInner**](SubmissionResponseConfigInner.md) | List of org-level variables for the customer per-tenant configuration. For example, a &#x60;subdomain&#x60; variable can be used in the ACS URL: &#x60;https://${org.subdomain}.example.com/saml/login&#x60; | [optional] 
**DefaultLogo** | Pointer to **bool** | Indicates whether the app submission uses a default logo or a custom-uploaded logo: * If &#x60;true&#x60;: Uses the default Okta-provided placeholder logo. * If &#x60;false&#x60;: Uses a custom logo type other than the default logo.  | [optional] [default to false]
**Description** | Pointer to **string** | A general description of your application and the benefits provided to your customers | [optional] 
**GlobalTokenRevocation** | Pointer to [**GlobalTokenRevocation**](GlobalTokenRevocation.md) |  | [optional] 
**Id** | Pointer to **string** | OIN Integration ID | [optional] [readonly] 
**LastPublished** | Pointer to **string** | Timestamp when the OIN Integration was last published | [optional] [readonly] 
**LastUpdated** | Pointer to **string** | Timestamp when the OIN Integration instance was last updated | [optional] [readonly] 
**LastUpdatedBy** | Pointer to **string** | ID of the user who made the last update | [optional] [readonly] 
**Logo** | Pointer to **string** | URL to an uploaded application logo. This logo appears next to your app integration name in the OIN catalog. You must first [Upload an OIN Integration logo](/openapi/okta-management/management/tag/YourOinIntegrations/#tag/YourOinIntegrations/operation/uploadSubmissionLogo) to obtain the logo URL before you can specify this value. | [optional] 
**Name** | Pointer to **string** | The app integration name. This is the main title used for your integration in the OIN catalog. | [optional] 
**OinFeatures** | Pointer to **string** | Type of feature supported by the OIN integration | [optional] [readonly] 
**Provisioning** | Pointer to [**ProvisioningDetails**](ProvisioningDetails.md) |  | [optional] 
**Sso** | Pointer to [**Sso**](Sso.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the OIN Integration submission | [optional] [readonly] 

## Methods

### NewSubmissionResponseWithApiResponse

`func NewSubmissionResponseWithApiResponse() *SubmissionResponseWithApiResponse`

NewSubmissionResponseWithApiResponse instantiates a new SubmissionResponseWithApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionResponseWithApiResponseWithDefaults

`func NewSubmissionResponseWithApiResponseWithDefaults() *SubmissionResponseWithApiResponse`

NewSubmissionResponseWithApiResponseWithDefaults instantiates a new SubmissionResponseWithApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiService

`func (o *SubmissionResponseWithApiResponse) GetApiService() ApiService`

GetApiService returns the ApiService field if non-nil, zero value otherwise.

### GetApiServiceOk

`func (o *SubmissionResponseWithApiResponse) GetApiServiceOk() (*ApiService, bool)`

GetApiServiceOk returns a tuple with the ApiService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiService

`func (o *SubmissionResponseWithApiResponse) SetApiService(v ApiService)`

SetApiService sets ApiService field to given value.

### HasApiService

`func (o *SubmissionResponseWithApiResponse) HasApiService() bool`

HasApiService returns a boolean if a field has been set.

### GetAppContactDetails

`func (o *SubmissionResponseWithApiResponse) GetAppContactDetails() []SubmissionResponseAppContactDetailsInner`

GetAppContactDetails returns the AppContactDetails field if non-nil, zero value otherwise.

### GetAppContactDetailsOk

`func (o *SubmissionResponseWithApiResponse) GetAppContactDetailsOk() (*[]SubmissionResponseAppContactDetailsInner, bool)`

GetAppContactDetailsOk returns a tuple with the AppContactDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppContactDetails

`func (o *SubmissionResponseWithApiResponse) SetAppContactDetails(v []SubmissionResponseAppContactDetailsInner)`

SetAppContactDetails sets AppContactDetails field to given value.

### HasAppContactDetails

`func (o *SubmissionResponseWithApiResponse) HasAppContactDetails() bool`

HasAppContactDetails returns a boolean if a field has been set.

### GetAuthSettings

`func (o *SubmissionResponseWithApiResponse) GetAuthSettings() AuthSettings`

GetAuthSettings returns the AuthSettings field if non-nil, zero value otherwise.

### GetAuthSettingsOk

`func (o *SubmissionResponseWithApiResponse) GetAuthSettingsOk() (*AuthSettings, bool)`

GetAuthSettingsOk returns a tuple with the AuthSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSettings

`func (o *SubmissionResponseWithApiResponse) SetAuthSettings(v AuthSettings)`

SetAuthSettings sets AuthSettings field to given value.

### HasAuthSettings

`func (o *SubmissionResponseWithApiResponse) HasAuthSettings() bool`

HasAuthSettings returns a boolean if a field has been set.

### GetCapabilities

`func (o *SubmissionResponseWithApiResponse) GetCapabilities() []SubmissionCapabilityEnhanced`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *SubmissionResponseWithApiResponse) GetCapabilitiesOk() (*[]SubmissionCapabilityEnhanced, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *SubmissionResponseWithApiResponse) SetCapabilities(v []SubmissionCapabilityEnhanced)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *SubmissionResponseWithApiResponse) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCaseLastUpdated

`func (o *SubmissionResponseWithApiResponse) GetCaseLastUpdated() string`

GetCaseLastUpdated returns the CaseLastUpdated field if non-nil, zero value otherwise.

### GetCaseLastUpdatedOk

`func (o *SubmissionResponseWithApiResponse) GetCaseLastUpdatedOk() (*string, bool)`

GetCaseLastUpdatedOk returns a tuple with the CaseLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseLastUpdated

`func (o *SubmissionResponseWithApiResponse) SetCaseLastUpdated(v string)`

SetCaseLastUpdated sets CaseLastUpdated field to given value.

### HasCaseLastUpdated

`func (o *SubmissionResponseWithApiResponse) HasCaseLastUpdated() bool`

HasCaseLastUpdated returns a boolean if a field has been set.

### GetConfig

`func (o *SubmissionResponseWithApiResponse) GetConfig() []SubmissionResponseConfigInner`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SubmissionResponseWithApiResponse) GetConfigOk() (*[]SubmissionResponseConfigInner, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SubmissionResponseWithApiResponse) SetConfig(v []SubmissionResponseConfigInner)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SubmissionResponseWithApiResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDefaultLogo

`func (o *SubmissionResponseWithApiResponse) GetDefaultLogo() bool`

GetDefaultLogo returns the DefaultLogo field if non-nil, zero value otherwise.

### GetDefaultLogoOk

`func (o *SubmissionResponseWithApiResponse) GetDefaultLogoOk() (*bool, bool)`

GetDefaultLogoOk returns a tuple with the DefaultLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLogo

`func (o *SubmissionResponseWithApiResponse) SetDefaultLogo(v bool)`

SetDefaultLogo sets DefaultLogo field to given value.

### HasDefaultLogo

`func (o *SubmissionResponseWithApiResponse) HasDefaultLogo() bool`

HasDefaultLogo returns a boolean if a field has been set.

### GetDescription

`func (o *SubmissionResponseWithApiResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubmissionResponseWithApiResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubmissionResponseWithApiResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubmissionResponseWithApiResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGlobalTokenRevocation

`func (o *SubmissionResponseWithApiResponse) GetGlobalTokenRevocation() GlobalTokenRevocation`

GetGlobalTokenRevocation returns the GlobalTokenRevocation field if non-nil, zero value otherwise.

### GetGlobalTokenRevocationOk

`func (o *SubmissionResponseWithApiResponse) GetGlobalTokenRevocationOk() (*GlobalTokenRevocation, bool)`

GetGlobalTokenRevocationOk returns a tuple with the GlobalTokenRevocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTokenRevocation

`func (o *SubmissionResponseWithApiResponse) SetGlobalTokenRevocation(v GlobalTokenRevocation)`

SetGlobalTokenRevocation sets GlobalTokenRevocation field to given value.

### HasGlobalTokenRevocation

`func (o *SubmissionResponseWithApiResponse) HasGlobalTokenRevocation() bool`

HasGlobalTokenRevocation returns a boolean if a field has been set.

### GetId

`func (o *SubmissionResponseWithApiResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubmissionResponseWithApiResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubmissionResponseWithApiResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubmissionResponseWithApiResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastPublished

`func (o *SubmissionResponseWithApiResponse) GetLastPublished() string`

GetLastPublished returns the LastPublished field if non-nil, zero value otherwise.

### GetLastPublishedOk

`func (o *SubmissionResponseWithApiResponse) GetLastPublishedOk() (*string, bool)`

GetLastPublishedOk returns a tuple with the LastPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPublished

`func (o *SubmissionResponseWithApiResponse) SetLastPublished(v string)`

SetLastPublished sets LastPublished field to given value.

### HasLastPublished

`func (o *SubmissionResponseWithApiResponse) HasLastPublished() bool`

HasLastPublished returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SubmissionResponseWithApiResponse) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SubmissionResponseWithApiResponse) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SubmissionResponseWithApiResponse) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SubmissionResponseWithApiResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *SubmissionResponseWithApiResponse) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *SubmissionResponseWithApiResponse) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *SubmissionResponseWithApiResponse) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *SubmissionResponseWithApiResponse) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetLogo

`func (o *SubmissionResponseWithApiResponse) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *SubmissionResponseWithApiResponse) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *SubmissionResponseWithApiResponse) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *SubmissionResponseWithApiResponse) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *SubmissionResponseWithApiResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubmissionResponseWithApiResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubmissionResponseWithApiResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubmissionResponseWithApiResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOinFeatures

`func (o *SubmissionResponseWithApiResponse) GetOinFeatures() string`

GetOinFeatures returns the OinFeatures field if non-nil, zero value otherwise.

### GetOinFeaturesOk

`func (o *SubmissionResponseWithApiResponse) GetOinFeaturesOk() (*string, bool)`

GetOinFeaturesOk returns a tuple with the OinFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOinFeatures

`func (o *SubmissionResponseWithApiResponse) SetOinFeatures(v string)`

SetOinFeatures sets OinFeatures field to given value.

### HasOinFeatures

`func (o *SubmissionResponseWithApiResponse) HasOinFeatures() bool`

HasOinFeatures returns a boolean if a field has been set.

### GetProvisioning

`func (o *SubmissionResponseWithApiResponse) GetProvisioning() ProvisioningDetails`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *SubmissionResponseWithApiResponse) GetProvisioningOk() (*ProvisioningDetails, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *SubmissionResponseWithApiResponse) SetProvisioning(v ProvisioningDetails)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *SubmissionResponseWithApiResponse) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetSso

`func (o *SubmissionResponseWithApiResponse) GetSso() Sso`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *SubmissionResponseWithApiResponse) GetSsoOk() (*Sso, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *SubmissionResponseWithApiResponse) SetSso(v Sso)`

SetSso sets Sso field to given value.

### HasSso

`func (o *SubmissionResponseWithApiResponse) HasSso() bool`

HasSso returns a boolean if a field has been set.

### GetStatus

`func (o *SubmissionResponseWithApiResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubmissionResponseWithApiResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubmissionResponseWithApiResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubmissionResponseWithApiResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


