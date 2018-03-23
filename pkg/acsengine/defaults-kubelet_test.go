package acsengine

import (
	"testing"

	"github.com/Azure/acs-engine/pkg/helpers"
)

func TestKubeletConfigUseCloudControllerManager(t *testing.T) {
	// Test UseCloudControllerManager = true
	cs := createContainerService("testcluster", defaultTestClusterVer, 3, 2)
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager = helpers.PointerToBool(true)
	setKubeletConfig(cs)
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--cloud-provider"] != "external" {
		t.Fatalf("got unexpected '--cloud-provider' kubelet config value for UseCloudControllerManager=true: %s",
			k["--cloud-provider"])
	}

	// Test UseCloudControllerManager = false
	cs = createContainerService("testcluster", defaultTestClusterVer, 3, 2)
	cs.Properties.OrchestratorProfile.KubernetesConfig.UseCloudControllerManager = helpers.PointerToBool(false)
	setKubeletConfig(cs)
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--cloud-provider"] != "azure" {
		t.Fatalf("got unexpected '--cloud-provider' kubelet config value for UseCloudControllerManager=false: %s",
			k["--cloud-provider"])
	}

}

func TestKubeletConfigNetworkPolicy(t *testing.T) {
	// Test NetworkPolicy = none
	cs := createContainerService("testcluster", defaultTestClusterVer, 3, 2)
	cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy = NetworkPolicyNone
	setKubeletConfig(cs)
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--network-plugin"] != NetworkPluginKubenet {
		t.Fatalf("got unexpected '--network-plugin' kubelet config value for NetworkPolicy=none: %s",
			k["--network-plugin"])
	}

	// Test NetworkPolicy = azure
	cs = createContainerService("testcluster", defaultTestClusterVer, 3, 2)
	cs.Properties.OrchestratorProfile.KubernetesConfig.NetworkPolicy = "azure"
	setKubeletConfig(cs)
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--network-plugin"] != "cni" {
		t.Fatalf("got unexpected '--network-plugin' kubelet config value for NetworkPolicy=azure: %s",
			k["--network-plugin"])
	}

}

func TestKubeletConfigEnableSecureKubelet(t *testing.T) {
	// Test EnableSecureKubelet = true
	cs := createContainerService("testcluster", defaultTestClusterVer, 3, 2)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet = helpers.PointerToBool(true)
	setKubeletConfig(cs)
	k := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	if k["--anonymous-auth"] != "false" {
		t.Fatalf("got unexpected '--anonymous-auth' kubelet config value for EnableSecureKubelet=true: %s",
			k["--anonymous-auth"])
	}
	if k["--authorization-mode"] != "Webhook" {
		t.Fatalf("got unexpected '--authorization-mode' kubelet config value for EnableSecureKubelet=true: %s",
			k["--authorization-mode"])
	}
	if k["--client-ca-file"] != "/etc/kubernetes/certs/ca.crt" {
		t.Fatalf("got unexpected '--client-ca-file' kubelet config value for EnableSecureKubelet=true: %s",
			k["--client-ca-file"])
	}

	// Test EnableSecureKubelet = false
	cs = createContainerService("testcluster", defaultTestClusterVer, 3, 2)
	cs.Properties.OrchestratorProfile.KubernetesConfig.EnableSecureKubelet = helpers.PointerToBool(false)
	setKubeletConfig(cs)
	k = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig
	for _, key := range []string{"--anonymous-auth", "--client-ca-file"} {
		if _, ok := k[key]; ok {
			t.Fatalf("got unexpected '%s' kubelet config value for EnableSecureKubelet=false: %s",
				key, k[key])
		}
	}

}

func TestKubeletConfigDefaults(t *testing.T) {
	clusterVer := "1.7.12"
	cs := createContainerService("testcluster", clusterVer, 3, 2)
	setKubeletConfig(cs)
	val, ok := cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig["--containerized"]
	if val != "" && !ok {
		t.Fatalf("expected --containerized kubelet option for Kubernetes v%s", clusterVer)
	}
	clusterVer = "1.8.10"
	cs = createContainerService("testcluster", clusterVer, 3, 2)
	setKubeletConfig(cs)
	val, ok = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig["--containerized"]
	if val != "" && !ok {
		t.Fatalf("expected --containerized kubelet option for Kubernetes v%s", clusterVer)
	}
	clusterVer = "1.9.6"
	cs = createContainerService("testcluster", clusterVer, 3, 2)
	setKubeletConfig(cs)
	val, ok = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig["--containerized"]
	if val != "" && !ok {
		t.Fatalf("expected --containerized kubelet option for Kubernetes v%s", clusterVer)
	}
	clusterVer = "1.10.0-rc.1"
	cs = createContainerService("testcluster", clusterVer, 3, 2)
	setKubeletConfig(cs)
	val, ok = cs.Properties.OrchestratorProfile.KubernetesConfig.KubeletConfig["--containerized"]
	if val != "" && !ok {
		t.Fatalf("expected --containerized kubelet option for Kubernetes v%s", clusterVer)
	}
}
