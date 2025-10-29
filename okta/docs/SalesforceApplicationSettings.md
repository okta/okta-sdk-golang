# SalesforceApplicationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmOptInStatus** | Pointer to **string** | The Governance Engine opt-in status for the app | [optional] [readonly] 
**IdentityStoreId** | Pointer to **string** | Identifies an additional identity store app, if your app supports it. The &#x60;identityStoreId&#x60; value must be a valid identity store app ID. This identity store app must be created in the same org as your app. | [optional] 
**ImplicitAssignment** | Pointer to **bool** | Controls whether Okta automatically assigns users to the app based on the user&#39;s role or group membership. | [optional] 
**InlineHookId** | Pointer to **string** | Identifier of an inline hook. Inline hooks are outbound calls from Okta to your own custom code, triggered at specific points in Okta process flows. They allow you to integrate custom functionality into those flows. See [Inline hooks](/openapi/okta-management/management/tag/InlineHook/). | [optional] 
**Notes** | Pointer to [**ApplicationSettingsNotes**](ApplicationSettingsNotes.md) |  | [optional] 
**Notifications** | Pointer to [**ApplicationSettingsNotifications**](ApplicationSettingsNotifications.md) |  | [optional] 
**App** | [**SalesforceApplicationSettingsApplication**](SalesforceApplicationSettingsApplication.md) |  | 
**SignOn** | Pointer to [**OINSaml20ApplicationSettingsSignOn**](OINSaml20ApplicationSettingsSignOn.md) |  | [optional] 

## Methods

### NewSalesforceApplicationSettings

`func NewSalesforceApplicationSettings(app SalesforceApplicationSettingsApplication, ) *SalesforceApplicationSettings`

NewSalesforceApplicationSettings instantiates a new SalesforceApplicationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceApplicationSettingsWithDefaults

`func NewSalesforceApplicationSettingsWithDefaults() *SalesforceApplicationSettings`

NewSalesforceApplicationSettingsWithDefaults instantiates a new SalesforceApplicationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmOptInStatus

`func (o *SalesforceApplicationSettings) GetEmOptInStatus() string`

GetEmOptInStatus returns the EmOptInStatus field if non-nil, zero value otherwise.

### GetEmOptInStatusOk

`func (o *SalesforceApplicationSettings) GetEmOptInStatusOk() (*string, bool)`

GetEmOptInStatusOk returns a tuple with the EmOptInStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmOptInStatus

`func (o *SalesforceApplicationSettings) SetEmOptInStatus(v string)`

SetEmOptInStatus sets EmOptInStatus field to given value.

### HasEmOptInStatus

`func (o *SalesforceApplicationSettings) HasEmOptInStatus() bool`

HasEmOptInStatus returns a boolean if a field has been set.

### GetIdentityStoreId

`func (o *SalesforceApplicationSettings) GetIdentityStoreId() string`

GetIdentityStoreId returns the IdentityStoreId field if non-nil, zero value otherwise.

### GetIdentityStoreIdOk

`func (o *SalesforceApplicationSettings) GetIdentityStoreIdOk() (*string, bool)`

GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStoreId

`func (o *SalesforceApplicationSettings) SetIdentityStoreId(v string)`

SetIdentityStoreId sets IdentityStoreId field to given value.

### HasIdentityStoreId

`func (o *SalesforceApplicationSettings) HasIdentityStoreId() bool`

HasIdentityStoreId returns a boolean if a field has been set.

### GetImplicitAssignment

`func (o *SalesforceApplicationSettings) GetImplicitAssignment() bool`

GetImplicitAssignment returns the ImplicitAssignment field if non-nil, zero value otherwise.

### GetImplicitAssignmentOk

`func (o *SalesforceApplicationSettings) GetImplicitAssignmentOk() (*bool, bool)`

GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitAssignment

`func (o *SalesforceApplicationSettings) SetImplicitAssignment(v bool)`

SetImplicitAssignment sets ImplicitAssignment field to given value.

### HasImplicitAssignment

`func (o *SalesforceApplicationSettings) HasImplicitAssignment() bool`

HasImplicitAssignment returns a boolean if a field has been set.

### GetInlineHookId

`func (o *SalesforceApplicationSettings) GetInlineHookId() string`

GetInlineHookId returns the InlineHookId field if non-nil, zero value otherwise.

### GetInlineHookIdOk

`func (o *SalesforceApplicationSettings) GetInlineHookIdOk() (*string, bool)`

GetInlineHookIdOk returns a tuple with the InlineHookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineHookId

`func (o *SalesforceApplicationSettings) SetInlineHookId(v string)`

SetInlineHookId sets InlineHookId field to given value.

### HasInlineHookId

`func (o *SalesforceApplicationSettings) HasInlineHookId() bool`

HasInlineHookId returns a boolean if a field has been set.

### GetNotes

`func (o *SalesforceApplicationSettings) GetNotes() ApplicationSettingsNotes`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SalesforceApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SalesforceApplicationSettings) SetNotes(v ApplicationSettingsNotes)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SalesforceApplicationSettings) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNotifications

`func (o *SalesforceApplicationSettings) GetNotifications() ApplicationSettingsNotifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *SalesforceApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *SalesforceApplicationSettings) SetNotifications(v ApplicationSettingsNotifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *SalesforceApplicationSettings) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetApp

`func (o *SalesforceApplicationSettings) GetApp() SalesforceApplicationSettingsApplication`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *SalesforceApplicationSettings) GetAppOk() (*SalesforceApplicationSettingsApplication, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *SalesforceApplicationSettings) SetApp(v SalesforceApplicationSettingsApplication)`

SetApp sets App field to given value.


### GetSignOn

`func (o *SalesforceApplicationSettings) GetSignOn() OINSaml20ApplicationSettingsSignOn`

GetSignOn returns the SignOn field if non-nil, zero value otherwise.

### GetSignOnOk

`func (o *SalesforceApplicationSettings) GetSignOnOk() (*OINSaml20ApplicationSettingsSignOn, bool)`

GetSignOnOk returns a tuple with the SignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOn

`func (o *SalesforceApplicationSettings) SetSignOn(v OINSaml20ApplicationSettingsSignOn)`

SetSignOn sets SignOn field to given value.

### HasSignOn

`func (o *SalesforceApplicationSettings) HasSignOn() bool`

HasSignOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


