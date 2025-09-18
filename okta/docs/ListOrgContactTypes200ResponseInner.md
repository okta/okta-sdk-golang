# ListOrgContactTypes200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactType** | Pointer to **string** | Type of contact | [optional] 
**Links** | Pointer to [**OrgTechnicalContactTypeLinks**](OrgTechnicalContactTypeLinks.md) |  | [optional] 

## Methods

### NewListOrgContactTypes200ResponseInner

`func NewListOrgContactTypes200ResponseInner() *ListOrgContactTypes200ResponseInner`

NewListOrgContactTypes200ResponseInner instantiates a new ListOrgContactTypes200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrgContactTypes200ResponseInnerWithDefaults

`func NewListOrgContactTypes200ResponseInnerWithDefaults() *ListOrgContactTypes200ResponseInner`

NewListOrgContactTypes200ResponseInnerWithDefaults instantiates a new ListOrgContactTypes200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactType

`func (o *ListOrgContactTypes200ResponseInner) GetContactType() string`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *ListOrgContactTypes200ResponseInner) GetContactTypeOk() (*string, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *ListOrgContactTypes200ResponseInner) SetContactType(v string)`

SetContactType sets ContactType field to given value.

### HasContactType

`func (o *ListOrgContactTypes200ResponseInner) HasContactType() bool`

HasContactType returns a boolean if a field has been set.

### GetLinks

`func (o *ListOrgContactTypes200ResponseInner) GetLinks() OrgTechnicalContactTypeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListOrgContactTypes200ResponseInner) GetLinksOk() (*OrgTechnicalContactTypeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListOrgContactTypes200ResponseInner) SetLinks(v OrgTechnicalContactTypeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListOrgContactTypes200ResponseInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


