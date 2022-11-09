# ResourceSetBindingMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the role was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique key for the role | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the role was last updated | [optional] [readonly] 
**Links** | Pointer to [**ApiTokenLink**](ApiTokenLink.md) |  | [optional] 

## Methods

### NewResourceSetBindingMember

`func NewResourceSetBindingMember() *ResourceSetBindingMember`

NewResourceSetBindingMember instantiates a new ResourceSetBindingMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSetBindingMemberWithDefaults

`func NewResourceSetBindingMemberWithDefaults() *ResourceSetBindingMember`

NewResourceSetBindingMemberWithDefaults instantiates a new ResourceSetBindingMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ResourceSetBindingMember) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResourceSetBindingMember) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResourceSetBindingMember) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResourceSetBindingMember) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *ResourceSetBindingMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceSetBindingMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceSetBindingMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceSetBindingMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ResourceSetBindingMember) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ResourceSetBindingMember) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ResourceSetBindingMember) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ResourceSetBindingMember) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceSetBindingMember) GetLinks() ApiTokenLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceSetBindingMember) GetLinksOk() (*ApiTokenLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceSetBindingMember) SetLinks(v ApiTokenLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceSetBindingMember) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


