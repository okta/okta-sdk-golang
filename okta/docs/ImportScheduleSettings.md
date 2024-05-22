# ImportScheduleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | **string** | The import schedule in UNIX cron format | 
**Timezone** | Pointer to **string** | The import schedule time zone in Internet Assigned Numbers Authority (IANA) time zone name format | [optional] 

## Methods

### NewImportScheduleSettings

`func NewImportScheduleSettings(expression string, ) *ImportScheduleSettings`

NewImportScheduleSettings instantiates a new ImportScheduleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportScheduleSettingsWithDefaults

`func NewImportScheduleSettingsWithDefaults() *ImportScheduleSettings`

NewImportScheduleSettingsWithDefaults instantiates a new ImportScheduleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *ImportScheduleSettings) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ImportScheduleSettings) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ImportScheduleSettings) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetTimezone

`func (o *ImportScheduleSettings) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ImportScheduleSettings) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ImportScheduleSettings) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ImportScheduleSettings) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


