# IdentityProviderProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AalValue** | Pointer to **NullableString** | The [authentication assurance level](https://developers.login.gov/oidc/#aal-values) (AAL) value for the Login.gov IdP. See [Add a Login.gov IdP](https://developer.okta.com/docs/guides/add-logingov-idp/). Applies to &#x60;LOGINGOV&#x60; and &#x60;LOGINGOV_SANDBOX&#x60; IdP types. | [optional] 
**AdditionalAmr** | Pointer to **[]string** | The additional Assurance Methods References (AMR) values for Smart Card IdPs. Applies to &#x60;X509&#x60; IdP type. | [optional] 
**AllowDynamicUserMatching** | Pointer to **bool** | Enables authenticated users to enroll a Smart Card if the users can&#39;t be identified using the configured &#x60;matchAttribute&#x60;. Applies to &#x60;X509&#x60; IdP type.  &gt; **Note:** Enrollment requires a PIN if the Smart Card is PIN-protected. And enrollment is also blocked if the X.509 certificate contains restricted attributes (&#x60;login&#x60;, &#x60;email&#x60;, &#x60;secondEmail&#x60;, and configured attributes that are used in multiple identifiers). | [optional] [default to false]
**AllowUserUpdates** | Pointer to **bool** | Enables attribute synchronization from the X.509 certificate to the user&#39;s Okta profile. If it&#39;s set to &#x60;true&#x60;, Okta overwrites user profile attributes with values found in the X.509 certificate. Applies to &#x60;X509&#x60; IdP type.  &gt; **Note:** Some attributes are restricted and won&#39;t be updated, including &#x60;login&#x60;, &#x60;email&#x60;, &#x60;secondEmail&#x60;, and configured attributes that are used in multiple identifiers.  | [optional] [default to false]
**IalValue** | Pointer to **NullableString** | The [type of identity verification](https://developers.login.gov/oidc/#ial-values) (IAL) value for the Login.gov IdP. See [Add a Login.gov IdP](https://developer.okta.com/docs/guides/add-logingov-idp/). Applies to &#x60;LOGINGOV&#x60; and &#x60;LOGINGOV_SANDBOX&#x60; IdP types. | [optional] 
**IdvMetadata** | Pointer to [**IdentityProviderPropertiesIdvMetadata**](IdentityProviderPropertiesIdvMetadata.md) |  | [optional] 
**InquiryTemplateId** | Pointer to **string** | The ID of the inquiry template from your Persona dashboard. The inquiry template always starts with &#x60;itmpl&#x60;. Applies to the legacy &#x60;IDV_PERSONA&#x60; IdP type. | [optional] 

## Methods

### NewIdentityProviderProperties

`func NewIdentityProviderProperties() *IdentityProviderProperties`

NewIdentityProviderProperties instantiates a new IdentityProviderProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderPropertiesWithDefaults

`func NewIdentityProviderPropertiesWithDefaults() *IdentityProviderProperties`

NewIdentityProviderPropertiesWithDefaults instantiates a new IdentityProviderProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAalValue

`func (o *IdentityProviderProperties) GetAalValue() string`

GetAalValue returns the AalValue field if non-nil, zero value otherwise.

### GetAalValueOk

`func (o *IdentityProviderProperties) GetAalValueOk() (*string, bool)`

GetAalValueOk returns a tuple with the AalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAalValue

`func (o *IdentityProviderProperties) SetAalValue(v string)`

SetAalValue sets AalValue field to given value.

### HasAalValue

`func (o *IdentityProviderProperties) HasAalValue() bool`

HasAalValue returns a boolean if a field has been set.

### SetAalValueNil

`func (o *IdentityProviderProperties) SetAalValueNil(b bool)`

 SetAalValueNil sets the value for AalValue to be an explicit nil

### UnsetAalValue
`func (o *IdentityProviderProperties) UnsetAalValue()`

UnsetAalValue ensures that no value is present for AalValue, not even an explicit nil
### GetAdditionalAmr

`func (o *IdentityProviderProperties) GetAdditionalAmr() []string`

GetAdditionalAmr returns the AdditionalAmr field if non-nil, zero value otherwise.

### GetAdditionalAmrOk

`func (o *IdentityProviderProperties) GetAdditionalAmrOk() (*[]string, bool)`

GetAdditionalAmrOk returns a tuple with the AdditionalAmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAmr

`func (o *IdentityProviderProperties) SetAdditionalAmr(v []string)`

SetAdditionalAmr sets AdditionalAmr field to given value.

### HasAdditionalAmr

`func (o *IdentityProviderProperties) HasAdditionalAmr() bool`

HasAdditionalAmr returns a boolean if a field has been set.

### SetAdditionalAmrNil

`func (o *IdentityProviderProperties) SetAdditionalAmrNil(b bool)`

 SetAdditionalAmrNil sets the value for AdditionalAmr to be an explicit nil

### UnsetAdditionalAmr
`func (o *IdentityProviderProperties) UnsetAdditionalAmr()`

UnsetAdditionalAmr ensures that no value is present for AdditionalAmr, not even an explicit nil
### GetAllowDynamicUserMatching

`func (o *IdentityProviderProperties) GetAllowDynamicUserMatching() bool`

GetAllowDynamicUserMatching returns the AllowDynamicUserMatching field if non-nil, zero value otherwise.

### GetAllowDynamicUserMatchingOk

`func (o *IdentityProviderProperties) GetAllowDynamicUserMatchingOk() (*bool, bool)`

GetAllowDynamicUserMatchingOk returns a tuple with the AllowDynamicUserMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDynamicUserMatching

`func (o *IdentityProviderProperties) SetAllowDynamicUserMatching(v bool)`

SetAllowDynamicUserMatching sets AllowDynamicUserMatching field to given value.

### HasAllowDynamicUserMatching

`func (o *IdentityProviderProperties) HasAllowDynamicUserMatching() bool`

HasAllowDynamicUserMatching returns a boolean if a field has been set.

### GetAllowUserUpdates

`func (o *IdentityProviderProperties) GetAllowUserUpdates() bool`

GetAllowUserUpdates returns the AllowUserUpdates field if non-nil, zero value otherwise.

### GetAllowUserUpdatesOk

`func (o *IdentityProviderProperties) GetAllowUserUpdatesOk() (*bool, bool)`

GetAllowUserUpdatesOk returns a tuple with the AllowUserUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserUpdates

`func (o *IdentityProviderProperties) SetAllowUserUpdates(v bool)`

SetAllowUserUpdates sets AllowUserUpdates field to given value.

### HasAllowUserUpdates

`func (o *IdentityProviderProperties) HasAllowUserUpdates() bool`

HasAllowUserUpdates returns a boolean if a field has been set.

### GetIalValue

`func (o *IdentityProviderProperties) GetIalValue() string`

GetIalValue returns the IalValue field if non-nil, zero value otherwise.

### GetIalValueOk

`func (o *IdentityProviderProperties) GetIalValueOk() (*string, bool)`

GetIalValueOk returns a tuple with the IalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIalValue

`func (o *IdentityProviderProperties) SetIalValue(v string)`

SetIalValue sets IalValue field to given value.

### HasIalValue

`func (o *IdentityProviderProperties) HasIalValue() bool`

HasIalValue returns a boolean if a field has been set.

### SetIalValueNil

`func (o *IdentityProviderProperties) SetIalValueNil(b bool)`

 SetIalValueNil sets the value for IalValue to be an explicit nil

### UnsetIalValue
`func (o *IdentityProviderProperties) UnsetIalValue()`

UnsetIalValue ensures that no value is present for IalValue, not even an explicit nil
### GetIdvMetadata

`func (o *IdentityProviderProperties) GetIdvMetadata() IdentityProviderPropertiesIdvMetadata`

GetIdvMetadata returns the IdvMetadata field if non-nil, zero value otherwise.

### GetIdvMetadataOk

`func (o *IdentityProviderProperties) GetIdvMetadataOk() (*IdentityProviderPropertiesIdvMetadata, bool)`

GetIdvMetadataOk returns a tuple with the IdvMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdvMetadata

`func (o *IdentityProviderProperties) SetIdvMetadata(v IdentityProviderPropertiesIdvMetadata)`

SetIdvMetadata sets IdvMetadata field to given value.

### HasIdvMetadata

`func (o *IdentityProviderProperties) HasIdvMetadata() bool`

HasIdvMetadata returns a boolean if a field has been set.

### GetInquiryTemplateId

`func (o *IdentityProviderProperties) GetInquiryTemplateId() string`

GetInquiryTemplateId returns the InquiryTemplateId field if non-nil, zero value otherwise.

### GetInquiryTemplateIdOk

`func (o *IdentityProviderProperties) GetInquiryTemplateIdOk() (*string, bool)`

GetInquiryTemplateIdOk returns a tuple with the InquiryTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInquiryTemplateId

`func (o *IdentityProviderProperties) SetInquiryTemplateId(v string)`

SetInquiryTemplateId sets InquiryTemplateId field to given value.

### HasInquiryTemplateId

`func (o *IdentityProviderProperties) HasInquiryTemplateId() bool`

HasInquiryTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


