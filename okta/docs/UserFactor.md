# UserFactor

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

`func (o *UserFactor) GetFactorType() FactorType`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *UserFactor) GetFactorTypeOk() (*FactorType, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *UserFactor) SetFactorType(v FactorType)`

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

### GetProvider

`func (o *UserFactor) GetProvider() FactorProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserFactor) GetProviderOk() (*FactorProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserFactor) SetProvider(v FactorProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserFactor) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *UserFactor) GetStatus() FactorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserFactor) GetStatusOk() (*FactorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserFactor) SetStatus(v FactorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserFactor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerify

`func (o *UserFactor) GetVerify() VerifyFactorRequest`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *UserFactor) GetVerifyOk() (*VerifyFactorRequest, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *UserFactor) SetVerify(v VerifyFactorRequest)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *UserFactor) HasVerify() bool`

HasVerify returns a boolean if a field has been set.

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

`func (o *UserFactor) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserFactor) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserFactor) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserFactor) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


