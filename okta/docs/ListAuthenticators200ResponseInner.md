# ListAuthenticators200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the Authenticator was created | [optional] [readonly] 
**Id** | Pointer to **string** | A unique identifier for the Authenticator | [optional] [readonly] 
**Key** | Pointer to **string** | A human-readable string that identifies the Authenticator | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the Authenticator was last modified | [optional] [readonly] 
**Name** | Pointer to **string** | Display name of the Authenticator | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The type of Authenticator | [optional] 
**Links** | Pointer to [**AuthenticatorLinks**](AuthenticatorLinks.md) |  | [optional] 
**Provider** | Pointer to [**AuthenticatorKeyDuoAllOfProvider**](AuthenticatorKeyDuoAllOfProvider.md) |  | [optional] 
**Settings** | Pointer to [**AuthenticatorKeyPhoneAllOfSettings**](AuthenticatorKeyPhoneAllOfSettings.md) |  | [optional] 

## Methods

### NewListAuthenticators200ResponseInner

`func NewListAuthenticators200ResponseInner() *ListAuthenticators200ResponseInner`

NewListAuthenticators200ResponseInner instantiates a new ListAuthenticators200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthenticators200ResponseInnerWithDefaults

`func NewListAuthenticators200ResponseInnerWithDefaults() *ListAuthenticators200ResponseInner`

NewListAuthenticators200ResponseInnerWithDefaults instantiates a new ListAuthenticators200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ListAuthenticators200ResponseInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListAuthenticators200ResponseInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListAuthenticators200ResponseInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListAuthenticators200ResponseInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *ListAuthenticators200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAuthenticators200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAuthenticators200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListAuthenticators200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *ListAuthenticators200ResponseInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ListAuthenticators200ResponseInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ListAuthenticators200ResponseInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ListAuthenticators200ResponseInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListAuthenticators200ResponseInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListAuthenticators200ResponseInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListAuthenticators200ResponseInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListAuthenticators200ResponseInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *ListAuthenticators200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListAuthenticators200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListAuthenticators200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListAuthenticators200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ListAuthenticators200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListAuthenticators200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListAuthenticators200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListAuthenticators200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ListAuthenticators200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListAuthenticators200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListAuthenticators200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListAuthenticators200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *ListAuthenticators200ResponseInner) GetLinks() AuthenticatorLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListAuthenticators200ResponseInner) GetLinksOk() (*AuthenticatorLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListAuthenticators200ResponseInner) SetLinks(v AuthenticatorLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListAuthenticators200ResponseInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetProvider

`func (o *ListAuthenticators200ResponseInner) GetProvider() AuthenticatorKeyDuoAllOfProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ListAuthenticators200ResponseInner) GetProviderOk() (*AuthenticatorKeyDuoAllOfProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ListAuthenticators200ResponseInner) SetProvider(v AuthenticatorKeyDuoAllOfProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ListAuthenticators200ResponseInner) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSettings

`func (o *ListAuthenticators200ResponseInner) GetSettings() AuthenticatorKeyPhoneAllOfSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ListAuthenticators200ResponseInner) GetSettingsOk() (*AuthenticatorKeyPhoneAllOfSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ListAuthenticators200ResponseInner) SetSettings(v AuthenticatorKeyPhoneAllOfSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ListAuthenticators200ResponseInner) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


