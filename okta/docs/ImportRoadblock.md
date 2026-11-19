# ImportRoadblock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedApps** | Pointer to [**[]ImportRoadblockAffectedApp**](ImportRoadblockAffectedApp.md) | The apps whose assignments the import would remove, each checked against its own app-level threshold. Only returned for app-level roadblocks (&#x60;type&#x60; is &#x60;GROUP_DELETE&#x60; or &#x60;GROUP_UPDATE&#x60; and the threshold crossed is an app-level one).  These are the apps that lose assignments, not the app instance whose import triggered the roadblock (&#x60;appInstanceId&#x60;). For example, deleting a group in an Active Directory source removes the assignments that the group grants to other apps. | [optional] [readonly] 
**AppInstanceId** | **string** | ID of the app instance whose import triggered the roadblock (the import source) | [readonly] 
**AppInstanceLabel** | Pointer to **string** | Label of the app instance | [optional] [readonly] 
**AppName** | Pointer to **string** | Key name of the app, for example, &#x60;active_directory&#x60; | [optional] [readonly] 
**Assignments** | Pointer to **int32** | Number of app assignments in the org used as the denominator for &#x60;unassignmentPercentage&#x60;. Only returned for org-level roadblocks. | [optional] [readonly] 
**Created** | **time.Time** | Timestamp when the roadblock was created | [readonly] 
**EntitlementChanges** | Pointer to [**ImportRoadblockEntitlementChanges**](ImportRoadblockEntitlementChanges.md) |  | [optional] 
**EntitlementSettings** | Pointer to [**ImportRoadblockEntitlementSettings**](ImportRoadblockEntitlementSettings.md) |  | [optional] 
**Id** | **string** | Unique identifier for the roadblock | [readonly] 
**PauseScope** | **string** | Indicates which user imports the roadblock pauses:  * &#x60;ORG&#x60;: Every user import in the org is paused. * &#x60;APP_INSTANCE&#x60;: Only the user imports of this app instance are paused. | [readonly] 
**Reason** | Pointer to **string** | A human-readable message that describes why the roadblock was created, in the org&#39;s language. Use it for display only. Don&#39;t parse it; use &#x60;type&#x60; and the count fields instead. | [optional] [readonly] 
**Status** | **string** | Status of the roadblock | [readonly] 
**Type** | **string** | The safeguard that triggered the roadblock | [readonly] 
**UnassignmentPercentage** | Pointer to **float32** | The share of app assignments in the org that the import would remove (&#x60;unassignments&#x60; / &#x60;assignments&#x60; × 100). Only returned for org-level roadblocks. | [optional] [readonly] 
**Unassignments** | Pointer to **int32** | Number of app assignments in the org that the import would remove. Only returned for org-level roadblocks (&#x60;pauseScope&#x60; is &#x60;ORG&#x60; and &#x60;type&#x60; is &#x60;USER_DELETE&#x60;, &#x60;GROUP_DELETE&#x60;, or &#x60;GROUP_UPDATE&#x60;). | [optional] [readonly] 

## Methods

### NewImportRoadblock

`func NewImportRoadblock(appInstanceId string, created time.Time, id string, pauseScope string, status string, type_ string, ) *ImportRoadblock`

NewImportRoadblock instantiates a new ImportRoadblock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportRoadblockWithDefaults

`func NewImportRoadblockWithDefaults() *ImportRoadblock`

NewImportRoadblockWithDefaults instantiates a new ImportRoadblock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedApps

`func (o *ImportRoadblock) GetAffectedApps() []ImportRoadblockAffectedApp`

GetAffectedApps returns the AffectedApps field if non-nil, zero value otherwise.

### GetAffectedAppsOk

`func (o *ImportRoadblock) GetAffectedAppsOk() (*[]ImportRoadblockAffectedApp, bool)`

GetAffectedAppsOk returns a tuple with the AffectedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedApps

`func (o *ImportRoadblock) SetAffectedApps(v []ImportRoadblockAffectedApp)`

SetAffectedApps sets AffectedApps field to given value.

### HasAffectedApps

`func (o *ImportRoadblock) HasAffectedApps() bool`

HasAffectedApps returns a boolean if a field has been set.

### GetAppInstanceId

`func (o *ImportRoadblock) GetAppInstanceId() string`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *ImportRoadblock) GetAppInstanceIdOk() (*string, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *ImportRoadblock) SetAppInstanceId(v string)`

SetAppInstanceId sets AppInstanceId field to given value.


### GetAppInstanceLabel

`func (o *ImportRoadblock) GetAppInstanceLabel() string`

GetAppInstanceLabel returns the AppInstanceLabel field if non-nil, zero value otherwise.

### GetAppInstanceLabelOk

`func (o *ImportRoadblock) GetAppInstanceLabelOk() (*string, bool)`

GetAppInstanceLabelOk returns a tuple with the AppInstanceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceLabel

`func (o *ImportRoadblock) SetAppInstanceLabel(v string)`

SetAppInstanceLabel sets AppInstanceLabel field to given value.

### HasAppInstanceLabel

`func (o *ImportRoadblock) HasAppInstanceLabel() bool`

HasAppInstanceLabel returns a boolean if a field has been set.

### GetAppName

`func (o *ImportRoadblock) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ImportRoadblock) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ImportRoadblock) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *ImportRoadblock) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAssignments

