# UserFactor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the Factor was enrolled | [optional] [readonly] 
**FactorType** | Pointer to **string** | Type of Factor | [optional] 
**Id** | Pointer to **string** | ID of the Factor | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the Factor was last updated | [optional] [readonly] 
**Profile** | Pointer to **map[string]interface{}** | Specific attributes related to the Factor | [optional] 
**Provider** | Pointer to **string** | Provider for the Factor | [optional] 
**Status** | Pointer to **string** | Status of the Factor | [optional] [readonly] 
**VendorName** | Pointer to **string** | Name of the Factor vendor. This is usually the same as the provider except for On-Prem MFA where it depends on administrator settings. | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUserFactor

`func NewUserFactor() *UserFactor`

NewUserFactor instantiates a new UserFactor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorWithDefaults

`func NewUserFactorWithDefaults() *UserFactor`

NewUserFactorWithDefaults instantiates a new UserFactor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *UserFactor) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserFactor) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserFactor) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserFactor) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFactorType

`func (o *UserFactor) GetFactorType() string`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *UserFactor) GetFactorTypeOk() (*string, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *UserFactor) SetFactorType(v string)`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *UserFactor) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### GetId

`func (o *UserFactor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserFactor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserFactor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserFactor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UserFactor) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UserFactor) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UserFactor) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UserFactor) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProfile

`func (o *UserFactor) GetProfile() map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserFactor) GetProfileOk() (*map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserFactor) SetProfile(v map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserFactor) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProvider

`func (o *UserFactor) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserFactor) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserFactor) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserFactor) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *UserFactor) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserFactor) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserFactor) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserFactor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVendorName

`func (o *UserFactor) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *UserFactor) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *UserFactor) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *UserFactor) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetEmbedded

`func (o *UserFactor) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *UserFactor) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *UserFactor) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *UserFactor) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *UserFactor) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserFactor) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserFactor) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserFactor) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


