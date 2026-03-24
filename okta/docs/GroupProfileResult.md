# GroupProfileResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the group | [optional] 
**Profile** | Pointer to **map[string]interface{}** | Map of requested attributes and their values | [optional] 

## Methods

### NewGroupProfileResult

`func NewGroupProfileResult() *GroupProfileResult`

NewGroupProfileResult instantiates a new GroupProfileResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupProfileResultWithDefaults

`func NewGroupProfileResultWithDefaults() *GroupProfileResult`

NewGroupProfileResultWithDefaults instantiates a new GroupProfileResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupProfileResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupProfileResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupProfileResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupProfileResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfile

`func (o *GroupProfileResult) GetProfile() map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GroupProfileResult) GetProfileOk() (*map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GroupProfileResult) SetProfile(v map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *GroupProfileResult) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


