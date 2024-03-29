/*
agent/service.proto

Testing AgentAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package agent

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/Gemini-Commerce/go-client-agent"
)

func Test_agent_AgentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AgentAPIService AgentCreateAgent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string

		resp, httpRes, err := apiClient.AgentAPI.AgentCreateAgent(context.Background(), tenantId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentAPIService AgentGetAgent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var id string

		resp, httpRes, err := apiClient.AgentAPI.AgentGetAgent(context.Background(), tenantId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentAPIService AgentListAgents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var pageSize int64

		resp, httpRes, err := apiClient.AgentAPI.AgentListAgents(context.Background(), tenantId, pageSize).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AgentAPIService AgentUpdateAgent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenantId string
		var id string

		resp, httpRes, err := apiClient.AgentAPI.AgentUpdateAgent(context.Background(), tenantId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
