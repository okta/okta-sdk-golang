# ActivateFactor200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the factor was enrolled | [optional] [readonly] 
**FactorType** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID of the factor | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the factor was last updated | [optional] [readonly] 
**Profile** | Pointer to [**UserFactorWebAuthnProfile**](UserFactorWebAuthnProfile.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Status of the factor | [optional] [readonly] 
**VendorName** | Pointer to **string** | Name of the factor vendor. This is usually the same as the provider except for On-Prem MFA, which depends on admin settings. | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**UserFactorLinks**](UserFactorLinks.md) |  | [optional] 

## Methods

### NewActivateFactor200Response

`func NewActivateFactor200Response() *ActivateFactor200Response`

NewActivateFactor200Response instantiates a new ActivateFactor200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivateFactor200ResponseWithDefaults

`func NewActivateFactor200ResponseWithDefaults() *ActivateFactor200Response`

NewActivateFactor200ResponseWithDefaults instantiates a new ActivateFactor200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ActivateFactor200Response) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ActivateFactor200Response) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ActivateFactor200Response) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ActivateFactor200Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFactorType

`func (o *ActivateFactor200Response) GetFactorType() interface{}`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *ActivateFactor200Response) GetFactorTypeOk() (*interface{}, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *ActivateFactor200Response) SetFactorType(v interface{})`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *ActivateFactor200Response) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### SetFactorTypeNil

`func (o *ActivateFactor200Response) SetFactorTypeNil(b bool)`

 SetFactorTypeNil sets the value for FactorType to be an explicit nil

### UnsetFactorType
`func (o *ActivateFactor200Response) UnsetFactorType()`

UnsetFactorType ensures that no value is present for FactorType, not even an explicit nil
### GetId

`func (o *ActivateFactor200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivateFactor200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivateFactor200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivateFactor200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ActivateFactor200Response) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ActivateFactor200Response) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ActivateFactor200Response) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ActivateFactor200Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProfile

`func (o *ActivateFactor200Response) GetProfile() UserFactorWebAuthnProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ActivateFactor200Response) GetProfileOk() (*UserFactorWebAuthnProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ActivateFactor200Response) SetProfile(v UserFactorWebAuthnProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ActivateFactor200Response) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProvider

`func (o *ActivateFactor200Response) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ActivateFactor200Response) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ActivateFactor200Response) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ActivateFactor200Response) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *ActivateFactor200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActivateFactor200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActivateFactor200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActivateFactor200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVendorName

`func (o *ActivateFactor200Response) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *ActivateFactor200Response) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *ActivateFactor200Response) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *ActivateFactor200Response) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetEmbedded

`func (o *ActivateFactor200Response) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ActivateFactor200Response) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ActivateFactor200Response) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ActivateFactor200Response) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *ActivateFactor200Response) GetLinks() UserFactorLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ActivateFactor200Response) GetLinksOk() (*UserFactorLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ActivateFactor200Response) SetLinks(v UserFactorLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ActivateFactor200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


