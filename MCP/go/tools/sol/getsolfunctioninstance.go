package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aws-telco-network-builder/mcp-server/config"
	"github.com/aws-telco-network-builder/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetsolfunctioninstanceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		vnfInstanceIdVal, ok := args["vnfInstanceId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: vnfInstanceId"), nil
		}
		vnfInstanceId, ok := vnfInstanceIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: vnfInstanceId"), nil
		}
		url := fmt.Sprintf("%s/sol/vnflcm/v1/vnf_instances/%s", cfg.BaseURL, vnfInstanceId)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.GetSolFunctionInstanceOutput
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetsolfunctioninstanceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_sol_vnflcm_v1_vnf_instances_vnfInstanceId",
		mcp.WithDescription("<p>Gets the details of a network function instance, including the instantation state and metadata from the function package descriptor in the network function package.</p> <p>A network function instance is a function in a function package .</p>"),
		mcp.WithString("vnfInstanceId", mcp.Required(), mcp.Description("ID of the network function.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetsolfunctioninstanceHandler(cfg),
	}
}
