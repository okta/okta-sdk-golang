# PotentialConnectionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PotentialConnection**](PotentialConnection.md) | Potential connections that can be established | 
**Links** | [**PotentialConnectionListLinks**](PotentialConnectionListLinks.md) |  | 

## Methods

### NewPotentialConnectionList

`func NewPotentialConnectionList(data []PotentialConnection, links PotentialConnectionListLinks, ) *PotentialConnectionList`

NewPotentialConnectionList instantiates a new PotentialConnectionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPotentialConnectionListWithDefaults

`func NewPotentialConnectionListWithDefaults() *PotentialConnectionList`

NewPotentialConnectionListWithDefaults instantiates a new PotentialConnectionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PotentialConnectionList) GetData() []PotentialConnection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PotentialConnectionList) GetDataOk() (*[]PotentialConnection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PotentialConnectionList) SetData(v []PotentialConnection)`

SetData sets Data field to given value.


### GetLinks

`func (o *PotentialConnectionList) GetLinks() PotentialConnectionListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PotentialConnectionList) GetLinksOk() (*PotentialConnectionListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PotentialConnectionList) SetLinks(v PotentialConnectionListLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


