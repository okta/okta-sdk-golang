# ManagedConnectionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ManagedConnection**](ManagedConnection.md) | All connections the agent has established | 
**Links** | [**ManagedConnectionListLinks**](ManagedConnectionListLinks.md) |  | 

## Methods

### NewManagedConnectionList

`func NewManagedConnectionList(data []ManagedConnection, links ManagedConnectionListLinks, ) *ManagedConnectionList`

NewManagedConnectionList instantiates a new ManagedConnectionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedConnectionListWithDefaults

`func NewManagedConnectionListWithDefaults() *ManagedConnectionList`

NewManagedConnectionListWithDefaults instantiates a new ManagedConnectionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ManagedConnectionList) GetData() []ManagedConnection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ManagedConnectionList) GetDataOk() (*[]ManagedConnection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ManagedConnectionList) SetData(v []ManagedConnection)`

SetData sets Data field to given value.


### GetLinks

`func (o *ManagedConnectionList) GetLinks() ManagedConnectionListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ManagedConnectionList) GetLinksOk() (*ManagedConnectionListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ManagedConnectionList) SetLinks(v ManagedConnectionListLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


