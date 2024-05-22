# EmailDomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsValidationRecords** | Pointer to [**[]EmailDomainDNSRecord**](EmailDomainDNSRecord.md) |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ValidationStatus** | Pointer to **string** |  | [optional] 
**ValidationSubdomain** | Pointer to **string** | The subdomain for the email sender&#39;s custom mail domain | [optional] [default to "mail"]
**DisplayName** | **string** |  | 
**UserName** | **string** |  | 

## Methods

### NewEmailDomainResponse

`func NewEmailDomainResponse(displayName string, userName string, ) *EmailDomainResponse`

NewEmailDomainResponse instantiates a new EmailDomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDomainResponseWithDefaults

`func NewEmailDomainResponseWithDefaults() *EmailDomainResponse`

NewEmailDomainResponseWithDefaults instantiates a new EmailDomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsValidationRecords

`func (o *EmailDomainResponse) GetDnsValidationRecords() []EmailDomainDNSRecord`

GetDnsValidationRecords returns the DnsValidationRecords field if non-nil, zero value otherwise.

### GetDnsValidationRecordsOk

`func (o *EmailDomainResponse) GetDnsValidationRecordsOk() (*[]EmailDomainDNSRecord, bool)`

GetDnsValidationRecordsOk returns a tuple with the DnsValidationRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsValidationRecords

`func (o *EmailDomainResponse) SetDnsValidationRecords(v []EmailDomainDNSRecord)`

SetDnsValidationRecords sets DnsValidationRecords field to given value.

### HasDnsValidationRecords

`func (o *EmailDomainResponse) HasDnsValidationRecords() bool`

HasDnsValidationRecords returns a boolean if a field has been set.

### GetDomain

`func (o *EmailDomainResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EmailDomainResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EmailDomainResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *EmailDomainResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetId

`func (o *EmailDomainResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailDomainResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailDomainResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailDomainResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValidationStatus

`func (o *EmailDomainResponse) GetValidationStatus() string`

GetValidationStatus returns the ValidationStatus field if non-nil, zero value otherwise.

### GetValidationStatusOk

`func (o *EmailDomainResponse) GetValidationStatusOk() (*string, bool)`

GetValidationStatusOk returns a tuple with the ValidationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationStatus

`func (o *EmailDomainResponse) SetValidationStatus(v string)`

SetValidationStatus sets ValidationStatus field to given value.

### HasValidationStatus

`func (o *EmailDomainResponse) HasValidationStatus() bool`

HasValidationStatus returns a boolean if a field has been set.

### GetValidationSubdomain

`func (o *EmailDomainResponse) GetValidationSubdomain() string`

GetValidationSubdomain returns the ValidationSubdomain field if non-nil, zero value otherwise.

### GetValidationSubdomainOk

`func (o *EmailDomainResponse) GetValidationSubdomainOk() (*string, bool)`

GetValidationSubdomainOk returns a tuple with the ValidationSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationSubdomain

`func (o *EmailDomainResponse) SetValidationSubdomain(v string)`

SetValidationSubdomain sets ValidationSubdomain field to given value.

### HasValidationSubdomain

`func (o *EmailDomainResponse) HasValidationSubdomain() bool`

HasValidationSubdomain returns a boolean if a field has been set.

### GetDisplayName

`func (o *EmailDomainResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EmailDomainResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EmailDomainResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetUserName

`func (o *EmailDomainResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *EmailDomainResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *EmailDomainResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


