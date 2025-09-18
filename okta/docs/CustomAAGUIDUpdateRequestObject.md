# CustomAAGUIDUpdateRequestObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttestationRootCertificates** | Pointer to [**[]AttestationRootCertificatesRequestInner**](AttestationRootCertificatesRequestInner.md) | Contains the certificate and information about it | [optional] 
**AuthenticatorCharacteristics** | Pointer to [**AAGUIDAuthenticatorCharacteristics**](AAGUIDAuthenticatorCharacteristics.md) |  | [optional] 
**Name** | Pointer to **string** | The product name associated with this AAGUID. | [optional] 

## Methods

### NewCustomAAGUIDUpdateRequestObject

`func NewCustomAAGUIDUpdateRequestObject() *CustomAAGUIDUpdateRequestObject`

NewCustomAAGUIDUpdateRequestObject instantiates a new CustomAAGUIDUpdateRequestObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAAGUIDUpdateRequestObjectWithDefaults

`func NewCustomAAGUIDUpdateRequestObjectWithDefaults() *CustomAAGUIDUpdateRequestObject`

NewCustomAAGUIDUpdateRequestObjectWithDefaults instantiates a new CustomAAGUIDUpdateRequestObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttestationRootCertificates

`func (o *CustomAAGUIDUpdateRequestObject) GetAttestationRootCertificates() []AttestationRootCertificatesRequestInner`

GetAttestationRootCertificates returns the AttestationRootCertificates field if non-nil, zero value otherwise.

### GetAttestationRootCertificatesOk

`func (o *CustomAAGUIDUpdateRequestObject) GetAttestationRootCertificatesOk() (*[]AttestationRootCertificatesRequestInner, bool)`

GetAttestationRootCertificatesOk returns a tuple with the AttestationRootCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationRootCertificates

`func (o *CustomAAGUIDUpdateRequestObject) SetAttestationRootCertificates(v []AttestationRootCertificatesRequestInner)`

SetAttestationRootCertificates sets AttestationRootCertificates field to given value.

### HasAttestationRootCertificates

`func (o *CustomAAGUIDUpdateRequestObject) HasAttestationRootCertificates() bool`

HasAttestationRootCertificates returns a boolean if a field has been set.

### GetAuthenticatorCharacteristics

`func (o *CustomAAGUIDUpdateRequestObject) GetAuthenticatorCharacteristics() AAGUIDAuthenticatorCharacteristics`

GetAuthenticatorCharacteristics returns the AuthenticatorCharacteristics field if non-nil, zero value otherwise.

### GetAuthenticatorCharacteristicsOk

`func (o *CustomAAGUIDUpdateRequestObject) GetAuthenticatorCharacteristicsOk() (*AAGUIDAuthenticatorCharacteristics, bool)`

GetAuthenticatorCharacteristicsOk returns a tuple with the AuthenticatorCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorCharacteristics

`func (o *CustomAAGUIDUpdateRequestObject) SetAuthenticatorCharacteristics(v AAGUIDAuthenticatorCharacteristics)`

SetAuthenticatorCharacteristics sets AuthenticatorCharacteristics field to given value.

### HasAuthenticatorCharacteristics

`func (o *CustomAAGUIDUpdateRequestObject) HasAuthenticatorCharacteristics() bool`

HasAuthenticatorCharacteristics returns a boolean if a field has been set.

### GetName

`func (o *CustomAAGUIDUpdateRequestObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomAAGUIDUpdateRequestObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomAAGUIDUpdateRequestObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomAAGUIDUpdateRequestObject) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


