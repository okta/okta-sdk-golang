# UpdateFeatureForApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to [**CapabilitiesCreateObject**](CapabilitiesCreateObject.md) |  | [optional] 
**Update** | Pointer to [**CapabilitiesUpdateObject**](CapabilitiesUpdateObject.md) |  | [optional] 
**ImportRules** | [**CapabilitiesImportRulesObject**](CapabilitiesImportRulesObject.md) |  | 
**ImportSettings** | [**CapabilitiesImportSettingsObject**](CapabilitiesImportSettingsObject.md) |  | 

## Methods

### NewUpdateFeatureForApplicationRequest

`func NewUpdateFeatureForApplicationRequest(importRules CapabilitiesImportRulesObject, importSettings CapabilitiesImportSettingsObject, ) *UpdateFeatureForApplicationRequest`

NewUpdateFeatureForApplicationRequest instantiates a new UpdateFeatureForApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFeatureForApplicationRequestWithDefaults

`func NewUpdateFeatureForApplicationRequestWithDefaults() *UpdateFeatureForApplicationRequest`

NewUpdateFeatureForApplicationRequestWithDefaults instantiates a new UpdateFeatureForApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *UpdateFeatureForApplicationRequest) GetCreate() CapabilitiesCreateObject`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *UpdateFeatureForApplicationRequest) GetCreateOk() (*CapabilitiesCreateObject, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *UpdateFeatureForApplicationRequest) SetCreate(v CapabilitiesCreateObject)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *UpdateFeatureForApplicationRequest) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *UpdateFeatureForApplicationRequest) GetUpdate() CapabilitiesUpdateObject`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *UpdateFeatureForApplicationRequest) GetUpdateOk() (*CapabilitiesUpdateObject, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *UpdateFeatureForApplicationRequest) SetUpdate(v CapabilitiesUpdateObject)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *UpdateFeatureForApplicationRequest) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetImportRules

`func (o *UpdateFeatureForApplicationRequest) GetImportRules() CapabilitiesImportRulesObject`

GetImportRules returns the ImportRules field if non-nil, zero value otherwise.

### GetImportRulesOk

`func (o *UpdateFeatureForApplicationRequest) GetImportRulesOk() (*CapabilitiesImportRulesObject, bool)`

GetImportRulesOk returns a tuple with the ImportRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportRules

`func (o *UpdateFeatureForApplicationRequest) SetImportRules(v CapabilitiesImportRulesObject)`

SetImportRules sets ImportRules field to given value.


### GetImportSettings

`func (o *UpdateFeatureForApplicationRequest) GetImportSettings() CapabilitiesImportSettingsObject`

GetImportSettings returns the ImportSettings field if non-nil, zero value otherwise.

### GetImportSettingsOk

`func (o *UpdateFeatureForApplicationRequest) GetImportSettingsOk() (*CapabilitiesImportSettingsObject, bool)`

GetImportSettingsOk returns a tuple with the ImportSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportSettings

`func (o *UpdateFeatureForApplicationRequest) SetImportSettings(v CapabilitiesImportSettingsObject)`

SetImportSettings sets ImportSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


