# Saml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acs** | [**[]SamlAcsInner**](SamlAcsInner.md) | List of Assertion Consumer Service (ACS) URLs. The default ACS URL is required and is indicated by a null &#x60;index&#x60; value. You can use the org-level variables you defined in the &#x60;config&#x60; array in the URL. For example: &#x60;https://${org.subdomain}.example.com/saml/login&#x60; | 
**Doc** | **string** | The URL to your customer-facing instructions for configuring your SAML integration. See [Customer configuration document guidelines](https://developer.okta.com/docs/guides/submit-app-prereq/main/#customer-configuration-document-guidelines). | 
**EntityId** | **string** | Globally unique name for your SAML entity. For instance, your Identity Provider (IdP) or Service Provider (SP) URL. | 

## Methods

### NewSaml

`func NewSaml(acs []SamlAcsInner, doc string, entityId string, ) *Saml`

NewSaml instantiates a new Saml object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlWithDefaults

`func NewSamlWithDefaults() *Saml`

NewSamlWithDefaults instantiates a new Saml object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcs

`func (o *Saml) GetAcs() []SamlAcsInner`

GetAcs returns the Acs field if non-nil, zero value otherwise.

### GetAcsOk

`func (o *Saml) GetAcsOk() (*[]SamlAcsInner, bool)`

GetAcsOk returns a tuple with the Acs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcs

`func (o *Saml) SetAcs(v []SamlAcsInner)`

SetAcs sets Acs field to given value.


### GetDoc

`func (o *Saml) GetDoc() string`

GetDoc returns the Doc field if non-nil, zero value otherwise.

### GetDocOk

`func (o *Saml) GetDocOk() (*string, bool)`

GetDocOk returns a tuple with the Doc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoc

`func (o *Saml) SetDoc(v string)`

SetDoc sets Doc field to given value.


### GetEntityId

`func (o *Saml) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Saml) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Saml) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


