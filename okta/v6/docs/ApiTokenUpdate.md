# ApiTokenUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | Pointer to **string** | The client name associated with the API Token | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The creation date of the API Token | [optional] [readonly] 
**Name** | Pointer to **string** | The name associated with the API Token | [optional] 
**Network** | Pointer to [**ApiTokenNetwork**](ApiTokenNetwork.md) |  | [optional] 
**UserId** | Pointer to **string** | The userId of the user who created the API Token | [optional] 

## Methods

### NewApiTokenUpdate

`func NewApiTokenUpdate() *ApiTokenUpdate`

NewApiTokenUpdate instantiates a new ApiTokenUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenUpdateWithDefaults

`func NewApiTokenUpdateWithDefaults() *ApiTokenUpdate`

NewApiTokenUpdateWithDefaults instantiates a new ApiTokenUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *ApiTokenUpdate) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ApiTokenUpdate) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ApiTokenUpdate) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ApiTokenUpdate) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetCreated

`func (o *ApiTokenUpdate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiTokenUpdate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiTokenUpdate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ApiTokenUpdate) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetName

`func (o *ApiTokenUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiTokenUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiTokenUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiTokenUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *ApiTokenUpdate) GetNetwork() ApiTokenNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ApiTokenUpdate) GetNetworkOk() (*ApiTokenNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ApiTokenUpdate) SetNetwork(v ApiTokenNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ApiTokenUpdate) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetUserId

`func (o *ApiTokenUpdate) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiTokenUpdate) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiTokenUpdate) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiTokenUpdate) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


