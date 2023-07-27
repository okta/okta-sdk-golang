# SamlApplicationSettingsSignOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcsEndpoints** | Pointer to [**[]AcsEndpoint**](AcsEndpoint.md) |  | [optional] 
**AllowMultipleAcsEndpoints** | Pointer to **bool** |  | [optional] 
**AssertionSigned** | Pointer to **bool** |  | [optional] 
**AttributeStatements** | Pointer to [**[]SamlAttributeStatement**](SamlAttributeStatement.md) |  | [optional] 
**Audience** | Pointer to **string** |  | [optional] 
**AudienceOverride** | Pointer to **string** |  | [optional] 
**AuthnContextClassRef** | Pointer to **string** |  | [optional] 
**DefaultRelayState** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **string** |  | [optional] 
**DestinationOverride** | Pointer to **string** |  | [optional] 
**DigestAlgorithm** | Pointer to **string** |  | [optional] 
**HonorForceAuthn** | Pointer to **bool** |  | [optional] 
**IdpIssuer** | Pointer to **string** |  | [optional] 
**InlineHooks** | Pointer to [**[]SignOnInlineHook**](SignOnInlineHook.md) |  | [optional] 
**Recipient** | Pointer to **string** |  | [optional] 
**RecipientOverride** | Pointer to **string** |  | [optional] 
**RequestCompressed** | Pointer to **bool** |  | [optional] 
**ResponseSigned** | Pointer to **bool** |  | [optional] 
**SignatureAlgorithm** | Pointer to **string** |  | [optional] 
**Slo** | Pointer to [**SingleLogout**](SingleLogout.md) |  | [optional] 
**SpCertificate** | Pointer to [**SpCertificate**](SpCertificate.md) |  | [optional] 
**SpIssuer** | Pointer to **string** |  | [optional] 
**SsoAcsUrl** | Pointer to **string** |  | [optional] 
**SsoAcsUrlOverride** | Pointer to **string** |  | [optional] 
**SubjectNameIdFormat** | Pointer to **string** |  | [optional] 
**SubjectNameIdTemplate** | Pointer to **string** |  | [optional] 

## Methods

### NewSamlApplicationSettingsSignOn

`func NewSamlApplicationSettingsSignOn() *SamlApplicationSettingsSignOn`

NewSamlApplicationSettingsSignOn instantiates a new SamlApplicationSettingsSignOn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlApplicationSettingsSignOnWithDefaults

`func NewSamlApplicationSettingsSignOnWithDefaults() *SamlApplicationSettingsSignOn`

NewSamlApplicationSettingsSignOnWithDefaults instantiates a new SamlApplicationSettingsSignOn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcsEndpoints

`func (o *SamlApplicationSettingsSignOn) GetAcsEndpoints() []AcsEndpoint`

GetAcsEndpoints returns the AcsEndpoints field if non-nil, zero value otherwise.

### GetAcsEndpointsOk

`func (o *SamlApplicationSettingsSignOn) GetAcsEndpointsOk() (*[]AcsEndpoint, bool)`

GetAcsEndpointsOk returns a tuple with the AcsEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsEndpoints

`func (o *SamlApplicationSettingsSignOn) SetAcsEndpoints(v []AcsEndpoint)`

SetAcsEndpoints sets AcsEndpoints field to given value.

### HasAcsEndpoints

`func (o *SamlApplicationSettingsSignOn) HasAcsEndpoints() bool`

HasAcsEndpoints returns a boolean if a field has been set.

### GetAllowMultipleAcsEndpoints

`func (o *SamlApplicationSettingsSignOn) GetAllowMultipleAcsEndpoints() bool`

GetAllowMultipleAcsEndpoints returns the AllowMultipleAcsEndpoints field if non-nil, zero value otherwise.

### GetAllowMultipleAcsEndpointsOk

`func (o *SamlApplicationSettingsSignOn) GetAllowMultipleAcsEndpointsOk() (*bool, bool)`

GetAllowMultipleAcsEndpointsOk returns a tuple with the AllowMultipleAcsEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleAcsEndpoints

