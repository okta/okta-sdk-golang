# PolicyMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**PolicyMappingLinks**](PolicyMappingLinks.md) |  | [optional] 

## Methods

### NewPolicyMapping

`func NewPolicyMapping() *PolicyMapping`

NewPolicyMapping instantiates a new PolicyMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyMappingWithDefaults

`func NewPolicyMappingWithDefaults() *PolicyMapping`

NewPolicyMappingWithDefaults instantiates a new PolicyMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyMapping) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyMapping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *PolicyMapping) GetLinks() PolicyMappingLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PolicyMapping) GetLinksOk() (*PolicyMappingLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PolicyMapping) SetLinks(v PolicyMappingLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PolicyMapping) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


