# SimulatePolicyBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppInstance** | **string** | The application instance ID for a simulate operation | 
**PolicyContext** | Pointer to [**PolicyContext**](PolicyContext.md) |  | [optional] 
**PolicyTypes** | Pointer to **[]string** | Supported policy types for a simulate operation. The default value, &#x60;null&#x60;, returns all types. | [optional] 

## Methods

### NewSimulatePolicyBody

`func NewSimulatePolicyBody(appInstance string, ) *SimulatePolicyBody`

NewSimulatePolicyBody instantiates a new SimulatePolicyBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulatePolicyBodyWithDefaults

`func NewSimulatePolicyBodyWithDefaults() *SimulatePolicyBody`

NewSimulatePolicyBodyWithDefaults instantiates a new SimulatePolicyBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppInstance

`func (o *SimulatePolicyBody) GetAppInstance() string`

GetAppInstance returns the AppInstance field if non-nil, zero value otherwise.

### GetAppInstanceOk

`func (o *SimulatePolicyBody) GetAppInstanceOk() (*string, bool)`

GetAppInstanceOk returns a tuple with the AppInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstance

`func (o *SimulatePolicyBody) SetAppInstance(v string)`

SetAppInstance sets AppInstance field to given value.


### GetPolicyContext

`func (o *SimulatePolicyBody) GetPolicyContext() PolicyContext`

GetPolicyContext returns the PolicyContext field if non-nil, zero value otherwise.

### GetPolicyContextOk

`func (o *SimulatePolicyBody) GetPolicyContextOk() (*PolicyContext, bool)`

GetPolicyContextOk returns a tuple with the PolicyContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyContext

`func (o *SimulatePolicyBody) SetPolicyContext(v PolicyContext)`

SetPolicyContext sets PolicyContext field to given value.

### HasPolicyContext

`func (o *SimulatePolicyBody) HasPolicyContext() bool`

HasPolicyContext returns a boolean if a field has been set.

### GetPolicyTypes

`func (o *SimulatePolicyBody) GetPolicyTypes() []string`

GetPolicyTypes returns the PolicyTypes field if non-nil, zero value otherwise.

### GetPolicyTypesOk

`func (o *SimulatePolicyBody) GetPolicyTypesOk() (*[]string, bool)`

GetPolicyTypesOk returns a tuple with the PolicyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTypes

`func (o *SimulatePolicyBody) SetPolicyTypes(v []string)`

SetPolicyTypes sets PolicyTypes field to given value.

### HasPolicyTypes

`func (o *SimulatePolicyBody) HasPolicyTypes() bool`

HasPolicyTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


