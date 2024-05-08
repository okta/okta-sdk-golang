# OSVersionConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DynamicVersionRequirement** | Pointer to [**OSVersionConstraintDynamicVersionRequirement**](OSVersionConstraintDynamicVersionRequirement.md) |  | [optional] 
**MajorVersionConstraint** | **string** | Indicates the Windows major version | 
**Minimum** | Pointer to **string** | The Windows device version must be equal to or newer than the specified version | [optional] 

## Methods

### NewOSVersionConstraint

`func NewOSVersionConstraint(majorVersionConstraint string, ) *OSVersionConstraint`

NewOSVersionConstraint instantiates a new OSVersionConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSVersionConstraintWithDefaults

`func NewOSVersionConstraintWithDefaults() *OSVersionConstraint`

NewOSVersionConstraintWithDefaults instantiates a new OSVersionConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDynamicVersionRequirement

`func (o *OSVersionConstraint) GetDynamicVersionRequirement() OSVersionConstraintDynamicVersionRequirement`

GetDynamicVersionRequirement returns the DynamicVersionRequirement field if non-nil, zero value otherwise.

### GetDynamicVersionRequirementOk

`func (o *OSVersionConstraint) GetDynamicVersionRequirementOk() (*OSVersionConstraintDynamicVersionRequirement, bool)`

GetDynamicVersionRequirementOk returns a tuple with the DynamicVersionRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicVersionRequirement

`func (o *OSVersionConstraint) SetDynamicVersionRequirement(v OSVersionConstraintDynamicVersionRequirement)`

SetDynamicVersionRequirement sets DynamicVersionRequirement field to given value.

### HasDynamicVersionRequirement

`func (o *OSVersionConstraint) HasDynamicVersionRequirement() bool`

HasDynamicVersionRequirement returns a boolean if a field has been set.

### GetMajorVersionConstraint

`func (o *OSVersionConstraint) GetMajorVersionConstraint() string`

GetMajorVersionConstraint returns the MajorVersionConstraint field if non-nil, zero value otherwise.

### GetMajorVersionConstraintOk

`func (o *OSVersionConstraint) GetMajorVersionConstraintOk() (*string, bool)`

GetMajorVersionConstraintOk returns a tuple with the MajorVersionConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorVersionConstraint

`func (o *OSVersionConstraint) SetMajorVersionConstraint(v string)`

SetMajorVersionConstraint sets MajorVersionConstraint field to given value.


### GetMinimum

`func (o *OSVersionConstraint) GetMinimum() string`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *OSVersionConstraint) GetMinimumOk() (*string, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *OSVersionConstraint) SetMinimum(v string)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *OSVersionConstraint) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


