package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-storage-blob-go/azblob"
	"log"
	"net/url"
	"os"
)

type BlobClient struct {
	accountName string
	accountKey  string
	credential  *azblob.SharedKeyCredential
	container   *azblob.ContainerURL
}

func NewBlobClient(accountName, accountKey, containerName string) (*BlobClient, error) {
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return nil, errors.New("Invalid credentials with error: " + err.Error())
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))
	//Create a ContainerURL object that wraps the container URL and a request pipeline to make requests.
	containerURL := azblob.NewContainerURL(*URL, p)
	//ctx := context.Background() // This example uses a never-expiring context
	//_, err = containerURL.Create(ctx, azblob.Metadata{}, azblob.PublicAccessNone)
	return &BlobClient{
		accountName: accountName,
		accountKey:  accountKey,
		credential:  credential,
		container:   &containerURL,
	}, nil
}

func (bc *BlobClient) Upload(filePath, fileName string) {
	ctx := context.Background()
	blobURL := bc.container.NewBlockBlobURL(fileName)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	_, err = azblob.UploadFileToBlockBlob(ctx, file, blobURL, azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024,
		Parallelism: 16})
}
