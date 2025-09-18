# RoleGovernanceSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | Pointer to **string** | &#x60;id&#x60; of the entitlement bundle | [optional] [readonly] 
**ExpirationDate** | Pointer to **time.Time** | The expiration date of the entitlement bundle | [optional] [readonly] 
**GrantId** | **string** | &#x60;id&#x60; of the grant | [readonly] 
**Type** | **string** | The grant type | 
**Links** | Pointer to [**RoleGovernanceSourceLinks**](RoleGovernanceSourceLinks.md) |  | [optional] 

## Methods

### NewRoleGovernanceSource

`func NewRoleGovernanceSource(grantId string, type_ string, ) *RoleGovernanceSource`

NewRoleGovernanceSource instantiates a new RoleGovernanceSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleGovernanceSourceWithDefaults

`func NewRoleGovernanceSourceWithDefaults() *RoleGovernanceSource`

NewRoleGovernanceSourceWithDefaults instantiates a new RoleGovernanceSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *RoleGovernanceSource) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *RoleGovernanceSource) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *RoleGovernanceSource) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *RoleGovernanceSource) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetExpirationDate

`func (o *RoleGovernanceSource) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *RoleGovernanceSource) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *RoleGovernanceSource) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *RoleGovernanceSource) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetGrantId

`func (o *RoleGovernanceSource) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *RoleGovernanceSource) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *RoleGovernanceSource) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.


### GetType

`func (o *RoleGovernanceSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleGovernanceSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleGovernanceSource) SetType(v string)`

SetType sets Type field to given value.


### GetLinks

`func (o *RoleGovernanceSource) GetLinks() RoleGovernanceSourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoleGovernanceSource) GetLinksOk() (*RoleGovernanceSourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoleGovernanceSource) SetLinks(v RoleGovernanceSourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RoleGovernanceSource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


