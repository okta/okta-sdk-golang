# OAuth2ScopeConsentGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Client ID of the app integration | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**CreatedBy** | Pointer to [**OAuth2Actor**](OAuth2Actor.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the Grant object | [optional] [readonly] 
**Issuer** | **string** | The issuer of your org authorization server. This is typically your Okta domain. | 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**ScopeId** | **string** | The name of the [Okta scope](https://developer.okta.com/docs/api/oauth2/#oauth-20-scopes) for which consent is granted | 
**Source** | Pointer to **string** | User type source that granted consent | [optional] [readonly] 
**Status** | Pointer to **string** | Status | [optional] [readonly] 
**UserId** | Pointer to **string** | User ID that granted consent (if &#x60;source&#x60; is &#x60;END_USER&#x60;) | [optional] [readonly] 
**Embedded** | Pointer to [**OAuth2ScopeConsentGrantEmbedded**](OAuth2ScopeConsentGrantEmbedded.md) |  | [optional] 
**Links** | Pointer to [**OAuth2ScopeConsentGrantLinks**](OAuth2ScopeConsentGrantLinks.md) |  | [optional] 

## Methods

### NewOAuth2ScopeConsentGrant

`func NewOAuth2ScopeConsentGrant(issuer string, scopeId string, ) *OAuth2ScopeConsentGrant`

NewOAuth2ScopeConsentGrant instantiates a new OAuth2ScopeConsentGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ScopeConsentGrantWithDefaults

`func NewOAuth2ScopeConsentGrantWithDefaults() *OAuth2ScopeConsentGrant`

NewOAuth2ScopeConsentGrantWithDefaults instantiates a new OAuth2ScopeConsentGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuth2ScopeConsentGrant) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2ScopeConsentGrant) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2ScopeConsentGrant) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuth2ScopeConsentGrant) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCreated

`func (o *OAuth2ScopeConsentGrant) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OAuth2ScopeConsentGrant) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OAuth2ScopeConsentGrant) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OAuth2ScopeConsentGrant) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *OAuth2ScopeConsentGrant) GetCreatedBy() OAuth2Actor`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OAuth2ScopeConsentGrant) GetCreatedByOk() (*OAuth2Actor, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OAuth2ScopeConsentGrant) SetCreatedBy(v OAuth2Actor)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *OAuth2ScopeConsentGrant) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetId

`func (o *OAuth2ScopeConsentGrant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2ScopeConsentGrant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2ScopeConsentGrant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuth2ScopeConsentGrant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuer

`func (o *OAuth2ScopeConsentGrant) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OAuth2ScopeConsentGrant) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OAuth2ScopeConsentGrant) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetLastUpdated

`func (o *OAuth2ScopeConsentGrant) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *OAuth2ScopeConsentGrant) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *OAuth2ScopeConsentGrant) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *OAuth2ScopeConsentGrant) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetScopeId

`func (o *OAuth2ScopeConsentGrant) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *OAuth2ScopeConsentGrant) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *OAuth2ScopeConsentGrant) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.


### GetSource

`func (o *OAuth2ScopeConsentGrant) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OAuth2ScopeConsentGrant) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OAuth2ScopeConsentGrant) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *OAuth2ScopeConsentGrant) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *OAuth2ScopeConsentGrant) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2ScopeConsentGrant) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2ScopeConsentGrant) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2ScopeConsentGrant) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *OAuth2ScopeConsentGrant) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OAuth2ScopeConsentGrant) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OAuth2ScopeConsentGrant) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OAuth2ScopeConsentGrant) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEmbedded

`func (o *OAuth2ScopeConsentGrant) GetEmbedded() OAuth2ScopeConsentGrantEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *OAuth2ScopeConsentGrant) GetEmbeddedOk() (*OAuth2ScopeConsentGrantEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *OAuth2ScopeConsentGrant) SetEmbedded(v OAuth2ScopeConsentGrantEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *OAuth2ScopeConsentGrant) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *OAuth2ScopeConsentGrant) GetLinks() OAuth2ScopeConsentGrantLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OAuth2ScopeConsentGrant) GetLinksOk() (*OAuth2ScopeConsentGrantLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OAuth2ScopeConsentGrant) SetLinks(v OAuth2ScopeConsentGrantLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OAuth2ScopeConsentGrant) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


