# DeprovisionTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppInstanceId** | **string** | ID of the app instance the task belongs to | 
**AppUserId** | **string** | ID of the user&#39;s assignment to the app instance | 
**CreatedDate** | **time.Time** | Timestamp when the task was created | [readonly] 
**EligibleActions** | **[]string** | Actions that you can currently perform on the task. Empty when none are available, which can also be the case for a &#x60;FAILED&#x60; task. Use this list instead of working out the actions from &#x60;outcome&#x60; or &#x60;reason&#x60;. | 
**Error** | **NullableString** | Human-readable error detail from the task. &#x60;null&#x60; when the task has no error text. | 
**Id** | **string** | Unique identifier of the task | [readonly] 
**LastUpdatedDate** | **time.Time** | Timestamp when the task was last updated | [readonly] 
**Outcome** | **string** | Outcome of the task | 
**Purpose** | **string** | Why the deprovision task was created | 
**Reason** | Pointer to **string** | Reason for the outcome. Not returned when the task has no reason. | [optional] 
**UserId** | **string** | ID of the user the task is for | 

## Methods

### NewDeprovisionTask

`func NewDeprovisionTask(appInstanceId string, appUserId string, createdDate time.Time, eligibleActions []string, error_ NullableString, id string, lastUpdatedDate time.Time, outcome string, purpose string, userId string, ) *DeprovisionTask`

NewDeprovisionTask instantiates a new DeprovisionTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprovisionTaskWithDefaults

`func NewDeprovisionTaskWithDefaults() *DeprovisionTask`

NewDeprovisionTaskWithDefaults instantiates a new DeprovisionTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppInstanceId

`func (o *DeprovisionTask) GetAppInstanceId() string`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *DeprovisionTask) GetAppInstanceIdOk() (*string, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *DeprovisionTask) SetAppInstanceId(v string)`

SetAppInstanceId sets AppInstanceId field to given value.


### GetAppUserId

`func (o *DeprovisionTask) GetAppUserId() string`

GetAppUserId returns the AppUserId field if non-nil, zero value otherwise.

### GetAppUserIdOk

`func (o *DeprovisionTask) GetAppUserIdOk() (*string, bool)`

GetAppUserIdOk returns a tuple with the AppUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUserId

`func (o *DeprovisionTask) SetAppUserId(v string)`

SetAppUserId sets AppUserId field to given value.


### GetCreatedDate

`func (o *DeprovisionTask) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *DeprovisionTask) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *DeprovisionTask) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetEligibleActions

`func (o *DeprovisionTask) GetEligibleActions() []string`

GetEligibleActions returns the EligibleActions field if non-nil, zero value otherwise.

### GetEligibleActionsOk

`func (o *DeprovisionTask) GetEligibleActionsOk() (*[]string, bool)`

GetEligibleActionsOk returns a tuple with the EligibleActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleActions

`func (o *DeprovisionTask) SetEligibleActions(v []string)`

SetEligibleActions sets EligibleActions field to given value.


### GetError

`func (o *DeprovisionTask) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DeprovisionTask) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DeprovisionTask) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *DeprovisionTask) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *DeprovisionTask) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetId

`func (o *DeprovisionTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeprovisionTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeprovisionTask) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedDate

`func (o *DeprovisionTask) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *DeprovisionTask) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *DeprovisionTask) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.


### GetOutcome

`func (o *DeprovisionTask) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *DeprovisionTask) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *DeprovisionTask) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetPurpose

`func (o *DeprovisionTask) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *DeprovisionTask) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *DeprovisionTask) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetReason

`func (o *DeprovisionTask) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DeprovisionTask) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DeprovisionTask) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DeprovisionTask) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetUserId

`func (o *DeprovisionTask) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DeprovisionTask) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DeprovisionTask) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


