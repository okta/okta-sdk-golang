# UserFactorSupported

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enrollment** | Pointer to **string** | Indicates if the Factor is required for the specified user | [optional] 
**FactorType** | Pointer to **string** | Type of Factor | [optional] 
**Provider** | Pointer to **string** | Provider for the Factor | [optional] 
**Status** | Pointer to **string** | Status of the Factor | [optional] [readonly] 
**VendorName** | Pointer to **string** | Name of the Factor vendor. This is usually the same as the provider except for On-Prem MFA where it depends on administrator settings. | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** | Embedded resources related to the Factor | [optional] [readonly] 
**Links** | Pointer to [**UserFactorLinks**](UserFactorLinks.md) |  | [optional] 

## Methods

### NewUserFactorSupported

`func NewUserFactorSupported() *UserFactorSupported`

NewUserFactorSupported instantiates a new UserFactorSupported object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorSupportedWithDefaults

`func NewUserFactorSupportedWithDefaults() *UserFactorSupported`

NewUserFactorSupportedWithDefaults instantiates a new UserFactorSupported object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnrollment

`func (o *UserFactorSupported) GetEnrollment() string`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *UserFactorSupported) GetEnrollmentOk() (*string, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *UserFactorSupported) SetEnrollment(v string)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *UserFactorSupported) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetFactorType

`func (o *UserFactorSupported) GetFactorType() string`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *UserFactorSupported) GetFactorTypeOk() (*string, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *UserFactorSupported) SetFactorType(v string)`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *UserFactorSupported) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### GetProvider

`func (o *UserFactorSupported) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserFactorSupported) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserFactorSupported) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserFactorSupported) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *UserFactorSupported) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserFactorSupported) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserFactorSupported) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserFactorSupported) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVendorName

`func (o *UserFactorSupported) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *UserFactorSupported) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *UserFactorSupported) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *UserFactorSupported) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetEmbedded

`func (o *UserFactorSupported) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *UserFactorSupported) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *UserFactorSupported) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *UserFactorSupported) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *UserFactorSupported) GetLinks() UserFactorLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserFactorSupported) GetLinksOk() (*UserFactorLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserFactorSupported) SetLinks(v UserFactorLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserFactorSupported) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


