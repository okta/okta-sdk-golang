# CatalogApplicationLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logo** | Pointer to [**[]HrefObjectLogoLink**](HrefObjectLogoLink.md) | List of app logo resources | [optional] 
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 

## Methods

### NewCatalogApplicationLinks

`func NewCatalogApplicationLinks() *CatalogApplicationLinks`

NewCatalogApplicationLinks instantiates a new CatalogApplicationLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogApplicationLinksWithDefaults

`func NewCatalogApplicationLinksWithDefaults() *CatalogApplicationLinks`

NewCatalogApplicationLinksWithDefaults instantiates a new CatalogApplicationLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogo

`func (o *CatalogApplicationLinks) GetLogo() []HrefObjectLogoLink`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CatalogApplicationLinks) GetLogoOk() (*[]HrefObjectLogoLink, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CatalogApplicationLinks) SetLogo(v []HrefObjectLogoLink)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CatalogApplicationLinks) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetSelf

`func (o *CatalogApplicationLinks) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *CatalogApplicationLinks) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *CatalogApplicationLinks) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *CatalogApplicationLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


