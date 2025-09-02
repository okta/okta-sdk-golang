# SamlSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HonorPersistentNameId** | Pointer to **bool** | Determines if the IdP should persist account linking when the incoming assertion NameID format is &#x60;urn:oasis:names:tc:SAML:2.0:nameid-format:persistent&#x60; | [optional] [default to true]
**NameFormat** | Pointer to **string** | SAML 2.0 Name Identifier formats | [optional] [default to "urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"]
**ParticipateSlo** | Pointer to **bool** | Set to &#x60;true&#x60; to have Okta send a logout request to the upstream IdP when a user signs out of Okta or a downstream app. | [optional] 
**SendApplicationContext** | Pointer to **bool** | Determines if the IdP should send the application context as &#x60;&lt;OktaAppInstanceId&gt;&#x60; and &#x60;&lt;OktaAppName&gt;&#x60; in the &#x60;&lt;saml2p:Extensions&gt;&#x60; element of the &#x60;&lt;AuthnRequest&gt;&#x60; message | [optional] [default to false]

## Methods

### NewSamlSettings

`func NewSamlSettings() *SamlSettings`

NewSamlSettings instantiates a new SamlSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlSettingsWithDefaults

`func NewSamlSettingsWithDefaults() *SamlSettings`

NewSamlSettingsWithDefaults instantiates a new SamlSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHonorPersistentNameId

`func (o *SamlSettings) GetHonorPersistentNameId() bool`

GetHonorPersistentNameId returns the HonorPersistentNameId field if non-nil, zero value otherwise.

### GetHonorPersistentNameIdOk

`func (o *SamlSettings) GetHonorPersistentNameIdOk() (*bool, bool)`

GetHonorPersistentNameIdOk returns a tuple with the HonorPersistentNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorPersistentNameId

`func (o *SamlSettings) SetHonorPersistentNameId(v bool)`

SetHonorPersistentNameId sets HonorPersistentNameId field to given value.

### HasHonorPersistentNameId

`func (o *SamlSettings) HasHonorPersistentNameId() bool`

HasHonorPersistentNameId returns a boolean if a field has been set.

### GetNameFormat

`func (o *SamlSettings) GetNameFormat() string`

GetNameFormat returns the NameFormat field if non-nil, zero value otherwise.

### GetNameFormatOk

`func (o *SamlSettings) GetNameFormatOk() (*string, bool)`

GetNameFormatOk returns a tuple with the NameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFormat

`func (o *SamlSettings) SetNameFormat(v string)`

SetNameFormat sets NameFormat field to given value.

### HasNameFormat

`func (o *SamlSettings) HasNameFormat() bool`

HasNameFormat returns a boolean if a field has been set.

### GetParticipateSlo

`func (o *SamlSettings) GetParticipateSlo() bool`

GetParticipateSlo returns the ParticipateSlo field if non-nil, zero value otherwise.

### GetParticipateSloOk

`func (o *SamlSettings) GetParticipateSloOk() (*bool, bool)`

GetParticipateSloOk returns a tuple with the ParticipateSlo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipateSlo

`func (o *SamlSettings) SetParticipateSlo(v bool)`

SetParticipateSlo sets ParticipateSlo field to given value.

### HasParticipateSlo

`func (o *SamlSettings) HasParticipateSlo() bool`

HasParticipateSlo returns a boolean if a field has been set.

### GetSendApplicationContext

`func (o *SamlSettings) GetSendApplicationContext() bool`

GetSendApplicationContext returns the SendApplicationContext field if non-nil, zero value otherwise.

### GetSendApplicationContextOk

`func (o *SamlSettings) GetSendApplicationContextOk() (*bool, bool)`

GetSendApplicationContextOk returns a tuple with the SendApplicationContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendApplicationContext

`func (o *SamlSettings) SetSendApplicationContext(v bool)`

SetSendApplicationContext sets SendApplicationContext field to given value.

### HasSendApplicationContext

`func (o *SamlSettings) HasSendApplicationContext() bool`

HasSendApplicationContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


