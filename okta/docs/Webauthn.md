# Webauthn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attestation** | Pointer to **string** | Base64-encoded attestation from the WebAuthn authenticator | [optional] 
**ClientData** | Pointer to **string** | Base64-encoded client data from the WebAuthn authenticator | [optional] 

## Methods

### NewWebauthn

`func NewWebauthn() *Webauthn`

NewWebauthn instantiates a new Webauthn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebauthnWithDefaults

`func NewWebauthnWithDefaults() *Webauthn`

NewWebauthnWithDefaults instantiates a new Webauthn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttestation

`func (o *Webauthn) GetAttestation() string`

GetAttestation returns the Attestation field if non-nil, zero value otherwise.

### GetAttestationOk

`func (o *Webauthn) GetAttestationOk() (*string, bool)`

GetAttestationOk returns a tuple with the Attestation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestation

`func (o *Webauthn) SetAttestation(v string)`

SetAttestation sets Attestation field to given value.

### HasAttestation

`func (o *Webauthn) HasAttestation() bool`

HasAttestation returns a boolean if a field has been set.

### GetClientData

`func (o *Webauthn) GetClientData() string`

GetClientData returns the ClientData field if non-nil, zero value otherwise.

### GetClientDataOk

`func (o *Webauthn) GetClientDataOk() (*string, bool)`

GetClientDataOk returns a tuple with the ClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientData

`func (o *Webauthn) SetClientData(v string)`

SetClientData sets ClientData field to given value.

### HasClientData

`func (o *Webauthn) HasClientData() bool`

HasClientData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


