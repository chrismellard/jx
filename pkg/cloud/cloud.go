package cloud

import (
	"k8s.io/client-go/kubernetes"
	"sort"
	"strings"
)

const (
	GKE        = "gke"
	OKE        = "oke"
	EKS        = "eks"
	AKS        = "aks"
	AWS        = "aws"
	PKS        = "pks"
	IKS        = "iks"
	MINIKUBE   = "minikube"
	MINISHIFT  = "minishift"
	KUBERNETES = "kubernetes"
	OPENSHIFT  = "openshift"
	ICP        = "icp"
	JX_INFRA   = "jx-infra"
	ALIBABA    = "alibaba"
)

var BootEnabledProviders = map[string]bool{
	GKE: true,
}

// KubernetesProviders list of all available Kubernetes providers
var KubernetesProviders = []string{MINIKUBE, GKE, OKE, AKS, AWS, EKS, KUBERNETES, IKS, OPENSHIFT, MINISHIFT, JX_INFRA, PKS, ICP, ALIBABA}

// KubernetesProviderOptions returns all the Kubernetes providers as a string
func KubernetesProviderOptions() string {
	values := []string{}
	values = append(values, KubernetesProviders...)
	sort.Strings(values)
	return strings.Join(values, ", ")
}

type CloudProvider interface {
	CreateServiceAccount(kubeClient kubernetes.Interface, serviceName, serviceAbbreviation, namespace, clusterName, projectID string, serviceAccountRoles []string, serviceAccountSecretKey string) (string, error)
	GenerateServiceAccountSecretName(releaseName string) string
	EnableAPIs(projectID string, apis ...string) error
}

type ExternalDNS interface {
	CreateExternalDNSServiceAccount(gcloud CloudProvider, kubeClient kubernetes.Interface, externalDNSName, namespace, clusterName, projectID string) (string, error)
	GetServiceAccountSecretKey() string
}
