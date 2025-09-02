# ChildOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | [**OrgCreationAdmin**](OrgCreationAdmin.md) |  | 
**Created** | Pointer to **time.Time** | Timestamp when the org was created | [optional] [readonly] 
**Edition** | **string** | Edition for the org. &#x60;SKU&#x60; is the only supported value. | 
**Id** | Pointer to **string** | Org ID | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the org was last updated | [optional] [readonly] 
**Name** | **string** | Unique name of the org. This name appears in the HTML &#x60;&lt;title&gt;&#x60; tag of the new org sign-in page. Only less than 4-width UTF-8 encoded characters are allowed. | 
**Settings** | Pointer to **map[string]interface{}** | Settings associated with the created org | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the org. &#x60;ACTIVE&#x60; is returned after the org is created. | [optional] [readonly] 
**Subdomain** | **string** | Subdomain of the org. Must be unique and include no spaces. | 
**Token** | Pointer to **string** | API token associated with the child org super admin account. Use this API token to provision resources (such as policies, apps, and groups) on the newly created child org. This token is revoked if the super admin account is deactivated. &gt; **Note:** If this API token expires, sign in to the Admin Console as the super admin user and create a new API token. See [Create an API token](https://developer.okta.com/docs/guides/create-an-api-token/). | [optional] [readonly] 
**TokenType** | Pointer to **string** | Type of returned &#x60;token&#x60;. See [Okta API tokens](https://developer.okta.com/docs/guides/create-an-api-token/main/#okta-api-tokens). | [optional] [readonly] 
**Website** | Pointer to **string** | Default website for the org | [optional] 
**Links** | Pointer to **map[string]interface{}** | Specifies available link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification | [optional] [readonly] 

## Methods

### NewChildOrg

`func NewChildOrg(admin OrgCreationAdmin, edition string, name string, subdomain string, ) *ChildOrg`

NewChildOrg instantiates a new ChildOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChildOrgWithDefaults

`func NewChildOrgWithDefaults() *ChildOrg`

NewChildOrgWithDefaults instantiates a new ChildOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *ChildOrg) GetAdmin() OrgCreationAdmin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ChildOrg) GetAdminOk() (*OrgCreationAdmin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ChildOrg) SetAdmin(v OrgCreationAdmin)`

SetAdmin sets Admin field to given value.


### GetCreated

`func (o *ChildOrg) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ChildOrg) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ChildOrg) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ChildOrg) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEdition

`func (o *ChildOrg) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *ChildOrg) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *ChildOrg) SetEdition(v string)`

SetEdition sets Edition field to given value.


### GetId

`func (o *ChildOrg) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChildOrg) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChildOrg) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChildOrg) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ChildOrg) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ChildOrg) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ChildOrg) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ChildOrg) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *ChildOrg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChildOrg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChildOrg) SetName(v string)`

SetName sets Name field to given value.


### GetSettings

`func (o *ChildOrg) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ChildOrg) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ChildOrg) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ChildOrg) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetStatus

`func (o *ChildOrg) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChildOrg) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChildOrg) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChildOrg) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubdomain

`func (o *ChildOrg) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *ChildOrg) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *ChildOrg) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.


### GetToken

`func (o *ChildOrg) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ChildOrg) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ChildOrg) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ChildOrg) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenType

`func (o *ChildOrg) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *ChildOrg) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *ChildOrg) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *ChildOrg) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetWebsite

`func (o *ChildOrg) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ChildOrg) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ChildOrg) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *ChildOrg) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetLinks

`func (o *ChildOrg) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ChildOrg) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ChildOrg) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ChildOrg) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


