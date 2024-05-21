# OSVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DynamicVersionRequirement** | Pointer to [**OSVersionDynamicVersionRequirement**](OSVersionDynamicVersionRequirement.md) |  | [optional] 
**Minimum** | Pointer to **string** | The device version must be equal to or newer than the specified version string (maximum of three components for iOS and macOS, and maximum of four components for Android) | [optional] 

## Methods

### NewOSVersion

`func NewOSVersion() *OSVersion`

NewOSVersion instantiates a new OSVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSVersionWithDefaults

`func NewOSVersionWithDefaults() *OSVersion`

NewOSVersionWithDefaults instantiates a new OSVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDynamicVersionRequirement

`func (o *OSVersion) GetDynamicVersionRequirement() OSVersionDynamicVersionRequirement`

GetDynamicVersionRequirement returns the DynamicVersionRequirement field if non-nil, zero value otherwise.

### GetDynamicVersionRequirementOk

`func (o *OSVersion) GetDynamicVersionRequirementOk() (*OSVersionDynamicVersionRequirement, bool)`

GetDynamicVersionRequirementOk returns a tuple with the DynamicVersionRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicVersionRequirement

`func (o *OSVersion) SetDynamicVersionRequirement(v OSVersionDynamicVersionRequirement)`

SetDynamicVersionRequirement sets DynamicVersionRequirement field to given value.

### HasDynamicVersionRequirement

`func (o *OSVersion) HasDynamicVersionRequirement() bool`

HasDynamicVersionRequirement returns a boolean if a field has been set.

### GetMinimum

`func (o *OSVersion) GetMinimum() string`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *OSVersion) GetMinimumOk() (*string, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *OSVersion) SetMinimum(v string)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *OSVersion) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


