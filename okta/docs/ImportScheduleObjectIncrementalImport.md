# ImportScheduleObjectIncrementalImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | **string** | The import schedule in UNIX cron format | 
**Timezone** | Pointer to **string** | The import schedule time zone in Internet Assigned Numbers Authority (IANA) time zone name format | [optional] 

## Methods

### NewImportScheduleObjectIncrementalImport

`func NewImportScheduleObjectIncrementalImport(expression string, ) *ImportScheduleObjectIncrementalImport`

NewImportScheduleObjectIncrementalImport instantiates a new ImportScheduleObjectIncrementalImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportScheduleObjectIncrementalImportWithDefaults

`func NewImportScheduleObjectIncrementalImportWithDefaults() *ImportScheduleObjectIncrementalImport`

NewImportScheduleObjectIncrementalImportWithDefaults instantiates a new ImportScheduleObjectIncrementalImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *ImportScheduleObjectIncrementalImport) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ImportScheduleObjectIncrementalImport) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ImportScheduleObjectIncrementalImport) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetTimezone

`func (o *ImportScheduleObjectIncrementalImport) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ImportScheduleObjectIncrementalImport) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ImportScheduleObjectIncrementalImport) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ImportScheduleObjectIncrementalImport) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


