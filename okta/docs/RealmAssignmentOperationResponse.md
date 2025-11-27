# RealmAssignmentOperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **time.Time** | Timestamp of when the operation completed | [optional] 
**Created** | **time.Time** | Timestamp of when the operation was created | 
**Id** | **string** | ID of the asynchronous operation | 
**Started** | Pointer to **time.Time** | Timestamp of when the operation started | [optional] 
**Status** | **string** | The status of the asynchronous operation | 
**Type** | **string** | The operation type | 
**AssignmentOperation** | Pointer to [**RealmAssignmentOperationResponseAllOfAssignmentOperation**](RealmAssignmentOperationResponseAllOfAssignmentOperation.md) |  | [optional] 
**NumUserMoved** | Pointer to **float32** | Number of users moved | [optional] [readonly] 
**RealmId** | Pointer to **string** | ID of the realm | [optional] [readonly] 
**RealmName** | Pointer to **string** | Name of the realm | [optional] [readonly] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewRealmAssignmentOperationResponse

`func NewRealmAssignmentOperationResponse(created time.Time, id string, status string, type_ string, ) *RealmAssignmentOperationResponse`

NewRealmAssignmentOperationResponse instantiates a new RealmAssignmentOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmAssignmentOperationResponseWithDefaults

`func NewRealmAssignmentOperationResponseWithDefaults() *RealmAssignmentOperationResponse`

NewRealmAssignmentOperationResponseWithDefaults instantiates a new RealmAssignmentOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *RealmAssignmentOperationResponse) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *RealmAssignmentOperationResponse) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *RealmAssignmentOperationResponse) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *RealmAssignmentOperationResponse) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetCreated

`func (o *RealmAssignmentOperationResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RealmAssignmentOperationResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RealmAssignmentOperationResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *RealmAssignmentOperationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmAssignmentOperationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmAssignmentOperationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetStarted

`func (o *RealmAssignmentOperationResponse) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *RealmAssignmentOperationResponse) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *RealmAssignmentOperationResponse) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *RealmAssignmentOperationResponse) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStatus

`func (o *RealmAssignmentOperationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RealmAssignmentOperationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RealmAssignmentOperationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *RealmAssignmentOperationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmAssignmentOperationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmAssignmentOperationResponse) SetType(v string)`

SetType sets Type field to given value.


### GetAssignmentOperation

`func (o *RealmAssignmentOperationResponse) GetAssignmentOperation() RealmAssignmentOperationResponseAllOfAssignmentOperation`

GetAssignmentOperation returns the AssignmentOperation field if non-nil, zero value otherwise.

### GetAssignmentOperationOk

`func (o *RealmAssignmentOperationResponse) GetAssignmentOperationOk() (*RealmAssignmentOperationResponseAllOfAssignmentOperation, bool)`

GetAssignmentOperationOk returns a tuple with the AssignmentOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentOperation

`func (o *RealmAssignmentOperationResponse) SetAssignmentOperation(v RealmAssignmentOperationResponseAllOfAssignmentOperation)`

SetAssignmentOperation sets AssignmentOperation field to given value.

### HasAssignmentOperation

`func (o *RealmAssignmentOperationResponse) HasAssignmentOperation() bool`

HasAssignmentOperation returns a boolean if a field has been set.

### GetNumUserMoved

`func (o *RealmAssignmentOperationResponse) GetNumUserMoved() float32`

GetNumUserMoved returns the NumUserMoved field if non-nil, zero value otherwise.

### GetNumUserMovedOk

`func (o *RealmAssignmentOperationResponse) GetNumUserMovedOk() (*float32, bool)`

GetNumUserMovedOk returns a tuple with the NumUserMoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUserMoved

`func (o *RealmAssignmentOperationResponse) SetNumUserMoved(v float32)`

SetNumUserMoved sets NumUserMoved field to given value.

### HasNumUserMoved

`func (o *RealmAssignmentOperationResponse) HasNumUserMoved() bool`

HasNumUserMoved returns a boolean if a field has been set.

### GetRealmId

`func (o *RealmAssignmentOperationResponse) GetRealmId() string`

GetRealmId returns the RealmId field if non-nil, zero value otherwise.

### GetRealmIdOk

`func (o *RealmAssignmentOperationResponse) GetRealmIdOk() (*string, bool)`

GetRealmIdOk returns a tuple with the RealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmId

`func (o *RealmAssignmentOperationResponse) SetRealmId(v string)`

SetRealmId sets RealmId field to given value.

### HasRealmId

`func (o *RealmAssignmentOperationResponse) HasRealmId() bool`

HasRealmId returns a boolean if a field has been set.

### GetRealmName

`func (o *RealmAssignmentOperationResponse) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *RealmAssignmentOperationResponse) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *RealmAssignmentOperationResponse) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.

### HasRealmName

`func (o *RealmAssignmentOperationResponse) HasRealmName() bool`

HasRealmName returns a boolean if a field has been set.

### GetLinks

`func (o *RealmAssignmentOperationResponse) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RealmAssignmentOperationResponse) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RealmAssignmentOperationResponse) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RealmAssignmentOperationResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


