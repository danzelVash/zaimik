package yandex_cloud

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
)

const (
	bucketName = "loan-companies-logos"
)

//func GetObjectList(bucketName string) []types.Object {
//	// Создаем кастомный обработчик эндпоинтов, который для сервиса S3 и региона ru-central1 выдаст корректный URL
//	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
//		if service == s3.ServiceID && region == "ru-central1" {
//			return aws.Endpoint{
//				PartitionID:   "yc",
//				URL:           "https://storage.yandexcloud.net",
//				SigningRegion: "ru-central1",
//			}, nil
//		}
//		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
//	})
//	// Подгружаем конфигрурацию из ~/.aws/*
//	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))
//	if err != nil {
//		logrus.Fatal(err, 1)
//	}
//
//	// Создаем клиента для доступа к хранилищу S3
//	client := s3.NewFromConfig(cfg)
//
//	// Запрашиваем список бакетов
//	result, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
//		Bucket: aws.String(bucketName),
//	})
//	if err != nil {
//		logrus.Fatal(err, 2)
//	}
//
//	return result.Contents
//}

func GetObjectFromYandexCloud(key string) ([]byte, error) {
	accessKeyId := os.Getenv("YANDEX_CLOUD_KEY_ID")
	secretKey := os.Getenv("YANDEX_CLOUD_SECRET_KEY")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     accessKeyId,
				SecretAccessKey: secretKey,
			},
		}),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "yc",
				URL:           "https://storage.yandexcloud.net",
				SigningRegion: "ru-central1",
			}, nil
		})),
	)

	if err != nil {
		return []byte{}, errors.Errorf("error while loading AWS config: %s", err.Error())
	}

	client := s3.NewFromConfig(cfg)

	bucket := bucketName

	resp, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return []byte{}, errors.Errorf("error while trying load object: %s", err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Errorf("memory leak: error closing response body: %s", err.Error())
		}
	}(resp.Body)

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.Errorf("error while reading resp.body: %s", err.Error())
	}

	return bytes, nil
}

// TODO create a function which creating bucket or folder and function which put object into folder

func PutObjectInBucket(key string, body io.Reader) error {

	accessKeyId := os.Getenv("YANDEX_CLOUD_KEY_ID")
	secretKey := os.Getenv("YANDEX_CLOUD_SECRET_KEY")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     accessKeyId,
				SecretAccessKey: secretKey,
			},
		}),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "yc",
				URL:           "https://storage.yandexcloud.net",
				SigningRegion: "ru-central1",
			}, nil
		})),
	)

	if err != nil {
		return errors.Errorf("error while loading AWS config: %s", err.Error())
	}

	client := s3.NewFromConfig(cfg)

	bucket := bucketName

	params := &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   body,
	}

	_, err = client.PutObject(context.Background(), params)
	return err
}

// TODO сделать удаление

func DeleteObjectFromCloud(ctx context.Context, ch chan<- error, key string) {
	accessKeyId := os.Getenv("YANDEX_CLOUD_KEY_ID")
	secretKey := os.Getenv("YANDEX_CLOUD_SECRET_KEY")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     accessKeyId,
				SecretAccessKey: secretKey,
			},
		}),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "yc",
				URL:           "https://storage.yandexcloud.net",
				SigningRegion: "ru-central1",
			}, nil
		})),
	)

	if err != nil {
		ch <- errors.Errorf("error while loading AWS config: %s", err.Error())
	}

	client := s3.NewFromConfig(cfg)

	bucket := bucketName

	params := &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}

	_, err = client.DeleteObject(ctx, params)
	ch <- err
}
