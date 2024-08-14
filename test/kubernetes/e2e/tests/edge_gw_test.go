package tests_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/solo-io/gloo/pkg/utils/env"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	. "github.com/solo-io/gloo/test/kubernetes/e2e/tests"
	"github.com/solo-io/gloo/test/kubernetes/testutils/gloogateway"
	"github.com/solo-io/gloo/test/kubernetes/testutils/helper"
	"github.com/solo-io/gloo/test/testutils"
	"github.com/solo-io/skv2/codegen/util"
)

// TestGlooGatewayEdgeGateway is the function which executes a series of tests against a given installation where
// the k8s Gateway controller is disabled
func TestGlooGatewayEdgeGateway(t *testing.T) {
	ctx := context.Background()
	installNs := env.GetOrDefault(testutils.InstallNamespace, "gloo-gateway-edge-test")
	testInstallation := e2e.CreateTestInstallation(
		t,
		&gloogateway.Context{
			InstallNamespace:   installNs,
			ValuesManifestFile: filepath.Join(util.MustGetThisDir(), "manifests", "edge-gateway-test-helm.yaml"),
		},
	)

	testHelper := e2e.MustTestHelper(ctx, testInstallation)
	// Set the env to the install namespace if it is not already set
	if os.Getenv(testutils.InstallNamespace) == "" {
		os.Setenv(testutils.InstallNamespace, installNs)
	}

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

	// Install Gloo Gateway with only Gloo Edge Gateway APIs enabled
	testInstallation.InstallGlooGateway(ctx, func(ctx context.Context) error {
		return testHelper.InstallGloo(ctx, 5*time.Minute, helper.WithExtraArgs("--values", testInstallation.Metadata.ValuesManifestFile))
	})

	EdgeGwSuiteRunner().Run(ctx, t, testInstallation)
}
