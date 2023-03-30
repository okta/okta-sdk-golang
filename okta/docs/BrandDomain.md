# BrandDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | Pointer to **string** |  | [optional] [readonly] 
**Links** | Pointer to [**BrandDomainLinks**](BrandDomainLinks.md) |  | [optional] 

## Methods

### NewBrandDomain

`func NewBrandDomain() *BrandDomain`

NewBrandDomain instantiates a new BrandDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandDomainWithDefaults

`func NewBrandDomainWithDefaults() *BrandDomain`

NewBrandDomainWithDefaults instantiates a new BrandDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *BrandDomain) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *BrandDomain) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *BrandDomain) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *BrandDomain) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetLinks

`func (o *BrandDomain) GetLinks() BrandDomainLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BrandDomain) GetLinksOk() (*BrandDomainLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BrandDomain) SetLinks(v BrandDomainLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BrandDomain) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


