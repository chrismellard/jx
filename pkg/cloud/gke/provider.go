package gke

import (
	"k8s.io/client-go/kubernetes"
)

type GKEProvider struct {
	gCloud GCloud
}

func (p *GKEProvider) CreateServiceAccount(kubeClient kubernetes.Interface, serviceName, serviceAbbreviation, namespace, clusterName, projectID string, serviceAccountRoles []string, serviceAccountSecretKey string) (string, error) {
	return p.gCloud.CreateGCPServiceAccount(kubeClient, serviceName, serviceAbbreviation, namespace, clusterName, projectID, serviceAccountRoles, serviceAccountSecretKey)
}

func (p *GKEProvider) GenerateServiceAccountSecretName(releaseName string) string {
	return GcpServiceAccountSecretName(releaseName)
}

func (p *GKEProvider) EnableAPIs(projectID string, apis ...string) error {
	return p.gCloud.EnableAPIs(projectID, apis...)
}
