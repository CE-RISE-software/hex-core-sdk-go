# \ModelsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRecord**](ModelsAPI.md#CreateRecord) | **Post** /models/{model}/versions/{version}:create | Validate and create a record
[**QueryRecords**](ModelsAPI.md#QueryRecords) | **Post** /models/{model}/versions/{version}:query | Query records for a model version
[**ValidateModelPayload**](ModelsAPI.md#ValidateModelPayload) | **Post** /models/{model}/versions/{version}:validate | Validate payload against model rules



## CreateRecord

> Record CreateRecord(ctx, model, version).IdempotencyKey(idempotencyKey).CreateRequest(createRequest).Execute()

Validate and create a record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CE-RISE-software/hex-core-sdk-go"
)

func main() {
	model := "model_example" // string | Model identifier.
	version := "version_example" // string | Model version.
	idempotencyKey := "idempotencyKey_example" // string | Client-generated key used to deduplicate retries for side-effecting requests.
	createRequest := *openapiclient.NewCreateRequest(map[string]interface{}{"key": interface{}(123)}) // CreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.CreateRecord(context.Background(), model, version).IdempotencyKey(idempotencyKey).CreateRequest(createRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.CreateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRecord`: Record
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.CreateRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model identifier. | 
**version** | **string** | Model version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **string** | Client-generated key used to deduplicate retries for side-effecting requests. | 
 **createRequest** | [**CreateRequest**](CreateRequest.md) |  | 

### Return type

[**Record**](Record.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryRecords

> QueryResponse QueryRecords(ctx, model, version).QueryRequest(queryRequest).Execute()

Query records for a model version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CE-RISE-software/hex-core-sdk-go"
)

func main() {
	model := "model_example" // string | Model identifier.
	version := "version_example" // string | Model version.
	queryRequest := *openapiclient.NewQueryRequest(*openapiclient.NewRecordQueryFilter([]openapiclient.RecordQueryCondition{*openapiclient.NewRecordQueryCondition("Field_example", "Op_example", interface{}(123))})) // QueryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.QueryRecords(context.Background(), model, version).QueryRequest(queryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.QueryRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryRecords`: QueryResponse
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.QueryRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model identifier. | 
**version** | **string** | Model version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **queryRequest** | [**QueryRequest**](QueryRequest.md) |  | 

### Return type

[**QueryResponse**](QueryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateModelPayload

> ValidationReport ValidateModelPayload(ctx, model, version).ValidateRequest(validateRequest).Execute()

Validate payload against model rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/CE-RISE-software/hex-core-sdk-go"
)

func main() {
	model := "model_example" // string | Model identifier, for example `re-indicators-specification`.
	version := "version_example" // string | Model version string, for example `0.0.3`.
	validateRequest := *openapiclient.NewValidateRequest(map[string]interface{}{"key": interface{}(123)}) // ValidateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.ValidateModelPayload(context.Background(), model, version).ValidateRequest(validateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.ValidateModelPayload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateModelPayload`: ValidationReport
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.ValidateModelPayload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** | Model identifier, for example &#x60;re-indicators-specification&#x60;. | 
**version** | **string** | Model version string, for example &#x60;0.0.3&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateModelPayloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateRequest** | [**ValidateRequest**](ValidateRequest.md) |  | 

### Return type

[**ValidationReport**](ValidationReport.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

