# OrgCrossAppAccessConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The ISO 8601 formatted date and time when the connection was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique identifier for the connection | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | The ISO 8601 formatted date and time when the connection was last updated | [optional] [readonly] 
**RequestingAppInstanceId** | Pointer to **string** | ID of the requesting app instance | [optional] 
**ResourceAppInstanceId** | Pointer to **string** | ID of the resource app instance | [optional] 
**Status** | Pointer to **string** | Indicates if the Cross App Access connection is active or inactive | [optional] 

## Methods

### NewOrgCrossAppAccessConnection

`func NewOrgCrossAppAccessConnection() *OrgCrossAppAccessConnection`

NewOrgCrossAppAccessConnection instantiates a new OrgCrossAppAccessConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgCrossAppAccessConnectionWithDefaults

`func NewOrgCrossAppAccessConnectionWithDefaults() *OrgCrossAppAccessConnection`

NewOrgCrossAppAccessConnectionWithDefaults instantiates a new OrgCrossAppAccessConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *OrgCrossAppAccessConnection) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OrgCrossAppAccessConnection) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OrgCrossAppAccessConnection) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OrgCrossAppAccessConnection) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *OrgCrossAppAccessConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgCrossAppAccessConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgCrossAppAccessConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrgCrossAppAccessConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *OrgCrossAppAccessConnection) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *OrgCrossAppAccessConnection) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *OrgCrossAppAccessConnection) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *OrgCrossAppAccessConnection) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRequestingAppInstanceId

`func (o *OrgCrossAppAccessConnection) GetRequestingAppInstanceId() string`

GetRequestingAppInstanceId returns the RequestingAppInstanceId field if non-nil, zero value otherwise.

### GetRequestingAppInstanceIdOk

`func (o *OrgCrossAppAccessConnection) GetRequestingAppInstanceIdOk() (*string, bool)`

GetRequestingAppInstanceIdOk returns a tuple with the RequestingAppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestingAppInstanceId

`func (o *OrgCrossAppAccessConnection) SetRequestingAppInstanceId(v string)`

SetRequestingAppInstanceId sets RequestingAppInstanceId field to given value.

### HasRequestingAppInstanceId

`func (o *OrgCrossAppAccessConnection) HasRequestingAppInstanceId() bool`

HasRequestingAppInstanceId returns a boolean if a field has been set.

### GetResourceAppInstanceId

`func (o *OrgCrossAppAccessConnection) GetResourceAppInstanceId() string`

GetResourceAppInstanceId returns the ResourceAppInstanceId field if non-nil, zero value otherwise.

### GetResourceAppInstanceIdOk

`func (o *OrgCrossAppAccessConnection) GetResourceAppInstanceIdOk() (*string, bool)`

GetResourceAppInstanceIdOk returns a tuple with the ResourceAppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAppInstanceId

`func (o *OrgCrossAppAccessConnection) SetResourceAppInstanceId(v string)`

SetResourceAppInstanceId sets ResourceAppInstanceId field to given value.

### HasResourceAppInstanceId

`func (o *OrgCrossAppAccessConnection) HasResourceAppInstanceId() bool`

HasResourceAppInstanceId returns a boolean if a field has been set.

### GetStatus

`func (o *OrgCrossAppAccessConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrgCrossAppAccessConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrgCrossAppAccessConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrgCrossAppAccessConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


