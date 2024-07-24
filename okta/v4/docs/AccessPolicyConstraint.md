# AccessPolicyConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationMethods** | Pointer to [**[]AuthenticationMethodObject**](AuthenticationMethodObject.md) | This property specifies the precise authenticator and method for authentication. &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt; | [optional] 
**ExcludedAuthenticationMethods** | Pointer to [**[]AuthenticationMethodObject**](AuthenticationMethodObject.md) | This property specifies the precise authenticator and method to exclude from authentication. &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt; | [optional] 
**Methods** | Pointer to **[]string** | The Authenticator methods that are permitted | [optional] 
**ReauthenticateIn** | Pointer to **string** | The duration after which the user must re-authenticate regardless of user activity. This re-authentication interval overrides the Verification Method object&#39;s &#x60;reauthenticateIn&#x60; interval. The supported values use ISO 8601 period format for recurring time intervals (for example, &#x60;PT1H&#x60;). | [optional] 
**Required** | Pointer to **bool** | This property indicates whether the knowledge or possession factor is required by the assurance. It&#39;s optional in the request, but is always returned in the response. By default, this field is &#x60;true&#x60;. If the knowledge or possession constraint has values for &#x60;excludedAuthenticationMethods&#x60; the &#x60;required&#x60; value is false. &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt; | [optional] 
**Types** | Pointer to **[]string** | The Authenticator types that are permitted | [optional] 

## Methods

### NewAccessPolicyConstraint

`func NewAccessPolicyConstraint() *AccessPolicyConstraint`

NewAccessPolicyConstraint instantiates a new AccessPolicyConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyConstraintWithDefaults

`func NewAccessPolicyConstraintWithDefaults() *AccessPolicyConstraint`

NewAccessPolicyConstraintWithDefaults instantiates a new AccessPolicyConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationMethods

`func (o *AccessPolicyConstraint) GetAuthenticationMethods() []AuthenticationMethodObject`

GetAuthenticationMethods returns the AuthenticationMethods field if non-nil, zero value otherwise.

### GetAuthenticationMethodsOk

`func (o *AccessPolicyConstraint) GetAuthenticationMethodsOk() (*[]AuthenticationMethodObject, bool)`

GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethods

`func (o *AccessPolicyConstraint) SetAuthenticationMethods(v []AuthenticationMethodObject)`

SetAuthenticationMethods sets AuthenticationMethods field to given value.

### HasAuthenticationMethods

`func (o *AccessPolicyConstraint) HasAuthenticationMethods() bool`

HasAuthenticationMethods returns a boolean if a field has been set.

### GetExcludedAuthenticationMethods

`func (o *AccessPolicyConstraint) GetExcludedAuthenticationMethods() []AuthenticationMethodObject`

GetExcludedAuthenticationMethods returns the ExcludedAuthenticationMethods field if non-nil, zero value otherwise.

### GetExcludedAuthenticationMethodsOk

`func (o *AccessPolicyConstraint) GetExcludedAuthenticationMethodsOk() (*[]AuthenticationMethodObject, bool)`

GetExcludedAuthenticationMethodsOk returns a tuple with the ExcludedAuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAuthenticationMethods

`func (o *AccessPolicyConstraint) SetExcludedAuthenticationMethods(v []AuthenticationMethodObject)`

SetExcludedAuthenticationMethods sets ExcludedAuthenticationMethods field to given value.

### HasExcludedAuthenticationMethods

`func (o *AccessPolicyConstraint) HasExcludedAuthenticationMethods() bool`

HasExcludedAuthenticationMethods returns a boolean if a field has been set.

### GetMethods

`func (o *AccessPolicyConstraint) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *AccessPolicyConstraint) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *AccessPolicyConstraint) SetMethods(v []string)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *AccessPolicyConstraint) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetReauthenticateIn

`func (o *AccessPolicyConstraint) GetReauthenticateIn() string`

GetReauthenticateIn returns the ReauthenticateIn field if non-nil, zero value otherwise.

### GetReauthenticateInOk

`func (o *AccessPolicyConstraint) GetReauthenticateInOk() (*string, bool)`

GetReauthenticateInOk returns a tuple with the ReauthenticateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthenticateIn

`func (o *AccessPolicyConstraint) SetReauthenticateIn(v string)`

SetReauthenticateIn sets ReauthenticateIn field to given value.

### HasReauthenticateIn

`func (o *AccessPolicyConstraint) HasReauthenticateIn() bool`

HasReauthenticateIn returns a boolean if a field has been set.

### GetRequired

`func (o *AccessPolicyConstraint) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AccessPolicyConstraint) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AccessPolicyConstraint) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *AccessPolicyConstraint) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetTypes

`func (o *AccessPolicyConstraint) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *AccessPolicyConstraint) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *AccessPolicyConstraint) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *AccessPolicyConstraint) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


