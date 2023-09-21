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
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/project-go-sdk/projectv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ProjectV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(projectService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(projectService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
				URL: "https://projectv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(projectService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PROJECT_URL": "https://projectv1/api",
				"PROJECT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				projectService, serviceErr := projectv1.NewProjectV1UsingExternalConfig(&projectv1.ProjectV1Options{
				})
				Expect(projectService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := projectService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != projectService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(projectService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(projectService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				projectService, serviceErr := projectv1.NewProjectV1UsingExternalConfig(&projectv1.ProjectV1Options{
					URL: "https://testService/api",
				})
				Expect(projectService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(projectService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := projectService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != projectService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(projectService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(projectService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				projectService, serviceErr := projectv1.NewProjectV1UsingExternalConfig(&projectv1.ProjectV1Options{
				})
				err := projectService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(projectService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := projectService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != projectService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(projectService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(projectService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PROJECT_URL": "https://projectv1/api",
				"PROJECT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			projectService, serviceErr := projectv1.NewProjectV1UsingExternalConfig(&projectv1.ProjectV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(projectService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PROJECT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			projectService, serviceErr := projectv1.NewProjectV1UsingExternalConfig(&projectv1.ProjectV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(projectService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = projectv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateProject(createProjectOptions *CreateProjectOptions) - Operation response error`, func() {
		createProjectPath := "/v1/projects"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["resource_group"]).To(Equal([]string{"Default"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"us-south"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProject with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectPrototypeDefinition model
				projectPrototypeDefinitionModel := new(projectv1.ProjectPrototypeDefinition)
				projectPrototypeDefinitionModel.Name = core.StringPtr("acme-microservice")
				projectPrototypeDefinitionModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				projectPrototypeDefinitionModel.DestroyOnDelete = core.BoolPtr(true)

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigPrototypeDefinitionBlock model
				projectConfigPrototypeDefinitionBlockModel := new(projectv1.ProjectConfigPrototypeDefinitionBlock)
				projectConfigPrototypeDefinitionBlockModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeDefinitionBlockModel.Description = core.StringPtr("testString")
				projectConfigPrototypeDefinitionBlockModel.Labels = []string{"testString"}
				projectConfigPrototypeDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypeDefinitionBlockModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypeDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.Definition = projectConfigPrototypeDefinitionBlockModel

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Definition = projectPrototypeDefinitionModel
				createProjectOptionsModel.Configs = []projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel}
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProject(createProjectOptions *CreateProjectOptions)`, func() {
		createProjectPath := "/v1/projects"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["resource_group"]).To(Equal([]string{"Default"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"us-south"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_error": false, "id": "ID", "location": "Location", "resource_group": "ResourceGroup", "state": "ready", "event_notifications_crn": "EventNotificationsCrn", "configs": [{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "href": "Href", "definition": {"name": "Name", "description": "Description"}}], "definition": {"name": "Name", "description": "Description", "destroy_on_delete": true}}`)
				}))
			})
			It(`Invoke CreateProject successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ProjectPrototypeDefinition model
				projectPrototypeDefinitionModel := new(projectv1.ProjectPrototypeDefinition)
				projectPrototypeDefinitionModel.Name = core.StringPtr("acme-microservice")
				projectPrototypeDefinitionModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				projectPrototypeDefinitionModel.DestroyOnDelete = core.BoolPtr(true)

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigPrototypeDefinitionBlock model
				projectConfigPrototypeDefinitionBlockModel := new(projectv1.ProjectConfigPrototypeDefinitionBlock)
				projectConfigPrototypeDefinitionBlockModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeDefinitionBlockModel.Description = core.StringPtr("testString")
				projectConfigPrototypeDefinitionBlockModel.Labels = []string{"testString"}
				projectConfigPrototypeDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypeDefinitionBlockModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypeDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.Definition = projectConfigPrototypeDefinitionBlockModel

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Definition = projectPrototypeDefinitionModel
				createProjectOptionsModel.Configs = []projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel}
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.CreateProjectWithContext(ctx, createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.CreateProjectWithContext(ctx, createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["resource_group"]).To(Equal([]string{"Default"}))
					Expect(req.URL.Query()["location"]).To(Equal([]string{"us-south"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_error": false, "id": "ID", "location": "Location", "resource_group": "ResourceGroup", "state": "ready", "event_notifications_crn": "EventNotificationsCrn", "configs": [{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "href": "Href", "definition": {"name": "Name", "description": "Description"}}], "definition": {"name": "Name", "description": "Description", "destroy_on_delete": true}}`)
				}))
			})
			It(`Invoke CreateProject successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.CreateProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProjectPrototypeDefinition model
				projectPrototypeDefinitionModel := new(projectv1.ProjectPrototypeDefinition)
				projectPrototypeDefinitionModel.Name = core.StringPtr("acme-microservice")
				projectPrototypeDefinitionModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				projectPrototypeDefinitionModel.DestroyOnDelete = core.BoolPtr(true)

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigPrototypeDefinitionBlock model
				projectConfigPrototypeDefinitionBlockModel := new(projectv1.ProjectConfigPrototypeDefinitionBlock)
				projectConfigPrototypeDefinitionBlockModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeDefinitionBlockModel.Description = core.StringPtr("testString")
				projectConfigPrototypeDefinitionBlockModel.Labels = []string{"testString"}
				projectConfigPrototypeDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypeDefinitionBlockModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypeDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.Definition = projectConfigPrototypeDefinitionBlockModel

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Definition = projectPrototypeDefinitionModel
				createProjectOptionsModel.Configs = []projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel}
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProject with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectPrototypeDefinition model
				projectPrototypeDefinitionModel := new(projectv1.ProjectPrototypeDefinition)
				projectPrototypeDefinitionModel.Name = core.StringPtr("acme-microservice")
				projectPrototypeDefinitionModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				projectPrototypeDefinitionModel.DestroyOnDelete = core.BoolPtr(true)

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigPrototypeDefinitionBlock model
				projectConfigPrototypeDefinitionBlockModel := new(projectv1.ProjectConfigPrototypeDefinitionBlock)
				projectConfigPrototypeDefinitionBlockModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeDefinitionBlockModel.Description = core.StringPtr("testString")
				projectConfigPrototypeDefinitionBlockModel.Labels = []string{"testString"}
				projectConfigPrototypeDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypeDefinitionBlockModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypeDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.Definition = projectConfigPrototypeDefinitionBlockModel

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Definition = projectPrototypeDefinitionModel
				createProjectOptionsModel.Configs = []projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel}
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProjectOptions model with no property values
				createProjectOptionsModelNew := new(projectv1.CreateProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.CreateProject(createProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateProject successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectPrototypeDefinition model
				projectPrototypeDefinitionModel := new(projectv1.ProjectPrototypeDefinition)
				projectPrototypeDefinitionModel.Name = core.StringPtr("acme-microservice")
				projectPrototypeDefinitionModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				projectPrototypeDefinitionModel.DestroyOnDelete = core.BoolPtr(true)

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigPrototypeDefinitionBlock model
				projectConfigPrototypeDefinitionBlockModel := new(projectv1.ProjectConfigPrototypeDefinitionBlock)
				projectConfigPrototypeDefinitionBlockModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeDefinitionBlockModel.Description = core.StringPtr("testString")
				projectConfigPrototypeDefinitionBlockModel.Labels = []string{"testString"}
				projectConfigPrototypeDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypeDefinitionBlockModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypeDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.Definition = projectConfigPrototypeDefinitionBlockModel

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Definition = projectPrototypeDefinitionModel
				createProjectOptionsModel.Configs = []projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel}
				createProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.CreateProject(createProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProjects(listProjectsOptions *ListProjectsOptions) - Operation response error`, func() {
		listProjectsPath := "/v1/projects"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProjects with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProjects(listProjectsOptions *ListProjectsOptions)`, func() {
		listProjectsPath := "/v1/projects"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 10, "total_count": 0, "first": {"href": "Href"}, "last": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "projects": [{"crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_error": false, "id": "ID", "location": "Location", "resource_group": "ResourceGroup", "state": "ready", "event_notifications_crn": "EventNotificationsCrn", "definition": {"name": "Name", "description": "Description", "destroy_on_delete": true}}]}`)
				}))
			})
			It(`Invoke ListProjects successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.ListProjectsWithContext(ctx, listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ListProjectsWithContext(ctx, listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 10, "total_count": 0, "first": {"href": "Href"}, "last": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "projects": [{"crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_error": false, "id": "ID", "location": "Location", "resource_group": "ResourceGroup", "state": "ready", "event_notifications_crn": "EventNotificationsCrn", "definition": {"name": "Name", "description": "Description", "destroy_on_delete": true}}]}`)
				}))
			})
			It(`Invoke ListProjects successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.ListProjects(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProjects with error: Operation request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListProjects successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := new(projectv1.ListProjectsOptions)
				listProjectsOptionsModel.Start = core.StringPtr("testString")
				listProjectsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listProjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.ListProjects(listProjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(projectv1.ProjectCollection)
				nextObject := new(projectv1.PaginationLink)
				nextObject.Href = core.StringPtr("ibm.com?start=abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(projectv1.ProjectCollection)
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextStart without any query params in the "Next" URL`, func() {
				responseObject := new(projectv1.ProjectCollection)
				nextObject := new(projectv1.PaginationLink)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?start=1"},"projects":[{"crn":"crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::","created_at":"2019-01-01T12:00:00.000Z","cumulative_needs_attention_view":[{"event":"Event","event_id":"EventID","config_id":"ConfigID","config_version":13}],"cumulative_needs_attention_view_error":false,"id":"ID","location":"Location","resource_group":"ResourceGroup","state":"ready","event_notifications_crn":"EventNotificationsCrn","definition":{"name":"Name","description":"Description","destroy_on_delete":true}}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"projects":[{"crn":"crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::","created_at":"2019-01-01T12:00:00.000Z","cumulative_needs_attention_view":[{"event":"Event","event_id":"EventID","config_id":"ConfigID","config_version":13}],"cumulative_needs_attention_view_error":false,"id":"ID","location":"Location","resource_group":"ResourceGroup","state":"ready","event_notifications_crn":"EventNotificationsCrn","definition":{"name":"Name","description":"Description","destroy_on_delete":true}}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use ProjectsPager.GetNext successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				listProjectsOptionsModel := &projectv1.ListProjectsOptions{
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := projectService.NewProjectsPager(listProjectsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []projectv1.ProjectCollectionMemberWithMetadata
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use ProjectsPager.GetAll successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				listProjectsOptionsModel := &projectv1.ListProjectsOptions{
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := projectService.NewProjectsPager(listProjectsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetProject(getProjectOptions *GetProjectOptions) - Operation response error`, func() {
		getProjectPath := "/v1/projects/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProject with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetProject(getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetProject(getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProject(getProjectOptions *GetProjectOptions)`, func() {
		getProjectPath := "/v1/projects/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_error": false, "id": "ID", "location": "Location", "resource_group": "ResourceGroup", "state": "ready", "event_notifications_crn": "EventNotificationsCrn", "configs": [{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "href": "Href", "definition": {"name": "Name", "description": "Description"}}], "definition": {"name": "Name", "description": "Description", "destroy_on_delete": true}}`)
				}))
			})
			It(`Invoke GetProject successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetProjectWithContext(ctx, getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetProject(getProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetProjectWithContext(ctx, getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_error": false, "id": "ID", "location": "Location", "resource_group": "ResourceGroup", "state": "ready", "event_notifications_crn": "EventNotificationsCrn", "configs": [{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "href": "Href", "definition": {"name": "Name", "description": "Description"}}], "definition": {"name": "Name", "description": "Description", "destroy_on_delete": true}}`)
				}))
			})
			It(`Invoke GetProject successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetProject(getProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProject with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetProject(getProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProjectOptions model with no property values
				getProjectOptionsModelNew := new(projectv1.GetProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.GetProject(getProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetProject successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetProjectOptions model
				getProjectOptionsModel := new(projectv1.GetProjectOptions)
				getProjectOptionsModel.ID = core.StringPtr("testString")
				getProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetProject(getProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProject(updateProjectOptions *UpdateProjectOptions) - Operation response error`, func() {
		updateProjectPath := "/v1/projects/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProject with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectPrototypePatchDefinitionBlock model
				projectPrototypePatchDefinitionBlockModel := new(projectv1.ProjectPrototypePatchDefinitionBlock)
				projectPrototypePatchDefinitionBlockModel.Name = core.StringPtr("acme-microservice")
				projectPrototypePatchDefinitionBlockModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				projectPrototypePatchDefinitionBlockModel.DestroyOnDelete = core.BoolPtr(true)

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Definition = projectPrototypePatchDefinitionBlockModel
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProject(updateProjectOptions *UpdateProjectOptions)`, func() {
		updateProjectPath := "/v1/projects/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_error": false, "id": "ID", "location": "Location", "resource_group": "ResourceGroup", "state": "ready", "event_notifications_crn": "EventNotificationsCrn", "configs": [{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "href": "Href", "definition": {"name": "Name", "description": "Description"}}], "definition": {"name": "Name", "description": "Description", "destroy_on_delete": true}}`)
				}))
			})
			It(`Invoke UpdateProject successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ProjectPrototypePatchDefinitionBlock model
				projectPrototypePatchDefinitionBlockModel := new(projectv1.ProjectPrototypePatchDefinitionBlock)
				projectPrototypePatchDefinitionBlockModel.Name = core.StringPtr("acme-microservice")
				projectPrototypePatchDefinitionBlockModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				projectPrototypePatchDefinitionBlockModel.DestroyOnDelete = core.BoolPtr(true)

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Definition = projectPrototypePatchDefinitionBlockModel
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.UpdateProjectWithContext(ctx, updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.UpdateProjectWithContext(ctx, updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_error": false, "id": "ID", "location": "Location", "resource_group": "ResourceGroup", "state": "ready", "event_notifications_crn": "EventNotificationsCrn", "configs": [{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "href": "Href", "definition": {"name": "Name", "description": "Description"}}], "definition": {"name": "Name", "description": "Description", "destroy_on_delete": true}}`)
				}))
			})
			It(`Invoke UpdateProject successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.UpdateProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProjectPrototypePatchDefinitionBlock model
				projectPrototypePatchDefinitionBlockModel := new(projectv1.ProjectPrototypePatchDefinitionBlock)
				projectPrototypePatchDefinitionBlockModel.Name = core.StringPtr("acme-microservice")
				projectPrototypePatchDefinitionBlockModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				projectPrototypePatchDefinitionBlockModel.DestroyOnDelete = core.BoolPtr(true)

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Definition = projectPrototypePatchDefinitionBlockModel
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProject with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectPrototypePatchDefinitionBlock model
				projectPrototypePatchDefinitionBlockModel := new(projectv1.ProjectPrototypePatchDefinitionBlock)
				projectPrototypePatchDefinitionBlockModel.Name = core.StringPtr("acme-microservice")
				projectPrototypePatchDefinitionBlockModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				projectPrototypePatchDefinitionBlockModel.DestroyOnDelete = core.BoolPtr(true)

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Definition = projectPrototypePatchDefinitionBlockModel
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProjectOptions model with no property values
				updateProjectOptionsModelNew := new(projectv1.UpdateProjectOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.UpdateProject(updateProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateProject successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectPrototypePatchDefinitionBlock model
				projectPrototypePatchDefinitionBlockModel := new(projectv1.ProjectPrototypePatchDefinitionBlock)
				projectPrototypePatchDefinitionBlockModel.Name = core.StringPtr("acme-microservice")
				projectPrototypePatchDefinitionBlockModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				projectPrototypePatchDefinitionBlockModel.DestroyOnDelete = core.BoolPtr(true)

				// Construct an instance of the UpdateProjectOptions model
				updateProjectOptionsModel := new(projectv1.UpdateProjectOptions)
				updateProjectOptionsModel.ID = core.StringPtr("testString")
				updateProjectOptionsModel.Definition = projectPrototypePatchDefinitionBlockModel
				updateProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.UpdateProject(updateProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteProject(deleteProjectOptions *DeleteProjectOptions)`, func() {
		deleteProjectPath := "/v1/projects/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProjectPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteProject successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectService.DeleteProject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteProjectOptions model
				deleteProjectOptionsModel := new(projectv1.DeleteProjectOptions)
				deleteProjectOptionsModel.ID = core.StringPtr("testString")
				deleteProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectService.DeleteProject(deleteProjectOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteProject with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteProjectOptions model
				deleteProjectOptionsModel := new(projectv1.DeleteProjectOptions)
				deleteProjectOptionsModel.ID = core.StringPtr("testString")
				deleteProjectOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectService.DeleteProject(deleteProjectOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteProjectOptions model with no property values
				deleteProjectOptionsModelNew := new(projectv1.DeleteProjectOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectService.DeleteProject(deleteProjectOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProjectEnvironment(createProjectEnvironmentOptions *CreateProjectEnvironmentOptions) - Operation response error`, func() {
		createProjectEnvironmentPath := "/v1/projects/testString/environments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectEnvironmentPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProjectEnvironment with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("API_KEY")
				projectConfigAuthModel.ApiKey = core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("region", core.StringPtr("us-south"))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the CreateProjectEnvironmentOptions model
				createProjectEnvironmentOptionsModel := new(projectv1.CreateProjectEnvironmentOptions)
				createProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				createProjectEnvironmentOptionsModel.Name = core.StringPtr("development")
				createProjectEnvironmentOptionsModel.Description = core.StringPtr("The environment 'development'")
				createProjectEnvironmentOptionsModel.Authorizations = projectConfigAuthModel
				createProjectEnvironmentOptionsModel.Inputs = inputVariableModel
				createProjectEnvironmentOptionsModel.ComplianceProfile = projectComplianceProfileModel
				createProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.CreateProjectEnvironment(createProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.CreateProjectEnvironment(createProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProjectEnvironment(createProjectEnvironmentOptions *CreateProjectEnvironmentOptions)`, func() {
		createProjectEnvironmentPath := "/v1/projects/testString/environments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectEnvironmentPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "inputs": {}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "metadata": {"project_id": "ProjectID", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke CreateProjectEnvironment successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("API_KEY")
				projectConfigAuthModel.ApiKey = core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("region", core.StringPtr("us-south"))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the CreateProjectEnvironmentOptions model
				createProjectEnvironmentOptionsModel := new(projectv1.CreateProjectEnvironmentOptions)
				createProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				createProjectEnvironmentOptionsModel.Name = core.StringPtr("development")
				createProjectEnvironmentOptionsModel.Description = core.StringPtr("The environment 'development'")
				createProjectEnvironmentOptionsModel.Authorizations = projectConfigAuthModel
				createProjectEnvironmentOptionsModel.Inputs = inputVariableModel
				createProjectEnvironmentOptionsModel.ComplianceProfile = projectComplianceProfileModel
				createProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.CreateProjectEnvironmentWithContext(ctx, createProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.CreateProjectEnvironment(createProjectEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.CreateProjectEnvironmentWithContext(ctx, createProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProjectEnvironmentPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "inputs": {}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "metadata": {"project_id": "ProjectID", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke CreateProjectEnvironment successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.CreateProjectEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("API_KEY")
				projectConfigAuthModel.ApiKey = core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("region", core.StringPtr("us-south"))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the CreateProjectEnvironmentOptions model
				createProjectEnvironmentOptionsModel := new(projectv1.CreateProjectEnvironmentOptions)
				createProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				createProjectEnvironmentOptionsModel.Name = core.StringPtr("development")
				createProjectEnvironmentOptionsModel.Description = core.StringPtr("The environment 'development'")
				createProjectEnvironmentOptionsModel.Authorizations = projectConfigAuthModel
				createProjectEnvironmentOptionsModel.Inputs = inputVariableModel
				createProjectEnvironmentOptionsModel.ComplianceProfile = projectComplianceProfileModel
				createProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.CreateProjectEnvironment(createProjectEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProjectEnvironment with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("API_KEY")
				projectConfigAuthModel.ApiKey = core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("region", core.StringPtr("us-south"))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the CreateProjectEnvironmentOptions model
				createProjectEnvironmentOptionsModel := new(projectv1.CreateProjectEnvironmentOptions)
				createProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				createProjectEnvironmentOptionsModel.Name = core.StringPtr("development")
				createProjectEnvironmentOptionsModel.Description = core.StringPtr("The environment 'development'")
				createProjectEnvironmentOptionsModel.Authorizations = projectConfigAuthModel
				createProjectEnvironmentOptionsModel.Inputs = inputVariableModel
				createProjectEnvironmentOptionsModel.ComplianceProfile = projectComplianceProfileModel
				createProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.CreateProjectEnvironment(createProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProjectEnvironmentOptions model with no property values
				createProjectEnvironmentOptionsModelNew := new(projectv1.CreateProjectEnvironmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.CreateProjectEnvironment(createProjectEnvironmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateProjectEnvironment successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("API_KEY")
				projectConfigAuthModel.ApiKey = core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("region", core.StringPtr("us-south"))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the CreateProjectEnvironmentOptions model
				createProjectEnvironmentOptionsModel := new(projectv1.CreateProjectEnvironmentOptions)
				createProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				createProjectEnvironmentOptionsModel.Name = core.StringPtr("development")
				createProjectEnvironmentOptionsModel.Description = core.StringPtr("The environment 'development'")
				createProjectEnvironmentOptionsModel.Authorizations = projectConfigAuthModel
				createProjectEnvironmentOptionsModel.Inputs = inputVariableModel
				createProjectEnvironmentOptionsModel.ComplianceProfile = projectComplianceProfileModel
				createProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.CreateProjectEnvironment(createProjectEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProjectEnvironments(listProjectEnvironmentsOptions *ListProjectEnvironmentsOptions) - Operation response error`, func() {
		listProjectEnvironmentsPath := "/v1/projects/testString/environments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProjectEnvironments with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListProjectEnvironmentsOptions model
				listProjectEnvironmentsOptionsModel := new(projectv1.ListProjectEnvironmentsOptions)
				listProjectEnvironmentsOptionsModel.ProjectID = core.StringPtr("testString")
				listProjectEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.ListProjectEnvironments(listProjectEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.ListProjectEnvironments(listProjectEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProjectEnvironments(listProjectEnvironmentsOptions *ListProjectEnvironmentsOptions)`, func() {
		listProjectEnvironmentsPath := "/v1/projects/testString/environments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environments": [{"id": "ID", "name": "Name", "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "inputs": {}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "metadata": {"project_id": "ProjectID", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}}]}`)
				}))
			})
			It(`Invoke ListProjectEnvironments successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ListProjectEnvironmentsOptions model
				listProjectEnvironmentsOptionsModel := new(projectv1.ListProjectEnvironmentsOptions)
				listProjectEnvironmentsOptionsModel.ProjectID = core.StringPtr("testString")
				listProjectEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.ListProjectEnvironmentsWithContext(ctx, listProjectEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.ListProjectEnvironments(listProjectEnvironmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ListProjectEnvironmentsWithContext(ctx, listProjectEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProjectEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environments": [{"id": "ID", "name": "Name", "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "inputs": {}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "metadata": {"project_id": "ProjectID", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}}]}`)
				}))
			})
			It(`Invoke ListProjectEnvironments successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.ListProjectEnvironments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProjectEnvironmentsOptions model
				listProjectEnvironmentsOptionsModel := new(projectv1.ListProjectEnvironmentsOptions)
				listProjectEnvironmentsOptionsModel.ProjectID = core.StringPtr("testString")
				listProjectEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.ListProjectEnvironments(listProjectEnvironmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProjectEnvironments with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListProjectEnvironmentsOptions model
				listProjectEnvironmentsOptionsModel := new(projectv1.ListProjectEnvironmentsOptions)
				listProjectEnvironmentsOptionsModel.ProjectID = core.StringPtr("testString")
				listProjectEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.ListProjectEnvironments(listProjectEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListProjectEnvironmentsOptions model with no property values
				listProjectEnvironmentsOptionsModelNew := new(projectv1.ListProjectEnvironmentsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.ListProjectEnvironments(listProjectEnvironmentsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListProjectEnvironments successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListProjectEnvironmentsOptions model
				listProjectEnvironmentsOptionsModel := new(projectv1.ListProjectEnvironmentsOptions)
				listProjectEnvironmentsOptionsModel.ProjectID = core.StringPtr("testString")
				listProjectEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.ListProjectEnvironments(listProjectEnvironmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProjectEnvironment(getProjectEnvironmentOptions *GetProjectEnvironmentOptions) - Operation response error`, func() {
		getProjectEnvironmentPath := "/v1/projects/testString/environments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectEnvironmentPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProjectEnvironment with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetProjectEnvironmentOptions model
				getProjectEnvironmentOptionsModel := new(projectv1.GetProjectEnvironmentOptions)
				getProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				getProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				getProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetProjectEnvironment(getProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetProjectEnvironment(getProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProjectEnvironment(getProjectEnvironmentOptions *GetProjectEnvironmentOptions)`, func() {
		getProjectEnvironmentPath := "/v1/projects/testString/environments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectEnvironmentPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "inputs": {}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "metadata": {"project_id": "ProjectID", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke GetProjectEnvironment successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetProjectEnvironmentOptions model
				getProjectEnvironmentOptionsModel := new(projectv1.GetProjectEnvironmentOptions)
				getProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				getProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				getProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetProjectEnvironmentWithContext(ctx, getProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetProjectEnvironment(getProjectEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetProjectEnvironmentWithContext(ctx, getProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProjectEnvironmentPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "inputs": {}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "metadata": {"project_id": "ProjectID", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke GetProjectEnvironment successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetProjectEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProjectEnvironmentOptions model
				getProjectEnvironmentOptionsModel := new(projectv1.GetProjectEnvironmentOptions)
				getProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				getProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				getProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetProjectEnvironment(getProjectEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProjectEnvironment with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetProjectEnvironmentOptions model
				getProjectEnvironmentOptionsModel := new(projectv1.GetProjectEnvironmentOptions)
				getProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				getProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				getProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetProjectEnvironment(getProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProjectEnvironmentOptions model with no property values
				getProjectEnvironmentOptionsModelNew := new(projectv1.GetProjectEnvironmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.GetProjectEnvironment(getProjectEnvironmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetProjectEnvironment successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetProjectEnvironmentOptions model
				getProjectEnvironmentOptionsModel := new(projectv1.GetProjectEnvironmentOptions)
				getProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				getProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				getProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetProjectEnvironment(getProjectEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProjectEnvironment(updateProjectEnvironmentOptions *UpdateProjectEnvironmentOptions) - Operation response error`, func() {
		updateProjectEnvironmentPath := "/v1/projects/testString/environments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectEnvironmentPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProjectEnvironment with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("API_KEY")
				projectConfigAuthModel.ApiKey = core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("region", core.StringPtr("us-south"))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectEnvironmentOptions model
				updateProjectEnvironmentOptionsModel := new(projectv1.UpdateProjectEnvironmentOptions)
				updateProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				updateProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				updateProjectEnvironmentOptionsModel.Name = core.StringPtr("dev")
				updateProjectEnvironmentOptionsModel.Description = core.StringPtr("The new environment 'dev'")
				updateProjectEnvironmentOptionsModel.Authorizations = projectConfigAuthModel
				updateProjectEnvironmentOptionsModel.Inputs = inputVariableModel
				updateProjectEnvironmentOptionsModel.ComplianceProfile = projectComplianceProfileModel
				updateProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.UpdateProjectEnvironment(updateProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.UpdateProjectEnvironment(updateProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProjectEnvironment(updateProjectEnvironmentOptions *UpdateProjectEnvironmentOptions)`, func() {
		updateProjectEnvironmentPath := "/v1/projects/testString/environments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectEnvironmentPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "inputs": {}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "metadata": {"project_id": "ProjectID", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke UpdateProjectEnvironment successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("API_KEY")
				projectConfigAuthModel.ApiKey = core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("region", core.StringPtr("us-south"))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectEnvironmentOptions model
				updateProjectEnvironmentOptionsModel := new(projectv1.UpdateProjectEnvironmentOptions)
				updateProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				updateProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				updateProjectEnvironmentOptionsModel.Name = core.StringPtr("dev")
				updateProjectEnvironmentOptionsModel.Description = core.StringPtr("The new environment 'dev'")
				updateProjectEnvironmentOptionsModel.Authorizations = projectConfigAuthModel
				updateProjectEnvironmentOptionsModel.Inputs = inputVariableModel
				updateProjectEnvironmentOptionsModel.ComplianceProfile = projectComplianceProfileModel
				updateProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.UpdateProjectEnvironmentWithContext(ctx, updateProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.UpdateProjectEnvironment(updateProjectEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.UpdateProjectEnvironmentWithContext(ctx, updateProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProjectEnvironmentPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "inputs": {}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "metadata": {"project_id": "ProjectID", "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z"}}`)
				}))
			})
			It(`Invoke UpdateProjectEnvironment successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.UpdateProjectEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("API_KEY")
				projectConfigAuthModel.ApiKey = core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("region", core.StringPtr("us-south"))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectEnvironmentOptions model
				updateProjectEnvironmentOptionsModel := new(projectv1.UpdateProjectEnvironmentOptions)
				updateProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				updateProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				updateProjectEnvironmentOptionsModel.Name = core.StringPtr("dev")
				updateProjectEnvironmentOptionsModel.Description = core.StringPtr("The new environment 'dev'")
				updateProjectEnvironmentOptionsModel.Authorizations = projectConfigAuthModel
				updateProjectEnvironmentOptionsModel.Inputs = inputVariableModel
				updateProjectEnvironmentOptionsModel.ComplianceProfile = projectComplianceProfileModel
				updateProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.UpdateProjectEnvironment(updateProjectEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProjectEnvironment with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("API_KEY")
				projectConfigAuthModel.ApiKey = core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("region", core.StringPtr("us-south"))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectEnvironmentOptions model
				updateProjectEnvironmentOptionsModel := new(projectv1.UpdateProjectEnvironmentOptions)
				updateProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				updateProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				updateProjectEnvironmentOptionsModel.Name = core.StringPtr("dev")
				updateProjectEnvironmentOptionsModel.Description = core.StringPtr("The new environment 'dev'")
				updateProjectEnvironmentOptionsModel.Authorizations = projectConfigAuthModel
				updateProjectEnvironmentOptionsModel.Inputs = inputVariableModel
				updateProjectEnvironmentOptionsModel.ComplianceProfile = projectComplianceProfileModel
				updateProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.UpdateProjectEnvironment(updateProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProjectEnvironmentOptions model with no property values
				updateProjectEnvironmentOptionsModelNew := new(projectv1.UpdateProjectEnvironmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.UpdateProjectEnvironment(updateProjectEnvironmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateProjectEnvironment successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("API_KEY")
				projectConfigAuthModel.ApiKey = core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("region", core.StringPtr("us-south"))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the UpdateProjectEnvironmentOptions model
				updateProjectEnvironmentOptionsModel := new(projectv1.UpdateProjectEnvironmentOptions)
				updateProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				updateProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				updateProjectEnvironmentOptionsModel.Name = core.StringPtr("dev")
				updateProjectEnvironmentOptionsModel.Description = core.StringPtr("The new environment 'dev'")
				updateProjectEnvironmentOptionsModel.Authorizations = projectConfigAuthModel
				updateProjectEnvironmentOptionsModel.Inputs = inputVariableModel
				updateProjectEnvironmentOptionsModel.ComplianceProfile = projectComplianceProfileModel
				updateProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.UpdateProjectEnvironment(updateProjectEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteProjectEnvironment(deleteProjectEnvironmentOptions *DeleteProjectEnvironmentOptions) - Operation response error`, func() {
		deleteProjectEnvironmentPath := "/v1/projects/testString/environments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProjectEnvironmentPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteProjectEnvironment with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteProjectEnvironmentOptions model
				deleteProjectEnvironmentOptionsModel := new(projectv1.DeleteProjectEnvironmentOptions)
				deleteProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				deleteProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				deleteProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.DeleteProjectEnvironment(deleteProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.DeleteProjectEnvironment(deleteProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteProjectEnvironment(deleteProjectEnvironmentOptions *DeleteProjectEnvironmentOptions)`, func() {
		deleteProjectEnvironmentPath := "/v1/projects/testString/environments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProjectEnvironmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID"}`)
				}))
			})
			It(`Invoke DeleteProjectEnvironment successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the DeleteProjectEnvironmentOptions model
				deleteProjectEnvironmentOptionsModel := new(projectv1.DeleteProjectEnvironmentOptions)
				deleteProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				deleteProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				deleteProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.DeleteProjectEnvironmentWithContext(ctx, deleteProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.DeleteProjectEnvironment(deleteProjectEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.DeleteProjectEnvironmentWithContext(ctx, deleteProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProjectEnvironmentPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID"}`)
				}))
			})
			It(`Invoke DeleteProjectEnvironment successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.DeleteProjectEnvironment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteProjectEnvironmentOptions model
				deleteProjectEnvironmentOptionsModel := new(projectv1.DeleteProjectEnvironmentOptions)
				deleteProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				deleteProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				deleteProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.DeleteProjectEnvironment(deleteProjectEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteProjectEnvironment with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteProjectEnvironmentOptions model
				deleteProjectEnvironmentOptionsModel := new(projectv1.DeleteProjectEnvironmentOptions)
				deleteProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				deleteProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				deleteProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.DeleteProjectEnvironment(deleteProjectEnvironmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteProjectEnvironmentOptions model with no property values
				deleteProjectEnvironmentOptionsModelNew := new(projectv1.DeleteProjectEnvironmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.DeleteProjectEnvironment(deleteProjectEnvironmentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteProjectEnvironment successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteProjectEnvironmentOptions model
				deleteProjectEnvironmentOptionsModel := new(projectv1.DeleteProjectEnvironmentOptions)
				deleteProjectEnvironmentOptionsModel.ProjectID = core.StringPtr("testString")
				deleteProjectEnvironmentOptionsModel.ID = core.StringPtr("testString")
				deleteProjectEnvironmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.DeleteProjectEnvironment(deleteProjectEnvironmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateConfig(createConfigOptions *CreateConfigOptions) - Operation response error`, func() {
		createConfigPath := "/v1/projects/testString/configs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createConfigPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateConfig with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name["account-stage"].input.account_id`))
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
				inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
				inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("IBMCLOUD_TOOLCHAIN_ENDPOINT", core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com"))

				// Construct an instance of the ProjectConfigPrototypeDefinitionBlock model
				projectConfigPrototypeDefinitionBlockModel := new(projectv1.ProjectConfigPrototypeDefinitionBlock)
				projectConfigPrototypeDefinitionBlockModel.Name = core.StringPtr("env-stage")
				projectConfigPrototypeDefinitionBlockModel.Description = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				projectConfigPrototypeDefinitionBlockModel.Labels = []string{"env:stage", "governance:test", "build:0"}
				projectConfigPrototypeDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypeDefinitionBlockModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypeDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ProjectID = core.StringPtr("testString")
				createConfigOptionsModel.Definition = projectConfigPrototypeDefinitionBlockModel
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateConfig(createConfigOptions *CreateConfigOptions)`, func() {
		createConfigPath := "/v1/projects/testString/configs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createConfigPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke CreateConfig successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name["account-stage"].input.account_id`))
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
				inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
				inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("IBMCLOUD_TOOLCHAIN_ENDPOINT", core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com"))

				// Construct an instance of the ProjectConfigPrototypeDefinitionBlock model
				projectConfigPrototypeDefinitionBlockModel := new(projectv1.ProjectConfigPrototypeDefinitionBlock)
				projectConfigPrototypeDefinitionBlockModel.Name = core.StringPtr("env-stage")
				projectConfigPrototypeDefinitionBlockModel.Description = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				projectConfigPrototypeDefinitionBlockModel.Labels = []string{"env:stage", "governance:test", "build:0"}
				projectConfigPrototypeDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypeDefinitionBlockModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypeDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ProjectID = core.StringPtr("testString")
				createConfigOptionsModel.Definition = projectConfigPrototypeDefinitionBlockModel
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.CreateConfigWithContext(ctx, createConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.CreateConfigWithContext(ctx, createConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createConfigPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke CreateConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.CreateConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name["account-stage"].input.account_id`))
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
				inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
				inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("IBMCLOUD_TOOLCHAIN_ENDPOINT", core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com"))

				// Construct an instance of the ProjectConfigPrototypeDefinitionBlock model
				projectConfigPrototypeDefinitionBlockModel := new(projectv1.ProjectConfigPrototypeDefinitionBlock)
				projectConfigPrototypeDefinitionBlockModel.Name = core.StringPtr("env-stage")
				projectConfigPrototypeDefinitionBlockModel.Description = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				projectConfigPrototypeDefinitionBlockModel.Labels = []string{"env:stage", "governance:test", "build:0"}
				projectConfigPrototypeDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypeDefinitionBlockModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypeDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ProjectID = core.StringPtr("testString")
				createConfigOptionsModel.Definition = projectConfigPrototypeDefinitionBlockModel
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name["account-stage"].input.account_id`))
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
				inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
				inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("IBMCLOUD_TOOLCHAIN_ENDPOINT", core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com"))

				// Construct an instance of the ProjectConfigPrototypeDefinitionBlock model
				projectConfigPrototypeDefinitionBlockModel := new(projectv1.ProjectConfigPrototypeDefinitionBlock)
				projectConfigPrototypeDefinitionBlockModel.Name = core.StringPtr("env-stage")
				projectConfigPrototypeDefinitionBlockModel.Description = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				projectConfigPrototypeDefinitionBlockModel.Labels = []string{"env:stage", "governance:test", "build:0"}
				projectConfigPrototypeDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypeDefinitionBlockModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypeDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ProjectID = core.StringPtr("testString")
				createConfigOptionsModel.Definition = projectConfigPrototypeDefinitionBlockModel
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateConfigOptions model with no property values
				createConfigOptionsModelNew := new(projectv1.CreateConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.CreateConfig(createConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name["account-stage"].input.account_id`))
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
				inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
				inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("IBMCLOUD_TOOLCHAIN_ENDPOINT", core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com"))

				// Construct an instance of the ProjectConfigPrototypeDefinitionBlock model
				projectConfigPrototypeDefinitionBlockModel := new(projectv1.ProjectConfigPrototypeDefinitionBlock)
				projectConfigPrototypeDefinitionBlockModel.Name = core.StringPtr("env-stage")
				projectConfigPrototypeDefinitionBlockModel.Description = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				projectConfigPrototypeDefinitionBlockModel.Labels = []string{"env:stage", "governance:test", "build:0"}
				projectConfigPrototypeDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypeDefinitionBlockModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypeDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ProjectID = core.StringPtr("testString")
				createConfigOptionsModel.Definition = projectConfigPrototypeDefinitionBlockModel
				createConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.CreateConfig(createConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListConfigs(listConfigsOptions *ListConfigsOptions) - Operation response error`, func() {
		listConfigsPath := "/v1/projects/testString/configs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListConfigs with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectv1.ListConfigsOptions)
				listConfigsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListConfigs(listConfigsOptions *ListConfigsOptions)`, func() {
		listConfigsPath := "/v1/projects/testString/configs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configs": [{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "href": "Href", "definition": {"name": "Name", "description": "Description"}}]}`)
				}))
			})
			It(`Invoke ListConfigs successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectv1.ListConfigsOptions)
				listConfigsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.ListConfigsWithContext(ctx, listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ListConfigsWithContext(ctx, listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configs": [{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "href": "Href", "definition": {"name": "Name", "description": "Description"}}]}`)
				}))
			})
			It(`Invoke ListConfigs successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.ListConfigs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectv1.ListConfigsOptions)
				listConfigsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListConfigs with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectv1.ListConfigsOptions)
				listConfigsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListConfigsOptions model with no property values
				listConfigsOptionsModelNew := new(projectv1.ListConfigsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.ListConfigs(listConfigsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListConfigs successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListConfigsOptions model
				listConfigsOptionsModel := new(projectv1.ListConfigsOptions)
				listConfigsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.ListConfigs(listConfigsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConfig(getConfigOptions *GetConfigOptions) - Operation response error`, func() {
		getConfigPath := "/v1/projects/testString/configs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConfig with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectv1.GetConfigOptions)
				getConfigOptionsModel.ProjectID = core.StringPtr("testString")
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConfig(getConfigOptions *GetConfigOptions)`, func() {
		getConfigPath := "/v1/projects/testString/configs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke GetConfig successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectv1.GetConfigOptions)
				getConfigOptionsModel.ProjectID = core.StringPtr("testString")
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetConfigWithContext(ctx, getConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetConfigWithContext(ctx, getConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke GetConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectv1.GetConfigOptions)
				getConfigOptionsModel.ProjectID = core.StringPtr("testString")
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectv1.GetConfigOptions)
				getConfigOptionsModel.ProjectID = core.StringPtr("testString")
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConfigOptions model with no property values
				getConfigOptionsModelNew := new(projectv1.GetConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.GetConfig(getConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetConfigOptions model
				getConfigOptionsModel := new(projectv1.GetConfigOptions)
				getConfigOptionsModel.ProjectID = core.StringPtr("testString")
				getConfigOptionsModel.ID = core.StringPtr("testString")
				getConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetConfig(getConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateConfig(updateConfigOptions *UpdateConfigOptions) - Operation response error`, func() {
		updateConfigPath := "/v1/projects/testString/configs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateConfig with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name["account-stage"].input.account_id`))
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
				inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
				inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigPrototypePatchDefinitionBlock model
				projectConfigPrototypePatchDefinitionBlockModel := new(projectv1.ProjectConfigPrototypePatchDefinitionBlock)
				projectConfigPrototypePatchDefinitionBlockModel.Name = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Description = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Labels = []string{"testString"}
				projectConfigPrototypePatchDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypePatchDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypePatchDefinitionBlockModel.LocatorID = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypePatchDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.Definition = projectConfigPrototypePatchDefinitionBlockModel
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateConfig(updateConfigOptions *UpdateConfigOptions)`, func() {
		updateConfigPath := "/v1/projects/testString/configs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke UpdateConfig successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name["account-stage"].input.account_id`))
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
				inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
				inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigPrototypePatchDefinitionBlock model
				projectConfigPrototypePatchDefinitionBlockModel := new(projectv1.ProjectConfigPrototypePatchDefinitionBlock)
				projectConfigPrototypePatchDefinitionBlockModel.Name = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Description = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Labels = []string{"testString"}
				projectConfigPrototypePatchDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypePatchDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypePatchDefinitionBlockModel.LocatorID = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypePatchDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.Definition = projectConfigPrototypePatchDefinitionBlockModel
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.UpdateConfigWithContext(ctx, updateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.UpdateConfigWithContext(ctx, updateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke UpdateConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.UpdateConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name["account-stage"].input.account_id`))
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
				inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
				inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigPrototypePatchDefinitionBlock model
				projectConfigPrototypePatchDefinitionBlockModel := new(projectv1.ProjectConfigPrototypePatchDefinitionBlock)
				projectConfigPrototypePatchDefinitionBlockModel.Name = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Description = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Labels = []string{"testString"}
				projectConfigPrototypePatchDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypePatchDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypePatchDefinitionBlockModel.LocatorID = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypePatchDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.Definition = projectConfigPrototypePatchDefinitionBlockModel
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name["account-stage"].input.account_id`))
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
				inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
				inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigPrototypePatchDefinitionBlock model
				projectConfigPrototypePatchDefinitionBlockModel := new(projectv1.ProjectConfigPrototypePatchDefinitionBlock)
				projectConfigPrototypePatchDefinitionBlockModel.Name = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Description = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Labels = []string{"testString"}
				projectConfigPrototypePatchDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypePatchDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypePatchDefinitionBlockModel.LocatorID = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypePatchDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.Definition = projectConfigPrototypePatchDefinitionBlockModel
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateConfigOptions model with no property values
				updateConfigOptionsModelNew := new(projectv1.UpdateConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.UpdateConfig(updateConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name["account-stage"].input.account_id`))
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
				inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
				inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				projectConfigSettingModel.SetProperty("foo", core.StringPtr("testString"))

				// Construct an instance of the ProjectConfigPrototypePatchDefinitionBlock model
				projectConfigPrototypePatchDefinitionBlockModel := new(projectv1.ProjectConfigPrototypePatchDefinitionBlock)
				projectConfigPrototypePatchDefinitionBlockModel.Name = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Description = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Labels = []string{"testString"}
				projectConfigPrototypePatchDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypePatchDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypePatchDefinitionBlockModel.LocatorID = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypePatchDefinitionBlockModel.Setting = projectConfigSettingModel

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.Definition = projectConfigPrototypePatchDefinitionBlockModel
				updateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.UpdateConfig(updateConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteConfig(deleteConfigOptions *DeleteConfigOptions) - Operation response error`, func() {
		deleteConfigPath := "/v1/projects/testString/configs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteConfig with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ProjectID = core.StringPtr("testString")
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteConfig(deleteConfigOptions *DeleteConfigOptions)`, func() {
		deleteConfigPath := "/v1/projects/testString/configs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID"}`)
				}))
			})
			It(`Invoke DeleteConfig successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ProjectID = core.StringPtr("testString")
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.DeleteConfigWithContext(ctx, deleteConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.DeleteConfigWithContext(ctx, deleteConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID"}`)
				}))
			})
			It(`Invoke DeleteConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.DeleteConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ProjectID = core.StringPtr("testString")
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ProjectID = core.StringPtr("testString")
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteConfigOptions model with no property values
				deleteConfigOptionsModelNew := new(projectv1.DeleteConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.DeleteConfig(deleteConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigOptions model
				deleteConfigOptionsModel := new(projectv1.DeleteConfigOptions)
				deleteConfigOptionsModel.ProjectID = core.StringPtr("testString")
				deleteConfigOptionsModel.ID = core.StringPtr("testString")
				deleteConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.DeleteConfig(deleteConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ForceApprove(forceApproveOptions *ForceApproveOptions) - Operation response error`, func() {
		forceApprovePath := "/v1/projects/testString/configs/testString/force_approve"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(forceApprovePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ForceApprove with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ForceApproveOptions model
				forceApproveOptionsModel := new(projectv1.ForceApproveOptions)
				forceApproveOptionsModel.ProjectID = core.StringPtr("testString")
				forceApproveOptionsModel.ID = core.StringPtr("testString")
				forceApproveOptionsModel.Comment = core.StringPtr("Approving the changes")
				forceApproveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.ForceApprove(forceApproveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.ForceApprove(forceApproveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ForceApprove(forceApproveOptions *ForceApproveOptions)`, func() {
		forceApprovePath := "/v1/projects/testString/configs/testString/force_approve"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(forceApprovePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke ForceApprove successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ForceApproveOptions model
				forceApproveOptionsModel := new(projectv1.ForceApproveOptions)
				forceApproveOptionsModel.ProjectID = core.StringPtr("testString")
				forceApproveOptionsModel.ID = core.StringPtr("testString")
				forceApproveOptionsModel.Comment = core.StringPtr("Approving the changes")
				forceApproveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.ForceApproveWithContext(ctx, forceApproveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.ForceApprove(forceApproveOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ForceApproveWithContext(ctx, forceApproveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(forceApprovePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke ForceApprove successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.ForceApprove(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ForceApproveOptions model
				forceApproveOptionsModel := new(projectv1.ForceApproveOptions)
				forceApproveOptionsModel.ProjectID = core.StringPtr("testString")
				forceApproveOptionsModel.ID = core.StringPtr("testString")
				forceApproveOptionsModel.Comment = core.StringPtr("Approving the changes")
				forceApproveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.ForceApprove(forceApproveOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ForceApprove with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ForceApproveOptions model
				forceApproveOptionsModel := new(projectv1.ForceApproveOptions)
				forceApproveOptionsModel.ProjectID = core.StringPtr("testString")
				forceApproveOptionsModel.ID = core.StringPtr("testString")
				forceApproveOptionsModel.Comment = core.StringPtr("Approving the changes")
				forceApproveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.ForceApprove(forceApproveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ForceApproveOptions model with no property values
				forceApproveOptionsModelNew := new(projectv1.ForceApproveOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.ForceApprove(forceApproveOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke ForceApprove successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ForceApproveOptions model
				forceApproveOptionsModel := new(projectv1.ForceApproveOptions)
				forceApproveOptionsModel.ProjectID = core.StringPtr("testString")
				forceApproveOptionsModel.ID = core.StringPtr("testString")
				forceApproveOptionsModel.Comment = core.StringPtr("Approving the changes")
				forceApproveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.ForceApprove(forceApproveOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Approve(approveOptions *ApproveOptions) - Operation response error`, func() {
		approvePath := "/v1/projects/testString/configs/testString/approve"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(approvePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Approve with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ApproveOptions model
				approveOptionsModel := new(projectv1.ApproveOptions)
				approveOptionsModel.ProjectID = core.StringPtr("testString")
				approveOptionsModel.ID = core.StringPtr("testString")
				approveOptionsModel.Comment = core.StringPtr("Approving the changes")
				approveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.Approve(approveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.Approve(approveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Approve(approveOptions *ApproveOptions)`, func() {
		approvePath := "/v1/projects/testString/configs/testString/approve"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(approvePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke Approve successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ApproveOptions model
				approveOptionsModel := new(projectv1.ApproveOptions)
				approveOptionsModel.ProjectID = core.StringPtr("testString")
				approveOptionsModel.ID = core.StringPtr("testString")
				approveOptionsModel.Comment = core.StringPtr("Approving the changes")
				approveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.ApproveWithContext(ctx, approveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.Approve(approveOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ApproveWithContext(ctx, approveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(approvePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke Approve successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.Approve(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ApproveOptions model
				approveOptionsModel := new(projectv1.ApproveOptions)
				approveOptionsModel.ProjectID = core.StringPtr("testString")
				approveOptionsModel.ID = core.StringPtr("testString")
				approveOptionsModel.Comment = core.StringPtr("Approving the changes")
				approveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.Approve(approveOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Approve with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ApproveOptions model
				approveOptionsModel := new(projectv1.ApproveOptions)
				approveOptionsModel.ProjectID = core.StringPtr("testString")
				approveOptionsModel.ID = core.StringPtr("testString")
				approveOptionsModel.Comment = core.StringPtr("Approving the changes")
				approveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.Approve(approveOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ApproveOptions model with no property values
				approveOptionsModelNew := new(projectv1.ApproveOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.Approve(approveOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke Approve successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ApproveOptions model
				approveOptionsModel := new(projectv1.ApproveOptions)
				approveOptionsModel.ProjectID = core.StringPtr("testString")
				approveOptionsModel.ID = core.StringPtr("testString")
				approveOptionsModel.Comment = core.StringPtr("Approving the changes")
				approveOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.Approve(approveOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ValidateConfig(validateConfigOptions *ValidateConfigOptions) - Operation response error`, func() {
		validateConfigPath := "/v1/projects/testString/configs/testString/validate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validateConfigPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ValidateConfig with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ValidateConfigOptions model
				validateConfigOptionsModel := new(projectv1.ValidateConfigOptions)
				validateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				validateConfigOptionsModel.ID = core.StringPtr("testString")
				validateConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				validateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.ValidateConfig(validateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.ValidateConfig(validateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ValidateConfig(validateConfigOptions *ValidateConfigOptions)`, func() {
		validateConfigPath := "/v1/projects/testString/configs/testString/validate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validateConfigPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke ValidateConfig successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ValidateConfigOptions model
				validateConfigOptionsModel := new(projectv1.ValidateConfigOptions)
				validateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				validateConfigOptionsModel.ID = core.StringPtr("testString")
				validateConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				validateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.ValidateConfigWithContext(ctx, validateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.ValidateConfig(validateConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ValidateConfigWithContext(ctx, validateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validateConfigPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke ValidateConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.ValidateConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ValidateConfigOptions model
				validateConfigOptionsModel := new(projectv1.ValidateConfigOptions)
				validateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				validateConfigOptionsModel.ID = core.StringPtr("testString")
				validateConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				validateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.ValidateConfig(validateConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ValidateConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ValidateConfigOptions model
				validateConfigOptionsModel := new(projectv1.ValidateConfigOptions)
				validateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				validateConfigOptionsModel.ID = core.StringPtr("testString")
				validateConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				validateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.ValidateConfig(validateConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ValidateConfigOptions model with no property values
				validateConfigOptionsModelNew := new(projectv1.ValidateConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.ValidateConfig(validateConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke ValidateConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ValidateConfigOptions model
				validateConfigOptionsModel := new(projectv1.ValidateConfigOptions)
				validateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				validateConfigOptionsModel.ID = core.StringPtr("testString")
				validateConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				validateConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.ValidateConfig(validateConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeployConfig(deployConfigOptions *DeployConfigOptions) - Operation response error`, func() {
		deployConfigPath := "/v1/projects/testString/configs/testString/deploy"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deployConfigPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeployConfig with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeployConfigOptions model
				deployConfigOptionsModel := new(projectv1.DeployConfigOptions)
				deployConfigOptionsModel.ProjectID = core.StringPtr("testString")
				deployConfigOptionsModel.ID = core.StringPtr("testString")
				deployConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.DeployConfig(deployConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.DeployConfig(deployConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeployConfig(deployConfigOptions *DeployConfigOptions)`, func() {
		deployConfigPath := "/v1/projects/testString/configs/testString/deploy"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deployConfigPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke DeployConfig successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the DeployConfigOptions model
				deployConfigOptionsModel := new(projectv1.DeployConfigOptions)
				deployConfigOptionsModel.ProjectID = core.StringPtr("testString")
				deployConfigOptionsModel.ID = core.StringPtr("testString")
				deployConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.DeployConfigWithContext(ctx, deployConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.DeployConfig(deployConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.DeployConfigWithContext(ctx, deployConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deployConfigPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke DeployConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.DeployConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeployConfigOptions model
				deployConfigOptionsModel := new(projectv1.DeployConfigOptions)
				deployConfigOptionsModel.ProjectID = core.StringPtr("testString")
				deployConfigOptionsModel.ID = core.StringPtr("testString")
				deployConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.DeployConfig(deployConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeployConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeployConfigOptions model
				deployConfigOptionsModel := new(projectv1.DeployConfigOptions)
				deployConfigOptionsModel.ProjectID = core.StringPtr("testString")
				deployConfigOptionsModel.ID = core.StringPtr("testString")
				deployConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.DeployConfig(deployConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeployConfigOptions model with no property values
				deployConfigOptionsModelNew := new(projectv1.DeployConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.DeployConfig(deployConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeployConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeployConfigOptions model
				deployConfigOptionsModel := new(projectv1.DeployConfigOptions)
				deployConfigOptionsModel.ProjectID = core.StringPtr("testString")
				deployConfigOptionsModel.ID = core.StringPtr("testString")
				deployConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.DeployConfig(deployConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UndeployConfig(undeployConfigOptions *UndeployConfigOptions)`, func() {
		undeployConfigPath := "/v1/projects/testString/configs/testString/undeploy"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(undeployConfigPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UndeployConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectService.UndeployConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UndeployConfigOptions model
				undeployConfigOptionsModel := new(projectv1.UndeployConfigOptions)
				undeployConfigOptionsModel.ProjectID = core.StringPtr("testString")
				undeployConfigOptionsModel.ID = core.StringPtr("testString")
				undeployConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectService.UndeployConfig(undeployConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UndeployConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the UndeployConfigOptions model
				undeployConfigOptionsModel := new(projectv1.UndeployConfigOptions)
				undeployConfigOptionsModel.ProjectID = core.StringPtr("testString")
				undeployConfigOptionsModel.ID = core.StringPtr("testString")
				undeployConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectService.UndeployConfig(undeployConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UndeployConfigOptions model with no property values
				undeployConfigOptionsModelNew := new(projectv1.UndeployConfigOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectService.UndeployConfig(undeployConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SyncConfig(syncConfigOptions *SyncConfigOptions)`, func() {
		syncConfigPath := "/v1/projects/testString/configs/testString/sync"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(syncConfigPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke SyncConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectService.SyncConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the SyncConfigOptions model
				syncConfigOptionsModel := new(projectv1.SyncConfigOptions)
				syncConfigOptionsModel.ProjectID = core.StringPtr("testString")
				syncConfigOptionsModel.ID = core.StringPtr("testString")
				syncConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectService.SyncConfig(syncConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke SyncConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the SyncConfigOptions model
				syncConfigOptionsModel := new(projectv1.SyncConfigOptions)
				syncConfigOptionsModel.ProjectID = core.StringPtr("testString")
				syncConfigOptionsModel.ID = core.StringPtr("testString")
				syncConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectService.SyncConfig(syncConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the SyncConfigOptions model with no property values
				syncConfigOptionsModelNew := new(projectv1.SyncConfigOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectService.SyncConfig(syncConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListConfigResources(listConfigResourcesOptions *ListConfigResourcesOptions) - Operation response error`, func() {
		listConfigResourcesPath := "/v1/projects/testString/configs/testString/resources"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigResourcesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListConfigResources with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListConfigResourcesOptions model
				listConfigResourcesOptionsModel := new(projectv1.ListConfigResourcesOptions)
				listConfigResourcesOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigResourcesOptionsModel.ID = core.StringPtr("testString")
				listConfigResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.ListConfigResources(listConfigResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.ListConfigResources(listConfigResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListConfigResources(listConfigResourcesOptions *ListConfigResourcesOptions)`, func() {
		listConfigResourcesPath := "/v1/projects/testString/configs/testString/resources"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigResourcesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resources": [{"resource_crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "resource_name": "ResourceName", "resource_type": "ResourceType", "resource_tainted": false, "resource_group_name": "ResourceGroupName"}], "resources_count": 14}`)
				}))
			})
			It(`Invoke ListConfigResources successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ListConfigResourcesOptions model
				listConfigResourcesOptionsModel := new(projectv1.ListConfigResourcesOptions)
				listConfigResourcesOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigResourcesOptionsModel.ID = core.StringPtr("testString")
				listConfigResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.ListConfigResourcesWithContext(ctx, listConfigResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.ListConfigResources(listConfigResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ListConfigResourcesWithContext(ctx, listConfigResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigResourcesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resources": [{"resource_crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "resource_name": "ResourceName", "resource_type": "ResourceType", "resource_tainted": false, "resource_group_name": "ResourceGroupName"}], "resources_count": 14}`)
				}))
			})
			It(`Invoke ListConfigResources successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.ListConfigResources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListConfigResourcesOptions model
				listConfigResourcesOptionsModel := new(projectv1.ListConfigResourcesOptions)
				listConfigResourcesOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigResourcesOptionsModel.ID = core.StringPtr("testString")
				listConfigResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.ListConfigResources(listConfigResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListConfigResources with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListConfigResourcesOptions model
				listConfigResourcesOptionsModel := new(projectv1.ListConfigResourcesOptions)
				listConfigResourcesOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigResourcesOptionsModel.ID = core.StringPtr("testString")
				listConfigResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.ListConfigResources(listConfigResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListConfigResourcesOptions model with no property values
				listConfigResourcesOptionsModelNew := new(projectv1.ListConfigResourcesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.ListConfigResources(listConfigResourcesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListConfigResources successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListConfigResourcesOptions model
				listConfigResourcesOptionsModel := new(projectv1.ListConfigResourcesOptions)
				listConfigResourcesOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigResourcesOptionsModel.ID = core.StringPtr("testString")
				listConfigResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.ListConfigResources(listConfigResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListConfigVersions(listConfigVersionsOptions *ListConfigVersionsOptions) - Operation response error`, func() {
		listConfigVersionsPath := "/v1/projects/testString/configs/testString/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigVersionsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListConfigVersions with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListConfigVersionsOptions model
				listConfigVersionsOptionsModel := new(projectv1.ListConfigVersionsOptions)
				listConfigVersionsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigVersionsOptionsModel.ID = core.StringPtr("testString")
				listConfigVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.ListConfigVersions(listConfigVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.ListConfigVersions(listConfigVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListConfigVersions(listConfigVersionsOptions *ListConfigVersionsOptions)`, func() {
		listConfigVersionsPath := "/v1/projects/testString/configs/testString/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"versions": [{"needs_attention_state": ["anyValue"], "state": "approved", "version": 7, "href": "Href"}]}`)
				}))
			})
			It(`Invoke ListConfigVersions successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ListConfigVersionsOptions model
				listConfigVersionsOptionsModel := new(projectv1.ListConfigVersionsOptions)
				listConfigVersionsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigVersionsOptionsModel.ID = core.StringPtr("testString")
				listConfigVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.ListConfigVersionsWithContext(ctx, listConfigVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.ListConfigVersions(listConfigVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ListConfigVersionsWithContext(ctx, listConfigVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listConfigVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"versions": [{"needs_attention_state": ["anyValue"], "state": "approved", "version": 7, "href": "Href"}]}`)
				}))
			})
			It(`Invoke ListConfigVersions successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.ListConfigVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListConfigVersionsOptions model
				listConfigVersionsOptionsModel := new(projectv1.ListConfigVersionsOptions)
				listConfigVersionsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigVersionsOptionsModel.ID = core.StringPtr("testString")
				listConfigVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.ListConfigVersions(listConfigVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListConfigVersions with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListConfigVersionsOptions model
				listConfigVersionsOptionsModel := new(projectv1.ListConfigVersionsOptions)
				listConfigVersionsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigVersionsOptionsModel.ID = core.StringPtr("testString")
				listConfigVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.ListConfigVersions(listConfigVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListConfigVersionsOptions model with no property values
				listConfigVersionsOptionsModelNew := new(projectv1.ListConfigVersionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.ListConfigVersions(listConfigVersionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListConfigVersions successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListConfigVersionsOptions model
				listConfigVersionsOptionsModel := new(projectv1.ListConfigVersionsOptions)
				listConfigVersionsOptionsModel.ProjectID = core.StringPtr("testString")
				listConfigVersionsOptionsModel.ID = core.StringPtr("testString")
				listConfigVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.ListConfigVersions(listConfigVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConfigVersion(getConfigVersionOptions *GetConfigVersionOptions) - Operation response error`, func() {
		getConfigVersionPath := "/v1/projects/testString/configs/testString/versions/38"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigVersionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConfigVersion with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetConfigVersionOptions model
				getConfigVersionOptionsModel := new(projectv1.GetConfigVersionOptions)
				getConfigVersionOptionsModel.ProjectID = core.StringPtr("testString")
				getConfigVersionOptionsModel.ID = core.StringPtr("testString")
				getConfigVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				getConfigVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.GetConfigVersion(getConfigVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.GetConfigVersion(getConfigVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConfigVersion(getConfigVersionOptions *GetConfigVersionOptions)`, func() {
		getConfigVersionPath := "/v1/projects/testString/configs/testString/versions/38"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke GetConfigVersion successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the GetConfigVersionOptions model
				getConfigVersionOptionsModel := new(projectv1.GetConfigVersionOptions)
				getConfigVersionOptionsModel.ProjectID = core.StringPtr("testString")
				getConfigVersionOptionsModel.ID = core.StringPtr("testString")
				getConfigVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				getConfigVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.GetConfigVersionWithContext(ctx, getConfigVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.GetConfigVersion(getConfigVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.GetConfigVersionWithContext(ctx, getConfigVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "project_id": "ProjectID", "version": 7, "is_draft": false, "needs_attention_state": ["anyValue"], "state": "approved", "update_available": false, "created_at": "2019-01-01T12:00:00.000Z", "updated_at": "2019-01-01T12:00:00.000Z", "last_approved": {"is_forced": true, "comment": "Comment", "timestamp": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "last_save": "2019-01-01T12:00:00.000Z", "last_validated": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}, "cost_estimate": {"version": "Version", "currency": "Currency", "totalHourlyCost": "TotalHourlyCost", "totalMonthlyCost": "TotalMonthlyCost", "pastTotalHourlyCost": "PastTotalHourlyCost", "pastTotalMonthlyCost": "PastTotalMonthlyCost", "diffTotalHourlyCost": "DiffTotalHourlyCost", "diffTotalMonthlyCost": "DiffTotalMonthlyCost", "timeGenerated": "2019-01-01T12:00:00.000Z", "user_id": "UserID"}, "cra_logs": {"cra_version": "CraVersion", "schema_version": "SchemaVersion", "status": "Status", "summary": {"anyKey": "anyValue"}, "timestamp": "2019-01-01T12:00:00.000Z"}}, "last_deployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "last_undeployed": {"href": "Href", "result": "failed", "job": {"id": "ID", "summary": {"plan_summary": {"anyKey": "anyValue"}, "apply_summary": {"anyKey": "anyValue"}, "destroy_summary": {"anyKey": "anyValue"}, "message_summary": {"anyKey": "anyValue"}, "plan_messages": {"anyKey": "anyValue"}, "apply_messages": {"anyKey": "anyValue"}, "destroy_messages": {"anyKey": "anyValue"}}}}, "definition": {"name": "Name", "description": "Description", "labels": ["Labels"], "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "input": {}, "setting": {}, "type": "terraform_template", "output": [{"name": "Name", "description": "Description", "value": "anyValue"}]}}`)
				}))
			})
			It(`Invoke GetConfigVersion successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.GetConfigVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConfigVersionOptions model
				getConfigVersionOptionsModel := new(projectv1.GetConfigVersionOptions)
				getConfigVersionOptionsModel.ProjectID = core.StringPtr("testString")
				getConfigVersionOptionsModel.ID = core.StringPtr("testString")
				getConfigVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				getConfigVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.GetConfigVersion(getConfigVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConfigVersion with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetConfigVersionOptions model
				getConfigVersionOptionsModel := new(projectv1.GetConfigVersionOptions)
				getConfigVersionOptionsModel.ProjectID = core.StringPtr("testString")
				getConfigVersionOptionsModel.ID = core.StringPtr("testString")
				getConfigVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				getConfigVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.GetConfigVersion(getConfigVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConfigVersionOptions model with no property values
				getConfigVersionOptionsModelNew := new(projectv1.GetConfigVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.GetConfigVersion(getConfigVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetConfigVersion successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the GetConfigVersionOptions model
				getConfigVersionOptionsModel := new(projectv1.GetConfigVersionOptions)
				getConfigVersionOptionsModel.ProjectID = core.StringPtr("testString")
				getConfigVersionOptionsModel.ID = core.StringPtr("testString")
				getConfigVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				getConfigVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.GetConfigVersion(getConfigVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteConfigVersion(deleteConfigVersionOptions *DeleteConfigVersionOptions) - Operation response error`, func() {
		deleteConfigVersionPath := "/v1/projects/testString/configs/testString/versions/38"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigVersionPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteConfigVersion with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigVersionOptions model
				deleteConfigVersionOptionsModel := new(projectv1.DeleteConfigVersionOptions)
				deleteConfigVersionOptionsModel.ProjectID = core.StringPtr("testString")
				deleteConfigVersionOptionsModel.ID = core.StringPtr("testString")
				deleteConfigVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				deleteConfigVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.DeleteConfigVersion(deleteConfigVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.DeleteConfigVersion(deleteConfigVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteConfigVersion(deleteConfigVersionOptions *DeleteConfigVersionOptions)`, func() {
		deleteConfigVersionPath := "/v1/projects/testString/configs/testString/versions/38"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigVersionPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID"}`)
				}))
			})
			It(`Invoke DeleteConfigVersion successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the DeleteConfigVersionOptions model
				deleteConfigVersionOptionsModel := new(projectv1.DeleteConfigVersionOptions)
				deleteConfigVersionOptionsModel.ProjectID = core.StringPtr("testString")
				deleteConfigVersionOptionsModel.ID = core.StringPtr("testString")
				deleteConfigVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				deleteConfigVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.DeleteConfigVersionWithContext(ctx, deleteConfigVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.DeleteConfigVersion(deleteConfigVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.DeleteConfigVersionWithContext(ctx, deleteConfigVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteConfigVersionPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID"}`)
				}))
			})
			It(`Invoke DeleteConfigVersion successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.DeleteConfigVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteConfigVersionOptions model
				deleteConfigVersionOptionsModel := new(projectv1.DeleteConfigVersionOptions)
				deleteConfigVersionOptionsModel.ProjectID = core.StringPtr("testString")
				deleteConfigVersionOptionsModel.ID = core.StringPtr("testString")
				deleteConfigVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				deleteConfigVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.DeleteConfigVersion(deleteConfigVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteConfigVersion with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigVersionOptions model
				deleteConfigVersionOptionsModel := new(projectv1.DeleteConfigVersionOptions)
				deleteConfigVersionOptionsModel.ProjectID = core.StringPtr("testString")
				deleteConfigVersionOptionsModel.ID = core.StringPtr("testString")
				deleteConfigVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				deleteConfigVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.DeleteConfigVersion(deleteConfigVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteConfigVersionOptions model with no property values
				deleteConfigVersionOptionsModelNew := new(projectv1.DeleteConfigVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.DeleteConfigVersion(deleteConfigVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteConfigVersion successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the DeleteConfigVersionOptions model
				deleteConfigVersionOptionsModel := new(projectv1.DeleteConfigVersionOptions)
				deleteConfigVersionOptionsModel.ProjectID = core.StringPtr("testString")
				deleteConfigVersionOptionsModel.ID = core.StringPtr("testString")
				deleteConfigVersionOptionsModel.Version = core.Int64Ptr(int64(38))
				deleteConfigVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.DeleteConfigVersion(deleteConfigVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListEnvironments(listEnvironmentsOptions *ListEnvironmentsOptions) - Operation response error`, func() {
		listEnvironmentsPath := "/v1/environments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))
					// TODO: Add check for name query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListEnvironments with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(projectv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Name = core.BoolPtr(false)
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListEnvironments(listEnvironmentsOptions *ListEnvironmentsOptions)`, func() {
		listEnvironmentsPath := "/v1/environments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for name query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environments": [{"name": "Name", "id": "ID"}]}`)
				}))
			})
			It(`Invoke ListEnvironments successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(projectv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Name = core.BoolPtr(false)
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.ListEnvironmentsWithContext(ctx, listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.ListEnvironmentsWithContext(ctx, listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listEnvironmentsPath))
					Expect(req.Method).To(Equal("GET"))

					// TODO: Add check for name query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"environments": [{"name": "Name", "id": "ID"}]}`)
				}))
			})
			It(`Invoke ListEnvironments successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.ListEnvironments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(projectv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Name = core.BoolPtr(false)
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListEnvironments with error: Operation request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(projectv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Name = core.BoolPtr(false)
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListEnvironments successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := new(projectv1.ListEnvironmentsOptions)
				listEnvironmentsOptionsModel.Name = core.BoolPtr(false)
				listEnvironmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.ListEnvironments(listEnvironmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			projectService, _ := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
				URL:           "http://projectv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewApproveOptions successfully`, func() {
				// Construct an instance of the ApproveOptions model
				projectID := "testString"
				id := "testString"
				approveOptionsModel := projectService.NewApproveOptions(projectID, id)
				approveOptionsModel.SetProjectID("testString")
				approveOptionsModel.SetID("testString")
				approveOptionsModel.SetComment("Approving the changes")
				approveOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(approveOptionsModel).ToNot(BeNil())
				Expect(approveOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(approveOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(approveOptionsModel.Comment).To(Equal(core.StringPtr("Approving the changes")))
				Expect(approveOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateConfigOptions successfully`, func() {
				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				Expect(projectConfigAuthTrustedProfileModel).ToNot(BeNil())
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")
				Expect(projectConfigAuthTrustedProfileModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigAuthTrustedProfileModel.TargetIamID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				Expect(projectConfigAuthModel).ToNot(BeNil())
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")
				Expect(projectConfigAuthModel.TrustedProfile).To(Equal(projectConfigAuthTrustedProfileModel))
				Expect(projectConfigAuthModel.Method).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigAuthModel.ApiKey).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				Expect(projectComplianceProfileModel).ToNot(BeNil())
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")
				Expect(projectComplianceProfileModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.InstanceLocation).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.ProfileName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				Expect(inputVariableModel).ToNot(BeNil())
				inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name["account-stage"].input.account_id`))
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
				inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
				inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))
				Expect(inputVariableModel.GetProperties()).ToNot(BeEmpty())
				Expect(inputVariableModel.GetProperty("account_id")).To(Equal(core.StringPtr(`$configs[].name["account-stage"].input.account_id`)))
				Expect(inputVariableModel.GetProperty("resource_group")).To(Equal(core.StringPtr("stage")))
				Expect(inputVariableModel.GetProperty("access_tags")).To(Equal(core.StringPtr(`["env:stage"]`)))
				Expect(inputVariableModel.GetProperty("logdna_name")).To(Equal(core.StringPtr("Name of the LogDNA stage service instance")))
				Expect(inputVariableModel.GetProperty("sysdig_name")).To(Equal(core.StringPtr("Name of the SysDig stage service instance")))

				inputVariableModel.SetProperties(nil)
				Expect(inputVariableModel.GetProperties()).To(BeEmpty())

				inputVariableModelExpectedMap := make(map[string]interface{})
				inputVariableModelExpectedMap["account_id"] = core.StringPtr(`$configs[].name["account-stage"].input.account_id`)
				inputVariableModelExpectedMap["resource_group"] = core.StringPtr("stage")
				inputVariableModelExpectedMap["access_tags"] = core.StringPtr(`["env:stage"]`)
				inputVariableModelExpectedMap["logdna_name"] = core.StringPtr("Name of the LogDNA stage service instance")
				inputVariableModelExpectedMap["sysdig_name"] = core.StringPtr("Name of the SysDig stage service instance")
				inputVariableModel.SetProperties(inputVariableModelExpectedMap)
				inputVariableModelActualMap := inputVariableModel.GetProperties()
				Expect(inputVariableModelActualMap).To(Equal(inputVariableModelExpectedMap))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				Expect(projectConfigSettingModel).ToNot(BeNil())
				projectConfigSettingModel.SetProperty("IBMCLOUD_TOOLCHAIN_ENDPOINT", core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com"))
				Expect(projectConfigSettingModel.GetProperties()).ToNot(BeEmpty())
				Expect(projectConfigSettingModel.GetProperty("IBMCLOUD_TOOLCHAIN_ENDPOINT")).To(Equal(core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")))

				projectConfigSettingModel.SetProperties(nil)
				Expect(projectConfigSettingModel.GetProperties()).To(BeEmpty())

				projectConfigSettingModelExpectedMap := make(map[string]interface{})
				projectConfigSettingModelExpectedMap["IBMCLOUD_TOOLCHAIN_ENDPOINT"] = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")
				projectConfigSettingModel.SetProperties(projectConfigSettingModelExpectedMap)
				projectConfigSettingModelActualMap := projectConfigSettingModel.GetProperties()
				Expect(projectConfigSettingModelActualMap).To(Equal(projectConfigSettingModelExpectedMap))

				// Construct an instance of the ProjectConfigPrototypeDefinitionBlock model
				projectConfigPrototypeDefinitionBlockModel := new(projectv1.ProjectConfigPrototypeDefinitionBlock)
				Expect(projectConfigPrototypeDefinitionBlockModel).ToNot(BeNil())
				projectConfigPrototypeDefinitionBlockModel.Name = core.StringPtr("env-stage")
				projectConfigPrototypeDefinitionBlockModel.Description = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				projectConfigPrototypeDefinitionBlockModel.Labels = []string{"env:stage", "governance:test", "build:0"}
				projectConfigPrototypeDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypeDefinitionBlockModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypeDefinitionBlockModel.Setting = projectConfigSettingModel
				Expect(projectConfigPrototypeDefinitionBlockModel.Name).To(Equal(core.StringPtr("env-stage")))
				Expect(projectConfigPrototypeDefinitionBlockModel.Description).To(Equal(core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")))
				Expect(projectConfigPrototypeDefinitionBlockModel.Labels).To(Equal([]string{"env:stage", "governance:test", "build:0"}))
				Expect(projectConfigPrototypeDefinitionBlockModel.Authorizations).To(Equal(projectConfigAuthModel))
				Expect(projectConfigPrototypeDefinitionBlockModel.ComplianceProfile).To(Equal(projectComplianceProfileModel))
				Expect(projectConfigPrototypeDefinitionBlockModel.LocatorID).To(Equal(core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")))
				Expect(projectConfigPrototypeDefinitionBlockModel.Input).To(Equal(inputVariableModel))
				Expect(projectConfigPrototypeDefinitionBlockModel.Setting).To(Equal(projectConfigSettingModel))

				// Construct an instance of the CreateConfigOptions model
				projectID := "testString"
				createConfigOptionsModel := projectService.NewCreateConfigOptions(projectID)
				createConfigOptionsModel.SetProjectID("testString")
				createConfigOptionsModel.SetDefinition(projectConfigPrototypeDefinitionBlockModel)
				createConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createConfigOptionsModel).ToNot(BeNil())
				Expect(createConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(createConfigOptionsModel.Definition).To(Equal(projectConfigPrototypeDefinitionBlockModel))
				Expect(createConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProjectEnvironmentOptions successfully`, func() {
				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				Expect(projectConfigAuthTrustedProfileModel).ToNot(BeNil())
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")
				Expect(projectConfigAuthTrustedProfileModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigAuthTrustedProfileModel.TargetIamID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				Expect(projectConfigAuthModel).ToNot(BeNil())
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("API_KEY")
				projectConfigAuthModel.ApiKey = core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9")
				Expect(projectConfigAuthModel.TrustedProfile).To(Equal(projectConfigAuthTrustedProfileModel))
				Expect(projectConfigAuthModel.Method).To(Equal(core.StringPtr("API_KEY")))
				Expect(projectConfigAuthModel.ApiKey).To(Equal(core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9")))

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				Expect(inputVariableModel).ToNot(BeNil())
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("region", core.StringPtr("us-south"))
				Expect(inputVariableModel.GetProperties()).ToNot(BeEmpty())
				Expect(inputVariableModel.GetProperty("resource_group")).To(Equal(core.StringPtr("stage")))
				Expect(inputVariableModel.GetProperty("region")).To(Equal(core.StringPtr("us-south")))

				inputVariableModel.SetProperties(nil)
				Expect(inputVariableModel.GetProperties()).To(BeEmpty())

				inputVariableModelExpectedMap := make(map[string]interface{})
				inputVariableModelExpectedMap["resource_group"] = core.StringPtr("stage")
				inputVariableModelExpectedMap["region"] = core.StringPtr("us-south")
				inputVariableModel.SetProperties(inputVariableModelExpectedMap)
				inputVariableModelActualMap := inputVariableModel.GetProperties()
				Expect(inputVariableModelActualMap).To(Equal(inputVariableModelExpectedMap))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				Expect(projectComplianceProfileModel).ToNot(BeNil())
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")
				Expect(projectComplianceProfileModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.InstanceLocation).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.ProfileName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateProjectEnvironmentOptions model
				projectID := "testString"
				createProjectEnvironmentOptionsModel := projectService.NewCreateProjectEnvironmentOptions(projectID)
				createProjectEnvironmentOptionsModel.SetProjectID("testString")
				createProjectEnvironmentOptionsModel.SetName("development")
				createProjectEnvironmentOptionsModel.SetDescription("The environment 'development'")
				createProjectEnvironmentOptionsModel.SetAuthorizations(projectConfigAuthModel)
				createProjectEnvironmentOptionsModel.SetInputs(inputVariableModel)
				createProjectEnvironmentOptionsModel.SetComplianceProfile(projectComplianceProfileModel)
				createProjectEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProjectEnvironmentOptionsModel).ToNot(BeNil())
				Expect(createProjectEnvironmentOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(createProjectEnvironmentOptionsModel.Name).To(Equal(core.StringPtr("development")))
				Expect(createProjectEnvironmentOptionsModel.Description).To(Equal(core.StringPtr("The environment 'development'")))
				Expect(createProjectEnvironmentOptionsModel.Authorizations).To(Equal(projectConfigAuthModel))
				Expect(createProjectEnvironmentOptionsModel.Inputs).To(Equal(inputVariableModel))
				Expect(createProjectEnvironmentOptionsModel.ComplianceProfile).To(Equal(projectComplianceProfileModel))
				Expect(createProjectEnvironmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProjectOptions successfully`, func() {
				// Construct an instance of the ProjectPrototypeDefinition model
				projectPrototypeDefinitionModel := new(projectv1.ProjectPrototypeDefinition)
				Expect(projectPrototypeDefinitionModel).ToNot(BeNil())
				projectPrototypeDefinitionModel.Name = core.StringPtr("acme-microservice")
				projectPrototypeDefinitionModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				projectPrototypeDefinitionModel.DestroyOnDelete = core.BoolPtr(true)
				Expect(projectPrototypeDefinitionModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(projectPrototypeDefinitionModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure.")))
				Expect(projectPrototypeDefinitionModel.DestroyOnDelete).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				Expect(projectConfigAuthTrustedProfileModel).ToNot(BeNil())
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")
				Expect(projectConfigAuthTrustedProfileModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigAuthTrustedProfileModel.TargetIamID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				Expect(projectConfigAuthModel).ToNot(BeNil())
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")
				Expect(projectConfigAuthModel.TrustedProfile).To(Equal(projectConfigAuthTrustedProfileModel))
				Expect(projectConfigAuthModel.Method).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigAuthModel.ApiKey).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				Expect(projectComplianceProfileModel).ToNot(BeNil())
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")
				Expect(projectComplianceProfileModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.InstanceLocation).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.ProfileName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				Expect(inputVariableModel).ToNot(BeNil())
				inputVariableModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(inputVariableModel.GetProperties()).ToNot(BeEmpty())
				Expect(inputVariableModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				inputVariableModel.SetProperties(nil)
				Expect(inputVariableModel.GetProperties()).To(BeEmpty())

				inputVariableModelExpectedMap := make(map[string]interface{})
				inputVariableModelExpectedMap["foo"] = core.StringPtr("testString")
				inputVariableModel.SetProperties(inputVariableModelExpectedMap)
				inputVariableModelActualMap := inputVariableModel.GetProperties()
				Expect(inputVariableModelActualMap).To(Equal(inputVariableModelExpectedMap))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				Expect(projectConfigSettingModel).ToNot(BeNil())
				projectConfigSettingModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(projectConfigSettingModel.GetProperties()).ToNot(BeEmpty())
				Expect(projectConfigSettingModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				projectConfigSettingModel.SetProperties(nil)
				Expect(projectConfigSettingModel.GetProperties()).To(BeEmpty())

				projectConfigSettingModelExpectedMap := make(map[string]interface{})
				projectConfigSettingModelExpectedMap["foo"] = core.StringPtr("testString")
				projectConfigSettingModel.SetProperties(projectConfigSettingModelExpectedMap)
				projectConfigSettingModelActualMap := projectConfigSettingModel.GetProperties()
				Expect(projectConfigSettingModelActualMap).To(Equal(projectConfigSettingModelExpectedMap))

				// Construct an instance of the ProjectConfigPrototypeDefinitionBlock model
				projectConfigPrototypeDefinitionBlockModel := new(projectv1.ProjectConfigPrototypeDefinitionBlock)
				Expect(projectConfigPrototypeDefinitionBlockModel).ToNot(BeNil())
				projectConfigPrototypeDefinitionBlockModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeDefinitionBlockModel.Description = core.StringPtr("testString")
				projectConfigPrototypeDefinitionBlockModel.Labels = []string{"testString"}
				projectConfigPrototypeDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypeDefinitionBlockModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypeDefinitionBlockModel.Setting = projectConfigSettingModel
				Expect(projectConfigPrototypeDefinitionBlockModel.Name).To(Equal(core.StringPtr("common-variables")))
				Expect(projectConfigPrototypeDefinitionBlockModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigPrototypeDefinitionBlockModel.Labels).To(Equal([]string{"testString"}))
				Expect(projectConfigPrototypeDefinitionBlockModel.Authorizations).To(Equal(projectConfigAuthModel))
				Expect(projectConfigPrototypeDefinitionBlockModel.ComplianceProfile).To(Equal(projectComplianceProfileModel))
				Expect(projectConfigPrototypeDefinitionBlockModel.LocatorID).To(Equal(core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")))
				Expect(projectConfigPrototypeDefinitionBlockModel.Input).To(Equal(inputVariableModel))
				Expect(projectConfigPrototypeDefinitionBlockModel.Setting).To(Equal(projectConfigSettingModel))

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				Expect(projectConfigPrototypeModel).ToNot(BeNil())
				projectConfigPrototypeModel.Definition = projectConfigPrototypeDefinitionBlockModel
				Expect(projectConfigPrototypeModel.Definition).To(Equal(projectConfigPrototypeDefinitionBlockModel))

				// Construct an instance of the CreateProjectOptions model
				resourceGroup := "Default"
				location := "us-south"
				createProjectOptionsModel := projectService.NewCreateProjectOptions(resourceGroup, location)
				createProjectOptionsModel.SetResourceGroup("Default")
				createProjectOptionsModel.SetLocation("us-south")
				createProjectOptionsModel.SetDefinition(projectPrototypeDefinitionModel)
				createProjectOptionsModel.SetConfigs([]projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel})
				createProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProjectOptionsModel).ToNot(BeNil())
				Expect(createProjectOptionsModel.ResourceGroup).To(Equal(core.StringPtr("Default")))
				Expect(createProjectOptionsModel.Location).To(Equal(core.StringPtr("us-south")))
				Expect(createProjectOptionsModel.Definition).To(Equal(projectPrototypeDefinitionModel))
				Expect(createProjectOptionsModel.Configs).To(Equal([]projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel}))
				Expect(createProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteConfigOptions successfully`, func() {
				// Construct an instance of the DeleteConfigOptions model
				projectID := "testString"
				id := "testString"
				deleteConfigOptionsModel := projectService.NewDeleteConfigOptions(projectID, id)
				deleteConfigOptionsModel.SetProjectID("testString")
				deleteConfigOptionsModel.SetID("testString")
				deleteConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteConfigOptionsModel).ToNot(BeNil())
				Expect(deleteConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteConfigVersionOptions successfully`, func() {
				// Construct an instance of the DeleteConfigVersionOptions model
				projectID := "testString"
				id := "testString"
				version := int64(38)
				deleteConfigVersionOptionsModel := projectService.NewDeleteConfigVersionOptions(projectID, id, version)
				deleteConfigVersionOptionsModel.SetProjectID("testString")
				deleteConfigVersionOptionsModel.SetID("testString")
				deleteConfigVersionOptionsModel.SetVersion(int64(38))
				deleteConfigVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteConfigVersionOptionsModel).ToNot(BeNil())
				Expect(deleteConfigVersionOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigVersionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigVersionOptionsModel.Version).To(Equal(core.Int64Ptr(int64(38))))
				Expect(deleteConfigVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProjectEnvironmentOptions successfully`, func() {
				// Construct an instance of the DeleteProjectEnvironmentOptions model
				projectID := "testString"
				id := "testString"
				deleteProjectEnvironmentOptionsModel := projectService.NewDeleteProjectEnvironmentOptions(projectID, id)
				deleteProjectEnvironmentOptionsModel.SetProjectID("testString")
				deleteProjectEnvironmentOptionsModel.SetID("testString")
				deleteProjectEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProjectEnvironmentOptionsModel).ToNot(BeNil())
				Expect(deleteProjectEnvironmentOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProjectEnvironmentOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProjectEnvironmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProjectOptions successfully`, func() {
				// Construct an instance of the DeleteProjectOptions model
				id := "testString"
				deleteProjectOptionsModel := projectService.NewDeleteProjectOptions(id)
				deleteProjectOptionsModel.SetID("testString")
				deleteProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProjectOptionsModel).ToNot(BeNil())
				Expect(deleteProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeployConfigOptions successfully`, func() {
				// Construct an instance of the DeployConfigOptions model
				projectID := "testString"
				id := "testString"
				deployConfigOptionsModel := projectService.NewDeployConfigOptions(projectID, id)
				deployConfigOptionsModel.SetProjectID("testString")
				deployConfigOptionsModel.SetID("testString")
				deployConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deployConfigOptionsModel).ToNot(BeNil())
				Expect(deployConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(deployConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deployConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewForceApproveOptions successfully`, func() {
				// Construct an instance of the ForceApproveOptions model
				projectID := "testString"
				id := "testString"
				forceApproveOptionsModel := projectService.NewForceApproveOptions(projectID, id)
				forceApproveOptionsModel.SetProjectID("testString")
				forceApproveOptionsModel.SetID("testString")
				forceApproveOptionsModel.SetComment("Approving the changes")
				forceApproveOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(forceApproveOptionsModel).ToNot(BeNil())
				Expect(forceApproveOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(forceApproveOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(forceApproveOptionsModel.Comment).To(Equal(core.StringPtr("Approving the changes")))
				Expect(forceApproveOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConfigOptions successfully`, func() {
				// Construct an instance of the GetConfigOptions model
				projectID := "testString"
				id := "testString"
				getConfigOptionsModel := projectService.NewGetConfigOptions(projectID, id)
				getConfigOptionsModel.SetProjectID("testString")
				getConfigOptionsModel.SetID("testString")
				getConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConfigOptionsModel).ToNot(BeNil())
				Expect(getConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(getConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConfigVersionOptions successfully`, func() {
				// Construct an instance of the GetConfigVersionOptions model
				projectID := "testString"
				id := "testString"
				version := int64(38)
				getConfigVersionOptionsModel := projectService.NewGetConfigVersionOptions(projectID, id, version)
				getConfigVersionOptionsModel.SetProjectID("testString")
				getConfigVersionOptionsModel.SetID("testString")
				getConfigVersionOptionsModel.SetVersion(int64(38))
				getConfigVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConfigVersionOptionsModel).ToNot(BeNil())
				Expect(getConfigVersionOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(getConfigVersionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getConfigVersionOptionsModel.Version).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getConfigVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectEnvironmentOptions successfully`, func() {
				// Construct an instance of the GetProjectEnvironmentOptions model
				projectID := "testString"
				id := "testString"
				getProjectEnvironmentOptionsModel := projectService.NewGetProjectEnvironmentOptions(projectID, id)
				getProjectEnvironmentOptionsModel.SetProjectID("testString")
				getProjectEnvironmentOptionsModel.SetID("testString")
				getProjectEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectEnvironmentOptionsModel).ToNot(BeNil())
				Expect(getProjectEnvironmentOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(getProjectEnvironmentOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getProjectEnvironmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProjectOptions successfully`, func() {
				// Construct an instance of the GetProjectOptions model
				id := "testString"
				getProjectOptionsModel := projectService.NewGetProjectOptions(id)
				getProjectOptionsModel.SetID("testString")
				getProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProjectOptionsModel).ToNot(BeNil())
				Expect(getProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListConfigResourcesOptions successfully`, func() {
				// Construct an instance of the ListConfigResourcesOptions model
				projectID := "testString"
				id := "testString"
				listConfigResourcesOptionsModel := projectService.NewListConfigResourcesOptions(projectID, id)
				listConfigResourcesOptionsModel.SetProjectID("testString")
				listConfigResourcesOptionsModel.SetID("testString")
				listConfigResourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listConfigResourcesOptionsModel).ToNot(BeNil())
				Expect(listConfigResourcesOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(listConfigResourcesOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listConfigResourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListConfigVersionsOptions successfully`, func() {
				// Construct an instance of the ListConfigVersionsOptions model
				projectID := "testString"
				id := "testString"
				listConfigVersionsOptionsModel := projectService.NewListConfigVersionsOptions(projectID, id)
				listConfigVersionsOptionsModel.SetProjectID("testString")
				listConfigVersionsOptionsModel.SetID("testString")
				listConfigVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listConfigVersionsOptionsModel).ToNot(BeNil())
				Expect(listConfigVersionsOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(listConfigVersionsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(listConfigVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListConfigsOptions successfully`, func() {
				// Construct an instance of the ListConfigsOptions model
				projectID := "testString"
				listConfigsOptionsModel := projectService.NewListConfigsOptions(projectID)
				listConfigsOptionsModel.SetProjectID("testString")
				listConfigsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listConfigsOptionsModel).ToNot(BeNil())
				Expect(listConfigsOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(listConfigsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListEnvironmentsOptions successfully`, func() {
				// Construct an instance of the ListEnvironmentsOptions model
				listEnvironmentsOptionsModel := projectService.NewListEnvironmentsOptions()
				listEnvironmentsOptionsModel.SetName(false)
				listEnvironmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listEnvironmentsOptionsModel).ToNot(BeNil())
				Expect(listEnvironmentsOptionsModel.Name).To(Equal(core.BoolPtr(false)))
				Expect(listEnvironmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProjectEnvironmentsOptions successfully`, func() {
				// Construct an instance of the ListProjectEnvironmentsOptions model
				projectID := "testString"
				listProjectEnvironmentsOptionsModel := projectService.NewListProjectEnvironmentsOptions(projectID)
				listProjectEnvironmentsOptionsModel.SetProjectID("testString")
				listProjectEnvironmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProjectEnvironmentsOptionsModel).ToNot(BeNil())
				Expect(listProjectEnvironmentsOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(listProjectEnvironmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProjectsOptions successfully`, func() {
				// Construct an instance of the ListProjectsOptions model
				listProjectsOptionsModel := projectService.NewListProjectsOptions()
				listProjectsOptionsModel.SetStart("testString")
				listProjectsOptionsModel.SetLimit(int64(10))
				listProjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProjectsOptionsModel).ToNot(BeNil())
				Expect(listProjectsOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(listProjectsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listProjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewProjectConfigPrototypeDefinitionBlock successfully`, func() {
				name := "testString"
				locatorID := "testString"
				_model, err := projectService.NewProjectConfigPrototypeDefinitionBlock(name, locatorID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProjectPrototypeDefinition successfully`, func() {
				name := "testString"
				_model, err := projectService.NewProjectPrototypeDefinition(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSyncConfigOptions successfully`, func() {
				// Construct an instance of the SyncConfigOptions model
				projectID := "testString"
				id := "testString"
				syncConfigOptionsModel := projectService.NewSyncConfigOptions(projectID, id)
				syncConfigOptionsModel.SetProjectID("testString")
				syncConfigOptionsModel.SetID("testString")
				syncConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(syncConfigOptionsModel).ToNot(BeNil())
				Expect(syncConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(syncConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(syncConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUndeployConfigOptions successfully`, func() {
				// Construct an instance of the UndeployConfigOptions model
				projectID := "testString"
				id := "testString"
				undeployConfigOptionsModel := projectService.NewUndeployConfigOptions(projectID, id)
				undeployConfigOptionsModel.SetProjectID("testString")
				undeployConfigOptionsModel.SetID("testString")
				undeployConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(undeployConfigOptionsModel).ToNot(BeNil())
				Expect(undeployConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(undeployConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(undeployConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateConfigOptions successfully`, func() {
				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				Expect(projectConfigAuthTrustedProfileModel).ToNot(BeNil())
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")
				Expect(projectConfigAuthTrustedProfileModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigAuthTrustedProfileModel.TargetIamID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				Expect(projectConfigAuthModel).ToNot(BeNil())
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")
				Expect(projectConfigAuthModel.TrustedProfile).To(Equal(projectConfigAuthTrustedProfileModel))
				Expect(projectConfigAuthModel.Method).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigAuthModel.ApiKey).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				Expect(projectComplianceProfileModel).ToNot(BeNil())
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")
				Expect(projectComplianceProfileModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.InstanceLocation).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.ProfileName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				Expect(inputVariableModel).ToNot(BeNil())
				inputVariableModel.SetProperty("account_id", core.StringPtr(`$configs[].name["account-stage"].input.account_id`))
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("access_tags", core.StringPtr(`["env:stage"]`))
				inputVariableModel.SetProperty("logdna_name", core.StringPtr("Name of the LogDNA stage service instance"))
				inputVariableModel.SetProperty("sysdig_name", core.StringPtr("Name of the SysDig stage service instance"))
				Expect(inputVariableModel.GetProperties()).ToNot(BeEmpty())
				Expect(inputVariableModel.GetProperty("account_id")).To(Equal(core.StringPtr(`$configs[].name["account-stage"].input.account_id`)))
				Expect(inputVariableModel.GetProperty("resource_group")).To(Equal(core.StringPtr("stage")))
				Expect(inputVariableModel.GetProperty("access_tags")).To(Equal(core.StringPtr(`["env:stage"]`)))
				Expect(inputVariableModel.GetProperty("logdna_name")).To(Equal(core.StringPtr("Name of the LogDNA stage service instance")))
				Expect(inputVariableModel.GetProperty("sysdig_name")).To(Equal(core.StringPtr("Name of the SysDig stage service instance")))

				inputVariableModel.SetProperties(nil)
				Expect(inputVariableModel.GetProperties()).To(BeEmpty())

				inputVariableModelExpectedMap := make(map[string]interface{})
				inputVariableModelExpectedMap["account_id"] = core.StringPtr(`$configs[].name["account-stage"].input.account_id`)
				inputVariableModelExpectedMap["resource_group"] = core.StringPtr("stage")
				inputVariableModelExpectedMap["access_tags"] = core.StringPtr(`["env:stage"]`)
				inputVariableModelExpectedMap["logdna_name"] = core.StringPtr("Name of the LogDNA stage service instance")
				inputVariableModelExpectedMap["sysdig_name"] = core.StringPtr("Name of the SysDig stage service instance")
				inputVariableModel.SetProperties(inputVariableModelExpectedMap)
				inputVariableModelActualMap := inputVariableModel.GetProperties()
				Expect(inputVariableModelActualMap).To(Equal(inputVariableModelExpectedMap))

				// Construct an instance of the ProjectConfigSetting model
				projectConfigSettingModel := new(projectv1.ProjectConfigSetting)
				Expect(projectConfigSettingModel).ToNot(BeNil())
				projectConfigSettingModel.SetProperty("foo", core.StringPtr("testString"))
				Expect(projectConfigSettingModel.GetProperties()).ToNot(BeEmpty())
				Expect(projectConfigSettingModel.GetProperty("foo")).To(Equal(core.StringPtr("testString")))

				projectConfigSettingModel.SetProperties(nil)
				Expect(projectConfigSettingModel.GetProperties()).To(BeEmpty())

				projectConfigSettingModelExpectedMap := make(map[string]interface{})
				projectConfigSettingModelExpectedMap["foo"] = core.StringPtr("testString")
				projectConfigSettingModel.SetProperties(projectConfigSettingModelExpectedMap)
				projectConfigSettingModelActualMap := projectConfigSettingModel.GetProperties()
				Expect(projectConfigSettingModelActualMap).To(Equal(projectConfigSettingModelExpectedMap))

				// Construct an instance of the ProjectConfigPrototypePatchDefinitionBlock model
				projectConfigPrototypePatchDefinitionBlockModel := new(projectv1.ProjectConfigPrototypePatchDefinitionBlock)
				Expect(projectConfigPrototypePatchDefinitionBlockModel).ToNot(BeNil())
				projectConfigPrototypePatchDefinitionBlockModel.Name = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Description = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Labels = []string{"testString"}
				projectConfigPrototypePatchDefinitionBlockModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypePatchDefinitionBlockModel.ComplianceProfile = projectComplianceProfileModel
				projectConfigPrototypePatchDefinitionBlockModel.LocatorID = core.StringPtr("testString")
				projectConfigPrototypePatchDefinitionBlockModel.Input = inputVariableModel
				projectConfigPrototypePatchDefinitionBlockModel.Setting = projectConfigSettingModel
				Expect(projectConfigPrototypePatchDefinitionBlockModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigPrototypePatchDefinitionBlockModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigPrototypePatchDefinitionBlockModel.Labels).To(Equal([]string{"testString"}))
				Expect(projectConfigPrototypePatchDefinitionBlockModel.Authorizations).To(Equal(projectConfigAuthModel))
				Expect(projectConfigPrototypePatchDefinitionBlockModel.ComplianceProfile).To(Equal(projectComplianceProfileModel))
				Expect(projectConfigPrototypePatchDefinitionBlockModel.LocatorID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigPrototypePatchDefinitionBlockModel.Input).To(Equal(inputVariableModel))
				Expect(projectConfigPrototypePatchDefinitionBlockModel.Setting).To(Equal(projectConfigSettingModel))

				// Construct an instance of the UpdateConfigOptions model
				projectID := "testString"
				id := "testString"
				updateConfigOptionsModel := projectService.NewUpdateConfigOptions(projectID, id)
				updateConfigOptionsModel.SetProjectID("testString")
				updateConfigOptionsModel.SetID("testString")
				updateConfigOptionsModel.SetDefinition(projectConfigPrototypePatchDefinitionBlockModel)
				updateConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateConfigOptionsModel).ToNot(BeNil())
				Expect(updateConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigOptionsModel.Definition).To(Equal(projectConfigPrototypePatchDefinitionBlockModel))
				Expect(updateConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProjectEnvironmentOptions successfully`, func() {
				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				Expect(projectConfigAuthTrustedProfileModel).ToNot(BeNil())
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")
				Expect(projectConfigAuthTrustedProfileModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigAuthTrustedProfileModel.TargetIamID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				Expect(projectConfigAuthModel).ToNot(BeNil())
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("API_KEY")
				projectConfigAuthModel.ApiKey = core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9")
				Expect(projectConfigAuthModel.TrustedProfile).To(Equal(projectConfigAuthTrustedProfileModel))
				Expect(projectConfigAuthModel.Method).To(Equal(core.StringPtr("API_KEY")))
				Expect(projectConfigAuthModel.ApiKey).To(Equal(core.StringPtr("TbcdlprpFODhkpns9e0daOWnAwd2tXwSYtPn8rpEd8d9")))

				// Construct an instance of the InputVariable model
				inputVariableModel := new(projectv1.InputVariable)
				Expect(inputVariableModel).ToNot(BeNil())
				inputVariableModel.SetProperty("resource_group", core.StringPtr("stage"))
				inputVariableModel.SetProperty("region", core.StringPtr("us-south"))
				Expect(inputVariableModel.GetProperties()).ToNot(BeEmpty())
				Expect(inputVariableModel.GetProperty("resource_group")).To(Equal(core.StringPtr("stage")))
				Expect(inputVariableModel.GetProperty("region")).To(Equal(core.StringPtr("us-south")))

				inputVariableModel.SetProperties(nil)
				Expect(inputVariableModel.GetProperties()).To(BeEmpty())

				inputVariableModelExpectedMap := make(map[string]interface{})
				inputVariableModelExpectedMap["resource_group"] = core.StringPtr("stage")
				inputVariableModelExpectedMap["region"] = core.StringPtr("us-south")
				inputVariableModel.SetProperties(inputVariableModelExpectedMap)
				inputVariableModelActualMap := inputVariableModel.GetProperties()
				Expect(inputVariableModelActualMap).To(Equal(inputVariableModelExpectedMap))

				// Construct an instance of the ProjectComplianceProfile model
				projectComplianceProfileModel := new(projectv1.ProjectComplianceProfile)
				Expect(projectComplianceProfileModel).ToNot(BeNil())
				projectComplianceProfileModel.ID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectComplianceProfileModel.ProfileName = core.StringPtr("testString")
				Expect(projectComplianceProfileModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.InstanceLocation).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(projectComplianceProfileModel.ProfileName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateProjectEnvironmentOptions model
				projectID := "testString"
				id := "testString"
				updateProjectEnvironmentOptionsModel := projectService.NewUpdateProjectEnvironmentOptions(projectID, id)
				updateProjectEnvironmentOptionsModel.SetProjectID("testString")
				updateProjectEnvironmentOptionsModel.SetID("testString")
				updateProjectEnvironmentOptionsModel.SetName("dev")
				updateProjectEnvironmentOptionsModel.SetDescription("The new environment 'dev'")
				updateProjectEnvironmentOptionsModel.SetAuthorizations(projectConfigAuthModel)
				updateProjectEnvironmentOptionsModel.SetInputs(inputVariableModel)
				updateProjectEnvironmentOptionsModel.SetComplianceProfile(projectComplianceProfileModel)
				updateProjectEnvironmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProjectEnvironmentOptionsModel).ToNot(BeNil())
				Expect(updateProjectEnvironmentOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectEnvironmentOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectEnvironmentOptionsModel.Name).To(Equal(core.StringPtr("dev")))
				Expect(updateProjectEnvironmentOptionsModel.Description).To(Equal(core.StringPtr("The new environment 'dev'")))
				Expect(updateProjectEnvironmentOptionsModel.Authorizations).To(Equal(projectConfigAuthModel))
				Expect(updateProjectEnvironmentOptionsModel.Inputs).To(Equal(inputVariableModel))
				Expect(updateProjectEnvironmentOptionsModel.ComplianceProfile).To(Equal(projectComplianceProfileModel))
				Expect(updateProjectEnvironmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProjectOptions successfully`, func() {
				// Construct an instance of the ProjectPrototypePatchDefinitionBlock model
				projectPrototypePatchDefinitionBlockModel := new(projectv1.ProjectPrototypePatchDefinitionBlock)
				Expect(projectPrototypePatchDefinitionBlockModel).ToNot(BeNil())
				projectPrototypePatchDefinitionBlockModel.Name = core.StringPtr("acme-microservice")
				projectPrototypePatchDefinitionBlockModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				projectPrototypePatchDefinitionBlockModel.DestroyOnDelete = core.BoolPtr(true)
				Expect(projectPrototypePatchDefinitionBlockModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(projectPrototypePatchDefinitionBlockModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure.")))
				Expect(projectPrototypePatchDefinitionBlockModel.DestroyOnDelete).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the UpdateProjectOptions model
				id := "testString"
				updateProjectOptionsModel := projectService.NewUpdateProjectOptions(id)
				updateProjectOptionsModel.SetID("testString")
				updateProjectOptionsModel.SetDefinition(projectPrototypePatchDefinitionBlockModel)
				updateProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProjectOptionsModel).ToNot(BeNil())
				Expect(updateProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateProjectOptionsModel.Definition).To(Equal(projectPrototypePatchDefinitionBlockModel))
				Expect(updateProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewValidateConfigOptions successfully`, func() {
				// Construct an instance of the ValidateConfigOptions model
				projectID := "testString"
				id := "testString"
				validateConfigOptionsModel := projectService.NewValidateConfigOptions(projectID, id)
				validateConfigOptionsModel.SetProjectID("testString")
				validateConfigOptionsModel.SetID("testString")
				validateConfigOptionsModel.SetXAuthRefreshToken("testString")
				validateConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(validateConfigOptionsModel).ToNot(BeNil())
				Expect(validateConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(validateConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(validateConfigOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(validateConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
