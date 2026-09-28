# Office365ClientCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Include** | Pointer to **[]string** | List of Office 365 client types to match | [optional] 

## Methods

### NewOffice365ClientCondition

`func NewOffice365ClientCondition() *Office365ClientCondition`

NewOffice365ClientCondition instantiates a new Office365ClientCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365ClientConditionWithDefaults

`func NewOffice365ClientConditionWithDefaults() *Office365ClientCondition`

NewOffice365ClientConditionWithDefaults instantiates a new Office365ClientCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInclude

`func (o *Office365ClientCondition) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *Office365ClientCondition) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *Office365ClientCondition) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *Office365ClientCondition) HasInclude() bool`

HasInclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


