# ApplicationEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **map[string]map[string]interface{}** | The specified [Application User](/openapi/okta-management/management/tag/ApplicationUsers/) assigned to the app | [optional] 

## Methods

### NewApplicationEmbedded

`func NewApplicationEmbedded() *ApplicationEmbedded`

NewApplicationEmbedded instantiates a new ApplicationEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationEmbeddedWithDefaults

`func NewApplicationEmbeddedWithDefaults() *ApplicationEmbedded`

NewApplicationEmbeddedWithDefaults instantiates a new ApplicationEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *ApplicationEmbedded) GetUser() map[string]map[string]interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ApplicationEmbedded) GetUserOk() (*map[string]map[string]interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ApplicationEmbedded) SetUser(v map[string]map[string]interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *ApplicationEmbedded) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


