# SyncPasswordTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppInstanceId** | **string** | ID of the app instance the task belongs to | 
**AppUserId** | **string** | ID of the user&#39;s assignment to the app instance | 
**CreatedDate** | **time.Time** | Timestamp when the task was created | [readonly] 
**EligibleActions** | **[]string** | Actions that you can currently perform on the task. Empty when none are available. Use this list instead of working out the actions from &#x60;outcome&#x60;. | 
**Id** | **string** | Unique identifier of the task | [readonly] 
**LastUpdatedDate** | **time.Time** | Timestamp when the task was last updated | [readonly] 
**Outcome** | **string** | Outcome of the task | 
**UserId** | **string** | ID of the user the task is for | 

## Methods

### NewSyncPasswordTask

`func NewSyncPasswordTask(appInstanceId string, appUserId string, createdDate time.Time, eligibleActions []string, id string, lastUpdatedDate time.Time, outcome string, userId string, ) *SyncPasswordTask`

NewSyncPasswordTask instantiates a new SyncPasswordTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncPasswordTaskWithDefaults

`func NewSyncPasswordTaskWithDefaults() *SyncPasswordTask`

NewSyncPasswordTaskWithDefaults instantiates a new SyncPasswordTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppInstanceId

`func (o *SyncPasswordTask) GetAppInstanceId() string`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *SyncPasswordTask) GetAppInstanceIdOk() (*string, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *SyncPasswordTask) SetAppInstanceId(v string)`

SetAppInstanceId sets AppInstanceId field to given value.


### GetAppUserId

`func (o *SyncPasswordTask) GetAppUserId() string`

GetAppUserId returns the AppUserId field if non-nil, zero value otherwise.

### GetAppUserIdOk

`func (o *SyncPasswordTask) GetAppUserIdOk() (*string, bool)`

GetAppUserIdOk returns a tuple with the AppUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUserId

`func (o *SyncPasswordTask) SetAppUserId(v string)`

SetAppUserId sets AppUserId field to given value.


### GetCreatedDate

`func (o *SyncPasswordTask) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *SyncPasswordTask) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *SyncPasswordTask) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetEligibleActions

`func (o *SyncPasswordTask) GetEligibleActions() []string`

GetEligibleActions returns the EligibleActions field if non-nil, zero value otherwise.

### GetEligibleActionsOk

`func (o *SyncPasswordTask) GetEligibleActionsOk() (*[]string, bool)`

GetEligibleActionsOk returns a tuple with the EligibleActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleActions

`func (o *SyncPasswordTask) SetEligibleActions(v []string)`

SetEligibleActions sets EligibleActions field to given value.


### GetId

`func (o *SyncPasswordTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncPasswordTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncPasswordTask) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedDate

`func (o *SyncPasswordTask) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *SyncPasswordTask) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *SyncPasswordTask) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.


### GetOutcome

`func (o *SyncPasswordTask) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *SyncPasswordTask) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *SyncPasswordTask) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetUserId

`func (o *SyncPasswordTask) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SyncPasswordTask) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SyncPasswordTask) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


