# IdentitySourceSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IdentitySourceId** | Pointer to **string** |  | [optional] [readonly] 
**ImportType** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewIdentitySourceSession

`func NewIdentitySourceSession() *IdentitySourceSession`

NewIdentitySourceSession instantiates a new IdentitySourceSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitySourceSessionWithDefaults

`func NewIdentitySourceSessionWithDefaults() *IdentitySourceSession`

NewIdentitySourceSessionWithDefaults instantiates a new IdentitySourceSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *IdentitySourceSession) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IdentitySourceSession) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IdentitySourceSession) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IdentitySourceSession) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *IdentitySourceSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentitySourceSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentitySourceSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentitySourceSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentitySourceId

`func (o *IdentitySourceSession) GetIdentitySourceId() string`

GetIdentitySourceId returns the IdentitySourceId field if non-nil, zero value otherwise.

### GetIdentitySourceIdOk

`func (o *IdentitySourceSession) GetIdentitySourceIdOk() (*string, bool)`

GetIdentitySourceIdOk returns a tuple with the IdentitySourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentitySourceId

`func (o *IdentitySourceSession) SetIdentitySourceId(v string)`

SetIdentitySourceId sets IdentitySourceId field to given value.

### HasIdentitySourceId

`func (o *IdentitySourceSession) HasIdentitySourceId() bool`

HasIdentitySourceId returns a boolean if a field has been set.

### GetImportType

`func (o *IdentitySourceSession) GetImportType() string`

GetImportType returns the ImportType field if non-nil, zero value otherwise.

### GetImportTypeOk

`func (o *IdentitySourceSession) GetImportTypeOk() (*string, bool)`

GetImportTypeOk returns a tuple with the ImportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportType

`func (o *IdentitySourceSession) SetImportType(v string)`

SetImportType sets ImportType field to given value.

### HasImportType

`func (o *IdentitySourceSession) HasImportType() bool`

HasImportType returns a boolean if a field has been set.

### GetLastUpdated

`func (o *IdentitySourceSession) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *IdentitySourceSession) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *IdentitySourceSession) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *IdentitySourceSession) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatus

`func (o *IdentitySourceSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentitySourceSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentitySourceSession) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentitySourceSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


