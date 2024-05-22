# Office365ApplicationSettingsApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain for your Office 365 account | 
**Domains** | Pointer to [**[]Office365Domain**](Office365Domain.md) | List of Office 365 domains | [optional] 
**MsftTenant** | **string** | Microsoft tenant name | 

## Methods

### NewOffice365ApplicationSettingsApplication

`func NewOffice365ApplicationSettingsApplication(domain string, msftTenant string, ) *Office365ApplicationSettingsApplication`

NewOffice365ApplicationSettingsApplication instantiates a new Office365ApplicationSettingsApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365ApplicationSettingsApplicationWithDefaults

`func NewOffice365ApplicationSettingsApplicationWithDefaults() *Office365ApplicationSettingsApplication`

NewOffice365ApplicationSettingsApplicationWithDefaults instantiates a new Office365ApplicationSettingsApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *Office365ApplicationSettingsApplication) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Office365ApplicationSettingsApplication) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Office365ApplicationSettingsApplication) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDomains

`func (o *Office365ApplicationSettingsApplication) GetDomains() []Office365Domain`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *Office365ApplicationSettingsApplication) GetDomainsOk() (*[]Office365Domain, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *Office365ApplicationSettingsApplication) SetDomains(v []Office365Domain)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *Office365ApplicationSettingsApplication) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetMsftTenant

`func (o *Office365ApplicationSettingsApplication) GetMsftTenant() string`

GetMsftTenant returns the MsftTenant field if non-nil, zero value otherwise.

### GetMsftTenantOk

`func (o *Office365ApplicationSettingsApplication) GetMsftTenantOk() (*string, bool)`

GetMsftTenantOk returns a tuple with the MsftTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsftTenant

`func (o *Office365ApplicationSettingsApplication) SetMsftTenant(v string)`

SetMsftTenant sets MsftTenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


