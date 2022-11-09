# DomainCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** |  | [optional] 
**CertificateChain** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DomainCertificateType**](DomainCertificateType.md) |  | [optional] 

## Methods

### NewDomainCertificate

`func NewDomainCertificate() *DomainCertificate`

NewDomainCertificate instantiates a new DomainCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainCertificateWithDefaults

`func NewDomainCertificateWithDefaults() *DomainCertificate`

NewDomainCertificateWithDefaults instantiates a new DomainCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *DomainCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *DomainCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *DomainCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *DomainCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCertificateChain

`func (o *DomainCertificate) GetCertificateChain() string`

GetCertificateChain returns the CertificateChain field if non-nil, zero value otherwise.

### GetCertificateChainOk

`func (o *DomainCertificate) GetCertificateChainOk() (*string, bool)`

GetCertificateChainOk returns a tuple with the CertificateChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChain

`func (o *DomainCertificate) SetCertificateChain(v string)`

SetCertificateChain sets CertificateChain field to given value.

### HasCertificateChain

`func (o *DomainCertificate) HasCertificateChain() bool`

HasCertificateChain returns a boolean if a field has been set.

### GetPrivateKey

`func (o *DomainCertificate) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *DomainCertificate) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *DomainCertificate) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *DomainCertificate) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetType

`func (o *DomainCertificate) GetType() DomainCertificateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainCertificate) GetTypeOk() (*DomainCertificateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainCertificate) SetType(v DomainCertificateType)`

SetType sets Type field to given value.

### HasType

`func (o *DomainCertificate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


