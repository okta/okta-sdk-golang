# UserIdentificationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityMethods** | Pointer to [**UserIdentificationSecurityMethods**](UserIdentificationSecurityMethods.md) |  | [optional] 

## Methods

### NewUserIdentificationSettings

`func NewUserIdentificationSettings() *UserIdentificationSettings`

NewUserIdentificationSettings instantiates a new UserIdentificationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentificationSettingsWithDefaults

`func NewUserIdentificationSettingsWithDefaults() *UserIdentificationSettings`

NewUserIdentificationSettingsWithDefaults instantiates a new UserIdentificationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityMethods

`func (o *UserIdentificationSettings) GetSecurityMethods() UserIdentificationSecurityMethods`

GetSecurityMethods returns the SecurityMethods field if non-nil, zero value otherwise.

### GetSecurityMethodsOk

`func (o *UserIdentificationSettings) GetSecurityMethodsOk() (*UserIdentificationSecurityMethods, bool)`

GetSecurityMethodsOk returns a tuple with the SecurityMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMethods

`func (o *UserIdentificationSettings) SetSecurityMethods(v UserIdentificationSecurityMethods)`

SetSecurityMethods sets SecurityMethods field to given value.

### HasSecurityMethods

`func (o *UserIdentificationSettings) HasSecurityMethods() bool`

HasSecurityMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


