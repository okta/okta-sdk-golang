# BundleEntitlementsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlements** | Pointer to [**[]BundleEntitlement**](BundleEntitlement.md) |  | [optional] 
**Links** | Pointer to [**NullableBundleEntitlementsResponseLinks**](BundleEntitlementsResponseLinks.md) |  | [optional] 

## Methods

### NewBundleEntitlementsResponse

`func NewBundleEntitlementsResponse() *BundleEntitlementsResponse`

NewBundleEntitlementsResponse instantiates a new BundleEntitlementsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleEntitlementsResponseWithDefaults

`func NewBundleEntitlementsResponseWithDefaults() *BundleEntitlementsResponse`

NewBundleEntitlementsResponseWithDefaults instantiates a new BundleEntitlementsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlements

`func (o *BundleEntitlementsResponse) GetEntitlements() []BundleEntitlement`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *BundleEntitlementsResponse) GetEntitlementsOk() (*[]BundleEntitlement, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *BundleEntitlementsResponse) SetEntitlements(v []BundleEntitlement)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *BundleEntitlementsResponse) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetLinks

`func (o *BundleEntitlementsResponse) GetLinks() BundleEntitlementsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BundleEntitlementsResponse) GetLinksOk() (*BundleEntitlementsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BundleEntitlementsResponse) SetLinks(v BundleEntitlementsResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BundleEntitlementsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *BundleEntitlementsResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *BundleEntitlementsResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


