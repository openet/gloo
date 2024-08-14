package tests_test

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/solo-io/skv2/codegen/util"

	"github.com/solo-io/gloo/test/kubernetes/e2e"
	. "github.com/solo-io/gloo/test/kubernetes/e2e/tests"
	"github.com/solo-io/gloo/test/kubernetes/testutils/gloogateway"
	"github.com/solo-io/gloo/test/kubernetes/testutils/helper"
)

// TestK8sGatewayNoValidation executes tests against a K8s Gateway gloo install with validation disabled
func TestK8sGatewayNoValidation(t *testing.T) {
	ctx := context.Background()
	testInstallation := e2e.CreateTestInstallation(
		t,
		&gloogateway.Context{
			InstallNamespace:       "k8s-gw-test-no-validation",
			ValuesManifestFile:     filepath.Join(util.MustGetThisDir(), "manifests", "k8s-gateway-no-webhook-validation-test-helm.yaml"),
			ValidationAlwaysAccept: true,
			K8sGatewayEnabled:      true,
		},
	)

	testHelper := e2e.MustTestHelper(ctx, testInstallation)

	// We register the cleanup function _before_ we actually perform the installation.
	// This allows us to uninstall Gloo Gateway, in case the original installation only completed partially
	t.Cleanup(func() {
		if t.Failed() {
			testInstallation.PreFailHandler(ctx)
		}

		testInstallation.UninstallGlooGateway(ctx, func(ctx context.Context) error {
			return testHelper.UninstallGlooAll()
		})
	})

	// Install Gloo Gateway
	testInstallation.InstallGlooGateway(ctx, func(ctx context.Context) error {
		return testHelper.InstallGloo(ctx, 5*time.Minute, helper.WithExtraArgs("--values", testInstallation.Metadata.ValuesManifestFile))
	})

	KubeGatewayNoValidationSuiteRunner().Run(ctx, t, testInstallation)
}