`func (o *SamlApplicationSettingsSignOn) SetAllowMultipleAcsEndpoints(v bool)`

SetAllowMultipleAcsEndpoints sets AllowMultipleAcsEndpoints field to given value.

### HasAllowMultipleAcsEndpoints

`func (o *SamlApplicationSettingsSignOn) HasAllowMultipleAcsEndpoints() bool`

HasAllowMultipleAcsEndpoints returns a boolean if a field has been set.

### GetAssertionSigned

`func (o *SamlApplicationSettingsSignOn) GetAssertionSigned() bool`

GetAssertionSigned returns the AssertionSigned field if non-nil, zero value otherwise.

### GetAssertionSignedOk

`func (o *SamlApplicationSettingsSignOn) GetAssertionSignedOk() (*bool, bool)`

GetAssertionSignedOk returns a tuple with the AssertionSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionSigned

`func (o *SamlApplicationSettingsSignOn) SetAssertionSigned(v bool)`

SetAssertionSigned sets AssertionSigned field to given value.

### HasAssertionSigned

`func (o *SamlApplicationSettingsSignOn) HasAssertionSigned() bool`

HasAssertionSigned returns a boolean if a field has been set.

### GetAttributeStatements

`func (o *SamlApplicationSettingsSignOn) GetAttributeStatements() []SamlAttributeStatement`

GetAttributeStatements returns the AttributeStatements field if non-nil, zero value otherwise.

### GetAttributeStatementsOk

`func (o *SamlApplicationSettingsSignOn) GetAttributeStatementsOk() (*[]SamlAttributeStatement, bool)`

GetAttributeStatementsOk returns a tuple with the AttributeStatements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeStatements

`func (o *SamlApplicationSettingsSignOn) SetAttributeStatements(v []SamlAttributeStatement)`

SetAttributeStatements sets AttributeStatements field to given value.

### HasAttributeStatements

`func (o *SamlApplicationSettingsSignOn) HasAttributeStatements() bool`

HasAttributeStatements returns a boolean if a field has been set.

### GetAudience

`func (o *SamlApplicationSettingsSignOn) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *SamlApplicationSettingsSignOn) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *SamlApplicationSettingsSignOn) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *SamlApplicationSettingsSignOn) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetAudienceOverride

`func (o *SamlApplicationSettingsSignOn) GetAudienceOverride() string`

GetAudienceOverride returns the AudienceOverride field if non-nil, zero value otherwise.

### GetAudienceOverrideOk

`func (o *SamlApplicationSettingsSignOn) GetAudienceOverrideOk() (*string, bool)`

GetAudienceOverrideOk returns a tuple with the AudienceOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceOverride

`func (o *SamlApplicationSettingsSignOn) SetAudienceOverride(v string)`

SetAudienceOverride sets AudienceOverride field to given value.

### HasAudienceOverride

`func (o *SamlApplicationSettingsSignOn) HasAudienceOverride() bool`

HasAudienceOverride returns a boolean if a field has been set.

### GetAuthnContextClassRef

`func (o *SamlApplicationSettingsSignOn) GetAuthnContextClassRef() string`

GetAuthnContextClassRef returns the AuthnContextClassRef field if non-nil, zero value otherwise.

### GetAuthnContextClassRefOk

`func (o *SamlApplicationSettingsSignOn) GetAuthnContextClassRefOk() (*string, bool)`

GetAuthnContextClassRefOk returns a tuple with the AuthnContextClassRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnContextClassRef

`func (o *SamlApplicationSettingsSignOn) SetAuthnContextClassRef(v string)`

SetAuthnContextClassRef sets AuthnContextClassRef field to given value.

### HasAuthnContextClassRef

`func (o *SamlApplicationSettingsSignOn) HasAuthnContextClassRef() bool`

HasAuthnContextClassRef returns a boolean if a field has been set.

### GetDefaultRelayState

`func (o *SamlApplicationSettingsSignOn) GetDefaultRelayState() string`

GetDefaultRelayState returns the DefaultRelayState field if non-nil, zero value otherwise.

### GetDefaultRelayStateOk

`func (o *SamlApplicationSettingsSignOn) GetDefaultRelayStateOk() (*string, bool)`

GetDefaultRelayStateOk returns a tuple with the DefaultRelayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRelayState

`func (o *SamlApplicationSettingsSignOn) SetDefaultRelayState(v string)`

SetDefaultRelayState sets DefaultRelayState field to given value.

### HasDefaultRelayState

`func (o *SamlApplicationSettingsSignOn) HasDefaultRelayState() bool`

HasDefaultRelayState returns a boolean if a field has been set.

### GetDestination

`func (o *SamlApplicationSettingsSignOn) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *SamlApplicationSettingsSignOn) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *SamlApplicationSettingsSignOn) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *SamlApplicationSettingsSignOn) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationOverride

