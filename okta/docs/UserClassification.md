# UserClassification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdated** | Pointer to **time.Time** | The timestamp when the user classification was last updated | [optional] [readonly] 
**Type** | Pointer to **string** | The type of user classification | [optional] 

## Methods

### NewUserClassification

`func NewUserClassification() *UserClassification`

NewUserClassification instantiates a new UserClassification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserClassificationWithDefaults

`func NewUserClassificationWithDefaults() *UserClassification`

NewUserClassificationWithDefaults instantiates a new UserClassification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdated

`func (o *UserClassification) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UserClassification) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UserClassification) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UserClassification) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetType

`func (o *UserClassification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserClassification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserClassification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserClassification) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


