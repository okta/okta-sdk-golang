# GroupOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name of the group owner | [optional] [readonly] 
**Id** | Pointer to **string** | The &#x60;id&#x60; of the group owner | [optional] 
**LastUpdated** | Pointer to **string** | Timestamp when the group owner was last updated | [optional] [readonly] 
**OriginId** | Pointer to **string** | The ID of the app instance if the &#x60;originType&#x60; is &#x60;APPLICATION&#x60;. This value is &#x60;NULL&#x60; if &#x60;originType&#x60; is &#x60;OKTA_DIRECTORY&#x60;. | [optional] 
**OriginType** | Pointer to **string** | The source where group ownership is managed | [optional] 
**Resolved** | Pointer to **bool** | If &#x60;originType&#x60;is APPLICATION, this parameter is set to &#x60;FALSE&#x60; until the owner’s &#x60;originId&#x60; is reconciled with an associated Okta ID. | [optional] 
**Type** | Pointer to **string** | The entity type of the owner | [optional] 

## Methods

### NewGroupOwner

`func NewGroupOwner() *GroupOwner`

NewGroupOwner instantiates a new GroupOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupOwnerWithDefaults

`func NewGroupOwnerWithDefaults() *GroupOwner`

NewGroupOwnerWithDefaults instantiates a new GroupOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *GroupOwner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GroupOwner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GroupOwner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GroupOwner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *GroupOwner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupOwner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupOwner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupOwner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GroupOwner) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GroupOwner) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GroupOwner) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GroupOwner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOriginId

`func (o *GroupOwner) GetOriginId() string`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *GroupOwner) GetOriginIdOk() (*string, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *GroupOwner) SetOriginId(v string)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *GroupOwner) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### GetOriginType

`func (o *GroupOwner) GetOriginType() string`

GetOriginType returns the OriginType field if non-nil, zero value otherwise.

### GetOriginTypeOk

`func (o *GroupOwner) GetOriginTypeOk() (*string, bool)`

GetOriginTypeOk returns a tuple with the OriginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginType

`func (o *GroupOwner) SetOriginType(v string)`

SetOriginType sets OriginType field to given value.

### HasOriginType

`func (o *GroupOwner) HasOriginType() bool`

HasOriginType returns a boolean if a field has been set.

### GetResolved

`func (o *GroupOwner) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *GroupOwner) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *GroupOwner) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *GroupOwner) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetType

`func (o *GroupOwner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupOwner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupOwner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GroupOwner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