`func (o *SamlApplicationSettingsSignOn) GetDestinationOverride() string`

GetDestinationOverride returns the DestinationOverride field if non-nil, zero value otherwise.

### GetDestinationOverrideOk

`func (o *SamlApplicationSettingsSignOn) GetDestinationOverrideOk() (*string, bool)`

GetDestinationOverrideOk returns a tuple with the DestinationOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationOverride

`func (o *SamlApplicationSettingsSignOn) SetDestinationOverride(v string)`

SetDestinationOverride sets DestinationOverride field to given value.

### HasDestinationOverride

`func (o *SamlApplicationSettingsSignOn) HasDestinationOverride() bool`

HasDestinationOverride returns a boolean if a field has been set.

### GetDigestAlgorithm

`func (o *SamlApplicationSettingsSignOn) GetDigestAlgorithm() string`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *SamlApplicationSettingsSignOn) GetDigestAlgorithmOk() (*string, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *SamlApplicationSettingsSignOn) SetDigestAlgorithm(v string)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *SamlApplicationSettingsSignOn) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetHonorForceAuthn

`func (o *SamlApplicationSettingsSignOn) GetHonorForceAuthn() bool`

GetHonorForceAuthn returns the HonorForceAuthn field if non-nil, zero value otherwise.

### GetHonorForceAuthnOk

`func (o *SamlApplicationSettingsSignOn) GetHonorForceAuthnOk() (*bool, bool)`

GetHonorForceAuthnOk returns a tuple with the HonorForceAuthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorForceAuthn

`func (o *SamlApplicationSettingsSignOn) SetHonorForceAuthn(v bool)`

SetHonorForceAuthn sets HonorForceAuthn field to given value.

### HasHonorForceAuthn

`func (o *SamlApplicationSettingsSignOn) HasHonorForceAuthn() bool`

HasHonorForceAuthn returns a boolean if a field has been set.

### GetIdpIssuer

`func (o *SamlApplicationSettingsSignOn) GetIdpIssuer() string`

GetIdpIssuer returns the IdpIssuer field if non-nil, zero value otherwise.

### GetIdpIssuerOk

`func (o *SamlApplicationSettingsSignOn) GetIdpIssuerOk() (*string, bool)`

GetIdpIssuerOk returns a tuple with the IdpIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpIssuer

`func (o *SamlApplicationSettingsSignOn) SetIdpIssuer(v string)`

SetIdpIssuer sets IdpIssuer field to given value.

### HasIdpIssuer

`func (o *SamlApplicationSettingsSignOn) HasIdpIssuer() bool`

HasIdpIssuer returns a boolean if a field has been set.

### GetInlineHooks

`func (o *SamlApplicationSettingsSignOn) GetInlineHooks() []SignOnInlineHook`

GetInlineHooks returns the InlineHooks field if non-nil, zero value otherwise.

### GetInlineHooksOk

`func (o *SamlApplicationSettingsSignOn) GetInlineHooksOk() (*[]SignOnInlineHook, bool)`

GetInlineHooksOk returns a tuple with the InlineHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineHooks

`func (o *SamlApplicationSettingsSignOn) SetInlineHooks(v []SignOnInlineHook)`

SetInlineHooks sets InlineHooks field to given value.

### HasInlineHooks

`func (o *SamlApplicationSettingsSignOn) HasInlineHooks() bool`

HasInlineHooks returns a boolean if a field has been set.

