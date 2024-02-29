# Go API client for agent

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: version not set
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import agent "github.com/Gemini-Commerce/go-client-agent"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `agent.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), agent.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `agent.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), agent.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `agent.ContextOperationServerIndices` and `agent.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), agent.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), agent.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://agent.api.gogemini.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AgentAPI* | [**AgentCreateAgent**](docs/AgentAPI.md#agentcreateagent) | **Post** /v1/{tenantId}/agent/create | 
*AgentAPI* | [**AgentGetAgent**](docs/AgentAPI.md#agentgetagent) | **Get** /v1/{tenantId}/agent/{id} | 
*AgentAPI* | [**AgentListAgents**](docs/AgentAPI.md#agentlistagents) | **Post** /v1/{tenantId}/agent/list/page-size/{pageSize} | 
*AgentAPI* | [**AgentUpdateAgent**](docs/AgentAPI.md#agentupdateagent) | **Put** /v1/{tenantId}/agent/{id} | 


## Documentation For Models

 - [AgentAgentEntity](docs/AgentAgentEntity.md)
 - [AgentCreateAgentRequest](docs/AgentCreateAgentRequest.md)
 - [AgentListAgentsRequest](docs/AgentListAgentsRequest.md)
 - [AgentListResponse](docs/AgentListResponse.md)
 - [AgentSortOrder](docs/AgentSortOrder.md)
 - [AgentUpdateAgentRequest](docs/AgentUpdateAgentRequest.md)
 - [AgentUpdatePayload](docs/AgentUpdatePayload.md)
 - [ListRequestFilters](docs/ListRequestFilters.md)
 - [ListRequestSort](docs/ListRequestSort.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [RpcStatus](docs/RpcStatus.md)
 - [SortSortField](docs/SortSortField.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### Authorization

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		agent.ContextAPIKeys,
		map[string]agent.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### geminiAuthorization

- **Type**: API key
- **API key parameter name**: X-Gem-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Gem-Token and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		agent.ContextAPIKeys,
		map[string]agent.APIKey{
			"X-Gem-Token": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


