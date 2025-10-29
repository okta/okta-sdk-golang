# LogIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Varies depending on the type of authentication. If authentication is SAML 2.0, &#x60;id&#x60; is the issuer in the SAML assertion. For social login, &#x60;id&#x60; is the issuer of the token. | [optional] [readonly] 
**Type** | Pointer to **string** | Information on the &#x60;issuer&#x60; and source of the SAML assertion or token | [optional] [readonly] 

## Methods

### NewLogIssuer

`func NewLogIssuer() *LogIssuer`

NewLogIssuer instantiates a new LogIssuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogIssuerWithDefaults

`func NewLogIssuerWithDefaults() *LogIssuer`

NewLogIssuerWithDefaults instantiates a new LogIssuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogIssuer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogIssuer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogIssuer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogIssuer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LogIssuer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogIssuer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogIssuer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LogIssuer) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


