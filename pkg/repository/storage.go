package repository

import (
	"context"
	"fmt"
	"go-naming-convention/pkg/common"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

// BlobStore .
type BlobStore struct {
	ctx       context.Context
	container azblob.ContainerURL
}

// NewBlobStore ...
func NewBlobStore(accountName string, accountKey string, containerName string) *BlobStore {
	// Create a default request pipeline using your storage account name and account key.
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		common.Logger.Error("Invalid credentials with error: " + err.Error())
	}

	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	// From the Azure portal, get your storage account blob service URL endpoint.
	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))

	// Create a ContainerURL object that wraps the container URL and a request
	// pipeline to make requests.
	containerURL := azblob.NewContainerURL(*URL, p)
	blobStore := new(BlobStore)
	blobStore.container = containerURL

	ctx := context.Background()
	_, err = containerURL.Create(ctx, azblob.Metadata{}, azblob.PublicAccessNone)
	// 409 is intentionally ignored
	blobStore.ctx = ctx
	return blobStore
}

// GetBlob downloads the specified blob content
func (blobStore *BlobStore) GetBlob(blobName string) (string, error) {

	blob := blobStore.container.NewBlobURL(blobName)
	resp, err := blob.Download(blobStore.ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false)

	if err != nil {
		return "", err
	}
	defer resp.Response().Body.Close()
	body, err := ioutil.ReadAll(resp.Body(azblob.RetryReaderOptions{}))
	return string(body), err
}

// CreateBlockBlob creates a new block blob
func (blobStore *BlobStore) CreateBlockBlob(blobName string, data string) (azblob.BlockBlobURL, error) {
	b := blobStore.container.NewBlockBlobURL(blobName)
	data = "blob created by Azure-Samples, okay to delete!"

	_, err := b.Upload(
		blobStore.ctx,
		strings.NewReader(data),
		azblob.BlobHTTPHeaders{
			ContentType: "text/plain",
		},
		azblob.Metadata{},
		azblob.BlobAccessConditions{},
	)
	return b, err
}
