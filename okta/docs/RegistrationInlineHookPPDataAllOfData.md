# RegistrationInlineHookPPDataAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**RegistrationInlineHookPPDataAllOfDataContext**](RegistrationInlineHookPPDataAllOfDataContext.md) |  | [optional] 
**Action** | Pointer to **string** | The default action the system takes. Set to &#x60;ALLOW&#x60;. &#x60;DENY&#x60; is never sent to your external service | [optional] 
**UserProfileUpdate** | Pointer to **map[string]interface{}** | Name-value pairs for each new attribute supplied by the user in the Progressive Profile form | [optional] 

## Methods

### NewRegistrationInlineHookPPDataAllOfData

`func NewRegistrationInlineHookPPDataAllOfData() *RegistrationInlineHookPPDataAllOfData`

NewRegistrationInlineHookPPDataAllOfData instantiates a new RegistrationInlineHookPPDataAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationInlineHookPPDataAllOfDataWithDefaults

`func NewRegistrationInlineHookPPDataAllOfDataWithDefaults() *RegistrationInlineHookPPDataAllOfData`

NewRegistrationInlineHookPPDataAllOfDataWithDefaults instantiates a new RegistrationInlineHookPPDataAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *RegistrationInlineHookPPDataAllOfData) GetContext() RegistrationInlineHookPPDataAllOfDataContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RegistrationInlineHookPPDataAllOfData) GetContextOk() (*RegistrationInlineHookPPDataAllOfDataContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RegistrationInlineHookPPDataAllOfData) SetContext(v RegistrationInlineHookPPDataAllOfDataContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *RegistrationInlineHookPPDataAllOfData) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetAction

`func (o *RegistrationInlineHookPPDataAllOfData) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RegistrationInlineHookPPDataAllOfData) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RegistrationInlineHookPPDataAllOfData) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RegistrationInlineHookPPDataAllOfData) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetUserProfileUpdate

`func (o *RegistrationInlineHookPPDataAllOfData) GetUserProfileUpdate() map[string]interface{}`

GetUserProfileUpdate returns the UserProfileUpdate field if non-nil, zero value otherwise.

### GetUserProfileUpdateOk

`func (o *RegistrationInlineHookPPDataAllOfData) GetUserProfileUpdateOk() (*map[string]interface{}, bool)`

GetUserProfileUpdateOk returns a tuple with the UserProfileUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfileUpdate

`func (o *RegistrationInlineHookPPDataAllOfData) SetUserProfileUpdate(v map[string]interface{})`

SetUserProfileUpdate sets UserProfileUpdate field to given value.

### HasUserProfileUpdate

`func (o *RegistrationInlineHookPPDataAllOfData) HasUserProfileUpdate() bool`

HasUserProfileUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


