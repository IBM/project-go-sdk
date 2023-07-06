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
			projectConfigAuthTrustedProfileModel := &projectv1.ProjectConfigAuthTrustedProfile{
				ID: core.StringPtr("testString"),
				TargetIamID: core.StringPtr("testString"),
			}

			projectConfigAuthModel := &projectv1.ProjectConfigAuth{
				TrustedProfile: projectConfigAuthTrustedProfileModel,
				Method: core.StringPtr("testString"),
				ApiKey: core.StringPtr("testString"),
			}

			projectConfigComplianceProfileModel := &projectv1.ProjectConfigComplianceProfile{
				ID: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				InstanceLocation: core.StringPtr("testString"),
				AttachmentID: core.StringPtr("testString"),
				ProfileName: core.StringPtr("testString"),
			}

			projectConfigInputVariableModel := &projectv1.ProjectConfigInputVariable{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			projectConfigSettingCollectionModel := &projectv1.ProjectConfigSettingCollection{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			projectConfigPrototypeModel := &projectv1.ProjectConfigPrototype{
				Name: core.StringPtr("common-variables"),
				Labels: []string{},
				Description: core.StringPtr("testString"),
				Authorizations: projectConfigAuthModel,
				ComplianceProfile: projectConfigComplianceProfileModel,
				LocatorID: core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"),
				Input: []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel},
				Setting: []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel},
			}

			createProjectOptions := &projectv1.CreateProjectOptions{
				ResourceGroup: core.StringPtr("Default"),
				Location: core.StringPtr("us-south"),
				Name: core.StringPtr("acme-microservice"),
				Description: core.StringPtr("A microservice to deploy on top of ACME infrastructure."),
				DestroyOnDelete: core.BoolPtr(true),
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
			projectConfigAuthTrustedProfileModel := &projectv1.ProjectConfigAuthTrustedProfile{
				ID: core.StringPtr("testString"),
				TargetIamID: core.StringPtr("testString"),
			}

			projectConfigAuthModel := &projectv1.ProjectConfigAuth{
				TrustedProfile: projectConfigAuthTrustedProfileModel,
				Method: core.StringPtr("testString"),
				ApiKey: core.StringPtr("testString"),
			}

			projectConfigComplianceProfileModel := &projectv1.ProjectConfigComplianceProfile{
				ID: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				InstanceLocation: core.StringPtr("testString"),
				AttachmentID: core.StringPtr("testString"),
				ProfileName: core.StringPtr("testString"),
			}

			projectConfigInputVariableModel := &projectv1.ProjectConfigInputVariable{
				Name: core.StringPtr("account_id"),
				Value: core.StringPtr(`$configs[].name[\"account-stage\"].input.account_id`),
			}

			projectConfigSettingCollectionModel := &projectv1.ProjectConfigSettingCollection{
				Name: core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT"),
				Value: core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com"),
			}

			createConfigOptions := &projectv1.CreateConfigOptions{
				ProjectID: &projectIdLink,
				Name: core.StringPtr("env-stage"),
				LocatorID: core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"),
				Labels: []string{"env:stage", "governance:test", "build:0"},
				Description: core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace."),
				Authorizations: projectConfigAuthModel,
				ComplianceProfile: projectConfigComplianceProfileModel,
				Input: []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel},
				Setting: []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel},
			}

			projectConfigDraftResponse, response, err := projectService.CreateConfig(createConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(projectConfigDraftResponse).ToNot(BeNil())

			configIdLink = *projectConfigDraftResponse.ID
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

			projectSummary, response, err := projectService.GetProject(getProjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectSummary).ToNot(BeNil())
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

			projectConfigGetResponse, response, err := projectService.GetConfig(getConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigGetResponse).ToNot(BeNil())
		})
	})

	Describe(`UpdateConfig - Update a configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateConfig(updateConfigOptions *UpdateConfigOptions)`, func() {
			projectConfigInputVariableModel := &projectv1.ProjectConfigInputVariable{
				Name: core.StringPtr("account_id"),
				Value: core.StringPtr(`$configs[].name[\"account-stage\"].input.account_id`),
			}

			projectConfigSettingCollectionModel := &projectv1.ProjectConfigSettingCollection{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			projectConfigAuthTrustedProfileModel := &projectv1.ProjectConfigAuthTrustedProfile{
				ID: core.StringPtr("testString"),
				TargetIamID: core.StringPtr("testString"),
			}

			projectConfigAuthModel := &projectv1.ProjectConfigAuth{
				TrustedProfile: projectConfigAuthTrustedProfileModel,
				Method: core.StringPtr("testString"),
				ApiKey: core.StringPtr("testString"),
			}

			projectConfigComplianceProfileModel := &projectv1.ProjectConfigComplianceProfile{
				ID: core.StringPtr("testString"),
				InstanceID: core.StringPtr("testString"),
				InstanceLocation: core.StringPtr("testString"),
				AttachmentID: core.StringPtr("testString"),
				ProfileName: core.StringPtr("testString"),
			}

			updateConfigOptions := &projectv1.UpdateConfigOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
				LocatorID: core.StringPtr("testString"),
				Input: []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel},
				Setting: []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel},
				Name: core.StringPtr("testString"),
				Labels: []string{"testString"},
				Description: core.StringPtr("testString"),
				Authorizations: projectConfigAuthModel,
				ComplianceProfile: projectConfigComplianceProfileModel,
			}

			projectConfigDraftResponse, response, err := projectService.UpdateConfig(updateConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDraftResponse).ToNot(BeNil())
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

			projectConfigGetResponse, response, err := projectService.Approve(approveOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(projectConfigGetResponse).ToNot(BeNil())
		})
	})

	Describe(`CheckConfig - Run a validation check`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CheckConfig(checkConfigOptions *CheckConfigOptions)`, func() {
			checkConfigOptions := &projectv1.CheckConfigOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
				XAuthRefreshToken: core.StringPtr("testString"),
				IsDraft: core.BoolPtr(true),
			}

			projectConfigGetResponse, response, err := projectService.CheckConfig(checkConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfigGetResponse).ToNot(BeNil())
		})
	})

	Describe(`InstallConfig - Deploy a configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InstallConfig(installConfigOptions *InstallConfigOptions)`, func() {
			installConfigOptions := &projectv1.InstallConfigOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
			}

			projectConfigGetResponse, response, err := projectService.InstallConfig(installConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfigGetResponse).ToNot(BeNil())
		})
	})

	Describe(`UninstallConfig - Destroy configuration resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UninstallConfig(uninstallConfigOptions *UninstallConfigOptions)`, func() {
			uninstallConfigOptions := &projectv1.UninstallConfigOptions{
				ProjectID: &projectIdLink,
				ID: &configIdLink,
			}

			response, err := projectService.UninstallConfig(uninstallConfigOptions)
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

	Describe(`ListConfigDrafts - Get a list of project configuration drafts`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListConfigDrafts(listConfigDraftsOptions *ListConfigDraftsOptions)`, func() {
			listConfigDraftsOptions := &projectv1.ListConfigDraftsOptions{
				ProjectID: &projectIdLink,
				ConfigID: core.StringPtr("testString"),
			}

			projectConfigDraftSummaryCollection, response, err := projectService.ListConfigDrafts(listConfigDraftsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDraftSummaryCollection).ToNot(BeNil())
		})
	})

	Describe(`GetConfigDraft - Get a project configuration draft`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetConfigDraft(getConfigDraftOptions *GetConfigDraftOptions)`, func() {
			getConfigDraftOptions := &projectv1.GetConfigDraftOptions{
				ProjectID: &projectIdLink,
				ConfigID: core.StringPtr("testString"),
				Version: core.Int64Ptr(int64(38)),
			}

			projectConfigDraftResponse, response, err := projectService.GetConfigDraft(getConfigDraftOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDraftResponse).ToNot(BeNil())
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
				DraftOnly: core.BoolPtr(false),
			}

			projectConfigDelete, response, err := projectService.DeleteConfig(deleteConfigOptions)
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
