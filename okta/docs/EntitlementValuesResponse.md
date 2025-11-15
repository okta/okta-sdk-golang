# EntitlementValuesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementValues** | Pointer to [**[]EntitlementValue**](EntitlementValue.md) | List of entitlement values for a bundle entitlement | [optional] 
**Links** | Pointer to [**EntitlementValuesResponseLinks**](EntitlementValuesResponseLinks.md) |  | [optional] 

## Methods

### NewEntitlementValuesResponse

`func NewEntitlementValuesResponse() *EntitlementValuesResponse`

NewEntitlementValuesResponse instantiates a new EntitlementValuesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementValuesResponseWithDefaults

`func NewEntitlementValuesResponseWithDefaults() *EntitlementValuesResponse`

NewEntitlementValuesResponseWithDefaults instantiates a new EntitlementValuesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementValues

`func (o *EntitlementValuesResponse) GetEntitlementValues() []EntitlementValue`

GetEntitlementValues returns the EntitlementValues field if non-nil, zero value otherwise.

### GetEntitlementValuesOk

`func (o *EntitlementValuesResponse) GetEntitlementValuesOk() (*[]EntitlementValue, bool)`

GetEntitlementValuesOk returns a tuple with the EntitlementValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementValues

`func (o *EntitlementValuesResponse) SetEntitlementValues(v []EntitlementValue)`

SetEntitlementValues sets EntitlementValues field to given value.

### HasEntitlementValues

`func (o *EntitlementValuesResponse) HasEntitlementValues() bool`

HasEntitlementValues returns a boolean if a field has been set.

### GetLinks

`func (o *EntitlementValuesResponse) GetLinks() EntitlementValuesResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntitlementValuesResponse) GetLinksOk() (*EntitlementValuesResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntitlementValuesResponse) SetLinks(v EntitlementValuesResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EntitlementValuesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


