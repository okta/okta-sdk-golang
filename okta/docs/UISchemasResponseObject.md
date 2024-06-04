# UISchemasResponseObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp when the UI Schema was created (ISO-86001) | [readonly] 
**Id** | **string** | Unique identifier for the UI Schema | [readonly] 
**LastUpdated** | **time.Time** | Timestamp when the UI Schema was last modified (ISO-86001) | [readonly] 
**UiSchema** | [**UISchemaObject**](UISchemaObject.md) |  | 
**Links** | [**LinksSelf**](LinksSelf.md) |  | 

## Methods

### NewUISchemasResponseObject

`func NewUISchemasResponseObject(created time.Time, id string, lastUpdated time.Time, uiSchema UISchemaObject, links LinksSelf, ) *UISchemasResponseObject`

NewUISchemasResponseObject instantiates a new UISchemasResponseObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUISchemasResponseObjectWithDefaults

`func NewUISchemasResponseObjectWithDefaults() *UISchemasResponseObject`

NewUISchemasResponseObjectWithDefaults instantiates a new UISchemasResponseObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *UISchemasResponseObject) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UISchemasResponseObject) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UISchemasResponseObject) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *UISchemasResponseObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UISchemasResponseObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UISchemasResponseObject) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdated

`func (o *UISchemasResponseObject) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UISchemasResponseObject) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UISchemasResponseObject) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetUiSchema

`func (o *UISchemasResponseObject) GetUiSchema() UISchemaObject`

GetUiSchema returns the UiSchema field if non-nil, zero value otherwise.

### GetUiSchemaOk

`func (o *UISchemasResponseObject) GetUiSchemaOk() (*UISchemaObject, bool)`

GetUiSchemaOk returns a tuple with the UiSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiSchema

`func (o *UISchemasResponseObject) SetUiSchema(v UISchemaObject)`

SetUiSchema sets UiSchema field to given value.


### GetLinks

`func (o *UISchemasResponseObject) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UISchemasResponseObject) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UISchemasResponseObject) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


