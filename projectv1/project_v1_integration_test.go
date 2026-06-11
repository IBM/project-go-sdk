//go:build integration

/**
 * (C) Copyright IBM Corp. 2026.
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
			projectDefinitionStoreModel := &projectv1.ProjectDefinitionStore{
				Type: core.StringPtr("gh"),
				URL: core.StringPtr("testString"),
				Token: core.StringPtr("testString"),
				ConfigDirectory: core.StringPtr("testString"),
			}

			projectTerraformEngineSettingsModel := &projectv1.ProjectTerraformEngineSettings{
				ID: core.StringPtr("testString"),
				Type: core.StringPtr("terraform-enterprise"),
			}

			projectPrototypeDefinitionModel := &projectv1.ProjectPrototypeDefinition{
				Name: core.StringPtr("acme-microservice"),
				Description: core.StringPtr("A microservice to deploy on top of ACME infrastructure."),
				AutoDeployMode: core.StringPtr("manual_approval"),
				MonitoringEnabled: core.BoolPtr(false),
				DestroyOnDelete: core.BoolPtr(true),
				Owner: core.StringPtr("testString"),
				Store: projectDefinitionStoreModel,
				TerraformEngine: projectTerraformEngineSettingsModel,
				AutoDeploy: core.BoolPtr(false),
			}

			projectComplianceProfileModel := &projectv1.ProjectComplianceProfileNullableObject{
			}

			stackMemberModel := &projectv1.StackMember{
				Name: core.StringPtr("testString"),
				ConfigID: core.StringPtr("testString"),
			}

			projectConfigUsesModel := &projectv1.ProjectConfigUses{
				ConfigID: core.StringPtr("testString"),
				ProjectID: core.StringPtr("testString"),
			}

			projectConfigAuthModel := &projectv1.ProjectConfigAuth{
				TrustedProfileID: core.StringPtr("testString"),
				Method: core.StringPtr("testString"),
				ApiKey: core.StringPtr("testString"),
			}

			projectConfigDefinitionPrototypeModel := &projectv1.ProjectConfigDefinitionPrototypeDAConfigDefinitionPropertiesPrototype{
				ComplianceProfile: projectComplianceProfileModel,
				LocatorID: core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"),
				Members: []projectv1.StackMember{*stackMemberModel},
				Uses: []projectv1.ProjectConfigUses{*projectConfigUsesModel},
				Description: core.StringPtr("The stage account configuration."),
				Name: core.StringPtr("account-stage"),
				Authorizations: projectConfigAuthModel,
				Inputs: map[string]interface{}{"anyKey": "anyValue"},
				Settings: map[string]interface{}{"anyKey": "anyValue"},
				EnvironmentID: core.StringPtr("testString"),
			}

			schematicsWorkspaceModel := &projectv1.SchematicsWorkspace{
				WorkspaceCrn: core.StringPtr("crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"),
			}

			projectConfigPrototypeModel := &projectv1.ProjectConfigPrototype{
				Definition: projectConfigDefinitionPrototypeModel,
				Schematics: schematicsWorkspaceModel,
			}

			environmentDefinitionRequiredPropertiesModel := &projectv1.EnvironmentDefinitionRequiredProperties{
				Description: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				Authorizations: projectConfigAuthModel,
				Inputs: map[string]interface{}{"anyKey": "anyValue"},
				ComplianceProfile: projectComplianceProfileModel,
			}

			environmentPrototypeModel := &projectv1.EnvironmentPrototype{
				Definition: environmentDefinitionRequiredPropertiesModel,
			}

			createProjectOptions := &projectv1.CreateProjectOptions{
				Definition: projectPrototypeDefinitionModel,
				Location: core.StringPtr("us-south"),
				ResourceGroup: core.StringPtr("Default"),
				Configs: []projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel},
				Environments: []projectv1.EnvironmentPrototype{*environmentPrototypeModel},
			}

			project, response, err := projectService.CreateProject(createProjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(project).ToNot(BeNil())
		})
	})

	Describe(`ListProjects - List projects`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProjects(listProjectsOptions *ListProjectsOptions) with pagination`, func(){
			listProjectsOptions := &projectv1.ListProjectsOptions{
				Token: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			listProjectsOptions.Token = nil
			listProjectsOptions.Limit = core.Int64Ptr(1)

			var allResults []projectv1.ProjectSummary
			for {
				projectCollection, response, err := projectService.ListProjects(listProjectsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(projectCollection).ToNot(BeNil())
				allResults = append(allResults, projectCollection.Projects...)

				listProjectsOptions.Token, err = projectCollection.GetNextToken()
				Expect(err).To(BeNil())

				if listProjectsOptions.Token == nil {
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

			var allResults []projectv1.ProjectSummary
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
				ID: core.StringPtr("testString"),
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
			projectDefinitionStoreModel := &projectv1.ProjectDefinitionStore{
				Type: core.StringPtr("gh"),
				URL: core.StringPtr("testString"),
				Token: core.StringPtr("testString"),
				ConfigDirectory: core.StringPtr("testString"),
			}

			projectTerraformEngineSettingsModel := &projectv1.ProjectTerraformEngineSettings{
				ID: core.StringPtr("testString"),
				Type: core.StringPtr("terraform-enterprise"),
			}

			projectDefinitionPatchModel := &projectv1.ProjectDefinitionPatch{
				Name: core.StringPtr("acme-microservice"),
				Description: core.StringPtr("A microservice to deploy on top of ACME infrastructure."),
				AutoDeployMode: core.StringPtr("auto_approval"),
				MonitoringEnabled: core.BoolPtr(true),
				DestroyOnDelete: core.BoolPtr(true),
				Owner: core.StringPtr("IBMid-31982730TY"),
				Store: projectDefinitionStoreModel,
				TerraformEngine: projectTerraformEngineSettingsModel,
				AutoDeploy: core.BoolPtr(true),
			}

			updateProjectOptions := &projectv1.UpdateProjectOptions{
				ID: core.StringPtr("testString"),
				Definition: projectDefinitionPatchModel,
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
				TrustedProfileID: core.StringPtr("Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12"),
				Method: core.StringPtr("trusted_profile"),
				ApiKey: core.StringPtr("testString"),
			}

			projectComplianceProfileModel := &projectv1.ProjectComplianceProfileV1{
				ID: core.StringPtr("some-profile-id"),
				InstanceID: core.StringPtr("some-instance-id"),
				InstanceLocation: core.StringPtr("us-south"),
				AttachmentID: core.StringPtr("some-attachment-id"),
				ProfileName: core.StringPtr("some-profile-name"),
				WpPolicyID: core.StringPtr("testString"),
				WpInstanceID: core.StringPtr("testString"),
				WpInstanceName: core.StringPtr("testString"),
				WpInstanceLocation: core.StringPtr("us-south"),
				WpZoneID: core.StringPtr("testString"),
				WpZoneName: core.StringPtr("testString"),
				WpPolicyName: core.StringPtr("testString"),
			}

			environmentDefinitionRequiredPropertiesModel := &projectv1.EnvironmentDefinitionRequiredProperties{
				Description: core.StringPtr("The environment development."),
				Name: core.StringPtr("development"),
				Authorizations: projectConfigAuthModel,
				Inputs: map[string]interface{}{"anyKey": "anyValue"},
				ComplianceProfile: projectComplianceProfileModel,
			}

			createProjectEnvironmentOptions := &projectv1.CreateProjectEnvironmentOptions{
				ProjectID: core.StringPtr("testString"),
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
		It(`ListProjectEnvironments(listProjectEnvironmentsOptions *ListProjectEnvironmentsOptions) with pagination`, func(){
			listProjectEnvironmentsOptions := &projectv1.ListProjectEnvironmentsOptions{
				ProjectID: core.StringPtr("testString"),
				Token: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			listProjectEnvironmentsOptions.Token = nil
			listProjectEnvironmentsOptions.Limit = core.Int64Ptr(1)

			var allResults []projectv1.Environment
			for {
				environmentCollection, response, err := projectService.ListProjectEnvironments(listProjectEnvironmentsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(environmentCollection).ToNot(BeNil())
				allResults = append(allResults, environmentCollection.Environments...)

				listProjectEnvironmentsOptions.Token, err = environmentCollection.GetNextToken()
				Expect(err).To(BeNil())

				if listProjectEnvironmentsOptions.Token == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListProjectEnvironments(listProjectEnvironmentsOptions *ListProjectEnvironmentsOptions) using ProjectEnvironmentsPager`, func(){
			listProjectEnvironmentsOptions := &projectv1.ListProjectEnvironmentsOptions{
				ProjectID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := projectService.NewProjectEnvironmentsPager(listProjectEnvironmentsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []projectv1.Environment
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = projectService.NewProjectEnvironmentsPager(listProjectEnvironmentsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListProjectEnvironments() returned a total of %d item(s) using ProjectEnvironmentsPager.\n", len(allResults))
		})
	})

	Describe(`GetProjectEnvironment - Get an environment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProjectEnvironment(getProjectEnvironmentOptions *GetProjectEnvironmentOptions)`, func() {
			getProjectEnvironmentOptions := &projectv1.GetProjectEnvironmentOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
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
				TrustedProfileID: core.StringPtr("Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12"),
				Method: core.StringPtr("trusted_profile"),
				ApiKey: core.StringPtr("testString"),
			}

			projectComplianceProfileModel := &projectv1.ProjectComplianceProfileV1{
				ID: core.StringPtr("some-profile-id"),
				InstanceID: core.StringPtr("some-instance-id"),
				InstanceLocation: core.StringPtr("us-south"),
				AttachmentID: core.StringPtr("some-attachment-id"),
				ProfileName: core.StringPtr("some-profile-name"),
				WpPolicyID: core.StringPtr("testString"),
				WpInstanceID: core.StringPtr("testString"),
				WpInstanceName: core.StringPtr("testString"),
				WpInstanceLocation: core.StringPtr("us-south"),
				WpZoneID: core.StringPtr("testString"),
				WpZoneName: core.StringPtr("testString"),
				WpPolicyName: core.StringPtr("testString"),
			}

			environmentDefinitionPropertiesPatchModel := &projectv1.EnvironmentDefinitionPropertiesPatch{
				Description: core.StringPtr("The environment development."),
				Name: core.StringPtr("development"),
				Authorizations: projectConfigAuthModel,
				Inputs: map[string]interface{}{"anyKey": "anyValue"},
				ComplianceProfile: projectComplianceProfileModel,
			}

			updateProjectEnvironmentOptions := &projectv1.UpdateProjectEnvironmentOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Definition: environmentDefinitionPropertiesPatchModel,
			}

			environment, response, err := projectService.UpdateProjectEnvironment(updateProjectEnvironmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environment).ToNot(BeNil())
		})
	})

	Describe(`CreateConfig - Add a new configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateConfig(createConfigOptions *CreateConfigOptions)`, func() {
			projectComplianceProfileModel := &projectv1.ProjectComplianceProfileNullableObject{
			}

			stackMemberModel := &projectv1.StackMember{
				Name: core.StringPtr("testString"),
				ConfigID: core.StringPtr("testString"),
			}

			projectConfigUsesModel := &projectv1.ProjectConfigUses{
				ConfigID: core.StringPtr("testString"),
				ProjectID: core.StringPtr("testString"),
			}

			projectConfigAuthModel := &projectv1.ProjectConfigAuth{
				TrustedProfileID: core.StringPtr("testString"),
				Method: core.StringPtr("testString"),
				ApiKey: core.StringPtr("testString"),
			}

			projectConfigDefinitionPrototypeModel := &projectv1.ProjectConfigDefinitionPrototypeDAConfigDefinitionPropertiesPrototype{
				ComplianceProfile: projectComplianceProfileModel,
				LocatorID: core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"),
				Members: []projectv1.StackMember{*stackMemberModel},
				Uses: []projectv1.ProjectConfigUses{*projectConfigUsesModel},
				Description: core.StringPtr("The stage environment configuration."),
				Name: core.StringPtr("env-stage"),
				Authorizations: projectConfigAuthModel,
				Inputs: map[string]interface{}{"anyKey": "anyValue"},
				Settings: map[string]interface{}{"anyKey": "anyValue"},
				EnvironmentID: core.StringPtr("testString"),
			}

			schematicsWorkspaceModel := &projectv1.SchematicsWorkspace{
				WorkspaceCrn: core.StringPtr("crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"),
			}

			createConfigOptions := &projectv1.CreateConfigOptions{
				ProjectID: core.StringPtr("testString"),
				Definition: projectConfigDefinitionPrototypeModel,
				Schematics: schematicsWorkspaceModel,
			}

			projectConfig, response, err := projectService.CreateConfig(createConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(projectConfig).ToNot(BeNil())
		})
	})

	Describe(`ListConfigs - List all project configurations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListConfigs(listConfigsOptions *ListConfigsOptions) with pagination`, func(){
			listConfigsOptions := &projectv1.ListConfigsOptions{
				ProjectID: core.StringPtr("testString"),
				Token: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			listConfigsOptions.Token = nil
			listConfigsOptions.Limit = core.Int64Ptr(1)

			var allResults []projectv1.ProjectConfigSummary
			for {
				projectConfigCollection, response, err := projectService.ListConfigs(listConfigsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(projectConfigCollection).ToNot(BeNil())
				allResults = append(allResults, projectConfigCollection.Configs...)

				listConfigsOptions.Token, err = projectConfigCollection.GetNextToken()
				Expect(err).To(BeNil())

				if listConfigsOptions.Token == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListConfigs(listConfigsOptions *ListConfigsOptions) using ConfigsPager`, func(){
			listConfigsOptions := &projectv1.ListConfigsOptions{
				ProjectID: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(10)),
			}

			// Test GetNext().
			pager, err := projectService.NewConfigsPager(listConfigsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []projectv1.ProjectConfigSummary
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = projectService.NewConfigsPager(listConfigsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListConfigs() returned a total of %d item(s) using ConfigsPager.\n", len(allResults))
		})
	})

	Describe(`GetConfig - Get a project configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetConfig(getConfigOptions *GetConfigOptions)`, func() {
			getConfigOptions := &projectv1.GetConfigOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
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
			projectComplianceProfileModel := &projectv1.ProjectComplianceProfileNullableObject{
			}

			stackMemberModel := &projectv1.StackMember{
				Name: core.StringPtr("testString"),
				ConfigID: core.StringPtr("testString"),
			}

			projectConfigUsesModel := &projectv1.ProjectConfigUses{
				ConfigID: core.StringPtr("testString"),
				ProjectID: core.StringPtr("testString"),
			}

			projectConfigAuthModel := &projectv1.ProjectConfigAuth{
				TrustedProfileID: core.StringPtr("testString"),
				Method: core.StringPtr("testString"),
				ApiKey: core.StringPtr("testString"),
			}

			projectConfigDefinitionPatchModel := &projectv1.ProjectConfigDefinitionPatchDAConfigDefinitionPropertiesPatch{
				ComplianceProfile: projectComplianceProfileModel,
				LocatorID: core.StringPtr("testString"),
				Members: []projectv1.StackMember{*stackMemberModel},
				Uses: []projectv1.ProjectConfigUses{*projectConfigUsesModel},
				Description: core.StringPtr("testString"),
				Name: core.StringPtr("env-stage"),
				Authorizations: projectConfigAuthModel,
				Inputs: map[string]interface{}{"anyKey": "anyValue"},
				Settings: map[string]interface{}{"anyKey": "anyValue"},
				EnvironmentID: core.StringPtr("testString"),
			}

			updateConfigOptions := &projectv1.UpdateConfigOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Definition: projectConfigDefinitionPatchModel,
			}

			projectConfig, response, err := projectService.UpdateConfig(updateConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfig).ToNot(BeNil())
		})
	})

	Describe(`ForceApprove - Force approve a project configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ForceApprove(forceApproveOptions *ForceApproveOptions)`, func() {
			forceApproveOptions := &projectv1.ForceApproveOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Comment: core.StringPtr("Approving the changes"),
			}

			projectConfigVersion, response, err := projectService.ForceApprove(forceApproveOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigVersion).ToNot(BeNil())
		})
	})

	Describe(`Approve - Approve and merge a configuration draft`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Approve(approveOptions *ApproveOptions)`, func() {
			approveOptions := &projectv1.ApproveOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Comment: core.StringPtr("Approving the changes"),
			}

			projectConfigVersion, response, err := projectService.Approve(approveOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigVersion).ToNot(BeNil())
		})
	})

	Describe(`ValidateConfig - Run a validation check`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ValidateConfig(validateConfigOptions *ValidateConfigOptions)`, func() {
			validateConfigOptions := &projectv1.ValidateConfigOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
			}

			projectConfigVersion, response, err := projectService.ValidateConfig(validateConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfigVersion).ToNot(BeNil())
		})
	})

	Describe(`CreatePrevalidate - Run a prevalidation check`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePrevalidate(createPrevalidateOptions *CreatePrevalidateOptions)`, func() {
			projectComplianceProfileModel := &projectv1.ProjectComplianceProfileNullableObject{
			}

			stackMemberModel := &projectv1.StackMember{
				Name: core.StringPtr("testString"),
				ConfigID: core.StringPtr("testString"),
			}

			projectConfigUsesModel := &projectv1.ProjectConfigUses{
				ConfigID: core.StringPtr("testString"),
				ProjectID: core.StringPtr("testString"),
			}

			projectConfigAuthModel := &projectv1.ProjectConfigAuth{
				TrustedProfileID: core.StringPtr("testString"),
				Method: core.StringPtr("testString"),
				ApiKey: core.StringPtr("testString"),
			}

			projectConfigDefinitionPrototypeModel := &projectv1.ProjectConfigDefinitionPrototypeDAConfigDefinitionPropertiesPrototype{
				ComplianceProfile: projectComplianceProfileModel,
				LocatorID: core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"),
				Members: []projectv1.StackMember{*stackMemberModel},
				Uses: []projectv1.ProjectConfigUses{*projectConfigUsesModel},
				Description: core.StringPtr("The stage environment configuration."),
				Name: core.StringPtr("env-stage"),
				Authorizations: projectConfigAuthModel,
				Inputs: map[string]interface{}{"anyKey": "anyValue"},
				Settings: map[string]interface{}{"anyKey": "anyValue"},
				EnvironmentID: core.StringPtr("testString"),
			}

			schematicsWorkspaceModel := &projectv1.SchematicsWorkspace{
				WorkspaceCrn: core.StringPtr("crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"),
			}

			createPrevalidateOptions := &projectv1.CreatePrevalidateOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Definition: projectConfigDefinitionPrototypeModel,
				Schematics: schematicsWorkspaceModel,
			}

			result, response, err := projectService.CreatePrevalidate(createPrevalidateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetPrevalidate - Get prevalidate results`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPrevalidate(getPrevalidateOptions *GetPrevalidateOptions)`, func() {
			getPrevalidateOptions := &projectv1.GetPrevalidateOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				ResultID: core.StringPtr("testString"),
			}

			prevalidateGetResponse, response, err := projectService.GetPrevalidate(getPrevalidateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(prevalidateGetResponse).ToNot(BeNil())
		})
	})

	Describe(`DeployConfig - Deploy a configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeployConfig(deployConfigOptions *DeployConfigOptions)`, func() {
			deployConfigOptions := &projectv1.DeployConfigOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
			}

			projectConfigVersion, response, err := projectService.DeployConfig(deployConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfigVersion).ToNot(BeNil())
		})
	})

	Describe(`UndeployConfig - Undeploy configuration resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UndeployConfig(undeployConfigOptions *UndeployConfigOptions)`, func() {
			undeployConfigOptions := &projectv1.UndeployConfigOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
			}

			projectConfigVersion, response, err := projectService.UndeployConfig(undeployConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfigVersion).ToNot(BeNil())
		})
	})

	Describe(`SyncConfig - Sync a project configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SyncConfig(syncConfigOptions *SyncConfigOptions)`, func() {
			schematicsWorkspaceModel := &projectv1.SchematicsWorkspace{
				WorkspaceCrn: core.StringPtr("crn:v1:staging:public:schematics:us-south:a/38acaf4469814090a4e675dc0c317a0d:95ad49de-ab96-4e7d-a08c-45c38aa448e6:workspace:us-south.workspace.service.e0106139"),
			}

			syncConfigOptions := &projectv1.SyncConfigOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Schematics: schematicsWorkspaceModel,
			}

			response, err := projectService.SyncConfig(syncConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`ListConfigResources - List all deployed resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListConfigResources(listConfigResourcesOptions *ListConfigResourcesOptions)`, func() {
			listConfigResourcesOptions := &projectv1.ListConfigResourcesOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
			}

			projectConfigResourceCollection, response, err := projectService.ListConfigResources(listConfigResourcesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigResourceCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateStackDefinition - Create a stack definition`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateStackDefinition(createStackDefinitionOptions *CreateStackDefinitionOptions)`, func() {
			stackDefinitionInputVariableModel := &projectv1.StackDefinitionInputVariable{
				Name: core.StringPtr("region"),
				Type: core.StringPtr("string"),
				Description: core.StringPtr("The IBM Cloud location where a resource is deployed."),
				Default: core.StringPtr("us-south"),
				Required: core.BoolPtr(true),
				Hidden: core.BoolPtr(false),
			}

			stackDefinitionOutputVariableModel := &projectv1.StackDefinitionOutputVariable{
				Name: core.StringPtr("vpc_cluster_id"),
				Value: core.StringPtr("cluster_id"),
			}

			stackDefinitionBlockPrototypeModel := &projectv1.StackDefinitionBlockPrototype{
				Inputs: []projectv1.StackDefinitionInputVariable{*stackDefinitionInputVariableModel},
				Outputs: []projectv1.StackDefinitionOutputVariable{*stackDefinitionOutputVariableModel},
			}

			createStackDefinitionOptions := &projectv1.CreateStackDefinitionOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				StackDefinition: stackDefinitionBlockPrototypeModel,
			}

			stackDefinition, response, err := projectService.CreateStackDefinition(createStackDefinitionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(stackDefinition).ToNot(BeNil())
		})
	})

	Describe(`GetStackDefinition - Get a stack definition`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetStackDefinition(getStackDefinitionOptions *GetStackDefinitionOptions)`, func() {
			getStackDefinitionOptions := &projectv1.GetStackDefinitionOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
			}

			stackDefinition, response, err := projectService.GetStackDefinition(getStackDefinitionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(stackDefinition).ToNot(BeNil())
		})
	})

	Describe(`UpdateStackDefinition - Update a stack definition`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateStackDefinition(updateStackDefinitionOptions *UpdateStackDefinitionOptions)`, func() {
			stackDefinitionInputVariableModel := &projectv1.StackDefinitionInputVariable{
				Name: core.StringPtr("region"),
				Type: core.StringPtr("string"),
				Description: core.StringPtr("The IBM Cloud location where a resource is deployed."),
				Default: core.StringPtr("eu-gb"),
				Required: core.BoolPtr(true),
				Hidden: core.BoolPtr(false),
			}

			stackDefinitionOutputVariableModel := &projectv1.StackDefinitionOutputVariable{
				Name: core.StringPtr("testString"),
				Value: "testString",
			}

			stackDefinitionBlockPrototypeModel := &projectv1.StackDefinitionBlockPrototype{
				Inputs: []projectv1.StackDefinitionInputVariable{*stackDefinitionInputVariableModel},
				Outputs: []projectv1.StackDefinitionOutputVariable{*stackDefinitionOutputVariableModel},
			}

			updateStackDefinitionOptions := &projectv1.UpdateStackDefinitionOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				StackDefinition: stackDefinitionBlockPrototypeModel,
			}

			stackDefinition, response, err := projectService.UpdateStackDefinition(updateStackDefinitionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(stackDefinition).ToNot(BeNil())
		})
	})

	Describe(`ExportStackDefinition - Export a deployable architecture stack to the private catalog`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ExportStackDefinition(exportStackDefinitionOptions *ExportStackDefinitionOptions)`, func() {
			stackDefinitionExportRequestModel := &projectv1.StackDefinitionExportRequestStackDefinitionExportCatalogRequest{
				CatalogID: core.StringPtr("01e1a9ad-534b-4ab9-996a-b8f2a8653d5c"),
				TargetVersion: core.StringPtr("testString"),
				Variation: core.StringPtr("testString"),
				Label: core.StringPtr("Stack Deployable Architecture"),
				Tags: []string{"testString"},
			}

			exportStackDefinitionOptions := &projectv1.ExportStackDefinitionOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Settings: stackDefinitionExportRequestModel,
			}

			stackDefinitionExportResponse, response, err := projectService.ExportStackDefinition(exportStackDefinitionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(stackDefinitionExportResponse).ToNot(BeNil())
		})
	})

	Describe(`ListConfigVersions - Get a list of project configuration versions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListConfigVersions(listConfigVersionsOptions *ListConfigVersionsOptions)`, func() {
			listConfigVersionsOptions := &projectv1.ListConfigVersionsOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
			}

			projectConfigVersionCollection, response, err := projectService.ListConfigVersions(listConfigVersionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigVersionCollection).ToNot(BeNil())
		})
	})

	Describe(`GetConfigVersion - Get a specific project configuration version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetConfigVersion(getConfigVersionOptions *GetConfigVersionOptions)`, func() {
			getConfigVersionOptions := &projectv1.GetConfigVersionOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Version: core.Int64Ptr(int64(0)),
			}

			projectConfigVersion, response, err := projectService.GetConfigVersion(getConfigVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigVersion).ToNot(BeNil())
		})
	})

	Describe(`DeleteProject - Delete a project`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProject(deleteProjectOptions *DeleteProjectOptions)`, func() {
			deleteProjectOptions := &projectv1.DeleteProjectOptions{
				ID: core.StringPtr("testString"),
			}

			projectDeleteResponse, response, err := projectService.DeleteProject(deleteProjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectDeleteResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteProjectEnvironment - Delete an environment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProjectEnvironment(deleteProjectEnvironmentOptions *DeleteProjectEnvironmentOptions)`, func() {
			deleteProjectEnvironmentOptions := &projectv1.DeleteProjectEnvironmentOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
			}

			environmentDeleteResponse, response, err := projectService.DeleteProjectEnvironment(deleteProjectEnvironmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environmentDeleteResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteConfig - Delete a configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteConfig(deleteConfigOptions *DeleteConfigOptions)`, func() {
			deleteConfigOptions := &projectv1.DeleteConfigOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
			}

			projectConfigDelete, response, err := projectService.DeleteConfig(deleteConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDelete).ToNot(BeNil())
		})
	})

	Describe(`DeleteConfigVersion - Delete a project configuration version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteConfigVersion(deleteConfigVersionOptions *DeleteConfigVersionOptions)`, func() {
			deleteConfigVersionOptions := &projectv1.DeleteConfigVersionOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Version: core.Int64Ptr(int64(0)),
			}

			projectConfigDelete, response, err := projectService.DeleteConfigVersion(deleteConfigVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDelete).ToNot(BeNil())
		})
	})

	Describe(`DeleteConfigVersionV2 - Delete a project configuration version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteConfigVersionV2(deleteConfigVersionV2Options *DeleteConfigVersionV2Options)`, func() {
			deleteConfigVersionV2Options := &projectv1.DeleteConfigVersionV2Options{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Version: core.Int64Ptr(int64(0)),
			}

			projectConfigDelete, response, err := projectService.DeleteConfigVersionV2(deleteConfigVersionV2Options)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDelete).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
