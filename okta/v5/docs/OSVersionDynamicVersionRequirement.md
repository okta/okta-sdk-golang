# OSVersionDynamicVersionRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Indicates the type of the dynamic OS version requirement | [optional] 
**DistanceFromLatestMajor** | Pointer to **int32** | Indicates the distance from the latest major version | [optional] 
**LatestSecurityPatch** | Pointer to **bool** | Indicates whether the device needs to be on the latest security patch | [optional] 

## Methods

### NewOSVersionDynamicVersionRequirement

`func NewOSVersionDynamicVersionRequirement() *OSVersionDynamicVersionRequirement`

NewOSVersionDynamicVersionRequirement instantiates a new OSVersionDynamicVersionRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSVersionDynamicVersionRequirementWithDefaults

`func NewOSVersionDynamicVersionRequirementWithDefaults() *OSVersionDynamicVersionRequirement`

NewOSVersionDynamicVersionRequirementWithDefaults instantiates a new OSVersionDynamicVersionRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OSVersionDynamicVersionRequirement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OSVersionDynamicVersionRequirement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OSVersionDynamicVersionRequirement) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OSVersionDynamicVersionRequirement) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDistanceFromLatestMajor

`func (o *OSVersionDynamicVersionRequirement) GetDistanceFromLatestMajor() int32`

GetDistanceFromLatestMajor returns the DistanceFromLatestMajor field if non-nil, zero value otherwise.

### GetDistanceFromLatestMajorOk

`func (o *OSVersionDynamicVersionRequirement) GetDistanceFromLatestMajorOk() (*int32, bool)`

GetDistanceFromLatestMajorOk returns a tuple with the DistanceFromLatestMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceFromLatestMajor

`func (o *OSVersionDynamicVersionRequirement) SetDistanceFromLatestMajor(v int32)`

SetDistanceFromLatestMajor sets DistanceFromLatestMajor field to given value.

### HasDistanceFromLatestMajor

`func (o *OSVersionDynamicVersionRequirement) HasDistanceFromLatestMajor() bool`

HasDistanceFromLatestMajor returns a boolean if a field has been set.

### GetLatestSecurityPatch

`func (o *OSVersionDynamicVersionRequirement) GetLatestSecurityPatch() bool`

GetLatestSecurityPatch returns the LatestSecurityPatch field if non-nil, zero value otherwise.

### GetLatestSecurityPatchOk

`func (o *OSVersionDynamicVersionRequirement) GetLatestSecurityPatchOk() (*bool, bool)`

GetLatestSecurityPatchOk returns a tuple with the LatestSecurityPatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSecurityPatch

`func (o *OSVersionDynamicVersionRequirement) SetLatestSecurityPatch(v bool)`

SetLatestSecurityPatch sets LatestSecurityPatch field to given value.

### HasLatestSecurityPatch

`func (o *OSVersionDynamicVersionRequirement) HasLatestSecurityPatch() bool`

HasLatestSecurityPatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


