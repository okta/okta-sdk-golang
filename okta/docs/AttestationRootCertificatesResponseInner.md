# AttestationRootCertificatesResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X5c** | Pointer to **string** | X.509 certificate chain | [optional] 
**X5tS256** | Pointer to **string** | SHA-256 hash (thumbprint) of the X.509 certificate | [optional] 
**Iss** | Pointer to **string** | Issuer of certificate | [optional] 
**Exp** | Pointer to **string** | Expiry date of certificate | [optional] 

## Methods

### NewAttestationRootCertificatesResponseInner

`func NewAttestationRootCertificatesResponseInner() *AttestationRootCertificatesResponseInner`

NewAttestationRootCertificatesResponseInner instantiates a new AttestationRootCertificatesResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttestationRootCertificatesResponseInnerWithDefaults

`func NewAttestationRootCertificatesResponseInnerWithDefaults() *AttestationRootCertificatesResponseInner`

NewAttestationRootCertificatesResponseInnerWithDefaults instantiates a new AttestationRootCertificatesResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX5c

`func (o *AttestationRootCertificatesResponseInner) GetX5c() string`

GetX5c returns the X5c field if non-nil, zero value otherwise.

### GetX5cOk

`func (o *AttestationRootCertificatesResponseInner) GetX5cOk() (*string, bool)`

GetX5cOk returns a tuple with the X5c field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5c

`func (o *AttestationRootCertificatesResponseInner) SetX5c(v string)`

SetX5c sets X5c field to given value.

### HasX5c

`func (o *AttestationRootCertificatesResponseInner) HasX5c() bool`

HasX5c returns a boolean if a field has been set.

### GetX5tS256

`func (o *AttestationRootCertificatesResponseInner) GetX5tS256() string`

GetX5tS256 returns the X5tS256 field if non-nil, zero value otherwise.

### GetX5tS256Ok

`func (o *AttestationRootCertificatesResponseInner) GetX5tS256Ok() (*string, bool)`

GetX5tS256Ok returns a tuple with the X5tS256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5tS256

`func (o *AttestationRootCertificatesResponseInner) SetX5tS256(v string)`

SetX5tS256 sets X5tS256 field to given value.

### HasX5tS256

`func (o *AttestationRootCertificatesResponseInner) HasX5tS256() bool`

HasX5tS256 returns a boolean if a field has been set.

### GetIss

`func (o *AttestationRootCertificatesResponseInner) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *AttestationRootCertificatesResponseInner) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *AttestationRootCertificatesResponseInner) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *AttestationRootCertificatesResponseInner) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetExp

`func (o *AttestationRootCertificatesResponseInner) GetExp() string`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *AttestationRootCertificatesResponseInner) GetExpOk() (*string, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *AttestationRootCertificatesResponseInner) SetExp(v string)`

SetExp sets Exp field to given value.

### HasExp

`func (o *AttestationRootCertificatesResponseInner) HasExp() bool`

HasExp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


