# AuthenticatorKeyCustomAppAllOfSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserVerification** | Pointer to **string** | User verification setting | [optional] 
**AppInstanceId** | Pointer to **string** | The application instance ID. For custom_app, you need to create an OIDC native app using the [Apps API](https://developer.okta.com/docs/reference/api/apps/) with &#x60;Authorization Code&#x60; and &#x60;Refresh Token&#x60; grant types. You can leave both &#x60;Sign-in redirect URIs&#x60; and &#x60;Sign-out redirect URIs&#x60; as the default values. | [optional] 

## Methods

### NewAuthenticatorKeyCustomAppAllOfSettings

`func NewAuthenticatorKeyCustomAppAllOfSettings() *AuthenticatorKeyCustomAppAllOfSettings`

NewAuthenticatorKeyCustomAppAllOfSettings instantiates a new AuthenticatorKeyCustomAppAllOfSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorKeyCustomAppAllOfSettingsWithDefaults

`func NewAuthenticatorKeyCustomAppAllOfSettingsWithDefaults() *AuthenticatorKeyCustomAppAllOfSettings`

NewAuthenticatorKeyCustomAppAllOfSettingsWithDefaults instantiates a new AuthenticatorKeyCustomAppAllOfSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserVerification

`func (o *AuthenticatorKeyCustomAppAllOfSettings) GetUserVerification() string`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *AuthenticatorKeyCustomAppAllOfSettings) GetUserVerificationOk() (*string, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *AuthenticatorKeyCustomAppAllOfSettings) SetUserVerification(v string)`

SetUserVerification sets UserVerification field to given value.

### HasUserVerification

`func (o *AuthenticatorKeyCustomAppAllOfSettings) HasUserVerification() bool`

HasUserVerification returns a boolean if a field has been set.

### GetAppInstanceId

`func (o *AuthenticatorKeyCustomAppAllOfSettings) GetAppInstanceId() string`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *AuthenticatorKeyCustomAppAllOfSettings) GetAppInstanceIdOk() (*string, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *AuthenticatorKeyCustomAppAllOfSettings) SetAppInstanceId(v string)`

SetAppInstanceId sets AppInstanceId field to given value.

### HasAppInstanceId

`func (o *AuthenticatorKeyCustomAppAllOfSettings) HasAppInstanceId() bool`

HasAppInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


