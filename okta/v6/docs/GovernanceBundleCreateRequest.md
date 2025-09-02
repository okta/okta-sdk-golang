# GovernanceBundleCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Entitlements** | Pointer to [**[]IAMBundleEntitlement**](IAMBundleEntitlement.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewGovernanceBundleCreateRequest

`func NewGovernanceBundleCreateRequest() *GovernanceBundleCreateRequest`

NewGovernanceBundleCreateRequest instantiates a new GovernanceBundleCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGovernanceBundleCreateRequestWithDefaults

`func NewGovernanceBundleCreateRequestWithDefaults() *GovernanceBundleCreateRequest`

NewGovernanceBundleCreateRequestWithDefaults instantiates a new GovernanceBundleCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GovernanceBundleCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GovernanceBundleCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GovernanceBundleCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GovernanceBundleCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntitlements

`func (o *GovernanceBundleCreateRequest) GetEntitlements() []IAMBundleEntitlement`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *GovernanceBundleCreateRequest) GetEntitlementsOk() (*[]IAMBundleEntitlement, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *GovernanceBundleCreateRequest) SetEntitlements(v []IAMBundleEntitlement)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *GovernanceBundleCreateRequest) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetName

`func (o *GovernanceBundleCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GovernanceBundleCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GovernanceBundleCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GovernanceBundleCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


