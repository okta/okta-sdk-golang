# ApplicationFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the feature | [optional] [readonly] 
**Name** | Pointer to **string** | Identifying name of the feature  | Value | Description   | | --------- | ------------- | | USER_PROVISIONING  | Represents the **To App** provisioning feature setting in the Admin Console | | INBOUND_PROVISIONING | Represents the **To Okta** provisioning feature setting in the Admin Console |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**ApplicationFeatureLinks**](ApplicationFeatureLinks.md) |  | [optional] 

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

`func (o *ApplicationFeature) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationFeature) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationFeature) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplicationFeature) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *ApplicationFeature) GetLinks() ApplicationFeatureLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationFeature) GetLinksOk() (*ApplicationFeatureLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationFeature) SetLinks(v ApplicationFeatureLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationFeature) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


