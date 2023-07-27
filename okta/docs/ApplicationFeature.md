# ApplicationFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**CapabilitiesObject**](CapabilitiesObject.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**EnabledStatus**](EnabledStatus.md) |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewApplicationFeature

`func NewApplicationFeature() *ApplicationFeature`

NewApplicationFeature instantiates a new ApplicationFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationFeatureWithDefaults

`func NewApplicationFeatureWithDefaults() *ApplicationFeature`

NewApplicationFeatureWithDefaults instantiates a new ApplicationFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *ApplicationFeature) GetCapabilities() CapabilitiesObject`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ApplicationFeature) GetCapabilitiesOk() (*CapabilitiesObject, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ApplicationFeature) SetCapabilities(v CapabilitiesObject)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ApplicationFeature) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationFeature) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationFeature) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationFeature) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationFeature) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ApplicationFeature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationFeature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationFeature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationFeature) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ApplicationFeature) GetStatus() EnabledStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationFeature) GetStatusOk() (*EnabledStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationFeature) SetStatus(v EnabledStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplicationFeature) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *ApplicationFeature) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationFeature) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationFeature) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationFeature) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


