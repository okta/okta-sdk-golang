# RegistrationGrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** | Type of registration grant | 
**RegistrationSecretId** | **string** | ID of the registration secret (from &#x60;POST /device-identity/api/v1/registration-secrets&#x60;) | 

## Methods

### NewRegistrationGrantRequest

`func NewRegistrationGrantRequest(grantType string, registrationSecretId string, ) *RegistrationGrantRequest`

NewRegistrationGrantRequest instantiates a new RegistrationGrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationGrantRequestWithDefaults

`func NewRegistrationGrantRequestWithDefaults() *RegistrationGrantRequest`

NewRegistrationGrantRequestWithDefaults instantiates a new RegistrationGrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *RegistrationGrantRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *RegistrationGrantRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *RegistrationGrantRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetRegistrationSecretId

`func (o *RegistrationGrantRequest) GetRegistrationSecretId() string`

GetRegistrationSecretId returns the RegistrationSecretId field if non-nil, zero value otherwise.

### GetRegistrationSecretIdOk

`func (o *RegistrationGrantRequest) GetRegistrationSecretIdOk() (*string, bool)`

GetRegistrationSecretIdOk returns a tuple with the RegistrationSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationSecretId

`func (o *RegistrationGrantRequest) SetRegistrationSecretId(v string)`

SetRegistrationSecretId sets RegistrationSecretId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


