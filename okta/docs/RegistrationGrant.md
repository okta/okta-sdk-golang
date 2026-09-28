# RegistrationGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumedAt** | Pointer to **NullableTime** | Timestamp when the grant was consumed by successful device registration. &#x60;null&#x60; while &#x60;PENDING&#x60;. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Timestamp when the grant was created | [optional] [readonly] 
**GrantType** | Pointer to **string** | Type of registration grant | [optional] 
**Id** | Pointer to **string** | Unique identifier for the registration grant (prefix &#x60;dpr&#x60;) | [optional] [readonly] 
**RegistrationSecretId** | Pointer to **string** | ID of the registration secret used for this grant | [optional] 
**Status** | Pointer to **string** | Status of the registration grant | [optional] 

## Methods

### NewRegistrationGrant

`func NewRegistrationGrant() *RegistrationGrant`

NewRegistrationGrant instantiates a new RegistrationGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationGrantWithDefaults

`func NewRegistrationGrantWithDefaults() *RegistrationGrant`

NewRegistrationGrantWithDefaults instantiates a new RegistrationGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumedAt

`func (o *RegistrationGrant) GetConsumedAt() time.Time`

GetConsumedAt returns the ConsumedAt field if non-nil, zero value otherwise.

### GetConsumedAtOk

`func (o *RegistrationGrant) GetConsumedAtOk() (*time.Time, bool)`

GetConsumedAtOk returns a tuple with the ConsumedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedAt

`func (o *RegistrationGrant) SetConsumedAt(v time.Time)`

SetConsumedAt sets ConsumedAt field to given value.

### HasConsumedAt

`func (o *RegistrationGrant) HasConsumedAt() bool`

HasConsumedAt returns a boolean if a field has been set.

### SetConsumedAtNil

`func (o *RegistrationGrant) SetConsumedAtNil(b bool)`

 SetConsumedAtNil sets the value for ConsumedAt to be an explicit nil

### UnsetConsumedAt
`func (o *RegistrationGrant) UnsetConsumedAt()`

UnsetConsumedAt ensures that no value is present for ConsumedAt, not even an explicit nil
### GetCreated

`func (o *RegistrationGrant) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RegistrationGrant) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RegistrationGrant) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RegistrationGrant) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetGrantType

`func (o *RegistrationGrant) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *RegistrationGrant) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *RegistrationGrant) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *RegistrationGrant) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetId

`func (o *RegistrationGrant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistrationGrant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistrationGrant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegistrationGrant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRegistrationSecretId

`func (o *RegistrationGrant) GetRegistrationSecretId() string`

GetRegistrationSecretId returns the RegistrationSecretId field if non-nil, zero value otherwise.

### GetRegistrationSecretIdOk

`func (o *RegistrationGrant) GetRegistrationSecretIdOk() (*string, bool)`

GetRegistrationSecretIdOk returns a tuple with the RegistrationSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationSecretId

`func (o *RegistrationGrant) SetRegistrationSecretId(v string)`

SetRegistrationSecretId sets RegistrationSecretId field to given value.

### HasRegistrationSecretId

`func (o *RegistrationGrant) HasRegistrationSecretId() bool`

HasRegistrationSecretId returns a boolean if a field has been set.

### GetStatus

`func (o *RegistrationGrant) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegistrationGrant) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegistrationGrant) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegistrationGrant) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


