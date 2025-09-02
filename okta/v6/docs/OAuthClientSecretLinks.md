# OAuthClientSecretLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activate** | Pointer to [**HrefObjectActivateLink**](HrefObjectActivateLink.md) |  | [optional] 
**Deactivate** | Pointer to [**HrefObjectDeactivateLink**](HrefObjectDeactivateLink.md) |  | [optional] 
**Delete** | Pointer to [**HrefObjectDeleteLink**](HrefObjectDeleteLink.md) |  | [optional] 

## Methods

### NewOAuthClientSecretLinks

`func NewOAuthClientSecretLinks() *OAuthClientSecretLinks`

NewOAuthClientSecretLinks instantiates a new OAuthClientSecretLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthClientSecretLinksWithDefaults

`func NewOAuthClientSecretLinksWithDefaults() *OAuthClientSecretLinks`

NewOAuthClientSecretLinksWithDefaults instantiates a new OAuthClientSecretLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivate

`func (o *OAuthClientSecretLinks) GetActivate() HrefObjectActivateLink`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *OAuthClientSecretLinks) GetActivateOk() (*HrefObjectActivateLink, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *OAuthClientSecretLinks) SetActivate(v HrefObjectActivateLink)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *OAuthClientSecretLinks) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### GetDeactivate

`func (o *OAuthClientSecretLinks) GetDeactivate() HrefObjectDeactivateLink`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *OAuthClientSecretLinks) GetDeactivateOk() (*HrefObjectDeactivateLink, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *OAuthClientSecretLinks) SetDeactivate(v HrefObjectDeactivateLink)`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *OAuthClientSecretLinks) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### GetDelete

`func (o *OAuthClientSecretLinks) GetDelete() HrefObjectDeleteLink`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *OAuthClientSecretLinks) GetDeleteOk() (*HrefObjectDeleteLink, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *OAuthClientSecretLinks) SetDelete(v HrefObjectDeleteLink)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *OAuthClientSecretLinks) HasDelete() bool`

HasDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


