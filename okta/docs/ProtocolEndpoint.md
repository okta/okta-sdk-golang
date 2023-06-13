# ProtocolEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binding** | Pointer to [**ProtocolEndpointBinding**](ProtocolEndpointBinding.md) |  | [optional] 
**Destination** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ProtocolEndpointType**](ProtocolEndpointType.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewProtocolEndpoint

`func NewProtocolEndpoint() *ProtocolEndpoint`

NewProtocolEndpoint instantiates a new ProtocolEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolEndpointWithDefaults

`func NewProtocolEndpointWithDefaults() *ProtocolEndpoint`

NewProtocolEndpointWithDefaults instantiates a new ProtocolEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinding

`func (o *ProtocolEndpoint) GetBinding() ProtocolEndpointBinding`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *ProtocolEndpoint) GetBindingOk() (*ProtocolEndpointBinding, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *ProtocolEndpoint) SetBinding(v ProtocolEndpointBinding)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *ProtocolEndpoint) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetDestination

`func (o *ProtocolEndpoint) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ProtocolEndpoint) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ProtocolEndpoint) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ProtocolEndpoint) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetType

`func (o *ProtocolEndpoint) GetType() ProtocolEndpointType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProtocolEndpoint) GetTypeOk() (*ProtocolEndpointType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProtocolEndpoint) SetType(v ProtocolEndpointType)`

SetType sets Type field to given value.

### HasType

`func (o *ProtocolEndpoint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *ProtocolEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProtocolEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProtocolEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProtocolEndpoint) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


