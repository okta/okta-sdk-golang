# ImportScheduleObjectFullImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | **string** | The import schedule in UNIX cron format | 
**Timezone** | Pointer to **string** | The import schedule time zone in Internet Assigned Numbers Authority (IANA) time zone name format | [optional] 

## Methods

### NewImportScheduleObjectFullImport

`func NewImportScheduleObjectFullImport(expression string, ) *ImportScheduleObjectFullImport`

NewImportScheduleObjectFullImport instantiates a new ImportScheduleObjectFullImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportScheduleObjectFullImportWithDefaults

`func NewImportScheduleObjectFullImportWithDefaults() *ImportScheduleObjectFullImport`

NewImportScheduleObjectFullImportWithDefaults instantiates a new ImportScheduleObjectFullImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *ImportScheduleObjectFullImport) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ImportScheduleObjectFullImport) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ImportScheduleObjectFullImport) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetTimezone

`func (o *ImportScheduleObjectFullImport) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ImportScheduleObjectFullImport) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ImportScheduleObjectFullImport) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ImportScheduleObjectFullImport) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


