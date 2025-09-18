# OrgBillingContactType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactType** | Pointer to **string** | Type of contact | [optional] 
**Links** | Pointer to [**OrgBillingContactTypeLinks**](OrgBillingContactTypeLinks.md) |  | [optional] 

## Methods

### NewOrgBillingContactType

`func NewOrgBillingContactType() *OrgBillingContactType`

NewOrgBillingContactType instantiates a new OrgBillingContactType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgBillingContactTypeWithDefaults

`func NewOrgBillingContactTypeWithDefaults() *OrgBillingContactType`

NewOrgBillingContactTypeWithDefaults instantiates a new OrgBillingContactType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactType

`func (o *OrgBillingContactType) GetContactType() string`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *OrgBillingContactType) GetContactTypeOk() (*string, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *OrgBillingContactType) SetContactType(v string)`

SetContactType sets ContactType field to given value.

### HasContactType

`func (o *OrgBillingContactType) HasContactType() bool`

HasContactType returns a boolean if a field has been set.

### GetLinks

`func (o *OrgBillingContactType) GetLinks() OrgBillingContactTypeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrgBillingContactType) GetLinksOk() (*OrgBillingContactTypeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrgBillingContactType) SetLinks(v OrgBillingContactTypeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrgBillingContactType) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


