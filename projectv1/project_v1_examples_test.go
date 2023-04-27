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

			projectConfigPrototypeModel := &projectv1.ProjectConfigPrototype{
				Name: core.StringPtr("common-variables"),
				LocatorID: core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"),
			}

			createProjectOptions := projectService.NewCreateProjectOptions(
				"testString",
				"testString",
				"acme-microservice",
			)
			createProjectOptions.SetDescription("A microservice to deploy on top of ACME infrastructure.")
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

			projectConfigInputVariableModel := &projectv1.ProjectConfigInputVariable{
				Name: core.StringPtr("account_id"),
				Value: core.StringPtr(`$configs[].name["account-stage"].input.account_id`),
			}

			projectConfigSettingCollectionModel := &projectv1.ProjectConfigSettingCollection{
				Name: core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT"),
				Value: core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com"),
			}

			createConfigOptions := projectService.NewCreateConfigOptions(
				"testString",
				"env-stage",
				"1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global",
			)
			createConfigOptions.SetLabels([]string{"env:stage", "governance:test", "build:0"})
			createConfigOptions.SetDescription("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
			createConfigOptions.SetInput([]projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel})
			createConfigOptions.SetSetting([]projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel})

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
				Complete: core.BoolPtr(false),
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

			updateProjectOptions := projectService.NewUpdateProjectOptions(
				"testString",
			)
			updateProjectOptions.SetName("acme-microservice")
			updateProjectOptions.SetDescription("A microservice to deploy on top of ACME infrastructure.")

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
		It(`ListConfigs request example`, func() {
			fmt.Println("\nListConfigs() result:")
			// begin-list_configs

			listConfigsOptions := projectService.NewListConfigsOptions(
				"testString",
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
				"testString",
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

			projectConfigInputVariableModel := &projectv1.ProjectConfigInputVariable{
				Name: core.StringPtr("account_id"),
				Value: core.StringPtr(`$configs[].name["account-stage"].input.account_id`),
			}

			projectConfigPatchRequestModel := &projectv1.ProjectConfigPatchRequestProjectConfigPatchSchematicsTemplate{
				Input: []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel},
			}

			updateConfigOptions := projectService.NewUpdateConfigOptions(
				"testString",
				"testString",
				projectConfigPatchRequestModel,
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
		It(`GetConfigDiff request example`, func() {
			fmt.Println("\nGetConfigDiff() result:")
			// begin-get_config_diff

			getConfigDiffOptions := projectService.NewGetConfigDiffOptions(
				"testString",
				"testString",
			)

			projectConfigDiff, response, err := projectService.GetConfigDiff(getConfigDiffOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfigDiff, "", "  ")
			fmt.Println(string(b))

			// end-get_config_diff

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectConfigDiff).ToNot(BeNil())
		})
		It(`ForceApprove request example`, func() {
			fmt.Println("\nForceApprove() result:")
			// begin-force_approve

			forceApproveOptions := projectService.NewForceApproveOptions(
				"testString",
				"testString",
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
				"testString",
				"testString",
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
		It(`CheckConfig request example`, func() {
			fmt.Println("\nCheckConfig() result:")
			// begin-check_config

			checkConfigOptions := projectService.NewCheckConfigOptions(
				"testString",
				"testString",
			)

			projectConfig, response, err := projectService.CheckConfig(checkConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-check_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`InstallConfig request example`, func() {
			fmt.Println("\nInstallConfig() result:")
			// begin-install_config

			installConfigOptions := projectService.NewInstallConfigOptions(
				"testString",
				"testString",
			)

			projectConfig, response, err := projectService.InstallConfig(installConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectConfig, "", "  ")
			fmt.Println(string(b))

			// end-install_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(projectConfig).ToNot(BeNil())
		})
		It(`UninstallConfig request example`, func() {
			// begin-uninstall_config

			uninstallConfigOptions := projectService.NewUninstallConfigOptions(
				"testString",
				"testString",
			)

			response, err := projectService.UninstallConfig(uninstallConfigOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from UninstallConfig(): %d\n", response.StatusCode)
			}

			// end-uninstall_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`GetSchematicsJob request example`, func() {
			fmt.Println("\nGetSchematicsJob() result:")
			// begin-get_schematics_job

			getSchematicsJobOptions := projectService.NewGetSchematicsJobOptions(
				"testString",
				"testString",
				"plan",
			)

			actionJob, response, err := projectService.GetSchematicsJob(getSchematicsJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(actionJob, "", "  ")
			fmt.Println(string(b))

			// end-get_schematics_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionJob).ToNot(BeNil())
		})
		It(`GetCostEstimate request example`, func() {
			fmt.Println("\nGetCostEstimate() result:")
			// begin-get_cost_estimate

			getCostEstimateOptions := projectService.NewGetCostEstimateOptions(
				"testString",
				"testString",
			)

			costEstimate, response, err := projectService.GetCostEstimate(getCostEstimateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(costEstimate, "", "  ")
			fmt.Println(string(b))

			// end-get_cost_estimate

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(costEstimate).ToNot(BeNil())
		})
		It(`PostCrnToken request example`, func() {
			fmt.Println("\nPostCrnToken() result:")
			// begin-post_crn_token

			postCrnTokenOptions := projectService.NewPostCrnTokenOptions(
				"testString",
			)

			projectCrnTokenResponse, response, err := projectService.PostCrnToken(postCrnTokenOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectCrnTokenResponse, "", "  ")
			fmt.Println(string(b))

			// end-post_crn_token

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectCrnTokenResponse).ToNot(BeNil())
		})
		It(`PostNotification request example`, func() {
			fmt.Println("\nPostNotification() result:")
			// begin-post_notification

			notificationEventModel := &projectv1.NotificationEvent{
				Event: core.StringPtr("project.create.failed"),
				Target: core.StringPtr("234234324-3444-4556-224232432"),
				Source: core.StringPtr("id.of.project.service.instance"),
				TriggeredBy: core.StringPtr("user-iam-id"),
				ActionURL: core.StringPtr("actionable/url"),
			}

			postNotificationOptions := projectService.NewPostNotificationOptions(
				"testString",
			)
			postNotificationOptions.SetNotifications([]projectv1.NotificationEvent{*notificationEventModel})

			notificationsPrototypePostResponse, response, err := projectService.PostNotification(postNotificationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(notificationsPrototypePostResponse, "", "  ")
			fmt.Println(string(b))

			// end-post_notification

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsPrototypePostResponse).ToNot(BeNil())
		})
		It(`GetNotifications request example`, func() {
			fmt.Println("\nGetNotifications() result:")
			// begin-get_notifications

			getNotificationsOptions := projectService.NewGetNotificationsOptions(
				"testString",
			)

			notificationsGetResponse, response, err := projectService.GetNotifications(getNotificationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(notificationsGetResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_notifications

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsGetResponse).ToNot(BeNil())
		})
		It(`PostEventNotificationsIntegration request example`, func() {
			fmt.Println("\nPostEventNotificationsIntegration() result:")
			// begin-post_event_notifications_integration

			postEventNotificationsIntegrationOptions := projectService.NewPostEventNotificationsIntegrationOptions(
				"testString",
				"CRN of event notifications instance",
			)
			postEventNotificationsIntegrationOptions.SetDescription("A sample project source.")
			postEventNotificationsIntegrationOptions.SetEventNotificationsSourceName("project 1 source name for event notifications")
			postEventNotificationsIntegrationOptions.SetEnabled(true)

			notificationsIntegrationPostResponse, response, err := projectService.PostEventNotificationsIntegration(postEventNotificationsIntegrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(notificationsIntegrationPostResponse, "", "  ")
			fmt.Println(string(b))

			// end-post_event_notifications_integration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsIntegrationPostResponse).ToNot(BeNil())
		})
		It(`GetEventNotificationsIntegration request example`, func() {
			fmt.Println("\nGetEventNotificationsIntegration() result:")
			// begin-get_event_notifications_integration

			getEventNotificationsIntegrationOptions := projectService.NewGetEventNotificationsIntegrationOptions(
				"testString",
			)

			notificationsIntegrationGetResponse, response, err := projectService.GetEventNotificationsIntegration(getEventNotificationsIntegrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(notificationsIntegrationGetResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_event_notifications_integration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsIntegrationGetResponse).ToNot(BeNil())
		})
		It(`PostTestEventNotification request example`, func() {
			fmt.Println("\nPostTestEventNotification() result:")
			// begin-post_test_event_notification

			postTestEventNotificationOptions := projectService.NewPostTestEventNotificationOptions(
				"testString",
			)
			postTestEventNotificationOptions.SetIbmendefaultlong("long test notification message")
			postTestEventNotificationOptions.SetIbmendefaultshort("Test notification")

			notificationsIntegrationTestPostResponse, response, err := projectService.PostTestEventNotification(postTestEventNotificationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(notificationsIntegrationTestPostResponse, "", "  ")
			fmt.Println(string(b))

			// end-post_test_event_notification

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(notificationsIntegrationTestPostResponse).ToNot(BeNil())
		})
		It(`DeleteProject request example`, func() {
			// begin-delete_project

			deleteProjectOptions := projectService.NewDeleteProjectOptions(
				"testString",
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
		It(`DeleteConfig request example`, func() {
			fmt.Println("\nDeleteConfig() result:")
			// begin-delete_config

			deleteConfigOptions := projectService.NewDeleteConfigOptions(
				"testString",
				"testString",
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
		It(`DeleteEventNotificationsIntegration request example`, func() {
			// begin-delete_event_notifications_integration

			deleteEventNotificationsIntegrationOptions := projectService.NewDeleteEventNotificationsIntegrationOptions(
				"testString",
			)

			response, err := projectService.DeleteEventNotificationsIntegration(deleteEventNotificationsIntegrationOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteEventNotificationsIntegration(): %d\n", response.StatusCode)
			}

			// end-delete_event_notifications_integration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
