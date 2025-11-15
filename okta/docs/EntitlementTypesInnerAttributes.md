# EntitlementTypesInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Required** | Pointer to **bool** | A boolean value to indicate if this entitlement type is required for the user | [optional] [default to false]
**Multivalued** | Pointer to **bool** | A boolean value to indicate if a user can have multiple entitlements of this type | [optional] [default to false]

## Methods

### NewEntitlementTypesInnerAttributes

`func NewEntitlementTypesInnerAttributes() *EntitlementTypesInnerAttributes`

NewEntitlementTypesInnerAttributes instantiates a new EntitlementTypesInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementTypesInnerAttributesWithDefaults

`func NewEntitlementTypesInnerAttributesWithDefaults() *EntitlementTypesInnerAttributes`

NewEntitlementTypesInnerAttributesWithDefaults instantiates a new EntitlementTypesInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequired

`func (o *EntitlementTypesInnerAttributes) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *EntitlementTypesInnerAttributes) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *EntitlementTypesInnerAttributes) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *EntitlementTypesInnerAttributes) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetMultivalued

`func (o *EntitlementTypesInnerAttributes) GetMultivalued() bool`

GetMultivalued returns the Multivalued field if non-nil, zero value otherwise.

### GetMultivaluedOk

`func (o *EntitlementTypesInnerAttributes) GetMultivaluedOk() (*bool, bool)`

GetMultivaluedOk returns a tuple with the Multivalued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultivalued

`func (o *EntitlementTypesInnerAttributes) SetMultivalued(v bool)`

SetMultivalued sets Multivalued field to given value.

### HasMultivalued

`func (o *EntitlementTypesInnerAttributes) HasMultivalued() bool`

HasMultivalued returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


