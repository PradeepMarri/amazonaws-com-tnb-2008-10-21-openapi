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

func GetsolnetworkpackageHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		nsdInfoIdVal, ok := args["nsdInfoId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: nsdInfoId"), nil
		}
		nsdInfoId, ok := nsdInfoIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: nsdInfoId"), nil
		}
		url := fmt.Sprintf("%s/sol/nsd/v1/ns_descriptors/%s", cfg.BaseURL, nsdInfoId)
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
		var result models.GetSolNetworkPackageOutput
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

func CreateGetsolnetworkpackageTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_sol_nsd_v1_ns_descriptors_nsdInfoId",
		mcp.WithDescription("<p>Gets the details of a network package.</p> <p>A network package is a .zip file in CSAR (Cloud Service Archive) format defines the function packages you want to deploy and the Amazon Web Services infrastructure you want to deploy them on.</p>"),
		mcp.WithString("nsdInfoId", mcp.Required(), mcp.Description("ID of the network service descriptor in the network package.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetsolnetworkpackageHandler(cfg),
	}
}
