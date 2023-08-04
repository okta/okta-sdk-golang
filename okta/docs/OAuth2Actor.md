# OAuth2Actor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuth2Actor

`func NewOAuth2Actor() *OAuth2Actor`

NewOAuth2Actor instantiates a new OAuth2Actor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ActorWithDefaults

`func NewOAuth2ActorWithDefaults() *OAuth2Actor`

NewOAuth2ActorWithDefaults instantiates a new OAuth2Actor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OAuth2Actor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2Actor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2Actor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuth2Actor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *OAuth2Actor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OAuth2Actor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OAuth2Actor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OAuth2Actor) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


