# IDVCredentialsClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The client ID that you generate in your IDV | 
**ClientSecret** | **string** | The client secret that you generate in your IDV | 

## Methods

### NewIDVCredentialsClient

`func NewIDVCredentialsClient(clientId string, clientSecret string, ) *IDVCredentialsClient`

NewIDVCredentialsClient instantiates a new IDVCredentialsClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIDVCredentialsClientWithDefaults

`func NewIDVCredentialsClientWithDefaults() *IDVCredentialsClient`

NewIDVCredentialsClientWithDefaults instantiates a new IDVCredentialsClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IDVCredentialsClient) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IDVCredentialsClient) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IDVCredentialsClient) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *IDVCredentialsClient) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IDVCredentialsClient) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IDVCredentialsClient) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


