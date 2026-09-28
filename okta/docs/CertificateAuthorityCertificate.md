# CertificateAuthorityCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aki** | Pointer to **string** | The authority key identifier of this certificate — the &#x60;ski&#x60; of the CA certificate that signed it. This is how the trust chain is walked. A self-signed root has no parent. | [optional] [readonly] 
**AuthorityInstanceId** | Pointer to **string** | The stable identifier of the authority this certificate belongs to | [optional] [readonly] 
**CaType** | Pointer to **string** | The family of the certificate authority. &#x60;OKTA_AS_CA&#x60; is the Okta CA used for SCEP enrollment and device management attestation; &#x60;DEVICE_ACCESS_OKTA_AS_CA&#x60; is the Device Access CA. The &#x60;caType&#x60; query parameter filters on the same values. | [optional] 
**CreatedDate** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**ExpirationDate** | Pointer to **time.Time** | Timestamp when this certificate expires | [optional] [readonly] 
**Jwk** | Pointer to [**CertificateAuthorityJsonWebKey**](CertificateAuthorityJsonWebKey.md) |  | [optional] 
**LastUpdatedDate** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**RotationState** | Pointer to **string** | The rotation stage of the CA certificate. &#x60;ACTIVE&#x60; indicates the authority&#39;s newest certificate, which is the one that a renewal migrates SCEP configurations onto. &#x60;RETIRING&#x60; indicates a certificate that has been superseded by a renewal but remains valid, allowing enrolled devices to continue functioning during migration. | [optional] [readonly] 
**Ski** | Pointer to **string** | The subject key identifier of this certificate, in uppercase hexadecimal. It identifies the certificate within its authority. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the certificate. &#x60;VALID&#x60; is in service and able to issue certificates, which is how Okta generates a replacement CA certificate. &#x60;INACTIVE&#x60; is a replacement that Okta generated but hasn&#39;t brought into service, which occurs only for an authority renewed before selective renewal was available; renewing such an authority activates it. | [optional] [readonly] 
**SubjectDN** | Pointer to **string** | The distinguished name in the subject of this certificate | [optional] [readonly] 
**Type** | Pointer to **string** | The position of the certificate authority in the hierarchy. A root is self-signed; an intermediate is signed by the root and signs the certificates Okta issues to devices. The &#x60;type&#x60; query parameter filters on the same values. | [optional] 

## Methods

### NewCertificateAuthorityCertificate

`func NewCertificateAuthorityCertificate() *CertificateAuthorityCertificate`

NewCertificateAuthorityCertificate instantiates a new CertificateAuthorityCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthorityCertificateWithDefaults

`func NewCertificateAuthorityCertificateWithDefaults() *CertificateAuthorityCertificate`

NewCertificateAuthorityCertificateWithDefaults instantiates a new CertificateAuthorityCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAki

`func (o *CertificateAuthorityCertificate) GetAki() string`

GetAki returns the Aki field if non-nil, zero value otherwise.

### GetAkiOk

`func (o *CertificateAuthorityCertificate) GetAkiOk() (*string, bool)`

GetAkiOk returns a tuple with the Aki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAki

`func (o *CertificateAuthorityCertificate) SetAki(v string)`

SetAki sets Aki field to given value.

### HasAki

`func (o *CertificateAuthorityCertificate) HasAki() bool`

HasAki returns a boolean if a field has been set.

### GetAuthorityInstanceId

`func (o *CertificateAuthorityCertificate) GetAuthorityInstanceId() string`

GetAuthorityInstanceId returns the AuthorityInstanceId field if non-nil, zero value otherwise.

### GetAuthorityInstanceIdOk

`func (o *CertificateAuthorityCertificate) GetAuthorityInstanceIdOk() (*string, bool)`

GetAuthorityInstanceIdOk returns a tuple with the AuthorityInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityInstanceId

`func (o *CertificateAuthorityCertificate) SetAuthorityInstanceId(v string)`

SetAuthorityInstanceId sets AuthorityInstanceId field to given value.

### HasAuthorityInstanceId

