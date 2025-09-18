# UserImportResponseCommandsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The command types supported for the import inline hook. When using the &#x60;com.okta.action.update&#x60; command to specify that the user should be treated as a match, you need to also provide a &#x60;com.okta.user.update&#x60; command that sets the ID of the Okta user. | [optional] 
**Value** | Pointer to **map[string]string** | The &#x60;value&#x60; object is the parameter to pass to the command. In the case of the &#x60;com.okta.appUser.profile.update&#x60; and &#x60;com.okta.user.profile.update&#x60; commands, the parameter should be a list of one or more profile attributes and the values you wish to set them to. In the case of the &#x60;com.okta.action.update&#x60; command, the parameter should be a &#x60;result&#x60; property set to either &#x60;CREATE_USER&#x60; or &#x60;LINK_USER&#x60;. | [optional] 

## Methods

### NewUserImportResponseCommandsInner

`func NewUserImportResponseCommandsInner() *UserImportResponseCommandsInner`

NewUserImportResponseCommandsInner instantiates a new UserImportResponseCommandsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserImportResponseCommandsInnerWithDefaults

`func NewUserImportResponseCommandsInnerWithDefaults() *UserImportResponseCommandsInner`

NewUserImportResponseCommandsInnerWithDefaults instantiates a new UserImportResponseCommandsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserImportResponseCommandsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserImportResponseCommandsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserImportResponseCommandsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserImportResponseCommandsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *UserImportResponseCommandsInner) GetValue() map[string]string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserImportResponseCommandsInner) GetValueOk() (*map[string]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserImportResponseCommandsInner) SetValue(v map[string]string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UserImportResponseCommandsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


