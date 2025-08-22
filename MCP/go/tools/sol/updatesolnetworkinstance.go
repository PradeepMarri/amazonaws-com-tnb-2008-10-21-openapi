package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/aws-telco-network-builder/mcp-server/config"
	"github.com/aws-telco-network-builder/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UpdatesolnetworkinstanceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/sol/nslcm/v1/ns_instances/%s/update", cfg.BaseURL, nsInstanceId)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.UpdateSolNetworkInstanceOutput
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

func CreateUpdatesolnetworkinstanceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_sol_nslcm_v1_ns_instances_nsInstanceId_update",
		mcp.WithDescription("<p>Update a network instance.</p> <p>A network instance is a single network created in Amazon Web Services TNB that can be deployed and on which life-cycle operations (like terminate, update, and delete) can be performed.</p>"),
		mcp.WithString("nsInstanceId", mcp.Required(), mcp.Description("ID of the network instance.")),
		mcp.WithObject("modifyVnfInfoData", mcp.Description("Input parameter: <p>Information parameters and/or the configurable properties for a network function.</p> <p>A network function instance is a function in a function package .</p>")),
		mcp.WithObject("tags", mcp.Description("Input parameter: A tag is a label that you assign to an Amazon Web Services resource. Each tag consists of a key and an optional value. When you use this API, the tags are transferred to the network operation that is created. Use tags to search and filter your resources or track your Amazon Web Services costs.")),
		mcp.WithString("updateType", mcp.Required(), mcp.Description("Input parameter: The type of update.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdatesolnetworkinstanceHandler(cfg),
	}
}
