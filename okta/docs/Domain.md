# Domain

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

## Methods

### NewDomain

`func NewDomain() *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandId

`func (o *Domain) GetBrandId() string`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *Domain) GetBrandIdOk() (*string, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *Domain) SetBrandId(v string)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *Domain) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### GetCertificateSourceType

`func (o *Domain) GetCertificateSourceType() DomainCertificateSourceType`

GetCertificateSourceType returns the CertificateSourceType field if non-nil, zero value otherwise.

### GetCertificateSourceTypeOk

`func (o *Domain) GetCertificateSourceTypeOk() (*DomainCertificateSourceType, bool)`

GetCertificateSourceTypeOk returns a tuple with the CertificateSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateSourceType

`func (o *Domain) SetCertificateSourceType(v DomainCertificateSourceType)`

SetCertificateSourceType sets CertificateSourceType field to given value.

### HasCertificateSourceType

`func (o *Domain) HasCertificateSourceType() bool`

HasCertificateSourceType returns a boolean if a field has been set.

### GetDnsRecords

`func (o *Domain) GetDnsRecords() []DNSRecord`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *Domain) GetDnsRecordsOk() (*[]DNSRecord, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *Domain) SetDnsRecords(v []DNSRecord)`

SetDnsRecords sets DnsRecords field to given value.

### HasDnsRecords

`func (o *Domain) HasDnsRecords() bool`

HasDnsRecords returns a boolean if a field has been set.

### GetDomain

`func (o *Domain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Domain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Domain) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Domain) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetId

`func (o *Domain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Domain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPublicCertificate

`func (o *Domain) GetPublicCertificate() DomainCertificateMetadata`

GetPublicCertificate returns the PublicCertificate field if non-nil, zero value otherwise.

### GetPublicCertificateOk

`func (o *Domain) GetPublicCertificateOk() (*DomainCertificateMetadata, bool)`

GetPublicCertificateOk returns a tuple with the PublicCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicCertificate

`func (o *Domain) SetPublicCertificate(v DomainCertificateMetadata)`

SetPublicCertificate sets PublicCertificate field to given value.

### HasPublicCertificate

`func (o *Domain) HasPublicCertificate() bool`

HasPublicCertificate returns a boolean if a field has been set.

### GetValidationStatus

`func (o *Domain) GetValidationStatus() DomainValidationStatus`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *Domain) GetValidationStatusOk() (*DomainValidationStatus, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *Domain) SetValidationStatus(v DomainValidationStatus)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *Domain) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


