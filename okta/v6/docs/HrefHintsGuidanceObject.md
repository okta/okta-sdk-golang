# HrefHintsGuidanceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allow** | Pointer to **[]string** |  | [optional] 
**Guidance** | Pointer to **[]string** | Specifies the URI to invoke for granting scope consent required to complete the OAuth 2.0 connection  | [optional] 

## Methods

### NewHrefHintsGuidanceObject

`func NewHrefHintsGuidanceObject() *HrefHintsGuidanceObject`

NewHrefHintsGuidanceObject instantiates a new HrefHintsGuidanceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHrefHintsGuidanceObjectWithDefaults

`func NewHrefHintsGuidanceObjectWithDefaults() *HrefHintsGuidanceObject`

NewHrefHintsGuidanceObjectWithDefaults instantiates a new HrefHintsGuidanceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllow

`func (o *HrefHintsGuidanceObject) GetAllow() []string`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *HrefHintsGuidanceObject) GetAllowOk() (*[]string, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *HrefHintsGuidanceObject) SetAllow(v []string)`

SetAllow sets Allow field to given value.

### HasAllow

`func (o *HrefHintsGuidanceObject) HasAllow() bool`

HasAllow returns a boolean if a field has been set.

### GetGuidance

`func (o *HrefHintsGuidanceObject) GetGuidance() []string`

GetGuidance returns the Guidance field if non-nil, zero value otherwise.

### GetGuidanceOk

`func (o *HrefHintsGuidanceObject) GetGuidanceOk() (*[]string, bool)`

GetGuidanceOk returns a tuple with the Guidance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidance

`func (o *HrefHintsGuidanceObject) SetGuidance(v []string)`

SetGuidance sets Guidance field to given value.

### HasGuidance

`func (o *HrefHintsGuidanceObject) HasGuidance() bool`

HasGuidance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


