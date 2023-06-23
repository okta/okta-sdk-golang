# LinkedObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associated** | Pointer to [**LinkedObjectDetails**](LinkedObjectDetails.md) |  | [optional] 
**Primary** | Pointer to [**LinkedObjectDetails**](LinkedObjectDetails.md) |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewLinkedObject

`func NewLinkedObject() *LinkedObject`

NewLinkedObject instantiates a new LinkedObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedObjectWithDefaults

`func NewLinkedObjectWithDefaults() *LinkedObject`

NewLinkedObjectWithDefaults instantiates a new LinkedObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociated

`func (o *LinkedObject) GetAssociated() LinkedObjectDetails`

GetAssociated returns the Associated field if non-nil, zero value otherwise.

### GetAssociatedOk

`func (o *LinkedObject) GetAssociatedOk() (*LinkedObjectDetails, bool)`

GetAssociatedOk returns a tuple with the Associated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociated

`func (o *LinkedObject) SetAssociated(v LinkedObjectDetails)`

SetAssociated sets Associated field to given value.

### HasAssociated

`func (o *LinkedObject) HasAssociated() bool`

HasAssociated returns a boolean if a field has been set.

### GetPrimary

`func (o *LinkedObject) GetPrimary() LinkedObjectDetails`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *LinkedObject) GetPrimaryOk() (*LinkedObjectDetails, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *LinkedObject) SetPrimary(v LinkedObjectDetails)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *LinkedObject) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetLinks

`func (o *LinkedObject) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LinkedObject) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LinkedObject) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LinkedObject) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


