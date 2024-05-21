# AutoLoginApplicationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityStoreId** | Pointer to **string** |  | [optional] 
**ImplicitAssignment** | Pointer to **bool** |  | [optional] 
**InlineHookId** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to [**ApplicationSettingsNotes**](ApplicationSettingsNotes.md) |  | [optional] 
**Notifications** | Pointer to [**ApplicationSettingsNotifications**](ApplicationSettingsNotifications.md) |  | [optional] 
**SignOn** | Pointer to [**AutoLoginApplicationSettingsSignOn**](AutoLoginApplicationSettingsSignOn.md) |  | [optional] 

## Methods

### NewAutoLoginApplicationSettings

`func NewAutoLoginApplicationSettings() *AutoLoginApplicationSettings`

NewAutoLoginApplicationSettings instantiates a new AutoLoginApplicationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoLoginApplicationSettingsWithDefaults

`func NewAutoLoginApplicationSettingsWithDefaults() *AutoLoginApplicationSettings`

NewAutoLoginApplicationSettingsWithDefaults instantiates a new AutoLoginApplicationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityStoreId

`func (o *AutoLoginApplicationSettings) GetIdentityStoreId() string`

GetIdentityStoreId returns the IdentityStoreId field if non-nil, zero value otherwise.

### GetIdentityStoreIdOk

`func (o *AutoLoginApplicationSettings) GetIdentityStoreIdOk() (*string, bool)`

GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStoreId

`func (o *AutoLoginApplicationSettings) SetIdentityStoreId(v string)`

SetIdentityStoreId sets IdentityStoreId field to given value.

### HasIdentityStoreId

`func (o *AutoLoginApplicationSettings) HasIdentityStoreId() bool`

HasIdentityStoreId returns a boolean if a field has been set.

### GetImplicitAssignment

`func (o *AutoLoginApplicationSettings) GetImplicitAssignment() bool`

GetImplicitAssignment returns the ImplicitAssignment field if non-nil, zero value otherwise.

### GetImplicitAssignmentOk

`func (o *AutoLoginApplicationSettings) GetImplicitAssignmentOk() (*bool, bool)`

GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitAssignment

`func (o *AutoLoginApplicationSettings) SetImplicitAssignment(v bool)`

SetImplicitAssignment sets ImplicitAssignment field to given value.

### HasImplicitAssignment

`func (o *AutoLoginApplicationSettings) HasImplicitAssignment() bool`

HasImplicitAssignment returns a boolean if a field has been set.

### GetInlineHookId

`func (o *AutoLoginApplicationSettings) GetInlineHookId() string`

GetInlineHookId returns the InlineHookId field if non-nil, zero value otherwise.

### GetInlineHookIdOk

`func (o *AutoLoginApplicationSettings) GetInlineHookIdOk() (*string, bool)`

GetInlineHookIdOk returns a tuple with the InlineHookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineHookId

`func (o *AutoLoginApplicationSettings) SetInlineHookId(v string)`

SetInlineHookId sets InlineHookId field to given value.

### HasInlineHookId

`func (o *AutoLoginApplicationSettings) HasInlineHookId() bool`

HasInlineHookId returns a boolean if a field has been set.

### GetNotes

`func (o *AutoLoginApplicationSettings) GetNotes() ApplicationSettingsNotes`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AutoLoginApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AutoLoginApplicationSettings) SetNotes(v ApplicationSettingsNotes)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AutoLoginApplicationSettings) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNotifications

`func (o *AutoLoginApplicationSettings) GetNotifications() ApplicationSettingsNotifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *AutoLoginApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *AutoLoginApplicationSettings) SetNotifications(v ApplicationSettingsNotifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *AutoLoginApplicationSettings) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetSignOn

`func (o *AutoLoginApplicationSettings) GetSignOn() AutoLoginApplicationSettingsSignOn`

GetSignOn returns the SignOn field if non-nil, zero value otherwise.

### GetSignOnOk

`func (o *AutoLoginApplicationSettings) GetSignOnOk() (*AutoLoginApplicationSettingsSignOn, bool)`

GetSignOnOk returns a tuple with the SignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOn

`func (o *AutoLoginApplicationSettings) SetSignOn(v AutoLoginApplicationSettingsSignOn)`

SetSignOn sets SignOn field to given value.

### HasSignOn

`func (o *AutoLoginApplicationSettings) HasSignOn() bool`

HasSignOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


