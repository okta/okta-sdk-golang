# SamlAcsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **float32** | Index of ACS URL. You can&#39;t reuse the same index in the ACS URL array. | [optional] 
**Url** | Pointer to **string** | Assertion Consumer Service (ACS) URL | [optional] 

## Methods

### NewSamlAcsInner

`func NewSamlAcsInner() *SamlAcsInner`

NewSamlAcsInner instantiates a new SamlAcsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlAcsInnerWithDefaults

`func NewSamlAcsInnerWithDefaults() *SamlAcsInner`

NewSamlAcsInnerWithDefaults instantiates a new SamlAcsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *SamlAcsInner) GetIndex() float32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SamlAcsInner) GetIndexOk() (*float32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SamlAcsInner) SetIndex(v float32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *SamlAcsInner) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetUrl

`func (o *SamlAcsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SamlAcsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SamlAcsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SamlAcsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


