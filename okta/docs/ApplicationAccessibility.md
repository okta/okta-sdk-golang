# ApplicationAccessibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorRedirectUrl** | Pointer to **string** | Custom error page URL for the app | [optional] 
**LoginRedirectUrl** | Pointer to **string** | Custom login page URL for the app &gt; **Note:** The &#x60;loginRedirectUrl&#x60; property is deprecated in Identity Engine. This property is used with the custom app login feature. Orgs that actively use this feature can continue to do so. See [Okta-hosted sign-in (redirect authentication)](https://developer.okta.com/docs/guides/redirect-authentication/) or [configure IdP routing rules](https://help.okta.com/okta_help.htm?type&#x3D;oie&amp;id&#x3D;ext-cfg-routing-rules) to redirect users to the appropriate sign-in app for orgs that don&#39;t use the custom app login feature. | [optional] 
**SelfService** | Pointer to **bool** | Represents whether the app can be self-assignable by users | [optional] 

## Methods

### NewApplicationAccessibility

`func NewApplicationAccessibility() *ApplicationAccessibility`

NewApplicationAccessibility instantiates a new ApplicationAccessibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAccessibilityWithDefaults

`func NewApplicationAccessibilityWithDefaults() *ApplicationAccessibility`

NewApplicationAccessibilityWithDefaults instantiates a new ApplicationAccessibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorRedirectUrl

`func (o *ApplicationAccessibility) GetErrorRedirectUrl() string`

GetErrorRedirectUrl returns the ErrorRedirectUrl field if non-nil, zero value otherwise.

### GetErrorRedirectUrlOk

`func (o *ApplicationAccessibility) GetErrorRedirectUrlOk() (*string, bool)`

GetErrorRedirectUrlOk returns a tuple with the ErrorRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRedirectUrl

`func (o *ApplicationAccessibility) SetErrorRedirectUrl(v string)`

SetErrorRedirectUrl sets ErrorRedirectUrl field to given value.

### HasErrorRedirectUrl

`func (o *ApplicationAccessibility) HasErrorRedirectUrl() bool`

HasErrorRedirectUrl returns a boolean if a field has been set.

### GetLoginRedirectUrl

`func (o *ApplicationAccessibility) GetLoginRedirectUrl() string`

GetLoginRedirectUrl returns the LoginRedirectUrl field if non-nil, zero value otherwise.

### GetLoginRedirectUrlOk

`func (o *ApplicationAccessibility) GetLoginRedirectUrlOk() (*string, bool)`

GetLoginRedirectUrlOk returns a tuple with the LoginRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginRedirectUrl

`func (o *ApplicationAccessibility) SetLoginRedirectUrl(v string)`

SetLoginRedirectUrl sets LoginRedirectUrl field to given value.

### HasLoginRedirectUrl

`func (o *ApplicationAccessibility) HasLoginRedirectUrl() bool`

HasLoginRedirectUrl returns a boolean if a field has been set.

### GetSelfService

`func (o *ApplicationAccessibility) GetSelfService() bool`

GetSelfService returns the SelfService field if non-nil, zero value otherwise.

### GetSelfServiceOk

`func (o *ApplicationAccessibility) GetSelfServiceOk() (*bool, bool)`

GetSelfServiceOk returns a tuple with the SelfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfService

`func (o *ApplicationAccessibility) SetSelfService(v bool)`

SetSelfService sets SelfService field to given value.

### HasSelfService

`func (o *ApplicationAccessibility) HasSelfService() bool`

HasSelfService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