### GetRecipient

`func (o *SamlApplicationSettingsSignOn) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *SamlApplicationSettingsSignOn) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *SamlApplicationSettingsSignOn) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *SamlApplicationSettingsSignOn) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetRecipientOverride

`func (o *SamlApplicationSettingsSignOn) GetRecipientOverride() string`

GetRecipientOverride returns the RecipientOverride field if non-nil, zero value otherwise.

### GetRecipientOverrideOk

`func (o *SamlApplicationSettingsSignOn) GetRecipientOverrideOk() (*string, bool)`

GetRecipientOverrideOk returns a tuple with the RecipientOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientOverride

`func (o *SamlApplicationSettingsSignOn) SetRecipientOverride(v string)`

SetRecipientOverride sets RecipientOverride field to given value.

### HasRecipientOverride

`func (o *SamlApplicationSettingsSignOn) HasRecipientOverride() bool`

HasRecipientOverride returns a boolean if a field has been set.

### GetRequestCompressed

`func (o *SamlApplicationSettingsSignOn) GetRequestCompressed() bool`

GetRequestCompressed returns the RequestCompressed field if non-nil, zero value otherwise.

### GetRequestCompressedOk

`func (o *SamlApplicationSettingsSignOn) GetRequestCompressedOk() (*bool, bool)`

GetRequestCompressedOk returns a tuple with the RequestCompressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCompressed

`func (o *SamlApplicationSettingsSignOn) SetRequestCompressed(v bool)`

SetRequestCompressed sets RequestCompressed field to given value.

### HasRequestCompressed

`func (o *SamlApplicationSettingsSignOn) HasRequestCompressed() bool`

HasRequestCompressed returns a boolean if a field has been set.

### GetResponseSigned

`func (o *SamlApplicationSettingsSignOn) GetResponseSigned() bool`

GetResponseSigned returns the ResponseSigned field if non-nil, zero value otherwise.

### GetResponseSignedOk

`func (o *SamlApplicationSettingsSignOn) GetResponseSignedOk() (*bool, bool)`

GetResponseSignedOk returns a tuple with the ResponseSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSigned

`func (o *SamlApplicationSettingsSignOn) SetResponseSigned(v bool)`

SetResponseSigned sets ResponseSigned field to given value.

### HasResponseSigned

`func (o *SamlApplicationSettingsSignOn) HasResponseSigned() bool`

HasResponseSigned returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *SamlApplicationSettingsSignOn) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *SamlApplicationSettingsSignOn) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *SamlApplicationSettingsSignOn) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *SamlApplicationSettingsSignOn) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetSlo

`func (o *SamlApplicationSettingsSignOn) GetSlo() SingleLogout`

GetSlo returns the Slo field if non-nil, zero value otherwise.

### GetSloOk

`func (o *SamlApplicationSettingsSignOn) GetSloOk() (*SingleLogout, bool)`

GetSloOk returns a tuple with the Slo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlo

`func (o *SamlApplicationSettingsSignOn) SetSlo(v SingleLogout)`

SetSlo sets Slo field to given value.

### HasSlo

`func (o *SamlApplicationSettingsSignOn) HasSlo() bool`

HasSlo returns a boolean if a field has been set.

### GetSpCertificate

`func (o *SamlApplicationSettingsSignOn) GetSpCertificate() SpCertificate`

GetSpCertificate returns the SpCertificate field if non-nil, zero value otherwise.

### GetSpCertificateOk

`func (o *SamlApplicationSettingsSignOn) GetSpCertificateOk() (*SpCertificate, bool)`

GetSpCertificateOk returns a tuple with the SpCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpCertificate

`func (o *SamlApplicationSettingsSignOn) SetSpCertificate(v SpCertificate)`

SetSpCertificate sets SpCertificate field to given value.

### HasSpCertificate

`func (o *SamlApplicationSettingsSignOn) HasSpCertificate() bool`

HasSpCertificate returns a boolean if a field has been set.

### GetSpIssuer

`func (o *SamlApplicationSettingsSignOn) GetSpIssuer() string`

GetSpIssuer returns the SpIssuer field if non-nil, zero value otherwise.

### GetSpIssuerOk

`func (o *SamlApplicationSettingsSignOn) GetSpIssuerOk() (*string, bool)`

GetSpIssuerOk returns a tuple with the SpIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpIssuer

`func (o *SamlApplicationSettingsSignOn) SetSpIssuer(v string)`

SetSpIssuer sets SpIssuer field to given value.

### HasSpIssuer

`func (o *SamlApplicationSettingsSignOn) HasSpIssuer() bool`

HasSpIssuer returns a boolean if a field has been set.

### GetSsoAcsUrl

`func (o *SamlApplicationSettingsSignOn) GetSsoAcsUrl() string`

GetSsoAcsUrl returns the SsoAcsUrl field if non-nil, zero value otherwise.

### GetSsoAcsUrlOk

`func (o *SamlApplicationSettingsSignOn) GetSsoAcsUrlOk() (*string, bool)`

GetSsoAcsUrlOk returns a tuple with the SsoAcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoAcsUrl

`func (o *SamlApplicationSettingsSignOn) SetSsoAcsUrl(v string)`

SetSsoAcsUrl sets SsoAcsUrl field to given value.

### HasSsoAcsUrl

`func (o *SamlApplicationSettingsSignOn) HasSsoAcsUrl() bool`

HasSsoAcsUrl returns a boolean if a field has been set.

### GetSsoAcsUrlOverride

`func (o *SamlApplicationSettingsSignOn) GetSsoAcsUrlOverride() string`

GetSsoAcsUrlOverride returns the SsoAcsUrlOverride field if non-nil, zero value otherwise.

### GetSsoAcsUrlOverrideOk

`func (o *SamlApplicationSettingsSignOn) GetSsoAcsUrlOverrideOk() (*string, bool)`

GetSsoAcsUrlOverrideOk returns a tuple with the SsoAcsUrlOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoAcsUrlOverride

`func (o *SamlApplicationSettingsSignOn) SetSsoAcsUrlOverride(v string)`

SetSsoAcsUrlOverride sets SsoAcsUrlOverride field to given value.

### HasSsoAcsUrlOverride

`func (o *SamlApplicationSettingsSignOn) HasSsoAcsUrlOverride() bool`

HasSsoAcsUrlOverride returns a boolean if a field has been set.

### GetSubjectNameIdFormat

`func (o *SamlApplicationSettingsSignOn) GetSubjectNameIdFormat() string`

GetSubjectNameIdFormat returns the SubjectNameIdFormat field if non-nil, zero value otherwise.

### GetSubjectNameIdFormatOk

`func (o *SamlApplicationSettingsSignOn) GetSubjectNameIdFormatOk() (*string, bool)`

GetSubjectNameIdFormatOk returns a tuple with the SubjectNameIdFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectNameIdFormat

`func (o *SamlApplicationSettingsSignOn) SetSubjectNameIdFormat(v string)`

SetSubjectNameIdFormat sets SubjectNameIdFormat field to given value.

### HasSubjectNameIdFormat

`func (o *SamlApplicationSettingsSignOn) HasSubjectNameIdFormat() bool`

HasSubjectNameIdFormat returns a boolean if a field has been set.

### GetSubjectNameIdTemplate

`func (o *SamlApplicationSettingsSignOn) GetSubjectNameIdTemplate() string`

GetSubjectNameIdTemplate returns the SubjectNameIdTemplate field if non-nil, zero value otherwise.

### GetSubjectNameIdTemplateOk

`func (o *SamlApplicationSettingsSignOn) GetSubjectNameIdTemplateOk() (*string, bool)`

GetSubjectNameIdTemplateOk returns a tuple with the SubjectNameIdTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectNameIdTemplate

`func (o *SamlApplicationSettingsSignOn) SetSubjectNameIdTemplate(v string)`

SetSubjectNameIdTemplate sets SubjectNameIdTemplate field to given value.

### HasSubjectNameIdTemplate

`func (o *SamlApplicationSettingsSignOn) HasSubjectNameIdTemplate() bool`

HasSubjectNameIdTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


