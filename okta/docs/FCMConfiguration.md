# FCMConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | (Optional) File name for Admin Console display | [optional] 
**ProjectId** | Pointer to **string** | Project ID of FCM configuration | [optional] [readonly] 
**ServiceAccountJson** | Pointer to **map[string]interface{}** | JSON containing the private service account key and service account details. See [Creating and managing service account keys](https://cloud.google.com/iam/docs/creating-managing-service-account-keys) for more information on creating service account keys in JSON. | [optional] 

## Methods

### NewFCMConfiguration

`func NewFCMConfiguration() *FCMConfiguration`

NewFCMConfiguration instantiates a new FCMConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFCMConfigurationWithDefaults

`func NewFCMConfigurationWithDefaults() *FCMConfiguration`

NewFCMConfigurationWithDefaults instantiates a new FCMConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *FCMConfiguration) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FCMConfiguration) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FCMConfiguration) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *FCMConfiguration) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetProjectId

`func (o *FCMConfiguration) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FCMConfiguration) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FCMConfiguration) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *FCMConfiguration) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetServiceAccountJson

`func (o *FCMConfiguration) GetServiceAccountJson() map[string]interface{}`

GetServiceAccountJson returns the ServiceAccountJson field if non-nil, zero value otherwise.

### GetServiceAccountJsonOk

`func (o *FCMConfiguration) GetServiceAccountJsonOk() (*map[string]interface{}, bool)`

GetServiceAccountJsonOk returns a tuple with the ServiceAccountJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountJson

`func (o *FCMConfiguration) SetServiceAccountJson(v map[string]interface{})`

SetServiceAccountJson sets ServiceAccountJson field to given value.

### HasServiceAccountJson

`func (o *FCMConfiguration) HasServiceAccountJson() bool`

HasServiceAccountJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


