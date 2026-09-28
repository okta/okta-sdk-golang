# RegistrationAuthorityConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AadAuthority** | Pointer to **string** | The Microsoft Entra ID tenant domain that Okta authenticates against, for example &#x60;example.onmicrosoft.com&#x60; | [optional] 
**AadClientId** | Pointer to **string** | The client ID of the Microsoft Entra ID application that Okta uses to validate Microsoft Intune SCEP requests | [optional] 
**AadClientSecret** | Pointer to **string** | The client secret of the Microsoft Entra ID application. Write-only — it&#39;s never returned in a response. | [optional] 
**Provider** | Pointer to **string** | The mobile device management provider that validates the challenge. Use &#x60;INTUNE&#x60; for Microsoft Intune. | [optional] 

## Methods

### NewRegistrationAuthorityConfigInfo

`func NewRegistrationAuthorityConfigInfo() *RegistrationAuthorityConfigInfo`

NewRegistrationAuthorityConfigInfo instantiates a new RegistrationAuthorityConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationAuthorityConfigInfoWithDefaults

`func NewRegistrationAuthorityConfigInfoWithDefaults() *RegistrationAuthorityConfigInfo`

NewRegistrationAuthorityConfigInfoWithDefaults instantiates a new RegistrationAuthorityConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAadAuthority

`func (o *RegistrationAuthorityConfigInfo) GetAadAuthority() string`

GetAadAuthority returns the AadAuthority field if non-nil, zero value otherwise.

### GetAadAuthorityOk

`func (o *RegistrationAuthorityConfigInfo) GetAadAuthorityOk() (*string, bool)`

GetAadAuthorityOk returns a tuple with the AadAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAadAuthority

`func (o *RegistrationAuthorityConfigInfo) SetAadAuthority(v string)`

SetAadAuthority sets AadAuthority field to given value.

### HasAadAuthority

`func (o *RegistrationAuthorityConfigInfo) HasAadAuthority() bool`

HasAadAuthority returns a boolean if a field has been set.

### GetAadClientId

`func (o *RegistrationAuthorityConfigInfo) GetAadClientId() string`

GetAadClientId returns the AadClientId field if non-nil, zero value otherwise.

### GetAadClientIdOk

`func (o *RegistrationAuthorityConfigInfo) GetAadClientIdOk() (*string, bool)`

GetAadClientIdOk returns a tuple with the AadClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAadClientId

`func (o *RegistrationAuthorityConfigInfo) SetAadClientId(v string)`

SetAadClientId sets AadClientId field to given value.

### HasAadClientId

`func (o *RegistrationAuthorityConfigInfo) HasAadClientId() bool`

HasAadClientId returns a boolean if a field has been set.

### GetAadClientSecret

`func (o *RegistrationAuthorityConfigInfo) GetAadClientSecret() string`

GetAadClientSecret returns the AadClientSecret field if non-nil, zero value otherwise.

### GetAadClientSecretOk

`func (o *RegistrationAuthorityConfigInfo) GetAadClientSecretOk() (*string, bool)`

GetAadClientSecretOk returns a tuple with the AadClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAadClientSecret

`func (o *RegistrationAuthorityConfigInfo) SetAadClientSecret(v string)`

SetAadClientSecret sets AadClientSecret field to given value.

### HasAadClientSecret

`func (o *RegistrationAuthorityConfigInfo) HasAadClientSecret() bool`

HasAadClientSecret returns a boolean if a field has been set.

### GetProvider

`func (o *RegistrationAuthorityConfigInfo) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *RegistrationAuthorityConfigInfo) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *RegistrationAuthorityConfigInfo) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *RegistrationAuthorityConfigInfo) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


