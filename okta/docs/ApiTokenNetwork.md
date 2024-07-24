# ApiTokenNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | Pointer to **string** | The connection type of the Network Condition | [optional] 
**Include** | Pointer to **[]string** | List of included IP network zones | [optional] 
**Exclude** | Pointer to **[]string** | List of excluded IP network zones | [optional] 

## Methods

### NewApiTokenNetwork

`func NewApiTokenNetwork() *ApiTokenNetwork`

NewApiTokenNetwork instantiates a new ApiTokenNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenNetworkWithDefaults

`func NewApiTokenNetworkWithDefaults() *ApiTokenNetwork`

NewApiTokenNetworkWithDefaults instantiates a new ApiTokenNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *ApiTokenNetwork) GetConnection() string`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ApiTokenNetwork) GetConnectionOk() (*string, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ApiTokenNetwork) SetConnection(v string)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ApiTokenNetwork) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetInclude

`func (o *ApiTokenNetwork) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *ApiTokenNetwork) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *ApiTokenNetwork) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *ApiTokenNetwork) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### GetExclude

`func (o *ApiTokenNetwork) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *ApiTokenNetwork) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *ApiTokenNetwork) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *ApiTokenNetwork) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


