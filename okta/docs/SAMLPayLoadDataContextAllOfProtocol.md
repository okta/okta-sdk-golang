# SAMLPayLoadDataContextAllOfProtocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of authentication protocol being used for the assertion | [optional] 
**Issuer** | Pointer to [**SAMLPayLoadDataContextAllOfProtocolIssuer**](SAMLPayLoadDataContextAllOfProtocolIssuer.md) |  | [optional] 

## Methods

### NewSAMLPayLoadDataContextAllOfProtocol

`func NewSAMLPayLoadDataContextAllOfProtocol() *SAMLPayLoadDataContextAllOfProtocol`

NewSAMLPayLoadDataContextAllOfProtocol instantiates a new SAMLPayLoadDataContextAllOfProtocol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLPayLoadDataContextAllOfProtocolWithDefaults

`func NewSAMLPayLoadDataContextAllOfProtocolWithDefaults() *SAMLPayLoadDataContextAllOfProtocol`

NewSAMLPayLoadDataContextAllOfProtocolWithDefaults instantiates a new SAMLPayLoadDataContextAllOfProtocol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SAMLPayLoadDataContextAllOfProtocol) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SAMLPayLoadDataContextAllOfProtocol) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SAMLPayLoadDataContextAllOfProtocol) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SAMLPayLoadDataContextAllOfProtocol) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIssuer

`func (o *SAMLPayLoadDataContextAllOfProtocol) GetIssuer() SAMLPayLoadDataContextAllOfProtocolIssuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SAMLPayLoadDataContextAllOfProtocol) GetIssuerOk() (*SAMLPayLoadDataContextAllOfProtocolIssuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SAMLPayLoadDataContextAllOfProtocol) SetIssuer(v SAMLPayLoadDataContextAllOfProtocolIssuer)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SAMLPayLoadDataContextAllOfProtocol) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


