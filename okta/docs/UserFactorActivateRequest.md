# UserFactorActivateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PassCode** | Pointer to **string** | OTP for the current time window | [optional] 
**ClientData** | Pointer to **string** | Base64-encoded client data from the WebAuthn authenticator | [optional] 
**RegistrationData** | Pointer to **string** | Base64-encoded registration data from the U2F token | [optional] 
**Attestation** | Pointer to **string** | Base64-encoded attestation from the WebAuthn authenticator | [optional] 

## Methods

### NewUserFactorActivateRequest

`func NewUserFactorActivateRequest() *UserFactorActivateRequest`

NewUserFactorActivateRequest instantiates a new UserFactorActivateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorActivateRequestWithDefaults

`func NewUserFactorActivateRequestWithDefaults() *UserFactorActivateRequest`

NewUserFactorActivateRequestWithDefaults instantiates a new UserFactorActivateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassCode

`func (o *UserFactorActivateRequest) GetPassCode() string`

GetPassCode returns the PassCode field if non-nil, zero value otherwise.

### GetPassCodeOk

`func (o *UserFactorActivateRequest) GetPassCodeOk() (*string, bool)`

GetPassCodeOk returns a tuple with the PassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCode

`func (o *UserFactorActivateRequest) SetPassCode(v string)`

SetPassCode sets PassCode field to given value.

### HasPassCode

`func (o *UserFactorActivateRequest) HasPassCode() bool`

HasPassCode returns a boolean if a field has been set.

### GetClientData

`func (o *UserFactorActivateRequest) GetClientData() string`

GetClientData returns the ClientData field if non-nil, zero value otherwise.

### GetClientDataOk

`func (o *UserFactorActivateRequest) GetClientDataOk() (*string, bool)`

GetClientDataOk returns a tuple with the ClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientData

`func (o *UserFactorActivateRequest) SetClientData(v string)`

SetClientData sets ClientData field to given value.

### HasClientData

`func (o *UserFactorActivateRequest) HasClientData() bool`

HasClientData returns a boolean if a field has been set.

### GetRegistrationData

`func (o *UserFactorActivateRequest) GetRegistrationData() string`

GetRegistrationData returns the RegistrationData field if non-nil, zero value otherwise.

### GetRegistrationDataOk

`func (o *UserFactorActivateRequest) GetRegistrationDataOk() (*string, bool)`

GetRegistrationDataOk returns a tuple with the RegistrationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationData

`func (o *UserFactorActivateRequest) SetRegistrationData(v string)`

SetRegistrationData sets RegistrationData field to given value.

### HasRegistrationData

`func (o *UserFactorActivateRequest) HasRegistrationData() bool`

HasRegistrationData returns a boolean if a field has been set.

### GetAttestation

`func (o *UserFactorActivateRequest) GetAttestation() string`

GetAttestation returns the Attestation field if non-nil, zero value otherwise.

### GetAttestationOk

`func (o *UserFactorActivateRequest) GetAttestationOk() (*string, bool)`

GetAttestationOk returns a tuple with the Attestation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestation

`func (o *UserFactorActivateRequest) SetAttestation(v string)`

SetAttestation sets Attestation field to given value.

### HasAttestation

`func (o *UserFactorActivateRequest) HasAttestation() bool`

HasAttestation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


