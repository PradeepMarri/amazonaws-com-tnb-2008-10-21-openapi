package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws-telco-network-builder/mcp-server/config"
	"github.com/aws-telco-network-builder/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UntagresourceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		resourceArnVal, ok := args["resourceArn"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resourceArn"), nil
		}
		resourceArn, ok := resourceArnVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resourceArn"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["tagKeys"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tagKeys=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/tags/%s#tagKeys%s", cfg.BaseURL, resourceArn, queryString)
		req, err := http.NewRequest("DELETE", url, nil)
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
		var result models.UntagResourceOutput
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

func CreateUntagresourceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_tags_resourceArn#tagKeys",
		mcp.WithDescription("<p>Untags an AWS TNB resource.</p> <p>A tag is a label that you assign to an Amazon Web Services resource. Each tag consists of a key and an optional value. You can use tags to search and filter your resources or track your Amazon Web Services costs.</p>"),
		mcp.WithString("resourceArn", mcp.Required(), mcp.Description("Resource ARN.")),
		mcp.WithArray("tagKeys", mcp.Required(), mcp.Description("Tag keys.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UntagresourceHandler(cfg),
	}
}
