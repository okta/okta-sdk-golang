# SecurityEventTokenRequestJwtHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | **string** | Algorithm used to sign or encrypt the JWT | 
**Kid** | **string** | Key ID used to sign or encrypt the JWT | 
**Typ** | **string** | The type of content being signed or encrypted | 

## Methods

### NewSecurityEventTokenRequestJwtHeader

`func NewSecurityEventTokenRequestJwtHeader(alg string, kid string, typ string, ) *SecurityEventTokenRequestJwtHeader`

NewSecurityEventTokenRequestJwtHeader instantiates a new SecurityEventTokenRequestJwtHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityEventTokenRequestJwtHeaderWithDefaults

`func NewSecurityEventTokenRequestJwtHeaderWithDefaults() *SecurityEventTokenRequestJwtHeader`

NewSecurityEventTokenRequestJwtHeaderWithDefaults instantiates a new SecurityEventTokenRequestJwtHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *SecurityEventTokenRequestJwtHeader) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *SecurityEventTokenRequestJwtHeader) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *SecurityEventTokenRequestJwtHeader) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetKid

`func (o *SecurityEventTokenRequestJwtHeader) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *SecurityEventTokenRequestJwtHeader) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *SecurityEventTokenRequestJwtHeader) SetKid(v string)`

SetKid sets Kid field to given value.


### GetTyp

`func (o *SecurityEventTokenRequestJwtHeader) GetTyp() string`

GetTyp returns the Typ field if non-nil, zero value otherwise.

### GetTypOk

`func (o *SecurityEventTokenRequestJwtHeader) GetTypOk() (*string, bool)`

GetTypOk returns a tuple with the Typ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTyp

`func (o *SecurityEventTokenRequestJwtHeader) SetTyp(v string)`

SetTyp sets Typ field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


