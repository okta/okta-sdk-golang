# CertificateAuthorityToScopeMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaId** | **string** | The &#x60;id&#x60; of the Trusted Certificate Authority that&#39;s mapped to the scope. | 
**Id** | **string** | The unique identifier of the Certificate Authority mapping. | [readonly] 
**Scope** | **string** | The scope (feature) that a Certificate Authority mapping is associated with. | 

## Methods

### NewCertificateAuthorityToScopeMapping

`func NewCertificateAuthorityToScopeMapping(caId string, id string, scope string, ) *CertificateAuthorityToScopeMapping`

NewCertificateAuthorityToScopeMapping instantiates a new CertificateAuthorityToScopeMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthorityToScopeMappingWithDefaults

`func NewCertificateAuthorityToScopeMappingWithDefaults() *CertificateAuthorityToScopeMapping`

NewCertificateAuthorityToScopeMappingWithDefaults instantiates a new CertificateAuthorityToScopeMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaId

`func (o *CertificateAuthorityToScopeMapping) GetCaId() string`

GetCaId returns the CaId field if non-nil, zero value otherwise.

### GetCaIdOk

`func (o *CertificateAuthorityToScopeMapping) GetCaIdOk() (*string, bool)`

GetCaIdOk returns a tuple with the CaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaId

`func (o *CertificateAuthorityToScopeMapping) SetCaId(v string)`

SetCaId sets CaId field to given value.


### GetId

`func (o *CertificateAuthorityToScopeMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateAuthorityToScopeMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateAuthorityToScopeMapping) SetId(v string)`

SetId sets Id field to given value.


### GetScope

`func (o *CertificateAuthorityToScopeMapping) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CertificateAuthorityToScopeMapping) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CertificateAuthorityToScopeMapping) SetScope(v string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


