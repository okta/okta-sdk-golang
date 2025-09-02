# WebAuthnRpIdDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsRecord** | Pointer to [**DNSRecord**](DNSRecord.md) |  | [optional] [readonly] 
**Name** | Pointer to **string** | The [RP ID](https://www.w3.org/TR/webauthn/#relying-party-identifier) domain value to be used for all WebAuthn operations.  If it isn&#39;t specified, the &#x60;domain&#x60; object isn&#39;t included in the request, and the domain value defaults to the domain of the current page (the domain of your org or a custom domain, for example).  &gt; **Note:** If you don&#39;t use a custom RP ID (the default behavior), the domain value defaults to the end user&#39;s current page. The domain value defaults to the full domain name of the page that the end user is on when they&#39;re attempting the WebAuthn credential operation (enrollment or verification). | [optional] 
**ValidationStatus** | Pointer to **string** | Indicates the validation status of the domain | [optional] [readonly] 

## Methods

### NewWebAuthnRpIdDomain

`func NewWebAuthnRpIdDomain() *WebAuthnRpIdDomain`

NewWebAuthnRpIdDomain instantiates a new WebAuthnRpIdDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebAuthnRpIdDomainWithDefaults

`func NewWebAuthnRpIdDomainWithDefaults() *WebAuthnRpIdDomain`

NewWebAuthnRpIdDomainWithDefaults instantiates a new WebAuthnRpIdDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsRecord

`func (o *WebAuthnRpIdDomain) GetDnsRecord() DNSRecord`

GetDnsRecord returns the DnsRecord field if non-nil, zero value otherwise.

### GetDnsRecordOk

`func (o *WebAuthnRpIdDomain) GetDnsRecordOk() (*DNSRecord, bool)`

GetDnsRecordOk returns a tuple with the DnsRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecord

`func (o *WebAuthnRpIdDomain) SetDnsRecord(v DNSRecord)`

SetDnsRecord sets DnsRecord field to given value.

### HasDnsRecord

`func (o *WebAuthnRpIdDomain) HasDnsRecord() bool`

HasDnsRecord returns a boolean if a field has been set.

### GetName

`func (o *WebAuthnRpIdDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebAuthnRpIdDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebAuthnRpIdDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebAuthnRpIdDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValidationStatus

`func (o *WebAuthnRpIdDomain) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *WebAuthnRpIdDomain) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *WebAuthnRpIdDomain) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *WebAuthnRpIdDomain) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


