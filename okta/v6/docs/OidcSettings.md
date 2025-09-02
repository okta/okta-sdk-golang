# OidcSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParticipateSlo** | Pointer to **bool** | Set to &#x60;true&#x60; to have Okta send a logout request to the upstream IdP when a user signs out of Okta or a downstream app. | [optional] 
**SendApplicationContext** | Pointer to **bool** | Determines if the IdP should send the application context as &#x60;OktaAppInstanceId&#x60; and &#x60;OktaAppName&#x60; params in the request | [optional] [default to false]

## Methods

### NewOidcSettings

`func NewOidcSettings() *OidcSettings`

NewOidcSettings instantiates a new OidcSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcSettingsWithDefaults

`func NewOidcSettingsWithDefaults() *OidcSettings`

NewOidcSettingsWithDefaults instantiates a new OidcSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParticipateSlo

`func (o *OidcSettings) GetParticipateSlo() bool`

GetParticipateSlo returns the ParticipateSlo field if non-nil, zero value otherwise.

### GetParticipateSloOk

`func (o *OidcSettings) GetParticipateSloOk() (*bool, bool)`

GetParticipateSloOk returns a tuple with the ParticipateSlo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipateSlo

`func (o *OidcSettings) SetParticipateSlo(v bool)`

SetParticipateSlo sets ParticipateSlo field to given value.

### HasParticipateSlo

`func (o *OidcSettings) HasParticipateSlo() bool`

HasParticipateSlo returns a boolean if a field has been set.

### GetSendApplicationContext

`func (o *OidcSettings) GetSendApplicationContext() bool`

GetSendApplicationContext returns the SendApplicationContext field if non-nil, zero value otherwise.

### GetSendApplicationContextOk

`func (o *OidcSettings) GetSendApplicationContextOk() (*bool, bool)`

GetSendApplicationContextOk returns a tuple with the SendApplicationContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendApplicationContext

`func (o *OidcSettings) SetSendApplicationContext(v bool)`

SetSendApplicationContext sets SendApplicationContext field to given value.

### HasSendApplicationContext

`func (o *OidcSettings) HasSendApplicationContext() bool`

HasSendApplicationContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


