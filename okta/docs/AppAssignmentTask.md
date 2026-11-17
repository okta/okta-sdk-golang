# AppAssignmentTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppInstanceId** | **string** | ID of the app instance the task belongs to | 
**AppUserId** | **string** | ID of the user&#39;s assignment to the app instance | 
**AssignmentType** | **string** | Type of assignment | 
**CreatedDate** | **time.Time** | Timestamp when the task was created | [readonly] 
**EligibleActions** | **[]string** | Actions that you can currently perform on the task. Empty when none are available, which can also be the case for a &#x60;FAILED&#x60; task. Use this list instead of working out the actions from &#x60;outcome&#x60; or &#x60;reason&#x60;. | 
**Error** | **NullableString** | Human-readable error detail from the task. &#x60;null&#x60; when the task has no error text. | 
**Id** | **string** | Unique identifier of the task | [readonly] 
**LastUpdatedDate** | **time.Time** | Timestamp when the task was last updated | [readonly] 
**Outcome** | **string** | Outcome of the task | 
**Reason** | Pointer to **string** | Reason for the outcome. Not returned when the task has no reason. | [optional] 
**UserId** | **string** | ID of the user the task is for | 

## Methods

### NewAppAssignmentTask

`func NewAppAssignmentTask(appInstanceId string, appUserId string, assignmentType string, createdDate time.Time, eligibleActions []string, error_ NullableString, id string, lastUpdatedDate time.Time, outcome string, userId string, ) *AppAssignmentTask`

NewAppAssignmentTask instantiates a new AppAssignmentTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAssignmentTaskWithDefaults

`func NewAppAssignmentTaskWithDefaults() *AppAssignmentTask`

NewAppAssignmentTaskWithDefaults instantiates a new AppAssignmentTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppInstanceId

`func (o *AppAssignmentTask) GetAppInstanceId() string`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *AppAssignmentTask) GetAppInstanceIdOk() (*string, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *AppAssignmentTask) SetAppInstanceId(v string)`

SetAppInstanceId sets AppInstanceId field to given value.


### GetAppUserId

`func (o *AppAssignmentTask) GetAppUserId() string`

GetAppUserId returns the AppUserId field if non-nil, zero value otherwise.

### GetAppUserIdOk

`func (o *AppAssignmentTask) GetAppUserIdOk() (*string, bool)`

GetAppUserIdOk returns a tuple with the AppUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUserId

`func (o *AppAssignmentTask) SetAppUserId(v string)`

SetAppUserId sets AppUserId field to given value.


### GetAssignmentType

`func (o *AppAssignmentTask) GetAssignmentType() string`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *AppAssignmentTask) GetAssignmentTypeOk() (*string, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *AppAssignmentTask) SetAssignmentType(v string)`

SetAssignmentType sets AssignmentType field to given value.


### GetCreatedDate

`func (o *AppAssignmentTask) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *AppAssignmentTask) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *AppAssignmentTask) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetEligibleActions

`func (o *AppAssignmentTask) GetEligibleActions() []string`

GetEligibleActions returns the EligibleActions field if non-nil, zero value otherwise.

### GetEligibleActionsOk

`func (o *AppAssignmentTask) GetEligibleActionsOk() (*[]string, bool)`

GetEligibleActionsOk returns a tuple with the EligibleActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleActions

`func (o *AppAssignmentTask) SetEligibleActions(v []string)`

SetEligibleActions sets EligibleActions field to given value.


### GetError

`func (o *AppAssignmentTask) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AppAssignmentTask) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AppAssignmentTask) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *AppAssignmentTask) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AppAssignmentTask) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetId

`func (o *AppAssignmentTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppAssignmentTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppAssignmentTask) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedDate

`func (o *AppAssignmentTask) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *AppAssignmentTask) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *AppAssignmentTask) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.


### GetOutcome

`func (o *AppAssignmentTask) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AppAssignmentTask) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AppAssignmentTask) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetReason

`func (o *AppAssignmentTask) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AppAssignmentTask) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AppAssignmentTask) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AppAssignmentTask) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetUserId

`func (o *AppAssignmentTask) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AppAssignmentTask) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AppAssignmentTask) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


