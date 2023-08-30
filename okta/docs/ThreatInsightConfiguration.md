# ThreatInsightConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Specifies how Okta responds to authentication requests from suspicious IP addresses | 
**Created** | Pointer to **time.Time** | Timestamp when the ThreatInsight Configuration object was created | [optional] [readonly] 
**ExcludeZones** | Pointer to **[]string** | Accepts a list of [Network Zone](/openapi/okta-management/management/tag/NetworkZone/) IDs. IPs in the excluded network zones aren&#39;t logged or blocked. This ensures that traffic from known, trusted IPs isn&#39;t accidentally logged or blocked. | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the ThreatInsight Configuration object was last updated | [optional] [readonly] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewThreatInsightConfiguration

`func NewThreatInsightConfiguration(action string, ) *ThreatInsightConfiguration`

NewThreatInsightConfiguration instantiates a new ThreatInsightConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatInsightConfigurationWithDefaults

`func NewThreatInsightConfigurationWithDefaults() *ThreatInsightConfiguration`

NewThreatInsightConfigurationWithDefaults instantiates a new ThreatInsightConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ThreatInsightConfiguration) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ThreatInsightConfiguration) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ThreatInsightConfiguration) SetAction(v string)`

SetAction sets Action field to given value.


### GetCreated

`func (o *ThreatInsightConfiguration) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ThreatInsightConfiguration) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ThreatInsightConfiguration) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ThreatInsightConfiguration) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExcludeZones

`func (o *ThreatInsightConfiguration) GetExcludeZones() []string`

GetExcludeZones returns the ExcludeZones field if non-nil, zero value otherwise.

### GetExcludeZonesOk

`func (o *ThreatInsightConfiguration) GetExcludeZonesOk() (*[]string, bool)`

GetExcludeZonesOk returns a tuple with the ExcludeZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeZones

`func (o *ThreatInsightConfiguration) SetExcludeZones(v []string)`

SetExcludeZones sets ExcludeZones field to given value.

### HasExcludeZones

`func (o *ThreatInsightConfiguration) HasExcludeZones() bool`

HasExcludeZones returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ThreatInsightConfiguration) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ThreatInsightConfiguration) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ThreatInsightConfiguration) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ThreatInsightConfiguration) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *ThreatInsightConfiguration) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ThreatInsightConfiguration) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ThreatInsightConfiguration) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ThreatInsightConfiguration) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


