# GroupQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | **[]string** | An array of LDAP group attribute names to retrieve. Restricted attributes: member, memberOf, * | 

## Methods

### NewGroupQueryRequest

`func NewGroupQueryRequest(attributes []string, ) *GroupQueryRequest`

NewGroupQueryRequest instantiates a new GroupQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupQueryRequestWithDefaults

`func NewGroupQueryRequestWithDefaults() *GroupQueryRequest`

NewGroupQueryRequestWithDefaults instantiates a new GroupQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *GroupQueryRequest) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GroupQueryRequest) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GroupQueryRequest) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


