# SamlClaimsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The attribute name | [optional] 
**Values** | Pointer to **[]string** | The Okta values inserted in the attribute statement | [optional] 

## Methods

### NewSamlClaimsInner

`func NewSamlClaimsInner() *SamlClaimsInner`

NewSamlClaimsInner instantiates a new SamlClaimsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlClaimsInnerWithDefaults

`func NewSamlClaimsInnerWithDefaults() *SamlClaimsInner`

NewSamlClaimsInnerWithDefaults instantiates a new SamlClaimsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SamlClaimsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlClaimsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlClaimsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SamlClaimsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValues

`func (o *SamlClaimsInner) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SamlClaimsInner) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SamlClaimsInner) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *SamlClaimsInner) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


