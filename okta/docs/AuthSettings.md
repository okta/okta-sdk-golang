# AuthSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **string** |  | 
**CustomSettings** | Pointer to [**CustomAuthSettings**](CustomAuthSettings.md) |  | [optional] 
**OAuth2Settings** | Pointer to [**OAuth2Settings**](OAuth2Settings.md) |  | [optional] 

## Methods

### NewAuthSettings

`func NewAuthSettings(authType string, ) *AuthSettings`

NewAuthSettings instantiates a new AuthSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthSettingsWithDefaults

`func NewAuthSettingsWithDefaults() *AuthSettings`

NewAuthSettingsWithDefaults instantiates a new AuthSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *AuthSettings) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AuthSettings) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AuthSettings) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetCustomSettings

`func (o *AuthSettings) GetCustomSettings() CustomAuthSettings`

GetCustomSettings returns the CustomSettings field if non-nil, zero value otherwise.

### GetCustomSettingsOk

`func (o *AuthSettings) GetCustomSettingsOk() (*CustomAuthSettings, bool)`

GetCustomSettingsOk returns a tuple with the CustomSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSettings

`func (o *AuthSettings) SetCustomSettings(v CustomAuthSettings)`

SetCustomSettings sets CustomSettings field to given value.

### HasCustomSettings

`func (o *AuthSettings) HasCustomSettings() bool`

HasCustomSettings returns a boolean if a field has been set.

### GetOAuth2Settings

`func (o *AuthSettings) GetOAuth2Settings() OAuth2Settings`

GetOAuth2Settings returns the OAuth2Settings field if non-nil, zero value otherwise.

### GetOAuth2SettingsOk

`func (o *AuthSettings) GetOAuth2SettingsOk() (*OAuth2Settings, bool)`

GetOAuth2SettingsOk returns a tuple with the OAuth2Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuth2Settings

`func (o *AuthSettings) SetOAuth2Settings(v OAuth2Settings)`

SetOAuth2Settings sets OAuth2Settings field to given value.

### HasOAuth2Settings

`func (o *AuthSettings) HasOAuth2Settings() bool`

HasOAuth2Settings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


