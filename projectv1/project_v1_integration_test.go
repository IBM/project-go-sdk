// +build integration

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

package projectv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/project-go-sdk/projectv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the projectv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`ProjectV1 Integration Tests`, func() {
	const externalConfigFile = "../project_v1.env"

	var (
		err          error
		projectService *projectv1.ProjectV1
		serviceURL   string
		config       map[string]string

		// Variables to hold link values
		configIdLink string
		projectIdLink string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(projectv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			projectServiceOptions := &projectv1.ProjectV1Options{}

			projectService, err = projectv1.NewProjectV1UsingExternalConfig(projectServiceOptions)
			Expect(err).To(BeNil())
			Expect(projectService).ToNot(BeNil())
			Expect(projectService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			projectService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateProject - Create a project`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProject(createProjectOptions *CreateProjectOptions)`, func() {
			projectPrototypeDefinitionModel := &projectv1.ProjectPrototypeDefinition{
				Name: core.StringPtr("acme-microservice"),
				Description: core.StringPtr("A microservice to deploy on top of ACME infrastructure."),
				DestroyOnDelete: core.BoolPtr(true),
			}

			projectConfigAuthModel := &projectv1.ProjectConfigAuth{
				TrustedProfileID: core.StringPtr("testString"),
				Method: core.StringPtr("api_key"),
				ApiKey: core.StringPtr("testString"),
			}

			projectComplianceProfileModel := &projectv1.ProjectComplianceProfile{
				ID: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				InstanceLocation: core.StringPtr("testString"),
				AttachmentID: core.StringPtr("testString"),
				ProfileName: core.StringPtr("testString"),
			}

			inputVariableModel := &projectv1.InputVariable{
			}
			inputVariableModel.SetProperty("foo", core.StringPtr("testString"))

			projectConfigSettingModel := &projectv1.ProjectConfigSetting{
			}
			projectConfigSettingModel.SetProperty("foo", core.StringPtr("testString"))

			projectConfigPrototypeDefinitionBlockModel := &projectv1.ProjectConfigPrototypeDefinitionBlock{
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Labels: []string{"testString"},
				Environment: core.StringPtr("testString"),
				Authorizations: projectConfigAuthModel,
				ComplianceProfile: projectComplianceProfileModel,
				LocatorID: core.StringPtr("testString"),
				Inputs: inputVariableModel,
				Settings: projectConfigSettingModel,
			}

			schematicsWorkspaceModel := &projectv1.SchematicsWorkspace{
				WorkspaceID: core.StringPtr("testString"),
			}

			projectConfigPrototypeModel := &projectv1.ProjectConfigPrototype{
				Definition: projectConfigPrototypeDefinitionBlockModel,
				Schematics: schematicsWorkspaceModel,
			}

			createProjectOptions := &projectv1.CreateProjectOptions{
				Definition: projectPrototypeDefinitionModel,
				Location: core.StringPtr("us-south"),
				ResourceGroup: core.StringPtr("Default"),
				Configs: []projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel},
			}

			project, response, err := projectService.CreateProject(createProjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(project).ToNot(BeNil())

			projectIdLink = *project.ID
			fmt.Fprintf(GinkgoWriter, "Saved projectIdLink value: %v\n", projectIdLink)
		})
	})

	Describe(`CreateConfig - Add a new configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateConfig(createConfigOptions *CreateConfigOptions)`, func() {
			projectConfigAuthModel := &projectv1.ProjectConfigAuth{
				TrustedProfileID: core.StringPtr("testString"),
				Method: core.StringPtr("api_key"),
				ApiKey: core.StringPtr("testString"),
			}

			projectComplianceProfileModel := &projectv1.ProjectComplianceProfile{
				ID: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				InstanceLocation: core.StringPtr("testString"),
				AttachmentID: core.StringPtr("testString"),
				ProfileName: core.StringPtr("testString"),
			}

			inputVariableModel := &projectv1.InputVariable{
			}
			inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name[\"account-stage\"].inputs.account_id`))
			inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
			inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
			inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
			inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))

			projectConfigSettingModel := &projectv1.ProjectConfigSetting{
			}
			projectConfigSettingModel.SetProperty("IBMCLOUD_TOOLCHAIN_ENDPOINT", core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com"))

			projectConfigPrototypeDefinitionBlockModel := &projectv1.ProjectConfigPrototypeDefinitionBlock{
				Name: core.StringPtr("env-stage"),
				Description: core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace."),
				Labels: []string{"env:stage", "governance:test", "build:0"},
				Environment: core.StringPtr("testString"),
				Authorizations: projectConfigAuthModel,
				ComplianceProfile: projectComplianceProfileModel,
				LocatorID: core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"),
				Inputs: inputVariableModel,
				Settings: projectConfigSettingModel,
			}

			schematicsWorkspaceModel := &projectv1.SchematicsWorkspace{
				WorkspaceID: core.StringPtr("testString"),
			}

			createConfigOptions := &projectv1.CreateConfigOptions{
				ProjectID: &projectIdLink,
				Definition: projectConfigPrototypeDefinitionBlockModel,
				Schematics: schematicsWorkspaceModel,
			}

			projectConfig, response, err := projectService.CreateConfig(createConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(projectConfig).ToNot(BeNil())

			configIdLink = *projectConfig.ID
			fmt.Fprintf(GinkgoWriter, "Saved configIdLink value: %v\n", configIdLink)
		})
	})

	Describe(`ListProjects - List projects`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProjects(listProjectsOptions *ListProjectsOptions) with pagination`, func(){
			listProjectsOptions := &projectv1.ListProjectsOptions{
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			listProjectsOptions.Start = nil
			listProjectsOptions.Limit = core.Int64Ptr(1)

			var allResults []projectv1.ProjectCollectionMemberWithMetadata
			for {
				projectCollection, response, err := projectService.ListProjects(listProjectsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(projectCollection).ToNot(BeNil())
				allResults = append(allResults, projectCollection.Projects...)

				listProjectsOptions.Start, err = projectCollection.GetNextStart()
				Expect(err).To(BeNil())

				if listProjectsOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListProjects(listProjectsOptions *ListProjectsOptions) using ProjectsPager`, func(){
			listProjectsOptions := &projectv1.ListProjectsOptions{
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := projectService.NewProjectsPager(listProjectsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []projectv1.ProjectCollectionMemberWithMetadata
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = projectService.NewProjectsPager(listProjectsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListProjects() returned a total of %d item(s) using ProjectsPager.\n", len(allResults))
		})
	})

	Describe(`GetProject - Get a project`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProject(getProjectOptions *GetProjectOptions)`, func() {
			getProjectOptions := &projectv1.GetProjectOptions{
				ID: &projectIdLink,
			}

			project, response, err := projectService.GetProject(getProjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(project).ToNot(BeNil())
		})
	})

	Describe(`UpdateProject - Update a project`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateProject(updateProjectOptions *UpdateProjectOptions)`, func() {
			projectPrototypePatchDefinitionBlockModel := &projectv1.ProjectPrototypePatchDefinitionBlock{
				Name: core.StringPtr("acme-microservice"),
				Description: core.StringPtr("A microservice to deploy on top of ACME infrastructure."),
				DestroyOnDelete: core.BoolPtr(true),
			}

			updateProjectOptions := &projectv1.UpdateProjectOptions{
				ID: &projectIdLink,
				Definition: projectPrototypePatchDefinitionBlockModel,
			}

			project, response, err := projectService.UpdateProject(updateProjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(project).ToNot(BeNil())
		})
	})

	Describe(`CreateProjectEnvironment - Create an environment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProjectEnvironment(createProjectEnvironmentOptions *CreateProjectEnvironmentOptions)`, func() {
			projectConfigAuthModel := &projectv1.ProjectConfigAuth{
				TrustedProfileID: core.StringPtr("testString"),
				Method: core.StringPtr("api_key"),
				ApiKey: core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9"),
			}

			inputVariableModel := &projectv1.InputVariable{
			}
			inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
			inputVariableModel.SetProperty("region", core.StringPtr("us-south"))

			projectComplianceProfileModel := &projectv1.ProjectComplianceProfile{
				ID: core.StringPtr("some-profile-id"),
				InstanceID: core.StringPtr("some-instance-id"),
				InstanceLocation: core.StringPtr("us-south"),
				AttachmentID: core.StringPtr("some-attachment-id"),
				ProfileName: core.StringPtr("some-profile-name"),
			}

			environmentDefinitionRequiredPropertiesModel := &projectv1.EnvironmentDefinitionRequiredProperties{
				Name: core.StringPtr("development"),
				Description: core.StringPtr("The environment 'development'"),
				Authorizations: projectConfigAuthModel,
				Inputs: inputVariableModel,
				ComplianceProfile: projectComplianceProfileModel,
			}

			createProjectEnvironmentOptions := &projectv1.CreateProjectEnvironmentOptions{
				ProjectID: &projectIdLink,
				Definition: environmentDefinitionRequiredPropertiesModel,
			}

			environment, response, err := projectService.CreateProjectEnvironment(createProjectEnvironmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(environment).ToNot(BeNil())
		})
	})

	Describe(`ListProjectEnvironments - List environments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProjectEnvironments(listProjectEnvironmentsOptions *ListProjectEnvironmentsOptions)`, func() {
			listProjectEnvironmentsOptions := &projectv1.ListProjectEnvironmentsOptions{
				ProjectID: &projectIdLink,
			}

			environmentListResponse, response, err := projectService.ListProjectEnvironments(listProjectEnvironmentsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environmentListResponse).ToNot(BeNil())
		})
	})

	Describe(`GetProjectEnvironment - Get an environment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProjectEnvironment(getProjectEnvironmentOptions *GetProjectEnvironmentOptions)`, func() {
			getProjectEnvironmentOptions := &projectv1.GetProjectEnvironmentOptions{
				ProjectID: &projectIdLink,
				ID: &projectIdLink,
			}

			environment, response, err := projectService.GetProjectEnvironment(getProjectEnvironmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environment).ToNot(BeNil())
		})
	})

	Describe(`UpdateProjectEnvironment - Update an environment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateProjectEnvironment(updateProjectEnvironmentOptions *UpdateProjectEnvironmentOptions)`, func() {
			projectConfigAuthModel := &projectv1.ProjectConfigAuth{
				TrustedProfileID: core.StringPtr("testString"),
				Method: core.StringPtr("api_key"),
				ApiKey: core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9"),
			}

			inputVariableModel := &projectv1.InputVariable{
			}
			inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
			inputVariableModel.SetProperty("region", core.StringPtr("us-south"))

			projectComplianceProfileModel := &projectv1.ProjectComplianceProfile{
				ID: core.StringPtr("some-profile-id"),
				InstanceID: core.StringPtr("some-instance-id"),
				InstanceLocation: core.StringPtr("us-south"),
				AttachmentID: core.StringPtr("some-attachment-id"),
				ProfileName: core.StringPtr("some-profile-name"),
			}

			environmentDefinitionPropertiesModel := &projectv1.EnvironmentDefinitionProperties{
				Name: core.StringPtr("development"),
				Description: core.StringPtr("The environment 'development'"),
				Authorizations: projectConfigAuthModel,
				Inputs: inputVariableModel,
				ComplianceProfile: projectComplianceProfileModel,
			}

			updateProjectEnvironmentOptions := &projectv1.UpdateProjectEnvironmentOptions{
				ProjectID: &projectIdLink,
				ID: &projectIdLink,
				Definition: environmentDefinitionPropertiesModel,
			}

			environment, response, err := projectService.UpdateProjectEnvironment(updateProjectEnvironmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environment).ToNot(BeNil())
		})
	})

	Describe(`ListConfigs - List all project configurations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListConfigs(listConfigsOptions *ListConfigsOptions)`, func() {
			listConfigsOptions := &projectv1.ListConfigsOptions{
				ProjectID: &projectIdLink,
			}

			projectConfigCollection, response, err := projectService.ListConfigs(listConfigsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigCollection).ToNot(BeNil())
		})
	})

	Describe(`GetConfig - Get a project configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetConfig(getConfigOptions *GetConfigOptions)`, func() {
			getConfigOptions := &projectv1.GetConfigOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
			}

			projectConfig, response, err := projectService.GetConfig(getConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfig).ToNot(BeNil())
		})
	})

	Describe(`UpdateConfig - Update a configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateConfig(updateConfigOptions *UpdateConfigOptions)`, func() {
			projectConfigAuthModel := &projectv1.ProjectConfigAuth{
				TrustedProfileID: core.StringPtr("testString"),
				Method: core.StringPtr("api_key"),
				ApiKey: core.StringPtr("testString"),
			}

			projectComplianceProfileModel := &projectv1.ProjectComplianceProfile{
				ID: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				InstanceLocation: core.StringPtr("testString"),
				AttachmentID: core.StringPtr("testString"),
				ProfileName: core.StringPtr("testString"),
			}

			inputVariableModel := &projectv1.InputVariable{
			}
			inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name[\"account-stage\"].inputs.account_id`))
			inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
			inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
			inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
			inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))

			projectConfigSettingModel := &projectv1.ProjectConfigSetting{
			}
			projectConfigSettingModel.SetProperty("foo", core.StringPtr("testString"))

			projectConfigPrototypePatchDefinitionBlockModel := &projectv1.ProjectConfigPrototypePatchDefinitionBlock{
				Name: core.StringPtr("testString"),
				Description: core.StringPtr("testString"),
				Labels: []string{"testString"},
				Environment: core.StringPtr("testString"),
				Authorizations: projectConfigAuthModel,
				ComplianceProfile: projectComplianceProfileModel,
				LocatorID: core.StringPtr("testString"),
				Inputs: inputVariableModel,
				Settings: projectConfigSettingModel,
			}

			updateConfigOptions := &projectv1.UpdateConfigOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
				Definition: projectConfigPrototypePatchDefinitionBlockModel,
			}

			projectConfig, response, err := projectService.UpdateConfig(updateConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfig).ToNot(BeNil())
		})
	})

	Describe(`ForceApprove - Force approve project configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ForceApprove(forceApproveOptions *ForceApproveOptions)`, func() {
			forceApproveOptions := &projectv1.ForceApproveOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
				Comment: core.StringPtr("Approving the changes"),
			}

			projectConfigVersion, response, err := projectService.ForceApprove(forceApproveOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(projectConfigVersion).ToNot(BeNil())
		})
	})

	Describe(`Approve - Approve and merge a configuration draft`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Approve(approveOptions *ApproveOptions)`, func() {
			approveOptions := &projectv1.ApproveOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
				Comment: core.StringPtr("Approving the changes"),
			}

			projectConfigVersion, response, err := projectService.Approve(approveOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(projectConfigVersion).ToNot(BeNil())
		})
	})

	Describe(`ValidateConfig - Run a validation check`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ValidateConfig(validateConfigOptions *ValidateConfigOptions)`, func() {
			validateConfigOptions := &projectv1.ValidateConfigOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
			}

			projectConfigVersion, response, err := projectService.ValidateConfig(validateConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfigVersion).ToNot(BeNil())
		})
	})

	Describe(`DeployConfig - Deploy a configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeployConfig(deployConfigOptions *DeployConfigOptions)`, func() {
			deployConfigOptions := &projectv1.DeployConfigOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
			}

			projectConfigVersion, response, err := projectService.DeployConfig(deployConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfigVersion).ToNot(BeNil())
		})
	})

	Describe(`UndeployConfig - Destroy configuration resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UndeployConfig(undeployConfigOptions *UndeployConfigOptions)`, func() {
			undeployConfigOptions := &projectv1.UndeployConfigOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
			}

			response, err := projectService.UndeployConfig(undeployConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`SyncConfig - Sync a project configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SyncConfig(syncConfigOptions *SyncConfigOptions)`, func() {
			schematicsWorkspaceModel := &projectv1.SchematicsWorkspace{
				WorkspaceID: core.StringPtr("us-south.workspace.service.e0106139"),
			}

			syncConfigOptions := &projectv1.SyncConfigOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
				Schematics: schematicsWorkspaceModel,
			}

			response, err := projectService.SyncConfig(syncConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`ListConfigResources - List the resources deployed by a configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListConfigResources(listConfigResourcesOptions *ListConfigResourcesOptions)`, func() {
			listConfigResourcesOptions := &projectv1.ListConfigResourcesOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
			}

			projectConfigResourceCollection, response, err := projectService.ListConfigResources(listConfigResourcesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigResourceCollection).ToNot(BeNil())
		})
	})

	Describe(`ListConfigVersions - Get a list of versions of a project configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListConfigVersions(listConfigVersionsOptions *ListConfigVersionsOptions)`, func() {
			listConfigVersionsOptions := &projectv1.ListConfigVersionsOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
			}

			projectConfigVersionSummaryCollection, response, err := projectService.ListConfigVersions(listConfigVersionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigVersionSummaryCollection).ToNot(BeNil())
		})
	})

	Describe(`GetConfigVersion - Get a specific version of a project configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetConfigVersion(getConfigVersionOptions *GetConfigVersionOptions)`, func() {
			getConfigVersionOptions := &projectv1.GetConfigVersionOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
				Version: core.Int64Ptr(int64(38)),
			}

			projectConfigVersion, response, err := projectService.GetConfigVersion(getConfigVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigVersion).ToNot(BeNil())
		})
	})

	Describe(`DeleteProjectEnvironment - Delete an environment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProjectEnvironment(deleteProjectEnvironmentOptions *DeleteProjectEnvironmentOptions)`, func() {
			deleteProjectEnvironmentOptions := &projectv1.DeleteProjectEnvironmentOptions{
				ProjectID: &projectIdLink,
				ID: &projectIdLink,
			}

			environmentDeleteResponse, response, err := projectService.DeleteProjectEnvironment(deleteProjectEnvironmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environmentDeleteResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteConfig - Delete a configuration in a project by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteConfig(deleteConfigOptions *DeleteConfigOptions)`, func() {
			deleteConfigOptions := &projectv1.DeleteConfigOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
			}

			projectConfigDelete, response, err := projectService.DeleteConfig(deleteConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDelete).ToNot(BeNil())
		})
	})

	Describe(`DeleteConfigVersion - Delete a configuration for the specified project ID and version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteConfigVersion(deleteConfigVersionOptions *DeleteConfigVersionOptions)`, func() {
			deleteConfigVersionOptions := &projectv1.DeleteConfigVersionOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
				Version: core.Int64Ptr(int64(38)),
			}

			projectConfigDelete, response, err := projectService.DeleteConfigVersion(deleteConfigVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDelete).ToNot(BeNil())
		})
	})

	Describe(`DeleteProject - Delete a project`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProject(deleteProjectOptions *DeleteProjectOptions)`, func() {
			deleteProjectOptions := &projectv1.DeleteProjectOptions{
				ID: &projectIdLink,
			}

			response, err := projectService.DeleteProject(deleteProjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
