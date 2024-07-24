# APIServiceIntegrationInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigGuideUrl** | Pointer to **string** | The URL to the API service integration configuration guide | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | Timestamp when the API Service Integration instance was created | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | The user ID of the API Service Integration instance creator | [optional] [readonly] 
**GrantedScopes** | Pointer to **[]string** | The list of Okta management scopes granted to the API Service Integration instance. See [Okta management OAuth 2.0 scopes](/oauth2/#okta-admin-management). | [optional] 
**Id** | Pointer to **string** | The ID of the API Service Integration instance | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the API service integration that corresponds with the &#x60;type&#x60; property. This is the full name of the API service integration listed in the Okta Integration Network (OIN) catalog. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the API service integration. This string is an underscore-concatenated, lowercased API service integration name. For example, &#x60;my_api_log_integration&#x60;. | [optional] 
**Links** | Pointer to [**APIServiceIntegrationLinks**](APIServiceIntegrationLinks.md) |  | [optional] 

## Methods

### NewAPIServiceIntegrationInstance

`func NewAPIServiceIntegrationInstance() *APIServiceIntegrationInstance`

NewAPIServiceIntegrationInstance instantiates a new APIServiceIntegrationInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServiceIntegrationInstanceWithDefaults

`func NewAPIServiceIntegrationInstanceWithDefaults() *APIServiceIntegrationInstance`

NewAPIServiceIntegrationInstanceWithDefaults instantiates a new APIServiceIntegrationInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigGuideUrl

`func (o *APIServiceIntegrationInstance) GetConfigGuideUrl() string`

GetConfigGuideUrl returns the ConfigGuideUrl field if non-nil, zero value otherwise.

### GetConfigGuideUrlOk

`func (o *APIServiceIntegrationInstance) GetConfigGuideUrlOk() (*string, bool)`

GetConfigGuideUrlOk returns a tuple with the ConfigGuideUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigGuideUrl

`func (o *APIServiceIntegrationInstance) SetConfigGuideUrl(v string)`

SetConfigGuideUrl sets ConfigGuideUrl field to given value.

### HasConfigGuideUrl

`func (o *APIServiceIntegrationInstance) HasConfigGuideUrl() bool`

HasConfigGuideUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *APIServiceIntegrationInstance) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APIServiceIntegrationInstance) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APIServiceIntegrationInstance) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *APIServiceIntegrationInstance) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *APIServiceIntegrationInstance) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *APIServiceIntegrationInstance) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *APIServiceIntegrationInstance) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *APIServiceIntegrationInstance) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetGrantedScopes

`func (o *APIServiceIntegrationInstance) GetGrantedScopes() []string`

GetGrantedScopes returns the GrantedScopes field if non-nil, zero value otherwise.

### GetGrantedScopesOk

`func (o *APIServiceIntegrationInstance) GetGrantedScopesOk() (*[]string, bool)`

GetGrantedScopesOk returns a tuple with the GrantedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedScopes

`func (o *APIServiceIntegrationInstance) SetGrantedScopes(v []string)`

SetGrantedScopes sets GrantedScopes field to given value.

### HasGrantedScopes

`func (o *APIServiceIntegrationInstance) HasGrantedScopes() bool`

HasGrantedScopes returns a boolean if a field has been set.

### GetId

`func (o *APIServiceIntegrationInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIServiceIntegrationInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIServiceIntegrationInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *APIServiceIntegrationInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *APIServiceIntegrationInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APIServiceIntegrationInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APIServiceIntegrationInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *APIServiceIntegrationInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *APIServiceIntegrationInstance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *APIServiceIntegrationInstance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *APIServiceIntegrationInstance) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *APIServiceIntegrationInstance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *APIServiceIntegrationInstance) GetLinks() APIServiceIntegrationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *APIServiceIntegrationInstance) GetLinksOk() (*APIServiceIntegrationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *APIServiceIntegrationInstance) SetLinks(v APIServiceIntegrationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *APIServiceIntegrationInstance) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


