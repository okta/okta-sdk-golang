# UserFactorSecurityQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorType** | Pointer to **interface{}** |  | [optional] 
**Profile** | Pointer to [**UserFactorSecurityQuestionProfile**](UserFactorSecurityQuestionProfile.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 

## Methods

### NewUserFactorSecurityQuestion

`func NewUserFactorSecurityQuestion() *UserFactorSecurityQuestion`

NewUserFactorSecurityQuestion instantiates a new UserFactorSecurityQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorSecurityQuestionWithDefaults

`func NewUserFactorSecurityQuestionWithDefaults() *UserFactorSecurityQuestion`

NewUserFactorSecurityQuestionWithDefaults instantiates a new UserFactorSecurityQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorType

`func (o *UserFactorSecurityQuestion) GetFactorType() interface{}`

GetFactorType returns the FactorType field if non-nil, zero value otherwise.

### GetFactorTypeOk

`func (o *UserFactorSecurityQuestion) GetFactorTypeOk() (*interface{}, bool)`

GetFactorTypeOk returns a tuple with the FactorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorType

`func (o *UserFactorSecurityQuestion) SetFactorType(v interface{})`

SetFactorType sets FactorType field to given value.

### HasFactorType

`func (o *UserFactorSecurityQuestion) HasFactorType() bool`

HasFactorType returns a boolean if a field has been set.

### SetFactorTypeNil

`func (o *UserFactorSecurityQuestion) SetFactorTypeNil(b bool)`

 SetFactorTypeNil sets the value for FactorType to be an explicit nil

### UnsetFactorType
`func (o *UserFactorSecurityQuestion) UnsetFactorType()`

UnsetFactorType ensures that no value is present for FactorType, not even an explicit nil
### GetProfile

`func (o *UserFactorSecurityQuestion) GetProfile() UserFactorSecurityQuestionProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserFactorSecurityQuestion) GetProfileOk() (*UserFactorSecurityQuestionProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserFactorSecurityQuestion) SetProfile(v UserFactorSecurityQuestionProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserFactorSecurityQuestion) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProvider

`func (o *UserFactorSecurityQuestion) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserFactorSecurityQuestion) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserFactorSecurityQuestion) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserFactorSecurityQuestion) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


