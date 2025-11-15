# ManagedConnectionListLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | 
**Next** | Pointer to [**HrefObjectNextLink**](HrefObjectNextLink.md) |  | [optional] 

## Methods

### NewManagedConnectionListLinks

`func NewManagedConnectionListLinks(self HrefObjectSelfLink, ) *ManagedConnectionListLinks`

NewManagedConnectionListLinks instantiates a new ManagedConnectionListLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedConnectionListLinksWithDefaults

`func NewManagedConnectionListLinksWithDefaults() *ManagedConnectionListLinks`

NewManagedConnectionListLinksWithDefaults instantiates a new ManagedConnectionListLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *ManagedConnectionListLinks) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ManagedConnectionListLinks) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ManagedConnectionListLinks) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.


### GetNext

`func (o *ManagedConnectionListLinks) GetNext() HrefObjectNextLink`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ManagedConnectionListLinks) GetNextOk() (*HrefObjectNextLink, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ManagedConnectionListLinks) SetNext(v HrefObjectNextLink)`

SetNext sets Next field to given value.

### HasNext

`func (o *ManagedConnectionListLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


