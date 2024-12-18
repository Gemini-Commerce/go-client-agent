/*
agent/service.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package agent

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type AgentAPI interface {

	/*
		AgentCreateAgent Method for AgentCreateAgent

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param tenantId
		@return ApiAgentCreateAgentRequest
	*/
	AgentCreateAgent(ctx context.Context, tenantId string) ApiAgentCreateAgentRequest

	// AgentCreateAgentExecute executes the request
	//  @return AgentAgentEntity
	AgentCreateAgentExecute(r ApiAgentCreateAgentRequest) (*AgentAgentEntity, *http.Response, error)

	/*
		AgentGetAgent Method for AgentGetAgent

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param tenantId
		@param id
		@return ApiAgentGetAgentRequest
	*/
	AgentGetAgent(ctx context.Context, tenantId string, id string) ApiAgentGetAgentRequest

	// AgentGetAgentExecute executes the request
	//  @return AgentAgentEntity
	AgentGetAgentExecute(r ApiAgentGetAgentRequest) (*AgentAgentEntity, *http.Response, error)

	/*
		AgentListAgents Method for AgentListAgents

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param tenantId
		@param pageSize
		@return ApiAgentListAgentsRequest
	*/
	AgentListAgents(ctx context.Context, tenantId string, pageSize int64) ApiAgentListAgentsRequest

	// AgentListAgentsExecute executes the request
	//  @return AgentListResponse
	AgentListAgentsExecute(r ApiAgentListAgentsRequest) (*AgentListResponse, *http.Response, error)

	/*
		AgentUpdateAgent Method for AgentUpdateAgent

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param tenantId
		@param id
		@return ApiAgentUpdateAgentRequest
	*/
	AgentUpdateAgent(ctx context.Context, tenantId string, id string) ApiAgentUpdateAgentRequest

	// AgentUpdateAgentExecute executes the request
	//  @return AgentAgentEntity
	AgentUpdateAgentExecute(r ApiAgentUpdateAgentRequest) (*AgentAgentEntity, *http.Response, error)
}

// AgentAPIService AgentAPI service
type AgentAPIService service

type ApiAgentCreateAgentRequest struct {
	ctx        context.Context
	ApiService AgentAPI
	tenantId   string
	body       *AgentCreateAgentRequest
}

func (r ApiAgentCreateAgentRequest) Body(body AgentCreateAgentRequest) ApiAgentCreateAgentRequest {
	r.body = &body
	return r
}

func (r ApiAgentCreateAgentRequest) Execute() (*AgentAgentEntity, *http.Response, error) {
	return r.ApiService.AgentCreateAgentExecute(r)
}

/*
AgentCreateAgent Method for AgentCreateAgent

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId
	@return ApiAgentCreateAgentRequest
*/
func (a *AgentAPIService) AgentCreateAgent(ctx context.Context, tenantId string) ApiAgentCreateAgentRequest {
	return ApiAgentCreateAgentRequest{
		ApiService: a,
		ctx:        ctx,
		tenantId:   tenantId,
	}
}

// Execute executes the request
//
//	@return AgentAgentEntity
func (a *AgentAPIService) AgentCreateAgentExecute(r ApiAgentCreateAgentRequest) (*AgentAgentEntity, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AgentAgentEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentAPIService.AgentCreateAgent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/{tenantId}/agent/create"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantId"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v RpcStatus
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAgentGetAgentRequest struct {
	ctx        context.Context
	ApiService AgentAPI
	tenantId   string
	id         string
}

func (r ApiAgentGetAgentRequest) Execute() (*AgentAgentEntity, *http.Response, error) {
	return r.ApiService.AgentGetAgentExecute(r)
}

/*
AgentGetAgent Method for AgentGetAgent

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId
	@param id
	@return ApiAgentGetAgentRequest
*/
func (a *AgentAPIService) AgentGetAgent(ctx context.Context, tenantId string, id string) ApiAgentGetAgentRequest {
	return ApiAgentGetAgentRequest{
		ApiService: a,
		ctx:        ctx,
		tenantId:   tenantId,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AgentAgentEntity
func (a *AgentAPIService) AgentGetAgentExecute(r ApiAgentGetAgentRequest) (*AgentAgentEntity, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AgentAgentEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentAPIService.AgentGetAgent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/{tenantId}/agent/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantId"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v RpcStatus
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAgentListAgentsRequest struct {
	ctx        context.Context
	ApiService AgentAPI
	tenantId   string
	pageSize   int64
	body       *AgentListAgentsRequest
}

func (r ApiAgentListAgentsRequest) Body(body AgentListAgentsRequest) ApiAgentListAgentsRequest {
	r.body = &body
	return r
}

func (r ApiAgentListAgentsRequest) Execute() (*AgentListResponse, *http.Response, error) {
	return r.ApiService.AgentListAgentsExecute(r)
}

/*
AgentListAgents Method for AgentListAgents

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId
	@param pageSize
	@return ApiAgentListAgentsRequest
*/
func (a *AgentAPIService) AgentListAgents(ctx context.Context, tenantId string, pageSize int64) ApiAgentListAgentsRequest {
	return ApiAgentListAgentsRequest{
		ApiService: a,
		ctx:        ctx,
		tenantId:   tenantId,
		pageSize:   pageSize,
	}
}

// Execute executes the request
//
//	@return AgentListResponse
func (a *AgentAPIService) AgentListAgentsExecute(r ApiAgentListAgentsRequest) (*AgentListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AgentListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentAPIService.AgentListAgents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/{tenantId}/agent/list/page-size/{pageSize}"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantId"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageSize"+"}", url.PathEscape(parameterValueToString(r.pageSize, "pageSize")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v RpcStatus
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAgentUpdateAgentRequest struct {
	ctx        context.Context
	ApiService AgentAPI
	tenantId   string
	id         string
	body       *AgentUpdateAgentRequest
}

func (r ApiAgentUpdateAgentRequest) Body(body AgentUpdateAgentRequest) ApiAgentUpdateAgentRequest {
	r.body = &body
	return r
}

func (r ApiAgentUpdateAgentRequest) Execute() (*AgentAgentEntity, *http.Response, error) {
	return r.ApiService.AgentUpdateAgentExecute(r)
}

/*
AgentUpdateAgent Method for AgentUpdateAgent

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId
	@param id
	@return ApiAgentUpdateAgentRequest
*/
func (a *AgentAPIService) AgentUpdateAgent(ctx context.Context, tenantId string, id string) ApiAgentUpdateAgentRequest {
	return ApiAgentUpdateAgentRequest{
		ApiService: a,
		ctx:        ctx,
		tenantId:   tenantId,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AgentAgentEntity
func (a *AgentAPIService) AgentUpdateAgentExecute(r ApiAgentUpdateAgentRequest) (*AgentAgentEntity, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AgentAgentEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentAPIService.AgentUpdateAgent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/{tenantId}/agent/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantId"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Authorization"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v RpcStatus
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
