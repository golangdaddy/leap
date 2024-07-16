package common

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func (self *GCPClients) GetObjectAndUnmarshal(ctx context.Context, bucket *storage.BucketHandle, objectName string, dst interface{}) error {
	b, err := self.GetObjectGCS(ctx, bucket, objectName)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

func (self *GCPClients) GetObjectGCS(ctx context.Context, bucket *storage.BucketHandle, objectName string) ([]byte, error) {
	r, err := bucket.Object(objectName).NewReader(context.Background())
	if err != nil {
		return nil, fmt.Errorf("storage.GetObject: %w", err)
	}
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("storage.GetObject: %w", err)
	}
	if err := r.Close(); err != nil {
		return nil, fmt.Errorf("storage.GetObject: %w", err)
	}
	return b, nil
}

// ListObjectGCS returns the contents of the directory with the given prefix
func (self *GCPClients) ListObjectsGCS(ctx context.Context, bucketName, prefix string) (names []string, err error) {

	it := self.GCS().Bucket(bucketName).Objects(ctx, &storage.Query{Prefix: prefix})
	for {
		var attrs *storage.ObjectAttrs
		attrs, err = it.Next()
		if err == iterator.Done {
			return names, nil
		}
		if err != nil {
			return
		}
		names = append(names, attrs.Name)
	}
	return names, nil
}

// WriteObjectGCS uploads a file to a bucket
func (self *GCPClients) WriteObjectGCS(ctx context.Context, bucket *storage.BucketHandle, objectName string, b []byte) error {

	obj := bucket.Object(objectName)

	// Create a writer for the object
	wc := obj.NewWriter(ctx)
	defer wc.Close()

	// Write content to the object
	if _, err := wc.Write(b); err != nil {
		return fmt.Errorf("failed to write to object: %v", err)
	}

	// Check for errors during the close operation
	if err := wc.Close(); err != nil {
		return fmt.Errorf("failed to close object writer: %v", err)
	}

	return nil
}
