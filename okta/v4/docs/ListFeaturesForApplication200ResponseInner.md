# ListFeaturesForApplication200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the feature | [optional] [readonly] 
**Name** | Pointer to **string** | Identifying name of the feature  | Value | Description   | | --------- | ------------- | | USER_PROVISIONING  | Represents the **To App** provisioning feature setting in the Admin Console | | INBOUND_PROVISIONING | Represents the **To Okta** provisioning feature setting in the Admin Console |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**ApplicationFeatureLinks**](ApplicationFeatureLinks.md) |  | [optional] 
**Capabilities** | Pointer to [**CapabilitiesInboundProvisioningObject**](CapabilitiesInboundProvisioningObject.md) |  | [optional] 

## Methods

### NewListFeaturesForApplication200ResponseInner

`func NewListFeaturesForApplication200ResponseInner() *ListFeaturesForApplication200ResponseInner`

NewListFeaturesForApplication200ResponseInner instantiates a new ListFeaturesForApplication200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFeaturesForApplication200ResponseInnerWithDefaults

`func NewListFeaturesForApplication200ResponseInnerWithDefaults() *ListFeaturesForApplication200ResponseInner`

NewListFeaturesForApplication200ResponseInnerWithDefaults instantiates a new ListFeaturesForApplication200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ListFeaturesForApplication200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListFeaturesForApplication200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListFeaturesForApplication200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListFeaturesForApplication200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ListFeaturesForApplication200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListFeaturesForApplication200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListFeaturesForApplication200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListFeaturesForApplication200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ListFeaturesForApplication200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListFeaturesForApplication200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListFeaturesForApplication200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListFeaturesForApplication200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *ListFeaturesForApplication200ResponseInner) GetLinks() ApplicationFeatureLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListFeaturesForApplication200ResponseInner) GetLinksOk() (*ApplicationFeatureLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListFeaturesForApplication200ResponseInner) SetLinks(v ApplicationFeatureLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListFeaturesForApplication200ResponseInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCapabilities

`func (o *ListFeaturesForApplication200ResponseInner) GetCapabilities() CapabilitiesInboundProvisioningObject`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ListFeaturesForApplication200ResponseInner) GetCapabilitiesOk() (*CapabilitiesInboundProvisioningObject, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ListFeaturesForApplication200ResponseInner) SetCapabilities(v CapabilitiesInboundProvisioningObject)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ListFeaturesForApplication200ResponseInner) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


