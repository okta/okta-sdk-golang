# OAuth2RefreshTokenScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the Scope | [optional] 
**DisplayName** | Pointer to **string** | Name of the end user displayed in a consent dialog | [optional] 
**Id** | Pointer to **string** | Scope object ID | [optional] [readonly] 
**Name** | Pointer to **string** | Scope name | [optional] 
**Links** | Pointer to [**OAuth2RefreshTokenScopeLinks**](OAuth2RefreshTokenScopeLinks.md) |  | [optional] 

## Methods

### NewOAuth2RefreshTokenScope

`func NewOAuth2RefreshTokenScope() *OAuth2RefreshTokenScope`

NewOAuth2RefreshTokenScope instantiates a new OAuth2RefreshTokenScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2RefreshTokenScopeWithDefaults

`func NewOAuth2RefreshTokenScopeWithDefaults() *OAuth2RefreshTokenScope`

NewOAuth2RefreshTokenScopeWithDefaults instantiates a new OAuth2RefreshTokenScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OAuth2RefreshTokenScope) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OAuth2RefreshTokenScope) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OAuth2RefreshTokenScope) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OAuth2RefreshTokenScope) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *OAuth2RefreshTokenScope) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OAuth2RefreshTokenScope) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OAuth2RefreshTokenScope) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OAuth2RefreshTokenScope) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *OAuth2RefreshTokenScope) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2RefreshTokenScope) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2RefreshTokenScope) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuth2RefreshTokenScope) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OAuth2RefreshTokenScope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuth2RefreshTokenScope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuth2RefreshTokenScope) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OAuth2RefreshTokenScope) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLinks

`func (o *OAuth2RefreshTokenScope) GetLinks() OAuth2RefreshTokenScopeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OAuth2RefreshTokenScope) GetLinksOk() (*OAuth2RefreshTokenScopeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OAuth2RefreshTokenScope) SetLinks(v OAuth2RefreshTokenScopeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OAuth2RefreshTokenScope) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


