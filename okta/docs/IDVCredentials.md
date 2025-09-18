# IDVCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bearer** | Pointer to [**IDVCredentialsBearer**](IDVCredentialsBearer.md) |  | [optional] 
**Client** | Pointer to [**IDVCredentialsClient**](IDVCredentialsClient.md) |  | [optional] 

## Methods

### NewIDVCredentials

`func NewIDVCredentials() *IDVCredentials`

NewIDVCredentials instantiates a new IDVCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIDVCredentialsWithDefaults

`func NewIDVCredentialsWithDefaults() *IDVCredentials`

NewIDVCredentialsWithDefaults instantiates a new IDVCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBearer

`func (o *IDVCredentials) GetBearer() IDVCredentialsBearer`

GetBearer returns the Bearer field if non-nil, zero value otherwise.

### GetBearerOk

`func (o *IDVCredentials) GetBearerOk() (*IDVCredentialsBearer, bool)`

GetBearerOk returns a tuple with the Bearer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearer

`func (o *IDVCredentials) SetBearer(v IDVCredentialsBearer)`

SetBearer sets Bearer field to given value.

### HasBearer

`func (o *IDVCredentials) HasBearer() bool`

HasBearer returns a boolean if a field has been set.

### GetClient

`func (o *IDVCredentials) GetClient() IDVCredentialsClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *IDVCredentials) GetClientOk() (*IDVCredentialsClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *IDVCredentials) SetClient(v IDVCredentialsClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *IDVCredentials) HasClient() bool`

HasClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


