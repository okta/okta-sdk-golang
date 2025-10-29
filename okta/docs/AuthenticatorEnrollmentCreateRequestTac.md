# AuthenticatorEnrollmentCreateRequestTac

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatorId** | **string** | Unique identifier of the TAC authenticator | 
**Profile** | Pointer to [**AuthenticatorProfileTacRequest**](AuthenticatorProfileTacRequest.md) |  | [optional] 

## Methods

### NewAuthenticatorEnrollmentCreateRequestTac

`func NewAuthenticatorEnrollmentCreateRequestTac(authenticatorId string, ) *AuthenticatorEnrollmentCreateRequestTac`

NewAuthenticatorEnrollmentCreateRequestTac instantiates a new AuthenticatorEnrollmentCreateRequestTac object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorEnrollmentCreateRequestTacWithDefaults

`func NewAuthenticatorEnrollmentCreateRequestTacWithDefaults() *AuthenticatorEnrollmentCreateRequestTac`

NewAuthenticatorEnrollmentCreateRequestTacWithDefaults instantiates a new AuthenticatorEnrollmentCreateRequestTac object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticatorId

`func (o *AuthenticatorEnrollmentCreateRequestTac) GetAuthenticatorId() string`

GetAuthenticatorId returns the AuthenticatorId field if non-nil, zero value otherwise.

### GetAuthenticatorIdOk

`func (o *AuthenticatorEnrollmentCreateRequestTac) GetAuthenticatorIdOk() (*string, bool)`

GetAuthenticatorIdOk returns a tuple with the AuthenticatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorId

`func (o *AuthenticatorEnrollmentCreateRequestTac) SetAuthenticatorId(v string)`

SetAuthenticatorId sets AuthenticatorId field to given value.


### GetProfile

`func (o *AuthenticatorEnrollmentCreateRequestTac) GetProfile() AuthenticatorProfileTacRequest`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AuthenticatorEnrollmentCreateRequestTac) GetProfileOk() (*AuthenticatorProfileTacRequest, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AuthenticatorEnrollmentCreateRequestTac) SetProfile(v AuthenticatorProfileTacRequest)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AuthenticatorEnrollmentCreateRequestTac) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


