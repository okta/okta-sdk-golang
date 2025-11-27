# \IdentitySourceAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentitySourceSession**](IdentitySourceAPI.md#CreateIdentitySourceSession) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions | Create an identity source session
[**DeleteIdentitySourceSession**](IdentitySourceAPI.md#DeleteIdentitySourceSession) | **Delete** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId} | Delete an identity source session
[**GetIdentitySourceSession**](IdentitySourceAPI.md#GetIdentitySourceSession) | **Get** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId} | Retrieve an identity source session
[**ListIdentitySourceSessions**](IdentitySourceAPI.md#ListIdentitySourceSessions) | **Get** /api/v1/identity-sources/{identitySourceId}/sessions | List all identity source sessions
[**StartImportFromIdentitySource**](IdentitySourceAPI.md#StartImportFromIdentitySource) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/start-import | Start the import from the identity source
[**UploadIdentitySourceDataForDelete**](IdentitySourceAPI.md#UploadIdentitySourceDataForDelete) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-delete | Upload the data to be deleted in Okta
[**UploadIdentitySourceDataForUpsert**](IdentitySourceAPI.md#UploadIdentitySourceDataForUpsert) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-upsert | Upload the data to be upserted in Okta
[**UploadIdentitySourceGroupMembershipsForDelete**](IdentitySourceAPI.md#UploadIdentitySourceGroupMembershipsForDelete) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-group-memberships-delete | Upload the group memberships to be deleted in Okta
[**UploadIdentitySourceGroupMembershipsForUpsert**](IdentitySourceAPI.md#UploadIdentitySourceGroupMembershipsForUpsert) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-group-memberships-upsert | Upload the group memberships to be upserted in Okta
[**UploadIdentitySourceGroupsDataForDelete**](IdentitySourceAPI.md#UploadIdentitySourceGroupsDataForDelete) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-groups-delete | Upload the group external IDs to be deleted in Okta
[**UploadIdentitySourceGroupsForUpsert**](IdentitySourceAPI.md#UploadIdentitySourceGroupsForUpsert) | **Post** /api/v1/identity-sources/{identitySourceId}/sessions/{sessionId}/bulk-groups-upsert | Upload the group profiles without memberships to be upserted in Okta



## CreateIdentitySourceSession

> IdentitySourceSession CreateIdentitySourceSession(ctx, identitySourceId).Execute()

Create an identity source session



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.CreateIdentitySourceSession(context.Background(), identitySourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.CreateIdentitySourceSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIdentitySourceSession`: IdentitySourceSession
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.CreateIdentitySourceSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentitySourceSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentitySourceSession**](IdentitySourceSession.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentitySourceSession

> DeleteIdentitySourceSession(ctx, identitySourceId, sessionId).Execute()

Delete an identity source session



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.DeleteIdentitySourceSession(context.Background(), identitySourceId, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.DeleteIdentitySourceSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentitySourceSessionRequest struct via the builder pattern


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


## GetIdentitySourceSession

> IdentitySourceSession GetIdentitySourceSession(ctx, identitySourceId, sessionId).Execute()

Retrieve an identity source session



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.GetIdentitySourceSession(context.Background(), identitySourceId, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.GetIdentitySourceSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentitySourceSession`: IdentitySourceSession
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.GetIdentitySourceSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitySourceSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdentitySourceSession**](IdentitySourceSession.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentitySourceSessions

> []IdentitySourceSession ListIdentitySourceSessions(ctx, identitySourceId).Execute()

List all identity source sessions



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.ListIdentitySourceSessions(context.Background(), identitySourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.ListIdentitySourceSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdentitySourceSessions`: []IdentitySourceSession
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.ListIdentitySourceSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIdentitySourceSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IdentitySourceSession**](IdentitySourceSession.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartImportFromIdentitySource

> IdentitySourceSession StartImportFromIdentitySource(ctx, identitySourceId, sessionId).Execute()

Start the import from the identity source



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitySourceAPI.StartImportFromIdentitySource(context.Background(), identitySourceId, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.StartImportFromIdentitySource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartImportFromIdentitySource`: IdentitySourceSession
	fmt.Fprintf(os.Stdout, "Response from `IdentitySourceAPI.StartImportFromIdentitySource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartImportFromIdentitySourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdentitySourceSession**](IdentitySourceSession.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdentitySourceDataForDelete

> UploadIdentitySourceDataForDelete(ctx, identitySourceId, sessionId).BulkDeleteRequestBody(bulkDeleteRequestBody).Execute()

Upload the data to be deleted in Okta



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session
	bulkDeleteRequestBody := *openapiclient.NewBulkDeleteRequestBody() // BulkDeleteRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.UploadIdentitySourceDataForDelete(context.Background(), identitySourceId, sessionId).BulkDeleteRequestBody(bulkDeleteRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.UploadIdentitySourceDataForDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadIdentitySourceDataForDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkDeleteRequestBody** | [**BulkDeleteRequestBody**](BulkDeleteRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdentitySourceDataForUpsert

> UploadIdentitySourceDataForUpsert(ctx, identitySourceId, sessionId).BulkUpsertRequestBody(bulkUpsertRequestBody).Execute()

Upload the data to be upserted in Okta



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session
	bulkUpsertRequestBody := *openapiclient.NewBulkUpsertRequestBody() // BulkUpsertRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.UploadIdentitySourceDataForUpsert(context.Background(), identitySourceId, sessionId).BulkUpsertRequestBody(bulkUpsertRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.UploadIdentitySourceDataForUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadIdentitySourceDataForUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkUpsertRequestBody** | [**BulkUpsertRequestBody**](BulkUpsertRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdentitySourceGroupMembershipsForDelete

> UploadIdentitySourceGroupMembershipsForDelete(ctx, identitySourceId, sessionId).BulkGroupMembershipsDeleteRequestBody(bulkGroupMembershipsDeleteRequestBody).Execute()

Upload the group memberships to be deleted in Okta



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session
	bulkGroupMembershipsDeleteRequestBody := *openapiclient.NewBulkGroupMembershipsDeleteRequestBody() // BulkGroupMembershipsDeleteRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.UploadIdentitySourceGroupMembershipsForDelete(context.Background(), identitySourceId, sessionId).BulkGroupMembershipsDeleteRequestBody(bulkGroupMembershipsDeleteRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.UploadIdentitySourceGroupMembershipsForDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadIdentitySourceGroupMembershipsForDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkGroupMembershipsDeleteRequestBody** | [**BulkGroupMembershipsDeleteRequestBody**](BulkGroupMembershipsDeleteRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdentitySourceGroupMembershipsForUpsert

> UploadIdentitySourceGroupMembershipsForUpsert(ctx, identitySourceId, sessionId).BulkGroupMembershipsUpsertRequestBody(bulkGroupMembershipsUpsertRequestBody).Execute()

Upload the group memberships to be upserted in Okta



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session
	bulkGroupMembershipsUpsertRequestBody := *openapiclient.NewBulkGroupMembershipsUpsertRequestBody() // BulkGroupMembershipsUpsertRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.UploadIdentitySourceGroupMembershipsForUpsert(context.Background(), identitySourceId, sessionId).BulkGroupMembershipsUpsertRequestBody(bulkGroupMembershipsUpsertRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.UploadIdentitySourceGroupMembershipsForUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadIdentitySourceGroupMembershipsForUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkGroupMembershipsUpsertRequestBody** | [**BulkGroupMembershipsUpsertRequestBody**](BulkGroupMembershipsUpsertRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdentitySourceGroupsDataForDelete

> UploadIdentitySourceGroupsDataForDelete(ctx, identitySourceId, sessionId).BulkGroupDeleteRequestBody(bulkGroupDeleteRequestBody).Execute()

Upload the group external IDs to be deleted in Okta



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session
	bulkGroupDeleteRequestBody := *openapiclient.NewBulkGroupDeleteRequestBody() // BulkGroupDeleteRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.UploadIdentitySourceGroupsDataForDelete(context.Background(), identitySourceId, sessionId).BulkGroupDeleteRequestBody(bulkGroupDeleteRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.UploadIdentitySourceGroupsDataForDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadIdentitySourceGroupsDataForDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkGroupDeleteRequestBody** | [**BulkGroupDeleteRequestBody**](BulkGroupDeleteRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdentitySourceGroupsForUpsert

> UploadIdentitySourceGroupsForUpsert(ctx, identitySourceId, sessionId).BulkGroupUpsertRequestBody(bulkGroupUpsertRequestBody).Execute()

Upload the group profiles without memberships to be upserted in Okta



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
	identitySourceId := "0oa3l6l6WK6h0R0QW0g4" // string | The ID of the identity source for which the session is created
	sessionId := "aps1qqonvr2SZv6o70h8" // string | The ID of the identity source session
	bulkGroupUpsertRequestBody := *openapiclient.NewBulkGroupUpsertRequestBody() // BulkGroupUpsertRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitySourceAPI.UploadIdentitySourceGroupsForUpsert(context.Background(), identitySourceId, sessionId).BulkGroupUpsertRequestBody(bulkGroupUpsertRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitySourceAPI.UploadIdentitySourceGroupsForUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identitySourceId** | **string** | The ID of the identity source for which the session is created | 
**sessionId** | **string** | The ID of the identity source session | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadIdentitySourceGroupsForUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkGroupUpsertRequestBody** | [**BulkGroupUpsertRequestBody**](BulkGroupUpsertRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

