# DNSRecordDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | Pointer to **string** | DNS TXT record expiration | [optional] 
**Fqdn** | Pointer to **string** | DNS record name | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**Values** | Pointer to **[]string** | DNS record value | [optional] 

## Methods

### NewDNSRecordDomains

`func NewDNSRecordDomains() *DNSRecordDomains`

NewDNSRecordDomains instantiates a new DNSRecordDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordDomainsWithDefaults

`func NewDNSRecordDomainsWithDefaults() *DNSRecordDomains`

NewDNSRecordDomainsWithDefaults instantiates a new DNSRecordDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiration

`func (o *DNSRecordDomains) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *DNSRecordDomains) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *DNSRecordDomains) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *DNSRecordDomains) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetFqdn

`func (o *DNSRecordDomains) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *DNSRecordDomains) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *DNSRecordDomains) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *DNSRecordDomains) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetRecordType

`func (o *DNSRecordDomains) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *DNSRecordDomains) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *DNSRecordDomains) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *DNSRecordDomains) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetValues

`func (o *DNSRecordDomains) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *DNSRecordDomains) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *DNSRecordDomains) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *DNSRecordDomains) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


