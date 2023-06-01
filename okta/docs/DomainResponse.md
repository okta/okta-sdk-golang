# DomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandId** | Pointer to **string** |  | [optional] 
**CertificateSourceType** | Pointer to [**DomainCertificateSourceType**](DomainCertificateSourceType.md) |  | [optional] 
**DnsRecords** | Pointer to [**[]DNSRecord**](DNSRecord.md) |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PublicCertificate** | Pointer to [**DomainCertificateMetadata**](DomainCertificateMetadata.md) |  | [optional] 
**ValidationStatus** | Pointer to [**DomainValidationStatus**](DomainValidationStatus.md) |  | [optional] 
**Links** | Pointer to [**DomainLinks**](DomainLinks.md) |  | [optional] 

## Methods

### NewDomainResponse

`func NewDomainResponse() *DomainResponse`

NewDomainResponse instantiates a new DomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainResponseWithDefaults

`func NewDomainResponseWithDefaults() *DomainResponse`

NewDomainResponseWithDefaults instantiates a new DomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandId

`func (o *DomainResponse) GetBrandId() string`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *DomainResponse) GetBrandIdOk() (*string, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *DomainResponse) SetBrandId(v string)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *DomainResponse) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### GetCertificateSourceType

`func (o *DomainResponse) GetCertificateSourceType() DomainCertificateSourceType`

GetCertificateSourceType returns the CertificateSourceType field if non-nil, zero value otherwise.

### GetCertificateSourceTypeOk

`func (o *DomainResponse) GetCertificateSourceTypeOk() (*DomainCertificateSourceType, bool)`

GetCertificateSourceTypeOk returns a tuple with the CertificateSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateSourceType

`func (o *DomainResponse) SetCertificateSourceType(v DomainCertificateSourceType)`

SetCertificateSourceType sets CertificateSourceType field to given value.

### HasCertificateSourceType

`func (o *DomainResponse) HasCertificateSourceType() bool`

HasCertificateSourceType returns a boolean if a field has been set.

### GetDnsRecords

`func (o *DomainResponse) GetDnsRecords() []DNSRecord`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *DomainResponse) GetDnsRecordsOk() (*[]DNSRecord, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *DomainResponse) SetDnsRecords(v []DNSRecord)`

SetDnsRecords sets DnsRecords field to given value.

### HasDnsRecords

`func (o *DomainResponse) HasDnsRecords() bool`

HasDnsRecords returns a boolean if a field has been set.

### GetDomain

`func (o *DomainResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DomainResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetId

`func (o *DomainResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPublicCertificate

`func (o *DomainResponse) GetPublicCertificate() DomainCertificateMetadata`

GetPublicCertificate returns the PublicCertificate field if non-nil, zero value otherwise.

### GetPublicCertificateOk

`func (o *DomainResponse) GetPublicCertificateOk() (*DomainCertificateMetadata, bool)`

GetPublicCertificateOk returns a tuple with the PublicCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicCertificate

`func (o *DomainResponse) SetPublicCertificate(v DomainCertificateMetadata)`

SetPublicCertificate sets PublicCertificate field to given value.

### HasPublicCertificate

`func (o *DomainResponse) HasPublicCertificate() bool`

HasPublicCertificate returns a boolean if a field has been set.

### GetValidationStatus

`func (o *DomainResponse) GetValidationStatus() DomainValidationStatus`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *DomainResponse) GetValidationStatusOk() (*DomainValidationStatus, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *DomainResponse) SetValidationStatus(v DomainValidationStatus)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *DomainResponse) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.

### GetLinks

`func (o *DomainResponse) GetLinks() DomainLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DomainResponse) GetLinksOk() (*DomainLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DomainResponse) SetLinks(v DomainLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DomainResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


