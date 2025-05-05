# IdentityProviderLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 
**Acs** | Pointer to [**HrefObject**](HrefObject.md) |  | [optional] 
**Authorize** | Pointer to [**HrefObject**](HrefObject.md) |  | [optional] 
**ClientRedirectUri** | Pointer to [**HrefObject**](HrefObject.md) |  | [optional] 
**Metadata** | Pointer to [**HrefObject**](HrefObject.md) |  | [optional] 
**Users** | Pointer to [**HrefObject**](HrefObject.md) |  | [optional] 
**Deactivate** | Pointer to [**HrefObject**](HrefObject.md) |  | [optional] 
**Activate** | Pointer to [**HrefObject**](HrefObject.md) |  | [optional] 
**Keys** | Pointer to [**HrefObject**](HrefObject.md) |  | [optional] 

## Methods

### NewIdentityProviderLinks

`func NewIdentityProviderLinks() *IdentityProviderLinks`

NewIdentityProviderLinks instantiates a new IdentityProviderLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderLinksWithDefaults

`func NewIdentityProviderLinksWithDefaults() *IdentityProviderLinks`

NewIdentityProviderLinksWithDefaults instantiates a new IdentityProviderLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *IdentityProviderLinks) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *IdentityProviderLinks) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *IdentityProviderLinks) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *IdentityProviderLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetAcs

`func (o *IdentityProviderLinks) GetAcs() HrefObject`

GetAcs returns the Acs field if non-nil, zero value otherwise.

### GetAcsOk

`func (o *IdentityProviderLinks) GetAcsOk() (*HrefObject, bool)`

GetAcsOk returns a tuple with the Acs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcs

`func (o *IdentityProviderLinks) SetAcs(v HrefObject)`

SetAcs sets Acs field to given value.

### HasAcs

`func (o *IdentityProviderLinks) HasAcs() bool`

HasAcs returns a boolean if a field has been set.

### GetAuthorize

`func (o *IdentityProviderLinks) GetAuthorize() HrefObject`

GetAuthorize returns the Authorize field if non-nil, zero value otherwise.

### GetAuthorizeOk

`func (o *IdentityProviderLinks) GetAuthorizeOk() (*HrefObject, bool)`

GetAuthorizeOk returns a tuple with the Authorize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorize

`func (o *IdentityProviderLinks) SetAuthorize(v HrefObject)`

SetAuthorize sets Authorize field to given value.

### HasAuthorize

`func (o *IdentityProviderLinks) HasAuthorize() bool`

HasAuthorize returns a boolean if a field has been set.

### GetClientRedirectUri

`func (o *IdentityProviderLinks) GetClientRedirectUri() HrefObject`

GetClientRedirectUri returns the ClientRedirectUri field if non-nil, zero value otherwise.

### GetClientRedirectUriOk

`func (o *IdentityProviderLinks) GetClientRedirectUriOk() (*HrefObject, bool)`

GetClientRedirectUriOk returns a tuple with the ClientRedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRedirectUri

`func (o *IdentityProviderLinks) SetClientRedirectUri(v HrefObject)`

SetClientRedirectUri sets ClientRedirectUri field to given value.

### HasClientRedirectUri

`func (o *IdentityProviderLinks) HasClientRedirectUri() bool`

HasClientRedirectUri returns a boolean if a field has been set.

### GetMetadata

`func (o *IdentityProviderLinks) GetMetadata() HrefObject`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IdentityProviderLinks) GetMetadataOk() (*HrefObject, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IdentityProviderLinks) SetMetadata(v HrefObject)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IdentityProviderLinks) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUsers

`func (o *IdentityProviderLinks) GetUsers() HrefObject`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *IdentityProviderLinks) GetUsersOk() (*HrefObject, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *IdentityProviderLinks) SetUsers(v HrefObject)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *IdentityProviderLinks) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetDeactivate

`func (o *IdentityProviderLinks) GetDeactivate() HrefObject`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *IdentityProviderLinks) GetDeactivateOk() (*HrefObject, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *IdentityProviderLinks) SetDeactivate(v HrefObject)`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *IdentityProviderLinks) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### GetActivate

`func (o *IdentityProviderLinks) GetActivate() HrefObject`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *IdentityProviderLinks) GetActivateOk() (*HrefObject, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *IdentityProviderLinks) SetActivate(v HrefObject)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *IdentityProviderLinks) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### GetKeys

`func (o *IdentityProviderLinks) GetKeys() HrefObject`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *IdentityProviderLinks) GetKeysOk() (*HrefObject, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *IdentityProviderLinks) SetKeys(v HrefObject)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *IdentityProviderLinks) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


