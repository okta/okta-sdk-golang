# \PolicyAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivatePolicy**](PolicyAPI.md#ActivatePolicy) | **Post** /api/v1/policies/{policyId}/lifecycle/activate | Activate a policy
[**ActivatePolicyRule**](PolicyAPI.md#ActivatePolicyRule) | **Post** /api/v1/policies/{policyId}/rules/{ruleId}/lifecycle/activate | Activate a policy rule
[**ClonePolicy**](PolicyAPI.md#ClonePolicy) | **Post** /api/v1/policies/{policyId}/clone | Clone an existing policy
[**CreatePolicy**](PolicyAPI.md#CreatePolicy) | **Post** /api/v1/policies | Create a policy
[**CreatePolicyRule**](PolicyAPI.md#CreatePolicyRule) | **Post** /api/v1/policies/{policyId}/rules | Create a policy rule
[**CreatePolicySimulation**](PolicyAPI.md#CreatePolicySimulation) | **Post** /api/v1/policies/simulate | Create a policy simulation
[**DeactivatePolicy**](PolicyAPI.md#DeactivatePolicy) | **Post** /api/v1/policies/{policyId}/lifecycle/deactivate | Deactivate a policy
[**DeactivatePolicyRule**](PolicyAPI.md#DeactivatePolicyRule) | **Post** /api/v1/policies/{policyId}/rules/{ruleId}/lifecycle/deactivate | Deactivate a policy rule
[**DeletePolicy**](PolicyAPI.md#DeletePolicy) | **Delete** /api/v1/policies/{policyId} | Delete a policy
[**DeletePolicyResourceMapping**](PolicyAPI.md#DeletePolicyResourceMapping) | **Delete** /api/v1/policies/{policyId}/mappings/{mappingId} | Delete a policy resource mapping
[**DeletePolicyRule**](PolicyAPI.md#DeletePolicyRule) | **Delete** /api/v1/policies/{policyId}/rules/{ruleId} | Delete a policy rule
[**GetPolicy**](PolicyAPI.md#GetPolicy) | **Get** /api/v1/policies/{policyId} | Retrieve a policy
[**GetPolicyMapping**](PolicyAPI.md#GetPolicyMapping) | **Get** /api/v1/policies/{policyId}/mappings/{mappingId} | Retrieve a policy resource mapping
[**GetPolicyRule**](PolicyAPI.md#GetPolicyRule) | **Get** /api/v1/policies/{policyId}/rules/{ruleId} | Retrieve a policy rule
[**ListPolicies**](PolicyAPI.md#ListPolicies) | **Get** /api/v1/policies | List all policies
[**ListPolicyApps**](PolicyAPI.md#ListPolicyApps) | **Get** /api/v1/policies/{policyId}/app | List all apps mapped to a policy
[**ListPolicyMappings**](PolicyAPI.md#ListPolicyMappings) | **Get** /api/v1/policies/{policyId}/mappings | List all resources mapped to a policy
[**ListPolicyRules**](PolicyAPI.md#ListPolicyRules) | **Get** /api/v1/policies/{policyId}/rules | List all policy rules
[**MapResourceToPolicy**](PolicyAPI.md#MapResourceToPolicy) | **Post** /api/v1/policies/{policyId}/mappings | Map a resource to a policy
[**ReplacePolicy**](PolicyAPI.md#ReplacePolicy) | **Put** /api/v1/policies/{policyId} | Replace a policy
[**ReplacePolicyRule**](PolicyAPI.md#ReplacePolicyRule) | **Put** /api/v1/policies/{policyId}/rules/{ruleId} | Replace a policy rule



## ActivatePolicy

> ActivatePolicy(ctx, policyId).Execute()

Activate a policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PolicyAPI.ActivatePolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.ActivatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivatePolicyRule

> ActivatePolicyRule(ctx, policyId, ruleId).Execute()

Activate a policy rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
	ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the policy rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PolicyAPI.ActivatePolicyRule(context.Background(), policyId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.ActivatePolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the policy rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivatePolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClonePolicy

> ListPolicies200ResponseInner ClonePolicy(ctx, policyId).Execute()

Clone an existing policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.ClonePolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.ClonePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClonePolicy`: ListPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.ClonePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiClonePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicy

> CreatePolicyRequest CreatePolicy(ctx).Policy(policy).Activate(activate).Execute()

Create a policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policy := openapiclient.createPolicy_request{AccessPolicy: openapiclient.NewAccessPolicy("Name_example", "Type_example")} // CreatePolicyRequest | 
	activate := true // bool | This query parameter is only valid for Classic Engine orgs. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.CreatePolicy(context.Background()).Policy(policy).Activate(activate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.CreatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePolicy`: CreatePolicyRequest
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.CreatePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policy** | [**CreatePolicyRequest**](CreatePolicyRequest.md) |  | 
 **activate** | **bool** | This query parameter is only valid for Classic Engine orgs. | [default to true]

### Return type

[**CreatePolicyRequest**](CreatePolicyRequest.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicyRule

> ListPolicyRules200ResponseInner CreatePolicyRule(ctx, policyId).PolicyRule(policyRule).Limit(limit).Activate(activate).Execute()

Create a policy rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
	policyRule := openapiclient.listPolicyRules_200_response_inner{AccessPolicyRule: openapiclient.NewAccessPolicyRule()} // ListPolicyRules200ResponseInner | 
	limit := "limit_example" // string | Defines the number of policy rules returned. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	activate := true // bool | Set this parameter to `false` to create an `INACTIVE` rule. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.CreatePolicyRule(context.Background(), policyId).PolicyRule(policyRule).Limit(limit).Activate(activate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.CreatePolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePolicyRule`: ListPolicyRules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.CreatePolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyRule** | [**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md) |  | 
 **limit** | **string** | Defines the number of policy rules returned. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 
 **activate** | **bool** | Set this parameter to &#x60;false&#x60; to create an &#x60;INACTIVE&#x60; rule. | [default to true]

### Return type

[**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicySimulation

> []SimulatePolicyEvaluations CreatePolicySimulation(ctx).SimulatePolicy(simulatePolicy).Expand(expand).Execute()

Create a policy simulation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	simulatePolicy := []openapiclient.SimulatePolicyBody{*openapiclient.NewSimulatePolicyBody("AppInstance_example")} // []SimulatePolicyBody | 
	expand := "EVALUATED" // string | Use `expand=EVALUATED` to include a list of evaluated but not matched policies and policy rules. Use `expand=RULE` to include details about why a rule condition wasn't matched. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.CreatePolicySimulation(context.Background()).SimulatePolicy(simulatePolicy).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.CreatePolicySimulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePolicySimulation`: []SimulatePolicyEvaluations
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.CreatePolicySimulation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicySimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **simulatePolicy** | [**[]SimulatePolicyBody**](SimulatePolicyBody.md) |  | 
 **expand** | **string** | Use &#x60;expand&#x3D;EVALUATED&#x60; to include a list of evaluated but not matched policies and policy rules. Use &#x60;expand&#x3D;RULE&#x60; to include details about why a rule condition wasn&#39;t matched. | 

### Return type

[**[]SimulatePolicyEvaluations**](SimulatePolicyEvaluations.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivatePolicy

> DeactivatePolicy(ctx, policyId).Execute()

Deactivate a policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PolicyAPI.DeactivatePolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.DeactivatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivatePolicyRule

> DeactivatePolicyRule(ctx, policyId, ruleId).Execute()

Deactivate a policy rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
	ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the policy rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PolicyAPI.DeactivatePolicyRule(context.Background(), policyId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.DeactivatePolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the policy rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivatePolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> DeletePolicy(ctx, policyId).Execute()

Delete a policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PolicyAPI.DeletePolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.DeletePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicyResourceMapping

> DeletePolicyResourceMapping(ctx, policyId, mappingId).Execute()

Delete a policy resource mapping



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
	mappingId := "maplr2rLjZ6NsGn1P0g3" // string | `id` of the policy resource Mapping

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PolicyAPI.DeletePolicyResourceMapping(context.Background(), policyId, mappingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.DeletePolicyResourceMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**mappingId** | **string** | &#x60;id&#x60; of the policy resource Mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyResourceMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicyRule

> DeletePolicyRule(ctx, policyId, ruleId).Execute()

Delete a policy rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
	ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the policy rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PolicyAPI.DeletePolicyRule(context.Background(), policyId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.DeletePolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the policy rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicy

> ListPolicies200ResponseInner GetPolicy(ctx, policyId).Expand(expand).Execute()

Retrieve a policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
	expand := "expand_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.GetPolicy(context.Background(), policyId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.GetPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicy`: ListPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.GetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** |  | [default to &quot;&quot;]

### Return type

[**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyMapping

> PolicyMapping GetPolicyMapping(ctx, policyId, mappingId).Execute()

Retrieve a policy resource mapping



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
	mappingId := "maplr2rLjZ6NsGn1P0g3" // string | `id` of the policy resource Mapping

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.GetPolicyMapping(context.Background(), policyId, mappingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.GetPolicyMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyMapping`: PolicyMapping
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.GetPolicyMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**mappingId** | **string** | &#x60;id&#x60; of the policy resource Mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PolicyMapping**](PolicyMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyRule

> ListPolicyRules200ResponseInner GetPolicyRule(ctx, policyId, ruleId).Execute()

Retrieve a policy rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
	ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the policy rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.GetPolicyRule(context.Background(), policyId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.GetPolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyRule`: ListPolicyRules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.GetPolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the policy rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> []ListPolicies200ResponseInner ListPolicies(ctx).Type_(type_).Status(status).Q(q).Expand(expand).SortBy(sortBy).Limit(limit).ResourceId(resourceId).After(after).Execute()

List all policies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	type_ := "type__example" // string | Specifies the type of policy to return. The following policy types are available only with the Okta Identity Engine - `ACCESS_POLICY`, <x-lifecycle class=\"ea\"></x-lifecycle> `DEVICE_SIGNAL_COLLECTION`, `PROFILE_ENROLLMENT`, `POST_AUTH_SESSION` and `ENTITY_RISK`.
	status := "status_example" // string | Refines the query by the `status` of the policy - `ACTIVE` or `INACTIVE` (optional)
	q := "q_example" // string | Refines the query by policy name prefix (startWith method) passed in as `q=string` (optional)
	expand := "expand_example" // string |  (optional) (default to "")
	sortBy := "sortBy_example" // string | Refines the query by sorting on the policy `name` in ascending order (optional)
	limit := "limit_example" // string | Defines the number of policies returned, see [Pagination](https://developer.okta.com/docs/api/#pagination) (optional)
	resourceId := "resourceId_example" // string | Reference to the associated authorization server (optional)
	after := "after_example" // string | End page cursor for pagination, see [Pagination](https://developer.okta.com/docs/api/#pagination) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.ListPolicies(context.Background()).Type_(type_).Status(status).Q(q).Expand(expand).SortBy(sortBy).Limit(limit).ResourceId(resourceId).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.ListPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPolicies`: []ListPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.ListPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Specifies the type of policy to return. The following policy types are available only with the Okta Identity Engine - &#x60;ACCESS_POLICY&#x60;, &lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; &#x60;DEVICE_SIGNAL_COLLECTION&#x60;, &#x60;PROFILE_ENROLLMENT&#x60;, &#x60;POST_AUTH_SESSION&#x60; and &#x60;ENTITY_RISK&#x60;. | 
 **status** | **string** | Refines the query by the &#x60;status&#x60; of the policy - &#x60;ACTIVE&#x60; or &#x60;INACTIVE&#x60; | 
 **q** | **string** | Refines the query by policy name prefix (startWith method) passed in as &#x60;q&#x3D;string&#x60; | 
 **expand** | **string** |  | [default to &quot;&quot;]
 **sortBy** | **string** | Refines the query by sorting on the policy &#x60;name&#x60; in ascending order | 
 **limit** | **string** | Defines the number of policies returned, see [Pagination](https://developer.okta.com/docs/api/#pagination) | 
 **resourceId** | **string** | Reference to the associated authorization server | 
 **after** | **string** | End page cursor for pagination, see [Pagination](https://developer.okta.com/docs/api/#pagination) | 

### Return type

[**[]ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicyApps

> []ListApplications200ResponseInner ListPolicyApps(ctx, policyId).Execute()

List all apps mapped to a policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.ListPolicyApps(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.ListPolicyApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPolicyApps`: []ListApplications200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.ListPolicyApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListApplications200ResponseInner**](ListApplications200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicyMappings

> []PolicyMapping ListPolicyMappings(ctx, policyId).Execute()

List all resources mapped to a policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.ListPolicyMappings(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.ListPolicyMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPolicyMappings`: []PolicyMapping
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.ListPolicyMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PolicyMapping**](PolicyMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicyRules

> []ListPolicyRules200ResponseInner ListPolicyRules(ctx, policyId).Limit(limit).Execute()

List all policy rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
	limit := "limit_example" // string | Defines the number of policy rules returned. See [Pagination](https://developer.okta.com/docs/api/#pagination). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.ListPolicyRules(context.Background(), policyId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.ListPolicyRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPolicyRules`: []ListPolicyRules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.ListPolicyRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **string** | Defines the number of policy rules returned. See [Pagination](https://developer.okta.com/docs/api/#pagination). | 

### Return type

[**[]ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MapResourceToPolicy

> PolicyMapping MapResourceToPolicy(ctx, policyId).PolicyMappingRequest(policyMappingRequest).Execute()

Map a resource to a policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
	policyMappingRequest := *openapiclient.NewPolicyMappingRequest() // PolicyMappingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.MapResourceToPolicy(context.Background(), policyId).PolicyMappingRequest(policyMappingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.MapResourceToPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MapResourceToPolicy`: PolicyMapping
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.MapResourceToPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiMapResourceToPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyMappingRequest** | [**PolicyMappingRequest**](PolicyMappingRequest.md) |  | 

### Return type

[**PolicyMapping**](PolicyMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplacePolicy

> ListPolicies200ResponseInner ReplacePolicy(ctx, policyId).Policy(policy).Execute()

Replace a policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
	policy := openapiclient.createPolicy_request{AccessPolicy: openapiclient.NewAccessPolicy("Name_example", "Type_example")} // CreatePolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.ReplacePolicy(context.Background(), policyId).Policy(policy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.ReplacePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplacePolicy`: ListPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.ReplacePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policy** | [**CreatePolicyRequest**](CreatePolicyRequest.md) |  | 

### Return type

[**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplacePolicyRule

> ListPolicyRules200ResponseInner ReplacePolicyRule(ctx, policyId, ruleId).PolicyRule(policyRule).Execute()

Replace a policy rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
	policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
	ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the policy rule
	policyRule := openapiclient.listPolicyRules_200_response_inner{AccessPolicyRule: openapiclient.NewAccessPolicyRule()} // ListPolicyRules200ResponseInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.ReplacePolicyRule(context.Background(), policyId, ruleId).PolicyRule(policyRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.ReplacePolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplacePolicyRule`: ListPolicyRules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.ReplacePolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the policy rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyRule** | [**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md) |  | 

### Return type

[**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

