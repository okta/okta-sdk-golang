# PostAPIServiceIntegrationInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantedScopes** | **[]string** | The list of Okta management scopes granted to the API Service Integration instance. See [Okta management OAuth 2.0 scopes](/oauth2/#okta-admin-management). | 
**Type** | **string** | The type of the API service integration. This string is an underscore-concatenated, lowercased API service integration name. For example, &#x60;my_api_log_integration&#x60;. | 

## Methods

### NewPostAPIServiceIntegrationInstanceRequest

`func NewPostAPIServiceIntegrationInstanceRequest(grantedScopes []string, type_ string, ) *PostAPIServiceIntegrationInstanceRequest`

NewPostAPIServiceIntegrationInstanceRequest instantiates a new PostAPIServiceIntegrationInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAPIServiceIntegrationInstanceRequestWithDefaults

`func NewPostAPIServiceIntegrationInstanceRequestWithDefaults() *PostAPIServiceIntegrationInstanceRequest`

NewPostAPIServiceIntegrationInstanceRequestWithDefaults instantiates a new PostAPIServiceIntegrationInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantedScopes

`func (o *PostAPIServiceIntegrationInstanceRequest) GetGrantedScopes() []string`

GetGrantedScopes returns the GrantedScopes field if non-nil, zero value otherwise.

### GetGrantedScopesOk

`func (o *PostAPIServiceIntegrationInstanceRequest) GetGrantedScopesOk() (*[]string, bool)`

GetGrantedScopesOk returns a tuple with the GrantedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedScopes

`func (o *PostAPIServiceIntegrationInstanceRequest) SetGrantedScopes(v []string)`

SetGrantedScopes sets GrantedScopes field to given value.


### GetType

`func (o *PostAPIServiceIntegrationInstanceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostAPIServiceIntegrationInstanceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostAPIServiceIntegrationInstanceRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


