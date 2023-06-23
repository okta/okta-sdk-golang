# OrgPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowEndUserFooter** | Pointer to **bool** |  | [optional] [readonly] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewOrgPreferences

`func NewOrgPreferences() *OrgPreferences`

NewOrgPreferences instantiates a new OrgPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgPreferencesWithDefaults

`func NewOrgPreferencesWithDefaults() *OrgPreferences`

NewOrgPreferencesWithDefaults instantiates a new OrgPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowEndUserFooter

`func (o *OrgPreferences) GetShowEndUserFooter() bool`

GetShowEndUserFooter returns the ShowEndUserFooter field if non-nil, zero value otherwise.

### GetShowEndUserFooterOk

`func (o *OrgPreferences) GetShowEndUserFooterOk() (*bool, bool)`

GetShowEndUserFooterOk returns a tuple with the ShowEndUserFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowEndUserFooter

`func (o *OrgPreferences) SetShowEndUserFooter(v bool)`

SetShowEndUserFooter sets ShowEndUserFooter field to given value.

### HasShowEndUserFooter

`func (o *OrgPreferences) HasShowEndUserFooter() bool`

HasShowEndUserFooter returns a boolean if a field has been set.

### GetLinks

`func (o *OrgPreferences) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrgPreferences) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrgPreferences) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrgPreferences) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


