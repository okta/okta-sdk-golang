# AuthenticatorKeyEmailAllOfSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedFor** | Pointer to **string** | The allowed types of uses for the Authenticator | [optional] 
**TokenLifetimeInMinutes** | Pointer to **float32** | Specifies the lifetime of an email token. Default value is 5 minutes. | [optional] [default to 5]

## Methods

### NewAuthenticatorKeyEmailAllOfSettings

`func NewAuthenticatorKeyEmailAllOfSettings() *AuthenticatorKeyEmailAllOfSettings`

NewAuthenticatorKeyEmailAllOfSettings instantiates a new AuthenticatorKeyEmailAllOfSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorKeyEmailAllOfSettingsWithDefaults

`func NewAuthenticatorKeyEmailAllOfSettingsWithDefaults() *AuthenticatorKeyEmailAllOfSettings`

NewAuthenticatorKeyEmailAllOfSettingsWithDefaults instantiates a new AuthenticatorKeyEmailAllOfSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedFor

`func (o *AuthenticatorKeyEmailAllOfSettings) GetAllowedFor() string`

GetAllowedFor returns the AllowedFor field if non-nil, zero value otherwise.

### GetAllowedForOk

`func (o *AuthenticatorKeyEmailAllOfSettings) GetAllowedForOk() (*string, bool)`

GetAllowedForOk returns a tuple with the AllowedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedFor

`func (o *AuthenticatorKeyEmailAllOfSettings) SetAllowedFor(v string)`

SetAllowedFor sets AllowedFor field to given value.

### HasAllowedFor

`func (o *AuthenticatorKeyEmailAllOfSettings) HasAllowedFor() bool`

HasAllowedFor returns a boolean if a field has been set.

### GetTokenLifetimeInMinutes

`func (o *AuthenticatorKeyEmailAllOfSettings) GetTokenLifetimeInMinutes() float32`

GetTokenLifetimeInMinutes returns the TokenLifetimeInMinutes field if non-nil, zero value otherwise.

### GetTokenLifetimeInMinutesOk

`func (o *AuthenticatorKeyEmailAllOfSettings) GetTokenLifetimeInMinutesOk() (*float32, bool)`

GetTokenLifetimeInMinutesOk returns a tuple with the TokenLifetimeInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLifetimeInMinutes

`func (o *AuthenticatorKeyEmailAllOfSettings) SetTokenLifetimeInMinutes(v float32)`

SetTokenLifetimeInMinutes sets TokenLifetimeInMinutes field to given value.

### HasTokenLifetimeInMinutes

`func (o *AuthenticatorKeyEmailAllOfSettings) HasTokenLifetimeInMinutes() bool`

HasTokenLifetimeInMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