`func (o *ImportRoadblock) GetAssignments() int32`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *ImportRoadblock) GetAssignmentsOk() (*int32, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *ImportRoadblock) SetAssignments(v int32)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *ImportRoadblock) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetCreated

`func (o *ImportRoadblock) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ImportRoadblock) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ImportRoadblock) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetEntitlementChanges

`func (o *ImportRoadblock) GetEntitlementChanges() ImportRoadblockEntitlementChanges`

GetEntitlementChanges returns the EntitlementChanges field if non-nil, zero value otherwise.

### GetEntitlementChangesOk

`func (o *ImportRoadblock) GetEntitlementChangesOk() (*ImportRoadblockEntitlementChanges, bool)`

GetEntitlementChangesOk returns a tuple with the EntitlementChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementChanges

`func (o *ImportRoadblock) SetEntitlementChanges(v ImportRoadblockEntitlementChanges)`

SetEntitlementChanges sets EntitlementChanges field to given value.

### HasEntitlementChanges

`func (o *ImportRoadblock) HasEntitlementChanges() bool`

HasEntitlementChanges returns a boolean if a field has been set.

### GetEntitlementSettings

`func (o *ImportRoadblock) GetEntitlementSettings() ImportRoadblockEntitlementSettings`

GetEntitlementSettings returns the EntitlementSettings field if non-nil, zero value otherwise.

### GetEntitlementSettingsOk

`func (o *ImportRoadblock) GetEntitlementSettingsOk() (*ImportRoadblockEntitlementSettings, bool)`

GetEntitlementSettingsOk returns a tuple with the EntitlementSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementSettings

`func (o *ImportRoadblock) SetEntitlementSettings(v ImportRoadblockEntitlementSettings)`

SetEntitlementSettings sets EntitlementSettings field to given value.

### HasEntitlementSettings

`func (o *ImportRoadblock) HasEntitlementSettings() bool`

HasEntitlementSettings returns a boolean if a field has been set.

### GetId

`func (o *ImportRoadblock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportRoadblock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportRoadblock) SetId(v string)`

SetId sets Id field to given value.


### GetPauseScope

`func (o *ImportRoadblock) GetPauseScope() string`

GetPauseScope returns the PauseScope field if non-nil, zero value otherwise.

### GetPauseScopeOk

`func (o *ImportRoadblock) GetPauseScopeOk() (*string, bool)`

GetPauseScopeOk returns a tuple with the PauseScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseScope

`func (o *ImportRoadblock) SetPauseScope(v string)`

SetPauseScope sets PauseScope field to given value.


### GetReason

`func (o *ImportRoadblock) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ImportRoadblock) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ImportRoadblock) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ImportRoadblock) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *ImportRoadblock) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImportRoadblock) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImportRoadblock) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *ImportRoadblock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImportRoadblock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImportRoadblock) SetType(v string)`

SetType sets Type field to given value.


### GetUnassignmentPercentage

`func (o *ImportRoadblock) GetUnassignmentPercentage() float32`

GetUnassignmentPercentage returns the UnassignmentPercentage field if non-nil, zero value otherwise.

### GetUnassignmentPercentageOk

`func (o *ImportRoadblock) GetUnassignmentPercentageOk() (*float32, bool)`

GetUnassignmentPercentageOk returns a tuple with the UnassignmentPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassignmentPercentage

`func (o *ImportRoadblock) SetUnassignmentPercentage(v float32)`

SetUnassignmentPercentage sets UnassignmentPercentage field to given value.

### HasUnassignmentPercentage

`func (o *ImportRoadblock) HasUnassignmentPercentage() bool`

HasUnassignmentPercentage returns a boolean if a field has been set.

### GetUnassignments

`func (o *ImportRoadblock) GetUnassignments() int32`

GetUnassignments returns the Unassignments field if non-nil, zero value otherwise.

### GetUnassignmentsOk

`func (o *ImportRoadblock) GetUnassignmentsOk() (*int32, bool)`

GetUnassignmentsOk returns a tuple with the Unassignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassignments

`func (o *ImportRoadblock) SetUnassignments(v int32)`

SetUnassignments sets Unassignments field to given value.

### HasUnassignments

`func (o *ImportRoadblock) HasUnassignments() bool`

HasUnassignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


