# DNSRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | Pointer to **string** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to [**DNSRecordType**](DNSRecordType.md) |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDNSRecord

`func NewDNSRecord() *DNSRecord`

NewDNSRecord instantiates a new DNSRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordWithDefaults

`func NewDNSRecordWithDefaults() *DNSRecord`

NewDNSRecordWithDefaults instantiates a new DNSRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiration

`func (o *DNSRecord) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *DNSRecord) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *DNSRecord) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *DNSRecord) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetFqdn

`func (o *DNSRecord) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *DNSRecord) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *DNSRecord) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *DNSRecord) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetRecordType

`func (o *DNSRecord) GetRecordType() DNSRecordType`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *DNSRecord) GetRecordTypeOk() (*DNSRecordType, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *DNSRecord) SetRecordType(v DNSRecordType)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *DNSRecord) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetValues

`func (o *DNSRecord) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *DNSRecord) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *DNSRecord) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *DNSRecord) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


