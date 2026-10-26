package photos

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type Store struct {
	client *s3.Client
	bucket string
}

func NewStore(ctx context.Context) (Store, error) {
	bucket := os.Getenv("PHOTOS_BUCKET")
	if bucket == "" {
		return Store{}, errors.New("PHOTOS_BUCKET is not set")
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return Store{}, err
	}

	// Local S3 servers (MinIO) are addressed by hostname, so they need path-style URLs.
	customEndpoint := os.Getenv("AWS_ENDPOINT_URL_S3") != ""

	return Store{
		client: s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = customEndpoint
		}),
		bucket: bucket,
	}, nil
}

func key(photoID uuid.UUID) string {
	return "photos/" + photoID.String()
}

func (s Store) Save(ctx context.Context, data []byte, photoID uuid.UUID) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key(photoID)),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(http.DetectContentType(data)),
	})
	return err
}

func (s Store) Get(ctx context.Context, photoID uuid.UUID) ([]byte, error) {
	out, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key(photoID)),
	})
	if err != nil {
		return nil, err
	}
	defer out.Body.Close()

	return io.ReadAll(out.Body)
}
