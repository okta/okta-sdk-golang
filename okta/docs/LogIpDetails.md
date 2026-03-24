# LogIpDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsNumber** | Pointer to **NullableInt32** | The [Autonomous system](https://docs.telemetry.mozilla.org/datasets/other/asn_aggregates/reference) number that&#39;s associated with the IP address | [optional] [readonly] 
**AsOrg** | Pointer to **NullableString** | The name associated with the Autonomous System Number (ASN) | [optional] [readonly] 
**Domain** | Pointer to **NullableString** | The domain name associated with the IP address | [optional] [readonly] 
**IpServiceCategories** | Pointer to [**[]LogIpServiceCategory**](LogIpServiceCategory.md) | The associated IP service categories for the IP address | [optional] [readonly] 
**Isp** | Pointer to **NullableString** | The internet service provider associated with the IP address | [optional] [readonly] 

## Methods

### NewLogIpDetails

`func NewLogIpDetails() *LogIpDetails`

NewLogIpDetails instantiates a new LogIpDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogIpDetailsWithDefaults

`func NewLogIpDetailsWithDefaults() *LogIpDetails`

NewLogIpDetailsWithDefaults instantiates a new LogIpDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsNumber

`func (o *LogIpDetails) GetAsNumber() int32`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *LogIpDetails) GetAsNumberOk() (*int32, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *LogIpDetails) SetAsNumber(v int32)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *LogIpDetails) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### SetAsNumberNil

`func (o *LogIpDetails) SetAsNumberNil(b bool)`

 SetAsNumberNil sets the value for AsNumber to be an explicit nil

### UnsetAsNumber
`func (o *LogIpDetails) UnsetAsNumber()`

UnsetAsNumber ensures that no value is present for AsNumber, not even an explicit nil
### GetAsOrg

`func (o *LogIpDetails) GetAsOrg() string`

GetAsOrg returns the AsOrg field if non-nil, zero value otherwise.

### GetAsOrgOk

`func (o *LogIpDetails) GetAsOrgOk() (*string, bool)`

GetAsOrgOk returns a tuple with the AsOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOrg

`func (o *LogIpDetails) SetAsOrg(v string)`

SetAsOrg sets AsOrg field to given value.

### HasAsOrg

`func (o *LogIpDetails) HasAsOrg() bool`

HasAsOrg returns a boolean if a field has been set.

### SetAsOrgNil

`func (o *LogIpDetails) SetAsOrgNil(b bool)`

 SetAsOrgNil sets the value for AsOrg to be an explicit nil

### UnsetAsOrg
`func (o *LogIpDetails) UnsetAsOrg()`

UnsetAsOrg ensures that no value is present for AsOrg, not even an explicit nil
### GetDomain

`func (o *LogIpDetails) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LogIpDetails) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LogIpDetails) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LogIpDetails) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *LogIpDetails) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *LogIpDetails) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetIpServiceCategories

`func (o *LogIpDetails) GetIpServiceCategories() []LogIpServiceCategory`

GetIpServiceCategories returns the IpServiceCategories field if non-nil, zero value otherwise.

### GetIpServiceCategoriesOk

`func (o *LogIpDetails) GetIpServiceCategoriesOk() (*[]LogIpServiceCategory, bool)`

GetIpServiceCategoriesOk returns a tuple with the IpServiceCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpServiceCategories

`func (o *LogIpDetails) SetIpServiceCategories(v []LogIpServiceCategory)`

SetIpServiceCategories sets IpServiceCategories field to given value.

### HasIpServiceCategories

`func (o *LogIpDetails) HasIpServiceCategories() bool`

HasIpServiceCategories returns a boolean if a field has been set.

### SetIpServiceCategoriesNil

`func (o *LogIpDetails) SetIpServiceCategoriesNil(b bool)`

 SetIpServiceCategoriesNil sets the value for IpServiceCategories to be an explicit nil

### UnsetIpServiceCategories
`func (o *LogIpDetails) UnsetIpServiceCategories()`

UnsetIpServiceCategories ensures that no value is present for IpServiceCategories, not even an explicit nil
### GetIsp

`func (o *LogIpDetails) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *LogIpDetails) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *LogIpDetails) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *LogIpDetails) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### SetIspNil

`func (o *LogIpDetails) SetIspNil(b bool)`

 SetIspNil sets the value for Isp to be an explicit nil

### UnsetIsp
`func (o *LogIpDetails) UnsetIsp()`

UnsetIsp ensures that no value is present for Isp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


