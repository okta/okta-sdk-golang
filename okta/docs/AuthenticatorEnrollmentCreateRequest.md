# AuthenticatorEnrollmentCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatorId** | **string** | Unique identifier of the &#x60;phone&#x60; authenticator | 
**Profile** | [**AuthenticatorProfile**](AuthenticatorProfile.md) |  | 

## Methods

### NewAuthenticatorEnrollmentCreateRequest

`func NewAuthenticatorEnrollmentCreateRequest(authenticatorId string, profile AuthenticatorProfile, ) *AuthenticatorEnrollmentCreateRequest`

NewAuthenticatorEnrollmentCreateRequest instantiates a new AuthenticatorEnrollmentCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorEnrollmentCreateRequestWithDefaults

`func NewAuthenticatorEnrollmentCreateRequestWithDefaults() *AuthenticatorEnrollmentCreateRequest`

NewAuthenticatorEnrollmentCreateRequestWithDefaults instantiates a new AuthenticatorEnrollmentCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticatorId

`func (o *AuthenticatorEnrollmentCreateRequest) GetAuthenticatorId() string`

GetAuthenticatorId returns the AuthenticatorId field if non-nil, zero value otherwise.

### GetAuthenticatorIdOk

`func (o *AuthenticatorEnrollmentCreateRequest) GetAuthenticatorIdOk() (*string, bool)`

GetAuthenticatorIdOk returns a tuple with the AuthenticatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorId

`func (o *AuthenticatorEnrollmentCreateRequest) SetAuthenticatorId(v string)`

SetAuthenticatorId sets AuthenticatorId field to given value.


### GetProfile

`func (o *AuthenticatorEnrollmentCreateRequest) GetProfile() AuthenticatorProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AuthenticatorEnrollmentCreateRequest) GetProfileOk() (*AuthenticatorProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AuthenticatorEnrollmentCreateRequest) SetProfile(v AuthenticatorProfile)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


