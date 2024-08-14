package validation_strict_warnings

import (
	"context"
	"fmt"
	"time"

	gloo_defaults "github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	testdefaults "github.com/solo-io/gloo/test/kubernetes/e2e/defaults"
	"github.com/solo-io/gloo/test/kubernetes/e2e/features/validation"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/stretchr/testify/suite"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ e2e.NewSuiteFunc = NewTestingSuite

// testingSuite is the entire Suite of tests for the webhook validation of strict validation feature (alwaysAccept=false, allowWarnings=false)
type testingSuite struct {
	suite.Suite

	ctx context.Context

	// testInstallation contains all the metadata/utilities necessary to execute a series of tests
	// against an installation of Gloo Gateway
	testInstallation *e2e.TestInstallation
}

func NewTestingSuite(ctx context.Context, testInst *e2e.TestInstallation) suite.TestingSuite {
	return &testingSuite{
		ctx:              ctx,
		testInstallation: testInst,
	}
}

// TestInvalidUpstreamMissingPort tests behaviors when Gloo rejects an invalid upstream with a missing port
func (s *testingSuite) TestInvalidUpstreamMissingPort() {
	s.T().Cleanup(func() {
		err := s.testInstallation.Actions.Kubectl().DeleteFileSafe(s.ctx, validation.ExampleVS, "-n", s.testInstallation.Metadata.InstallNamespace)
		s.Assert().NoError(err, "can delete "+validation.ExampleVS)

		// Delete can fail with strict validation if VS is not deleted first from snapshot, so try multiple times so that snapshot has time to update
		s.Assert().Eventually(func() bool {
			err := s.testInstallation.Actions.Kubectl().DeleteFileSafe(s.ctx, validation.ExampleUpstream, "-n", s.testInstallation.Metadata.InstallNamespace)
			return err == nil
		}, time.Minute, 5*time.Second, "can delete "+validation.ExampleUpstream)

		err = s.testInstallation.Actions.Kubectl().DeleteFileSafe(s.ctx, testdefaults.NginxPodManifest)
		s.Assert().NoError(err, "can delete "+testdefaults.NginxPodManifest)
	})

	err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, testdefaults.NginxPodManifest)
	s.Assert().NoError(err)
	// Check that test resources are running
	s.testInstallation.Assertions.EventuallyPodsRunning(s.ctx, testdefaults.NginxPod.ObjectMeta.GetNamespace(), metav1.ListOptions{
		LabelSelector: "app.kubernetes.io/name=nginx",
	})

	// Upstream is only rejected when the upstream plugin is run when a valid cluster is present
	err = s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, validation.ExampleUpstream, "-n", s.testInstallation.Metadata.InstallNamespace)
	s.Assert().NoError(err, "can apply valid upstream")
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
		func() (resources.InputResource, error) {
			return s.testInstallation.ResourceClients.UpstreamClient().Read(s.testInstallation.Metadata.InstallNamespace, validation.ExampleUpstreamName, clients.ReadOpts{Ctx: s.ctx})
		},
		core.Status_Accepted,
		gloo_defaults.GlooReporter,
	)
	err = s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, validation.ExampleVS, "-n", s.testInstallation.Metadata.InstallNamespace)
	s.Assert().NoError(err, "can apply valid virtual service")
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
		func() (resources.InputResource, error) {
			return s.testInstallation.ResourceClients.VirtualServiceClient().Read(s.testInstallation.Metadata.InstallNamespace, validation.ExampleVsName, clients.ReadOpts{Ctx: s.ctx})
		},
		core.Status_Accepted,
		gloo_defaults.GlooReporter,
	)

	output, err := s.testInstallation.Actions.Kubectl().ApplyFileWithOutput(s.ctx, validation.InvalidUpstreamNoPort, "-n", s.testInstallation.Metadata.InstallNamespace)
	s.Assert().Contains(output, fmt.Sprintf(`admission webhook "gloo.%s.svc" denied the request`, s.testInstallation.Metadata.InstallNamespace))
	s.Assert().Contains(output, `Validating *v1.Upstream failed: validating *v1.Upstream name:"invalid-us"`)
	s.Assert().Contains(output, "port cannot be empty for host")
}

// TestRejectsInvalidVSMissingUpstream tests behaviors when Gloo rejects invalid VirtualService resources due to missing upstream
func (s *testingSuite) TestRejectsInvalidVSMissingUpstream() {
	output, err := s.testInstallation.Actions.Kubectl().ApplyFileWithOutput(s.ctx, validation.InvalidVirtualMissingUpstream, "-n", s.testInstallation.Metadata.InstallNamespace)
	s.Assert().Error(err)
	s.Assert().Contains(output, fmt.Sprintf(`admission webhook "gloo.%s.svc" denied the request`, s.testInstallation.Metadata.InstallNamespace))
	s.Assert().Contains(output, `Validating *v1.VirtualService failed: validating *v1.VirtualService name:"no-upstream-vs"`)
	s.Assert().Contains(output, fmt.Sprintf(`Route Warning: InvalidDestinationWarning. Reason: *v1.Upstream { %s.does-not-exist } not found`, s.testInstallation.Metadata.InstallNamespace))
}
