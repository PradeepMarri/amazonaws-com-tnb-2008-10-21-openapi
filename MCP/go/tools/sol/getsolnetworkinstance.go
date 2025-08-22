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

func GetsolnetworkinstanceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		nsInstanceIdVal, ok := args["nsInstanceId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: nsInstanceId"), nil
		}
		nsInstanceId, ok := nsInstanceIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: nsInstanceId"), nil
		}
		url := fmt.Sprintf("%s/sol/nslcm/v1/ns_instances/%s", cfg.BaseURL, nsInstanceId)
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
		var result models.GetSolNetworkInstanceOutput
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

func CreateGetsolnetworkinstanceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_sol_nslcm_v1_ns_instances_nsInstanceId",
		mcp.WithDescription("<p>Gets the details of the network instance.</p> <p>A network instance is a single network created in Amazon Web Services TNB that can be deployed and on which life-cycle operations (like terminate, update, and delete) can be performed.</p>"),
		mcp.WithString("nsInstanceId", mcp.Required(), mcp.Description("ID of the network instance.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetsolnetworkinstanceHandler(cfg),
	}
}
