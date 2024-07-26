# PrivilegedResourceAccountApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerDetails** | Pointer to [**AppAccountContainerDetails**](AppAccountContainerDetails.md) |  | [optional] 
**Credentials** | [**PrivilegedResourceCredentials**](PrivilegedResourceCredentials.md) |  | 

## Methods

### NewPrivilegedResourceAccountApp

`func NewPrivilegedResourceAccountApp(credentials PrivilegedResourceCredentials, ) *PrivilegedResourceAccountApp`

NewPrivilegedResourceAccountApp instantiates a new PrivilegedResourceAccountApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegedResourceAccountAppWithDefaults

`func NewPrivilegedResourceAccountAppWithDefaults() *PrivilegedResourceAccountApp`

NewPrivilegedResourceAccountAppWithDefaults instantiates a new PrivilegedResourceAccountApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerDetails

`func (o *PrivilegedResourceAccountApp) GetContainerDetails() AppAccountContainerDetails`

GetContainerDetails returns the ContainerDetails field if non-nil, zero value otherwise.

### GetContainerDetailsOk

`func (o *PrivilegedResourceAccountApp) GetContainerDetailsOk() (*AppAccountContainerDetails, bool)`

GetContainerDetailsOk returns a tuple with the ContainerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDetails

`func (o *PrivilegedResourceAccountApp) SetContainerDetails(v AppAccountContainerDetails)`

SetContainerDetails sets ContainerDetails field to given value.

### HasContainerDetails

`func (o *PrivilegedResourceAccountApp) HasContainerDetails() bool`

HasContainerDetails returns a boolean if a field has been set.

### GetCredentials

`func (o *PrivilegedResourceAccountApp) GetCredentials() PrivilegedResourceCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *PrivilegedResourceAccountApp) GetCredentialsOk() (*PrivilegedResourceCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *PrivilegedResourceAccountApp) SetCredentials(v PrivilegedResourceCredentials)`

SetCredentials sets Credentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


