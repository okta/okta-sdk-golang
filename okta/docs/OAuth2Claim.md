# OAuth2Claim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysIncludeInToken** | Pointer to **bool** |  | [optional] 
**ClaimType** | Pointer to [**OAuth2ClaimType**](OAuth2ClaimType.md) |  | [optional] 
**Conditions** | Pointer to [**OAuth2ClaimConditions**](OAuth2ClaimConditions.md) |  | [optional] 
**GroupFilterType** | Pointer to [**OAuth2ClaimGroupFilterType**](OAuth2ClaimGroupFilterType.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**LifecycleStatus**](LifecycleStatus.md) |  | [optional] 
**System** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**ValueType** | Pointer to [**OAuth2ClaimValueType**](OAuth2ClaimValueType.md) |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewOAuth2Claim

`func NewOAuth2Claim() *OAuth2Claim`

NewOAuth2Claim instantiates a new OAuth2Claim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClaimWithDefaults

`func NewOAuth2ClaimWithDefaults() *OAuth2Claim`

NewOAuth2ClaimWithDefaults instantiates a new OAuth2Claim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysIncludeInToken

`func (o *OAuth2Claim) GetAlwaysIncludeInToken() bool`

GetAlwaysIncludeInToken returns the AlwaysIncludeInToken field if non-nil, zero value otherwise.

### GetAlwaysIncludeInTokenOk

`func (o *OAuth2Claim) GetAlwaysIncludeInTokenOk() (*bool, bool)`

GetAlwaysIncludeInTokenOk returns a tuple with the AlwaysIncludeInToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysIncludeInToken

`func (o *OAuth2Claim) SetAlwaysIncludeInToken(v bool)`

SetAlwaysIncludeInToken sets AlwaysIncludeInToken field to given value.

### HasAlwaysIncludeInToken

`func (o *OAuth2Claim) HasAlwaysIncludeInToken() bool`

HasAlwaysIncludeInToken returns a boolean if a field has been set.

### GetClaimType

`func (o *OAuth2Claim) GetClaimType() OAuth2ClaimType`

GetClaimType returns the ClaimType field if non-nil, zero value otherwise.

### GetClaimTypeOk

`func (o *OAuth2Claim) GetClaimTypeOk() (*OAuth2ClaimType, bool)`

GetClaimTypeOk returns a tuple with the ClaimType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimType

`func (o *OAuth2Claim) SetClaimType(v OAuth2ClaimType)`

SetClaimType sets ClaimType field to given value.

### HasClaimType

`func (o *OAuth2Claim) HasClaimType() bool`

HasClaimType returns a boolean if a field has been set.

### GetConditions

`func (o *OAuth2Claim) GetConditions() OAuth2ClaimConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *OAuth2Claim) GetConditionsOk() (*OAuth2ClaimConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *OAuth2Claim) SetConditions(v OAuth2ClaimConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *OAuth2Claim) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetGroupFilterType

`func (o *OAuth2Claim) GetGroupFilterType() OAuth2ClaimGroupFilterType`

GetGroupFilterType returns the GroupFilterType field if non-nil, zero value otherwise.

### GetGroupFilterTypeOk

`func (o *OAuth2Claim) GetGroupFilterTypeOk() (*OAuth2ClaimGroupFilterType, bool)`

GetGroupFilterTypeOk returns a tuple with the GroupFilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupFilterType

`func (o *OAuth2Claim) SetGroupFilterType(v OAuth2ClaimGroupFilterType)`

SetGroupFilterType sets GroupFilterType field to given value.

### HasGroupFilterType

`func (o *OAuth2Claim) HasGroupFilterType() bool`

HasGroupFilterType returns a boolean if a field has been set.

### GetId

`func (o *OAuth2Claim) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2Claim) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2Claim) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuth2Claim) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OAuth2Claim) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuth2Claim) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuth2Claim) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OAuth2Claim) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *OAuth2Claim) GetStatus() LifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2Claim) GetStatusOk() (*LifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2Claim) SetStatus(v LifecycleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2Claim) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *OAuth2Claim) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *OAuth2Claim) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *OAuth2Claim) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *OAuth2Claim) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetValue

`func (o *OAuth2Claim) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OAuth2Claim) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OAuth2Claim) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *OAuth2Claim) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueType

`func (o *OAuth2Claim) GetValueType() OAuth2ClaimValueType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *OAuth2Claim) GetValueTypeOk() (*OAuth2ClaimValueType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *OAuth2Claim) SetValueType(v OAuth2ClaimValueType)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *OAuth2Claim) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetLinks

`func (o *OAuth2Claim) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OAuth2Claim) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OAuth2Claim) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OAuth2Claim) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


