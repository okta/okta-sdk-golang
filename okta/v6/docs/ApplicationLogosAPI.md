# \ApplicationLogosAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadApplicationLogo**](ApplicationLogosAPI.md#UploadApplicationLogo) | **Post** /api/v1/apps/{appId}/logo | Upload an application logo



## UploadApplicationLogo

> UploadApplicationLogo(ctx, appId).File(file).Execute()

Upload an application logo



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
	appId := "0oafxqCAJWWGELFTYASJ" // string | Application ID
	file := os.NewFile(1234, "some_file") // *os.File | The image file containing the logo.  The file must be in PNG, JPG, SVG, or GIF format, and less than one MB in size. For best results, use an image with a transparent background and a square dimension of 200 x 200 pixels to prevent upscaling.  > **Notes:** > * Only SVG files encoded in UTF-8 are supported. For example, `<xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\">` is a valid SVG file declaration. > * `multipart/form-data` isn't supported for Python. Remove the `\\\"Content-Type\\\": \\\"multipart/form-data\\\"` line if you use the Python request sample code. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationLogosAPI.UploadApplicationLogo(context.Background(), appId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLogosAPI.UploadApplicationLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadApplicationLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The image file containing the logo.  The file must be in PNG, JPG, SVG, or GIF format, and less than one MB in size. For best results, use an image with a transparent background and a square dimension of 200 x 200 pixels to prevent upscaling.  &gt; **Notes:** &gt; * Only SVG files encoded in UTF-8 are supported. For example, &#x60;&lt;xml version&#x3D;\\\&quot;1.0\\\&quot; encoding&#x3D;\\\&quot;UTF-8\\\&quot;&gt;&#x60; is a valid SVG file declaration. &gt; * &#x60;multipart/form-data&#x60; isn&#39;t supported for Python. Remove the &#x60;\\\&quot;Content-Type\\\&quot;: \\\&quot;multipart/form-data\\\&quot;&#x60; line if you use the Python request sample code.  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

