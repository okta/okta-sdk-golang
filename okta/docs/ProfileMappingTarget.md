# ProfileMappingTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the application instance or UserType | [optional] [readonly] 
**Name** | Pointer to **string** | Variable name of the application instance or name of the referenced userType | [optional] [readonly] 
**Type** | Pointer to **string** | Type of user referenced in the mapping | [optional] [readonly] 
**Links** | Pointer to [**SourceLinks**](SourceLinks.md) |  | [optional] 

## Methods

### NewProfileMappingTarget

`func NewProfileMappingTarget() *ProfileMappingTarget`

NewProfileMappingTarget instantiates a new ProfileMappingTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileMappingTargetWithDefaults

`func NewProfileMappingTargetWithDefaults() *ProfileMappingTarget`

NewProfileMappingTargetWithDefaults instantiates a new ProfileMappingTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProfileMappingTarget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileMappingTarget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileMappingTarget) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileMappingTarget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProfileMappingTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProfileMappingTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProfileMappingTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProfileMappingTarget) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ProfileMappingTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProfileMappingTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProfileMappingTarget) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProfileMappingTarget) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *ProfileMappingTarget) GetLinks() SourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProfileMappingTarget) GetLinksOk() (*SourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProfileMappingTarget) SetLinks(v SourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProfileMappingTarget) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


