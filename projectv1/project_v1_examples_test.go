// +build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/project-go-sdk/projectv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the project service.
//
// The following configuration properties are assumed to be defined:
// PROJECT_URL=<service base url>
// PROJECT_AUTH_TYPE=iam
// PROJECT_APIKEY=<IAM apikey>
// PROJECT_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`ProjectV1 Examples Tests`, func() {

	const externalConfigFile = "../project_v1.env"

	var (
		projectService *projectv1.ProjectV1
		config       map[string]string

		// Variables to hold link values
		configIdLink string
		projectIdLink string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(projectv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			projectServiceOptions := &projectv1.ProjectV1Options{}

			projectService, err = projectv1.NewProjectV1UsingExternalConfig(projectServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(projectService).ToNot(BeNil())
		})
	})

	Describe(`ProjectV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProject request example`, func() {
			fmt.Println("\nCreateProject() result:")
			// begin-create_project

			projectPrototypeDefinitionModel := &projectv1.ProjectPrototypeDefinition{
				Name: core.StringPtr("acme-microservice"),
				Description: core.StringPtr("A microservice to deploy on top of ACME infrastructure."),
			}

			projectConfigPrototypeDefinitionBlockModel := &projectv1.ProjectConfigPrototypeDefinitionBlock{
				Name: core.StringPtr("common-variables"),
				LocatorID: core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"),
			}

			projectConfigPrototypeModel := &projectv1.ProjectConfigPrototype{
				Definition: projectConfigPrototypeDefinitionBlockModel,
			}

			createProjectOptions := projectService.NewCreateProjectOptions(
				projectPrototypeDefinitionModel,
				"us-south",
				"Default",
			)
			createProjectOptions.SetConfigs([]projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel})

			project, response, err := projectService.CreateProject(createProjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(project, "", "  ")
			fmt.Println(string(b))

			// end-create_project

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(project).ToNot(BeNil())

			projectIdLink = *project.ID
			fmt.Fprintf(GinkgoWriter, "Saved projectIdLink value: %v\n", projectIdLink)
		})
		It(`CreateConfig request example`, func() {
			fmt.Println("\nCreateConfig() result:")
			// begin-create_config

			inputVariableModel := &projectv1.InputVariable{
			}
			inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name["account-stage"].input.account_id`))
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
				LocatorID: core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"),
				Input: inputVariableModel,
				Setting: projectConfigSettingModel,
			}

			createConfigOptions := projectService.NewCreateConfigOptions(
				projectIdLink,
				projectConfigPrototypeDefinitionBlockModel,
			)

			projectConfig, response, err := projectService.CreateConfig(createConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-create_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(projectConfig).ToNot(BeNil())

			configIdLink = *projectConfig.ID
			fmt.Fprintf(GinkgoWriter, "Saved configIdLink value: %v\n", configIdLink)
		})
		It(`ListProjects request example`, func() {
			fmt.Println("\nListProjects() result:")
			// begin-list_projects
			listProjectsOptions := &projectv1.ListProjectsOptions{
				Limit: core.Int64Ptr(int64(10)),
			}

			pager, err := projectService.NewProjectsPager(listProjectsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []projectv1.ProjectCollectionMemberWithMetadata
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_projects
		})
		It(`GetProject request example`, func() {
			fmt.Println("\nGetProject() result:")
			// begin-get_project

			getProjectOptions := projectService.NewGetProjectOptions(
				projectIdLink,
			)

			project, response, err := projectService.GetProject(getProjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(project, "", "  ")
			fmt.Println(string(b))

			// end-get_project

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(project).ToNot(BeNil())
		})
		It(`UpdateProject request example`, func() {
			fmt.Println("\nUpdateProject() result:")
			// begin-update_project

			projectPrototypePatchDefinitionBlockModel := &projectv1.ProjectPrototypePatchDefinitionBlock{
				Name: core.StringPtr("acme-microservice"),
				Description: core.StringPtr("A microservice to deploy on top of ACME infrastructure."),
			}

			updateProjectOptions := projectService.NewUpdateProjectOptions(
				projectIdLink,
				projectPrototypePatchDefinitionBlockModel,
			)

			project, response, err := projectService.UpdateProject(updateProjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(project, "", "  ")
			fmt.Println(string(b))

			// end-update_project

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(project).ToNot(BeNil())
		})
		It(`CreateProjectEnvironment request example`, func() {
			fmt.Println("\nCreateProjectEnvironment() result:")
			// begin-create_project_environment

			projectConfigAuthModel := &projectv1.ProjectConfigAuth{
				Method: core.StringPtr("API_KEY"),
				ApiKey: core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9"),
			}

			inputVariableModel := &projectv1.InputVariable{
			}
			inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
			inputVariableModel.SetProperty("region", core.StringPtr("us-south"))

			createProjectEnvironmentOptions := projectService.NewCreateProjectEnvironmentOptions(
				projectIdLink,
			)
			createProjectEnvironmentOptions.SetName("development")
			createProjectEnvironmentOptions.SetDescription("The environment 'development'")
			createProjectEnvironmentOptions.SetAuthorizations(projectConfigAuthModel)
			createProjectEnvironmentOptions.SetInputs(inputVariableModel)

			environmentResponse, response, err := projectService.CreateProjectEnvironment(createProjectEnvironmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(environmentResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_project_environment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(environmentResponse).ToNot(BeNil())
		})
		It(`ListProjectEnvironments request example`, func() {
			fmt.Println("\nListProjectEnvironments() result:")
			// begin-list_project_environments

			listProjectEnvironmentsOptions := projectService.NewListProjectEnvironmentsOptions(
				projectIdLink,
			)

			environmentListResponse, response, err := projectService.ListProjectEnvironments(listProjectEnvironmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(environmentListResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_project_environments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environmentListResponse).ToNot(BeNil())
		})
		It(`GetProjectEnvironment request example`, func() {
			fmt.Println("\nGetProjectEnvironment() result:")
			// begin-get_project_environment

			getProjectEnvironmentOptions := projectService.NewGetProjectEnvironmentOptions(
				projectIdLink,
				projectIdLink,
			)

			environmentResponse, response, err := projectService.GetProjectEnvironment(getProjectEnvironmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(environmentResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_project_environment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environmentResponse).ToNot(BeNil())
		})
		It(`UpdateProjectEnvironment request example`, func() {
			fmt.Println("\nUpdateProjectEnvironment() result:")
			// begin-update_project_environment

			projectConfigAuthModel := &projectv1.ProjectConfigAuth{
				Method: core.StringPtr("API_KEY"),
				ApiKey: core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9"),
			}

			inputVariableModel := &projectv1.InputVariable{
			}
			inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
			inputVariableModel.SetProperty("region", core.StringPtr("us-south"))

			updateProjectEnvironmentOptions := projectService.NewUpdateProjectEnvironmentOptions(
				projectIdLink,
				projectIdLink,
			)
			updateProjectEnvironmentOptions.SetName("dev")
			updateProjectEnvironmentOptions.SetDescription("The new environment 'dev'")
			updateProjectEnvironmentOptions.SetAuthorizations(projectConfigAuthModel)
			updateProjectEnvironmentOptions.SetInputs(inputVariableModel)

			environmentResponse, response, err := projectService.UpdateProjectEnvironment(updateProjectEnvironmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(environmentResponse, "", "  ")
			fmt.Println(string(b))

			// end-update_project_environment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environmentResponse).ToNot(BeNil())
		})
		It(`ListConfigs request example`, func() {
			fmt.Println("\nListConfigs() result:")
			// begin-list_configs

			listConfigsOptions := projectService.NewListConfigsOptions(
				projectIdLink,
			)

			projectConfigCollection, response, err := projectService.ListConfigs(listConfigsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfigCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_configs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigCollection).ToNot(BeNil())
		})
		It(`GetConfig request example`, func() {
			fmt.Println("\nGetConfig() result:")
			// begin-get_config

			getConfigOptions := projectService.NewGetConfigOptions(
				projectIdLink,
				configIdLink,
			)

			projectConfig, response, err := projectService.GetConfig(getConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-get_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`UpdateConfig request example`, func() {
			fmt.Println("\nUpdateConfig() result:")
			// begin-update_config

			inputVariableModel := &projectv1.InputVariable{
			}
			inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name["account-stage"].input.account_id`))
			inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
			inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
			inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
			inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))

			projectConfigPrototypePatchDefinitionBlockModel := &projectv1.ProjectConfigPrototypePatchDefinitionBlock{
				Input: inputVariableModel,
			}

			updateConfigOptions := projectService.NewUpdateConfigOptions(
				projectIdLink,
				configIdLink,
				projectConfigPrototypePatchDefinitionBlockModel,
			)

			projectConfig, response, err := projectService.UpdateConfig(updateConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-update_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`ForceApprove request example`, func() {
			fmt.Println("\nForceApprove() result:")
			// begin-force_approve

			forceApproveOptions := projectService.NewForceApproveOptions(
				projectIdLink,
				configIdLink,
			)
			forceApproveOptions.SetComment("Approving the changes")

			projectConfig, response, err := projectService.ForceApprove(forceApproveOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-force_approve

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`Approve request example`, func() {
			fmt.Println("\nApprove() result:")
			// begin-approve

			approveOptions := projectService.NewApproveOptions(
				projectIdLink,
				configIdLink,
			)
			approveOptions.SetComment("Approving the changes")

			projectConfig, response, err := projectService.Approve(approveOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-approve

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`ValidateConfig request example`, func() {
			fmt.Println("\nValidateConfig() result:")
			// begin-validate_config

			validateConfigOptions := projectService.NewValidateConfigOptions(
				projectIdLink,
				configIdLink,
			)

			projectConfig, response, err := projectService.ValidateConfig(validateConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-validate_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`DeployConfig request example`, func() {
			fmt.Println("\nDeployConfig() result:")
			// begin-deploy_config

			deployConfigOptions := projectService.NewDeployConfigOptions(
				projectIdLink,
				configIdLink,
			)

			projectConfig, response, err := projectService.DeployConfig(deployConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-deploy_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`UndeployConfig request example`, func() {
			// begin-undeploy_config

			undeployConfigOptions := projectService.NewUndeployConfigOptions(
				projectIdLink,
				configIdLink,
			)

			response, err := projectService.UndeployConfig(undeployConfigOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from UndeployConfig(): %d\n", response.StatusCode)
			}

			// end-undeploy_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`SyncConfig request example`, func() {
			// begin-sync_config

			syncConfigOptions := projectService.NewSyncConfigOptions(
				projectIdLink,
				configIdLink,
			)
			syncConfigOptions.SetWorkspaceID("us-south.workspace.service.e0106139")

			response, err := projectService.SyncConfig(syncConfigOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from SyncConfig(): %d\n", response.StatusCode)
			}

			// end-sync_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`ListConfigResources request example`, func() {
			fmt.Println("\nListConfigResources() result:")
			// begin-list_config_resources

			listConfigResourcesOptions := projectService.NewListConfigResourcesOptions(
				projectIdLink,
				configIdLink,
			)

			projectConfigResourceCollection, response, err := projectService.ListConfigResources(listConfigResourcesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfigResourceCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_config_resources

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigResourceCollection).ToNot(BeNil())
		})
		It(`ListConfigVersions request example`, func() {
			fmt.Println("\nListConfigVersions() result:")
			// begin-list_config_versions

			listConfigVersionsOptions := projectService.NewListConfigVersionsOptions(
				projectIdLink,
				configIdLink,
			)

			projectConfigVersionSummaryCollection, response, err := projectService.ListConfigVersions(listConfigVersionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfigVersionSummaryCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_config_versions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigVersionSummaryCollection).ToNot(BeNil())
		})
		It(`GetConfigVersion request example`, func() {
			fmt.Println("\nGetConfigVersion() result:")
			// begin-get_config_version

			getConfigVersionOptions := projectService.NewGetConfigVersionOptions(
				projectIdLink,
				configIdLink,
				int64(38),
			)

			projectConfig, response, err := projectService.GetConfigVersion(getConfigVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-get_config_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`ListEnvironments request example`, func() {
			fmt.Println("\nListEnvironments() result:")
			// begin-list_environments

			listEnvironmentsOptions := projectService.NewListEnvironmentsOptions()

			environmentListIamResponse, response, err := projectService.ListEnvironments(listEnvironmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(environmentListIamResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_environments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environmentListIamResponse).ToNot(BeNil())
		})
		It(`DeleteProjectEnvironment request example`, func() {
			fmt.Println("\nDeleteProjectEnvironment() result:")
			// begin-delete_project_environment

			deleteProjectEnvironmentOptions := projectService.NewDeleteProjectEnvironmentOptions(
				projectIdLink,
				projectIdLink,
			)

			environmentDeleteResponse, response, err := projectService.DeleteProjectEnvironment(deleteProjectEnvironmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(environmentDeleteResponse, "", "  ")
			fmt.Println(string(b))

			// end-delete_project_environment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environmentDeleteResponse).ToNot(BeNil())
		})
		It(`DeleteConfig request example`, func() {
			fmt.Println("\nDeleteConfig() result:")
			// begin-delete_config

			deleteConfigOptions := projectService.NewDeleteConfigOptions(
				projectIdLink,
				configIdLink,
			)

			projectConfigDelete, response, err := projectService.DeleteConfig(deleteConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfigDelete, "", "  ")
			fmt.Println(string(b))

			// end-delete_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDelete).ToNot(BeNil())
		})
		It(`DeleteConfigVersion request example`, func() {
			fmt.Println("\nDeleteConfigVersion() result:")
			// begin-delete_config_version

			deleteConfigVersionOptions := projectService.NewDeleteConfigVersionOptions(
				projectIdLink,
				configIdLink,
				int64(38),
			)

			projectConfigDelete, response, err := projectService.DeleteConfigVersion(deleteConfigVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfigDelete, "", "  ")
			fmt.Println(string(b))

			// end-delete_config_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDelete).ToNot(BeNil())
		})
		It(`DeleteProject request example`, func() {
			// begin-delete_project

			deleteProjectOptions := projectService.NewDeleteProjectOptions(
				projectIdLink,
			)

			response, err := projectService.DeleteProject(deleteProjectOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteProject(): %d\n", response.StatusCode)
			}

			// end-delete_project

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
