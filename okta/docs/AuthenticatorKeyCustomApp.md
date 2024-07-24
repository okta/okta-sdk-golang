# AuthenticatorKeyCustomApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreeToTerms** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that the administrator accepts the [terms](https://www.okta.com/privacy-policy/)for creating a new authenticator. Okta requires that you accept the terms when creating a new &#x60;custom_app&#x60; authenticator. Other authenticators don&#39;t require this field. | [optional] 
**Provider** | Pointer to [**AuthenticatorKeyCustomAppAllOfProvider**](AuthenticatorKeyCustomAppAllOfProvider.md) |  | [optional] 
**Settings** | Pointer to [**AuthenticatorKeyCustomAppAllOfSettings**](AuthenticatorKeyCustomAppAllOfSettings.md) |  | [optional] 

## Methods

### NewAuthenticatorKeyCustomApp

`func NewAuthenticatorKeyCustomApp() *AuthenticatorKeyCustomApp`

NewAuthenticatorKeyCustomApp instantiates a new AuthenticatorKeyCustomApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorKeyCustomAppWithDefaults

`func NewAuthenticatorKeyCustomAppWithDefaults() *AuthenticatorKeyCustomApp`

NewAuthenticatorKeyCustomAppWithDefaults instantiates a new AuthenticatorKeyCustomApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreeToTerms

`func (o *AuthenticatorKeyCustomApp) GetAgreeToTerms() bool`

GetAgreeToTerms returns the AgreeToTerms field if non-nil, zero value otherwise.

### GetAgreeToTermsOk

`func (o *AuthenticatorKeyCustomApp) GetAgreeToTermsOk() (*bool, bool)`

GetAgreeToTermsOk returns a tuple with the AgreeToTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreeToTerms

`func (o *AuthenticatorKeyCustomApp) SetAgreeToTerms(v bool)`

SetAgreeToTerms sets AgreeToTerms field to given value.

### HasAgreeToTerms

`func (o *AuthenticatorKeyCustomApp) HasAgreeToTerms() bool`

HasAgreeToTerms returns a boolean if a field has been set.

### GetProvider

`func (o *AuthenticatorKeyCustomApp) GetProvider() AuthenticatorKeyCustomAppAllOfProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AuthenticatorKeyCustomApp) GetProviderOk() (*AuthenticatorKeyCustomAppAllOfProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AuthenticatorKeyCustomApp) SetProvider(v AuthenticatorKeyCustomAppAllOfProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AuthenticatorKeyCustomApp) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSettings

`func (o *AuthenticatorKeyCustomApp) GetSettings() AuthenticatorKeyCustomAppAllOfSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AuthenticatorKeyCustomApp) GetSettingsOk() (*AuthenticatorKeyCustomAppAllOfSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AuthenticatorKeyCustomApp) SetSettings(v AuthenticatorKeyCustomAppAllOfSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *AuthenticatorKeyCustomApp) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


