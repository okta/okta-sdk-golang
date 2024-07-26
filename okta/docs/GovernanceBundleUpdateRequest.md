# GovernanceBundleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Entitlements** | Pointer to [**[]IAMBundleEntitlement**](IAMBundleEntitlement.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewGovernanceBundleUpdateRequest

`func NewGovernanceBundleUpdateRequest() *GovernanceBundleUpdateRequest`

NewGovernanceBundleUpdateRequest instantiates a new GovernanceBundleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGovernanceBundleUpdateRequestWithDefaults

`func NewGovernanceBundleUpdateRequestWithDefaults() *GovernanceBundleUpdateRequest`

NewGovernanceBundleUpdateRequestWithDefaults instantiates a new GovernanceBundleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GovernanceBundleUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GovernanceBundleUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GovernanceBundleUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GovernanceBundleUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntitlements

`func (o *GovernanceBundleUpdateRequest) GetEntitlements() []IAMBundleEntitlement`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *GovernanceBundleUpdateRequest) GetEntitlementsOk() (*[]IAMBundleEntitlement, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *GovernanceBundleUpdateRequest) SetEntitlements(v []IAMBundleEntitlement)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *GovernanceBundleUpdateRequest) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetName

`func (o *GovernanceBundleUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GovernanceBundleUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GovernanceBundleUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GovernanceBundleUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


