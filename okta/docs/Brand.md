# Brand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreeToCustomPrivacyPolicy** | Pointer to **bool** |  | [optional] 
**CustomPrivacyPolicyUrl** | Pointer to **string** |  | [optional] 
**DefaultApp** | Pointer to [**BrandDefaultApp**](BrandDefaultApp.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IsDefault** | Pointer to **bool** |  | [optional] [readonly] 
**Locale** | Pointer to **string** | The language specified as an [IETF BCP 47 language tag](https://datatracker.ietf.org/doc/html/rfc5646). | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RemovePoweredByOkta** | Pointer to **bool** |  | [optional] 
**Links** | Pointer to [**BrandLinks**](BrandLinks.md) |  | [optional] 

## Methods

### NewBrand

`func NewBrand() *Brand`

NewBrand instantiates a new Brand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandWithDefaults

`func NewBrandWithDefaults() *Brand`

NewBrandWithDefaults instantiates a new Brand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreeToCustomPrivacyPolicy

`func (o *Brand) GetAgreeToCustomPrivacyPolicy() bool`

GetAgreeToCustomPrivacyPolicy returns the AgreeToCustomPrivacyPolicy field if non-nil, zero value otherwise.

### GetAgreeToCustomPrivacyPolicyOk

`func (o *Brand) GetAgreeToCustomPrivacyPolicyOk() (*bool, bool)`

GetAgreeToCustomPrivacyPolicyOk returns a tuple with the AgreeToCustomPrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreeToCustomPrivacyPolicy

`func (o *Brand) SetAgreeToCustomPrivacyPolicy(v bool)`

SetAgreeToCustomPrivacyPolicy sets AgreeToCustomPrivacyPolicy field to given value.

### HasAgreeToCustomPrivacyPolicy

`func (o *Brand) HasAgreeToCustomPrivacyPolicy() bool`

HasAgreeToCustomPrivacyPolicy returns a boolean if a field has been set.

### GetCustomPrivacyPolicyUrl

`func (o *Brand) GetCustomPrivacyPolicyUrl() string`

GetCustomPrivacyPolicyUrl returns the CustomPrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetCustomPrivacyPolicyUrlOk

`func (o *Brand) GetCustomPrivacyPolicyUrlOk() (*string, bool)`

GetCustomPrivacyPolicyUrlOk returns a tuple with the CustomPrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrivacyPolicyUrl

`func (o *Brand) SetCustomPrivacyPolicyUrl(v string)`

SetCustomPrivacyPolicyUrl sets CustomPrivacyPolicyUrl field to given value.

### HasCustomPrivacyPolicyUrl

`func (o *Brand) HasCustomPrivacyPolicyUrl() bool`

HasCustomPrivacyPolicyUrl returns a boolean if a field has been set.

### GetDefaultApp

`func (o *Brand) GetDefaultApp() BrandDefaultApp`

GetDefaultApp returns the DefaultApp field if non-nil, zero value otherwise.

### GetDefaultAppOk

`func (o *Brand) GetDefaultAppOk() (*BrandDefaultApp, bool)`

GetDefaultAppOk returns a tuple with the DefaultApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApp

`func (o *Brand) SetDefaultApp(v BrandDefaultApp)`

SetDefaultApp sets DefaultApp field to given value.

### HasDefaultApp

`func (o *Brand) HasDefaultApp() bool`

HasDefaultApp returns a boolean if a field has been set.

### GetId

`func (o *Brand) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Brand) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Brand) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Brand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *Brand) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Brand) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Brand) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Brand) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLocale

`func (o *Brand) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Brand) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Brand) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *Brand) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetName

`func (o *Brand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Brand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Brand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Brand) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemovePoweredByOkta

`func (o *Brand) GetRemovePoweredByOkta() bool`

GetRemovePoweredByOkta returns the RemovePoweredByOkta field if non-nil, zero value otherwise.

### GetRemovePoweredByOktaOk

`func (o *Brand) GetRemovePoweredByOktaOk() (*bool, bool)`

GetRemovePoweredByOktaOk returns a tuple with the RemovePoweredByOkta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePoweredByOkta

`func (o *Brand) SetRemovePoweredByOkta(v bool)`

SetRemovePoweredByOkta sets RemovePoweredByOkta field to given value.

### HasRemovePoweredByOkta

`func (o *Brand) HasRemovePoweredByOkta() bool`

HasRemovePoweredByOkta returns a boolean if a field has been set.

### GetLinks

`func (o *Brand) GetLinks() BrandLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Brand) GetLinksOk() (*BrandLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Brand) SetLinks(v BrandLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Brand) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


