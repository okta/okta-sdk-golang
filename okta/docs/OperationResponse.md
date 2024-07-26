# OperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentOperation** | Pointer to [**OperationResponseAssignmentOperation**](OperationResponseAssignmentOperation.md) |  | [optional] 
**Completed** | Pointer to **time.Time** |  | [optional] [readonly] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**NumUserMoved** | Pointer to **float32** |  | [optional] [readonly] 
**RealmId** | Pointer to **string** |  | [optional] [readonly] 
**RealmName** | Pointer to **string** |  | [optional] [readonly] 
**Started** | Pointer to **time.Time** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewOperationResponse

`func NewOperationResponse() *OperationResponse`

NewOperationResponse instantiates a new OperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseWithDefaults

`func NewOperationResponseWithDefaults() *OperationResponse`

NewOperationResponseWithDefaults instantiates a new OperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentOperation

`func (o *OperationResponse) GetAssignmentOperation() OperationResponseAssignmentOperation`

GetAssignmentOperation returns the AssignmentOperation field if non-nil, zero value otherwise.

### GetAssignmentOperationOk

`func (o *OperationResponse) GetAssignmentOperationOk() (*OperationResponseAssignmentOperation, bool)`

GetAssignmentOperationOk returns a tuple with the AssignmentOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentOperation

`func (o *OperationResponse) SetAssignmentOperation(v OperationResponseAssignmentOperation)`

SetAssignmentOperation sets AssignmentOperation field to given value.

### HasAssignmentOperation

`func (o *OperationResponse) HasAssignmentOperation() bool`

HasAssignmentOperation returns a boolean if a field has been set.

### GetCompleted

`func (o *OperationResponse) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *OperationResponse) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *OperationResponse) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *OperationResponse) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetCreated

`func (o *OperationResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OperationResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OperationResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OperationResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *OperationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperationResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OperationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumUserMoved

`func (o *OperationResponse) GetNumUserMoved() float32`

GetNumUserMoved returns the NumUserMoved field if non-nil, zero value otherwise.

### GetNumUserMovedOk

`func (o *OperationResponse) GetNumUserMovedOk() (*float32, bool)`

GetNumUserMovedOk returns a tuple with the NumUserMoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUserMoved

`func (o *OperationResponse) SetNumUserMoved(v float32)`

SetNumUserMoved sets NumUserMoved field to given value.

### HasNumUserMoved

`func (o *OperationResponse) HasNumUserMoved() bool`

HasNumUserMoved returns a boolean if a field has been set.

### GetRealmId

`func (o *OperationResponse) GetRealmId() string`

GetRealmId returns the RealmId field if non-nil, zero value otherwise.

### GetRealmIdOk

`func (o *OperationResponse) GetRealmIdOk() (*string, bool)`

GetRealmIdOk returns a tuple with the RealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmId

`func (o *OperationResponse) SetRealmId(v string)`

SetRealmId sets RealmId field to given value.

### HasRealmId

`func (o *OperationResponse) HasRealmId() bool`

HasRealmId returns a boolean if a field has been set.

### GetRealmName

`func (o *OperationResponse) GetRealmName() string`

GetRealmName returns the RealmName field if non-nil, zero value otherwise.

### GetRealmNameOk

`func (o *OperationResponse) GetRealmNameOk() (*string, bool)`

GetRealmNameOk returns a tuple with the RealmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmName

`func (o *OperationResponse) SetRealmName(v string)`

SetRealmName sets RealmName field to given value.

### HasRealmName

`func (o *OperationResponse) HasRealmName() bool`

HasRealmName returns a boolean if a field has been set.

### GetStarted

`func (o *OperationResponse) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *OperationResponse) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *OperationResponse) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *OperationResponse) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStatus

`func (o *OperationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OperationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OperationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OperationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *OperationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OperationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OperationResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OperationResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *OperationResponse) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OperationResponse) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OperationResponse) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OperationResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


