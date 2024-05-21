# APIServiceIntegrationInstanceSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | **string** | The OAuth 2.0 client secret string. The client secret string is returned in the response of a Secret creation request. In other responses (such as list, activate, or deactivate requests), the client secret is returned as an undisclosed hashed value. | [readonly] 
**Created** | **string** | Timestamp when the API Service Integration instance Secret was created | [readonly] 
**Id** | **string** | The ID of the API Service Integration instance Secret | [readonly] 
**LastUpdated** | **string** | Timestamp when the API Service Integration instance Secret was updated | [readonly] 
**SecretHash** | **string** | OAuth 2.0 client secret string hash | [readonly] 
**Status** | **string** | Status of the API Service Integration instance Secret | 
**Links** | [**APIServiceIntegrationSecretLinks**](APIServiceIntegrationSecretLinks.md) |  | 

## Methods

### NewAPIServiceIntegrationInstanceSecret

`func NewAPIServiceIntegrationInstanceSecret(clientSecret string, created string, id string, lastUpdated string, secretHash string, status string, links APIServiceIntegrationSecretLinks, ) *APIServiceIntegrationInstanceSecret`

NewAPIServiceIntegrationInstanceSecret instantiates a new APIServiceIntegrationInstanceSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServiceIntegrationInstanceSecretWithDefaults

`func NewAPIServiceIntegrationInstanceSecretWithDefaults() *APIServiceIntegrationInstanceSecret`

NewAPIServiceIntegrationInstanceSecretWithDefaults instantiates a new APIServiceIntegrationInstanceSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *APIServiceIntegrationInstanceSecret) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *APIServiceIntegrationInstanceSecret) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *APIServiceIntegrationInstanceSecret) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetCreated

`func (o *APIServiceIntegrationInstanceSecret) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *APIServiceIntegrationInstanceSecret) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *APIServiceIntegrationInstanceSecret) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetId

`func (o *APIServiceIntegrationInstanceSecret) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIServiceIntegrationInstanceSecret) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIServiceIntegrationInstanceSecret) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdated

`func (o *APIServiceIntegrationInstanceSecret) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *APIServiceIntegrationInstanceSecret) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *APIServiceIntegrationInstanceSecret) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.


### GetSecretHash

`func (o *APIServiceIntegrationInstanceSecret) GetSecretHash() string`

GetSecretHash returns the SecretHash field if non-nil, zero value otherwise.

### GetSecretHashOk

`func (o *APIServiceIntegrationInstanceSecret) GetSecretHashOk() (*string, bool)`

GetSecretHashOk returns a tuple with the SecretHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretHash

`func (o *APIServiceIntegrationInstanceSecret) SetSecretHash(v string)`

SetSecretHash sets SecretHash field to given value.


### GetStatus

`func (o *APIServiceIntegrationInstanceSecret) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *APIServiceIntegrationInstanceSecret) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *APIServiceIntegrationInstanceSecret) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLinks

`func (o *APIServiceIntegrationInstanceSecret) GetLinks() APIServiceIntegrationSecretLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *APIServiceIntegrationInstanceSecret) GetLinksOk() (*APIServiceIntegrationSecretLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *APIServiceIntegrationInstanceSecret) SetLinks(v APIServiceIntegrationSecretLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


