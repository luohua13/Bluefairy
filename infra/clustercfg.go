package infra

import (
	"context"
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"os"
	"strings"
)

var (
	kubeconfig string
)

func init() {
	// TODO: Fix this to allow double vendoring this library but still register flags on behalf of users
	flag.StringVar(&kubeconfig, "kubeconfig", "",
		"Paths to a kubeconfig. Only required if out-of-cluster.")
}

const (
	MultiClusterSecretLabel = "istio/multiCluster"
	RemoteSecretPrefix      = "istio-remote-secret-"
)

type ClusterID string

func ClusterNameFromRemoteSecretName(name string) ClusterID {
	return ClusterID(strings.TrimPrefix(name, RemoteSecretPrefix))
}

func GetKubernetesConfig() (cfgArray []*rest.Config, err error) {
	cfg, err := getConfig()
	if err != nil {
		fmt.Println("111111")
		return nil, err
	}
	cfgArray = append(cfgArray, cfg)
	localClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	req, err := labels.NewRequirement(MultiClusterSecretLabel, selection.Equals, []string{"true"})
	if err != nil {
		return nil, err
	}
	labelSelector := labels.NewSelector().Add(*req)

	secretList, err := localClient.CoreV1().Secrets("istio-system").
		List(context.TODO(), metav1.ListOptions{LabelSelector: labelSelector.String()})
	if err != nil {
		return nil, err
	}

	for _, secret := range secretList.Items {
		restCfg, err := restConfigFromRemoteSecret(&secret)
		if err != nil {
			return nil, err
		}
		cfgArray = append(cfgArray, restCfg)
	}

	return
}

func restConfigFromRemoteSecret(secret *corev1.Secret) (*rest.Config, error) {
	remoteClusterName := ClusterNameFromRemoteSecretName(secret.GetName())
	cfg := secret.Data[string(remoteClusterName)]
	if cfg == nil {
		return nil, fmt.Errorf("kubeconfig is nil in secret data ")
	}

	restCfg, err := clientcmd.RESTConfigFromKubeConfig(cfg)
	if err != nil {
		return nil, err
	}

	return restCfg, nil
}

func getConfig() (*rest.Config, error) {
	if len(kubeconfig) > 0 {
		return loadConfigWithContext("", &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig}, "")
	}

	cfg, err := rest.InClusterConfig()
	if err != nil {
		homeDir, _ := os.UserHomeDir()
		cfg, err = clientcmd.BuildConfigFromFlags("", homeDir+"/.kube/config")
		if err != nil {
			return nil, err
		}
	}
	return cfg, err
}

func loadConfigWithContext(apiServerURL string, loader clientcmd.ClientConfigLoader, context string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loader,
		&clientcmd.ConfigOverrides{
			ClusterInfo: clientcmdapi.Cluster{
				Server: apiServerURL,
			},
			CurrentContext: context,
		}).ClientConfig()
}
