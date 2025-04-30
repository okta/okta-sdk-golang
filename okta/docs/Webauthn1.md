# Webauthn1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatorData** | Pointer to **string** | Base64-encoded authenticator data from the WebAuthn authenticator | [optional] 
**ClientData** | Pointer to **string** | Base64-encoded client data from the WebAuthn authenticator | [optional] 
**SignatureData** | Pointer to **string** | Base64-encoded signature data from the WebAuthn authenticator | [optional] 

## Methods

### NewWebauthn1

`func NewWebauthn1() *Webauthn1`

NewWebauthn1 instantiates a new Webauthn1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebauthn1WithDefaults

`func NewWebauthn1WithDefaults() *Webauthn1`

NewWebauthn1WithDefaults instantiates a new Webauthn1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticatorData

`func (o *Webauthn1) GetAuthenticatorData() string`

GetAuthenticatorData returns the AuthenticatorData field if non-nil, zero value otherwise.

### GetAuthenticatorDataOk

`func (o *Webauthn1) GetAuthenticatorDataOk() (*string, bool)`

GetAuthenticatorDataOk returns a tuple with the AuthenticatorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorData

`func (o *Webauthn1) SetAuthenticatorData(v string)`

SetAuthenticatorData sets AuthenticatorData field to given value.

### HasAuthenticatorData

`func (o *Webauthn1) HasAuthenticatorData() bool`

HasAuthenticatorData returns a boolean if a field has been set.

### GetClientData

`func (o *Webauthn1) GetClientData() string`

GetClientData returns the ClientData field if non-nil, zero value otherwise.

### GetClientDataOk

`func (o *Webauthn1) GetClientDataOk() (*string, bool)`

GetClientDataOk returns a tuple with the ClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientData

`func (o *Webauthn1) SetClientData(v string)`

SetClientData sets ClientData field to given value.

### HasClientData

`func (o *Webauthn1) HasClientData() bool`

HasClientData returns a boolean if a field has been set.

### GetSignatureData

`func (o *Webauthn1) GetSignatureData() string`

GetSignatureData returns the SignatureData field if non-nil, zero value otherwise.

### GetSignatureDataOk

`func (o *Webauthn1) GetSignatureDataOk() (*string, bool)`

GetSignatureDataOk returns a tuple with the SignatureData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureData

`func (o *Webauthn1) SetSignatureData(v string)`

SetSignatureData sets SignatureData field to given value.

### HasSignatureData

`func (o *Webauthn1) HasSignatureData() bool`

HasSignatureData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