`func (o *CertificateAuthorityCertificate) HasAuthorityInstanceId() bool`

HasAuthorityInstanceId returns a boolean if a field has been set.

### GetCaType

`func (o *CertificateAuthorityCertificate) GetCaType() string`

GetCaType returns the CaType field if non-nil, zero value otherwise.

### GetCaTypeOk

`func (o *CertificateAuthorityCertificate) GetCaTypeOk() (*string, bool)`

GetCaTypeOk returns a tuple with the CaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaType

`func (o *CertificateAuthorityCertificate) SetCaType(v string)`

SetCaType sets CaType field to given value.

### HasCaType

`func (o *CertificateAuthorityCertificate) HasCaType() bool`

HasCaType returns a boolean if a field has been set.

### GetCreatedDate

`func (o *CertificateAuthorityCertificate) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CertificateAuthorityCertificate) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CertificateAuthorityCertificate) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *CertificateAuthorityCertificate) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CertificateAuthorityCertificate) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CertificateAuthorityCertificate) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CertificateAuthorityCertificate) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CertificateAuthorityCertificate) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetJwk

`func (o *CertificateAuthorityCertificate) GetJwk() CertificateAuthorityJsonWebKey`

GetJwk returns the Jwk field if non-nil, zero value otherwise.

### GetJwkOk

`func (o *CertificateAuthorityCertificate) GetJwkOk() (*CertificateAuthorityJsonWebKey, bool)`

GetJwkOk returns a tuple with the Jwk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwk

`func (o *CertificateAuthorityCertificate) SetJwk(v CertificateAuthorityJsonWebKey)`

SetJwk sets Jwk field to given value.

### HasJwk

`func (o *CertificateAuthorityCertificate) HasJwk() bool`

HasJwk returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *CertificateAuthorityCertificate) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *CertificateAuthorityCertificate) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *CertificateAuthorityCertificate) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *CertificateAuthorityCertificate) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetRotationState

`func (o *CertificateAuthorityCertificate) GetRotationState() string`

GetRotationState returns the RotationState field if non-nil, zero value otherwise.

### GetRotationStateOk

`func (o *CertificateAuthorityCertificate) GetRotationStateOk() (*string, bool)`

GetRotationStateOk returns a tuple with the RotationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationState

`func (o *CertificateAuthorityCertificate) SetRotationState(v string)`

SetRotationState sets RotationState field to given value.

### HasRotationState

`func (o *CertificateAuthorityCertificate) HasRotationState() bool`

HasRotationState returns a boolean if a field has been set.

### GetSki

`func (o *CertificateAuthorityCertificate) GetSki() string`

GetSki returns the Ski field if non-nil, zero value otherwise.

### GetSkiOk

`func (o *CertificateAuthorityCertificate) GetSkiOk() (*string, bool)`

GetSkiOk returns a tuple with the Ski field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSki

`func (o *CertificateAuthorityCertificate) SetSki(v string)`

SetSki sets Ski field to given value.

### HasSki

`func (o *CertificateAuthorityCertificate) HasSki() bool`

HasSki returns a boolean if a field has been set.

### GetStatus

`func (o *CertificateAuthorityCertificate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateAuthorityCertificate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateAuthorityCertificate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CertificateAuthorityCertificate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubjectDN

`func (o *CertificateAuthorityCertificate) GetSubjectDN() string`

GetSubjectDN returns the SubjectDN field if non-nil, zero value otherwise.

### GetSubjectDNOk

`func (o *CertificateAuthorityCertificate) GetSubjectDNOk() (*string, bool)`

GetSubjectDNOk returns a tuple with the SubjectDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDN

`func (o *CertificateAuthorityCertificate) SetSubjectDN(v string)`

SetSubjectDN sets SubjectDN field to given value.

### HasSubjectDN

`func (o *CertificateAuthorityCertificate) HasSubjectDN() bool`

HasSubjectDN returns a boolean if a field has been set.

### GetType

`func (o *CertificateAuthorityCertificate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateAuthorityCertificate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateAuthorityCertificate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CertificateAuthorityCertificate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


