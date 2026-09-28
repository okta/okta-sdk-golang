# CertificateAuthorityToScopeMappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaType** | Pointer to **string** | The type of the Certificate Authority. | [optional] 
**CreatedDate** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**ExpirationDate** | Pointer to **time.Time** | Timestamp when the object expires | [optional] [readonly] 
**Id** | **string** | The &#x60;id&#x60; of the Certificate Authority. | [readonly] 
**IssuerDN** | Pointer to **string** | The issuer distinguished name parsed from the Certificate Authority&#39;s trust anchor. | [optional] [readonly] 
**LastUpdatedDate** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**Mapping** | Pointer to [**CertificateAuthorityToScopeMappingResponseMapping**](CertificateAuthorityToScopeMappingResponseMapping.md) |  | [optional] 
**Status** | Pointer to **string** | The validity status of the Certificate Authority&#39;s trust anchor. | [optional] 
**SubjectDN** | Pointer to **string** | The subject distinguished name parsed from the Certificate Authority&#39;s trust anchor. | [optional] [readonly] 

## Methods

### NewCertificateAuthorityToScopeMappingResponse

`func NewCertificateAuthorityToScopeMappingResponse(id string, ) *CertificateAuthorityToScopeMappingResponse`

NewCertificateAuthorityToScopeMappingResponse instantiates a new CertificateAuthorityToScopeMappingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthorityToScopeMappingResponseWithDefaults

`func NewCertificateAuthorityToScopeMappingResponseWithDefaults() *CertificateAuthorityToScopeMappingResponse`

NewCertificateAuthorityToScopeMappingResponseWithDefaults instantiates a new CertificateAuthorityToScopeMappingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaType

`func (o *CertificateAuthorityToScopeMappingResponse) GetCaType() string`

GetCaType returns the CaType field if non-nil, zero value otherwise.

### GetCaTypeOk

`func (o *CertificateAuthorityToScopeMappingResponse) GetCaTypeOk() (*string, bool)`

GetCaTypeOk returns a tuple with the CaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaType

`func (o *CertificateAuthorityToScopeMappingResponse) SetCaType(v string)`

SetCaType sets CaType field to given value.

### HasCaType

`func (o *CertificateAuthorityToScopeMappingResponse) HasCaType() bool`

HasCaType returns a boolean if a field has been set.

### GetCreatedDate

`func (o *CertificateAuthorityToScopeMappingResponse) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CertificateAuthorityToScopeMappingResponse) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CertificateAuthorityToScopeMappingResponse) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *CertificateAuthorityToScopeMappingResponse) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CertificateAuthorityToScopeMappingResponse) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CertificateAuthorityToScopeMappingResponse) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CertificateAuthorityToScopeMappingResponse) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CertificateAuthorityToScopeMappingResponse) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetId

`func (o *CertificateAuthorityToScopeMappingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateAuthorityToScopeMappingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateAuthorityToScopeMappingResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIssuerDN

`func (o *CertificateAuthorityToScopeMappingResponse) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *CertificateAuthorityToScopeMappingResponse) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *CertificateAuthorityToScopeMappingResponse) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *CertificateAuthorityToScopeMappingResponse) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *CertificateAuthorityToScopeMappingResponse) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *CertificateAuthorityToScopeMappingResponse) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *CertificateAuthorityToScopeMappingResponse) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *CertificateAuthorityToScopeMappingResponse) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetMapping

`func (o *CertificateAuthorityToScopeMappingResponse) GetMapping() CertificateAuthorityToScopeMappingResponseMapping`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *CertificateAuthorityToScopeMappingResponse) GetMappingOk() (*CertificateAuthorityToScopeMappingResponseMapping, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *CertificateAuthorityToScopeMappingResponse) SetMapping(v CertificateAuthorityToScopeMappingResponseMapping)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *CertificateAuthorityToScopeMappingResponse) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetStatus

`func (o *CertificateAuthorityToScopeMappingResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateAuthorityToScopeMappingResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateAuthorityToScopeMappingResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CertificateAuthorityToScopeMappingResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubjectDN

`func (o *CertificateAuthorityToScopeMappingResponse) GetSubjectDN() string`

GetSubjectDN returns the SubjectDN field if non-nil, zero value otherwise.

### GetSubjectDNOk

`func (o *CertificateAuthorityToScopeMappingResponse) GetSubjectDNOk() (*string, bool)`

GetSubjectDNOk returns a tuple with the SubjectDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDN

`func (o *CertificateAuthorityToScopeMappingResponse) SetSubjectDN(v string)`

SetSubjectDN sets SubjectDN field to given value.

### HasSubjectDN

`func (o *CertificateAuthorityToScopeMappingResponse) HasSubjectDN() bool`

HasSubjectDN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


