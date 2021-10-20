package pkg

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"net/url"
)

type StorageAccountClient struct {
	accountName string
	accountKey  string
	credential  *azblob.SharedKeyCredential
	serviceUrl  *azblob.ServiceURL
}

func NewStorageAccountClient(accountName, accountKey string) (*StorageAccountClient, error) {
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return nil, errors.New("Invalid credentials with error: " + err.Error())
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net", accountName))
	serviceUrl := azblob.NewServiceURL(*URL, p)
	return &StorageAccountClient{
		accountName: accountName,
		accountKey: accountKey,
		credential: credential,
		serviceUrl: &serviceUrl,
	}, nil
}

func (client *StorageAccountClient) ListContainer() {
	//ctx := context.Background()
	//client.serviceUrl.ListContainersSegment(ctx, )
}
