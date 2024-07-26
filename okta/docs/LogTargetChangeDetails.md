# LogTargetChangeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **map[string]map[string]interface{}** | The original properties of the target | [optional] 
**To** | Pointer to **map[string]map[string]interface{}** | The updated properties of the target | [optional] 

## Methods

### NewLogTargetChangeDetails

`func NewLogTargetChangeDetails() *LogTargetChangeDetails`

NewLogTargetChangeDetails instantiates a new LogTargetChangeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogTargetChangeDetailsWithDefaults

`func NewLogTargetChangeDetailsWithDefaults() *LogTargetChangeDetails`

NewLogTargetChangeDetailsWithDefaults instantiates a new LogTargetChangeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *LogTargetChangeDetails) GetFrom() map[string]map[string]interface{}`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *LogTargetChangeDetails) GetFromOk() (*map[string]map[string]interface{}, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *LogTargetChangeDetails) SetFrom(v map[string]map[string]interface{})`

SetFrom sets From field to given value.

### HasFrom

`func (o *LogTargetChangeDetails) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *LogTargetChangeDetails) GetTo() map[string]map[string]interface{}`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LogTargetChangeDetails) GetToOk() (*map[string]map[string]interface{}, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LogTargetChangeDetails) SetTo(v map[string]map[string]interface{})`

SetTo sets To field to given value.

### HasTo

`func (o *LogTargetChangeDetails) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


