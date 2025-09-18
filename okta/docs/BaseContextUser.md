# BaseContextUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the user | [optional] 
**PasswordChanged** | Pointer to **time.Time** | The timestamp when the user&#39;s password was last updated | [optional] 
**Profile** | Pointer to [**BaseContextUserProfile**](BaseContextUserProfile.md) |  | [optional] 
**Links** | Pointer to [**BaseContextUserLinks**](BaseContextUserLinks.md) |  | [optional] 

## Methods

### NewBaseContextUser

`func NewBaseContextUser() *BaseContextUser`

NewBaseContextUser instantiates a new BaseContextUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseContextUserWithDefaults

`func NewBaseContextUserWithDefaults() *BaseContextUser`

NewBaseContextUserWithDefaults instantiates a new BaseContextUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseContextUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseContextUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseContextUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseContextUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPasswordChanged

`func (o *BaseContextUser) GetPasswordChanged() time.Time`

GetPasswordChanged returns the PasswordChanged field if non-nil, zero value otherwise.

### GetPasswordChangedOk

`func (o *BaseContextUser) GetPasswordChangedOk() (*time.Time, bool)`

GetPasswordChangedOk returns a tuple with the PasswordChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChanged

`func (o *BaseContextUser) SetPasswordChanged(v time.Time)`

SetPasswordChanged sets PasswordChanged field to given value.

### HasPasswordChanged

`func (o *BaseContextUser) HasPasswordChanged() bool`

HasPasswordChanged returns a boolean if a field has been set.

### GetProfile

`func (o *BaseContextUser) GetProfile() BaseContextUserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *BaseContextUser) GetProfileOk() (*BaseContextUserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *BaseContextUser) SetProfile(v BaseContextUserProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *BaseContextUser) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetLinks

`func (o *BaseContextUser) GetLinks() BaseContextUserLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BaseContextUser) GetLinksOk() (*BaseContextUserLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BaseContextUser) SetLinks(v BaseContextUserLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BaseContextUser) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


