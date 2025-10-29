# PolicyAccountLinkFilterUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclude** | Pointer to **[]string** | Specifies the blocklist of user identifiers to exclude from account linking | [optional] 
**ExcludeAdmins** | Pointer to **bool** | Specifies whether admin users should be excluded from account linking | [optional] [default to false]

## Methods

### NewPolicyAccountLinkFilterUsers

`func NewPolicyAccountLinkFilterUsers() *PolicyAccountLinkFilterUsers`

NewPolicyAccountLinkFilterUsers instantiates a new PolicyAccountLinkFilterUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAccountLinkFilterUsersWithDefaults

`func NewPolicyAccountLinkFilterUsersWithDefaults() *PolicyAccountLinkFilterUsers`

NewPolicyAccountLinkFilterUsersWithDefaults instantiates a new PolicyAccountLinkFilterUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclude

`func (o *PolicyAccountLinkFilterUsers) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *PolicyAccountLinkFilterUsers) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *PolicyAccountLinkFilterUsers) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *PolicyAccountLinkFilterUsers) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetExcludeAdmins

`func (o *PolicyAccountLinkFilterUsers) GetExcludeAdmins() bool`

GetExcludeAdmins returns the ExcludeAdmins field if non-nil, zero value otherwise.

### GetExcludeAdminsOk

`func (o *PolicyAccountLinkFilterUsers) GetExcludeAdminsOk() (*bool, bool)`

GetExcludeAdminsOk returns a tuple with the ExcludeAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAdmins

`func (o *PolicyAccountLinkFilterUsers) SetExcludeAdmins(v bool)`

SetExcludeAdmins sets ExcludeAdmins field to given value.

### HasExcludeAdmins

`func (o *PolicyAccountLinkFilterUsers) HasExcludeAdmins() bool`

HasExcludeAdmins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


