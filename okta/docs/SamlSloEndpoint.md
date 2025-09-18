# SamlSloEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binding** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | URL of the binding-specific IdP endpoint where Okta sends a &#x60;&lt;LogoutRequest&gt;&#x60; | [optional] 

## Methods

### NewSamlSloEndpoint

`func NewSamlSloEndpoint() *SamlSloEndpoint`

NewSamlSloEndpoint instantiates a new SamlSloEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlSloEndpointWithDefaults

`func NewSamlSloEndpointWithDefaults() *SamlSloEndpoint`

NewSamlSloEndpointWithDefaults instantiates a new SamlSloEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinding

`func (o *SamlSloEndpoint) GetBinding() string`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *SamlSloEndpoint) GetBindingOk() (*string, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *SamlSloEndpoint) SetBinding(v string)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *SamlSloEndpoint) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetUrl

`func (o *SamlSloEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SamlSloEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SamlSloEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SamlSloEndpoint) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


