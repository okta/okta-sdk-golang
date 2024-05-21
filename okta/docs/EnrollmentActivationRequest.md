# EnrollmentActivationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredResponses** | Pointer to [**[]WebAuthnCredResponse**](WebAuthnCredResponse.md) | List of credential responses from the fulfillment provider | [optional] 
**FulfillmentProvider** | Pointer to **string** | Name of the fulfillment provider for the WebAuthn Preregistration Factor | [optional] 
**PinResponseJwe** | Pointer to **string** | Encrypted JWE of PIN response from the fulfillment provider | [optional] 
**Serial** | Pointer to **string** | Serial number of the YubiKey | [optional] 
**UserId** | Pointer to **string** | ID of an existing Okta user | [optional] 
**Version** | Pointer to **string** | Firmware version of the YubiKey | [optional] 
**YubicoSigningJwks** | Pointer to [**[]ECKeyJWK**](ECKeyJWK.md) | List of usable signing keys from Yubico (in JWKS format) used to verify the JWS inside the JWE | [optional] 

## Methods

### NewEnrollmentActivationRequest

`func NewEnrollmentActivationRequest() *EnrollmentActivationRequest`

NewEnrollmentActivationRequest instantiates a new EnrollmentActivationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentActivationRequestWithDefaults

`func NewEnrollmentActivationRequestWithDefaults() *EnrollmentActivationRequest`

NewEnrollmentActivationRequestWithDefaults instantiates a new EnrollmentActivationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredResponses

`func (o *EnrollmentActivationRequest) GetCredResponses() []WebAuthnCredResponse`

GetCredResponses returns the CredResponses field if non-nil, zero value otherwise.

### GetCredResponsesOk

`func (o *EnrollmentActivationRequest) GetCredResponsesOk() (*[]WebAuthnCredResponse, bool)`

GetCredResponsesOk returns a tuple with the CredResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredResponses

`func (o *EnrollmentActivationRequest) SetCredResponses(v []WebAuthnCredResponse)`

SetCredResponses sets CredResponses field to given value.

### HasCredResponses

`func (o *EnrollmentActivationRequest) HasCredResponses() bool`

HasCredResponses returns a boolean if a field has been set.

### GetFulfillmentProvider

`func (o *EnrollmentActivationRequest) GetFulfillmentProvider() string`

GetFulfillmentProvider returns the FulfillmentProvider field if non-nil, zero value otherwise.

### GetFulfillmentProviderOk

`func (o *EnrollmentActivationRequest) GetFulfillmentProviderOk() (*string, bool)`

GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentProvider

`func (o *EnrollmentActivationRequest) SetFulfillmentProvider(v string)`

SetFulfillmentProvider sets FulfillmentProvider field to given value.

### HasFulfillmentProvider

`func (o *EnrollmentActivationRequest) HasFulfillmentProvider() bool`

HasFulfillmentProvider returns a boolean if a field has been set.

### GetPinResponseJwe

`func (o *EnrollmentActivationRequest) GetPinResponseJwe() string`

GetPinResponseJwe returns the PinResponseJwe field if non-nil, zero value otherwise.

### GetPinResponseJweOk

`func (o *EnrollmentActivationRequest) GetPinResponseJweOk() (*string, bool)`

GetPinResponseJweOk returns a tuple with the PinResponseJwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinResponseJwe

`func (o *EnrollmentActivationRequest) SetPinResponseJwe(v string)`

SetPinResponseJwe sets PinResponseJwe field to given value.

### HasPinResponseJwe

`func (o *EnrollmentActivationRequest) HasPinResponseJwe() bool`

HasPinResponseJwe returns a boolean if a field has been set.

### GetSerial

`func (o *EnrollmentActivationRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *EnrollmentActivationRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *EnrollmentActivationRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *EnrollmentActivationRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetUserId

`func (o *EnrollmentActivationRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EnrollmentActivationRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EnrollmentActivationRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EnrollmentActivationRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVersion

`func (o *EnrollmentActivationRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EnrollmentActivationRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EnrollmentActivationRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EnrollmentActivationRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetYubicoSigningJwks

`func (o *EnrollmentActivationRequest) GetYubicoSigningJwks() []ECKeyJWK`

GetYubicoSigningJwks returns the YubicoSigningJwks field if non-nil, zero value otherwise.

### GetYubicoSigningJwksOk

`func (o *EnrollmentActivationRequest) GetYubicoSigningJwksOk() (*[]ECKeyJWK, bool)`

GetYubicoSigningJwksOk returns a tuple with the YubicoSigningJwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYubicoSigningJwks

`func (o *EnrollmentActivationRequest) SetYubicoSigningJwks(v []ECKeyJWK)`

SetYubicoSigningJwks sets YubicoSigningJwks field to given value.

### HasYubicoSigningJwks

`func (o *EnrollmentActivationRequest) HasYubicoSigningJwks() bool`

HasYubicoSigningJwks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


