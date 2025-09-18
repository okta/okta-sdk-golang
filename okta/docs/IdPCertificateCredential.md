# IdPCertificateCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X5c** | **[]string** | Base64-encoded X.509 certificate chain with DER encoding | 

## Methods

### NewIdPCertificateCredential

`func NewIdPCertificateCredential(x5c []string, ) *IdPCertificateCredential`

NewIdPCertificateCredential instantiates a new IdPCertificateCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdPCertificateCredentialWithDefaults

`func NewIdPCertificateCredentialWithDefaults() *IdPCertificateCredential`

NewIdPCertificateCredentialWithDefaults instantiates a new IdPCertificateCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX5c

`func (o *IdPCertificateCredential) GetX5c() []string`

GetX5c returns the X5c field if non-nil, zero value otherwise.

### GetX5cOk

`func (o *IdPCertificateCredential) GetX5cOk() (*[]string, bool)`

GetX5cOk returns a tuple with the X5c field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5c

`func (o *IdPCertificateCredential) SetX5c(v []string)`

SetX5c sets X5c field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


