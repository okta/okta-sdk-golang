# ListNetworkZones200ResponseInner

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
**Asns** | Pointer to [**EnhancedDynamicNetworkZoneAllOfAsns**](EnhancedDynamicNetworkZoneAllOfAsns.md) |  | [optional] 
**ProxyType** | Pointer to **string** | The proxy type used for a Dynamic Network Zone | [optional] 
**Locations** | Pointer to [**EnhancedDynamicNetworkZoneAllOfLocations**](EnhancedDynamicNetworkZoneAllOfLocations.md) |  | [optional] 
**IpServiceCategories** | Pointer to [**EnhancedDynamicNetworkZoneAllOfIpServiceCategories**](EnhancedDynamicNetworkZoneAllOfIpServiceCategories.md) |  | [optional] 

## Methods

### NewListNetworkZones200ResponseInner

`func NewListNetworkZones200ResponseInner(name string, type_ string, ) *ListNetworkZones200ResponseInner`

NewListNetworkZones200ResponseInner instantiates a new ListNetworkZones200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkZones200ResponseInnerWithDefaults

`func NewListNetworkZones200ResponseInnerWithDefaults() *ListNetworkZones200ResponseInner`

NewListNetworkZones200ResponseInnerWithDefaults instantiates a new ListNetworkZones200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ListNetworkZones200ResponseInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListNetworkZones200ResponseInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListNetworkZones200ResponseInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListNetworkZones200ResponseInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *ListNetworkZones200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListNetworkZones200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListNetworkZones200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListNetworkZones200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListNetworkZones200ResponseInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListNetworkZones200ResponseInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListNetworkZones200ResponseInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListNetworkZones200ResponseInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *ListNetworkZones200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListNetworkZones200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListNetworkZones200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ListNetworkZones200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListNetworkZones200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListNetworkZones200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListNetworkZones200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *ListNetworkZones200ResponseInner) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ListNetworkZones200ResponseInner) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ListNetworkZones200ResponseInner) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *ListNetworkZones200ResponseInner) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetType

`func (o *ListNetworkZones200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListNetworkZones200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListNetworkZones200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetUsage

`func (o *ListNetworkZones200ResponseInner) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ListNetworkZones200ResponseInner) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ListNetworkZones200ResponseInner) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ListNetworkZones200ResponseInner) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetLinks

`func (o *ListNetworkZones200ResponseInner) GetLinks() LinksSelfAndLifecycle`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListNetworkZones200ResponseInner) GetLinksOk() (*LinksSelfAndLifecycle, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListNetworkZones200ResponseInner) SetLinks(v LinksSelfAndLifecycle)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListNetworkZones200ResponseInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAsns

`func (o *ListNetworkZones200ResponseInner) GetAsns() EnhancedDynamicNetworkZoneAllOfAsns`

GetAsns returns the Asns field if non-nil, zero value otherwise.

### GetAsnsOk

`func (o *ListNetworkZones200ResponseInner) GetAsnsOk() (*EnhancedDynamicNetworkZoneAllOfAsns, bool)`

GetAsnsOk returns a tuple with the Asns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsns

`func (o *ListNetworkZones200ResponseInner) SetAsns(v EnhancedDynamicNetworkZoneAllOfAsns)`

SetAsns sets Asns field to given value.

### HasAsns

`func (o *ListNetworkZones200ResponseInner) HasAsns() bool`

HasAsns returns a boolean if a field has been set.

### GetProxyType

`func (o *ListNetworkZones200ResponseInner) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *ListNetworkZones200ResponseInner) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *ListNetworkZones200ResponseInner) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *ListNetworkZones200ResponseInner) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetLocations

`func (o *ListNetworkZones200ResponseInner) GetLocations() EnhancedDynamicNetworkZoneAllOfLocations`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ListNetworkZones200ResponseInner) GetLocationsOk() (*EnhancedDynamicNetworkZoneAllOfLocations, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ListNetworkZones200ResponseInner) SetLocations(v EnhancedDynamicNetworkZoneAllOfLocations)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ListNetworkZones200ResponseInner) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetIpServiceCategories

`func (o *ListNetworkZones200ResponseInner) GetIpServiceCategories() EnhancedDynamicNetworkZoneAllOfIpServiceCategories`

GetIpServiceCategories returns the IpServiceCategories field if non-nil, zero value otherwise.

### GetIpServiceCategoriesOk

`func (o *ListNetworkZones200ResponseInner) GetIpServiceCategoriesOk() (*EnhancedDynamicNetworkZoneAllOfIpServiceCategories, bool)`

GetIpServiceCategoriesOk returns a tuple with the IpServiceCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpServiceCategories

`func (o *ListNetworkZones200ResponseInner) SetIpServiceCategories(v EnhancedDynamicNetworkZoneAllOfIpServiceCategories)`

SetIpServiceCategories sets IpServiceCategories field to given value.

### HasIpServiceCategories

`func (o *ListNetworkZones200ResponseInner) HasIpServiceCategories() bool`

HasIpServiceCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


