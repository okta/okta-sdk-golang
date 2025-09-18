# AuthorizationServerPolicyAllOfLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 
**Activate** | Pointer to [**HrefObjectActivateLink**](HrefObjectActivateLink.md) |  | [optional] 
**Deactivate** | Pointer to [**HrefObjectDeactivateLink**](HrefObjectDeactivateLink.md) |  | [optional] 
**Rules** | Pointer to [**HrefObject**](HrefObject.md) | Link to the authorization server policy&#39;s rules | [optional] 

## Methods

### NewAuthorizationServerPolicyAllOfLinks

`func NewAuthorizationServerPolicyAllOfLinks() *AuthorizationServerPolicyAllOfLinks`

NewAuthorizationServerPolicyAllOfLinks instantiates a new AuthorizationServerPolicyAllOfLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerPolicyAllOfLinksWithDefaults

`func NewAuthorizationServerPolicyAllOfLinksWithDefaults() *AuthorizationServerPolicyAllOfLinks`

NewAuthorizationServerPolicyAllOfLinksWithDefaults instantiates a new AuthorizationServerPolicyAllOfLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *AuthorizationServerPolicyAllOfLinks) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AuthorizationServerPolicyAllOfLinks) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AuthorizationServerPolicyAllOfLinks) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AuthorizationServerPolicyAllOfLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetActivate

`func (o *AuthorizationServerPolicyAllOfLinks) GetActivate() HrefObjectActivateLink`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *AuthorizationServerPolicyAllOfLinks) GetActivateOk() (*HrefObjectActivateLink, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *AuthorizationServerPolicyAllOfLinks) SetActivate(v HrefObjectActivateLink)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *AuthorizationServerPolicyAllOfLinks) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### GetDeactivate

`func (o *AuthorizationServerPolicyAllOfLinks) GetDeactivate() HrefObjectDeactivateLink`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *AuthorizationServerPolicyAllOfLinks) GetDeactivateOk() (*HrefObjectDeactivateLink, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *AuthorizationServerPolicyAllOfLinks) SetDeactivate(v HrefObjectDeactivateLink)`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *AuthorizationServerPolicyAllOfLinks) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### GetRules

`func (o *AuthorizationServerPolicyAllOfLinks) GetRules() HrefObject`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AuthorizationServerPolicyAllOfLinks) GetRulesOk() (*HrefObject, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AuthorizationServerPolicyAllOfLinks) SetRules(v HrefObject)`

SetRules sets Rules field to given value.

### HasRules

`func (o *AuthorizationServerPolicyAllOfLinks) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


