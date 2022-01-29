package main

import (
	authenticationv1alpha1 "github.com/gardener/gardener/pkg/apis/authentication/v1alpha1"
	gardenerv1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	operationsv1alpha1 "github.com/gardener/gardener/pkg/apis/operations/v1alpha1"
	resourcesv1alpha1 "github.com/gardener/gardener/pkg/apis/resources/v1alpha1"
	settingsv1alpha1 "github.com/gardener/gardener/pkg/apis/settings/v1alpha1"
)

func main() {
	useGardenerAPIs()
}

func useGardenerAPIs() {
	_ = &gardenerv1beta1.Shoot{}
	_ = &extensionsv1alpha1.Infrastructure{}
	_ = &authenticationv1alpha1.AdminKubeconfigRequest{}
	_ = &operationsv1alpha1.Bastion{}
	_ = &resourcesv1alpha1.ManagedResource{}
	_ = &settingsv1alpha1.OpenIDConnectPreset{}
}
