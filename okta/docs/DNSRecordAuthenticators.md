# DNSRecordAuthenticators

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** | The DNS record name | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**VerificationValue** | Pointer to **string** | The DNS record verification value | [optional] 

## Methods

### NewDNSRecordAuthenticators

`func NewDNSRecordAuthenticators() *DNSRecordAuthenticators`

NewDNSRecordAuthenticators instantiates a new DNSRecordAuthenticators object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordAuthenticatorsWithDefaults

`func NewDNSRecordAuthenticatorsWithDefaults() *DNSRecordAuthenticators`

NewDNSRecordAuthenticatorsWithDefaults instantiates a new DNSRecordAuthenticators object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *DNSRecordAuthenticators) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *DNSRecordAuthenticators) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *DNSRecordAuthenticators) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *DNSRecordAuthenticators) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetRecordType

`func (o *DNSRecordAuthenticators) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *DNSRecordAuthenticators) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *DNSRecordAuthenticators) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *DNSRecordAuthenticators) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetVerificationValue

`func (o *DNSRecordAuthenticators) GetVerificationValue() string`

GetVerificationValue returns the VerificationValue field if non-nil, zero value otherwise.

### GetVerificationValueOk

`func (o *DNSRecordAuthenticators) GetVerificationValueOk() (*string, bool)`

GetVerificationValueOk returns a tuple with the VerificationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationValue

`func (o *DNSRecordAuthenticators) SetVerificationValue(v string)`

SetVerificationValue sets VerificationValue field to given value.

### HasVerificationValue

`func (o *DNSRecordAuthenticators) HasVerificationValue() bool`

HasVerificationValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


