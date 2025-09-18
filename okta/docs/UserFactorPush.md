# UserFactorPush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**UserFactorPushProfile**](UserFactorPushProfile.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 

## Methods

### NewUserFactorPush

`func NewUserFactorPush() *UserFactorPush`

NewUserFactorPush instantiates a new UserFactorPush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorPushWithDefaults

`func NewUserFactorPushWithDefaults() *UserFactorPush`

NewUserFactorPushWithDefaults instantiates a new UserFactorPush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *UserFactorPush) GetProfile() UserFactorPushProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserFactorPush) GetProfileOk() (*UserFactorPushProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserFactorPush) SetProfile(v UserFactorPushProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserFactorPush) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProvider

`func (o *UserFactorPush) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserFactorPush) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserFactorPush) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserFactorPush) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


