# APNSConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | (Optional) File name for Admin Console display | [optional] 
**KeyId** | Pointer to **string** | 10-character Key ID obtained from the Apple developer account | [optional] 
**TeamId** | Pointer to **string** | 10-character Team ID used to develop the iOS app | [optional] 
**TokenSigningKey** | Pointer to **string** | APNs private authentication token signing key | [optional] 

## Methods

### NewAPNSConfiguration

`func NewAPNSConfiguration() *APNSConfiguration`

NewAPNSConfiguration instantiates a new APNSConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPNSConfigurationWithDefaults

`func NewAPNSConfigurationWithDefaults() *APNSConfiguration`

NewAPNSConfigurationWithDefaults instantiates a new APNSConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *APNSConfiguration) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *APNSConfiguration) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *APNSConfiguration) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *APNSConfiguration) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetKeyId

`func (o *APNSConfiguration) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *APNSConfiguration) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *APNSConfiguration) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *APNSConfiguration) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetTeamId

`func (o *APNSConfiguration) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *APNSConfiguration) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *APNSConfiguration) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *APNSConfiguration) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetTokenSigningKey

`func (o *APNSConfiguration) GetTokenSigningKey() string`

GetTokenSigningKey returns the TokenSigningKey field if non-nil, zero value otherwise.

### GetTokenSigningKeyOk

`func (o *APNSConfiguration) GetTokenSigningKeyOk() (*string, bool)`

GetTokenSigningKeyOk returns a tuple with the TokenSigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSigningKey

`func (o *APNSConfiguration) SetTokenSigningKey(v string)`

SetTokenSigningKey sets TokenSigningKey field to given value.

### HasTokenSigningKey

`func (o *APNSConfiguration) HasTokenSigningKey() bool`

HasTokenSigningKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


