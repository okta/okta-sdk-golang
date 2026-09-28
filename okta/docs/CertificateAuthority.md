# CertificateAuthority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorityInstanceId** | Pointer to **string** | The stable identifier of the authority. For the authority&#39;s original certificate, this is also that certificate&#39;s &#x60;ski&#x60;, because a CA&#39;s own subject key identifier becomes the authority&#39;s identifier when the CA is created. | [optional] [readonly] 
**CaType** | Pointer to **string** | The family of the certificate authority. &#x60;OKTA_AS_CA&#x60; is the Okta CA used for SCEP enrollment and device management attestation; &#x60;DEVICE_ACCESS_OKTA_AS_CA&#x60; is the Device Access CA. The &#x60;caType&#x60; query parameter filters on the same values. | [optional] 
**Status** | Pointer to **string** | The status of the certificate. &#x60;VALID&#x60; is in service and able to issue certificates, which is how Okta generates a replacement CA certificate. &#x60;INACTIVE&#x60; is a replacement that Okta generated but hasn&#39;t brought into service, which occurs only for an authority renewed before selective renewal was available; renewing such an authority activates it. | [optional] [readonly] 
**SubjectDN** | Pointer to **string** | The distinguished name in the subject of the authority&#39;s certificate | [optional] [readonly] 
**Type** | Pointer to **string** | The position of the certificate authority in the hierarchy. A root is self-signed; an intermediate is signed by the root and signs the certificates Okta issues to devices. The &#x60;type&#x60; query parameter filters on the same values. | [optional] 
**Embedded** | Pointer to [**CertificateAuthorityEmbedded**](CertificateAuthorityEmbedded.md) |  | [optional] 

## Methods

### NewCertificateAuthority

`func NewCertificateAuthority() *CertificateAuthority`

NewCertificateAuthority instantiates a new CertificateAuthority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthorityWithDefaults

`func NewCertificateAuthorityWithDefaults() *CertificateAuthority`

NewCertificateAuthorityWithDefaults instantiates a new CertificateAuthority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorityInstanceId

`func (o *CertificateAuthority) GetAuthorityInstanceId() string`

GetAuthorityInstanceId returns the AuthorityInstanceId field if non-nil, zero value otherwise.

### GetAuthorityInstanceIdOk

`func (o *CertificateAuthority) GetAuthorityInstanceIdOk() (*string, bool)`

GetAuthorityInstanceIdOk returns a tuple with the AuthorityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityInstanceId

`func (o *CertificateAuthority) SetAuthorityInstanceId(v string)`

SetAuthorityInstanceId sets AuthorityInstanceId field to given value.

### HasAuthorityInstanceId

`func (o *CertificateAuthority) HasAuthorityInstanceId() bool`

HasAuthorityInstanceId returns a boolean if a field has been set.

### GetCaType

`func (o *CertificateAuthority) GetCaType() string`

GetCaType returns the CaType field if non-nil, zero value otherwise.

### GetCaTypeOk

`func (o *CertificateAuthority) GetCaTypeOk() (*string, bool)`

GetCaTypeOk returns a tuple with the CaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaType

`func (o *CertificateAuthority) SetCaType(v string)`

SetCaType sets CaType field to given value.

### HasCaType

`func (o *CertificateAuthority) HasCaType() bool`

HasCaType returns a boolean if a field has been set.

### GetStatus

`func (o *CertificateAuthority) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateAuthority) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateAuthority) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CertificateAuthority) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubjectDN

`func (o *CertificateAuthority) GetSubjectDN() string`

GetSubjectDN returns the SubjectDN field if non-nil, zero value otherwise.

### GetSubjectDNOk

`func (o *CertificateAuthority) GetSubjectDNOk() (*string, bool)`

GetSubjectDNOk returns a tuple with the SubjectDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDN

`func (o *CertificateAuthority) SetSubjectDN(v string)`

SetSubjectDN sets SubjectDN field to given value.

### HasSubjectDN

`func (o *CertificateAuthority) HasSubjectDN() bool`

HasSubjectDN returns a boolean if a field has been set.

### GetType

`func (o *CertificateAuthority) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateAuthority) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateAuthority) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CertificateAuthority) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmbedded

`func (o *CertificateAuthority) GetEmbedded() CertificateAuthorityEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *CertificateAuthority) GetEmbeddedOk() (*CertificateAuthorityEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *CertificateAuthority) SetEmbedded(v CertificateAuthorityEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *CertificateAuthority) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


