# SelectiveRenewalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewParallelConfigs** | Pointer to [**[]RegistrationAuthority**](RegistrationAuthority.md) | The configurations created against the renewed certificate, each naming the configuration it was cloned from in its own &#x60;sourceConfigId&#x60;. Empty if the request asked for none. | [optional] 
**RenewedCertificate** | Pointer to [**CertificateAuthorityCertificate**](CertificateAuthorityCertificate.md) |  | [optional] 
**RolledOverConfigs** | Pointer to [**[]RegistrationAuthority**](RegistrationAuthority.md) | The configurations now bound to the renewed certificate. Empty if the request rolled none over. | [optional] 

## Methods

### NewSelectiveRenewalResponse

`func NewSelectiveRenewalResponse() *SelectiveRenewalResponse`

NewSelectiveRenewalResponse instantiates a new SelectiveRenewalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectiveRenewalResponseWithDefaults

`func NewSelectiveRenewalResponseWithDefaults() *SelectiveRenewalResponse`

NewSelectiveRenewalResponseWithDefaults instantiates a new SelectiveRenewalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewParallelConfigs

`func (o *SelectiveRenewalResponse) GetNewParallelConfigs() []RegistrationAuthority`

GetNewParallelConfigs returns the NewParallelConfigs field if non-nil, zero value otherwise.

### GetNewParallelConfigsOk

`func (o *SelectiveRenewalResponse) GetNewParallelConfigsOk() (*[]RegistrationAuthority, bool)`

GetNewParallelConfigsOk returns a tuple with the NewParallelConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewParallelConfigs

`func (o *SelectiveRenewalResponse) SetNewParallelConfigs(v []RegistrationAuthority)`

SetNewParallelConfigs sets NewParallelConfigs field to given value.

### HasNewParallelConfigs

`func (o *SelectiveRenewalResponse) HasNewParallelConfigs() bool`

HasNewParallelConfigs returns a boolean if a field has been set.

### GetRenewedCertificate

`func (o *SelectiveRenewalResponse) GetRenewedCertificate() CertificateAuthorityCertificate`

GetRenewedCertificate returns the RenewedCertificate field if non-nil, zero value otherwise.

### GetRenewedCertificateOk

`func (o *SelectiveRenewalResponse) GetRenewedCertificateOk() (*CertificateAuthorityCertificate, bool)`

GetRenewedCertificateOk returns a tuple with the RenewedCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewedCertificate

`func (o *SelectiveRenewalResponse) SetRenewedCertificate(v CertificateAuthorityCertificate)`

SetRenewedCertificate sets RenewedCertificate field to given value.

### HasRenewedCertificate

`func (o *SelectiveRenewalResponse) HasRenewedCertificate() bool`

HasRenewedCertificate returns a boolean if a field has been set.

### GetRolledOverConfigs

`func (o *SelectiveRenewalResponse) GetRolledOverConfigs() []RegistrationAuthority`

GetRolledOverConfigs returns the RolledOverConfigs field if non-nil, zero value otherwise.

### GetRolledOverConfigsOk

`func (o *SelectiveRenewalResponse) GetRolledOverConfigsOk() (*[]RegistrationAuthority, bool)`

GetRolledOverConfigsOk returns a tuple with the RolledOverConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolledOverConfigs

`func (o *SelectiveRenewalResponse) SetRolledOverConfigs(v []RegistrationAuthority)`

SetRolledOverConfigs sets RolledOverConfigs field to given value.

### HasRolledOverConfigs

`func (o *SelectiveRenewalResponse) HasRolledOverConfigs() bool`

HasRolledOverConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


