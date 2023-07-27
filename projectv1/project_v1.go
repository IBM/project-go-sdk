/**
 * (C) Copyright IBM Corp. 2023.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.75.0-726bc7e3-20230713-221716
 */

// Package projectv1 : Operations and models for the ProjectV1 service
package projectv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/project-go-sdk/common"
	"github.com/go-openapi/strfmt"
)

// ProjectV1 : This document is the **REST API specification** for the Projects Service. The Projects service provides
// the capability to manage Infrastructure as Code in IBM Cloud.
//
// API Version: 1.0.0
type ProjectV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://projects.api.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "project"

// ProjectV1Options : Service options
type ProjectV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewProjectV1UsingExternalConfig : constructs an instance of ProjectV1 with passed in options and external configuration.
func NewProjectV1UsingExternalConfig(options *ProjectV1Options) (project *ProjectV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	project, err = NewProjectV1(options)
	if err != nil {
		return
	}

	err = project.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = project.Service.SetServiceURL(options.URL)
	}
	return
}

// NewProjectV1 : constructs an instance of ProjectV1 with passed in options.
func NewProjectV1(options *ProjectV1Options) (service *ProjectV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &ProjectV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "project" suitable for processing requests.
func (project *ProjectV1) Clone() *ProjectV1 {
	if core.IsNil(project) {
		return nil
	}
	clone := *project
	clone.Service = project.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (project *ProjectV1) SetServiceURL(url string) error {
	return project.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (project *ProjectV1) GetServiceURL() string {
	return project.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (project *ProjectV1) SetDefaultHeaders(headers http.Header) {
	project.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (project *ProjectV1) SetEnableGzipCompression(enableGzip bool) {
	project.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (project *ProjectV1) GetEnableGzipCompression() bool {
	return project.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (project *ProjectV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	project.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (project *ProjectV1) DisableRetries() {
	project.Service.DisableRetries()
}

// CreateProject : Create a project
// Create a new project and asynchronously setup the tools to manage it. Add a deployable architecture by customizing
// the configuration. After the changes are validated and approved, deploy the resources that the project configures.
func (project *ProjectV1) CreateProject(createProjectOptions *CreateProjectOptions) (result *ProjectCanonical, response *core.DetailedResponse, err error) {
	return project.CreateProjectWithContext(context.Background(), createProjectOptions)
}

// CreateProjectWithContext is an alternate form of the CreateProject method which supports a Context parameter
func (project *ProjectV1) CreateProjectWithContext(ctx context.Context, createProjectOptions *CreateProjectOptions) (result *ProjectCanonical, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createProjectOptions, "createProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createProjectOptions, "createProjectOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "CreateProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("resource_group", fmt.Sprint(*createProjectOptions.ResourceGroup))
	builder.AddQuery("location", fmt.Sprint(*createProjectOptions.Location))

	body := make(map[string]interface{})
	if createProjectOptions.Name != nil {
		body["name"] = createProjectOptions.Name
	}
	if createProjectOptions.Description != nil {
		body["description"] = createProjectOptions.Description
	}
	if createProjectOptions.DestroyOnDelete != nil {
		body["destroy_on_delete"] = createProjectOptions.DestroyOnDelete
	}
	if createProjectOptions.Configs != nil {
		body["configs"] = createProjectOptions.Configs
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectCanonical)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListProjects : List projects
// List existing projects. Projects are sorted by ID.
func (project *ProjectV1) ListProjects(listProjectsOptions *ListProjectsOptions) (result *ProjectCollection, response *core.DetailedResponse, err error) {
	return project.ListProjectsWithContext(context.Background(), listProjectsOptions)
}

// ListProjectsWithContext is an alternate form of the ListProjects method which supports a Context parameter
func (project *ProjectV1) ListProjectsWithContext(ctx context.Context, listProjectsOptions *ListProjectsOptions) (result *ProjectCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listProjectsOptions, "listProjectsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listProjectsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "ListProjects")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listProjectsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listProjectsOptions.Start))
	}
	if listProjectsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listProjectsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetProject : Get a project
// Get information about a project.
func (project *ProjectV1) GetProject(getProjectOptions *GetProjectOptions) (result *ProjectResponseSummary, response *core.DetailedResponse, err error) {
	return project.GetProjectWithContext(context.Background(), getProjectOptions)
}

// GetProjectWithContext is an alternate form of the GetProject method which supports a Context parameter
func (project *ProjectV1) GetProjectWithContext(ctx context.Context, getProjectOptions *GetProjectOptions) (result *ProjectResponseSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getProjectOptions, "getProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getProjectOptions, "getProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getProjectOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "GetProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectResponseSummary)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateProject : Update a project
// Update a project by the ID.
func (project *ProjectV1) UpdateProject(updateProjectOptions *UpdateProjectOptions) (result *ProjectResponseSummary, response *core.DetailedResponse, err error) {
	return project.UpdateProjectWithContext(context.Background(), updateProjectOptions)
}

// UpdateProjectWithContext is an alternate form of the UpdateProject method which supports a Context parameter
func (project *ProjectV1) UpdateProjectWithContext(ctx context.Context, updateProjectOptions *UpdateProjectOptions) (result *ProjectResponseSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateProjectOptions, "updateProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateProjectOptions, "updateProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateProjectOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "UpdateProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateProjectOptions.Name != nil {
		body["name"] = updateProjectOptions.Name
	}
	if updateProjectOptions.Description != nil {
		body["description"] = updateProjectOptions.Description
	}
	if updateProjectOptions.DestroyOnDelete != nil {
		body["destroy_on_delete"] = updateProjectOptions.DestroyOnDelete
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectResponseSummary)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteProject : Delete a project
// Delete a project document by the ID. A project can only be deleted after deleting all of its artifacts.
func (project *ProjectV1) DeleteProject(deleteProjectOptions *DeleteProjectOptions) (response *core.DetailedResponse, err error) {
	return project.DeleteProjectWithContext(context.Background(), deleteProjectOptions)
}

// DeleteProjectWithContext is an alternate form of the DeleteProject method which supports a Context parameter
func (project *ProjectV1) DeleteProjectWithContext(ctx context.Context, deleteProjectOptions *DeleteProjectOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteProjectOptions, "deleteProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteProjectOptions, "deleteProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteProjectOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "DeleteProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = project.Service.Request(request, nil)

	return
}

// CreateConfig : Add a new configuration
// Add a new configuration to a project.
func (project *ProjectV1) CreateConfig(createConfigOptions *CreateConfigOptions) (result *ProjectConfigVersionResponse, response *core.DetailedResponse, err error) {
	return project.CreateConfigWithContext(context.Background(), createConfigOptions)
}

// CreateConfigWithContext is an alternate form of the CreateConfig method which supports a Context parameter
func (project *ProjectV1) CreateConfigWithContext(ctx context.Context, createConfigOptions *CreateConfigOptions) (result *ProjectConfigVersionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createConfigOptions, "createConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createConfigOptions, "createConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *createConfigOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{project_id}/configs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "CreateConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createConfigOptions.Name != nil {
		body["name"] = createConfigOptions.Name
	}
	if createConfigOptions.LocatorID != nil {
		body["locator_id"] = createConfigOptions.LocatorID
	}
	if createConfigOptions.Labels != nil {
		body["labels"] = createConfigOptions.Labels
	}
	if createConfigOptions.Description != nil {
		body["description"] = createConfigOptions.Description
	}
	if createConfigOptions.Authorizations != nil {
		body["authorizations"] = createConfigOptions.Authorizations
	}
	if createConfigOptions.ComplianceProfile != nil {
		body["compliance_profile"] = createConfigOptions.ComplianceProfile
	}
	if createConfigOptions.Input != nil {
		body["input"] = createConfigOptions.Input
	}
	if createConfigOptions.Setting != nil {
		body["setting"] = createConfigOptions.Setting
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigVersionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListConfigs : List all project configurations
// The collection of configurations that are returned.
func (project *ProjectV1) ListConfigs(listConfigsOptions *ListConfigsOptions) (result *ProjectConfigCollection, response *core.DetailedResponse, err error) {
	return project.ListConfigsWithContext(context.Background(), listConfigsOptions)
}

// ListConfigsWithContext is an alternate form of the ListConfigs method which supports a Context parameter
func (project *ProjectV1) ListConfigsWithContext(ctx context.Context, listConfigsOptions *ListConfigsOptions) (result *ProjectConfigCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listConfigsOptions, "listConfigsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listConfigsOptions, "listConfigsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *listConfigsOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{project_id}/configs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listConfigsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "ListConfigs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetConfig : Get a project configuration
// Returns the specified project configuration in a specific project.
func (project *ProjectV1) GetConfig(getConfigOptions *GetConfigOptions) (result *ProjectConfigGetResponse, response *core.DetailedResponse, err error) {
	return project.GetConfigWithContext(context.Background(), getConfigOptions)
}

// GetConfigWithContext is an alternate form of the GetConfig method which supports a Context parameter
func (project *ProjectV1) GetConfigWithContext(ctx context.Context, getConfigOptions *GetConfigOptions) (result *ProjectConfigGetResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getConfigOptions, "getConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getConfigOptions, "getConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *getConfigOptions.ProjectID,
		"id": *getConfigOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{project_id}/configs/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "GetConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigGetResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateConfig : Update a configuration
// Update a configuration in a project by the ID.
func (project *ProjectV1) UpdateConfig(updateConfigOptions *UpdateConfigOptions) (result *ProjectConfigVersionResponse, response *core.DetailedResponse, err error) {
	return project.UpdateConfigWithContext(context.Background(), updateConfigOptions)
}

// UpdateConfigWithContext is an alternate form of the UpdateConfig method which supports a Context parameter
func (project *ProjectV1) UpdateConfigWithContext(ctx context.Context, updateConfigOptions *UpdateConfigOptions) (result *ProjectConfigVersionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateConfigOptions, "updateConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateConfigOptions, "updateConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *updateConfigOptions.ProjectID,
		"id": *updateConfigOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{project_id}/configs/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "UpdateConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateConfigOptions.LocatorID != nil {
		body["locator_id"] = updateConfigOptions.LocatorID
	}
	if updateConfigOptions.Input != nil {
		body["input"] = updateConfigOptions.Input
	}
	if updateConfigOptions.Setting != nil {
		body["setting"] = updateConfigOptions.Setting
	}
	if updateConfigOptions.Name != nil {
		body["name"] = updateConfigOptions.Name
	}
	if updateConfigOptions.Labels != nil {
		body["labels"] = updateConfigOptions.Labels
	}
	if updateConfigOptions.Description != nil {
		body["description"] = updateConfigOptions.Description
	}
	if updateConfigOptions.Authorizations != nil {
		body["authorizations"] = updateConfigOptions.Authorizations
	}
	if updateConfigOptions.ComplianceProfile != nil {
		body["compliance_profile"] = updateConfigOptions.ComplianceProfile
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigVersionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteConfig : Delete a configuration in a project by ID
// Delete a configuration in a project. Deleting the configuration will also destroy all the resources deployed by the
// configuration if the query parameter `destroy` is specified.
func (project *ProjectV1) DeleteConfig(deleteConfigOptions *DeleteConfigOptions) (result *ProjectConfigDelete, response *core.DetailedResponse, err error) {
	return project.DeleteConfigWithContext(context.Background(), deleteConfigOptions)
}

// DeleteConfigWithContext is an alternate form of the DeleteConfig method which supports a Context parameter
func (project *ProjectV1) DeleteConfigWithContext(ctx context.Context, deleteConfigOptions *DeleteConfigOptions) (result *ProjectConfigDelete, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteConfigOptions, "deleteConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteConfigOptions, "deleteConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *deleteConfigOptions.ProjectID,
		"id": *deleteConfigOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{project_id}/configs/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "DeleteConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if deleteConfigOptions.DraftOnly != nil {
		builder.AddQuery("draft_only", fmt.Sprint(*deleteConfigOptions.DraftOnly))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigDelete)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ForceApprove : Force approve project configuration
// Force approve configuration edits to the main configuration with an approving comment.
func (project *ProjectV1) ForceApprove(forceApproveOptions *ForceApproveOptions) (result *ProjectConfigVersionResponse, response *core.DetailedResponse, err error) {
	return project.ForceApproveWithContext(context.Background(), forceApproveOptions)
}

// ForceApproveWithContext is an alternate form of the ForceApprove method which supports a Context parameter
func (project *ProjectV1) ForceApproveWithContext(ctx context.Context, forceApproveOptions *ForceApproveOptions) (result *ProjectConfigVersionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(forceApproveOptions, "forceApproveOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(forceApproveOptions, "forceApproveOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *forceApproveOptions.ProjectID,
		"id": *forceApproveOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{project_id}/configs/{id}/force_approve`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range forceApproveOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "ForceApprove")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if forceApproveOptions.Comment != nil {
		body["comment"] = forceApproveOptions.Comment
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigVersionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// Approve : Approve and merge a configuration draft
// Approve and merge configuration edits to the main configuration.
func (project *ProjectV1) Approve(approveOptions *ApproveOptions) (result *ProjectConfigVersionResponse, response *core.DetailedResponse, err error) {
	return project.ApproveWithContext(context.Background(), approveOptions)
}

// ApproveWithContext is an alternate form of the Approve method which supports a Context parameter
func (project *ProjectV1) ApproveWithContext(ctx context.Context, approveOptions *ApproveOptions) (result *ProjectConfigVersionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(approveOptions, "approveOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(approveOptions, "approveOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *approveOptions.ProjectID,
		"id": *approveOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{project_id}/configs/{id}/approve`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range approveOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "Approve")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if approveOptions.Comment != nil {
		body["comment"] = approveOptions.Comment
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigVersionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CheckConfig : Run a validation check
// Run a validation check on a given configuration in project. The check includes creating or updating the associated
// schematics workspace with a plan job, running the CRA scans, and cost estimatation.
func (project *ProjectV1) CheckConfig(checkConfigOptions *CheckConfigOptions) (result *ProjectConfigVersionResponse, response *core.DetailedResponse, err error) {
	return project.CheckConfigWithContext(context.Background(), checkConfigOptions)
}

// CheckConfigWithContext is an alternate form of the CheckConfig method which supports a Context parameter
func (project *ProjectV1) CheckConfigWithContext(ctx context.Context, checkConfigOptions *CheckConfigOptions) (result *ProjectConfigVersionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(checkConfigOptions, "checkConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(checkConfigOptions, "checkConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *checkConfigOptions.ProjectID,
		"id": *checkConfigOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{project_id}/configs/{id}/check`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range checkConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "CheckConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if checkConfigOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*checkConfigOptions.XAuthRefreshToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigVersionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// InstallConfig : Deploy a configuration
// Deploy a project's configuration. It's an asynchronous operation that can be tracked using the get project
// configuration API with full metadata.
func (project *ProjectV1) InstallConfig(installConfigOptions *InstallConfigOptions) (result *ProjectConfigVersionResponse, response *core.DetailedResponse, err error) {
	return project.InstallConfigWithContext(context.Background(), installConfigOptions)
}

// InstallConfigWithContext is an alternate form of the InstallConfig method which supports a Context parameter
func (project *ProjectV1) InstallConfigWithContext(ctx context.Context, installConfigOptions *InstallConfigOptions) (result *ProjectConfigVersionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(installConfigOptions, "installConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(installConfigOptions, "installConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *installConfigOptions.ProjectID,
		"id": *installConfigOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{project_id}/configs/{id}/install`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range installConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "InstallConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigVersionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UninstallConfig : Destroy configuration resources
// Destroy a project's configuration resources. The operation destroys all the resources that are deployed with the
// specific configuration. You can track it by using the get project configuration API with full metadata.
func (project *ProjectV1) UninstallConfig(uninstallConfigOptions *UninstallConfigOptions) (response *core.DetailedResponse, err error) {
	return project.UninstallConfigWithContext(context.Background(), uninstallConfigOptions)
}

// UninstallConfigWithContext is an alternate form of the UninstallConfig method which supports a Context parameter
func (project *ProjectV1) UninstallConfigWithContext(ctx context.Context, uninstallConfigOptions *UninstallConfigOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(uninstallConfigOptions, "uninstallConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(uninstallConfigOptions, "uninstallConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *uninstallConfigOptions.ProjectID,
		"id": *uninstallConfigOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{project_id}/configs/{id}/uninstall`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range uninstallConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "UninstallConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = project.Service.Request(request, nil)

	return
}

// ListConfigResources : List the resources deployed by a configuration
// A list of resources deployed by a configuraton.
func (project *ProjectV1) ListConfigResources(listConfigResourcesOptions *ListConfigResourcesOptions) (result *ProjectConfigResourceCollection, response *core.DetailedResponse, err error) {
	return project.ListConfigResourcesWithContext(context.Background(), listConfigResourcesOptions)
}

// ListConfigResourcesWithContext is an alternate form of the ListConfigResources method which supports a Context parameter
func (project *ProjectV1) ListConfigResourcesWithContext(ctx context.Context, listConfigResourcesOptions *ListConfigResourcesOptions) (result *ProjectConfigResourceCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listConfigResourcesOptions, "listConfigResourcesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listConfigResourcesOptions, "listConfigResourcesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *listConfigResourcesOptions.ProjectID,
		"id": *listConfigResourcesOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{project_id}/configs/{id}/resources`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listConfigResourcesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "ListConfigResources")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigResourceCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListConfigVersions : Get a list of versions of a project configuration
// Returns a list of previous and current versions of a project configuration in a specific project.
func (project *ProjectV1) ListConfigVersions(listConfigVersionsOptions *ListConfigVersionsOptions) (result *ProjectConfigVersionSummaryCollection, response *core.DetailedResponse, err error) {
	return project.ListConfigVersionsWithContext(context.Background(), listConfigVersionsOptions)
}

// ListConfigVersionsWithContext is an alternate form of the ListConfigVersions method which supports a Context parameter
func (project *ProjectV1) ListConfigVersionsWithContext(ctx context.Context, listConfigVersionsOptions *ListConfigVersionsOptions) (result *ProjectConfigVersionSummaryCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listConfigVersionsOptions, "listConfigVersionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listConfigVersionsOptions, "listConfigVersionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *listConfigVersionsOptions.ProjectID,
		"id": *listConfigVersionsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{project_id}/configs/{id}/versions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listConfigVersionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "ListConfigVersions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigVersionSummaryCollection)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetConfigVersion : Get a specific version of a project configuration
// Returns a specific version of a project configuration in a specific project.
func (project *ProjectV1) GetConfigVersion(getConfigVersionOptions *GetConfigVersionOptions) (result *ProjectConfigVersionResponse, response *core.DetailedResponse, err error) {
	return project.GetConfigVersionWithContext(context.Background(), getConfigVersionOptions)
}

// GetConfigVersionWithContext is an alternate form of the GetConfigVersion method which supports a Context parameter
func (project *ProjectV1) GetConfigVersionWithContext(ctx context.Context, getConfigVersionOptions *GetConfigVersionOptions) (result *ProjectConfigVersionResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getConfigVersionOptions, "getConfigVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getConfigVersionOptions, "getConfigVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *getConfigVersionOptions.ProjectID,
		"id": *getConfigVersionOptions.ID,
		"version": fmt.Sprint(*getConfigVersionOptions.Version),
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{project_id}/configs/{id}/versions/{version}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConfigVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "GetConfigVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigVersionResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ApproveOptions : The Approve options.
type ApproveOptions struct {
	// The unique project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The unique config ID.
	ID *string `json:"id" validate:"required,ne="`

	// Notes on the project draft action. If this is a forced approve on the draft configuration, a non-empty comment is
	// required.
	Comment *string `json:"comment,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewApproveOptions : Instantiate ApproveOptions
func (*ProjectV1) NewApproveOptions(projectID string, id string) *ApproveOptions {
	return &ApproveOptions{
		ProjectID: core.StringPtr(projectID),
		ID: core.StringPtr(id),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *ApproveOptions) SetProjectID(projectID string) *ApproveOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetID : Allow user to set ID
func (_options *ApproveOptions) SetID(id string) *ApproveOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetComment : Allow user to set Comment
func (_options *ApproveOptions) SetComment(comment string) *ApproveOptions {
	_options.Comment = core.StringPtr(comment)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ApproveOptions) SetHeaders(param map[string]string) *ApproveOptions {
	options.Headers = param
	return options
}

// CheckConfigOptions : The CheckConfig options.
type CheckConfigOptions struct {
	// The unique project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The unique config ID.
	ID *string `json:"id" validate:"required,ne="`

	// The IAM refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCheckConfigOptions : Instantiate CheckConfigOptions
func (*ProjectV1) NewCheckConfigOptions(projectID string, id string) *CheckConfigOptions {
	return &CheckConfigOptions{
		ProjectID: core.StringPtr(projectID),
		ID: core.StringPtr(id),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *CheckConfigOptions) SetProjectID(projectID string) *CheckConfigOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetID : Allow user to set ID
func (_options *CheckConfigOptions) SetID(id string) *CheckConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (_options *CheckConfigOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *CheckConfigOptions {
	_options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CheckConfigOptions) SetHeaders(param map[string]string) *CheckConfigOptions {
	options.Headers = param
	return options
}

// CreateConfigOptions : The CreateConfig options.
type CreateConfigOptions struct {
	// The unique project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The name of the configuration.
	Name *string `json:"name" validate:"required"`

	// A dotted value of catalogID.versionID.
	LocatorID *string `json:"locator_id" validate:"required"`

	// A collection of configuration labels.
	Labels []string `json:"labels,omitempty"`

	// The description of the project configuration.
	Description *string `json:"description,omitempty"`

	// The authorization for a configuration.
	// You can authorize by using a trusted profile or an API key in Secrets Manager.
	Authorizations *ProjectConfigAuth `json:"authorizations,omitempty"`

	// The profile required for compliance.
	ComplianceProfile *ProjectConfigComplianceProfile `json:"compliance_profile,omitempty"`

	// The inputs of a Schematics template property.
	Input []ProjectConfigInputVariable `json:"input,omitempty"`

	// Schematics environment variables to use to deploy the configuration. Settings are only available if they were
	// specified when the configuration was initially created.
	Setting []ProjectConfigSettingCollection `json:"setting,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateConfigOptions : Instantiate CreateConfigOptions
func (*ProjectV1) NewCreateConfigOptions(projectID string, name string, locatorID string) *CreateConfigOptions {
	return &CreateConfigOptions{
		ProjectID: core.StringPtr(projectID),
		Name: core.StringPtr(name),
		LocatorID: core.StringPtr(locatorID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *CreateConfigOptions) SetProjectID(projectID string) *CreateConfigOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateConfigOptions) SetName(name string) *CreateConfigOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetLocatorID : Allow user to set LocatorID
func (_options *CreateConfigOptions) SetLocatorID(locatorID string) *CreateConfigOptions {
	_options.LocatorID = core.StringPtr(locatorID)
	return _options
}

// SetLabels : Allow user to set Labels
func (_options *CreateConfigOptions) SetLabels(labels []string) *CreateConfigOptions {
	_options.Labels = labels
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateConfigOptions) SetDescription(description string) *CreateConfigOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetAuthorizations : Allow user to set Authorizations
func (_options *CreateConfigOptions) SetAuthorizations(authorizations *ProjectConfigAuth) *CreateConfigOptions {
	_options.Authorizations = authorizations
	return _options
}

// SetComplianceProfile : Allow user to set ComplianceProfile
func (_options *CreateConfigOptions) SetComplianceProfile(complianceProfile *ProjectConfigComplianceProfile) *CreateConfigOptions {
	_options.ComplianceProfile = complianceProfile
	return _options
}

// SetInput : Allow user to set Input
func (_options *CreateConfigOptions) SetInput(input []ProjectConfigInputVariable) *CreateConfigOptions {
	_options.Input = input
	return _options
}

// SetSetting : Allow user to set Setting
func (_options *CreateConfigOptions) SetSetting(setting []ProjectConfigSettingCollection) *CreateConfigOptions {
	_options.Setting = setting
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateConfigOptions) SetHeaders(param map[string]string) *CreateConfigOptions {
	options.Headers = param
	return options
}

// CreateProjectOptions : The CreateProject options.
type CreateProjectOptions struct {
	// The resource group where the project's data and tools are created.
	ResourceGroup *string `json:"resource_group" validate:"required"`

	// The location where the project's data and tools are created.
	Location *string `json:"location" validate:"required"`

	// The name of the project.
	Name *string `json:"name" validate:"required"`

	// A brief explanation of the project's use in the configuration of a deployable architecture. It is possible to create
	// a project without providing a description.
	Description *string `json:"description,omitempty"`

	// The policy that indicates whether the resources are destroyed or not when a project is deleted.
	DestroyOnDelete *bool `json:"destroy_on_delete,omitempty"`

	// The project configurations. If configurations are not included, the project resource is persisted without this
	// property.
	Configs []ProjectConfig `json:"configs,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateProjectOptions : Instantiate CreateProjectOptions
func (*ProjectV1) NewCreateProjectOptions(resourceGroup string, location string, name string) *CreateProjectOptions {
	return &CreateProjectOptions{
		ResourceGroup: core.StringPtr(resourceGroup),
		Location: core.StringPtr(location),
		Name: core.StringPtr(name),
	}
}

// SetResourceGroup : Allow user to set ResourceGroup
func (_options *CreateProjectOptions) SetResourceGroup(resourceGroup string) *CreateProjectOptions {
	_options.ResourceGroup = core.StringPtr(resourceGroup)
	return _options
}

// SetLocation : Allow user to set Location
func (_options *CreateProjectOptions) SetLocation(location string) *CreateProjectOptions {
	_options.Location = core.StringPtr(location)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateProjectOptions) SetName(name string) *CreateProjectOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateProjectOptions) SetDescription(description string) *CreateProjectOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetDestroyOnDelete : Allow user to set DestroyOnDelete
func (_options *CreateProjectOptions) SetDestroyOnDelete(destroyOnDelete bool) *CreateProjectOptions {
	_options.DestroyOnDelete = core.BoolPtr(destroyOnDelete)
	return _options
}

// SetConfigs : Allow user to set Configs
func (_options *CreateProjectOptions) SetConfigs(configs []ProjectConfig) *CreateProjectOptions {
	_options.Configs = configs
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateProjectOptions) SetHeaders(param map[string]string) *CreateProjectOptions {
	options.Headers = param
	return options
}

// CumulativeNeedsAttention : CumulativeNeedsAttention struct
type CumulativeNeedsAttention struct {
	// The event name.
	Event *string `json:"event,omitempty"`

	// A unique ID for that individual event.
	EventID *string `json:"event_id,omitempty"`

	// A unique ID for the configuration.
	ConfigID *string `json:"config_id,omitempty"`

	// The version number of the configuration.
	ConfigVersion *int64 `json:"config_version,omitempty"`
}

// UnmarshalCumulativeNeedsAttention unmarshals an instance of CumulativeNeedsAttention from the specified map of raw messages.
func UnmarshalCumulativeNeedsAttention(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CumulativeNeedsAttention)
	err = core.UnmarshalPrimitive(m, "event", &obj.Event)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "event_id", &obj.EventID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "config_id", &obj.ConfigID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "config_version", &obj.ConfigVersion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteConfigOptions : The DeleteConfig options.
type DeleteConfigOptions struct {
	// The unique project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The unique config ID.
	ID *string `json:"id" validate:"required,ne="`

	// The flag to determine if only the draft version should be deleted.
	DraftOnly *bool `json:"draft_only,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteConfigOptions : Instantiate DeleteConfigOptions
func (*ProjectV1) NewDeleteConfigOptions(projectID string, id string) *DeleteConfigOptions {
	return &DeleteConfigOptions{
		ProjectID: core.StringPtr(projectID),
		ID: core.StringPtr(id),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *DeleteConfigOptions) SetProjectID(projectID string) *DeleteConfigOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetID : Allow user to set ID
func (_options *DeleteConfigOptions) SetID(id string) *DeleteConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetDraftOnly : Allow user to set DraftOnly
func (_options *DeleteConfigOptions) SetDraftOnly(draftOnly bool) *DeleteConfigOptions {
	_options.DraftOnly = core.BoolPtr(draftOnly)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteConfigOptions) SetHeaders(param map[string]string) *DeleteConfigOptions {
	options.Headers = param
	return options
}

// DeleteProjectOptions : The DeleteProject options.
type DeleteProjectOptions struct {
	// The unique project ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteProjectOptions : Instantiate DeleteProjectOptions
func (*ProjectV1) NewDeleteProjectOptions(id string) *DeleteProjectOptions {
	return &DeleteProjectOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteProjectOptions) SetID(id string) *DeleteProjectOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteProjectOptions) SetHeaders(param map[string]string) *DeleteProjectOptions {
	options.Headers = param
	return options
}

// ForceApproveOptions : The ForceApprove options.
type ForceApproveOptions struct {
	// The unique project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The unique config ID.
	ID *string `json:"id" validate:"required,ne="`

	// Notes on the project draft action. If this is a forced approve on the draft configuration, a non-empty comment is
	// required.
	Comment *string `json:"comment,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewForceApproveOptions : Instantiate ForceApproveOptions
func (*ProjectV1) NewForceApproveOptions(projectID string, id string) *ForceApproveOptions {
	return &ForceApproveOptions{
		ProjectID: core.StringPtr(projectID),
		ID: core.StringPtr(id),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *ForceApproveOptions) SetProjectID(projectID string) *ForceApproveOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetID : Allow user to set ID
func (_options *ForceApproveOptions) SetID(id string) *ForceApproveOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetComment : Allow user to set Comment
func (_options *ForceApproveOptions) SetComment(comment string) *ForceApproveOptions {
	_options.Comment = core.StringPtr(comment)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ForceApproveOptions) SetHeaders(param map[string]string) *ForceApproveOptions {
	options.Headers = param
	return options
}

// GetConfigOptions : The GetConfig options.
type GetConfigOptions struct {
	// The unique project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The unique config ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetConfigOptions : Instantiate GetConfigOptions
func (*ProjectV1) NewGetConfigOptions(projectID string, id string) *GetConfigOptions {
	return &GetConfigOptions{
		ProjectID: core.StringPtr(projectID),
		ID: core.StringPtr(id),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *GetConfigOptions) SetProjectID(projectID string) *GetConfigOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetID : Allow user to set ID
func (_options *GetConfigOptions) SetID(id string) *GetConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetConfigOptions) SetHeaders(param map[string]string) *GetConfigOptions {
	options.Headers = param
	return options
}

// GetConfigVersionOptions : The GetConfigVersion options.
type GetConfigVersionOptions struct {
	// The unique project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The unique config ID.
	ID *string `json:"id" validate:"required,ne="`

	// The configuration version.
	Version *int64 `json:"version" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetConfigVersionOptions : Instantiate GetConfigVersionOptions
func (*ProjectV1) NewGetConfigVersionOptions(projectID string, id string, version int64) *GetConfigVersionOptions {
	return &GetConfigVersionOptions{
		ProjectID: core.StringPtr(projectID),
		ID: core.StringPtr(id),
		Version: core.Int64Ptr(version),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *GetConfigVersionOptions) SetProjectID(projectID string) *GetConfigVersionOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetID : Allow user to set ID
func (_options *GetConfigVersionOptions) SetID(id string) *GetConfigVersionOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *GetConfigVersionOptions) SetVersion(version int64) *GetConfigVersionOptions {
	_options.Version = core.Int64Ptr(version)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetConfigVersionOptions) SetHeaders(param map[string]string) *GetConfigVersionOptions {
	options.Headers = param
	return options
}

// GetProjectOptions : The GetProject options.
type GetProjectOptions struct {
	// The unique project ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProjectOptions : Instantiate GetProjectOptions
func (*ProjectV1) NewGetProjectOptions(id string) *GetProjectOptions {
	return &GetProjectOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetProjectOptions) SetID(id string) *GetProjectOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetProjectOptions) SetHeaders(param map[string]string) *GetProjectOptions {
	options.Headers = param
	return options
}

// InputVariable : InputVariable struct
type InputVariable struct {
	// The variable name.
	Name *string `json:"name" validate:"required"`

	// The variable type.
	Type *string `json:"type" validate:"required"`

	// Can be any value - a string, number, boolean, array, or object.
	Value interface{} `json:"value,omitempty"`

	// Whether the variable is required or not.
	Required *bool `json:"required,omitempty"`
}

// Constants associated with the InputVariable.Type property.
// The variable type.
const (
	InputVariable_Type_Array = "array"
	InputVariable_Type_Boolean = "boolean"
	InputVariable_Type_Float = "float"
	InputVariable_Type_Int = "int"
	InputVariable_Type_Number = "number"
	InputVariable_Type_Object = "object"
	InputVariable_Type_Password = "password"
	InputVariable_Type_String = "string"
)

// UnmarshalInputVariable unmarshals an instance of InputVariable from the specified map of raw messages.
func UnmarshalInputVariable(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InputVariable)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "required", &obj.Required)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstallConfigOptions : The InstallConfig options.
type InstallConfigOptions struct {
	// The unique project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The unique config ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInstallConfigOptions : Instantiate InstallConfigOptions
func (*ProjectV1) NewInstallConfigOptions(projectID string, id string) *InstallConfigOptions {
	return &InstallConfigOptions{
		ProjectID: core.StringPtr(projectID),
		ID: core.StringPtr(id),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *InstallConfigOptions) SetProjectID(projectID string) *InstallConfigOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetID : Allow user to set ID
func (_options *InstallConfigOptions) SetID(id string) *InstallConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *InstallConfigOptions) SetHeaders(param map[string]string) *InstallConfigOptions {
	options.Headers = param
	return options
}

// ListConfigResourcesOptions : The ListConfigResources options.
type ListConfigResourcesOptions struct {
	// The unique project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The unique config ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListConfigResourcesOptions : Instantiate ListConfigResourcesOptions
func (*ProjectV1) NewListConfigResourcesOptions(projectID string, id string) *ListConfigResourcesOptions {
	return &ListConfigResourcesOptions{
		ProjectID: core.StringPtr(projectID),
		ID: core.StringPtr(id),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *ListConfigResourcesOptions) SetProjectID(projectID string) *ListConfigResourcesOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetID : Allow user to set ID
func (_options *ListConfigResourcesOptions) SetID(id string) *ListConfigResourcesOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListConfigResourcesOptions) SetHeaders(param map[string]string) *ListConfigResourcesOptions {
	options.Headers = param
	return options
}

// ListConfigVersionsOptions : The ListConfigVersions options.
type ListConfigVersionsOptions struct {
	// The unique project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The unique config ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListConfigVersionsOptions : Instantiate ListConfigVersionsOptions
func (*ProjectV1) NewListConfigVersionsOptions(projectID string, id string) *ListConfigVersionsOptions {
	return &ListConfigVersionsOptions{
		ProjectID: core.StringPtr(projectID),
		ID: core.StringPtr(id),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *ListConfigVersionsOptions) SetProjectID(projectID string) *ListConfigVersionsOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetID : Allow user to set ID
func (_options *ListConfigVersionsOptions) SetID(id string) *ListConfigVersionsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListConfigVersionsOptions) SetHeaders(param map[string]string) *ListConfigVersionsOptions {
	options.Headers = param
	return options
}

// ListConfigsOptions : The ListConfigs options.
type ListConfigsOptions struct {
	// The unique project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListConfigsOptions : Instantiate ListConfigsOptions
func (*ProjectV1) NewListConfigsOptions(projectID string) *ListConfigsOptions {
	return &ListConfigsOptions{
		ProjectID: core.StringPtr(projectID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *ListConfigsOptions) SetProjectID(projectID string) *ListConfigsOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListConfigsOptions) SetHeaders(param map[string]string) *ListConfigsOptions {
	options.Headers = param
	return options
}

// ListProjectsOptions : The ListProjects options.
type ListProjectsOptions struct {
	// Marks the last entry that is returned on the page. The server uses this parameter to determine the first entry that
	// is returned on the next page. If this parameter is not specified, the logical first page is returned.
	Start *string `json:"start,omitempty"`

	// Determine the maximum number of resources to return. The number of resources that are returned is the same, with the
	// exception of the last page.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListProjectsOptions : Instantiate ListProjectsOptions
func (*ProjectV1) NewListProjectsOptions() *ListProjectsOptions {
	return &ListProjectsOptions{}
}

// SetStart : Allow user to set Start
func (_options *ListProjectsOptions) SetStart(start string) *ListProjectsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListProjectsOptions) SetLimit(limit int64) *ListProjectsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListProjectsOptions) SetHeaders(param map[string]string) *ListProjectsOptions {
	options.Headers = param
	return options
}

// OutputValue : OutputValue struct
type OutputValue struct {
	// The variable name.
	Name *string `json:"name" validate:"required"`

	// A short explanation of the output value.
	Description *string `json:"description,omitempty"`

	// Can be any value - a string, number, boolean, array, or object.
	Value interface{} `json:"value,omitempty"`
}

// UnmarshalOutputValue unmarshals an instance of OutputValue from the specified map of raw messages.
func UnmarshalOutputValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OutputValue)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PaginationLink : A pagination link.
type PaginationLink struct {
	// A relative URL.
	Href *string `json:"href" validate:"required"`

	// A pagination token.
	Start *string `json:"start,omitempty"`
}

// UnmarshalPaginationLink unmarshals an instance of PaginationLink from the specified map of raw messages.
func UnmarshalPaginationLink(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginationLink)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectCanonical : The project returned in the response body.
type ProjectCanonical struct {
	// An IBM Cloud resource name, which uniquely identifies a resource.
	Crn *string `json:"crn" validate:"required"`

	// A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time
	// format as specified by RFC 3339.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The cumulative list of needs attention items for a project. If the view is successfully retrieved, an array which
	// could be empty is returned.
	CumulativeNeedsAttentionView []CumulativeNeedsAttention `json:"cumulative_needs_attention_view,omitempty"`

	// True indicates that the fetch of the needs attention items failed. It only exists if there was an error while
	// retrieving the cumulative needs attention view.
	CumulativeNeedsAttentionViewError *bool `json:"cumulative_needs_attention_view_error,omitempty"`

	// The unique ID of a project.
	ID *string `json:"id" validate:"required"`

	// The IBM Cloud location where a resource is deployed.
	Location *string `json:"location" validate:"required"`

	// The resource group where the project's data and tools are created.
	ResourceGroup *string `json:"resource_group" validate:"required"`

	// The project status value.
	State *string `json:"state" validate:"required"`

	// The CRN of the event notifications instance if one is connected to this project.
	EventNotificationsCrn *string `json:"event_notifications_crn,omitempty"`

	// The name of the project.
	Name *string `json:"name" validate:"required"`

	// A brief explanation of the project's use in the configuration of a deployable architecture. It is possible to create
	// a project without providing a description.
	Description *string `json:"description" validate:"required"`

	// The policy that indicates whether the resources are destroyed or not when a project is deleted.
	DestroyOnDelete *bool `json:"destroy_on_delete" validate:"required"`

	// The definition of the project.
	Definition *ProjectDefinitionResponse `json:"definition,omitempty"`

	// The project configurations. These configurations are only included in the response of creating a project if a
	// configs array is specified in the request payload.
	Configs []ProjectConfigCollectionMember `json:"configs,omitempty"`
}

// Constants associated with the ProjectCanonical.State property.
// The project status value.
const (
	ProjectCanonical_State_Deleting = "deleting"
	ProjectCanonical_State_DeletingFailed = "deleting_failed"
	ProjectCanonical_State_Ready = "ready"
)

// UnmarshalProjectCanonical unmarshals an instance of ProjectCanonical from the specified map of raw messages.
func UnmarshalProjectCanonical(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectCanonical)
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cumulative_needs_attention_view", &obj.CumulativeNeedsAttentionView, UnmarshalCumulativeNeedsAttention)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cumulative_needs_attention_view_error", &obj.CumulativeNeedsAttentionViewError)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group", &obj.ResourceGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "event_notifications_crn", &obj.EventNotificationsCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "destroy_on_delete", &obj.DestroyOnDelete)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "definition", &obj.Definition, UnmarshalProjectDefinitionResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalProjectConfigCollectionMember)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectCollection : Projects list.
type ProjectCollection struct {
	// A pagination limit.
	Limit *int64 `json:"limit" validate:"required"`

	// Get the occurrencies of the total projects.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// A pagination link.
	First *PaginationLink `json:"first" validate:"required"`

	// A pagination link.
	Last *PaginationLink `json:"last,omitempty"`

	// A pagination link.
	Previous *PaginationLink `json:"previous,omitempty"`

	// A pagination link.
	Next *PaginationLink `json:"next,omitempty"`

	// An array of projects.
	Projects []ProjectCanonical `json:"projects,omitempty"`
}

// UnmarshalProjectCollection unmarshals an instance of ProjectCollection from the specified map of raw messages.
func UnmarshalProjectCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectCollection)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "projects", &obj.Projects, UnmarshalProjectCanonical)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ProjectCollection) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// ProjectConfig : The input of a project configuration.
type ProjectConfig struct {
	// The name of the configuration.
	Name *string `json:"name" validate:"required"`

	// A collection of configuration labels.
	Labels []string `json:"labels,omitempty"`

	// The description of the project configuration.
	Description *string `json:"description,omitempty"`

	// The authorization for a configuration.
	// You can authorize by using a trusted profile or an API key in Secrets Manager.
	Authorizations *ProjectConfigAuth `json:"authorizations,omitempty"`

	// The profile required for compliance.
	ComplianceProfile *ProjectConfigComplianceProfile `json:"compliance_profile,omitempty"`

	// A dotted value of catalogID.versionID.
	LocatorID *string `json:"locator_id" validate:"required"`

	// The inputs of a Schematics template property.
	Input []ProjectConfigInputVariable `json:"input,omitempty"`

	// Schematics environment variables to use to deploy the configuration. Settings are only available if they were
	// specified when the configuration was initially created.
	Setting []ProjectConfigSettingCollection `json:"setting,omitempty"`
}

// NewProjectConfig : Instantiate ProjectConfig (Generic Model Constructor)
func (*ProjectV1) NewProjectConfig(name string, locatorID string) (_model *ProjectConfig, err error) {
	_model = &ProjectConfig{
		Name: core.StringPtr(name),
		LocatorID: core.StringPtr(locatorID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalProjectConfig unmarshals an instance of ProjectConfig from the specified map of raw messages.
func UnmarshalProjectConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfig)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authorizations", &obj.Authorizations, UnmarshalProjectConfigAuth)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "compliance_profile", &obj.ComplianceProfile, UnmarshalProjectConfigComplianceProfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "locator_id", &obj.LocatorID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalProjectConfigInputVariable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "setting", &obj.Setting, UnmarshalProjectConfigSettingCollection)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigAuth : The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager.
type ProjectConfigAuth struct {
	// The trusted profile for authorizations.
	TrustedProfile *ProjectConfigAuthTrustedProfile `json:"trusted_profile,omitempty"`

	// The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets
	// Manager.
	Method *string `json:"method,omitempty"`

	// The IBM Cloud API Key.
	ApiKey *string `json:"api_key,omitempty"`
}

// UnmarshalProjectConfigAuth unmarshals an instance of ProjectConfigAuth from the specified map of raw messages.
func UnmarshalProjectConfigAuth(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigAuth)
	err = core.UnmarshalModel(m, "trusted_profile", &obj.TrustedProfile, UnmarshalProjectConfigAuthTrustedProfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "api_key", &obj.ApiKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigAuthTrustedProfile : The trusted profile for authorizations.
type ProjectConfigAuthTrustedProfile struct {
	// The unique ID of a project.
	ID *string `json:"id,omitempty"`

	// The unique ID of a project.
	TargetIamID *string `json:"target_iam_id,omitempty"`
}

// UnmarshalProjectConfigAuthTrustedProfile unmarshals an instance of ProjectConfigAuthTrustedProfile from the specified map of raw messages.
func UnmarshalProjectConfigAuthTrustedProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigAuthTrustedProfile)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_iam_id", &obj.TargetIamID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigCollection : The project configuration list.
type ProjectConfigCollection struct {
	// The collection list operation response schema that should define the array property with the name "configs".
	Configs []ProjectConfigCollectionMember `json:"configs,omitempty"`
}

// UnmarshalProjectConfigCollection unmarshals an instance of ProjectConfigCollection from the specified map of raw messages.
func UnmarshalProjectConfigCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigCollection)
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalProjectConfigCollectionMember)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigCollectionMember : The configuration metadata.
type ProjectConfigCollectionMember struct {
	// The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
	ID *string `json:"id" validate:"required"`

	// The unique ID of a project.
	ProjectID *string `json:"project_id" validate:"required"`

	// The version of the configuration.
	Version *int64 `json:"version" validate:"required"`

	// The flag that indicates whether the version of the configuration is draft, or active.
	IsDraft *bool `json:"is_draft" validate:"required"`

	// The needs attention state of a configuration.
	NeedsAttentionState []interface{} `json:"needs_attention_state,omitempty"`

	// The state of the configuration.
	State *string `json:"state" validate:"required"`

	// The pipeline state of the configuration. It only exists after the first configuration validation.
	PipelineState *string `json:"pipeline_state,omitempty"`

	// The flag that indicates whether a configuration update is available.
	UpdateAvailable *bool `json:"update_available" validate:"required"`

	// A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time
	// format as specified by RFC 3339.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time
	// format as specified by RFC 3339.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// The last approved metadata of the configuration.
	LastApproved *ProjectConfigMetadataLastApproved `json:"last_approved,omitempty"`

	// A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time
	// format as specified by RFC 3339.
	LastSave *strfmt.DateTime `json:"last_save,omitempty"`

	// The project configuration version.
	ActiveDraft *ProjectConfigVersionSummary `json:"active_draft,omitempty"`

	// The project configuration definition.
	Definition *ProjectConfigDefinition `json:"definition" validate:"required"`

	// A relative URL.
	Href *string `json:"href" validate:"required"`
}

// Constants associated with the ProjectConfigCollectionMember.State property.
// The state of the configuration.
const (
	ProjectConfigCollectionMember_State_Approved = "approved"
	ProjectConfigCollectionMember_State_Deleted = "deleted"
	ProjectConfigCollectionMember_State_Deleting = "deleting"
	ProjectConfigCollectionMember_State_DeletingFailed = "deleting_failed"
	ProjectConfigCollectionMember_State_Discarded = "discarded"
	ProjectConfigCollectionMember_State_Draft = "draft"
	ProjectConfigCollectionMember_State_Installed = "installed"
	ProjectConfigCollectionMember_State_InstalledFailed = "installed_failed"
	ProjectConfigCollectionMember_State_Installing = "installing"
	ProjectConfigCollectionMember_State_Superceded = "superceded"
	ProjectConfigCollectionMember_State_Uninstalling = "uninstalling"
	ProjectConfigCollectionMember_State_UninstallingFailed = "uninstalling_failed"
	ProjectConfigCollectionMember_State_Validated = "validated"
	ProjectConfigCollectionMember_State_Validating = "validating"
	ProjectConfigCollectionMember_State_ValidatingFailed = "validating_failed"
)

// Constants associated with the ProjectConfigCollectionMember.PipelineState property.
// The pipeline state of the configuration. It only exists after the first configuration validation.
const (
	ProjectConfigCollectionMember_PipelineState_PipelineFailed = "pipeline_failed"
	ProjectConfigCollectionMember_PipelineState_PipelineRunning = "pipeline_running"
	ProjectConfigCollectionMember_PipelineState_PipelineSucceeded = "pipeline_succeeded"
)

// UnmarshalProjectConfigCollectionMember unmarshals an instance of ProjectConfigCollectionMember from the specified map of raw messages.
func UnmarshalProjectConfigCollectionMember(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigCollectionMember)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_draft", &obj.IsDraft)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "needs_attention_state", &obj.NeedsAttentionState)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pipeline_state", &obj.PipelineState)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "update_available", &obj.UpdateAvailable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last_approved", &obj.LastApproved, UnmarshalProjectConfigMetadataLastApproved)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_save", &obj.LastSave)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "active_draft", &obj.ActiveDraft, UnmarshalProjectConfigVersionSummary)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "definition", &obj.Definition, UnmarshalProjectConfigDefinition)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigComplianceProfile : The profile required for compliance.
type ProjectConfigComplianceProfile struct {
	// The unique ID of a project.
	ID *string `json:"id,omitempty"`

	// The unique ID of a project.
	InstanceID *string `json:"instance_id,omitempty"`

	// The location of the compliance instance.
	InstanceLocation *string `json:"instance_location,omitempty"`

	// The unique ID of a project.
	AttachmentID *string `json:"attachment_id,omitempty"`

	// The name of the compliance profile.
	ProfileName *string `json:"profile_name,omitempty"`
}

// UnmarshalProjectConfigComplianceProfile unmarshals an instance of ProjectConfigComplianceProfile from the specified map of raw messages.
func UnmarshalProjectConfigComplianceProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigComplianceProfile)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_location", &obj.InstanceLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "attachment_id", &obj.AttachmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_name", &obj.ProfileName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigDefinition : The project configuration definition.
type ProjectConfigDefinition struct {
	// The name of the configuration.
	Name *string `json:"name" validate:"required"`

	// A collection of configuration labels.
	Labels []string `json:"labels,omitempty"`

	// The description of the project configuration.
	Description *string `json:"description,omitempty"`

	// The authorization for a configuration.
	// You can authorize by using a trusted profile or an API key in Secrets Manager.
	Authorizations *ProjectConfigAuth `json:"authorizations,omitempty"`

	// The profile required for compliance.
	ComplianceProfile *ProjectConfigComplianceProfile `json:"compliance_profile,omitempty"`

	// A dotted value of catalogID.versionID.
	LocatorID *string `json:"locator_id" validate:"required"`

	// The type of a project configuration manual property.
	Type *string `json:"type" validate:"required"`

	// The outputs of a Schematics template property.
	Input []InputVariable `json:"input,omitempty"`

	// The outputs of a Schematics template property.
	Output []OutputValue `json:"output,omitempty"`

	// Schematics environment variables to use to deploy the configuration. Settings are only available if they were
	// specified when the configuration was initially created.
	Setting []ProjectConfigSettingCollection `json:"setting,omitempty"`
}

// Constants associated with the ProjectConfigDefinition.Type property.
// The type of a project configuration manual property.
const (
	ProjectConfigDefinition_Type_SchematicsBlueprint = "schematics_blueprint"
	ProjectConfigDefinition_Type_TerraformTemplate = "terraform_template"
)

// UnmarshalProjectConfigDefinition unmarshals an instance of ProjectConfigDefinition from the specified map of raw messages.
func UnmarshalProjectConfigDefinition(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigDefinition)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authorizations", &obj.Authorizations, UnmarshalProjectConfigAuth)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "compliance_profile", &obj.ComplianceProfile, UnmarshalProjectConfigComplianceProfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "locator_id", &obj.LocatorID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalInputVariable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalOutputValue)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "setting", &obj.Setting, UnmarshalProjectConfigSettingCollection)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigDelete : Deletes the configuration response.
type ProjectConfigDelete struct {
	// The unique ID of a project.
	ID *string `json:"id" validate:"required"`
}

// UnmarshalProjectConfigDelete unmarshals an instance of ProjectConfigDelete from the specified map of raw messages.
func UnmarshalProjectConfigDelete(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigDelete)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigGetResponse : The configuration metadata.
type ProjectConfigGetResponse struct {
	// The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
	ID *string `json:"id" validate:"required"`

	// The unique ID of a project.
	ProjectID *string `json:"project_id" validate:"required"`

	// The version of the configuration.
	Version *int64 `json:"version" validate:"required"`

	// The flag that indicates whether the version of the configuration is draft, or active.
	IsDraft *bool `json:"is_draft" validate:"required"`

	// The needs attention state of a configuration.
	NeedsAttentionState []interface{} `json:"needs_attention_state,omitempty"`

	// The state of the configuration.
	State *string `json:"state" validate:"required"`

	// The pipeline state of the configuration. It only exists after the first configuration validation.
	PipelineState *string `json:"pipeline_state,omitempty"`

	// The flag that indicates whether a configuration update is available.
	UpdateAvailable *bool `json:"update_available" validate:"required"`

	// A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time
	// format as specified by RFC 3339.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time
	// format as specified by RFC 3339.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// The last approved metadata of the configuration.
	LastApproved *ProjectConfigMetadataLastApproved `json:"last_approved,omitempty"`

	// A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time
	// format as specified by RFC 3339.
	LastSave *strfmt.DateTime `json:"last_save,omitempty"`

	// The summaries of jobs that were performed on the configuration.
	JobSummary *ProjectConfigMetadataJobSummary `json:"job_summary,omitempty"`

	// The Code Risk Analyzer logs of the configuration.
	CraLogs *ProjectConfigMetadataCraLogs `json:"cra_logs,omitempty"`

	// The cost estimate of the configuration.
	// It only exists after the first configuration validation.
	CostEstimate *ProjectConfigMetadataCostEstimate `json:"cost_estimate,omitempty"`

	// The summaries of jobs that were performed on the configuration.
	LastDeploymentJobSummary *ProjectConfigMetadataJobSummary `json:"last_deployment_job_summary,omitempty"`

	// The project configuration version.
	ActiveDraft *ProjectConfigVersionSummary `json:"active_draft,omitempty"`

	// The project configuration definition.
	Definition *ProjectConfigDefinition `json:"definition" validate:"required"`
}

// Constants associated with the ProjectConfigGetResponse.State property.
// The state of the configuration.
const (
	ProjectConfigGetResponse_State_Approved = "approved"
	ProjectConfigGetResponse_State_Deleted = "deleted"
	ProjectConfigGetResponse_State_Deleting = "deleting"
	ProjectConfigGetResponse_State_DeletingFailed = "deleting_failed"
	ProjectConfigGetResponse_State_Discarded = "discarded"
	ProjectConfigGetResponse_State_Draft = "draft"
	ProjectConfigGetResponse_State_Installed = "installed"
	ProjectConfigGetResponse_State_InstalledFailed = "installed_failed"
	ProjectConfigGetResponse_State_Installing = "installing"
	ProjectConfigGetResponse_State_Superceded = "superceded"
	ProjectConfigGetResponse_State_Uninstalling = "uninstalling"
	ProjectConfigGetResponse_State_UninstallingFailed = "uninstalling_failed"
	ProjectConfigGetResponse_State_Validated = "validated"
	ProjectConfigGetResponse_State_Validating = "validating"
	ProjectConfigGetResponse_State_ValidatingFailed = "validating_failed"
)

// Constants associated with the ProjectConfigGetResponse.PipelineState property.
// The pipeline state of the configuration. It only exists after the first configuration validation.
const (
	ProjectConfigGetResponse_PipelineState_PipelineFailed = "pipeline_failed"
	ProjectConfigGetResponse_PipelineState_PipelineRunning = "pipeline_running"
	ProjectConfigGetResponse_PipelineState_PipelineSucceeded = "pipeline_succeeded"
)

// UnmarshalProjectConfigGetResponse unmarshals an instance of ProjectConfigGetResponse from the specified map of raw messages.
func UnmarshalProjectConfigGetResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigGetResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_draft", &obj.IsDraft)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "needs_attention_state", &obj.NeedsAttentionState)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pipeline_state", &obj.PipelineState)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "update_available", &obj.UpdateAvailable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last_approved", &obj.LastApproved, UnmarshalProjectConfigMetadataLastApproved)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_save", &obj.LastSave)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "job_summary", &obj.JobSummary, UnmarshalProjectConfigMetadataJobSummary)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cra_logs", &obj.CraLogs, UnmarshalProjectConfigMetadataCraLogs)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cost_estimate", &obj.CostEstimate, UnmarshalProjectConfigMetadataCostEstimate)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last_deployment_job_summary", &obj.LastDeploymentJobSummary, UnmarshalProjectConfigMetadataJobSummary)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "active_draft", &obj.ActiveDraft, UnmarshalProjectConfigVersionSummary)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "definition", &obj.Definition, UnmarshalProjectConfigDefinition)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigInputVariable : ProjectConfigInputVariable struct
type ProjectConfigInputVariable struct {
	// The variable name.
	Name *string `json:"name" validate:"required"`

	// Can be any value - a string, number, boolean, array, or object.
	Value interface{} `json:"value,omitempty"`
}

// NewProjectConfigInputVariable : Instantiate ProjectConfigInputVariable (Generic Model Constructor)
func (*ProjectV1) NewProjectConfigInputVariable(name string) (_model *ProjectConfigInputVariable, err error) {
	_model = &ProjectConfigInputVariable{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalProjectConfigInputVariable unmarshals an instance of ProjectConfigInputVariable from the specified map of raw messages.
func UnmarshalProjectConfigInputVariable(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigInputVariable)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigMetadataCostEstimate : The cost estimate of the configuration. It only exists after the first configuration validation.
type ProjectConfigMetadataCostEstimate struct {
	// The version of the cost estimate of the configuration.
	Version *string `json:"version,omitempty"`

	// The currency of the cost estimate of the configuration.
	Currency *string `json:"currency,omitempty"`

	// The total hourly cost estimate of the configuration.
	TotalHourlyCost *string `json:"totalHourlyCost,omitempty"`

	// The total monthly cost estimate of the configuration.
	TotalMonthlyCost *string `json:"totalMonthlyCost,omitempty"`

	// The past total hourly cost estimate of the configuration.
	PastTotalHourlyCost *string `json:"pastTotalHourlyCost,omitempty"`

	// The past total monthly cost estimate of the configuration.
	PastTotalMonthlyCost *string `json:"pastTotalMonthlyCost,omitempty"`

	// The difference between current and past total hourly cost estimates of the configuration.
	DiffTotalHourlyCost *string `json:"diffTotalHourlyCost,omitempty"`

	// The difference between current and past total monthly cost estimates of the configuration.
	DiffTotalMonthlyCost *string `json:"diffTotalMonthlyCost,omitempty"`

	// A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time
	// format as specified by RFC 3339.
	TimeGenerated *strfmt.DateTime `json:"timeGenerated,omitempty"`

	// The unique ID of a project.
	UserID *string `json:"user_id,omitempty"`
}

// UnmarshalProjectConfigMetadataCostEstimate unmarshals an instance of ProjectConfigMetadataCostEstimate from the specified map of raw messages.
func UnmarshalProjectConfigMetadataCostEstimate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigMetadataCostEstimate)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency", &obj.Currency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "totalHourlyCost", &obj.TotalHourlyCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "totalMonthlyCost", &obj.TotalMonthlyCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pastTotalHourlyCost", &obj.PastTotalHourlyCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pastTotalMonthlyCost", &obj.PastTotalMonthlyCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "diffTotalHourlyCost", &obj.DiffTotalHourlyCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "diffTotalMonthlyCost", &obj.DiffTotalMonthlyCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timeGenerated", &obj.TimeGenerated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigMetadataCraLogs : The Code Risk Analyzer logs of the configuration.
type ProjectConfigMetadataCraLogs struct {
	// The version of the Code Risk Analyzer logs of the configuration.
	CraVersion *string `json:"cra_version,omitempty"`

	// The schema version of Code Risk Analyzer logs of the configuration.
	SchemaVersion *string `json:"schema_version,omitempty"`

	// The status of the Code Risk Analyzer logs of the configuration.
	Status *string `json:"status,omitempty"`

	// The summary of the Code Risk Analyzer logs of the configuration.
	Summary map[string]interface{} `json:"summary,omitempty"`

	// A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time
	// format as specified by RFC 3339.
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// UnmarshalProjectConfigMetadataCraLogs unmarshals an instance of ProjectConfigMetadataCraLogs from the specified map of raw messages.
func UnmarshalProjectConfigMetadataCraLogs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigMetadataCraLogs)
	err = core.UnmarshalPrimitive(m, "cra_version", &obj.CraVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "schema_version", &obj.SchemaVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "summary", &obj.Summary)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigMetadataJobSummary : The summaries of jobs that were performed on the configuration.
type ProjectConfigMetadataJobSummary struct {
	// The summary of the plan jobs on the configuration.
	PlanSummary map[string]interface{} `json:"plan_summary,omitempty"`

	// The summary of the apply jobs on the configuration.
	ApplySummary map[string]interface{} `json:"apply_summary,omitempty"`

	// The summary of the destroy jobs on the configuration.
	DestroySummary map[string]interface{} `json:"destroy_summary,omitempty"`

	// The message summaries of jobs on the configuration.
	MessageSummary map[string]interface{} `json:"message_summary,omitempty"`

	// The messages of plan jobs on the configuration.
	PlanMessages map[string]interface{} `json:"plan_messages,omitempty"`

	// The messages of apply jobs on the configuration.
	ApplyMessages map[string]interface{} `json:"apply_messages,omitempty"`

	// The messages of destroy jobs on the configuration.
	DestroyMessages map[string]interface{} `json:"destroy_messages,omitempty"`
}

// UnmarshalProjectConfigMetadataJobSummary unmarshals an instance of ProjectConfigMetadataJobSummary from the specified map of raw messages.
func UnmarshalProjectConfigMetadataJobSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigMetadataJobSummary)
	err = core.UnmarshalPrimitive(m, "plan_summary", &obj.PlanSummary)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "apply_summary", &obj.ApplySummary)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "destroy_summary", &obj.DestroySummary)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message_summary", &obj.MessageSummary)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "plan_messages", &obj.PlanMessages)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "apply_messages", &obj.ApplyMessages)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "destroy_messages", &obj.DestroyMessages)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigMetadataLastApproved : The last approved metadata of the configuration.
type ProjectConfigMetadataLastApproved struct {
	// The flag that indicates whether the approval was forced approved.
	IsForced *bool `json:"is_forced" validate:"required"`

	// The comment left by the user who approved the configuration.
	Comment *string `json:"comment,omitempty"`

	// A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time
	// format as specified by RFC 3339.
	Timestamp *strfmt.DateTime `json:"timestamp" validate:"required"`

	// The unique ID of a project.
	UserID *string `json:"user_id" validate:"required"`
}

// UnmarshalProjectConfigMetadataLastApproved unmarshals an instance of ProjectConfigMetadataLastApproved from the specified map of raw messages.
func UnmarshalProjectConfigMetadataLastApproved(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigMetadataLastApproved)
	err = core.UnmarshalPrimitive(m, "is_forced", &obj.IsForced)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "comment", &obj.Comment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigResource : ProjectConfigResource struct
type ProjectConfigResource struct {
	// An IBM Cloud resource name, which uniquely identifies a resource.
	ResourceCrn *string `json:"resource_crn,omitempty"`

	// The name of the resource.
	ResourceName *string `json:"resource_name,omitempty"`

	// The resource type.
	ResourceType *string `json:"resource_type,omitempty"`

	// The flag that indicates whether the status of the resource is tainted.
	ResourceTainted *bool `json:"resource_tainted,omitempty"`

	// The resource group of the resource.
	ResourceGroupName *string `json:"resource_group_name,omitempty"`
}

// UnmarshalProjectConfigResource unmarshals an instance of ProjectConfigResource from the specified map of raw messages.
func UnmarshalProjectConfigResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigResource)
	err = core.UnmarshalPrimitive(m, "resource_crn", &obj.ResourceCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_name", &obj.ResourceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_type", &obj.ResourceType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_tainted", &obj.ResourceTainted)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_name", &obj.ResourceGroupName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigResourceCollection : The project configuration resource list.
type ProjectConfigResourceCollection struct {
	// The collection list operation response schema that defines the array property with the name `resources`.
	Resources []ProjectConfigResource `json:"resources,omitempty"`

	// The total number of resources deployed by the configuration.
	ResourcesCount *int64 `json:"resources_count" validate:"required"`
}

// UnmarshalProjectConfigResourceCollection unmarshals an instance of ProjectConfigResourceCollection from the specified map of raw messages.
func UnmarshalProjectConfigResourceCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigResourceCollection)
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalProjectConfigResource)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resources_count", &obj.ResourcesCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigSettingCollection : ProjectConfigSettingCollection struct
type ProjectConfigSettingCollection struct {
	// The name of the configuration setting.
	Name *string `json:"name" validate:"required"`

	// The value of the configuration setting.
	Value *string `json:"value" validate:"required"`
}

// NewProjectConfigSettingCollection : Instantiate ProjectConfigSettingCollection (Generic Model Constructor)
func (*ProjectV1) NewProjectConfigSettingCollection(name string, value string) (_model *ProjectConfigSettingCollection, err error) {
	_model = &ProjectConfigSettingCollection{
		Name: core.StringPtr(name),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalProjectConfigSettingCollection unmarshals an instance of ProjectConfigSettingCollection from the specified map of raw messages.
func UnmarshalProjectConfigSettingCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigSettingCollection)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigVersionResponse : The project configuration version.
type ProjectConfigVersionResponse struct {
	// The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
	ID *string `json:"id" validate:"required"`

	// The unique ID of a project.
	ProjectID *string `json:"project_id" validate:"required"`

	// The version of the configuration.
	Version *int64 `json:"version" validate:"required"`

	// The flag that indicates whether the version of the configuration is draft, or active.
	IsDraft *bool `json:"is_draft" validate:"required"`

	// The needs attention state of a configuration.
	NeedsAttentionState []interface{} `json:"needs_attention_state,omitempty"`

	// The state of the configuration.
	State *string `json:"state" validate:"required"`

	// The pipeline state of the configuration. It only exists after the first configuration validation.
	PipelineState *string `json:"pipeline_state,omitempty"`

	// The flag that indicates whether a configuration update is available.
	UpdateAvailable *bool `json:"update_available" validate:"required"`

	// A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time
	// format as specified by RFC 3339.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time
	// format as specified by RFC 3339.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// The last approved metadata of the configuration.
	LastApproved *ProjectConfigMetadataLastApproved `json:"last_approved,omitempty"`

	// A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time
	// format as specified by RFC 3339.
	LastSave *strfmt.DateTime `json:"last_save,omitempty"`

	// The summaries of jobs that were performed on the configuration.
	JobSummary *ProjectConfigMetadataJobSummary `json:"job_summary,omitempty"`

	// The Code Risk Analyzer logs of the configuration.
	CraLogs *ProjectConfigMetadataCraLogs `json:"cra_logs,omitempty"`

	// The cost estimate of the configuration.
	// It only exists after the first configuration validation.
	CostEstimate *ProjectConfigMetadataCostEstimate `json:"cost_estimate,omitempty"`

	// The summaries of jobs that were performed on the configuration.
	LastDeploymentJobSummary *ProjectConfigMetadataJobSummary `json:"last_deployment_job_summary,omitempty"`

	// The project configuration definition.
	Definition *ProjectConfigDefinition `json:"definition" validate:"required"`
}

// Constants associated with the ProjectConfigVersionResponse.State property.
// The state of the configuration.
const (
	ProjectConfigVersionResponse_State_Approved = "approved"
	ProjectConfigVersionResponse_State_Deleted = "deleted"
	ProjectConfigVersionResponse_State_Deleting = "deleting"
	ProjectConfigVersionResponse_State_DeletingFailed = "deleting_failed"
	ProjectConfigVersionResponse_State_Discarded = "discarded"
	ProjectConfigVersionResponse_State_Draft = "draft"
	ProjectConfigVersionResponse_State_Installed = "installed"
	ProjectConfigVersionResponse_State_InstalledFailed = "installed_failed"
	ProjectConfigVersionResponse_State_Installing = "installing"
	ProjectConfigVersionResponse_State_Superceded = "superceded"
	ProjectConfigVersionResponse_State_Uninstalling = "uninstalling"
	ProjectConfigVersionResponse_State_UninstallingFailed = "uninstalling_failed"
	ProjectConfigVersionResponse_State_Validated = "validated"
	ProjectConfigVersionResponse_State_Validating = "validating"
	ProjectConfigVersionResponse_State_ValidatingFailed = "validating_failed"
)

// Constants associated with the ProjectConfigVersionResponse.PipelineState property.
// The pipeline state of the configuration. It only exists after the first configuration validation.
const (
	ProjectConfigVersionResponse_PipelineState_PipelineFailed = "pipeline_failed"
	ProjectConfigVersionResponse_PipelineState_PipelineRunning = "pipeline_running"
	ProjectConfigVersionResponse_PipelineState_PipelineSucceeded = "pipeline_succeeded"
)

// UnmarshalProjectConfigVersionResponse unmarshals an instance of ProjectConfigVersionResponse from the specified map of raw messages.
func UnmarshalProjectConfigVersionResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigVersionResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_draft", &obj.IsDraft)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "needs_attention_state", &obj.NeedsAttentionState)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pipeline_state", &obj.PipelineState)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "update_available", &obj.UpdateAvailable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last_approved", &obj.LastApproved, UnmarshalProjectConfigMetadataLastApproved)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_save", &obj.LastSave)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "job_summary", &obj.JobSummary, UnmarshalProjectConfigMetadataJobSummary)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cra_logs", &obj.CraLogs, UnmarshalProjectConfigMetadataCraLogs)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cost_estimate", &obj.CostEstimate, UnmarshalProjectConfigMetadataCostEstimate)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last_deployment_job_summary", &obj.LastDeploymentJobSummary, UnmarshalProjectConfigMetadataJobSummary)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "definition", &obj.Definition, UnmarshalProjectConfigDefinition)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigVersionSummary : The project configuration version.
type ProjectConfigVersionSummary struct {
	// The version number of the configuration.
	Version *int64 `json:"version" validate:"required"`

	// The state of the configuration draft.
	State *string `json:"state" validate:"required"`

	// The pipeline state of the configuration. It only exists after the first configuration validation.
	PipelineState *string `json:"pipeline_state,omitempty"`

	// A relative URL.
	Href *string `json:"href,omitempty"`
}

// Constants associated with the ProjectConfigVersionSummary.State property.
// The state of the configuration draft.
const (
	ProjectConfigVersionSummary_State_Active = "active"
	ProjectConfigVersionSummary_State_Discarded = "discarded"
	ProjectConfigVersionSummary_State_Merged = "merged"
)

// Constants associated with the ProjectConfigVersionSummary.PipelineState property.
// The pipeline state of the configuration. It only exists after the first configuration validation.
const (
	ProjectConfigVersionSummary_PipelineState_PipelineFailed = "pipeline_failed"
	ProjectConfigVersionSummary_PipelineState_PipelineRunning = "pipeline_running"
	ProjectConfigVersionSummary_PipelineState_PipelineSucceeded = "pipeline_succeeded"
)

// UnmarshalProjectConfigVersionSummary unmarshals an instance of ProjectConfigVersionSummary from the specified map of raw messages.
func UnmarshalProjectConfigVersionSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigVersionSummary)
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pipeline_state", &obj.PipelineState)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigVersionSummaryCollection : The project configuration version list.
type ProjectConfigVersionSummaryCollection struct {
	// The collection list operation response schema that defines the array property with the name `versions`.
	Versions []ProjectConfigVersionSummary `json:"versions,omitempty"`
}

// UnmarshalProjectConfigVersionSummaryCollection unmarshals an instance of ProjectConfigVersionSummaryCollection from the specified map of raw messages.
func UnmarshalProjectConfigVersionSummaryCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigVersionSummaryCollection)
	err = core.UnmarshalModel(m, "versions", &obj.Versions, UnmarshalProjectConfigVersionSummary)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectDefinitionResponse : The definition of the project.
type ProjectDefinitionResponse struct {
	// The name of the project.
	Name *string `json:"name" validate:"required"`

	// A brief explanation of the project's use in the configuration of a deployable architecture. It is possible to create
	// a project without providing a description.
	Description *string `json:"description" validate:"required"`

	// The policy that indicates whether the resources are destroyed or not when a project is deleted.
	DestroyOnDelete *bool `json:"destroy_on_delete" validate:"required"`
}

// UnmarshalProjectDefinitionResponse unmarshals an instance of ProjectDefinitionResponse from the specified map of raw messages.
func UnmarshalProjectDefinitionResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectDefinitionResponse)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "destroy_on_delete", &obj.DestroyOnDelete)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectResponseSummary : The project returned in the response body.
type ProjectResponseSummary struct {
	// An IBM Cloud resource name, which uniquely identifies a resource.
	Crn *string `json:"crn" validate:"required"`

	// A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time
	// format as specified by RFC 3339.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The cumulative list of needs attention items for a project. If the view is successfully retrieved, an array which
	// could be empty is returned.
	CumulativeNeedsAttentionView []CumulativeNeedsAttention `json:"cumulative_needs_attention_view,omitempty"`

	// True indicates that the fetch of the needs attention items failed. It only exists if there was an error while
	// retrieving the cumulative needs attention view.
	CumulativeNeedsAttentionViewError *bool `json:"cumulative_needs_attention_view_error,omitempty"`

	// The unique ID of a project.
	ID *string `json:"id,omitempty"`

	// The IBM Cloud location where a resource is deployed.
	Location *string `json:"location" validate:"required"`

	// The resource group where the project's data and tools are created.
	ResourceGroup *string `json:"resource_group" validate:"required"`

	// The project status value.
	State *string `json:"state" validate:"required"`

	// The CRN of the event notifications instance if one is connected to this project.
	EventNotificationsCrn *string `json:"event_notifications_crn,omitempty"`

	// The name of the project.
	Name *string `json:"name" validate:"required"`

	// A brief explanation of the project's use in the configuration of a deployable architecture. It is possible to create
	// a project without providing a description.
	Description *string `json:"description" validate:"required"`

	// The policy that indicates whether the resources are destroyed or not when a project is deleted.
	DestroyOnDelete *bool `json:"destroy_on_delete" validate:"required"`

	// The definition of the project.
	Definition *ProjectDefinitionResponse `json:"definition,omitempty"`
}

// Constants associated with the ProjectResponseSummary.State property.
// The project status value.
const (
	ProjectResponseSummary_State_Deleting = "deleting"
	ProjectResponseSummary_State_DeletingFailed = "deleting_failed"
	ProjectResponseSummary_State_Ready = "ready"
)

// UnmarshalProjectResponseSummary unmarshals an instance of ProjectResponseSummary from the specified map of raw messages.
func UnmarshalProjectResponseSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectResponseSummary)
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cumulative_needs_attention_view", &obj.CumulativeNeedsAttentionView, UnmarshalCumulativeNeedsAttention)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cumulative_needs_attention_view_error", &obj.CumulativeNeedsAttentionViewError)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group", &obj.ResourceGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "event_notifications_crn", &obj.EventNotificationsCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "destroy_on_delete", &obj.DestroyOnDelete)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "definition", &obj.Definition, UnmarshalProjectDefinitionResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UninstallConfigOptions : The UninstallConfig options.
type UninstallConfigOptions struct {
	// The unique project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The unique config ID.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUninstallConfigOptions : Instantiate UninstallConfigOptions
func (*ProjectV1) NewUninstallConfigOptions(projectID string, id string) *UninstallConfigOptions {
	return &UninstallConfigOptions{
		ProjectID: core.StringPtr(projectID),
		ID: core.StringPtr(id),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *UninstallConfigOptions) SetProjectID(projectID string) *UninstallConfigOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetID : Allow user to set ID
func (_options *UninstallConfigOptions) SetID(id string) *UninstallConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UninstallConfigOptions) SetHeaders(param map[string]string) *UninstallConfigOptions {
	options.Headers = param
	return options
}

// UpdateConfigOptions : The UpdateConfig options.
type UpdateConfigOptions struct {
	// The unique project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The unique config ID.
	ID *string `json:"id" validate:"required,ne="`

	// A dotted value of catalogID.versionID.
	LocatorID *string `json:"locator_id,omitempty"`

	// The inputs of a Schematics template property.
	Input []ProjectConfigInputVariable `json:"input,omitempty"`

	// Schematics environment variables to use to deploy the configuration. Settings are only available if they were
	// specified when the configuration was initially created.
	Setting []ProjectConfigSettingCollection `json:"setting,omitempty"`

	// The configuration name.
	Name *string `json:"name,omitempty"`

	// The configuration labels.
	Labels []string `json:"labels,omitempty"`

	// A project configuration description.
	Description *string `json:"description,omitempty"`

	// The authorization for a configuration.
	// You can authorize by using a trusted profile or an API key in Secrets Manager.
	Authorizations *ProjectConfigAuth `json:"authorizations,omitempty"`

	// The profile required for compliance.
	ComplianceProfile *ProjectConfigComplianceProfile `json:"compliance_profile,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateConfigOptions : Instantiate UpdateConfigOptions
func (*ProjectV1) NewUpdateConfigOptions(projectID string, id string) *UpdateConfigOptions {
	return &UpdateConfigOptions{
		ProjectID: core.StringPtr(projectID),
		ID: core.StringPtr(id),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *UpdateConfigOptions) SetProjectID(projectID string) *UpdateConfigOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetID : Allow user to set ID
func (_options *UpdateConfigOptions) SetID(id string) *UpdateConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetLocatorID : Allow user to set LocatorID
func (_options *UpdateConfigOptions) SetLocatorID(locatorID string) *UpdateConfigOptions {
	_options.LocatorID = core.StringPtr(locatorID)
	return _options
}

// SetInput : Allow user to set Input
func (_options *UpdateConfigOptions) SetInput(input []ProjectConfigInputVariable) *UpdateConfigOptions {
	_options.Input = input
	return _options
}

// SetSetting : Allow user to set Setting
func (_options *UpdateConfigOptions) SetSetting(setting []ProjectConfigSettingCollection) *UpdateConfigOptions {
	_options.Setting = setting
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateConfigOptions) SetName(name string) *UpdateConfigOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetLabels : Allow user to set Labels
func (_options *UpdateConfigOptions) SetLabels(labels []string) *UpdateConfigOptions {
	_options.Labels = labels
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateConfigOptions) SetDescription(description string) *UpdateConfigOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetAuthorizations : Allow user to set Authorizations
func (_options *UpdateConfigOptions) SetAuthorizations(authorizations *ProjectConfigAuth) *UpdateConfigOptions {
	_options.Authorizations = authorizations
	return _options
}

// SetComplianceProfile : Allow user to set ComplianceProfile
func (_options *UpdateConfigOptions) SetComplianceProfile(complianceProfile *ProjectConfigComplianceProfile) *UpdateConfigOptions {
	_options.ComplianceProfile = complianceProfile
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateConfigOptions) SetHeaders(param map[string]string) *UpdateConfigOptions {
	options.Headers = param
	return options
}

// UpdateProjectOptions : The UpdateProject options.
type UpdateProjectOptions struct {
	// The unique project ID.
	ID *string `json:"id" validate:"required,ne="`

	// The project name.
	Name *string `json:"name,omitempty"`

	// The description of the project.
	Description *string `json:"description,omitempty"`

	// The policy that indicates whether the resources are destroyed or not when a project is deleted.
	DestroyOnDelete *bool `json:"destroy_on_delete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateProjectOptions : Instantiate UpdateProjectOptions
func (*ProjectV1) NewUpdateProjectOptions(id string) *UpdateProjectOptions {
	return &UpdateProjectOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *UpdateProjectOptions) SetID(id string) *UpdateProjectOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateProjectOptions) SetName(name string) *UpdateProjectOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateProjectOptions) SetDescription(description string) *UpdateProjectOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetDestroyOnDelete : Allow user to set DestroyOnDelete
func (_options *UpdateProjectOptions) SetDestroyOnDelete(destroyOnDelete bool) *UpdateProjectOptions {
	_options.DestroyOnDelete = core.BoolPtr(destroyOnDelete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateProjectOptions) SetHeaders(param map[string]string) *UpdateProjectOptions {
	options.Headers = param
	return options
}

//
// ProjectsPager can be used to simplify the use of the "ListProjects" method.
//
type ProjectsPager struct {
	hasNext bool
	options *ListProjectsOptions
	client  *ProjectV1
	pageContext struct {
		next *string
	}
}

// NewProjectsPager returns a new ProjectsPager instance.
func (project *ProjectV1) NewProjectsPager(options *ListProjectsOptions) (pager *ProjectsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListProjectsOptions = *options
	pager = &ProjectsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  project,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ProjectsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ProjectsPager) GetNextWithContext(ctx context.Context) (page []ProjectCanonical, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListProjectsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Projects

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ProjectsPager) GetAllWithContext(ctx context.Context) (allItems []ProjectCanonical, err error) {
	for pager.HasNext() {
		var nextPage []ProjectCanonical
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ProjectsPager) GetNext() (page []ProjectCanonical, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ProjectsPager) GetAll() (allItems []ProjectCanonical, err error) {
	return pager.GetAllWithContext(context.Background())
}
