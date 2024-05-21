# OrgContactTypeObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactType** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewOrgContactTypeObj

`func NewOrgContactTypeObj() *OrgContactTypeObj`

NewOrgContactTypeObj instantiates a new OrgContactTypeObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgContactTypeObjWithDefaults

`func NewOrgContactTypeObjWithDefaults() *OrgContactTypeObj`

NewOrgContactTypeObjWithDefaults instantiates a new OrgContactTypeObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactType

`func (o *OrgContactTypeObj) GetContactType() string`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *OrgContactTypeObj) GetContactTypeOk() (*string, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *OrgContactTypeObj) SetContactType(v string)`

SetContactType sets ContactType field to given value.

### HasContactType

`func (o *OrgContactTypeObj) HasContactType() bool`

HasContactType returns a boolean if a field has been set.

### GetLinks

`func (o *OrgContactTypeObj) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrgContactTypeObj) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrgContactTypeObj) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrgContactTypeObj) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


