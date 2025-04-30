# AAGUIDGroupObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aaguids** | Pointer to **[]string** | A list of YubiKey hardware FIDO2 Authenticator Attestation Global Unique Identifiers (AAGUIDs). The available [AAGUIDs](https://support.yubico.com/hc/en-us/articles/360016648959-YubiKey-Hardware-FIDO2-AAGUIDs) (opens new window) are provided by the FIDO Alliance Metadata Service. | [optional] 
**Name** | Pointer to **string** | A name to identify the group of YubiKey hardware FIDO2 AAGUIDs | [optional] 

## Methods

### NewAAGUIDGroupObject

`func NewAAGUIDGroupObject() *AAGUIDGroupObject`

NewAAGUIDGroupObject instantiates a new AAGUIDGroupObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAAGUIDGroupObjectWithDefaults

`func NewAAGUIDGroupObjectWithDefaults() *AAGUIDGroupObject`

NewAAGUIDGroupObjectWithDefaults instantiates a new AAGUIDGroupObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAaguids

`func (o *AAGUIDGroupObject) GetAaguids() []string`

GetAaguids returns the Aaguids field if non-nil, zero value otherwise.

### GetAaguidsOk

`func (o *AAGUIDGroupObject) GetAaguidsOk() (*[]string, bool)`

GetAaguidsOk returns a tuple with the Aaguids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaguids

`func (o *AAGUIDGroupObject) SetAaguids(v []string)`

SetAaguids sets Aaguids field to given value.

### HasAaguids

`func (o *AAGUIDGroupObject) HasAaguids() bool`

HasAaguids returns a boolean if a field has been set.

### GetName

`func (o *AAGUIDGroupObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AAGUIDGroupObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AAGUIDGroupObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AAGUIDGroupObject) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


