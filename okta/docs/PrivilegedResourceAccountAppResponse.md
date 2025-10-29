# PrivilegedResourceAccountAppResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**PrivilegedResourceCredentials**](PrivilegedResourceCredentials.md) |  | [optional] 
**Profile** | Pointer to **map[string]interface{}** | Specific profile properties for the privileged resource | [optional] [readonly] 

## Methods

### NewPrivilegedResourceAccountAppResponse

`func NewPrivilegedResourceAccountAppResponse() *PrivilegedResourceAccountAppResponse`

NewPrivilegedResourceAccountAppResponse instantiates a new PrivilegedResourceAccountAppResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegedResourceAccountAppResponseWithDefaults

`func NewPrivilegedResourceAccountAppResponseWithDefaults() *PrivilegedResourceAccountAppResponse`

NewPrivilegedResourceAccountAppResponseWithDefaults instantiates a new PrivilegedResourceAccountAppResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *PrivilegedResourceAccountAppResponse) GetCredentials() PrivilegedResourceCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *PrivilegedResourceAccountAppResponse) GetCredentialsOk() (*PrivilegedResourceCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *PrivilegedResourceAccountAppResponse) SetCredentials(v PrivilegedResourceCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *PrivilegedResourceAccountAppResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetProfile

`func (o *PrivilegedResourceAccountAppResponse) GetProfile() map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PrivilegedResourceAccountAppResponse) GetProfileOk() (*map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PrivilegedResourceAccountAppResponse) SetProfile(v map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PrivilegedResourceAccountAppResponse) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


