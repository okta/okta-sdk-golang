# CimdClientEntityLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**CimdClientEntityLinksSelf**](CimdClientEntityLinksSelf.md) |  | [optional] 
**MetadataRequirements** | Pointer to [**CimdClientEntityLinksMetadataRequirements**](CimdClientEntityLinksMetadataRequirements.md) |  | [optional] 
**Bindings** | Pointer to [**[]CimdClientEntityLinksSelf**](CimdClientEntityLinksSelf.md) | Links to the entities that are bound to the CIMD client entity | [optional] 

## Methods

### NewCimdClientEntityLinks

`func NewCimdClientEntityLinks() *CimdClientEntityLinks`

NewCimdClientEntityLinks instantiates a new CimdClientEntityLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCimdClientEntityLinksWithDefaults

`func NewCimdClientEntityLinksWithDefaults() *CimdClientEntityLinks`

NewCimdClientEntityLinksWithDefaults instantiates a new CimdClientEntityLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *CimdClientEntityLinks) GetSelf() CimdClientEntityLinksSelf`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *CimdClientEntityLinks) GetSelfOk() (*CimdClientEntityLinksSelf, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *CimdClientEntityLinks) SetSelf(v CimdClientEntityLinksSelf)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *CimdClientEntityLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetMetadataRequirements

`func (o *CimdClientEntityLinks) GetMetadataRequirements() CimdClientEntityLinksMetadataRequirements`

GetMetadataRequirements returns the MetadataRequirements field if non-nil, zero value otherwise.

### GetMetadataRequirementsOk

`func (o *CimdClientEntityLinks) GetMetadataRequirementsOk() (*CimdClientEntityLinksMetadataRequirements, bool)`

GetMetadataRequirementsOk returns a tuple with the MetadataRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataRequirements

`func (o *CimdClientEntityLinks) SetMetadataRequirements(v CimdClientEntityLinksMetadataRequirements)`

SetMetadataRequirements sets MetadataRequirements field to given value.

### HasMetadataRequirements

`func (o *CimdClientEntityLinks) HasMetadataRequirements() bool`

HasMetadataRequirements returns a boolean if a field has been set.

### GetBindings

`func (o *CimdClientEntityLinks) GetBindings() []CimdClientEntityLinksSelf`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *CimdClientEntityLinks) GetBindingsOk() (*[]CimdClientEntityLinksSelf, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *CimdClientEntityLinks) SetBindings(v []CimdClientEntityLinksSelf)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *CimdClientEntityLinks) HasBindings() bool`

HasBindings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


