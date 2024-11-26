# UserFactorVerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PassCode** | Pointer to **string** | OTP for the current time window | [optional] 
**Answer** | Pointer to **string** | Answer to the question | [optional] 
**ClientData** | Pointer to **string** | Base64-encoded client data from the WebAuthn authenticator | [optional] 
**SignatureData** | Pointer to **string** | Base64-encoded signature data from the WebAuthn authenticator | [optional] 
**AuthenticatorData** | Pointer to **string** | Base64-encoded authenticator data from the WebAuthn authenticator | [optional] 

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

### GetSignatureData

`func (o *UserFactorVerifyRequest) GetSignatureData() string`

GetSignatureData returns the SignatureData field if non-nil, zero value otherwise.

### GetSignatureDataOk

`func (o *UserFactorVerifyRequest) GetSignatureDataOk() (*string, bool)`

GetSignatureDataOk returns a tuple with the SignatureData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureData

`func (o *UserFactorVerifyRequest) SetSignatureData(v string)`

SetSignatureData sets SignatureData field to given value.

### HasSignatureData

`func (o *UserFactorVerifyRequest) HasSignatureData() bool`

HasSignatureData returns a boolean if a field has been set.

### GetAuthenticatorData

`func (o *UserFactorVerifyRequest) GetAuthenticatorData() string`

GetAuthenticatorData returns the AuthenticatorData field if non-nil, zero value otherwise.

### GetAuthenticatorDataOk

`func (o *UserFactorVerifyRequest) GetAuthenticatorDataOk() (*string, bool)`

GetAuthenticatorDataOk returns a tuple with the AuthenticatorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorData

`func (o *UserFactorVerifyRequest) SetAuthenticatorData(v string)`

SetAuthenticatorData sets AuthenticatorData field to given value.

### HasAuthenticatorData

`func (o *UserFactorVerifyRequest) HasAuthenticatorData() bool`

HasAuthenticatorData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


