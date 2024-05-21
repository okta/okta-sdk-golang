# OpenIdConnectApplicationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityStoreId** | Pointer to **string** |  | [optional] 
**ImplicitAssignment** | Pointer to **bool** |  | [optional] 
**InlineHookId** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to [**ApplicationSettingsNotes**](ApplicationSettingsNotes.md) |  | [optional] 
**Notifications** | Pointer to [**ApplicationSettingsNotifications**](ApplicationSettingsNotifications.md) |  | [optional] 
**OauthClient** | Pointer to [**OpenIdConnectApplicationSettingsClient**](OpenIdConnectApplicationSettingsClient.md) |  | [optional] 

## Methods

### NewOpenIdConnectApplicationSettings

`func NewOpenIdConnectApplicationSettings() *OpenIdConnectApplicationSettings`

NewOpenIdConnectApplicationSettings instantiates a new OpenIdConnectApplicationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConnectApplicationSettingsWithDefaults

`func NewOpenIdConnectApplicationSettingsWithDefaults() *OpenIdConnectApplicationSettings`

NewOpenIdConnectApplicationSettingsWithDefaults instantiates a new OpenIdConnectApplicationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityStoreId

`func (o *OpenIdConnectApplicationSettings) GetIdentityStoreId() string`

GetIdentityStoreId returns the IdentityStoreId field if non-nil, zero value otherwise.

### GetIdentityStoreIdOk

`func (o *OpenIdConnectApplicationSettings) GetIdentityStoreIdOk() (*string, bool)`

GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStoreId

`func (o *OpenIdConnectApplicationSettings) SetIdentityStoreId(v string)`

SetIdentityStoreId sets IdentityStoreId field to given value.

### HasIdentityStoreId

`func (o *OpenIdConnectApplicationSettings) HasIdentityStoreId() bool`

HasIdentityStoreId returns a boolean if a field has been set.

### GetImplicitAssignment

`func (o *OpenIdConnectApplicationSettings) GetImplicitAssignment() bool`

GetImplicitAssignment returns the ImplicitAssignment field if non-nil, zero value otherwise.

### GetImplicitAssignmentOk

`func (o *OpenIdConnectApplicationSettings) GetImplicitAssignmentOk() (*bool, bool)`

GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitAssignment

`func (o *OpenIdConnectApplicationSettings) SetImplicitAssignment(v bool)`

SetImplicitAssignment sets ImplicitAssignment field to given value.

### HasImplicitAssignment

`func (o *OpenIdConnectApplicationSettings) HasImplicitAssignment() bool`

HasImplicitAssignment returns a boolean if a field has been set.

### GetInlineHookId

`func (o *OpenIdConnectApplicationSettings) GetInlineHookId() string`

GetInlineHookId returns the InlineHookId field if non-nil, zero value otherwise.

### GetInlineHookIdOk

`func (o *OpenIdConnectApplicationSettings) GetInlineHookIdOk() (*string, bool)`

GetInlineHookIdOk returns a tuple with the InlineHookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineHookId

`func (o *OpenIdConnectApplicationSettings) SetInlineHookId(v string)`

SetInlineHookId sets InlineHookId field to given value.

### HasInlineHookId

`func (o *OpenIdConnectApplicationSettings) HasInlineHookId() bool`

HasInlineHookId returns a boolean if a field has been set.

### GetNotes

`func (o *OpenIdConnectApplicationSettings) GetNotes() ApplicationSettingsNotes`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OpenIdConnectApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OpenIdConnectApplicationSettings) SetNotes(v ApplicationSettingsNotes)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OpenIdConnectApplicationSettings) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNotifications

`func (o *OpenIdConnectApplicationSettings) GetNotifications() ApplicationSettingsNotifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *OpenIdConnectApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *OpenIdConnectApplicationSettings) SetNotifications(v ApplicationSettingsNotifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *OpenIdConnectApplicationSettings) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetOauthClient

`func (o *OpenIdConnectApplicationSettings) GetOauthClient() OpenIdConnectApplicationSettingsClient`

GetOauthClient returns the OauthClient field if non-nil, zero value otherwise.

### GetOauthClientOk

`func (o *OpenIdConnectApplicationSettings) GetOauthClientOk() (*OpenIdConnectApplicationSettingsClient, bool)`

GetOauthClientOk returns a tuple with the OauthClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClient

`func (o *OpenIdConnectApplicationSettings) SetOauthClient(v OpenIdConnectApplicationSettingsClient)`

SetOauthClient sets OauthClient field to given value.

### HasOauthClient

`func (o *OpenIdConnectApplicationSettings) HasOauthClient() bool`

HasOauthClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


