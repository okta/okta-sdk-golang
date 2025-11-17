# EntitlementTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The entitlement type name | 
**Description** | Pointer to **string** | Description of the entitlement type | [optional] 
**Endpoint** | **string** | URL of the entitlement type endpoint | 
**Attributes** | [**EntitlementTypesInnerAttributes**](EntitlementTypesInnerAttributes.md) |  | 
**Mappings** | [**EntitlementTypesInnerMappings**](EntitlementTypesInnerMappings.md) |  | 

## Methods

### NewEntitlementTypesInner

`func NewEntitlementTypesInner(name string, endpoint string, attributes EntitlementTypesInnerAttributes, mappings EntitlementTypesInnerMappings, ) *EntitlementTypesInner`

NewEntitlementTypesInner instantiates a new EntitlementTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementTypesInnerWithDefaults

`func NewEntitlementTypesInnerWithDefaults() *EntitlementTypesInner`

NewEntitlementTypesInnerWithDefaults instantiates a new EntitlementTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EntitlementTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementTypesInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EntitlementTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoint

`func (o *EntitlementTypesInner) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *EntitlementTypesInner) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *EntitlementTypesInner) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetAttributes

`func (o *EntitlementTypesInner) GetAttributes() EntitlementTypesInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EntitlementTypesInner) GetAttributesOk() (*EntitlementTypesInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EntitlementTypesInner) SetAttributes(v EntitlementTypesInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetMappings

`func (o *EntitlementTypesInner) GetMappings() EntitlementTypesInnerMappings`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *EntitlementTypesInner) GetMappingsOk() (*EntitlementTypesInnerMappings, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *EntitlementTypesInner) SetMappings(v EntitlementTypesInnerMappings)`

SetMappings sets Mappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


