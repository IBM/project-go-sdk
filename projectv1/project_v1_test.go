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

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigComplianceProfile model
				projectConfigComplianceProfileModel := new(projectv1.ProjectConfigComplianceProfile)
				projectConfigComplianceProfileModel.ID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectConfigComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("testString")
				projectConfigInputVariableModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.ID = core.StringPtr("testString")
				projectConfigPrototypeModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeModel.Labels = []string{"testString"}
				projectConfigPrototypeModel.Description = core.StringPtr("testString")
				projectConfigPrototypeModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeModel.ComplianceProfile = projectConfigComplianceProfileModel
				projectConfigPrototypeModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPrototypeModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				createProjectOptionsModel.DestroyOnDelete = core.BoolPtr(true)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "destroy_on_delete": true, "id": "ID", "crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}], "metadata": {"crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}`)
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

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigComplianceProfile model
				projectConfigComplianceProfileModel := new(projectv1.ProjectConfigComplianceProfile)
				projectConfigComplianceProfileModel.ID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectConfigComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("testString")
				projectConfigInputVariableModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.ID = core.StringPtr("testString")
				projectConfigPrototypeModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeModel.Labels = []string{"testString"}
				projectConfigPrototypeModel.Description = core.StringPtr("testString")
				projectConfigPrototypeModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeModel.ComplianceProfile = projectConfigComplianceProfileModel
				projectConfigPrototypeModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPrototypeModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				createProjectOptionsModel.DestroyOnDelete = core.BoolPtr(true)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "destroy_on_delete": true, "id": "ID", "crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}], "metadata": {"crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}`)
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

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigComplianceProfile model
				projectConfigComplianceProfileModel := new(projectv1.ProjectConfigComplianceProfile)
				projectConfigComplianceProfileModel.ID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectConfigComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("testString")
				projectConfigInputVariableModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.ID = core.StringPtr("testString")
				projectConfigPrototypeModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeModel.Labels = []string{"testString"}
				projectConfigPrototypeModel.Description = core.StringPtr("testString")
				projectConfigPrototypeModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeModel.ComplianceProfile = projectConfigComplianceProfileModel
				projectConfigPrototypeModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPrototypeModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				createProjectOptionsModel.DestroyOnDelete = core.BoolPtr(true)
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

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigComplianceProfile model
				projectConfigComplianceProfileModel := new(projectv1.ProjectConfigComplianceProfile)
				projectConfigComplianceProfileModel.ID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectConfigComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("testString")
				projectConfigInputVariableModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.ID = core.StringPtr("testString")
				projectConfigPrototypeModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeModel.Labels = []string{"testString"}
				projectConfigPrototypeModel.Description = core.StringPtr("testString")
				projectConfigPrototypeModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeModel.ComplianceProfile = projectConfigComplianceProfileModel
				projectConfigPrototypeModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPrototypeModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				createProjectOptionsModel.DestroyOnDelete = core.BoolPtr(true)
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

				// Construct an instance of the ProjectConfigAuthTrustedProfile model
				projectConfigAuthTrustedProfileModel := new(projectv1.ProjectConfigAuthTrustedProfile)
				projectConfigAuthTrustedProfileModel.ID = core.StringPtr("testString")
				projectConfigAuthTrustedProfileModel.TargetIamID = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigAuth model
				projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
				projectConfigAuthModel.TrustedProfile = projectConfigAuthTrustedProfileModel
				projectConfigAuthModel.Method = core.StringPtr("testString")
				projectConfigAuthModel.ApiKey = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigComplianceProfile model
				projectConfigComplianceProfileModel := new(projectv1.ProjectConfigComplianceProfile)
				projectConfigComplianceProfileModel.ID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectConfigComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("testString")
				projectConfigInputVariableModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				projectConfigPrototypeModel.ID = core.StringPtr("testString")
				projectConfigPrototypeModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeModel.Labels = []string{"testString"}
				projectConfigPrototypeModel.Description = core.StringPtr("testString")
				projectConfigPrototypeModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeModel.ComplianceProfile = projectConfigComplianceProfileModel
				projectConfigPrototypeModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPrototypeModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the CreateProjectOptions model
				createProjectOptionsModel := new(projectv1.CreateProjectOptions)
				createProjectOptionsModel.ResourceGroup = core.StringPtr("Default")
				createProjectOptionsModel.Location = core.StringPtr("us-south")
				createProjectOptionsModel.Name = core.StringPtr("acme-microservice")
				createProjectOptionsModel.Description = core.StringPtr("A microservice to deploy on top of ACME infrastructure.")
				createProjectOptionsModel.DestroyOnDelete = core.BoolPtr(true)
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
					fmt.Fprintf(res, "%s", `{"limit": 10, "total_count": 0, "first": {"href": "Href", "start": "Start"}, "last": {"href": "Href", "start": "Start"}, "previous": {"href": "Href", "start": "Start"}, "next": {"href": "Href", "start": "Start"}, "projects": [{"id": "ID", "name": "Name", "description": "Description", "metadata": {"crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}]}`)
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
					fmt.Fprintf(res, "%s", `{"limit": 10, "total_count": 0, "first": {"href": "Href", "start": "Start"}, "last": {"href": "Href", "start": "Start"}, "previous": {"href": "Href", "start": "Start"}, "next": {"href": "Href", "start": "Start"}, "projects": [{"id": "ID", "name": "Name", "description": "Description", "metadata": {"crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}]}`)
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
				nextObject.Start = core.StringPtr("abc-123")
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
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"projects":[{"id":"ID","name":"Name","description":"Description","metadata":{"crn":"crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::","created_at":"2019-01-01T12:00:00.000Z","cumulative_needs_attention_view":[{"event":"Event","event_id":"EventID","config_id":"ConfigID","config_version":13}],"cumulative_needs_attention_view_err":"CumulativeNeedsAttentionViewErr","location":"Location","resource_group":"ResourceGroup","state":"State","event_notifications_crn":"EventNotificationsCrn"}}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"projects":[{"id":"ID","name":"Name","description":"Description","metadata":{"crn":"crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::","created_at":"2019-01-01T12:00:00.000Z","cumulative_needs_attention_view":[{"event":"Event","event_id":"EventID","config_id":"ConfigID","config_version":13}],"cumulative_needs_attention_view_err":"CumulativeNeedsAttentionViewErr","location":"Location","resource_group":"ResourceGroup","state":"State","event_notifications_crn":"EventNotificationsCrn"}}],"total_count":2,"limit":1}`)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "destroy_on_delete": true, "id": "ID", "crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}], "metadata": {"crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}`)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "destroy_on_delete": true, "id": "ID", "crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}], "metadata": {"crn": "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::", "created_at": "2019-01-01T12:00:00.000Z", "cumulative_needs_attention_view": [{"event": "Event", "event_id": "EventID", "config_id": "ConfigID", "config_version": 13}], "cumulative_needs_attention_view_err": "CumulativeNeedsAttentionViewErr", "location": "Location", "resource_group": "ResourceGroup", "state": "State", "event_notifications_crn": "EventNotificationsCrn"}}`)
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
	Describe(`DeleteProject(deleteProjectOptions *DeleteProjectOptions)`, func() {
		deleteProjectPath := "/v1/projects/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProjectPath))
					Expect(req.Method).To(Equal("DELETE"))

					// TODO: Add check for destroy query parameter
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
				deleteProjectOptionsModel.Destroy = core.BoolPtr(true)
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
				deleteProjectOptionsModel.Destroy = core.BoolPtr(true)
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

				// Construct an instance of the ProjectConfigComplianceProfile model
				projectConfigComplianceProfileModel := new(projectv1.ProjectConfigComplianceProfile)
				projectConfigComplianceProfileModel.ID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectConfigComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")
				projectConfigInputVariableModel.Value = core.StringPtr(`$configs[].name["account-stage"].input.account_id`)

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				projectConfigSettingCollectionModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ProjectID = core.StringPtr("testString")
				createConfigOptionsModel.Name = core.StringPtr("env-stage")
				createConfigOptionsModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.Labels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.Description = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				createConfigOptionsModel.Authorizations = projectConfigAuthModel
				createConfigOptionsModel.ComplianceProfile = projectConfigComplianceProfileModel
				createConfigOptionsModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				createConfigOptionsModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}`)
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

				// Construct an instance of the ProjectConfigComplianceProfile model
				projectConfigComplianceProfileModel := new(projectv1.ProjectConfigComplianceProfile)
				projectConfigComplianceProfileModel.ID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectConfigComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")
				projectConfigInputVariableModel.Value = core.StringPtr(`$configs[].name["account-stage"].input.account_id`)

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				projectConfigSettingCollectionModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ProjectID = core.StringPtr("testString")
				createConfigOptionsModel.Name = core.StringPtr("env-stage")
				createConfigOptionsModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.Labels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.Description = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				createConfigOptionsModel.Authorizations = projectConfigAuthModel
				createConfigOptionsModel.ComplianceProfile = projectConfigComplianceProfileModel
				createConfigOptionsModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				createConfigOptionsModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}`)
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

				// Construct an instance of the ProjectConfigComplianceProfile model
				projectConfigComplianceProfileModel := new(projectv1.ProjectConfigComplianceProfile)
				projectConfigComplianceProfileModel.ID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectConfigComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")
				projectConfigInputVariableModel.Value = core.StringPtr(`$configs[].name["account-stage"].input.account_id`)

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				projectConfigSettingCollectionModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ProjectID = core.StringPtr("testString")
				createConfigOptionsModel.Name = core.StringPtr("env-stage")
				createConfigOptionsModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.Labels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.Description = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				createConfigOptionsModel.Authorizations = projectConfigAuthModel
				createConfigOptionsModel.ComplianceProfile = projectConfigComplianceProfileModel
				createConfigOptionsModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				createConfigOptionsModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}
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

				// Construct an instance of the ProjectConfigComplianceProfile model
				projectConfigComplianceProfileModel := new(projectv1.ProjectConfigComplianceProfile)
				projectConfigComplianceProfileModel.ID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectConfigComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")
				projectConfigInputVariableModel.Value = core.StringPtr(`$configs[].name["account-stage"].input.account_id`)

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				projectConfigSettingCollectionModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ProjectID = core.StringPtr("testString")
				createConfigOptionsModel.Name = core.StringPtr("env-stage")
				createConfigOptionsModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.Labels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.Description = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				createConfigOptionsModel.Authorizations = projectConfigAuthModel
				createConfigOptionsModel.ComplianceProfile = projectConfigComplianceProfileModel
				createConfigOptionsModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				createConfigOptionsModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}
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

				// Construct an instance of the ProjectConfigComplianceProfile model
				projectConfigComplianceProfileModel := new(projectv1.ProjectConfigComplianceProfile)
				projectConfigComplianceProfileModel.ID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectConfigComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.ProfileName = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")
				projectConfigInputVariableModel.Value = core.StringPtr(`$configs[].name["account-stage"].input.account_id`)

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				projectConfigSettingCollectionModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")

				// Construct an instance of the CreateConfigOptions model
				createConfigOptionsModel := new(projectv1.CreateConfigOptions)
				createConfigOptionsModel.ProjectID = core.StringPtr("testString")
				createConfigOptionsModel.Name = core.StringPtr("env-stage")
				createConfigOptionsModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.ID = core.StringPtr("testString")
				createConfigOptionsModel.Labels = []string{"env:stage", "governance:test", "build:0"}
				createConfigOptionsModel.Description = core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				createConfigOptionsModel.Authorizations = projectConfigAuthModel
				createConfigOptionsModel.ComplianceProfile = projectConfigComplianceProfileModel
				createConfigOptionsModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				createConfigOptionsModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}
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
					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
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
				listConfigsOptionsModel.Version = core.StringPtr("active")
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

					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}]}`)
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
				listConfigsOptionsModel.Version = core.StringPtr("active")
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

					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configs": [{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}]}`)
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
				listConfigsOptionsModel.Version = core.StringPtr("active")
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
				listConfigsOptionsModel.Version = core.StringPtr("active")
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
				listConfigsOptionsModel.Version = core.StringPtr("active")
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
					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
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
				getConfigOptionsModel.Version = core.StringPtr("active")
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

					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}`)
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
				getConfigOptionsModel.Version = core.StringPtr("active")
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

					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}`)
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
				getConfigOptionsModel.Version = core.StringPtr("active")
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
				getConfigOptionsModel.Version = core.StringPtr("active")
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
				getConfigOptionsModel.Version = core.StringPtr("active")
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

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")
				projectConfigInputVariableModel.Value = core.StringPtr(`$configs[].name["account-stage"].input.account_id`)

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPatchRequestProjectConfigPatchSchematicsTemplate model
				projectConfigPatchRequestModel := new(projectv1.ProjectConfigPatchRequestProjectConfigPatchSchematicsTemplate)
				projectConfigPatchRequestModel.Name = core.StringPtr("testString")
				projectConfigPatchRequestModel.Labels = []string{"testString"}
				projectConfigPatchRequestModel.Description = core.StringPtr("testString")
				projectConfigPatchRequestModel.LocatorID = core.StringPtr("testString")
				projectConfigPatchRequestModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPatchRequestModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = projectConfigPatchRequestModel
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}`)
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

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")
				projectConfigInputVariableModel.Value = core.StringPtr(`$configs[].name["account-stage"].input.account_id`)

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPatchRequestProjectConfigPatchSchematicsTemplate model
				projectConfigPatchRequestModel := new(projectv1.ProjectConfigPatchRequestProjectConfigPatchSchematicsTemplate)
				projectConfigPatchRequestModel.Name = core.StringPtr("testString")
				projectConfigPatchRequestModel.Labels = []string{"testString"}
				projectConfigPatchRequestModel.Description = core.StringPtr("testString")
				projectConfigPatchRequestModel.LocatorID = core.StringPtr("testString")
				projectConfigPatchRequestModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPatchRequestModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = projectConfigPatchRequestModel
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}`)
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

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")
				projectConfigInputVariableModel.Value = core.StringPtr(`$configs[].name["account-stage"].input.account_id`)

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPatchRequestProjectConfigPatchSchematicsTemplate model
				projectConfigPatchRequestModel := new(projectv1.ProjectConfigPatchRequestProjectConfigPatchSchematicsTemplate)
				projectConfigPatchRequestModel.Name = core.StringPtr("testString")
				projectConfigPatchRequestModel.Labels = []string{"testString"}
				projectConfigPatchRequestModel.Description = core.StringPtr("testString")
				projectConfigPatchRequestModel.LocatorID = core.StringPtr("testString")
				projectConfigPatchRequestModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPatchRequestModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = projectConfigPatchRequestModel
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

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")
				projectConfigInputVariableModel.Value = core.StringPtr(`$configs[].name["account-stage"].input.account_id`)

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPatchRequestProjectConfigPatchSchematicsTemplate model
				projectConfigPatchRequestModel := new(projectv1.ProjectConfigPatchRequestProjectConfigPatchSchematicsTemplate)
				projectConfigPatchRequestModel.Name = core.StringPtr("testString")
				projectConfigPatchRequestModel.Labels = []string{"testString"}
				projectConfigPatchRequestModel.Description = core.StringPtr("testString")
				projectConfigPatchRequestModel.LocatorID = core.StringPtr("testString")
				projectConfigPatchRequestModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPatchRequestModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = projectConfigPatchRequestModel
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

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")
				projectConfigInputVariableModel.Value = core.StringPtr(`$configs[].name["account-stage"].input.account_id`)

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")

				// Construct an instance of the ProjectConfigPatchRequestProjectConfigPatchSchematicsTemplate model
				projectConfigPatchRequestModel := new(projectv1.ProjectConfigPatchRequestProjectConfigPatchSchematicsTemplate)
				projectConfigPatchRequestModel.Name = core.StringPtr("testString")
				projectConfigPatchRequestModel.Labels = []string{"testString"}
				projectConfigPatchRequestModel.Description = core.StringPtr("testString")
				projectConfigPatchRequestModel.LocatorID = core.StringPtr("testString")
				projectConfigPatchRequestModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPatchRequestModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}

				// Construct an instance of the UpdateConfigOptions model
				updateConfigOptionsModel := new(projectv1.UpdateConfigOptions)
				updateConfigOptionsModel.ProjectID = core.StringPtr("testString")
				updateConfigOptionsModel.ID = core.StringPtr("testString")
				updateConfigOptionsModel.ProjectConfig = projectConfigPatchRequestModel
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
					// TODO: Add check for draft_only query parameter
					// TODO: Add check for destroy query parameter
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
				deleteConfigOptionsModel.DraftOnly = core.BoolPtr(false)
				deleteConfigOptionsModel.Destroy = core.BoolPtr(true)
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

					// TODO: Add check for draft_only query parameter
					// TODO: Add check for destroy query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name"}`)
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
				deleteConfigOptionsModel.DraftOnly = core.BoolPtr(false)
				deleteConfigOptionsModel.Destroy = core.BoolPtr(true)
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

					// TODO: Add check for draft_only query parameter
					// TODO: Add check for destroy query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name"}`)
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
				deleteConfigOptionsModel.DraftOnly = core.BoolPtr(false)
				deleteConfigOptionsModel.Destroy = core.BoolPtr(true)
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
				deleteConfigOptionsModel.DraftOnly = core.BoolPtr(false)
				deleteConfigOptionsModel.Destroy = core.BoolPtr(true)
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
				deleteConfigOptionsModel.DraftOnly = core.BoolPtr(false)
				deleteConfigOptionsModel.Destroy = core.BoolPtr(true)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}`)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}`)
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
	Describe(`CheckConfig(checkConfigOptions *CheckConfigOptions) - Operation response error`, func() {
		checkConfigPath := "/v1/projects/testString/configs/testString/check"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(checkConfigPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CheckConfig with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the CheckConfigOptions model
				checkConfigOptionsModel := new(projectv1.CheckConfigOptions)
				checkConfigOptionsModel.ProjectID = core.StringPtr("testString")
				checkConfigOptionsModel.ID = core.StringPtr("testString")
				checkConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				checkConfigOptionsModel.Version = core.StringPtr("active")
				checkConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.CheckConfig(checkConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.CheckConfig(checkConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CheckConfig(checkConfigOptions *CheckConfigOptions)`, func() {
		checkConfigPath := "/v1/projects/testString/configs/testString/check"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(checkConfigPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke CheckConfig successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the CheckConfigOptions model
				checkConfigOptionsModel := new(projectv1.CheckConfigOptions)
				checkConfigOptionsModel.ProjectID = core.StringPtr("testString")
				checkConfigOptionsModel.ID = core.StringPtr("testString")
				checkConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				checkConfigOptionsModel.Version = core.StringPtr("active")
				checkConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.CheckConfigWithContext(ctx, checkConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.CheckConfig(checkConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.CheckConfigWithContext(ctx, checkConfigOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(checkConfigPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.Header["X-Auth-Refresh-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Refresh-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"active"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke CheckConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.CheckConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CheckConfigOptions model
				checkConfigOptionsModel := new(projectv1.CheckConfigOptions)
				checkConfigOptionsModel.ProjectID = core.StringPtr("testString")
				checkConfigOptionsModel.ID = core.StringPtr("testString")
				checkConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				checkConfigOptionsModel.Version = core.StringPtr("active")
				checkConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.CheckConfig(checkConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CheckConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the CheckConfigOptions model
				checkConfigOptionsModel := new(projectv1.CheckConfigOptions)
				checkConfigOptionsModel.ProjectID = core.StringPtr("testString")
				checkConfigOptionsModel.ID = core.StringPtr("testString")
				checkConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				checkConfigOptionsModel.Version = core.StringPtr("active")
				checkConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.CheckConfig(checkConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CheckConfigOptions model with no property values
				checkConfigOptionsModelNew := new(projectv1.CheckConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.CheckConfig(checkConfigOptionsModelNew)
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
			It(`Invoke CheckConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the CheckConfigOptions model
				checkConfigOptionsModel := new(projectv1.CheckConfigOptions)
				checkConfigOptionsModel.ProjectID = core.StringPtr("testString")
				checkConfigOptionsModel.ID = core.StringPtr("testString")
				checkConfigOptionsModel.XAuthRefreshToken = core.StringPtr("testString")
				checkConfigOptionsModel.Version = core.StringPtr("active")
				checkConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.CheckConfig(checkConfigOptionsModel)
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
	Describe(`InstallConfig(installConfigOptions *InstallConfigOptions) - Operation response error`, func() {
		installConfigPath := "/v1/projects/testString/configs/testString/install"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(installConfigPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke InstallConfig with error: Operation response processing error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the InstallConfigOptions model
				installConfigOptionsModel := new(projectv1.InstallConfigOptions)
				installConfigOptionsModel.ProjectID = core.StringPtr("testString")
				installConfigOptionsModel.ID = core.StringPtr("testString")
				installConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := projectService.InstallConfig(installConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				projectService.EnableRetries(0, 0)
				result, response, operationErr = projectService.InstallConfig(installConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`InstallConfig(installConfigOptions *InstallConfigOptions)`, func() {
		installConfigPath := "/v1/projects/testString/configs/testString/install"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(installConfigPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke InstallConfig successfully with retries`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())
				projectService.EnableRetries(0, 0)

				// Construct an instance of the InstallConfigOptions model
				installConfigOptionsModel := new(projectv1.InstallConfigOptions)
				installConfigOptionsModel.ProjectID = core.StringPtr("testString")
				installConfigOptionsModel.ID = core.StringPtr("testString")
				installConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := projectService.InstallConfigWithContext(ctx, installConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				projectService.DisableRetries()
				result, response, operationErr := projectService.InstallConfig(installConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = projectService.InstallConfigWithContext(ctx, installConfigOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(installConfigPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "labels": ["Labels"], "description": "Description", "authorizations": {"trusted_profile": {"id": "ID", "target_iam_id": "TargetIamID"}, "method": "Method", "api_key": "ApiKey"}, "compliance_profile": {"id": "ID", "instance_id": "InstanceID", "instance_location": "InstanceLocation", "attachment_id": "AttachmentID", "profile_name": "ProfileName"}, "locator_id": "LocatorID", "type": "terraform_template", "input": [{"name": "Name", "type": "array", "value": "anyValue", "required": true}], "output": [{"name": "Name", "description": "Description", "value": "anyValue"}], "setting": [{"name": "Name", "value": "Value"}]}`)
				}))
			})
			It(`Invoke InstallConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := projectService.InstallConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the InstallConfigOptions model
				installConfigOptionsModel := new(projectv1.InstallConfigOptions)
				installConfigOptionsModel.ProjectID = core.StringPtr("testString")
				installConfigOptionsModel.ID = core.StringPtr("testString")
				installConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = projectService.InstallConfig(installConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke InstallConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the InstallConfigOptions model
				installConfigOptionsModel := new(projectv1.InstallConfigOptions)
				installConfigOptionsModel.ProjectID = core.StringPtr("testString")
				installConfigOptionsModel.ID = core.StringPtr("testString")
				installConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := projectService.InstallConfig(installConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the InstallConfigOptions model with no property values
				installConfigOptionsModelNew := new(projectv1.InstallConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = projectService.InstallConfig(installConfigOptionsModelNew)
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
			It(`Invoke InstallConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the InstallConfigOptions model
				installConfigOptionsModel := new(projectv1.InstallConfigOptions)
				installConfigOptionsModel.ProjectID = core.StringPtr("testString")
				installConfigOptionsModel.ID = core.StringPtr("testString")
				installConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := projectService.InstallConfig(installConfigOptionsModel)
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
	Describe(`UninstallConfig(uninstallConfigOptions *UninstallConfigOptions)`, func() {
		uninstallConfigPath := "/v1/projects/testString/configs/testString/uninstall"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(uninstallConfigPath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke UninstallConfig successfully`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := projectService.UninstallConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UninstallConfigOptions model
				uninstallConfigOptionsModel := new(projectv1.UninstallConfigOptions)
				uninstallConfigOptionsModel.ProjectID = core.StringPtr("testString")
				uninstallConfigOptionsModel.ID = core.StringPtr("testString")
				uninstallConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = projectService.UninstallConfig(uninstallConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UninstallConfig with error: Operation validation and request error`, func() {
				projectService, serviceErr := projectv1.NewProjectV1(&projectv1.ProjectV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(projectService).ToNot(BeNil())

				// Construct an instance of the UninstallConfigOptions model
				uninstallConfigOptionsModel := new(projectv1.UninstallConfigOptions)
				uninstallConfigOptionsModel.ProjectID = core.StringPtr("testString")
				uninstallConfigOptionsModel.ID = core.StringPtr("testString")
				uninstallConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := projectService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := projectService.UninstallConfig(uninstallConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UninstallConfigOptions model with no property values
				uninstallConfigOptionsModelNew := new(projectv1.UninstallConfigOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = projectService.UninstallConfig(uninstallConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
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
			It(`Invoke NewCheckConfigOptions successfully`, func() {
				// Construct an instance of the CheckConfigOptions model
				projectID := "testString"
				id := "testString"
				checkConfigOptionsModel := projectService.NewCheckConfigOptions(projectID, id)
				checkConfigOptionsModel.SetProjectID("testString")
				checkConfigOptionsModel.SetID("testString")
				checkConfigOptionsModel.SetXAuthRefreshToken("testString")
				checkConfigOptionsModel.SetVersion("active")
				checkConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(checkConfigOptionsModel).ToNot(BeNil())
				Expect(checkConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(checkConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(checkConfigOptionsModel.XAuthRefreshToken).To(Equal(core.StringPtr("testString")))
				Expect(checkConfigOptionsModel.Version).To(Equal(core.StringPtr("active")))
				Expect(checkConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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

				// Construct an instance of the ProjectConfigComplianceProfile model
				projectConfigComplianceProfileModel := new(projectv1.ProjectConfigComplianceProfile)
				Expect(projectConfigComplianceProfileModel).ToNot(BeNil())
				projectConfigComplianceProfileModel.ID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectConfigComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.ProfileName = core.StringPtr("testString")
				Expect(projectConfigComplianceProfileModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigComplianceProfileModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigComplianceProfileModel.InstanceLocation).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigComplianceProfileModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigComplianceProfileModel.ProfileName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				Expect(projectConfigInputVariableModel).ToNot(BeNil())
				projectConfigInputVariableModel.Name = core.StringPtr("account_id")
				projectConfigInputVariableModel.Value = core.StringPtr(`$configs[].name["account-stage"].input.account_id`)
				Expect(projectConfigInputVariableModel.Name).To(Equal(core.StringPtr("account_id")))
				Expect(projectConfigInputVariableModel.Value).To(Equal(core.StringPtr(`$configs[].name["account-stage"].input.account_id`)))

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				Expect(projectConfigSettingCollectionModel).ToNot(BeNil())
				projectConfigSettingCollectionModel.Name = core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")
				projectConfigSettingCollectionModel.Value = core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")
				Expect(projectConfigSettingCollectionModel.Name).To(Equal(core.StringPtr("IBMCLOUD_TOOLCHAIN_ENDPOINT")))
				Expect(projectConfigSettingCollectionModel.Value).To(Equal(core.StringPtr("https://api.us-south.devops.dev.cloud.ibm.com")))

				// Construct an instance of the CreateConfigOptions model
				projectID := "testString"
				createConfigOptionsName := "env-stage"
				createConfigOptionsLocatorID := "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"
				createConfigOptionsModel := projectService.NewCreateConfigOptions(projectID, createConfigOptionsName, createConfigOptionsLocatorID)
				createConfigOptionsModel.SetProjectID("testString")
				createConfigOptionsModel.SetName("env-stage")
				createConfigOptionsModel.SetLocatorID("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				createConfigOptionsModel.SetID("testString")
				createConfigOptionsModel.SetLabels([]string{"env:stage", "governance:test", "build:0"})
				createConfigOptionsModel.SetDescription("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")
				createConfigOptionsModel.SetAuthorizations(projectConfigAuthModel)
				createConfigOptionsModel.SetComplianceProfile(projectConfigComplianceProfileModel)
				createConfigOptionsModel.SetInput([]projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel})
				createConfigOptionsModel.SetSetting([]projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel})
				createConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createConfigOptionsModel).ToNot(BeNil())
				Expect(createConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(createConfigOptionsModel.Name).To(Equal(core.StringPtr("env-stage")))
				Expect(createConfigOptionsModel.LocatorID).To(Equal(core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")))
				Expect(createConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createConfigOptionsModel.Labels).To(Equal([]string{"env:stage", "governance:test", "build:0"}))
				Expect(createConfigOptionsModel.Description).To(Equal(core.StringPtr("Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace.")))
				Expect(createConfigOptionsModel.Authorizations).To(Equal(projectConfigAuthModel))
				Expect(createConfigOptionsModel.ComplianceProfile).To(Equal(projectConfigComplianceProfileModel))
				Expect(createConfigOptionsModel.Input).To(Equal([]projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}))
				Expect(createConfigOptionsModel.Setting).To(Equal([]projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}))
				Expect(createConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProjectOptions successfully`, func() {
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

				// Construct an instance of the ProjectConfigComplianceProfile model
				projectConfigComplianceProfileModel := new(projectv1.ProjectConfigComplianceProfile)
				Expect(projectConfigComplianceProfileModel).ToNot(BeNil())
				projectConfigComplianceProfileModel.ID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.InstanceLocation = core.StringPtr("testString")
				projectConfigComplianceProfileModel.AttachmentID = core.StringPtr("testString")
				projectConfigComplianceProfileModel.ProfileName = core.StringPtr("testString")
				Expect(projectConfigComplianceProfileModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigComplianceProfileModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigComplianceProfileModel.InstanceLocation).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigComplianceProfileModel.AttachmentID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigComplianceProfileModel.ProfileName).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigInputVariable model
				projectConfigInputVariableModel := new(projectv1.ProjectConfigInputVariable)
				Expect(projectConfigInputVariableModel).ToNot(BeNil())
				projectConfigInputVariableModel.Name = core.StringPtr("testString")
				projectConfigInputVariableModel.Value = core.StringPtr("testString")
				Expect(projectConfigInputVariableModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigInputVariableModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigSettingCollection model
				projectConfigSettingCollectionModel := new(projectv1.ProjectConfigSettingCollection)
				Expect(projectConfigSettingCollectionModel).ToNot(BeNil())
				projectConfigSettingCollectionModel.Name = core.StringPtr("testString")
				projectConfigSettingCollectionModel.Value = core.StringPtr("testString")
				Expect(projectConfigSettingCollectionModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigSettingCollectionModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ProjectConfigPrototype model
				projectConfigPrototypeModel := new(projectv1.ProjectConfigPrototype)
				Expect(projectConfigPrototypeModel).ToNot(BeNil())
				projectConfigPrototypeModel.ID = core.StringPtr("testString")
				projectConfigPrototypeModel.Name = core.StringPtr("common-variables")
				projectConfigPrototypeModel.Labels = []string{"testString"}
				projectConfigPrototypeModel.Description = core.StringPtr("testString")
				projectConfigPrototypeModel.Authorizations = projectConfigAuthModel
				projectConfigPrototypeModel.ComplianceProfile = projectConfigComplianceProfileModel
				projectConfigPrototypeModel.LocatorID = core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")
				projectConfigPrototypeModel.Input = []projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}
				projectConfigPrototypeModel.Setting = []projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}
				Expect(projectConfigPrototypeModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigPrototypeModel.Name).To(Equal(core.StringPtr("common-variables")))
				Expect(projectConfigPrototypeModel.Labels).To(Equal([]string{"testString"}))
				Expect(projectConfigPrototypeModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigPrototypeModel.Authorizations).To(Equal(projectConfigAuthModel))
				Expect(projectConfigPrototypeModel.ComplianceProfile).To(Equal(projectConfigComplianceProfileModel))
				Expect(projectConfigPrototypeModel.LocatorID).To(Equal(core.StringPtr("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global")))
				Expect(projectConfigPrototypeModel.Input).To(Equal([]projectv1.ProjectConfigInputVariable{*projectConfigInputVariableModel}))
				Expect(projectConfigPrototypeModel.Setting).To(Equal([]projectv1.ProjectConfigSettingCollection{*projectConfigSettingCollectionModel}))

				// Construct an instance of the CreateProjectOptions model
				resourceGroup := "Default"
				location := "us-south"
				createProjectOptionsName := "acme-microservice"
				createProjectOptionsModel := projectService.NewCreateProjectOptions(resourceGroup, location, createProjectOptionsName)
				createProjectOptionsModel.SetResourceGroup("Default")
				createProjectOptionsModel.SetLocation("us-south")
				createProjectOptionsModel.SetName("acme-microservice")
				createProjectOptionsModel.SetDescription("A microservice to deploy on top of ACME infrastructure.")
				createProjectOptionsModel.SetDestroyOnDelete(true)
				createProjectOptionsModel.SetConfigs([]projectv1.ProjectConfigPrototype{*projectConfigPrototypeModel})
				createProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProjectOptionsModel).ToNot(BeNil())
				Expect(createProjectOptionsModel.ResourceGroup).To(Equal(core.StringPtr("Default")))
				Expect(createProjectOptionsModel.Location).To(Equal(core.StringPtr("us-south")))
				Expect(createProjectOptionsModel.Name).To(Equal(core.StringPtr("acme-microservice")))
				Expect(createProjectOptionsModel.Description).To(Equal(core.StringPtr("A microservice to deploy on top of ACME infrastructure.")))
				Expect(createProjectOptionsModel.DestroyOnDelete).To(Equal(core.BoolPtr(true)))
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
				deleteConfigOptionsModel.SetDraftOnly(false)
				deleteConfigOptionsModel.SetDestroy(true)
				deleteConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteConfigOptionsModel).ToNot(BeNil())
				Expect(deleteConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteConfigOptionsModel.DraftOnly).To(Equal(core.BoolPtr(false)))
				Expect(deleteConfigOptionsModel.Destroy).To(Equal(core.BoolPtr(true)))
				Expect(deleteConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProjectOptions successfully`, func() {
				// Construct an instance of the DeleteProjectOptions model
				id := "testString"
				deleteProjectOptionsModel := projectService.NewDeleteProjectOptions(id)
				deleteProjectOptionsModel.SetID("testString")
				deleteProjectOptionsModel.SetDestroy(true)
				deleteProjectOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProjectOptionsModel).ToNot(BeNil())
				Expect(deleteProjectOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProjectOptionsModel.Destroy).To(Equal(core.BoolPtr(true)))
				Expect(deleteProjectOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConfigOptions successfully`, func() {
				// Construct an instance of the GetConfigOptions model
				projectID := "testString"
				id := "testString"
				getConfigOptionsModel := projectService.NewGetConfigOptions(projectID, id)
				getConfigOptionsModel.SetProjectID("testString")
				getConfigOptionsModel.SetID("testString")
				getConfigOptionsModel.SetVersion("active")
				getConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConfigOptionsModel).ToNot(BeNil())
				Expect(getConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(getConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getConfigOptionsModel.Version).To(Equal(core.StringPtr("active")))
				Expect(getConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewInstallConfigOptions successfully`, func() {
				// Construct an instance of the InstallConfigOptions model
				projectID := "testString"
				id := "testString"
				installConfigOptionsModel := projectService.NewInstallConfigOptions(projectID, id)
				installConfigOptionsModel.SetProjectID("testString")
				installConfigOptionsModel.SetID("testString")
				installConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(installConfigOptionsModel).ToNot(BeNil())
				Expect(installConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(installConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(installConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListConfigsOptions successfully`, func() {
				// Construct an instance of the ListConfigsOptions model
				projectID := "testString"
				listConfigsOptionsModel := projectService.NewListConfigsOptions(projectID)
				listConfigsOptionsModel.SetProjectID("testString")
				listConfigsOptionsModel.SetVersion("active")
				listConfigsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listConfigsOptionsModel).ToNot(BeNil())
				Expect(listConfigsOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(listConfigsOptionsModel.Version).To(Equal(core.StringPtr("active")))
				Expect(listConfigsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewProjectConfigInputVariable successfully`, func() {
				name := "testString"
				_model, err := projectService.NewProjectConfigInputVariable(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProjectConfigPrototype successfully`, func() {
				name := "testString"
				locatorID := "testString"
				_model, err := projectService.NewProjectConfigPrototype(name, locatorID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProjectConfigSettingCollection successfully`, func() {
				name := "testString"
				value := "testString"
				_model, err := projectService.NewProjectConfigSettingCollection(name, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUninstallConfigOptions successfully`, func() {
				// Construct an instance of the UninstallConfigOptions model
				projectID := "testString"
				id := "testString"
				uninstallConfigOptionsModel := projectService.NewUninstallConfigOptions(projectID, id)
				uninstallConfigOptionsModel.SetProjectID("testString")
				uninstallConfigOptionsModel.SetID("testString")
				uninstallConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uninstallConfigOptionsModel).ToNot(BeNil())
				Expect(uninstallConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(uninstallConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(uninstallConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateConfigOptions successfully`, func() {
				// Construct an instance of the ProjectConfigPatchRequestProjectConfigManualProperty model
				projectConfigPatchRequestModel := new(projectv1.ProjectConfigPatchRequestProjectConfigManualProperty)
				Expect(projectConfigPatchRequestModel).ToNot(BeNil())
				projectConfigPatchRequestModel.Name = core.StringPtr("testString")
				projectConfigPatchRequestModel.Labels = []string{"testString"}
				projectConfigPatchRequestModel.Description = core.StringPtr("testString")
				projectConfigPatchRequestModel.Type = core.StringPtr("manual")
				projectConfigPatchRequestModel.ExternalResourcesAccount = core.StringPtr("testString")
				Expect(projectConfigPatchRequestModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigPatchRequestModel.Labels).To(Equal([]string{"testString"}))
				Expect(projectConfigPatchRequestModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(projectConfigPatchRequestModel.Type).To(Equal(core.StringPtr("manual")))
				Expect(projectConfigPatchRequestModel.ExternalResourcesAccount).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateConfigOptions model
				projectID := "testString"
				id := "testString"
				var projectConfig projectv1.ProjectConfigPatchRequestIntf = nil
				updateConfigOptionsModel := projectService.NewUpdateConfigOptions(projectID, id, projectConfig)
				updateConfigOptionsModel.SetProjectID("testString")
				updateConfigOptionsModel.SetID("testString")
				updateConfigOptionsModel.SetProjectConfig(projectConfigPatchRequestModel)
				updateConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateConfigOptionsModel).ToNot(BeNil())
				Expect(updateConfigOptionsModel.ProjectID).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigOptionsModel.ProjectConfig).To(Equal(projectConfigPatchRequestModel))
				Expect(updateConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewProjectConfigPatchRequestProjectConfigManualProperty successfully`, func() {
				typeVar := "manual"
				_model, err := projectService.NewProjectConfigPatchRequestProjectConfigManualProperty(typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
