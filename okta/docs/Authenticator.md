# Authenticator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Key** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to [**AuthenticatorProvider**](AuthenticatorProvider.md) |  | [optional] 
**Settings** | Pointer to [**AuthenticatorSettings**](AuthenticatorSettings.md) |  | [optional] 
**Status** | Pointer to [**AuthenticatorStatus**](AuthenticatorStatus.md) |  | [optional] 
**Type** | Pointer to [**AuthenticatorType**](AuthenticatorType.md) |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewAuthenticator

`func NewAuthenticator() *Authenticator`

NewAuthenticator instantiates a new Authenticator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorWithDefaults

`func NewAuthenticatorWithDefaults() *Authenticator`

NewAuthenticatorWithDefaults instantiates a new Authenticator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Authenticator) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Authenticator) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Authenticator) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Authenticator) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *Authenticator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Authenticator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Authenticator) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Authenticator) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *Authenticator) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Authenticator) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Authenticator) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Authenticator) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Authenticator) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Authenticator) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Authenticator) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Authenticator) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *Authenticator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Authenticator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Authenticator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Authenticator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *Authenticator) GetProvider() AuthenticatorProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Authenticator) GetProviderOk() (*AuthenticatorProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Authenticator) SetProvider(v AuthenticatorProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Authenticator) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSettings

`func (o *Authenticator) GetSettings() AuthenticatorSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Authenticator) GetSettingsOk() (*AuthenticatorSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Authenticator) SetSettings(v AuthenticatorSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Authenticator) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetStatus

`func (o *Authenticator) GetStatus() AuthenticatorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Authenticator) GetStatusOk() (*AuthenticatorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Authenticator) SetStatus(v AuthenticatorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Authenticator) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Authenticator) GetType() AuthenticatorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Authenticator) GetTypeOk() (*AuthenticatorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Authenticator) SetType(v AuthenticatorType)`

SetType sets Type field to given value.

### HasType

`func (o *Authenticator) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *Authenticator) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Authenticator) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Authenticator) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Authenticator) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


