# OSVersionConstraintDynamicVersionRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Indicates the type of the dynamic Windows version requirement | [optional] 
**DistanceFromLatestMajor** | Pointer to **int32** | Indicates the distance from the latest Windows major version | [optional] 
**LatestSecurityPatch** | Pointer to **bool** | Indicates whether the policy requires Windows devices to be on the latest security patch | [optional] 

## Methods

### NewOSVersionConstraintDynamicVersionRequirement

`func NewOSVersionConstraintDynamicVersionRequirement() *OSVersionConstraintDynamicVersionRequirement`

NewOSVersionConstraintDynamicVersionRequirement instantiates a new OSVersionConstraintDynamicVersionRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSVersionConstraintDynamicVersionRequirementWithDefaults

`func NewOSVersionConstraintDynamicVersionRequirementWithDefaults() *OSVersionConstraintDynamicVersionRequirement`

NewOSVersionConstraintDynamicVersionRequirementWithDefaults instantiates a new OSVersionConstraintDynamicVersionRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OSVersionConstraintDynamicVersionRequirement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OSVersionConstraintDynamicVersionRequirement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OSVersionConstraintDynamicVersionRequirement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OSVersionConstraintDynamicVersionRequirement) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDistanceFromLatestMajor

`func (o *OSVersionConstraintDynamicVersionRequirement) GetDistanceFromLatestMajor() int32`

GetDistanceFromLatestMajor returns the DistanceFromLatestMajor field if non-nil, zero value otherwise.

### GetDistanceFromLatestMajorOk

`func (o *OSVersionConstraintDynamicVersionRequirement) GetDistanceFromLatestMajorOk() (*int32, bool)`

GetDistanceFromLatestMajorOk returns a tuple with the DistanceFromLatestMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceFromLatestMajor

`func (o *OSVersionConstraintDynamicVersionRequirement) SetDistanceFromLatestMajor(v int32)`

SetDistanceFromLatestMajor sets DistanceFromLatestMajor field to given value.

### HasDistanceFromLatestMajor

`func (o *OSVersionConstraintDynamicVersionRequirement) HasDistanceFromLatestMajor() bool`

HasDistanceFromLatestMajor returns a boolean if a field has been set.

### GetLatestSecurityPatch

`func (o *OSVersionConstraintDynamicVersionRequirement) GetLatestSecurityPatch() bool`

GetLatestSecurityPatch returns the LatestSecurityPatch field if non-nil, zero value otherwise.

### GetLatestSecurityPatchOk

`func (o *OSVersionConstraintDynamicVersionRequirement) GetLatestSecurityPatchOk() (*bool, bool)`

GetLatestSecurityPatchOk returns a tuple with the LatestSecurityPatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSecurityPatch

`func (o *OSVersionConstraintDynamicVersionRequirement) SetLatestSecurityPatch(v bool)`

SetLatestSecurityPatch sets LatestSecurityPatch field to given value.

### HasLatestSecurityPatch

`func (o *OSVersionConstraintDynamicVersionRequirement) HasLatestSecurityPatch() bool`

HasLatestSecurityPatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


