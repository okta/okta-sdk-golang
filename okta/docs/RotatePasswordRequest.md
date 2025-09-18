# RotatePasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | The password associated with the privileged resource | 
**SecretVersionId** | **string** | The version ID of the password secret from the OPA vault | 

## Methods

### NewRotatePasswordRequest

`func NewRotatePasswordRequest(password string, secretVersionId string, ) *RotatePasswordRequest`

NewRotatePasswordRequest instantiates a new RotatePasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatePasswordRequestWithDefaults

`func NewRotatePasswordRequestWithDefaults() *RotatePasswordRequest`

NewRotatePasswordRequestWithDefaults instantiates a new RotatePasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *RotatePasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RotatePasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RotatePasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSecretVersionId

`func (o *RotatePasswordRequest) GetSecretVersionId() string`

GetSecretVersionId returns the SecretVersionId field if non-nil, zero value otherwise.

### GetSecretVersionIdOk

`func (o *RotatePasswordRequest) GetSecretVersionIdOk() (*string, bool)`

GetSecretVersionIdOk returns a tuple with the SecretVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVersionId

`func (o *RotatePasswordRequest) SetSecretVersionId(v string)`

SetSecretVersionId sets SecretVersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


