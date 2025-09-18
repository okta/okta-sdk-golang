# RegistrationInlineHookSSRDataAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**RegistrationInlineHookSSRDataAllOfDataContext**](RegistrationInlineHookSSRDataAllOfDataContext.md) |  | [optional] 
**Action** | Pointer to **string** | The default action the system will take. Will be &#x60;ALLOW&#x60;. &#x60;DENY&#x60; will never be sent to your external service. | [optional] 
**UserProfile** | Pointer to **map[string]interface{}** | The name-value pairs for each registration-related attribute supplied by the user in the Profile Enrollment form. | [optional] 

## Methods

### NewRegistrationInlineHookSSRDataAllOfData

`func NewRegistrationInlineHookSSRDataAllOfData() *RegistrationInlineHookSSRDataAllOfData`

NewRegistrationInlineHookSSRDataAllOfData instantiates a new RegistrationInlineHookSSRDataAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationInlineHookSSRDataAllOfDataWithDefaults

`func NewRegistrationInlineHookSSRDataAllOfDataWithDefaults() *RegistrationInlineHookSSRDataAllOfData`

NewRegistrationInlineHookSSRDataAllOfDataWithDefaults instantiates a new RegistrationInlineHookSSRDataAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *RegistrationInlineHookSSRDataAllOfData) GetContext() RegistrationInlineHookSSRDataAllOfDataContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RegistrationInlineHookSSRDataAllOfData) GetContextOk() (*RegistrationInlineHookSSRDataAllOfDataContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RegistrationInlineHookSSRDataAllOfData) SetContext(v RegistrationInlineHookSSRDataAllOfDataContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *RegistrationInlineHookSSRDataAllOfData) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetAction

`func (o *RegistrationInlineHookSSRDataAllOfData) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RegistrationInlineHookSSRDataAllOfData) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RegistrationInlineHookSSRDataAllOfData) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RegistrationInlineHookSSRDataAllOfData) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetUserProfile

`func (o *RegistrationInlineHookSSRDataAllOfData) GetUserProfile() map[string]interface{}`

GetUserProfile returns the UserProfile field if non-nil, zero value otherwise.

### GetUserProfileOk

`func (o *RegistrationInlineHookSSRDataAllOfData) GetUserProfileOk() (*map[string]interface{}, bool)`

GetUserProfileOk returns a tuple with the UserProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfile

`func (o *RegistrationInlineHookSSRDataAllOfData) SetUserProfile(v map[string]interface{})`

SetUserProfile sets UserProfile field to given value.

### HasUserProfile

`func (o *RegistrationInlineHookSSRDataAllOfData) HasUserProfile() bool`

HasUserProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


