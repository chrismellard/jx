package externaldns

import (
	"github.com/jenkins-x/jx/pkg/cloud"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
)

const (
	// ServiceAccountSecretKey is the key for the external dns service account secret
	serviceAccountSecretKey = "credentials.json"
	// DefaultExternalDNSAbbreviation appended to the GCP service account
	defaultExternalDNSAbbreviation = "dn"
)

var (
	serviceAccountRoles = []string{
		"roles/dns.admin",
	}
)

type GKEExternalDNS struct {
}

// CreateExternalDNSGCPServiceAccount creates a service account in GCP for ExternalDNS
func (s *GKEExternalDNS) CreateExternalDNSServiceAccount(gcloud cloud.CloudProvider, kubeClient kubernetes.Interface, externalDNSName, namespace, clusterName, projectID string) (string, error) {
	gcpServiceAccountSecretName, err := gcloud.CreateServiceAccount(kubeClient, externalDNSName, defaultExternalDNSAbbreviation, namespace, clusterName, projectID, serviceAccountRoles, serviceAccountSecretKey)
	if err != nil {
		return "", errors.Wrap(err, "creating the ExternalDNS GCP Service Account")
	}
	return gcpServiceAccountSecretName, nil
}

func (s *GKEExternalDNS) GetServiceAccountSecretKey() string {
	return serviceAccountSecretKey
}
