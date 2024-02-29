# GeminiCommerce\Agent\AgentAPI

All URIs are relative to *https://agent.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentCreateAgent**](AgentAPI.md#AgentCreateAgent) | **Post** /v1/{tenantId}/agent/create | 
[**AgentGetAgent**](AgentAPI.md#AgentGetAgent) | **Get** /v1/{tenantId}/agent/{id} | 
[**AgentListAgents**](AgentAPI.md#AgentListAgents) | **Post** /v1/{tenantId}/agent/list/page-size/{pageSize} | 
[**AgentUpdateAgent**](AgentAPI.md#AgentUpdateAgent) | **Put** /v1/{tenantId}/agent/{id} | 



## AgentCreateAgent

> AgentAgentEntity AgentCreateAgent(ctx, tenantId).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-agent"
)

func main() {
	tenantId := "tenantId_example" // string | 
	body := *openapiclient.NewAgentCreateAgentRequest() // AgentCreateAgentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.AgentCreateAgent(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.AgentCreateAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentCreateAgent`: AgentAgentEntity
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.AgentCreateAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentCreateAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AgentCreateAgentRequest**](AgentCreateAgentRequest.md) |  | 

### Return type

[**AgentAgentEntity**](AgentAgentEntity.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentGetAgent

> AgentAgentEntity AgentGetAgent(ctx, tenantId, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-agent"
)

func main() {
	tenantId := "tenantId_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.AgentGetAgent(context.Background(), tenantId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.AgentGetAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentGetAgent`: AgentAgentEntity
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.AgentGetAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentGetAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AgentAgentEntity**](AgentAgentEntity.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentListAgents

> AgentListResponse AgentListAgents(ctx, tenantId, pageSize).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-agent"
)

func main() {
	tenantId := "tenantId_example" // string | 
	pageSize := int64(789) // int64 | 
	body := *openapiclient.NewAgentListAgentsRequest() // AgentListAgentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.AgentListAgents(context.Background(), tenantId, pageSize).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.AgentListAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentListAgents`: AgentListResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.AgentListAgents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**pageSize** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentListAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AgentListAgentsRequest**](AgentListAgentsRequest.md) |  | 

### Return type

[**AgentListResponse**](AgentListResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AgentUpdateAgent

> AgentAgentEntity AgentUpdateAgent(ctx, tenantId, id).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-agent"
)

func main() {
	tenantId := "tenantId_example" // string | 
	id := "id_example" // string | 
	body := *openapiclient.NewAgentUpdateAgentRequest() // AgentUpdateAgentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentAPI.AgentUpdateAgent(context.Background(), tenantId, id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentAPI.AgentUpdateAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentUpdateAgent`: AgentAgentEntity
	fmt.Fprintf(os.Stdout, "Response from `AgentAPI.AgentUpdateAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAgentUpdateAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AgentUpdateAgentRequest**](AgentUpdateAgentRequest.md) |  | 

### Return type

[**AgentAgentEntity**](AgentAgentEntity.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

