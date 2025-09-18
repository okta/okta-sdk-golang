# UserImportRequestDataAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** | The current default action that results when Okta imports a user. The two possible values are &#x60;CREATE_USER&#x60; and &#x60;LINK_USER&#x60;. You can change the action that is taken by means of the commands object you return. | [optional] 

## Methods

### NewUserImportRequestDataAction

`func NewUserImportRequestDataAction() *UserImportRequestDataAction`

NewUserImportRequestDataAction instantiates a new UserImportRequestDataAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserImportRequestDataActionWithDefaults

`func NewUserImportRequestDataActionWithDefaults() *UserImportRequestDataAction`

NewUserImportRequestDataActionWithDefaults instantiates a new UserImportRequestDataAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *UserImportRequestDataAction) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UserImportRequestDataAction) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UserImportRequestDataAction) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *UserImportRequestDataAction) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


