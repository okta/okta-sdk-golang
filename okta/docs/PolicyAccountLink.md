# PolicyAccountLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**PolicyAccountLinkAction**](PolicyAccountLinkAction.md) |  | [optional] 
**Filter** | Pointer to [**PolicyAccountLinkFilter**](PolicyAccountLinkFilter.md) |  | [optional] 

## Methods

### NewPolicyAccountLink

`func NewPolicyAccountLink() *PolicyAccountLink`

NewPolicyAccountLink instantiates a new PolicyAccountLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAccountLinkWithDefaults

`func NewPolicyAccountLinkWithDefaults() *PolicyAccountLink`

NewPolicyAccountLinkWithDefaults instantiates a new PolicyAccountLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PolicyAccountLink) GetAction() PolicyAccountLinkAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicyAccountLink) GetActionOk() (*PolicyAccountLinkAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicyAccountLink) SetAction(v PolicyAccountLinkAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *PolicyAccountLink) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetFilter

`func (o *PolicyAccountLink) GetFilter() PolicyAccountLinkFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PolicyAccountLink) GetFilterOk() (*PolicyAccountLinkFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PolicyAccountLink) SetFilter(v PolicyAccountLinkFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PolicyAccountLink) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


