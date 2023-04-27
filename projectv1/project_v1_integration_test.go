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
			projectConfigInputVariableModel := &projectv1.ProjectConfigInputVariable{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			projectConfigSettingCollectionModel := &projectv1.ProjectConfigSettingCollection{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			projectConfigPrototypeModel := &projectv1.ProjectConfigPrototype{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("common-variables"),
				Labels: []string{"testString"},
				Description: core.StringPtr("testString"),
				LocatorID: core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"),
				Input: []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel},
				Setting: []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel},
			}

			createProjectOptions := &projectv1.CreateProjectOptions{
				ResourceGroup: core.StringPtr("testString"),
				Location: core.StringPtr("testString"),
				Name: core.StringPtr("acme-microservice"),
				Description: core.StringPtr("A microservice to deploy on top of ACME infrastructure."),
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
			projectConfigInputVariableModel := &projectv1.ProjectConfigInputVariable{
				Name: core.StringPtr("account_id"),
				Value: core.StringPtr(`$configs[].name[\"account-stage\"].input.account_id`),
			}

			projectConfigSettingCollectionModel := &projectv1.ProjectConfigSettingCollection{
				Name: core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT"),
				Value: core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com"),
			}

			createConfigOptions := &projectv1.CreateConfigOptions{
				ProjectID: core.StringPtr("testString"),
				Name: core.StringPtr("env-stage"),
				LocatorID: core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"),
				ID: core.StringPtr("testString"),
				Labels: []string{"env:stage", "governance:test", "build:0"},
				Description: core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace."),
				Input: []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel},
				Setting: []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel},
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
				Complete: core.BoolPtr(false),
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
				Complete: core.BoolPtr(false),
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
				ExcludeConfigs: core.BoolPtr(false),
				Complete: core.BoolPtr(false),
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
			updateProjectOptions := &projectv1.UpdateProjectOptions{
				ID: core.StringPtr("testString"),
				Name: core.StringPtr("acme-microservice"),
				Description: core.StringPtr("A microservice to deploy on top of ACME infrastructure."),
			}

			project, response, err := projectService.UpdateProject(updateProjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(project).ToNot(BeNil())
		})
	})

	Describe(`ListConfigs - List all project configurations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListConfigs(listConfigsOptions *ListConfigsOptions)`, func() {
			listConfigsOptions := &projectv1.ListConfigsOptions{
				ProjectID: core.StringPtr("testString"),
				Version: core.StringPtr("active"),
				Complete: core.BoolPtr(false),
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
				ProjectID: core.StringPtr("testString"),
				ID: &configIdLink,
				Version: core.StringPtr("active"),
				Complete: core.BoolPtr(false),
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
			projectConfigInputVariableModel := &projectv1.ProjectConfigInputVariable{
				Name: core.StringPtr("account_id"),
				Value: core.StringPtr(`$configs[].name[\"account-stage\"].input.account_id`),
			}

			projectConfigSettingCollectionModel := &projectv1.ProjectConfigSettingCollection{
				Name: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			projectConfigPatchRequestModel := &projectv1.ProjectConfigPatchRequestProjectConfigPatchSchematicsTemplate{
				Name: core.StringPtr("testString"),
				Labels: []string{"testString"},
				Description: core.StringPtr("testString"),
				LocatorID: core.StringPtr("testString"),
				Input: []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel},
				Setting: []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel},
			}

			updateConfigOptions := &projectv1.UpdateConfigOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				ProjectConfig: projectConfigPatchRequestModel,
				Complete: core.BoolPtr(false),
			}

			projectConfig, response, err := projectService.UpdateConfig(updateConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfig).ToNot(BeNil())
		})
	})

	Describe(`GetConfigDiff - Get a diff summary of a project configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetConfigDiff(getConfigDiffOptions *GetConfigDiffOptions)`, func() {
			getConfigDiffOptions := &projectv1.GetConfigDiffOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
			}

			projectConfigDiff, response, err := projectService.GetConfigDiff(getConfigDiffOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDiff).ToNot(BeNil())
		})
	})

	Describe(`ForceApprove - Force approve project configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ForceApprove(forceApproveOptions *ForceApproveOptions)`, func() {
			forceApproveOptions := &projectv1.ForceApproveOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Comment: core.StringPtr("Approving the changes"),
				Complete: core.BoolPtr(false),
			}

			projectConfig, response, err := projectService.ForceApprove(forceApproveOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(projectConfig).ToNot(BeNil())
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
				Complete: core.BoolPtr(false),
			}

			projectConfig, response, err := projectService.Approve(approveOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(projectConfig).ToNot(BeNil())
		})
	})

	Describe(`CheckConfig - Run a validation check`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CheckConfig(checkConfigOptions *CheckConfigOptions)`, func() {
			checkConfigOptions := &projectv1.CheckConfigOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				XAuthRefreshToken: core.StringPtr("testString"),
				Complete: core.BoolPtr(false),
				Version: core.StringPtr("active"),
			}

			projectConfig, response, err := projectService.CheckConfig(checkConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfig).ToNot(BeNil())
		})
	})

	Describe(`InstallConfig - Deploy a configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InstallConfig(installConfigOptions *InstallConfigOptions)`, func() {
			installConfigOptions := &projectv1.InstallConfigOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Complete: core.BoolPtr(false),
			}

			projectConfig, response, err := projectService.InstallConfig(installConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfig).ToNot(BeNil())
		})
	})

	Describe(`UninstallConfig - Destroy configuration resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UninstallConfig(uninstallConfigOptions *UninstallConfigOptions)`, func() {
			uninstallConfigOptions := &projectv1.UninstallConfigOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
			}

			response, err := projectService.UninstallConfig(uninstallConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`GetSchematicsJob - View the latest schematics job`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSchematicsJob(getSchematicsJobOptions *GetSchematicsJobOptions)`, func() {
			getSchematicsJobOptions := &projectv1.GetSchematicsJobOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Action: core.StringPtr("plan"),
				Since: core.Int64Ptr(int64(38)),
			}

			actionJob, response, err := projectService.GetSchematicsJob(getSchematicsJobOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionJob).ToNot(BeNil())
		})
	})

	Describe(`GetCostEstimate - Get the cost estimate`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCostEstimate(getCostEstimateOptions *GetCostEstimateOptions)`, func() {
			getCostEstimateOptions := &projectv1.GetCostEstimateOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				Version: core.StringPtr("active"),
			}

			costEstimate, response, err := projectService.GetCostEstimate(getCostEstimateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(costEstimate).ToNot(BeNil())
		})
	})

	Describe(`PostCrnToken - Creates a project CRN token`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostCrnToken(postCrnTokenOptions *PostCrnTokenOptions)`, func() {
			postCrnTokenOptions := &projectv1.PostCrnTokenOptions{
				ID: core.StringPtr("testString"),
			}

			projectCrnTokenResponse, response, err := projectService.PostCrnToken(postCrnTokenOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectCrnTokenResponse).ToNot(BeNil())
		})
	})

	Describe(`PostNotification - Add notifications`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostNotification(postNotificationOptions *PostNotificationOptions)`, func() {
			notificationEventModel := &projectv1.NotificationEvent{
				Event: core.StringPtr("project.create.failed"),
				Target: core.StringPtr("234234324-3444-4556-224232432"),
				Source: core.StringPtr("id.of.project.service.instance"),
				TriggeredBy: core.StringPtr("user-iam-id"),
				ActionURL: core.StringPtr("actionable/url"),
				Data: map[string]interface{}{"anyKey": "anyValue"},
			}

			postNotificationOptions := &projectv1.PostNotificationOptions{
				ID: core.StringPtr("testString"),
				Notifications: []projectv1.NotificationEvent{*notificationEventModel},
			}

			notificationsPrototypePostResponse, response, err := projectService.PostNotification(postNotificationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsPrototypePostResponse).ToNot(BeNil())
		})
	})

	Describe(`GetNotifications - Get events by project ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetNotifications(getNotificationsOptions *GetNotificationsOptions)`, func() {
			getNotificationsOptions := &projectv1.GetNotificationsOptions{
				ID: core.StringPtr("testString"),
			}

			notificationsGetResponse, response, err := projectService.GetNotifications(getNotificationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsGetResponse).ToNot(BeNil())
		})
	})

	Describe(`PostEventNotificationsIntegration - Connect to a event notifications instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostEventNotificationsIntegration(postEventNotificationsIntegrationOptions *PostEventNotificationsIntegrationOptions)`, func() {
			postEventNotificationsIntegrationOptions := &projectv1.PostEventNotificationsIntegrationOptions{
				ID: core.StringPtr("testString"),
				InstanceCrn: core.StringPtr("CRN of event notifications instance"),
				Description: core.StringPtr("A sample project source."),
				EventNotificationsSourceName: core.StringPtr("project 1 source name for event notifications"),
				Enabled: core.BoolPtr(true),
			}

			notificationsIntegrationPostResponse, response, err := projectService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsIntegrationPostResponse).ToNot(BeNil())
		})
	})

	Describe(`GetEventNotificationsIntegration - Get event notification source details by project ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEventNotificationsIntegration(getEventNotificationsIntegrationOptions *GetEventNotificationsIntegrationOptions)`, func() {
			getEventNotificationsIntegrationOptions := &projectv1.GetEventNotificationsIntegrationOptions{
				ID: core.StringPtr("testString"),
			}

			notificationsIntegrationGetResponse, response, err := projectService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsIntegrationGetResponse).ToNot(BeNil())
		})
	})

	Describe(`PostTestEventNotification - Send notification to event notifications instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PostTestEventNotification(postTestEventNotificationOptions *PostTestEventNotificationOptions)`, func() {
			postTestEventNotificationOptions := &projectv1.PostTestEventNotificationOptions{
				ID: core.StringPtr("testString"),
				Ibmendefaultlong: core.StringPtr("long test notification message"),
				Ibmendefaultshort: core.StringPtr("Test notification"),
			}

			notificationsIntegrationTestPostResponse, response, err := projectService.PostTestEventNotification(postTestEventNotificationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsIntegrationTestPostResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteProject - Delete a project`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProject(deleteProjectOptions *DeleteProjectOptions)`, func() {
			deleteProjectOptions := &projectv1.DeleteProjectOptions{
				ID: core.StringPtr("testString"),
				Destroy: core.BoolPtr(false),
			}

			response, err := projectService.DeleteProject(deleteProjectOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteConfig - Delete a configuration in a project by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteConfig(deleteConfigOptions *DeleteConfigOptions)`, func() {
			deleteConfigOptions := &projectv1.DeleteConfigOptions{
				ProjectID: core.StringPtr("testString"),
				ID: core.StringPtr("testString"),
				DraftOnly: core.BoolPtr(false),
				Destroy: core.BoolPtr(false),
			}

			projectConfigDelete, response, err := projectService.DeleteConfig(deleteConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDelete).ToNot(BeNil())
		})
	})

	Describe(`DeleteEventNotificationsIntegration - Delete an event notifications connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteEventNotificationsIntegration(deleteEventNotificationsIntegrationOptions *DeleteEventNotificationsIntegrationOptions)`, func() {
			deleteEventNotificationsIntegrationOptions := &projectv1.DeleteEventNotificationsIntegrationOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := projectService.DeleteEventNotificationsIntegration(deleteEventNotificationsIntegrationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
