# PrivilegedResourceUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to **map[string]interface{}** | Specific profile properties for the privileged resource | [optional] [readonly] 
**UserName** | Pointer to **string** | The username associated with the privileged resource | [optional] 

## Methods

### NewPrivilegedResourceUpdateRequest

`func NewPrivilegedResourceUpdateRequest() *PrivilegedResourceUpdateRequest`

NewPrivilegedResourceUpdateRequest instantiates a new PrivilegedResourceUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegedResourceUpdateRequestWithDefaults

`func NewPrivilegedResourceUpdateRequestWithDefaults() *PrivilegedResourceUpdateRequest`

NewPrivilegedResourceUpdateRequestWithDefaults instantiates a new PrivilegedResourceUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *PrivilegedResourceUpdateRequest) GetProfile() map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PrivilegedResourceUpdateRequest) GetProfileOk() (*map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PrivilegedResourceUpdateRequest) SetProfile(v map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PrivilegedResourceUpdateRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetUserName

`func (o *PrivilegedResourceUpdateRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *PrivilegedResourceUpdateRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *PrivilegedResourceUpdateRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *PrivilegedResourceUpdateRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


