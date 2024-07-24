# EmailDomainResponseWithEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**DnsValidationRecords** | Pointer to [**[]EmailDomainDNSRecord**](EmailDomainDNSRecord.md) |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ValidationStatus** | Pointer to **string** |  | [optional] 
**ValidationSubdomain** | Pointer to **string** | The subdomain for the email sender&#39;s custom mail domain | [optional] [default to "mail"]
**Embedded** | Pointer to [**EmailDomainResponseWithEmbeddedEmbedded**](EmailDomainResponseWithEmbeddedEmbedded.md) |  | [optional] 

## Methods

### NewEmailDomainResponseWithEmbedded

`func NewEmailDomainResponseWithEmbedded() *EmailDomainResponseWithEmbedded`

NewEmailDomainResponseWithEmbedded instantiates a new EmailDomainResponseWithEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDomainResponseWithEmbeddedWithDefaults

`func NewEmailDomainResponseWithEmbeddedWithDefaults() *EmailDomainResponseWithEmbedded`

NewEmailDomainResponseWithEmbeddedWithDefaults instantiates a new EmailDomainResponseWithEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *EmailDomainResponseWithEmbedded) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EmailDomainResponseWithEmbedded) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EmailDomainResponseWithEmbedded) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EmailDomainResponseWithEmbedded) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUserName

`func (o *EmailDomainResponseWithEmbedded) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *EmailDomainResponseWithEmbedded) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *EmailDomainResponseWithEmbedded) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *EmailDomainResponseWithEmbedded) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetDnsValidationRecords

`func (o *EmailDomainResponseWithEmbedded) GetDnsValidationRecords() []EmailDomainDNSRecord`

GetDnsValidationRecords returns the DnsValidationRecords field if non-nil, zero value otherwise.

### GetDnsValidationRecordsOk

`func (o *EmailDomainResponseWithEmbedded) GetDnsValidationRecordsOk() (*[]EmailDomainDNSRecord, bool)`

GetDnsValidationRecordsOk returns a tuple with the DnsValidationRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsValidationRecords

`func (o *EmailDomainResponseWithEmbedded) SetDnsValidationRecords(v []EmailDomainDNSRecord)`

SetDnsValidationRecords sets DnsValidationRecords field to given value.

### HasDnsValidationRecords

`func (o *EmailDomainResponseWithEmbedded) HasDnsValidationRecords() bool`

HasDnsValidationRecords returns a boolean if a field has been set.

### GetDomain

`func (o *EmailDomainResponseWithEmbedded) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EmailDomainResponseWithEmbedded) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EmailDomainResponseWithEmbedded) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *EmailDomainResponseWithEmbedded) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetId

`func (o *EmailDomainResponseWithEmbedded) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailDomainResponseWithEmbedded) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailDomainResponseWithEmbedded) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailDomainResponseWithEmbedded) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValidationStatus

`func (o *EmailDomainResponseWithEmbedded) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *EmailDomainResponseWithEmbedded) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *EmailDomainResponseWithEmbedded) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *EmailDomainResponseWithEmbedded) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.

### GetValidationSubdomain

`func (o *EmailDomainResponseWithEmbedded) GetValidationSubdomain() string`

GetValidationSubdomain returns the ValidationSubdomain field if non-nil, zero value otherwise.

### GetValidationSubdomainOk

`func (o *EmailDomainResponseWithEmbedded) GetValidationSubdomainOk() (*string, bool)`

GetValidationSubdomainOk returns a tuple with the ValidationSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationSubdomain

`func (o *EmailDomainResponseWithEmbedded) SetValidationSubdomain(v string)`

SetValidationSubdomain sets ValidationSubdomain field to given value.

### HasValidationSubdomain

`func (o *EmailDomainResponseWithEmbedded) HasValidationSubdomain() bool`

HasValidationSubdomain returns a boolean if a field has been set.

### GetEmbedded

`func (o *EmailDomainResponseWithEmbedded) GetEmbedded() EmailDomainResponseWithEmbeddedEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *EmailDomainResponseWithEmbedded) GetEmbeddedOk() (*EmailDomainResponseWithEmbeddedEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *EmailDomainResponseWithEmbedded) SetEmbedded(v EmailDomainResponseWithEmbeddedEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *EmailDomainResponseWithEmbedded) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


