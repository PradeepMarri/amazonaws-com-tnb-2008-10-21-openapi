package main

import (
	"github.com/aws-telco-network-builder/mcp-server/config"
	"github.com/aws-telco-network-builder/mcp-server/models"
	tools_sol "github.com/aws-telco-network-builder/mcp-server/tools/sol"
	tools_tags "github.com/aws-telco-network-builder/mcp-server/tools/tags"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_sol.CreateValidatesolfunctionpackagecontentTool(cfg),
		tools_sol.CreatePutsolfunctionpackagecontentTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_sol.CreateListsolfunctioninstancesTool(cfg),
		tools_sol.CreateListsolnetworkpackagesTool(cfg),
		tools_sol.CreateCreatesolnetworkpackageTool(cfg),
		tools_sol.CreateDeletesolfunctionpackageTool(cfg),
		tools_sol.CreateGetsolfunctionpackageTool(cfg),
		tools_sol.CreateUpdatesolfunctionpackageTool(cfg),
		tools_sol.CreateGetsolfunctionpackagedescriptorTool(cfg),
		tools_sol.CreateGetsolfunctioninstanceTool(cfg),
		tools_sol.CreateListsolnetworkinstancesTool(cfg),
		tools_sol.CreateCreatesolnetworkinstanceTool(cfg),
		tools_sol.CreateValidatesolnetworkpackagecontentTool(cfg),
		tools_sol.CreateDeletesolnetworkinstanceTool(cfg),
		tools_sol.CreateGetsolnetworkinstanceTool(cfg),
		tools_sol.CreateCreatesolfunctionpackageTool(cfg),
		tools_sol.CreateListsolfunctionpackagesTool(cfg),
		tools_sol.CreateUpdatesolnetworkpackageTool(cfg),
		tools_sol.CreateDeletesolnetworkpackageTool(cfg),
		tools_sol.CreateGetsolnetworkpackageTool(cfg),
		tools_sol.CreateCancelsolnetworkoperationTool(cfg),
		tools_sol.CreatePutsolnetworkpackagecontentTool(cfg),
		tools_sol.CreateGetsolnetworkpackagecontentTool(cfg),
		tools_sol.CreateTerminatesolnetworkinstanceTool(cfg),
		tools_sol.CreateGetsolfunctionpackagecontentTool(cfg),
		tools_sol.CreateGetsolnetworkpackagedescriptorTool(cfg),
		tools_sol.CreateInstantiatesolnetworkinstanceTool(cfg),
		tools_sol.CreateUpdatesolnetworkinstanceTool(cfg),
		tools_sol.CreateGetsolnetworkoperationTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_sol.CreateListsolnetworkoperationsTool(cfg),
	}
}
