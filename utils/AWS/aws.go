package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	u "github.com/Raise-the-Voices/raise-the-voices/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var bucket string = os.Getenv("bucket")
var region string = os.Getenv("region")
var bucketURL string = fmt.Sprintf("%s.s3-%s.amazonaws.com", bucket, region)

var svc = generateNewSession()

func generateNewSession() *s3.S3 {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(os.Getenv("awsID"), os.Getenv("secretKEY"), ""),
	})

	// Create S3 service client
	return s3.New(sess)
}

// PutObjectToAWSS3 Upload file to S3
func PutObjectToAWSS3(finalName, filetype string, fileReader *bytes.Reader) error {
	_, errPUT := svc.PutObject(&s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &finalName,
		Body:        fileReader,
		ContentType: aws.String(filetype),
	})
	if errPUT != nil {
		return errPUT
	}
	return nil
}

// DeleteObjectAWSS3 Delete file to S3
func DeleteObjectAWSS3(ProfileImageURL string) (string, error) {
	tempProfileURL := strings.Split(ProfileImageURL, "/")
	victimURL := tempProfileURL[2]
	objectKey := tempProfileURL[3]
	err := errors.New("the object is not hosted in S3")

	if victimURL == bucketURL {
		thumbURL, thumbExtension := u.FindThumbnailURL(ProfileImageURL)
		if u.IsThisItAnImage(thumbExtension) {
			msj, errDelete := deleteObject(objectKey)
			if errDelete != nil {
				return msj, errDelete
			}
			return deleteObject(thumbURL)
		}
		return deleteObject(objectKey)
	}
	return "the image is not hosted on our server", err
}

// GenerateProfileS3URL join var to make profile image URL var
func GenerateProfileS3URL(finalName string) string {
	return fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", bucket, region, finalName)
}

func deleteObject(objectKey string) (string, error) {
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return "an error occurred deleting the photo", err
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return "an error occurred deleting the photo", err
	}
	return "the image was successfully deleted", nil
}
