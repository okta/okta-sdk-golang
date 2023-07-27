# AuthenticatorSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedFor** | Pointer to [**AllowedForEnum**](AllowedForEnum.md) |  | [optional] 
**AppInstanceId** | Pointer to **string** |  | [optional] 
**ChannelBinding** | Pointer to [**ChannelBinding**](ChannelBinding.md) |  | [optional] 
**Compliance** | Pointer to [**Compliance**](Compliance.md) |  | [optional] 
**TokenLifetimeInMinutes** | Pointer to **int32** |  | [optional] 
**UserVerification** | Pointer to [**UserVerificationEnum**](UserVerificationEnum.md) |  | [optional] 

## Methods

### NewAuthenticatorSettings

`func NewAuthenticatorSettings() *AuthenticatorSettings`

NewAuthenticatorSettings instantiates a new AuthenticatorSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorSettingsWithDefaults

`func NewAuthenticatorSettingsWithDefaults() *AuthenticatorSettings`

NewAuthenticatorSettingsWithDefaults instantiates a new AuthenticatorSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedFor

`func (o *AuthenticatorSettings) GetAllowedFor() AllowedForEnum`

GetAllowedFor returns the AllowedFor field if non-nil, zero value otherwise.

### GetAllowedForOk

`func (o *AuthenticatorSettings) GetAllowedForOk() (*AllowedForEnum, bool)`

GetAllowedForOk returns a tuple with the AllowedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedFor

`func (o *AuthenticatorSettings) SetAllowedFor(v AllowedForEnum)`

SetAllowedFor sets AllowedFor field to given value.

### HasAllowedFor

`func (o *AuthenticatorSettings) HasAllowedFor() bool`

HasAllowedFor returns a boolean if a field has been set.

### GetAppInstanceId

`func (o *AuthenticatorSettings) GetAppInstanceId() string`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *AuthenticatorSettings) GetAppInstanceIdOk() (*string, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *AuthenticatorSettings) SetAppInstanceId(v string)`

SetAppInstanceId sets AppInstanceId field to given value.

### HasAppInstanceId

`func (o *AuthenticatorSettings) HasAppInstanceId() bool`

HasAppInstanceId returns a boolean if a field has been set.

### GetChannelBinding

`func (o *AuthenticatorSettings) GetChannelBinding() ChannelBinding`

GetChannelBinding returns the ChannelBinding field if non-nil, zero value otherwise.

### GetChannelBindingOk

`func (o *AuthenticatorSettings) GetChannelBindingOk() (*ChannelBinding, bool)`

GetChannelBindingOk returns a tuple with the ChannelBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelBinding

`func (o *AuthenticatorSettings) SetChannelBinding(v ChannelBinding)`

SetChannelBinding sets ChannelBinding field to given value.

### HasChannelBinding

`func (o *AuthenticatorSettings) HasChannelBinding() bool`

HasChannelBinding returns a boolean if a field has been set.

### GetCompliance

`func (o *AuthenticatorSettings) GetCompliance() Compliance`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *AuthenticatorSettings) GetComplianceOk() (*Compliance, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *AuthenticatorSettings) SetCompliance(v Compliance)`

SetCompliance sets Compliance field to given value.

### HasCompliance

`func (o *AuthenticatorSettings) HasCompliance() bool`

HasCompliance returns a boolean if a field has been set.

### GetTokenLifetimeInMinutes

`func (o *AuthenticatorSettings) GetTokenLifetimeInMinutes() int32`

GetTokenLifetimeInMinutes returns the TokenLifetimeInMinutes field if non-nil, zero value otherwise.

### GetTokenLifetimeInMinutesOk

`func (o *AuthenticatorSettings) GetTokenLifetimeInMinutesOk() (*int32, bool)`

GetTokenLifetimeInMinutesOk returns a tuple with the TokenLifetimeInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLifetimeInMinutes

`func (o *AuthenticatorSettings) SetTokenLifetimeInMinutes(v int32)`

SetTokenLifetimeInMinutes sets TokenLifetimeInMinutes field to given value.

### HasTokenLifetimeInMinutes

`func (o *AuthenticatorSettings) HasTokenLifetimeInMinutes() bool`

HasTokenLifetimeInMinutes returns a boolean if a field has been set.

### GetUserVerification

`func (o *AuthenticatorSettings) GetUserVerification() UserVerificationEnum`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *AuthenticatorSettings) GetUserVerificationOk() (*UserVerificationEnum, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *AuthenticatorSettings) SetUserVerification(v UserVerificationEnum)`

SetUserVerification sets UserVerification field to given value.

### HasUserVerification

`func (o *AuthenticatorSettings) HasUserVerification() bool`

HasUserVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


