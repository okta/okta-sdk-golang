# CertificateAuthorityList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CertificateAuthority**](CertificateAuthority.md) | The certificate authorities, roots before intermediates | [optional] 

## Methods

### NewCertificateAuthorityList

`func NewCertificateAuthorityList() *CertificateAuthorityList`

NewCertificateAuthorityList instantiates a new CertificateAuthorityList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthorityListWithDefaults

`func NewCertificateAuthorityListWithDefaults() *CertificateAuthorityList`

NewCertificateAuthorityListWithDefaults instantiates a new CertificateAuthorityList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CertificateAuthorityList) GetData() []CertificateAuthority`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CertificateAuthorityList) GetDataOk() (*[]CertificateAuthority, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CertificateAuthorityList) SetData(v []CertificateAuthority)`

SetData sets Data field to given value.

### HasData

`func (o *CertificateAuthorityList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


