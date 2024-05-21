# AuthServerLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 
**Claims** | Pointer to [**AuthServerLinksAllOfClaims**](AuthServerLinksAllOfClaims.md) |  | [optional] 
**Deactivate** | Pointer to [**HrefObjectDeactivateLink**](HrefObjectDeactivateLink.md) |  | [optional] 
**Metadata** | Pointer to [**[]HrefObject**](HrefObject.md) | Link to the authorization server metadata | [optional] 
**Policies** | Pointer to [**AuthServerLinksAllOfPolicies**](AuthServerLinksAllOfPolicies.md) |  | [optional] 
**RotateKey** | Pointer to [**AuthServerLinksAllOfRotateKey**](AuthServerLinksAllOfRotateKey.md) |  | [optional] 
**Scopes** | Pointer to [**AuthServerLinksAllOfScopes**](AuthServerLinksAllOfScopes.md) |  | [optional] 

## Methods

### NewAuthServerLinks

`func NewAuthServerLinks() *AuthServerLinks`

NewAuthServerLinks instantiates a new AuthServerLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthServerLinksWithDefaults

`func NewAuthServerLinksWithDefaults() *AuthServerLinks`

NewAuthServerLinksWithDefaults instantiates a new AuthServerLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *AuthServerLinks) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AuthServerLinks) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AuthServerLinks) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AuthServerLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetClaims

`func (o *AuthServerLinks) GetClaims() AuthServerLinksAllOfClaims`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *AuthServerLinks) GetClaimsOk() (*AuthServerLinksAllOfClaims, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *AuthServerLinks) SetClaims(v AuthServerLinksAllOfClaims)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *AuthServerLinks) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetDeactivate

`func (o *AuthServerLinks) GetDeactivate() HrefObjectDeactivateLink`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *AuthServerLinks) GetDeactivateOk() (*HrefObjectDeactivateLink, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *AuthServerLinks) SetDeactivate(v HrefObjectDeactivateLink)`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *AuthServerLinks) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### GetMetadata

`func (o *AuthServerLinks) GetMetadata() []HrefObject`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AuthServerLinks) GetMetadataOk() (*[]HrefObject, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AuthServerLinks) SetMetadata(v []HrefObject)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AuthServerLinks) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPolicies

`func (o *AuthServerLinks) GetPolicies() AuthServerLinksAllOfPolicies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AuthServerLinks) GetPoliciesOk() (*AuthServerLinksAllOfPolicies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AuthServerLinks) SetPolicies(v AuthServerLinksAllOfPolicies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *AuthServerLinks) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetRotateKey

`func (o *AuthServerLinks) GetRotateKey() AuthServerLinksAllOfRotateKey`

GetRotateKey returns the RotateKey field if non-nil, zero value otherwise.

### GetRotateKeyOk

`func (o *AuthServerLinks) GetRotateKeyOk() (*AuthServerLinksAllOfRotateKey, bool)`

GetRotateKeyOk returns a tuple with the RotateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateKey

`func (o *AuthServerLinks) SetRotateKey(v AuthServerLinksAllOfRotateKey)`

SetRotateKey sets RotateKey field to given value.

### HasRotateKey

`func (o *AuthServerLinks) HasRotateKey() bool`

HasRotateKey returns a boolean if a field has been set.

### GetScopes

`func (o *AuthServerLinks) GetScopes() AuthServerLinksAllOfScopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthServerLinks) GetScopesOk() (*AuthServerLinksAllOfScopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthServerLinks) SetScopes(v AuthServerLinksAllOfScopes)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AuthServerLinks) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


