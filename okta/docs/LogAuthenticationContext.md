# LogAuthenticationContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationProvider** | Pointer to [**LogAuthenticationProvider**](LogAuthenticationProvider.md) |  | [optional] 
**AuthenticationStep** | Pointer to **int32** |  | [optional] [readonly] 
**CredentialProvider** | Pointer to [**LogCredentialProvider**](LogCredentialProvider.md) |  | [optional] 
**CredentialType** | Pointer to [**LogCredentialType**](LogCredentialType.md) |  | [optional] 
**ExternalSessionId** | Pointer to **string** |  | [optional] [readonly] 
**Interface** | Pointer to **string** |  | [optional] [readonly] 
**Issuer** | Pointer to [**LogIssuer**](LogIssuer.md) |  | [optional] 

## Methods

### NewLogAuthenticationContext

`func NewLogAuthenticationContext() *LogAuthenticationContext`

NewLogAuthenticationContext instantiates a new LogAuthenticationContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogAuthenticationContextWithDefaults

`func NewLogAuthenticationContextWithDefaults() *LogAuthenticationContext`

NewLogAuthenticationContextWithDefaults instantiates a new LogAuthenticationContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationProvider

`func (o *LogAuthenticationContext) GetAuthenticationProvider() LogAuthenticationProvider`

GetAuthenticationProvider returns the AuthenticationProvider field if non-nil, zero value otherwise.

### GetAuthenticationProviderOk

`func (o *LogAuthenticationContext) GetAuthenticationProviderOk() (*LogAuthenticationProvider, bool)`

GetAuthenticationProviderOk returns a tuple with the AuthenticationProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProvider

`func (o *LogAuthenticationContext) SetAuthenticationProvider(v LogAuthenticationProvider)`

SetAuthenticationProvider sets AuthenticationProvider field to given value.

### HasAuthenticationProvider

`func (o *LogAuthenticationContext) HasAuthenticationProvider() bool`

HasAuthenticationProvider returns a boolean if a field has been set.

### GetAuthenticationStep

`func (o *LogAuthenticationContext) GetAuthenticationStep() int32`

GetAuthenticationStep returns the AuthenticationStep field if non-nil, zero value otherwise.

### GetAuthenticationStepOk

`func (o *LogAuthenticationContext) GetAuthenticationStepOk() (*int32, bool)`

GetAuthenticationStepOk returns a tuple with the AuthenticationStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationStep

`func (o *LogAuthenticationContext) SetAuthenticationStep(v int32)`

SetAuthenticationStep sets AuthenticationStep field to given value.

### HasAuthenticationStep

`func (o *LogAuthenticationContext) HasAuthenticationStep() bool`

HasAuthenticationStep returns a boolean if a field has been set.

### GetCredentialProvider

`func (o *LogAuthenticationContext) GetCredentialProvider() LogCredentialProvider`

GetCredentialProvider returns the CredentialProvider field if non-nil, zero value otherwise.

### GetCredentialProviderOk

`func (o *LogAuthenticationContext) GetCredentialProviderOk() (*LogCredentialProvider, bool)`

GetCredentialProviderOk returns a tuple with the CredentialProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialProvider

`func (o *LogAuthenticationContext) SetCredentialProvider(v LogCredentialProvider)`

SetCredentialProvider sets CredentialProvider field to given value.

### HasCredentialProvider

`func (o *LogAuthenticationContext) HasCredentialProvider() bool`

HasCredentialProvider returns a boolean if a field has been set.

### GetCredentialType

`func (o *LogAuthenticationContext) GetCredentialType() LogCredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *LogAuthenticationContext) GetCredentialTypeOk() (*LogCredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *LogAuthenticationContext) SetCredentialType(v LogCredentialType)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *LogAuthenticationContext) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetExternalSessionId

`func (o *LogAuthenticationContext) GetExternalSessionId() string`

GetExternalSessionId returns the ExternalSessionId field if non-nil, zero value otherwise.

### GetExternalSessionIdOk

`func (o *LogAuthenticationContext) GetExternalSessionIdOk() (*string, bool)`

GetExternalSessionIdOk returns a tuple with the ExternalSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSessionId

`func (o *LogAuthenticationContext) SetExternalSessionId(v string)`

SetExternalSessionId sets ExternalSessionId field to given value.

### HasExternalSessionId

`func (o *LogAuthenticationContext) HasExternalSessionId() bool`

HasExternalSessionId returns a boolean if a field has been set.

### GetInterface

`func (o *LogAuthenticationContext) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *LogAuthenticationContext) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *LogAuthenticationContext) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *LogAuthenticationContext) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetIssuer

`func (o *LogAuthenticationContext) GetIssuer() LogIssuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *LogAuthenticationContext) GetIssuerOk() (*LogIssuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *LogAuthenticationContext) SetIssuer(v LogIssuer)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *LogAuthenticationContext) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


