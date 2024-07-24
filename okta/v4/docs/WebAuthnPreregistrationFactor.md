# WebAuthnPreregistrationFactor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp indicating when the Factor was enrolled | [optional] [readonly] 
**FactorType** | Pointer to **string** | Type of Factor | [optional] 
**Id** | Pointer to **string** | ID of the Factor | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp indicating when the Factor was last updated | [optional] [readonly] 
**Profile** | Pointer to **map[string]interface{}** | Specific attributes related to the Factor | [optional] 
**Provider** | Pointer to **string** | Provider for the Factor | [optional] 
**Status** | Pointer to **string** | Status of the Factor | [optional] [readonly] 
**VendorName** | Pointer to **string** | Name of the Factor vendor. This is usually the same as the provider. | [optional] [readonly] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewWebAuthnPreregistrationFactor

`func NewWebAuthnPreregistrationFactor() *WebAuthnPreregistrationFactor`

NewWebAuthnPreregistrationFactor instantiates a new WebAuthnPreregistrationFactor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebAuthnPreregistrationFactorWithDefaults

`func NewWebAuthnPreregistrationFactorWithDefaults() *WebAuthnPreregistrationFactor`

NewWebAuthnPreregistrationFactorWithDefaults instantiates a new WebAuthnPreregistrationFactor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *WebAuthnPreregistrationFactor) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebAuthnPreregistrationFactor) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebAuthnPreregistrationFactor) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WebAuthnPreregistrationFactor) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFactorType

`func (o *WebAuthnPreregistrationFactor) GetFactorType() string`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *WebAuthnPreregistrationFactor) GetFactorTypeOk() (*string, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *WebAuthnPreregistrationFactor) SetFactorType(v string)`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *WebAuthnPreregistrationFactor) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### GetId

`func (o *WebAuthnPreregistrationFactor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebAuthnPreregistrationFactor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebAuthnPreregistrationFactor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebAuthnPreregistrationFactor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *WebAuthnPreregistrationFactor) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WebAuthnPreregistrationFactor) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WebAuthnPreregistrationFactor) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *WebAuthnPreregistrationFactor) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProfile

`func (o *WebAuthnPreregistrationFactor) GetProfile() map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebAuthnPreregistrationFactor) GetProfileOk() (*map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebAuthnPreregistrationFactor) SetProfile(v map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *WebAuthnPreregistrationFactor) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProvider

`func (o *WebAuthnPreregistrationFactor) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *WebAuthnPreregistrationFactor) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *WebAuthnPreregistrationFactor) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *WebAuthnPreregistrationFactor) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *WebAuthnPreregistrationFactor) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebAuthnPreregistrationFactor) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebAuthnPreregistrationFactor) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebAuthnPreregistrationFactor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVendorName

`func (o *WebAuthnPreregistrationFactor) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *WebAuthnPreregistrationFactor) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *WebAuthnPreregistrationFactor) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *WebAuthnPreregistrationFactor) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetLinks

`func (o *WebAuthnPreregistrationFactor) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WebAuthnPreregistrationFactor) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WebAuthnPreregistrationFactor) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WebAuthnPreregistrationFactor) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


