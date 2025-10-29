# UserActivationToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationToken** | Pointer to **string** | Token received as part of an activation user request. If a password was set before the user was activated, then user must sign in with their password or the &#x60;activationToken&#x60; and not the activation link. More information about using the &#x60;activationToken&#x60; to login can be found in the [Authentication API](https://developer.okta.com/docs/reference/api/authn/#primary-authentication-with-activation-token). | [optional] [readonly] 
**ActivationUrl** | Pointer to **string** | If &#x60;sendEmail&#x60; is &#x60;false&#x60;, returns an activation link for the user to set up their account. The activation token can be used to create a custom activation link. | [optional] [readonly] 

## Methods

### NewUserActivationToken

`func NewUserActivationToken() *UserActivationToken`

NewUserActivationToken instantiates a new UserActivationToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivationTokenWithDefaults

`func NewUserActivationTokenWithDefaults() *UserActivationToken`

NewUserActivationTokenWithDefaults instantiates a new UserActivationToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationToken

`func (o *UserActivationToken) GetActivationToken() string`

GetActivationToken returns the ActivationToken field if non-nil, zero value otherwise.

### GetActivationTokenOk

`func (o *UserActivationToken) GetActivationTokenOk() (*string, bool)`

GetActivationTokenOk returns a tuple with the ActivationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationToken

`func (o *UserActivationToken) SetActivationToken(v string)`

SetActivationToken sets ActivationToken field to given value.

### HasActivationToken

`func (o *UserActivationToken) HasActivationToken() bool`

HasActivationToken returns a boolean if a field has been set.

### GetActivationUrl

`func (o *UserActivationToken) GetActivationUrl() string`

GetActivationUrl returns the ActivationUrl field if non-nil, zero value otherwise.

### GetActivationUrlOk

`func (o *UserActivationToken) GetActivationUrlOk() (*string, bool)`

GetActivationUrlOk returns a tuple with the ActivationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationUrl

`func (o *UserActivationToken) SetActivationUrl(v string)`

SetActivationUrl sets ActivationUrl field to given value.

### HasActivationUrl

`func (o *UserActivationToken) HasActivationUrl() bool`

HasActivationUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


