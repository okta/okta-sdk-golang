# ResendEnrollFactorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the Factor was enrolled | [optional] [readonly] 
**FactorType** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID of the Factor | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the Factor was last updated | [optional] [readonly] 
**Profile** | Pointer to [**UserFactorSMSProfile**](UserFactorSMSProfile.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Status of the Factor | [optional] [readonly] 
**VendorName** | Pointer to **string** | Name of the Factor vendor. This is usually the same as the provider except for On-Prem MFA where it depends on administrator settings. | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResendEnrollFactorRequest

`func NewResendEnrollFactorRequest() *ResendEnrollFactorRequest`

NewResendEnrollFactorRequest instantiates a new ResendEnrollFactorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResendEnrollFactorRequestWithDefaults

`func NewResendEnrollFactorRequestWithDefaults() *ResendEnrollFactorRequest`

NewResendEnrollFactorRequestWithDefaults instantiates a new ResendEnrollFactorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ResendEnrollFactorRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResendEnrollFactorRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResendEnrollFactorRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResendEnrollFactorRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFactorType

`func (o *ResendEnrollFactorRequest) GetFactorType() interface{}`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *ResendEnrollFactorRequest) GetFactorTypeOk() (*interface{}, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *ResendEnrollFactorRequest) SetFactorType(v interface{})`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *ResendEnrollFactorRequest) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### SetFactorTypeNil

`func (o *ResendEnrollFactorRequest) SetFactorTypeNil(b bool)`

 SetFactorTypeNil sets the value for FactorType to be an explicit nil

### UnsetFactorType
`func (o *ResendEnrollFactorRequest) UnsetFactorType()`

UnsetFactorType ensures that no value is present for FactorType, not even an explicit nil
### GetId

`func (o *ResendEnrollFactorRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResendEnrollFactorRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResendEnrollFactorRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResendEnrollFactorRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ResendEnrollFactorRequest) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ResendEnrollFactorRequest) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ResendEnrollFactorRequest) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ResendEnrollFactorRequest) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProfile

`func (o *ResendEnrollFactorRequest) GetProfile() UserFactorSMSProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ResendEnrollFactorRequest) GetProfileOk() (*UserFactorSMSProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ResendEnrollFactorRequest) SetProfile(v UserFactorSMSProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ResendEnrollFactorRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProvider

`func (o *ResendEnrollFactorRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ResendEnrollFactorRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ResendEnrollFactorRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ResendEnrollFactorRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *ResendEnrollFactorRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResendEnrollFactorRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResendEnrollFactorRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResendEnrollFactorRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVendorName

`func (o *ResendEnrollFactorRequest) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *ResendEnrollFactorRequest) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *ResendEnrollFactorRequest) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *ResendEnrollFactorRequest) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetEmbedded

`func (o *ResendEnrollFactorRequest) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ResendEnrollFactorRequest) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ResendEnrollFactorRequest) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ResendEnrollFactorRequest) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *ResendEnrollFactorRequest) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResendEnrollFactorRequest) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResendEnrollFactorRequest) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResendEnrollFactorRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


