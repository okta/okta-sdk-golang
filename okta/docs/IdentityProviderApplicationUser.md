# IdentityProviderApplicationUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**ExternalId** | Pointer to **string** | Unique IdP-specific identifier for the user | [optional] [readonly] 
**Id** | Pointer to **string** | Unique key of the user | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**Profile** | Pointer to **map[string]map[string]interface{}** | IdP-specific profile for the user.  IdP user profiles are IdP-specific but may be customized by the Profile Editor in the Admin Console.  &gt; **Note:** Okta variable names have reserved characters that may conflict with the name of an IdP assertion attribute. You can use the **External name** to define the attribute name as defined in an IdP assertion such as a SAML attribute name. | [optional] 
**Embedded** | Pointer to **map[string]map[string]interface{}** | Embedded resources related to the IdP user | [optional] [readonly] 
**Links** | Pointer to [**IdentityProviderApplicationUserLinks**](IdentityProviderApplicationUserLinks.md) |  | [optional] 

## Methods

### NewIdentityProviderApplicationUser

`func NewIdentityProviderApplicationUser() *IdentityProviderApplicationUser`

NewIdentityProviderApplicationUser instantiates a new IdentityProviderApplicationUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderApplicationUserWithDefaults

`func NewIdentityProviderApplicationUserWithDefaults() *IdentityProviderApplicationUser`

NewIdentityProviderApplicationUserWithDefaults instantiates a new IdentityProviderApplicationUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *IdentityProviderApplicationUser) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IdentityProviderApplicationUser) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IdentityProviderApplicationUser) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IdentityProviderApplicationUser) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExternalId

`func (o *IdentityProviderApplicationUser) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *IdentityProviderApplicationUser) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *IdentityProviderApplicationUser) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *IdentityProviderApplicationUser) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *IdentityProviderApplicationUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProviderApplicationUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProviderApplicationUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProviderApplicationUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *IdentityProviderApplicationUser) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *IdentityProviderApplicationUser) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *IdentityProviderApplicationUser) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *IdentityProviderApplicationUser) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProfile

`func (o *IdentityProviderApplicationUser) GetProfile() map[string]map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *IdentityProviderApplicationUser) GetProfileOk() (*map[string]map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *IdentityProviderApplicationUser) SetProfile(v map[string]map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *IdentityProviderApplicationUser) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetEmbedded

`func (o *IdentityProviderApplicationUser) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *IdentityProviderApplicationUser) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *IdentityProviderApplicationUser) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *IdentityProviderApplicationUser) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *IdentityProviderApplicationUser) GetLinks() IdentityProviderApplicationUserLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IdentityProviderApplicationUser) GetLinksOk() (*IdentityProviderApplicationUserLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IdentityProviderApplicationUser) SetLinks(v IdentityProviderApplicationUserLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IdentityProviderApplicationUser) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


