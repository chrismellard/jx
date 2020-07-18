package storage

import (
	"fmt"
	"github.com/jenkins-x/jx-logging/pkg/log"
	"github.com/jenkins-x/jx/v2/pkg/cloud/aks"
	"github.com/jenkins-x/jx/v2/pkg/cloud/buckets"
	"github.com/jenkins-x/jx/v2/pkg/config"
	"github.com/jenkins-x/jx/v2/pkg/util"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"io"
	"strings"
)

type AKSBucketProvider struct {
	Requirements *config.RequirementsConfig
	AzureStorage aks.AzureStorage
}

func (b *AKSBucketProvider) CreateNewBucketForCluster(clusterName string, bucketKind string) (string, error) {
	uuid4, _ := uuid.NewV4()
	bucketName := fmt.Sprintf("%s-%s-%s", clusterName, bucketKind, uuid4.String())

	// Max length is 63, https://docs.microsoft.com/en-us/rest/api/storageservices/naming-and-referencing-containers--blobs--and-metadata
	if len(bucketName) > 63 {
		bucketName = bucketName[:63]
	}
	bucketName = strings.TrimRight(bucketName, "-")
	bucketURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s", b.Requirements.Velero.ServiceAccount, bucketName)
	err := b.EnsureBucketIsCreated(bucketURL)
	if err != nil {
		return bucketURL, errors.Wrapf(err, "failed to create bucket %s", bucketURL)
	}

	return bucketURL, nil
}

func (b *AKSBucketProvider) EnsureBucketIsCreated(bucketURL string) error {

	exists, err := b.AzureStorage.ContainerExists(bucketURL)
	if err != nil {
		return errors.Wrap(err, "checking if the provided container exists")
	}
	if exists {
		return nil
	}

	log.Logger().Infof("The bucket %s does not exist so lets create it", util.ColorInfo(bucketURL))
	err = b.AzureStorage.CreateContainer(bucketURL)
	if err != nil {
		return errors.Wrapf(err, "there was a problem creating the bucket with URL %s",
			bucketURL)
	}
	return nil
}

func (b *AKSBucketProvider) UploadFileToBucket(r io.Reader, outputName string, bucketURL string) (string, error) {
	return "", nil
}

func (b *AKSBucketProvider) DownloadFileFromBucket(bucketURL string) (io.ReadCloser, error) {
	return closerReader{}, nil
}

func NewAKSBucketProvider(requirements *config.RequirementsConfig) buckets.Provider {
	return &AKSBucketProvider{
		Requirements: requirements,
		AzureStorage: aks.NewAzureRunner(),
	}
}

type closerReader struct {
}

func (closerReader) Close() error {
	return nil
}

func (closerReader) Read(p []byte) (n int, err error) {
	return 0, nil
}
