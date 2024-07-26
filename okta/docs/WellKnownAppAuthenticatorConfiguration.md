# WellKnownAppAuthenticatorConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppAuthenticatorEnrollEndpoint** | Pointer to **string** | The authenticator enrollment endpoint | [optional] 
**AuthenticatorId** | Pointer to **string** | The unique identifier of the app authenticator | [optional] 
**CreatedDate** | Pointer to **time.Time** | Timestamp when the Authenticator was created | [optional] 
**Key** | Pointer to **string** | A human-readable string that identifies the Authenticator | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the Authenticator was last modified | [optional] 
**Name** | Pointer to **string** | The authenticator display name | [optional] 
**OrgId** | Pointer to **string** | The &#x60;id&#x60; of the Okta Org | [optional] 
**Settings** | Pointer to [**WellKnownAppAuthenticatorConfigurationSettings**](WellKnownAppAuthenticatorConfigurationSettings.md) |  | [optional] 
**SupportedMethods** | Pointer to [**[]SupportedMethods**](SupportedMethods.md) |  | [optional] 
**Type** | Pointer to **string** | The type of Authenticator | [optional] 

## Methods

### NewWellKnownAppAuthenticatorConfiguration

`func NewWellKnownAppAuthenticatorConfiguration() *WellKnownAppAuthenticatorConfiguration`

NewWellKnownAppAuthenticatorConfiguration instantiates a new WellKnownAppAuthenticatorConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWellKnownAppAuthenticatorConfigurationWithDefaults

`func NewWellKnownAppAuthenticatorConfigurationWithDefaults() *WellKnownAppAuthenticatorConfiguration`

NewWellKnownAppAuthenticatorConfigurationWithDefaults instantiates a new WellKnownAppAuthenticatorConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppAuthenticatorEnrollEndpoint

`func (o *WellKnownAppAuthenticatorConfiguration) GetAppAuthenticatorEnrollEndpoint() string`

GetAppAuthenticatorEnrollEndpoint returns the AppAuthenticatorEnrollEndpoint field if non-nil, zero value otherwise.

### GetAppAuthenticatorEnrollEndpointOk

`func (o *WellKnownAppAuthenticatorConfiguration) GetAppAuthenticatorEnrollEndpointOk() (*string, bool)`

GetAppAuthenticatorEnrollEndpointOk returns a tuple with the AppAuthenticatorEnrollEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAuthenticatorEnrollEndpoint

`func (o *WellKnownAppAuthenticatorConfiguration) SetAppAuthenticatorEnrollEndpoint(v string)`

SetAppAuthenticatorEnrollEndpoint sets AppAuthenticatorEnrollEndpoint field to given value.

### HasAppAuthenticatorEnrollEndpoint

`func (o *WellKnownAppAuthenticatorConfiguration) HasAppAuthenticatorEnrollEndpoint() bool`

HasAppAuthenticatorEnrollEndpoint returns a boolean if a field has been set.

### GetAuthenticatorId

`func (o *WellKnownAppAuthenticatorConfiguration) GetAuthenticatorId() string`

GetAuthenticatorId returns the AuthenticatorId field if non-nil, zero value otherwise.

### GetAuthenticatorIdOk

`func (o *WellKnownAppAuthenticatorConfiguration) GetAuthenticatorIdOk() (*string, bool)`

GetAuthenticatorIdOk returns a tuple with the AuthenticatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorId

`func (o *WellKnownAppAuthenticatorConfiguration) SetAuthenticatorId(v string)`

SetAuthenticatorId sets AuthenticatorId field to given value.

### HasAuthenticatorId

`func (o *WellKnownAppAuthenticatorConfiguration) HasAuthenticatorId() bool`

HasAuthenticatorId returns a boolean if a field has been set.

### GetCreatedDate

`func (o *WellKnownAppAuthenticatorConfiguration) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *WellKnownAppAuthenticatorConfiguration) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *WellKnownAppAuthenticatorConfiguration) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *WellKnownAppAuthenticatorConfiguration) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetKey

`func (o *WellKnownAppAuthenticatorConfiguration) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WellKnownAppAuthenticatorConfiguration) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WellKnownAppAuthenticatorConfiguration) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *WellKnownAppAuthenticatorConfiguration) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WellKnownAppAuthenticatorConfiguration) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WellKnownAppAuthenticatorConfiguration) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WellKnownAppAuthenticatorConfiguration) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WellKnownAppAuthenticatorConfiguration) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *WellKnownAppAuthenticatorConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WellKnownAppAuthenticatorConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WellKnownAppAuthenticatorConfiguration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WellKnownAppAuthenticatorConfiguration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *WellKnownAppAuthenticatorConfiguration) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WellKnownAppAuthenticatorConfiguration) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WellKnownAppAuthenticatorConfiguration) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *WellKnownAppAuthenticatorConfiguration) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSettings

`func (o *WellKnownAppAuthenticatorConfiguration) GetSettings() WellKnownAppAuthenticatorConfigurationSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *WellKnownAppAuthenticatorConfiguration) GetSettingsOk() (*WellKnownAppAuthenticatorConfigurationSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *WellKnownAppAuthenticatorConfiguration) SetSettings(v WellKnownAppAuthenticatorConfigurationSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *WellKnownAppAuthenticatorConfiguration) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSupportedMethods

`func (o *WellKnownAppAuthenticatorConfiguration) GetSupportedMethods() []SupportedMethods`

GetSupportedMethods returns the SupportedMethods field if non-nil, zero value otherwise.

### GetSupportedMethodsOk

`func (o *WellKnownAppAuthenticatorConfiguration) GetSupportedMethodsOk() (*[]SupportedMethods, bool)`

GetSupportedMethodsOk returns a tuple with the SupportedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMethods

`func (o *WellKnownAppAuthenticatorConfiguration) SetSupportedMethods(v []SupportedMethods)`

SetSupportedMethods sets SupportedMethods field to given value.

### HasSupportedMethods

`func (o *WellKnownAppAuthenticatorConfiguration) HasSupportedMethods() bool`

HasSupportedMethods returns a boolean if a field has been set.

### GetType

`func (o *WellKnownAppAuthenticatorConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WellKnownAppAuthenticatorConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WellKnownAppAuthenticatorConfiguration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WellKnownAppAuthenticatorConfiguration) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


