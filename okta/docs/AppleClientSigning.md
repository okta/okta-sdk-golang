# AppleClientSigning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kid** | Pointer to **string** | The key ID that you obtained from Apple when you created the private key for the client | [optional] 
**PrivateKey** | Pointer to **string** | The PKCS \\#8 encoded private key that you created for the client and downloaded from Apple | [optional] 
**TeamId** | Pointer to **string** | The Team ID associated with your Apple developer account | [optional] 

## Methods

### NewAppleClientSigning

`func NewAppleClientSigning() *AppleClientSigning`

NewAppleClientSigning instantiates a new AppleClientSigning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleClientSigningWithDefaults

`func NewAppleClientSigningWithDefaults() *AppleClientSigning`

NewAppleClientSigningWithDefaults instantiates a new AppleClientSigning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKid

`func (o *AppleClientSigning) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *AppleClientSigning) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *AppleClientSigning) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *AppleClientSigning) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetPrivateKey

`func (o *AppleClientSigning) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *AppleClientSigning) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *AppleClientSigning) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *AppleClientSigning) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetTeamId

`func (o *AppleClientSigning) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *AppleClientSigning) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *AppleClientSigning) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *AppleClientSigning) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


