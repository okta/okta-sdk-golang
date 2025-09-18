# SamlSsoEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binding** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **string** | URI reference that indicates the address to which the &#x60;&lt;AuthnRequest&gt;&#x60; message is sent. The &#x60;destination&#x60; property is required if request signatures are specified. See [SAML 2.0 Request Algorithm object](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProvider/#tag/IdentityProvider/operation/createIdentityProvider!path&#x3D;protocol/0/algorithms/request&amp;t&#x3D;request). | [optional] 
**Url** | Pointer to **string** | URL of the binding-specific endpoint to send an &#x60;&lt;AuthnRequest&gt;&#x60; message to the IdP. The value of &#x60;url&#x60; defaults to the same value as the &#x60;sso&#x60; endpoint if omitted during creation of a new IdP instance. The &#x60;url&#x60; should be the same value as the &#x60;Location&#x60; attribute for a published binding in the IdP&#39;s SAML Metadata &#x60;IDPSSODescriptor&#x60;. | [optional] 

## Methods

### NewSamlSsoEndpoint

`func NewSamlSsoEndpoint() *SamlSsoEndpoint`

NewSamlSsoEndpoint instantiates a new SamlSsoEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlSsoEndpointWithDefaults

`func NewSamlSsoEndpointWithDefaults() *SamlSsoEndpoint`

NewSamlSsoEndpointWithDefaults instantiates a new SamlSsoEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinding

`func (o *SamlSsoEndpoint) GetBinding() string`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *SamlSsoEndpoint) GetBindingOk() (*string, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *SamlSsoEndpoint) SetBinding(v string)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *SamlSsoEndpoint) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetDestination

`func (o *SamlSsoEndpoint) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *SamlSsoEndpoint) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *SamlSsoEndpoint) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *SamlSsoEndpoint) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetUrl

`func (o *SamlSsoEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SamlSsoEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SamlSsoEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SamlSsoEndpoint) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


