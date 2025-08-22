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

func GetsolfunctionpackagedescriptorHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		vnfPkgIdVal, ok := args["vnfPkgId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: vnfPkgId"), nil
		}
		vnfPkgId, ok := vnfPkgIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: vnfPkgId"), nil
		}
		url := fmt.Sprintf("%s/sol/vnfpkgm/v1/vnf_packages/%s/vnfd#Accept", cfg.BaseURL, vnfPkgId)
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
		if val, ok := args["Accept"]; ok {
			req.Header.Set("Accept", fmt.Sprintf("%v", val))
		}

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
		var result models.GetSolFunctionPackageDescriptorOutput
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

func CreateGetsolfunctionpackagedescriptorTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_sol_vnfpkgm_v1_vnf_packages_vnfPkgId_vnfd#Accept",
		mcp.WithDescription("<p>Gets a function package descriptor in a function package.</p> <p>A function package descriptor is a .yaml file in a function package that uses the TOSCA standard to describe how the network function in the function package should run on your network.</p> <p>A function package is a .zip file in CSAR (Cloud Service Archive) format that contains a network function (an ETSI standard telecommunication application) and function package descriptor that uses the TOSCA standard to describe how the network functions should run on your network.</p>"),
		mcp.WithString("Accept", mcp.Required(), mcp.Description("Indicates which content types, expressed as MIME types, the client is able to understand.")),
		mcp.WithString("vnfPkgId", mcp.Required(), mcp.Description("ID of the function package.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetsolfunctionpackagedescriptorHandler(cfg),
	}
}
