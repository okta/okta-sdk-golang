# MacOSAccountProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountUuid** | Pointer to **string** | Unique identifier for the macOS account | [optional] 
**FullName** | Pointer to **string** | Full name of the account user | [optional] 
**Type** | Pointer to **string** | Profile type discriminator | [optional] 
**Username** | Pointer to **string** | Username of the account | [optional] 

## Methods

### NewMacOSAccountProfile

`func NewMacOSAccountProfile() *MacOSAccountProfile`

NewMacOSAccountProfile instantiates a new MacOSAccountProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacOSAccountProfileWithDefaults

`func NewMacOSAccountProfileWithDefaults() *MacOSAccountProfile`

NewMacOSAccountProfileWithDefaults instantiates a new MacOSAccountProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountUuid

`func (o *MacOSAccountProfile) GetAccountUuid() string`

GetAccountUuid returns the AccountUuid field if non-nil, zero value otherwise.

### GetAccountUuidOk

`func (o *MacOSAccountProfile) GetAccountUuidOk() (*string, bool)`

GetAccountUuidOk returns a tuple with the AccountUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUuid

`func (o *MacOSAccountProfile) SetAccountUuid(v string)`

SetAccountUuid sets AccountUuid field to given value.

### HasAccountUuid

`func (o *MacOSAccountProfile) HasAccountUuid() bool`

HasAccountUuid returns a boolean if a field has been set.

### GetFullName

`func (o *MacOSAccountProfile) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *MacOSAccountProfile) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *MacOSAccountProfile) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *MacOSAccountProfile) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetType

`func (o *MacOSAccountProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MacOSAccountProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MacOSAccountProfile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MacOSAccountProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *MacOSAccountProfile) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MacOSAccountProfile) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MacOSAccountProfile) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MacOSAccountProfile) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


