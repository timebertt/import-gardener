package main

import (
	gardenerv1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
)

func main() {
	useGardenerAPIs()
}

func useGardenerAPIs() {
	_ = &gardenerv1beta1.Shoot{}
	_ = &extensionsv1alpha1.Infrastructure{}
}
