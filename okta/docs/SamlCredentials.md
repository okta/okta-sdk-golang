# SamlCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signing** | Pointer to [**SamlSigningCredentials**](SamlSigningCredentials.md) |  | [optional] 
**Trust** | Pointer to [**SamlTrustCredentials**](SamlTrustCredentials.md) |  | [optional] 

## Methods

### NewSamlCredentials

`func NewSamlCredentials() *SamlCredentials`

NewSamlCredentials instantiates a new SamlCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlCredentialsWithDefaults

`func NewSamlCredentialsWithDefaults() *SamlCredentials`

NewSamlCredentialsWithDefaults instantiates a new SamlCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigning

`func (o *SamlCredentials) GetSigning() SamlSigningCredentials`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *SamlCredentials) GetSigningOk() (*SamlSigningCredentials, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *SamlCredentials) SetSigning(v SamlSigningCredentials)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *SamlCredentials) HasSigning() bool`

HasSigning returns a boolean if a field has been set.

### GetTrust

`func (o *SamlCredentials) GetTrust() SamlTrustCredentials`

GetTrust returns the Trust field if non-nil, zero value otherwise.

### GetTrustOk

`func (o *SamlCredentials) GetTrustOk() (*SamlTrustCredentials, bool)`

GetTrustOk returns a tuple with the Trust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrust

`func (o *SamlCredentials) SetTrust(v SamlTrustCredentials)`

SetTrust sets Trust field to given value.

### HasTrust

`func (o *SamlCredentials) HasTrust() bool`

HasTrust returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


