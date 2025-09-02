# PrivilegedResourceAccountAppRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerDetails** | Pointer to [**AppAccountContainerDetails**](AppAccountContainerDetails.md) |  | [optional] 
**Credentials** | Pointer to [**PrivilegedResourceCredentials**](PrivilegedResourceCredentials.md) |  | [optional] 

## Methods

### NewPrivilegedResourceAccountAppRequest

`func NewPrivilegedResourceAccountAppRequest() *PrivilegedResourceAccountAppRequest`

NewPrivilegedResourceAccountAppRequest instantiates a new PrivilegedResourceAccountAppRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegedResourceAccountAppRequestWithDefaults

`func NewPrivilegedResourceAccountAppRequestWithDefaults() *PrivilegedResourceAccountAppRequest`

NewPrivilegedResourceAccountAppRequestWithDefaults instantiates a new PrivilegedResourceAccountAppRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerDetails

`func (o *PrivilegedResourceAccountAppRequest) GetContainerDetails() AppAccountContainerDetails`

GetContainerDetails returns the ContainerDetails field if non-nil, zero value otherwise.

### GetContainerDetailsOk

`func (o *PrivilegedResourceAccountAppRequest) GetContainerDetailsOk() (*AppAccountContainerDetails, bool)`

GetContainerDetailsOk returns a tuple with the ContainerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDetails

`func (o *PrivilegedResourceAccountAppRequest) SetContainerDetails(v AppAccountContainerDetails)`

SetContainerDetails sets ContainerDetails field to given value.

### HasContainerDetails

`func (o *PrivilegedResourceAccountAppRequest) HasContainerDetails() bool`

HasContainerDetails returns a boolean if a field has been set.

### GetCredentials

`func (o *PrivilegedResourceAccountAppRequest) GetCredentials() PrivilegedResourceCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *PrivilegedResourceAccountAppRequest) GetCredentialsOk() (*PrivilegedResourceCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *PrivilegedResourceAccountAppRequest) SetCredentials(v PrivilegedResourceCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *PrivilegedResourceAccountAppRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


