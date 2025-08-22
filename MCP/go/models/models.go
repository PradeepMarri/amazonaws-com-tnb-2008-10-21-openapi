package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// InstantiateSolNetworkInstanceInput represents the InstantiateSolNetworkInstanceInput schema from the OpenAPI specification
type InstantiateSolNetworkInstanceInput struct {
	Additionalparamsforns interface{} `json:"additionalParamsForNs,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// NetworkArtifactMeta represents the NetworkArtifactMeta schema from the OpenAPI specification
type NetworkArtifactMeta struct {
	Overrides interface{} `json:"overrides,omitempty"`
}

// GetSolNetworkPackageMetadata represents the GetSolNetworkPackageMetadata schema from the OpenAPI specification
type GetSolNetworkPackageMetadata struct {
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
	Nsd interface{} `json:"nsd,omitempty"`
}

// PutSolNetworkPackageContentInput represents the PutSolNetworkPackageContentInput schema from the OpenAPI specification
type PutSolNetworkPackageContentInput struct {
	File interface{} `json:"file"`
}

// ListSolNetworkOperationsOutput represents the ListSolNetworkOperationsOutput schema from the OpenAPI specification
type ListSolNetworkOperationsOutput struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Networkoperations interface{} `json:"networkOperations,omitempty"`
}

// UpdateSolFunctionPackageInput represents the UpdateSolFunctionPackageInput schema from the OpenAPI specification
type UpdateSolFunctionPackageInput struct {
	Operationalstate interface{} `json:"operationalState"`
}

// GetSolVnfInfo represents the GetSolVnfInfo schema from the OpenAPI specification
type GetSolVnfInfo struct {
	Vnfstate interface{} `json:"vnfState,omitempty"`
	Vnfcresourceinfo interface{} `json:"vnfcResourceInfo,omitempty"`
}

// GetSolNetworkInstanceOutput represents the GetSolNetworkInstanceOutput schema from the OpenAPI specification
type GetSolNetworkInstanceOutput struct {
	Arn interface{} `json:"arn"`
	Lcmopinfo LcmOperationInfo `json:"lcmOpInfo,omitempty"` // <p>Lifecycle management operation details on the network instance.</p> <p>Lifecycle management operations are deploy, update, or delete operations.</p>
	Nsinstancedescription interface{} `json:"nsInstanceDescription"`
	Metadata GetSolNetworkInstanceMetadata `json:"metadata"` // <p>The metadata of a network instance.</p> <p>A network instance is a single network created in Amazon Web Services TNB that can be deployed and on which life-cycle operations (like terminate, update, and delete) can be performed.</p>
	Nsinstancename interface{} `json:"nsInstanceName"`
	Nsstate interface{} `json:"nsState,omitempty"`
	Nsdid interface{} `json:"nsdId"`
	Id interface{} `json:"id"`
	Nsdinfoid interface{} `json:"nsdInfoId"`
	Tags interface{} `json:"tags,omitempty"`
}

// InstantiateSolNetworkInstanceOutput represents the InstantiateSolNetworkInstanceOutput schema from the OpenAPI specification
type InstantiateSolNetworkInstanceOutput struct {
	Nslcmopoccid interface{} `json:"nsLcmOpOccId"`
	Tags interface{} `json:"tags,omitempty"`
}

// GetSolFunctionPackageMetadata represents the GetSolFunctionPackageMetadata schema from the OpenAPI specification
type GetSolFunctionPackageMetadata struct {
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
	Vnfd interface{} `json:"vnfd,omitempty"`
}

// GetSolNetworkInstanceInput represents the GetSolNetworkInstanceInput schema from the OpenAPI specification
type GetSolNetworkInstanceInput struct {
}

// StringMap represents the StringMap schema from the OpenAPI specification
type StringMap struct {
}

// PutSolNetworkPackageContentOutput represents the PutSolNetworkPackageContentOutput schema from the OpenAPI specification
type PutSolNetworkPackageContentOutput struct {
	Id interface{} `json:"id"`
	Metadata interface{} `json:"metadata"`
	Nsdid interface{} `json:"nsdId"`
	Nsdname interface{} `json:"nsdName"`
	Nsdversion interface{} `json:"nsdVersion"`
	Vnfpkgids interface{} `json:"vnfPkgIds"`
	Arn interface{} `json:"arn"`
}

// PutSolFunctionPackageContentInput represents the PutSolFunctionPackageContentInput schema from the OpenAPI specification
type PutSolFunctionPackageContentInput struct {
	File interface{} `json:"file"`
}

// GetSolVnfcResourceInfoMetadata represents the GetSolVnfcResourceInfoMetadata schema from the OpenAPI specification
type GetSolVnfcResourceInfoMetadata struct {
	Cluster interface{} `json:"cluster,omitempty"`
	Helmchart interface{} `json:"helmChart,omitempty"`
	Nodegroup interface{} `json:"nodeGroup,omitempty"`
}

// ListSolNetworkOperationsMetadata represents the ListSolNetworkOperationsMetadata schema from the OpenAPI specification
type ListSolNetworkOperationsMetadata struct {
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
}

// Document represents the Document schema from the OpenAPI specification
type Document struct {
}

// ValidateSolFunctionPackageContentOutput represents the ValidateSolFunctionPackageContentOutput schema from the OpenAPI specification
type ValidateSolFunctionPackageContentOutput struct {
	Vnfproductname interface{} `json:"vnfProductName"`
	Vnfprovider interface{} `json:"vnfProvider"`
	Vnfdid interface{} `json:"vnfdId"`
	Vnfdversion interface{} `json:"vnfdVersion"`
	Id interface{} `json:"id"`
	Metadata interface{} `json:"metadata"`
}

// UpdateSolNetworkModify represents the UpdateSolNetworkModify schema from the OpenAPI specification
type UpdateSolNetworkModify struct {
	Vnfconfigurableproperties interface{} `json:"vnfConfigurableProperties"`
	Vnfinstanceid interface{} `json:"vnfInstanceId"`
}

// CancelSolNetworkOperationInput represents the CancelSolNetworkOperationInput schema from the OpenAPI specification
type CancelSolNetworkOperationInput struct {
}

// UntagResourceInput represents the UntagResourceInput schema from the OpenAPI specification
type UntagResourceInput struct {
}

// GetSolNetworkPackageContentInput represents the GetSolNetworkPackageContentInput schema from the OpenAPI specification
type GetSolNetworkPackageContentInput struct {
}

// ListSolNetworkPackageInfo represents the ListSolNetworkPackageInfo schema from the OpenAPI specification
type ListSolNetworkPackageInfo struct {
	Nsdusagestate interface{} `json:"nsdUsageState"`
	Id interface{} `json:"id"`
	Nsdinvariantid interface{} `json:"nsdInvariantId,omitempty"`
	Nsdonboardingstate interface{} `json:"nsdOnboardingState"`
	Nsdversion interface{} `json:"nsdVersion,omitempty"`
	Vnfpkgids interface{} `json:"vnfPkgIds,omitempty"`
	Arn interface{} `json:"arn"`
	Nsdid interface{} `json:"nsdId,omitempty"`
	Nsdname interface{} `json:"nsdName,omitempty"`
	Metadata interface{} `json:"metadata"`
	Nsddesigner interface{} `json:"nsdDesigner,omitempty"`
	Nsdoperationalstate interface{} `json:"nsdOperationalState"`
}

// GetSolInstantiatedVnfInfo represents the GetSolInstantiatedVnfInfo schema from the OpenAPI specification
type GetSolInstantiatedVnfInfo struct {
	Vnfstate interface{} `json:"vnfState,omitempty"`
}

// TerminateSolNetworkInstanceOutput represents the TerminateSolNetworkInstanceOutput schema from the OpenAPI specification
type TerminateSolNetworkInstanceOutput struct {
	Nslcmopoccid interface{} `json:"nsLcmOpOccId,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// ListSolNetworkOperationsInput represents the ListSolNetworkOperationsInput schema from the OpenAPI specification
type ListSolNetworkOperationsInput struct {
}

// FunctionArtifactMeta represents the FunctionArtifactMeta schema from the OpenAPI specification
type FunctionArtifactMeta struct {
	Overrides interface{} `json:"overrides,omitempty"`
}

// GetSolFunctionPackageDescriptorInput represents the GetSolFunctionPackageDescriptorInput schema from the OpenAPI specification
type GetSolFunctionPackageDescriptorInput struct {
}

// ValidateSolFunctionPackageContentMetadata represents the ValidateSolFunctionPackageContentMetadata schema from the OpenAPI specification
type ValidateSolFunctionPackageContentMetadata struct {
	Vnfd FunctionArtifactMeta `json:"vnfd,omitempty"` // <p>Metadata for function package artifacts.</p> <p>Artifacts are the contents of the package descriptor file and the state of the package.</p>
}

// ListSolNetworkInstancesOutput represents the ListSolNetworkInstancesOutput schema from the OpenAPI specification
type ListSolNetworkInstancesOutput struct {
	Networkinstances interface{} `json:"networkInstances,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListSolFunctionInstanceMetadata represents the ListSolFunctionInstanceMetadata schema from the OpenAPI specification
type ListSolFunctionInstanceMetadata struct {
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
}

// ListSolFunctionPackagesOutput represents the ListSolFunctionPackagesOutput schema from the OpenAPI specification
type ListSolFunctionPackagesOutput struct {
	Functionpackages interface{} `json:"functionPackages"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListTagsForResourceInput represents the ListTagsForResourceInput schema from the OpenAPI specification
type ListTagsForResourceInput struct {
}

// GetSolFunctionPackageContentOutput represents the GetSolFunctionPackageContentOutput schema from the OpenAPI specification
type GetSolFunctionPackageContentOutput struct {
	Packagecontent interface{} `json:"packageContent,omitempty"`
}

// GetSolFunctionPackageContentInput represents the GetSolFunctionPackageContentInput schema from the OpenAPI specification
type GetSolFunctionPackageContentInput struct {
}

// ToscaOverride represents the ToscaOverride schema from the OpenAPI specification
type ToscaOverride struct {
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// UntagResourceOutput represents the UntagResourceOutput schema from the OpenAPI specification
type UntagResourceOutput struct {
}

// ListSolNetworkInstanceMetadata represents the ListSolNetworkInstanceMetadata schema from the OpenAPI specification
type ListSolNetworkInstanceMetadata struct {
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
}

// GetSolVnfcResourceInfo represents the GetSolVnfcResourceInfo schema from the OpenAPI specification
type GetSolVnfcResourceInfo struct {
	Metadata interface{} `json:"metadata,omitempty"`
}

// UpdateSolNetworkInstanceOutput represents the UpdateSolNetworkInstanceOutput schema from the OpenAPI specification
type UpdateSolNetworkInstanceOutput struct {
	Tags interface{} `json:"tags,omitempty"`
	Nslcmopoccid interface{} `json:"nsLcmOpOccId,omitempty"`
}

// GetSolNetworkOperationMetadata represents the GetSolNetworkOperationMetadata schema from the OpenAPI specification
type GetSolNetworkOperationMetadata struct {
	Lastmodified interface{} `json:"lastModified"`
	Createdat interface{} `json:"createdAt"`
}

// GetSolNetworkOperationInput represents the GetSolNetworkOperationInput schema from the OpenAPI specification
type GetSolNetworkOperationInput struct {
}

// CreateSolFunctionPackageInput represents the CreateSolFunctionPackageInput schema from the OpenAPI specification
type CreateSolFunctionPackageInput struct {
	Tags interface{} `json:"tags,omitempty"`
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// ValidateSolFunctionPackageContentInput represents the ValidateSolFunctionPackageContentInput schema from the OpenAPI specification
type ValidateSolFunctionPackageContentInput struct {
	File interface{} `json:"file"`
}

// ValidateSolNetworkPackageContentMetadata represents the ValidateSolNetworkPackageContentMetadata schema from the OpenAPI specification
type ValidateSolNetworkPackageContentMetadata struct {
	Nsd NetworkArtifactMeta `json:"nsd,omitempty"` // <p>Metadata for network package artifacts.</p> <p>Artifacts are the contents of the package descriptor file and the state of the package.</p>
}

// ValidateSolNetworkPackageContentInput represents the ValidateSolNetworkPackageContentInput schema from the OpenAPI specification
type ValidateSolNetworkPackageContentInput struct {
	File interface{} `json:"file"`
}

// ProblemDetails represents the ProblemDetails schema from the OpenAPI specification
type ProblemDetails struct {
	Detail interface{} `json:"detail"`
	Title interface{} `json:"title,omitempty"`
}

// PutSolFunctionPackageContentOutput represents the PutSolFunctionPackageContentOutput schema from the OpenAPI specification
type PutSolFunctionPackageContentOutput struct {
	Vnfdversion interface{} `json:"vnfdVersion"`
	Id interface{} `json:"id"`
	Metadata interface{} `json:"metadata"`
	Vnfproductname interface{} `json:"vnfProductName"`
	Vnfprovider interface{} `json:"vnfProvider"`
	Vnfdid interface{} `json:"vnfdId"`
}

// ListTagsForResourceOutput represents the ListTagsForResourceOutput schema from the OpenAPI specification
type ListTagsForResourceOutput struct {
	Tags interface{} `json:"tags"`
}

// GetSolFunctionInstanceInput represents the GetSolFunctionInstanceInput schema from the OpenAPI specification
type GetSolFunctionInstanceInput struct {
}

// CreateSolNetworkPackageInput represents the CreateSolNetworkPackageInput schema from the OpenAPI specification
type CreateSolNetworkPackageInput struct {
	Tags interface{} `json:"tags,omitempty"`
}

// ListSolFunctionInstancesOutput represents the ListSolFunctionInstancesOutput schema from the OpenAPI specification
type ListSolFunctionInstancesOutput struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Functioninstances interface{} `json:"functionInstances,omitempty"`
}

// UpdateSolNetworkInstanceInput represents the UpdateSolNetworkInstanceInput schema from the OpenAPI specification
type UpdateSolNetworkInstanceInput struct {
	Tags interface{} `json:"tags,omitempty"`
	Updatetype interface{} `json:"updateType"`
	Modifyvnfinfodata interface{} `json:"modifyVnfInfoData,omitempty"`
}

// ListSolFunctionPackageMetadata represents the ListSolFunctionPackageMetadata schema from the OpenAPI specification
type ListSolFunctionPackageMetadata struct {
	Lastmodified interface{} `json:"lastModified"`
	Createdat interface{} `json:"createdAt"`
}

// UpdateSolNetworkPackageOutput represents the UpdateSolNetworkPackageOutput schema from the OpenAPI specification
type UpdateSolNetworkPackageOutput struct {
	Nsdoperationalstate interface{} `json:"nsdOperationalState"`
}

// GetSolFunctionInstanceOutput represents the GetSolFunctionInstanceOutput schema from the OpenAPI specification
type GetSolFunctionInstanceOutput struct {
	Vnfpkgid interface{} `json:"vnfPkgId"`
	Vnfprovider interface{} `json:"vnfProvider,omitempty"`
	Arn interface{} `json:"arn"`
	Vnfproductname interface{} `json:"vnfProductName,omitempty"`
	Vnfdversion interface{} `json:"vnfdVersion,omitempty"`
	Instantiatedvnfinfo GetSolVnfInfo `json:"instantiatedVnfInfo,omitempty"` // <p>Information about the network function.</p> <p>A network function instance is a function in a function package .</p>
	Instantiationstate interface{} `json:"instantiationState"`
	Metadata GetSolFunctionInstanceMetadata `json:"metadata"` // <p>The metadata of a network function instance.</p> <p>A network function instance is a function in a function package .</p>
	Nsinstanceid interface{} `json:"nsInstanceId"`
	Vnfdid interface{} `json:"vnfdId"`
	Id interface{} `json:"id"`
	Tags interface{} `json:"tags,omitempty"`
}

// UpdateSolFunctionPackageOutput represents the UpdateSolFunctionPackageOutput schema from the OpenAPI specification
type UpdateSolFunctionPackageOutput struct {
	Operationalstate interface{} `json:"operationalState"`
}

// CreateSolFunctionPackageOutput represents the CreateSolFunctionPackageOutput schema from the OpenAPI specification
type CreateSolFunctionPackageOutput struct {
	Onboardingstate interface{} `json:"onboardingState"`
	Operationalstate interface{} `json:"operationalState"`
	Tags interface{} `json:"tags,omitempty"`
	Usagestate interface{} `json:"usageState"`
	Arn interface{} `json:"arn"`
	Id interface{} `json:"id"`
}

// GetSolFunctionPackageOutput represents the GetSolFunctionPackageOutput schema from the OpenAPI specification
type GetSolFunctionPackageOutput struct {
	Arn interface{} `json:"arn"`
	Usagestate interface{} `json:"usageState"`
	Vnfprovider interface{} `json:"vnfProvider,omitempty"`
	Vnfdversion interface{} `json:"vnfdVersion,omitempty"`
	Metadata GetSolFunctionPackageMetadata `json:"metadata,omitempty"` // <p>Metadata related to the function package.</p> <p>A function package is a .zip file in CSAR (Cloud Service Archive) format that contains a network function (an ETSI standard telecommunication application) and function package descriptor that uses the TOSCA standard to describe how the network functions should run on your network.</p>
	Operationalstate interface{} `json:"operationalState"`
	Tags interface{} `json:"tags,omitempty"`
	Id interface{} `json:"id"`
	Onboardingstate interface{} `json:"onboardingState"`
	Vnfproductname interface{} `json:"vnfProductName,omitempty"`
	Vnfdid interface{} `json:"vnfdId,omitempty"`
}

// PutSolNetworkPackageContentMetadata represents the PutSolNetworkPackageContentMetadata schema from the OpenAPI specification
type PutSolNetworkPackageContentMetadata struct {
	Nsd NetworkArtifactMeta `json:"nsd,omitempty"` // <p>Metadata for network package artifacts.</p> <p>Artifacts are the contents of the package descriptor file and the state of the package.</p>
}

// UpdateSolNetworkPackageInput represents the UpdateSolNetworkPackageInput schema from the OpenAPI specification
type UpdateSolNetworkPackageInput struct {
	Nsdoperationalstate interface{} `json:"nsdOperationalState"`
}

// LcmOperationInfo represents the LcmOperationInfo schema from the OpenAPI specification
type LcmOperationInfo struct {
	Nslcmopoccid interface{} `json:"nsLcmOpOccId"`
}

// GetSolNetworkPackageContentOutput represents the GetSolNetworkPackageContentOutput schema from the OpenAPI specification
type GetSolNetworkPackageContentOutput struct {
	Nsdcontent interface{} `json:"nsdContent,omitempty"`
}

// DeleteSolNetworkInstanceInput represents the DeleteSolNetworkInstanceInput schema from the OpenAPI specification
type DeleteSolNetworkInstanceInput struct {
}

// TagResourceInput represents the TagResourceInput schema from the OpenAPI specification
type TagResourceInput struct {
	Tags interface{} `json:"tags"`
}

// GetSolNetworkOperationTaskDetails represents the GetSolNetworkOperationTaskDetails schema from the OpenAPI specification
type GetSolNetworkOperationTaskDetails struct {
	Taskstatus interface{} `json:"taskStatus,omitempty"`
	Taskcontext interface{} `json:"taskContext,omitempty"`
	Taskendtime interface{} `json:"taskEndTime,omitempty"`
	Taskerrordetails interface{} `json:"taskErrorDetails,omitempty"`
	Taskname interface{} `json:"taskName,omitempty"`
	Taskstarttime interface{} `json:"taskStartTime,omitempty"`
}

// DeleteSolNetworkPackageInput represents the DeleteSolNetworkPackageInput schema from the OpenAPI specification
type DeleteSolNetworkPackageInput struct {
}

// GetSolNetworkPackageInput represents the GetSolNetworkPackageInput schema from the OpenAPI specification
type GetSolNetworkPackageInput struct {
}

// PutSolFunctionPackageContentMetadata represents the PutSolFunctionPackageContentMetadata schema from the OpenAPI specification
type PutSolFunctionPackageContentMetadata struct {
	Vnfd FunctionArtifactMeta `json:"vnfd,omitempty"` // <p>Metadata for function package artifacts.</p> <p>Artifacts are the contents of the package descriptor file and the state of the package.</p>
}

// TagResourceOutput represents the TagResourceOutput schema from the OpenAPI specification
type TagResourceOutput struct {
}

// CreateSolNetworkInstanceInput represents the CreateSolNetworkInstanceInput schema from the OpenAPI specification
type CreateSolNetworkInstanceInput struct {
	Nsdescription interface{} `json:"nsDescription,omitempty"`
	Nsname interface{} `json:"nsName"`
	Nsdinfoid interface{} `json:"nsdInfoId"`
	Tags interface{} `json:"tags,omitempty"`
}

// TerminateSolNetworkInstanceInput represents the TerminateSolNetworkInstanceInput schema from the OpenAPI specification
type TerminateSolNetworkInstanceInput struct {
	Tags interface{} `json:"tags,omitempty"`
}

// CreateSolNetworkPackageOutput represents the CreateSolNetworkPackageOutput schema from the OpenAPI specification
type CreateSolNetworkPackageOutput struct {
	Nsdonboardingstate interface{} `json:"nsdOnboardingState"`
	Nsdoperationalstate interface{} `json:"nsdOperationalState"`
	Nsdusagestate interface{} `json:"nsdUsageState"`
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn"`
	Id interface{} `json:"id"`
}

// ListSolNetworkPackagesInput represents the ListSolNetworkPackagesInput schema from the OpenAPI specification
type ListSolNetworkPackagesInput struct {
}

// ValidateSolNetworkPackageContentOutput represents the ValidateSolNetworkPackageContentOutput schema from the OpenAPI specification
type ValidateSolNetworkPackageContentOutput struct {
	Metadata interface{} `json:"metadata"`
	Nsdid interface{} `json:"nsdId"`
	Nsdname interface{} `json:"nsdName"`
	Nsdversion interface{} `json:"nsdVersion"`
	Vnfpkgids interface{} `json:"vnfPkgIds"`
	Arn interface{} `json:"arn"`
	Id interface{} `json:"id"`
}

// ListSolFunctionInstanceInfo represents the ListSolFunctionInstanceInfo schema from the OpenAPI specification
type ListSolFunctionInstanceInfo struct {
	Id interface{} `json:"id"`
	Instantiatedvnfinfo GetSolInstantiatedVnfInfo `json:"instantiatedVnfInfo,omitempty"` // <p>Information about a network function.</p> <p>A network instance is a single network created in Amazon Web Services TNB that can be deployed and on which life-cycle operations (like terminate, update, and delete) can be performed.</p>
	Instantiationstate interface{} `json:"instantiationState"`
	Metadata interface{} `json:"metadata"`
	Nsinstanceid interface{} `json:"nsInstanceId"`
	Vnfpkgid interface{} `json:"vnfPkgId"`
	Vnfpkgname interface{} `json:"vnfPkgName,omitempty"`
	Arn interface{} `json:"arn"`
}

// ErrorInfo represents the ErrorInfo schema from the OpenAPI specification
type ErrorInfo struct {
	Cause interface{} `json:"cause,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

// GetSolNetworkPackageOutput represents the GetSolNetworkPackageOutput schema from the OpenAPI specification
type GetSolNetworkPackageOutput struct {
	Nsdversion interface{} `json:"nsdVersion"`
	Id interface{} `json:"id"`
	Nsdonboardingstate interface{} `json:"nsdOnboardingState"`
	Nsdname interface{} `json:"nsdName"`
	Nsdoperationalstate interface{} `json:"nsdOperationalState"`
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn"`
	Metadata GetSolNetworkPackageMetadata `json:"metadata"` // <p>Metadata associated with a network package.</p> <p>A network package is a .zip file in CSAR (Cloud Service Archive) format defines the function packages you want to deploy and the Amazon Web Services infrastructure you want to deploy them on.</p>
	Vnfpkgids interface{} `json:"vnfPkgIds"`
	Nsdid interface{} `json:"nsdId"`
	Nsdusagestate interface{} `json:"nsdUsageState"`
}

// DeleteSolFunctionPackageInput represents the DeleteSolFunctionPackageInput schema from the OpenAPI specification
type DeleteSolFunctionPackageInput struct {
}

// GetSolNetworkPackageDescriptorOutput represents the GetSolNetworkPackageDescriptorOutput schema from the OpenAPI specification
type GetSolNetworkPackageDescriptorOutput struct {
	Nsd interface{} `json:"nsd,omitempty"`
}

// ListSolNetworkPackagesOutput represents the ListSolNetworkPackagesOutput schema from the OpenAPI specification
type ListSolNetworkPackagesOutput struct {
	Networkpackages interface{} `json:"networkPackages"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// GetSolNetworkInstanceMetadata represents the GetSolNetworkInstanceMetadata schema from the OpenAPI specification
type GetSolNetworkInstanceMetadata struct {
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
}

// CreateSolNetworkInstanceOutput represents the CreateSolNetworkInstanceOutput schema from the OpenAPI specification
type CreateSolNetworkInstanceOutput struct {
	Id interface{} `json:"id"`
	Nsinstancename interface{} `json:"nsInstanceName"`
	Nsdinfoid interface{} `json:"nsdInfoId"`
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn"`
}

// GetSolFunctionPackageInput represents the GetSolFunctionPackageInput schema from the OpenAPI specification
type GetSolFunctionPackageInput struct {
}

// GetSolNetworkOperationOutput represents the GetSolNetworkOperationOutput schema from the OpenAPI specification
type GetSolNetworkOperationOutput struct {
	ErrorField interface{} `json:"error,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lcmoperationtype interface{} `json:"lcmOperationType,omitempty"`
	Nsinstanceid interface{} `json:"nsInstanceId,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Tasks interface{} `json:"tasks,omitempty"`
	Operationstate interface{} `json:"operationState,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	Arn interface{} `json:"arn"`
}

// GetSolFunctionInstanceMetadata represents the GetSolFunctionInstanceMetadata schema from the OpenAPI specification
type GetSolFunctionInstanceMetadata struct {
	Lastmodified interface{} `json:"lastModified"`
	Createdat interface{} `json:"createdAt"`
}

// ListSolNetworkInstanceInfo represents the ListSolNetworkInstanceInfo schema from the OpenAPI specification
type ListSolNetworkInstanceInfo struct {
	Nsstate interface{} `json:"nsState"`
	Nsdid interface{} `json:"nsdId"`
	Nsdinfoid interface{} `json:"nsdInfoId"`
	Arn interface{} `json:"arn"`
	Id interface{} `json:"id"`
	Metadata interface{} `json:"metadata"`
	Nsinstancedescription interface{} `json:"nsInstanceDescription"`
	Nsinstancename interface{} `json:"nsInstanceName"`
}

// GetSolFunctionPackageDescriptorOutput represents the GetSolFunctionPackageDescriptorOutput schema from the OpenAPI specification
type GetSolFunctionPackageDescriptorOutput struct {
	Vnfd interface{} `json:"vnfd,omitempty"`
}

// ListSolNetworkPackageMetadata represents the ListSolNetworkPackageMetadata schema from the OpenAPI specification
type ListSolNetworkPackageMetadata struct {
	Createdat interface{} `json:"createdAt"`
	Lastmodified interface{} `json:"lastModified"`
}

// GetSolNetworkPackageDescriptorInput represents the GetSolNetworkPackageDescriptorInput schema from the OpenAPI specification
type GetSolNetworkPackageDescriptorInput struct {
}

// ListSolNetworkInstancesInput represents the ListSolNetworkInstancesInput schema from the OpenAPI specification
type ListSolNetworkInstancesInput struct {
}

// ListSolFunctionPackagesInput represents the ListSolFunctionPackagesInput schema from the OpenAPI specification
type ListSolFunctionPackagesInput struct {
}

// ListSolFunctionPackageInfo represents the ListSolFunctionPackageInfo schema from the OpenAPI specification
type ListSolFunctionPackageInfo struct {
	Operationalstate interface{} `json:"operationalState"`
	Vnfproductname interface{} `json:"vnfProductName,omitempty"`
	Id interface{} `json:"id"`
	Metadata interface{} `json:"metadata,omitempty"`
	Vnfprovider interface{} `json:"vnfProvider,omitempty"`
	Vnfdversion interface{} `json:"vnfdVersion,omitempty"`
	Arn interface{} `json:"arn"`
	Onboardingstate interface{} `json:"onboardingState"`
	Usagestate interface{} `json:"usageState"`
	Vnfdid interface{} `json:"vnfdId,omitempty"`
}

// ListSolFunctionInstancesInput represents the ListSolFunctionInstancesInput schema from the OpenAPI specification
type ListSolFunctionInstancesInput struct {
}

// ListSolNetworkOperationsInfo represents the ListSolNetworkOperationsInfo schema from the OpenAPI specification
type ListSolNetworkOperationsInfo struct {
	Lcmoperationtype interface{} `json:"lcmOperationType"`
	Metadata interface{} `json:"metadata,omitempty"`
	Nsinstanceid interface{} `json:"nsInstanceId"`
	Operationstate interface{} `json:"operationState"`
	Arn interface{} `json:"arn"`
	ErrorField interface{} `json:"error,omitempty"`
	Id interface{} `json:"id"`
}
