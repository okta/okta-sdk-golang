# DomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateSourceType** | **string** | Certificate source type that indicates whether the certificate is provided by the user or Okta. | 
**Domain** | **string** | Custom domain name | 

## Methods

### NewDomainRequest

`func NewDomainRequest(certificateSourceType string, domain string, ) *DomainRequest`

NewDomainRequest instantiates a new DomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRequestWithDefaults

`func NewDomainRequestWithDefaults() *DomainRequest`

NewDomainRequestWithDefaults instantiates a new DomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateSourceType

`func (o *DomainRequest) GetCertificateSourceType() string`

GetCertificateSourceType returns the CertificateSourceType field if non-nil, zero value otherwise.

### GetCertificateSourceTypeOk

`func (o *DomainRequest) GetCertificateSourceTypeOk() (*string, bool)`

GetCertificateSourceTypeOk returns a tuple with the CertificateSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateSourceType

`func (o *DomainRequest) SetCertificateSourceType(v string)`

SetCertificateSourceType sets CertificateSourceType field to given value.


### GetDomain

`func (o *DomainRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


