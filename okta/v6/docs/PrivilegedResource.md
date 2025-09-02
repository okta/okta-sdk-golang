# PrivilegedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**CredentialSyncInfo** | Pointer to [**CredentialSyncInfo**](CredentialSyncInfo.md) |  | [optional] 
**Id** | Pointer to **string** | ID of the privileged resource | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**ResourceType** | **string** | The type of the resource | 
**Status** | Pointer to **string** | Current status of the privileged resource | [optional] [readonly] 

## Methods

### NewPrivilegedResource

`func NewPrivilegedResource(resourceType string, ) *PrivilegedResource`

NewPrivilegedResource instantiates a new PrivilegedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegedResourceWithDefaults

`func NewPrivilegedResourceWithDefaults() *PrivilegedResource`

NewPrivilegedResourceWithDefaults instantiates a new PrivilegedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PrivilegedResource) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PrivilegedResource) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PrivilegedResource) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PrivilegedResource) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCredentialSyncInfo

`func (o *PrivilegedResource) GetCredentialSyncInfo() CredentialSyncInfo`

GetCredentialSyncInfo returns the CredentialSyncInfo field if non-nil, zero value otherwise.

### GetCredentialSyncInfoOk

`func (o *PrivilegedResource) GetCredentialSyncInfoOk() (*CredentialSyncInfo, bool)`

GetCredentialSyncInfoOk returns a tuple with the CredentialSyncInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialSyncInfo

`func (o *PrivilegedResource) SetCredentialSyncInfo(v CredentialSyncInfo)`

SetCredentialSyncInfo sets CredentialSyncInfo field to given value.

### HasCredentialSyncInfo

`func (o *PrivilegedResource) HasCredentialSyncInfo() bool`

HasCredentialSyncInfo returns a boolean if a field has been set.

### GetId

`func (o *PrivilegedResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivilegedResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivilegedResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrivilegedResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PrivilegedResource) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PrivilegedResource) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PrivilegedResource) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PrivilegedResource) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetResourceType

`func (o *PrivilegedResource) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *PrivilegedResource) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *PrivilegedResource) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetStatus

`func (o *PrivilegedResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivilegedResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivilegedResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrivilegedResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


