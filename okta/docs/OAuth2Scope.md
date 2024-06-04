# OAuth2Scope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consent** | Pointer to **string** | Indicates whether a consent dialog is needed for the Scope | [optional] [default to "IMPLICIT"]
**Default** | Pointer to **bool** | Indicates if this Scope is a default scope | [optional] 
**Description** | Pointer to **string** | Description of the Scope | [optional] 
**DisplayName** | Pointer to **string** | Name of the end user displayed in a consent dialog | [optional] 
**Id** | Pointer to **string** | Scope object ID | [optional] [readonly] 
**MetadataPublish** | Pointer to **string** | Indicates whether the Scope is included in the metadata | [optional] [default to "NO_CLIENTS"]
**Name** | Pointer to **string** | Scope name | [optional] 
**Optional** | Pointer to **bool** |  | [optional] 
**System** | Pointer to **bool** | Indicates if Okta created the Scope | [optional] 

## Methods

### NewOAuth2Scope

`func NewOAuth2Scope() *OAuth2Scope`

NewOAuth2Scope instantiates a new OAuth2Scope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ScopeWithDefaults

`func NewOAuth2ScopeWithDefaults() *OAuth2Scope`

NewOAuth2ScopeWithDefaults instantiates a new OAuth2Scope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsent

`func (o *OAuth2Scope) GetConsent() string`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *OAuth2Scope) GetConsentOk() (*string, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *OAuth2Scope) SetConsent(v string)`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *OAuth2Scope) HasConsent() bool`

HasConsent returns a boolean if a field has been set.

### GetDefault

`func (o *OAuth2Scope) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *OAuth2Scope) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *OAuth2Scope) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *OAuth2Scope) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *OAuth2Scope) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OAuth2Scope) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OAuth2Scope) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OAuth2Scope) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *OAuth2Scope) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OAuth2Scope) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OAuth2Scope) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OAuth2Scope) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *OAuth2Scope) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2Scope) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2Scope) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuth2Scope) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadataPublish

`func (o *OAuth2Scope) GetMetadataPublish() string`

GetMetadataPublish returns the MetadataPublish field if non-nil, zero value otherwise.

### GetMetadataPublishOk

`func (o *OAuth2Scope) GetMetadataPublishOk() (*string, bool)`

GetMetadataPublishOk returns a tuple with the MetadataPublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataPublish

`func (o *OAuth2Scope) SetMetadataPublish(v string)`

SetMetadataPublish sets MetadataPublish field to given value.

### HasMetadataPublish

`func (o *OAuth2Scope) HasMetadataPublish() bool`

HasMetadataPublish returns a boolean if a field has been set.

### GetName

`func (o *OAuth2Scope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuth2Scope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuth2Scope) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OAuth2Scope) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptional

`func (o *OAuth2Scope) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *OAuth2Scope) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *OAuth2Scope) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *OAuth2Scope) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetSystem

`func (o *OAuth2Scope) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *OAuth2Scope) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *OAuth2Scope) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *OAuth2Scope) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


