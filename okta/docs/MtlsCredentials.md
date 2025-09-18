# MtlsCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trust** | Pointer to [**MtlsTrustCredentials**](MtlsTrustCredentials.md) |  | [optional] 

## Methods

### NewMtlsCredentials

`func NewMtlsCredentials() *MtlsCredentials`

NewMtlsCredentials instantiates a new MtlsCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMtlsCredentialsWithDefaults

`func NewMtlsCredentialsWithDefaults() *MtlsCredentials`

NewMtlsCredentialsWithDefaults instantiates a new MtlsCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrust

`func (o *MtlsCredentials) GetTrust() MtlsTrustCredentials`

GetTrust returns the Trust field if non-nil, zero value otherwise.

### GetTrustOk

`func (o *MtlsCredentials) GetTrustOk() (*MtlsTrustCredentials, bool)`

GetTrustOk returns a tuple with the Trust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrust

`func (o *MtlsCredentials) SetTrust(v MtlsTrustCredentials)`

SetTrust sets Trust field to given value.

### HasTrust

`func (o *MtlsCredentials) HasTrust() bool`

HasTrust returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


