# SamlAcsEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binding** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Determines whether to publish an instance-specific (trust) or organization (shared) ACS endpoint in the SAML metadata | [optional] [default to "INSTANCE"]

## Methods

### NewSamlAcsEndpoint

`func NewSamlAcsEndpoint() *SamlAcsEndpoint`

NewSamlAcsEndpoint instantiates a new SamlAcsEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlAcsEndpointWithDefaults

`func NewSamlAcsEndpointWithDefaults() *SamlAcsEndpoint`

NewSamlAcsEndpointWithDefaults instantiates a new SamlAcsEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinding

`func (o *SamlAcsEndpoint) GetBinding() string`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *SamlAcsEndpoint) GetBindingOk() (*string, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *SamlAcsEndpoint) SetBinding(v string)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *SamlAcsEndpoint) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetType

`func (o *SamlAcsEndpoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SamlAcsEndpoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SamlAcsEndpoint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SamlAcsEndpoint) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


