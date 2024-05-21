# SwaApplicationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityStoreId** | Pointer to **string** |  | [optional] 
**ImplicitAssignment** | Pointer to **bool** |  | [optional] 
**InlineHookId** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to [**ApplicationSettingsNotes**](ApplicationSettingsNotes.md) |  | [optional] 
**Notifications** | Pointer to [**ApplicationSettingsNotifications**](ApplicationSettingsNotifications.md) |  | [optional] 
**App** | Pointer to [**SwaApplicationSettingsApplication**](SwaApplicationSettingsApplication.md) |  | [optional] 

## Methods

### NewSwaApplicationSettings

`func NewSwaApplicationSettings() *SwaApplicationSettings`

NewSwaApplicationSettings instantiates a new SwaApplicationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwaApplicationSettingsWithDefaults

`func NewSwaApplicationSettingsWithDefaults() *SwaApplicationSettings`

NewSwaApplicationSettingsWithDefaults instantiates a new SwaApplicationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityStoreId

`func (o *SwaApplicationSettings) GetIdentityStoreId() string`

GetIdentityStoreId returns the IdentityStoreId field if non-nil, zero value otherwise.

### GetIdentityStoreIdOk

`func (o *SwaApplicationSettings) GetIdentityStoreIdOk() (*string, bool)`

GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStoreId

`func (o *SwaApplicationSettings) SetIdentityStoreId(v string)`

SetIdentityStoreId sets IdentityStoreId field to given value.

### HasIdentityStoreId

`func (o *SwaApplicationSettings) HasIdentityStoreId() bool`

HasIdentityStoreId returns a boolean if a field has been set.

### GetImplicitAssignment

`func (o *SwaApplicationSettings) GetImplicitAssignment() bool`

GetImplicitAssignment returns the ImplicitAssignment field if non-nil, zero value otherwise.

### GetImplicitAssignmentOk

`func (o *SwaApplicationSettings) GetImplicitAssignmentOk() (*bool, bool)`

GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitAssignment

`func (o *SwaApplicationSettings) SetImplicitAssignment(v bool)`

SetImplicitAssignment sets ImplicitAssignment field to given value.

### HasImplicitAssignment

`func (o *SwaApplicationSettings) HasImplicitAssignment() bool`

HasImplicitAssignment returns a boolean if a field has been set.

### GetInlineHookId

`func (o *SwaApplicationSettings) GetInlineHookId() string`

GetInlineHookId returns the InlineHookId field if non-nil, zero value otherwise.

### GetInlineHookIdOk

`func (o *SwaApplicationSettings) GetInlineHookIdOk() (*string, bool)`

GetInlineHookIdOk returns a tuple with the InlineHookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineHookId

`func (o *SwaApplicationSettings) SetInlineHookId(v string)`

SetInlineHookId sets InlineHookId field to given value.

### HasInlineHookId

`func (o *SwaApplicationSettings) HasInlineHookId() bool`

HasInlineHookId returns a boolean if a field has been set.

### GetNotes

`func (o *SwaApplicationSettings) GetNotes() ApplicationSettingsNotes`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SwaApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SwaApplicationSettings) SetNotes(v ApplicationSettingsNotes)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SwaApplicationSettings) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNotifications

`func (o *SwaApplicationSettings) GetNotifications() ApplicationSettingsNotifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *SwaApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *SwaApplicationSettings) SetNotifications(v ApplicationSettingsNotifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *SwaApplicationSettings) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetApp

`func (o *SwaApplicationSettings) GetApp() SwaApplicationSettingsApplication`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *SwaApplicationSettings) GetAppOk() (*SwaApplicationSettingsApplication, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *SwaApplicationSettings) SetApp(v SwaApplicationSettingsApplication)`

SetApp sets App field to given value.

### HasApp

`func (o *SwaApplicationSettings) HasApp() bool`

HasApp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


