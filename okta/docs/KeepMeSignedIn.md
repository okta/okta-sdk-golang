# KeepMeSignedIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostAuth** | Pointer to **string** | Whether the post-authentication [Keep Me Signed In (KMSI)](https://help.okta.com/oie/en-us/content/topics/security/stay-signed-in.htm) flow is allowed | [optional] 
**PostAuthPromptFrequency** | Pointer to **string** | A time duration specified as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations). | [optional] 

## Methods

### NewKeepMeSignedIn

`func NewKeepMeSignedIn() *KeepMeSignedIn`

NewKeepMeSignedIn instantiates a new KeepMeSignedIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeepMeSignedInWithDefaults

`func NewKeepMeSignedInWithDefaults() *KeepMeSignedIn`

NewKeepMeSignedInWithDefaults instantiates a new KeepMeSignedIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostAuth

`func (o *KeepMeSignedIn) GetPostAuth() string`

GetPostAuth returns the PostAuth field if non-nil, zero value otherwise.

### GetPostAuthOk

`func (o *KeepMeSignedIn) GetPostAuthOk() (*string, bool)`

GetPostAuthOk returns a tuple with the PostAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostAuth

`func (o *KeepMeSignedIn) SetPostAuth(v string)`

SetPostAuth sets PostAuth field to given value.

### HasPostAuth

`func (o *KeepMeSignedIn) HasPostAuth() bool`

HasPostAuth returns a boolean if a field has been set.

### GetPostAuthPromptFrequency

`func (o *KeepMeSignedIn) GetPostAuthPromptFrequency() string`

GetPostAuthPromptFrequency returns the PostAuthPromptFrequency field if non-nil, zero value otherwise.

### GetPostAuthPromptFrequencyOk

`func (o *KeepMeSignedIn) GetPostAuthPromptFrequencyOk() (*string, bool)`

GetPostAuthPromptFrequencyOk returns a tuple with the PostAuthPromptFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostAuthPromptFrequency

`func (o *KeepMeSignedIn) SetPostAuthPromptFrequency(v string)`

SetPostAuthPromptFrequency sets PostAuthPromptFrequency field to given value.

### HasPostAuthPromptFrequency

`func (o *KeepMeSignedIn) HasPostAuthPromptFrequency() bool`

HasPostAuthPromptFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


