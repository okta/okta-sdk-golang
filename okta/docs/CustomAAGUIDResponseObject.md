# CustomAAGUIDResponseObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aaguid** | Pointer to **string** | A unique 128-bit identifier that&#39;s assigned to a specific model of security key or authenticator | [optional] 
**AttestationRootCertificates** | Pointer to [**[]AttestationRootCertificatesResponseInner**](AttestationRootCertificatesResponseInner.md) |  | [optional] 
**AuthenticatorCharacteristics** | Pointer to [**AAGUIDAuthenticatorCharacteristics**](AAGUIDAuthenticatorCharacteristics.md) |  | [optional] 
**Name** | Pointer to **string** | The product name associated with the AAGUID | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewCustomAAGUIDResponseObject

`func NewCustomAAGUIDResponseObject() *CustomAAGUIDResponseObject`

NewCustomAAGUIDResponseObject instantiates a new CustomAAGUIDResponseObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAAGUIDResponseObjectWithDefaults

`func NewCustomAAGUIDResponseObjectWithDefaults() *CustomAAGUIDResponseObject`

NewCustomAAGUIDResponseObjectWithDefaults instantiates a new CustomAAGUIDResponseObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAaguid

`func (o *CustomAAGUIDResponseObject) GetAaguid() string`

GetAaguid returns the Aaguid field if non-nil, zero value otherwise.

### GetAaguidOk

`func (o *CustomAAGUIDResponseObject) GetAaguidOk() (*string, bool)`

GetAaguidOk returns a tuple with the Aaguid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaguid

`func (o *CustomAAGUIDResponseObject) SetAaguid(v string)`

SetAaguid sets Aaguid field to given value.

### HasAaguid

`func (o *CustomAAGUIDResponseObject) HasAaguid() bool`

HasAaguid returns a boolean if a field has been set.

### GetAttestationRootCertificates

`func (o *CustomAAGUIDResponseObject) GetAttestationRootCertificates() []AttestationRootCertificatesResponseInner`

GetAttestationRootCertificates returns the AttestationRootCertificates field if non-nil, zero value otherwise.

### GetAttestationRootCertificatesOk

`func (o *CustomAAGUIDResponseObject) GetAttestationRootCertificatesOk() (*[]AttestationRootCertificatesResponseInner, bool)`

GetAttestationRootCertificatesOk returns a tuple with the AttestationRootCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationRootCertificates

`func (o *CustomAAGUIDResponseObject) SetAttestationRootCertificates(v []AttestationRootCertificatesResponseInner)`

SetAttestationRootCertificates sets AttestationRootCertificates field to given value.

### HasAttestationRootCertificates

`func (o *CustomAAGUIDResponseObject) HasAttestationRootCertificates() bool`

HasAttestationRootCertificates returns a boolean if a field has been set.

### GetAuthenticatorCharacteristics

`func (o *CustomAAGUIDResponseObject) GetAuthenticatorCharacteristics() AAGUIDAuthenticatorCharacteristics`

GetAuthenticatorCharacteristics returns the AuthenticatorCharacteristics field if non-nil, zero value otherwise.

### GetAuthenticatorCharacteristicsOk

`func (o *CustomAAGUIDResponseObject) GetAuthenticatorCharacteristicsOk() (*AAGUIDAuthenticatorCharacteristics, bool)`

GetAuthenticatorCharacteristicsOk returns a tuple with the AuthenticatorCharacteristics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorCharacteristics

`func (o *CustomAAGUIDResponseObject) SetAuthenticatorCharacteristics(v AAGUIDAuthenticatorCharacteristics)`

SetAuthenticatorCharacteristics sets AuthenticatorCharacteristics field to given value.

### HasAuthenticatorCharacteristics

`func (o *CustomAAGUIDResponseObject) HasAuthenticatorCharacteristics() bool`

HasAuthenticatorCharacteristics returns a boolean if a field has been set.

### GetName

`func (o *CustomAAGUIDResponseObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomAAGUIDResponseObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomAAGUIDResponseObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomAAGUIDResponseObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLinks

`func (o *CustomAAGUIDResponseObject) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomAAGUIDResponseObject) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomAAGUIDResponseObject) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CustomAAGUIDResponseObject) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


