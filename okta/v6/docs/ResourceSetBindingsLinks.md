# ResourceSetBindingsLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 
**Next** | Pointer to [**HrefObject**](HrefObject.md) | Link to the next list of bindings for the specified resource set | [optional] 
**ResourceSet** | Pointer to [**HrefObjectResourceSetLink**](HrefObjectResourceSetLink.md) |  | [optional] 

## Methods

### NewResourceSetBindingsLinks

`func NewResourceSetBindingsLinks() *ResourceSetBindingsLinks`

NewResourceSetBindingsLinks instantiates a new ResourceSetBindingsLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSetBindingsLinksWithDefaults

`func NewResourceSetBindingsLinksWithDefaults() *ResourceSetBindingsLinks`

NewResourceSetBindingsLinksWithDefaults instantiates a new ResourceSetBindingsLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *ResourceSetBindingsLinks) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ResourceSetBindingsLinks) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ResourceSetBindingsLinks) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ResourceSetBindingsLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetNext

`func (o *ResourceSetBindingsLinks) GetNext() HrefObject`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResourceSetBindingsLinks) GetNextOk() (*HrefObject, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResourceSetBindingsLinks) SetNext(v HrefObject)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResourceSetBindingsLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResourceSet

`func (o *ResourceSetBindingsLinks) GetResourceSet() HrefObjectResourceSetLink`

GetResourceSet returns the ResourceSet field if non-nil, zero value otherwise.

### GetResourceSetOk

`func (o *ResourceSetBindingsLinks) GetResourceSetOk() (*HrefObjectResourceSetLink, bool)`

GetResourceSetOk returns a tuple with the ResourceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSet

`func (o *ResourceSetBindingsLinks) SetResourceSet(v HrefObjectResourceSetLink)`

SetResourceSet sets ResourceSet field to given value.

### HasResourceSet

`func (o *ResourceSetBindingsLinks) HasResourceSet() bool`

HasResourceSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


