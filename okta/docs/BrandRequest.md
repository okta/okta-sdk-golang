# BrandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreeToCustomPrivacyPolicy** | Pointer to **bool** | Consent for updating the custom privacy URL. Not required when resetting the URL. | [optional] 
**CustomPrivacyPolicyUrl** | Pointer to **string** | Custom privacy policy URL | [optional] 
**DefaultApp** | Pointer to [**DefaultApp**](DefaultApp.md) |  | [optional] 
**EmailDomainId** | Pointer to **string** | The ID of the email domain | [optional] 
**Locale** | Pointer to **string** | The language specified as an [IETF BCP 47 language tag](https://datatracker.ietf.org/doc/html/rfc5646) | [optional] 
**Name** | **string** | The name of the Brand | 
**RemovePoweredByOkta** | Pointer to **bool** | Removes \&quot;Powered by Okta\&quot; from the sign-in page in redirect authentication deployments, and \&quot;© [current year] Okta, Inc.\&quot; from the Okta End-User Dashboard | [optional] [default to false]

## Methods

### NewBrandRequest

`func NewBrandRequest(name string, ) *BrandRequest`

NewBrandRequest instantiates a new BrandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandRequestWithDefaults

`func NewBrandRequestWithDefaults() *BrandRequest`

NewBrandRequestWithDefaults instantiates a new BrandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreeToCustomPrivacyPolicy

`func (o *BrandRequest) GetAgreeToCustomPrivacyPolicy() bool`

GetAgreeToCustomPrivacyPolicy returns the AgreeToCustomPrivacyPolicy field if non-nil, zero value otherwise.

### GetAgreeToCustomPrivacyPolicyOk

`func (o *BrandRequest) GetAgreeToCustomPrivacyPolicyOk() (*bool, bool)`

GetAgreeToCustomPrivacyPolicyOk returns a tuple with the AgreeToCustomPrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreeToCustomPrivacyPolicy

`func (o *BrandRequest) SetAgreeToCustomPrivacyPolicy(v bool)`

SetAgreeToCustomPrivacyPolicy sets AgreeToCustomPrivacyPolicy field to given value.

### HasAgreeToCustomPrivacyPolicy

`func (o *BrandRequest) HasAgreeToCustomPrivacyPolicy() bool`

HasAgreeToCustomPrivacyPolicy returns a boolean if a field has been set.

### GetCustomPrivacyPolicyUrl

`func (o *BrandRequest) GetCustomPrivacyPolicyUrl() string`

GetCustomPrivacyPolicyUrl returns the CustomPrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetCustomPrivacyPolicyUrlOk

`func (o *BrandRequest) GetCustomPrivacyPolicyUrlOk() (*string, bool)`

GetCustomPrivacyPolicyUrlOk returns a tuple with the CustomPrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrivacyPolicyUrl

`func (o *BrandRequest) SetCustomPrivacyPolicyUrl(v string)`

SetCustomPrivacyPolicyUrl sets CustomPrivacyPolicyUrl field to given value.

### HasCustomPrivacyPolicyUrl

`func (o *BrandRequest) HasCustomPrivacyPolicyUrl() bool`

HasCustomPrivacyPolicyUrl returns a boolean if a field has been set.

### GetDefaultApp

`func (o *BrandRequest) GetDefaultApp() DefaultApp`

GetDefaultApp returns the DefaultApp field if non-nil, zero value otherwise.

### GetDefaultAppOk

`func (o *BrandRequest) GetDefaultAppOk() (*DefaultApp, bool)`

GetDefaultAppOk returns a tuple with the DefaultApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApp

`func (o *BrandRequest) SetDefaultApp(v DefaultApp)`

SetDefaultApp sets DefaultApp field to given value.

### HasDefaultApp

`func (o *BrandRequest) HasDefaultApp() bool`

HasDefaultApp returns a boolean if a field has been set.

### GetEmailDomainId

`func (o *BrandRequest) GetEmailDomainId() string`

GetEmailDomainId returns the EmailDomainId field if non-nil, zero value otherwise.

### GetEmailDomainIdOk

`func (o *BrandRequest) GetEmailDomainIdOk() (*string, bool)`

GetEmailDomainIdOk returns a tuple with the EmailDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomainId

`func (o *BrandRequest) SetEmailDomainId(v string)`

SetEmailDomainId sets EmailDomainId field to given value.

### HasEmailDomainId

`func (o *BrandRequest) HasEmailDomainId() bool`

HasEmailDomainId returns a boolean if a field has been set.

### GetLocale

`func (o *BrandRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *BrandRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *BrandRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *BrandRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetName

`func (o *BrandRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BrandRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BrandRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRemovePoweredByOkta

`func (o *BrandRequest) GetRemovePoweredByOkta() bool`

GetRemovePoweredByOkta returns the RemovePoweredByOkta field if non-nil, zero value otherwise.

### GetRemovePoweredByOktaOk

`func (o *BrandRequest) GetRemovePoweredByOktaOk() (*bool, bool)`

GetRemovePoweredByOktaOk returns a tuple with the RemovePoweredByOkta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePoweredByOkta

`func (o *BrandRequest) SetRemovePoweredByOkta(v bool)`

SetRemovePoweredByOkta sets RemovePoweredByOkta field to given value.

### HasRemovePoweredByOkta

`func (o *BrandRequest) HasRemovePoweredByOkta() bool`

HasRemovePoweredByOkta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


