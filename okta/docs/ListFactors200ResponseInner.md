# ListFactors200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**FactorType** | Pointer to [**FactorType**](FactorType.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Provider** | Pointer to [**FactorProvider**](FactorProvider.md) |  | [optional] 
**Status** | Pointer to [**FactorStatus**](FactorStatus.md) |  | [optional] 
**Verify** | Pointer to [**VerifyFactorRequest**](VerifyFactorRequest.md) |  | [optional] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Profile** | Pointer to [**WebAuthnUserFactorProfile**](WebAuthnUserFactorProfile.md) |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**FactorResult** | Pointer to [**FactorResultType**](FactorResultType.md) |  | [optional] 
**FactorProfileId** | Pointer to **string** |  | [optional] 

## Methods

### NewListFactors200ResponseInner

`func NewListFactors200ResponseInner() *ListFactors200ResponseInner`

NewListFactors200ResponseInner instantiates a new ListFactors200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFactors200ResponseInnerWithDefaults

`func NewListFactors200ResponseInnerWithDefaults() *ListFactors200ResponseInner`

NewListFactors200ResponseInnerWithDefaults instantiates a new ListFactors200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ListFactors200ResponseInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListFactors200ResponseInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListFactors200ResponseInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListFactors200ResponseInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetFactorType

`func (o *ListFactors200ResponseInner) GetFactorType() FactorType`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *ListFactors200ResponseInner) GetFactorTypeOk() (*FactorType, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *ListFactors200ResponseInner) SetFactorType(v FactorType)`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *ListFactors200ResponseInner) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### GetId

`func (o *ListFactors200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListFactors200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListFactors200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListFactors200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListFactors200ResponseInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListFactors200ResponseInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListFactors200ResponseInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListFactors200ResponseInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProvider

`func (o *ListFactors200ResponseInner) GetProvider() FactorProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ListFactors200ResponseInner) GetProviderOk() (*FactorProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ListFactors200ResponseInner) SetProvider(v FactorProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ListFactors200ResponseInner) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *ListFactors200ResponseInner) GetStatus() FactorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListFactors200ResponseInner) GetStatusOk() (*FactorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListFactors200ResponseInner) SetStatus(v FactorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListFactors200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerify

`func (o *ListFactors200ResponseInner) GetVerify() VerifyFactorRequest`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *ListFactors200ResponseInner) GetVerifyOk() (*VerifyFactorRequest, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *ListFactors200ResponseInner) SetVerify(v VerifyFactorRequest)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *ListFactors200ResponseInner) HasVerify() bool`

HasVerify returns a boolean if a field has been set.

### GetEmbedded

`func (o *ListFactors200ResponseInner) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ListFactors200ResponseInner) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ListFactors200ResponseInner) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ListFactors200ResponseInner) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *ListFactors200ResponseInner) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListFactors200ResponseInner) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListFactors200ResponseInner) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListFactors200ResponseInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetProfile

`func (o *ListFactors200ResponseInner) GetProfile() WebAuthnUserFactorProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ListFactors200ResponseInner) GetProfileOk() (*WebAuthnUserFactorProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ListFactors200ResponseInner) SetProfile(v WebAuthnUserFactorProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ListFactors200ResponseInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ListFactors200ResponseInner) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ListFactors200ResponseInner) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ListFactors200ResponseInner) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ListFactors200ResponseInner) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFactorResult

`func (o *ListFactors200ResponseInner) GetFactorResult() FactorResultType`

GetFactorResult returns the FactorResult field if non-nil, zero value otherwise.

### GetFactorResultOk

`func (o *ListFactors200ResponseInner) GetFactorResultOk() (*FactorResultType, bool)`

GetFactorResultOk returns a tuple with the FactorResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorResult

`func (o *ListFactors200ResponseInner) SetFactorResult(v FactorResultType)`

SetFactorResult sets FactorResult field to given value.

### HasFactorResult

`func (o *ListFactors200ResponseInner) HasFactorResult() bool`

HasFactorResult returns a boolean if a field has been set.

### GetFactorProfileId

`func (o *ListFactors200ResponseInner) GetFactorProfileId() string`

GetFactorProfileId returns the FactorProfileId field if non-nil, zero value otherwise.

### GetFactorProfileIdOk

`func (o *ListFactors200ResponseInner) GetFactorProfileIdOk() (*string, bool)`

GetFactorProfileIdOk returns a tuple with the FactorProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorProfileId

`func (o *ListFactors200ResponseInner) SetFactorProfileId(v string)`

SetFactorProfileId sets FactorProfileId field to given value.

### HasFactorProfileId

`func (o *ListFactors200ResponseInner) HasFactorProfileId() bool`

HasFactorProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


