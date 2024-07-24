# NetworkZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique identifier for the Network Zone | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the object was last modified | [optional] [readonly] 
**Name** | **string** | Unique name for this Network Zone | 
**Status** | Pointer to **string** | Network Zone status | [optional] 
**System** | Pointer to **bool** | Indicates a system Network Zone: * &#x60;true&#x60; for system Network Zones * &#x60;false&#x60; for custom Network Zones  The Okta org provides the following default system Network Zones: * &#x60;LegacyIpZone&#x60; * &#x60;BlockedIpZone&#x60; * &lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; &#x60;DefaultEnhancedDynamicZone&#x60;  Admins can modify the name of the default system Network Zone and add up to 5000 gateway or proxy IP entries.  | [optional] [readonly] 
**Type** | **string** | The type of Network Zone | 
**Usage** | Pointer to **string** | The usage of the Network Zone | [optional] 
**Links** | Pointer to [**LinksSelfAndLifecycle**](LinksSelfAndLifecycle.md) |  | [optional] 

## Methods

### NewNetworkZone

`func NewNetworkZone(name string, type_ string, ) *NetworkZone`

NewNetworkZone instantiates a new NetworkZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkZoneWithDefaults

`func NewNetworkZoneWithDefaults() *NetworkZone`

NewNetworkZoneWithDefaults instantiates a new NetworkZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *NetworkZone) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *NetworkZone) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *NetworkZone) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *NetworkZone) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *NetworkZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NetworkZone) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NetworkZone) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NetworkZone) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NetworkZone) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *NetworkZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkZone) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *NetworkZone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkZone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkZone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkZone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *NetworkZone) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *NetworkZone) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *NetworkZone) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *NetworkZone) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetType

`func (o *NetworkZone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkZone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkZone) SetType(v string)`

SetType sets Type field to given value.


### GetUsage

`func (o *NetworkZone) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *NetworkZone) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *NetworkZone) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *NetworkZone) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetLinks

`func (o *NetworkZone) GetLinks() LinksSelfAndLifecycle`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkZone) GetLinksOk() (*LinksSelfAndLifecycle, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkZone) SetLinks(v LinksSelfAndLifecycle)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkZone) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


