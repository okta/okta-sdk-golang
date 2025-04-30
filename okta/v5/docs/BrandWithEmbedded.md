# BrandWithEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 
**AgreeToCustomPrivacyPolicy** | Pointer to **bool** | Consent for updating the custom privacy URL. Not required when resetting the URL. | [optional] 
**CustomPrivacyPolicyUrl** | Pointer to **string** | Custom privacy policy URL | [optional] 
**DefaultApp** | Pointer to [**DefaultApp**](DefaultApp.md) |  | [optional] 
**EmailDomainId** | Pointer to **string** | The ID of the email domain | [optional] 
**Id** | Pointer to **string** | The Brand ID | [optional] [readonly] 
**IsDefault** | Pointer to **bool** | If &#x60;true&#x60;, the Brand is used for the Okta subdomain | [optional] [readonly] 
**Locale** | Pointer to **string** | The language specified as an [IETF BCP 47 language tag](https://datatracker.ietf.org/doc/html/rfc5646) | [optional] 
**Name** | Pointer to **string** | The name of the Brand | [optional] 
**RemovePoweredByOkta** | Pointer to **bool** | Removes \&quot;Powered by Okta\&quot; from the sign-in page in redirect authentication deployments, and \&quot;© [current year] Okta, Inc.\&quot; from the Okta End-User Dashboard | [optional] [default to false]

## Methods

### NewBrandWithEmbedded

`func NewBrandWithEmbedded() *BrandWithEmbedded`

NewBrandWithEmbedded instantiates a new BrandWithEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandWithEmbeddedWithDefaults

`func NewBrandWithEmbeddedWithDefaults() *BrandWithEmbedded`

NewBrandWithEmbeddedWithDefaults instantiates a new BrandWithEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *BrandWithEmbedded) GetEmbedded() map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *BrandWithEmbedded) GetEmbeddedOk() (*map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *BrandWithEmbedded) SetEmbedded(v map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *BrandWithEmbedded) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *BrandWithEmbedded) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BrandWithEmbedded) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BrandWithEmbedded) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BrandWithEmbedded) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAgreeToCustomPrivacyPolicy

`func (o *BrandWithEmbedded) GetAgreeToCustomPrivacyPolicy() bool`

GetAgreeToCustomPrivacyPolicy returns the AgreeToCustomPrivacyPolicy field if non-nil, zero value otherwise.

### GetAgreeToCustomPrivacyPolicyOk

`func (o *BrandWithEmbedded) GetAgreeToCustomPrivacyPolicyOk() (*bool, bool)`

GetAgreeToCustomPrivacyPolicyOk returns a tuple with the AgreeToCustomPrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreeToCustomPrivacyPolicy

`func (o *BrandWithEmbedded) SetAgreeToCustomPrivacyPolicy(v bool)`

SetAgreeToCustomPrivacyPolicy sets AgreeToCustomPrivacyPolicy field to given value.

### HasAgreeToCustomPrivacyPolicy

`func (o *BrandWithEmbedded) HasAgreeToCustomPrivacyPolicy() bool`

HasAgreeToCustomPrivacyPolicy returns a boolean if a field has been set.

### GetCustomPrivacyPolicyUrl

`func (o *BrandWithEmbedded) GetCustomPrivacyPolicyUrl() string`

GetCustomPrivacyPolicyUrl returns the CustomPrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetCustomPrivacyPolicyUrlOk

`func (o *BrandWithEmbedded) GetCustomPrivacyPolicyUrlOk() (*string, bool)`

GetCustomPrivacyPolicyUrlOk returns a tuple with the CustomPrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrivacyPolicyUrl

`func (o *BrandWithEmbedded) SetCustomPrivacyPolicyUrl(v string)`

SetCustomPrivacyPolicyUrl sets CustomPrivacyPolicyUrl field to given value.

### HasCustomPrivacyPolicyUrl

`func (o *BrandWithEmbedded) HasCustomPrivacyPolicyUrl() bool`

HasCustomPrivacyPolicyUrl returns a boolean if a field has been set.

### GetDefaultApp

`func (o *BrandWithEmbedded) GetDefaultApp() DefaultApp`

GetDefaultApp returns the DefaultApp field if non-nil, zero value otherwise.

### GetDefaultAppOk

`func (o *BrandWithEmbedded) GetDefaultAppOk() (*DefaultApp, bool)`

GetDefaultAppOk returns a tuple with the DefaultApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApp

`func (o *BrandWithEmbedded) SetDefaultApp(v DefaultApp)`

SetDefaultApp sets DefaultApp field to given value.

### HasDefaultApp

`func (o *BrandWithEmbedded) HasDefaultApp() bool`

HasDefaultApp returns a boolean if a field has been set.

### GetEmailDomainId

`func (o *BrandWithEmbedded) GetEmailDomainId() string`

GetEmailDomainId returns the EmailDomainId field if non-nil, zero value otherwise.

### GetEmailDomainIdOk

`func (o *BrandWithEmbedded) GetEmailDomainIdOk() (*string, bool)`

GetEmailDomainIdOk returns a tuple with the EmailDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomainId

`func (o *BrandWithEmbedded) SetEmailDomainId(v string)`

SetEmailDomainId sets EmailDomainId field to given value.

### HasEmailDomainId

`func (o *BrandWithEmbedded) HasEmailDomainId() bool`

HasEmailDomainId returns a boolean if a field has been set.

### GetId

`func (o *BrandWithEmbedded) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BrandWithEmbedded) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BrandWithEmbedded) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BrandWithEmbedded) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *BrandWithEmbedded) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *BrandWithEmbedded) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *BrandWithEmbedded) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *BrandWithEmbedded) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLocale

`func (o *BrandWithEmbedded) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *BrandWithEmbedded) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *BrandWithEmbedded) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *BrandWithEmbedded) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetName

`func (o *BrandWithEmbedded) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BrandWithEmbedded) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BrandWithEmbedded) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BrandWithEmbedded) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemovePoweredByOkta

`func (o *BrandWithEmbedded) GetRemovePoweredByOkta() bool`

GetRemovePoweredByOkta returns the RemovePoweredByOkta field if non-nil, zero value otherwise.

### GetRemovePoweredByOktaOk

`func (o *BrandWithEmbedded) GetRemovePoweredByOktaOk() (*bool, bool)`

GetRemovePoweredByOktaOk returns a tuple with the RemovePoweredByOkta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePoweredByOkta

`func (o *BrandWithEmbedded) SetRemovePoweredByOkta(v bool)`

SetRemovePoweredByOkta sets RemovePoweredByOkta field to given value.

### HasRemovePoweredByOkta

`func (o *BrandWithEmbedded) HasRemovePoweredByOkta() bool`

HasRemovePoweredByOkta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


