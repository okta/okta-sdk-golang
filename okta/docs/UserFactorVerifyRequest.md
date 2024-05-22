# UserFactorVerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationToken** | Pointer to **string** |  | [optional] 
**Answer** | Pointer to **string** | Answer to the question | [optional] 
**Attestation** | Pointer to **string** | Base64-encoded attestation from the WebAuthn JavaScript call | [optional] 
**ClientData** | Pointer to **string** | Base64-encoded client data from the WebAuthn authenticator | [optional] 
**NextPassCode** | Pointer to **int32** | OTP for the next time window | [optional] 
**PassCode** | Pointer to **string** | OTP for the current time window | [optional] 
**RegistrationData** | Pointer to **string** | Base64-encoded registration data from the U2F JavaScript call | [optional] 
**StateToken** | Pointer to **string** |  | [optional] 

## Methods

### NewUserFactorVerifyRequest

`func NewUserFactorVerifyRequest() *UserFactorVerifyRequest`

NewUserFactorVerifyRequest instantiates a new UserFactorVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorVerifyRequestWithDefaults

`func NewUserFactorVerifyRequestWithDefaults() *UserFactorVerifyRequest`

NewUserFactorVerifyRequestWithDefaults instantiates a new UserFactorVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationToken

`func (o *UserFactorVerifyRequest) GetActivationToken() string`

GetActivationToken returns the ActivationToken field if non-nil, zero value otherwise.

### GetActivationTokenOk

`func (o *UserFactorVerifyRequest) GetActivationTokenOk() (*string, bool)`

GetActivationTokenOk returns a tuple with the ActivationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationToken

`func (o *UserFactorVerifyRequest) SetActivationToken(v string)`

SetActivationToken sets ActivationToken field to given value.

### HasActivationToken

`func (o *UserFactorVerifyRequest) HasActivationToken() bool`

HasActivationToken returns a boolean if a field has been set.

### GetAnswer

`func (o *UserFactorVerifyRequest) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *UserFactorVerifyRequest) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *UserFactorVerifyRequest) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *UserFactorVerifyRequest) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetAttestation

`func (o *UserFactorVerifyRequest) GetAttestation() string`

GetAttestation returns the Attestation field if non-nil, zero value otherwise.

### GetAttestationOk

`func (o *UserFactorVerifyRequest) GetAttestationOk() (*string, bool)`

GetAttestationOk returns a tuple with the Attestation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestation

`func (o *UserFactorVerifyRequest) SetAttestation(v string)`

SetAttestation sets Attestation field to given value.

### HasAttestation

`func (o *UserFactorVerifyRequest) HasAttestation() bool`

HasAttestation returns a boolean if a field has been set.

### GetClientData

`func (o *UserFactorVerifyRequest) GetClientData() string`

GetClientData returns the ClientData field if non-nil, zero value otherwise.

### GetClientDataOk

`func (o *UserFactorVerifyRequest) GetClientDataOk() (*string, bool)`

GetClientDataOk returns a tuple with the ClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientData

`func (o *UserFactorVerifyRequest) SetClientData(v string)`

SetClientData sets ClientData field to given value.

### HasClientData

`func (o *UserFactorVerifyRequest) HasClientData() bool`

HasClientData returns a boolean if a field has been set.

### GetNextPassCode

`func (o *UserFactorVerifyRequest) GetNextPassCode() int32`

GetNextPassCode returns the NextPassCode field if non-nil, zero value otherwise.

### GetNextPassCodeOk

`func (o *UserFactorVerifyRequest) GetNextPassCodeOk() (*int32, bool)`

GetNextPassCodeOk returns a tuple with the NextPassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPassCode

`func (o *UserFactorVerifyRequest) SetNextPassCode(v int32)`

SetNextPassCode sets NextPassCode field to given value.

### HasNextPassCode

`func (o *UserFactorVerifyRequest) HasNextPassCode() bool`

HasNextPassCode returns a boolean if a field has been set.

### GetPassCode

`func (o *UserFactorVerifyRequest) GetPassCode() string`

GetPassCode returns the PassCode field if non-nil, zero value otherwise.

### GetPassCodeOk

`func (o *UserFactorVerifyRequest) GetPassCodeOk() (*string, bool)`

GetPassCodeOk returns a tuple with the PassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCode

`func (o *UserFactorVerifyRequest) SetPassCode(v string)`

SetPassCode sets PassCode field to given value.

### HasPassCode

`func (o *UserFactorVerifyRequest) HasPassCode() bool`

HasPassCode returns a boolean if a field has been set.

### GetRegistrationData

`func (o *UserFactorVerifyRequest) GetRegistrationData() string`

GetRegistrationData returns the RegistrationData field if non-nil, zero value otherwise.

### GetRegistrationDataOk

`func (o *UserFactorVerifyRequest) GetRegistrationDataOk() (*string, bool)`

GetRegistrationDataOk returns a tuple with the RegistrationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationData

`func (o *UserFactorVerifyRequest) SetRegistrationData(v string)`

SetRegistrationData sets RegistrationData field to given value.

### HasRegistrationData

`func (o *UserFactorVerifyRequest) HasRegistrationData() bool`

HasRegistrationData returns a boolean if a field has been set.

### GetStateToken

`func (o *UserFactorVerifyRequest) GetStateToken() string`

GetStateToken returns the StateToken field if non-nil, zero value otherwise.

### GetStateTokenOk

`func (o *UserFactorVerifyRequest) GetStateTokenOk() (*string, bool)`

GetStateTokenOk returns a tuple with the StateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateToken

`func (o *UserFactorVerifyRequest) SetStateToken(v string)`

SetStateToken sets StateToken field to given value.

### HasStateToken

`func (o *UserFactorVerifyRequest) HasStateToken() bool`

HasStateToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


