# OAuth2ClientSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | Pointer to **string** | The OAuth 2.0 client secret string | [optional] [readonly] 
**Created** | Pointer to **string** | Timestamp when the OAuth Client 2.0 Secret was created | [optional] [readonly] 
**Id** | Pointer to **string** | The unique ID of the OAuth Client Secret | [optional] [readonly] 
**LastUpdated** | Pointer to **string** | Timestamp when the OAuth Client 2.0 Secret was updated | [optional] [readonly] 
**SecretHash** | Pointer to **string** | OAuth 2.0 client secret string hash | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 Client Secret | [optional] [default to "ACTIVE"]
**Links** | Pointer to [**OAuthClientSecretLinks**](OAuthClientSecretLinks.md) |  | [optional] 

## Methods

### NewOAuth2ClientSecret

`func NewOAuth2ClientSecret() *OAuth2ClientSecret`

NewOAuth2ClientSecret instantiates a new OAuth2ClientSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientSecretWithDefaults

`func NewOAuth2ClientSecretWithDefaults() *OAuth2ClientSecret`

NewOAuth2ClientSecretWithDefaults instantiates a new OAuth2ClientSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *OAuth2ClientSecret) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuth2ClientSecret) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuth2ClientSecret) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuth2ClientSecret) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCreated

`func (o *OAuth2ClientSecret) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OAuth2ClientSecret) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OAuth2ClientSecret) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OAuth2ClientSecret) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *OAuth2ClientSecret) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2ClientSecret) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2ClientSecret) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuth2ClientSecret) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *OAuth2ClientSecret) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *OAuth2ClientSecret) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *OAuth2ClientSecret) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *OAuth2ClientSecret) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetSecretHash

`func (o *OAuth2ClientSecret) GetSecretHash() string`

GetSecretHash returns the SecretHash field if non-nil, zero value otherwise.

### GetSecretHashOk

`func (o *OAuth2ClientSecret) GetSecretHashOk() (*string, bool)`

GetSecretHashOk returns a tuple with the SecretHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretHash

`func (o *OAuth2ClientSecret) SetSecretHash(v string)`

SetSecretHash sets SecretHash field to given value.

### HasSecretHash

`func (o *OAuth2ClientSecret) HasSecretHash() bool`

HasSecretHash returns a boolean if a field has been set.

### GetStatus

`func (o *OAuth2ClientSecret) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2ClientSecret) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2ClientSecret) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2ClientSecret) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *OAuth2ClientSecret) GetLinks() OAuthClientSecretLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OAuth2ClientSecret) GetLinksOk() (*OAuthClientSecretLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OAuth2ClientSecret) SetLinks(v OAuthClientSecretLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OAuth2ClientSecret) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


