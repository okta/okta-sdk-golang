# ServiceAccountDetailsAppAccountSub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppGlobalName** | Pointer to **string** | The name of the SaaS app in the Okta Integration Network catalog | [optional] [readonly] 
**AppInstanceName** | Pointer to **string** | The instance name of the SaaS app | [optional] [readonly] 
**Credentials** | [**AppServiceAccountCredentials**](AppServiceAccountCredentials.md) |  | 
**OktaApplicationId** | **string** | The Okta app instance ID of the SaaS app | 

## Methods

### NewServiceAccountDetailsAppAccountSub

`func NewServiceAccountDetailsAppAccountSub(credentials AppServiceAccountCredentials, oktaApplicationId string, ) *ServiceAccountDetailsAppAccountSub`

NewServiceAccountDetailsAppAccountSub instantiates a new ServiceAccountDetailsAppAccountSub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountDetailsAppAccountSubWithDefaults

`func NewServiceAccountDetailsAppAccountSubWithDefaults() *ServiceAccountDetailsAppAccountSub`

NewServiceAccountDetailsAppAccountSubWithDefaults instantiates a new ServiceAccountDetailsAppAccountSub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppGlobalName

`func (o *ServiceAccountDetailsAppAccountSub) GetAppGlobalName() string`

GetAppGlobalName returns the AppGlobalName field if non-nil, zero value otherwise.

### GetAppGlobalNameOk

`func (o *ServiceAccountDetailsAppAccountSub) GetAppGlobalNameOk() (*string, bool)`

GetAppGlobalNameOk returns a tuple with the AppGlobalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppGlobalName

`func (o *ServiceAccountDetailsAppAccountSub) SetAppGlobalName(v string)`

SetAppGlobalName sets AppGlobalName field to given value.

### HasAppGlobalName

`func (o *ServiceAccountDetailsAppAccountSub) HasAppGlobalName() bool`

HasAppGlobalName returns a boolean if a field has been set.

### GetAppInstanceName

`func (o *ServiceAccountDetailsAppAccountSub) GetAppInstanceName() string`

GetAppInstanceName returns the AppInstanceName field if non-nil, zero value otherwise.

### GetAppInstanceNameOk

`func (o *ServiceAccountDetailsAppAccountSub) GetAppInstanceNameOk() (*string, bool)`

GetAppInstanceNameOk returns a tuple with the AppInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceName

`func (o *ServiceAccountDetailsAppAccountSub) SetAppInstanceName(v string)`

SetAppInstanceName sets AppInstanceName field to given value.

### HasAppInstanceName

`func (o *ServiceAccountDetailsAppAccountSub) HasAppInstanceName() bool`

HasAppInstanceName returns a boolean if a field has been set.

### GetCredentials

`func (o *ServiceAccountDetailsAppAccountSub) GetCredentials() AppServiceAccountCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ServiceAccountDetailsAppAccountSub) GetCredentialsOk() (*AppServiceAccountCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ServiceAccountDetailsAppAccountSub) SetCredentials(v AppServiceAccountCredentials)`

SetCredentials sets Credentials field to given value.


### GetOktaApplicationId

`func (o *ServiceAccountDetailsAppAccountSub) GetOktaApplicationId() string`

GetOktaApplicationId returns the OktaApplicationId field if non-nil, zero value otherwise.

### GetOktaApplicationIdOk

`func (o *ServiceAccountDetailsAppAccountSub) GetOktaApplicationIdOk() (*string, bool)`

GetOktaApplicationIdOk returns a tuple with the OktaApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaApplicationId

`func (o *ServiceAccountDetailsAppAccountSub) SetOktaApplicationId(v string)`

SetOktaApplicationId sets OktaApplicationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


