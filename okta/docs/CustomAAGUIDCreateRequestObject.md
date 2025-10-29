# CustomAAGUIDCreateRequestObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aaguid** | Pointer to **string** | An Authenticator Attestation Global Unique Identifier (AAGUID) is a 128-bit identifier indicating the model. | [optional] 
**AttestationRootCertificates** | Pointer to [**[]AttestationRootCertificatesRequestInner**](AttestationRootCertificatesRequestInner.md) | Contains the certificate and information about it | [optional] 
**AuthenticatorCharacteristics** | Pointer to [**AAGUIDAuthenticatorCharacteristics**](AAGUIDAuthenticatorCharacteristics.md) |  | [optional] 

## Methods

### NewCustomAAGUIDCreateRequestObject

`func NewCustomAAGUIDCreateRequestObject() *CustomAAGUIDCreateRequestObject`

NewCustomAAGUIDCreateRequestObject instantiates a new CustomAAGUIDCreateRequestObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAAGUIDCreateRequestObjectWithDefaults

`func NewCustomAAGUIDCreateRequestObjectWithDefaults() *CustomAAGUIDCreateRequestObject`

NewCustomAAGUIDCreateRequestObjectWithDefaults instantiates a new CustomAAGUIDCreateRequestObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAaguid

`func (o *CustomAAGUIDCreateRequestObject) GetAaguid() string`

GetAaguid returns the Aaguid field if non-nil, zero value otherwise.

### GetAaguidOk

`func (o *CustomAAGUIDCreateRequestObject) GetAaguidOk() (*string, bool)`

GetAaguidOk returns a tuple with the Aaguid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaguid

`func (o *CustomAAGUIDCreateRequestObject) SetAaguid(v string)`

SetAaguid sets Aaguid field to given value.

### HasAaguid

`func (o *CustomAAGUIDCreateRequestObject) HasAaguid() bool`

HasAaguid returns a boolean if a field has been set.

### GetAttestationRootCertificates

`func (o *CustomAAGUIDCreateRequestObject) GetAttestationRootCertificates() []AttestationRootCertificatesRequestInner`

GetAttestationRootCertificates returns the AttestationRootCertificates field if non-nil, zero value otherwise.

### GetAttestationRootCertificatesOk

`func (o *CustomAAGUIDCreateRequestObject) GetAttestationRootCertificatesOk() (*[]AttestationRootCertificatesRequestInner, bool)`

GetAttestationRootCertificatesOk returns a tuple with the AttestationRootCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationRootCertificates

`func (o *CustomAAGUIDCreateRequestObject) SetAttestationRootCertificates(v []AttestationRootCertificatesRequestInner)`

SetAttestationRootCertificates sets AttestationRootCertificates field to given value.

### HasAttestationRootCertificates

`func (o *CustomAAGUIDCreateRequestObject) HasAttestationRootCertificates() bool`

HasAttestationRootCertificates returns a boolean if a field has been set.

### GetAuthenticatorCharacteristics

`func (o *CustomAAGUIDCreateRequestObject) GetAuthenticatorCharacteristics() AAGUIDAuthenticatorCharacteristics`

GetAuthenticatorCharacteristics returns the AuthenticatorCharacteristics field if non-nil, zero value otherwise.

### GetAuthenticatorCharacteristicsOk

`func (o *CustomAAGUIDCreateRequestObject) GetAuthenticatorCharacteristicsOk() (*AAGUIDAuthenticatorCharacteristics, bool)`

GetAuthenticatorCharacteristicsOk returns a tuple with the AuthenticatorCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorCharacteristics

`func (o *CustomAAGUIDCreateRequestObject) SetAuthenticatorCharacteristics(v AAGUIDAuthenticatorCharacteristics)`

SetAuthenticatorCharacteristics sets AuthenticatorCharacteristics field to given value.

### HasAuthenticatorCharacteristics

`func (o *CustomAAGUIDCreateRequestObject) HasAuthenticatorCharacteristics() bool`

HasAuthenticatorCharacteristics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


