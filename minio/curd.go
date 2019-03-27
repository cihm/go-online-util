package minio

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/minio/minio-go"
)

var (
	minioClient *minio.Client
)

func Flow(endpoint, accessKeyID, secretAccessKey string) {

	//endpoint := "192.168.99.100:9001" // Server 位址
	//accessKeyID := "543534543543" // 在建立 Minio Server 時可設定
	//secretAccessKey := "zuf+34543543rfdgt" // 在建立 Minio Server 時可設定
	useSSL := false // 是否使用 https
	var err error
	// Initialize minio client object.
	minioClient, err = minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%#v\n", minioClient) // minioClient is now setup

	bucketName := "bucket1"
	bucketName2 := "bucket2"
	createBucket(bucketName, "us-east-1")
	createBucket(bucketName2, "us-east-1")
	uploadFile(bucketName, "cat2.png", "C:/Users/1409035/Dropbox/集智/wireframe/12_結果頁.png", "application/octet-stream")
	moveFile(bucketName, bucketName2, "cat2.png", "cat2Rename.png")
	getFileURL(bucketName2, "cat2Rename.png", 60)

	//test folder
	//uploadFile(bucketName, "folder/cat2.png", "C:/Users/1409035/Dropbox/集智/wireframe/12_結果頁.png", "application/octet-stream")
	//getFileURL(bucketName, "folder/cat2.png", 60)

	setBucketReadOnly("three-kingdoms", "output-st*")
	getBucketPolicy("three-kingdoms")
	//http://domainname:9000/three-kingdoms/output-stream1.m4s
	//就可以直接拿 不用prsignURL
}

// GetBucketPolicy get bucket policy
func getBucketPolicy(bucketName string) {
	policy, err := minioClient.GetBucketPolicy(bucketName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(policy)
}

func setBucketReadOnly(bucketName, prefix string) {
	if len(prefix) == 0 {
		prefix = "*"
	} else {
		prefix = prefix + "*"
	}
	policy := `{
        "Version":"2012-10-17",
        "Statement": [
            {
                "Action": ["s3:GetObject"],
                "Effect":"Allow",
                "Principal":{
                    "AWS":["*"]
                },
                "Resource": ["arn:aws:s3:::` + bucketName + `/` + prefix + `"]
            }
        ]
    }`
	err := minioClient.SetBucketPolicy(bucketName, policy)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// GetFileURL, Get file url
// 檔案來源 bucket, 檔名, url 有效時間(second)
func getFileURL(bucketName, objectName string, duration int64) {
	// Set request parameters
	reqParams := make(url.Values)
	//reqParams.Set("response-content-disposition", "attachment; filename=\""+objectName+".jpg\"")
	reqParams.Set("response-content-disposition", "attachment; filename=\""+objectName+"\"")

	// Gernerate presigned get object url.
	presignedURL, err := minioClient.PresignedGetObject(bucketName, objectName, time.Duration(duration)*time.Second, reqParams)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(presignedURL)
}

// RenameFile, Rename file
// 檔案來源 bucket, 檔案目的地 bucket, 來源檔名, 目標檔名
func moveFile(srcBucketName, dstBucketName, srcName, dstName string) {
	// 創建來源 src
	src := minio.NewSourceInfo(srcBucketName, srcName, nil)

	// 創建目的地 dst
	dst, err := minio.NewDestinationInfo(dstBucketName, dstName, nil, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Initiate copy object. 複製檔案
	err = minioClient.CopyObject(dst, src)
	if err != nil {
		log.Fatalln(err)
	}

	// 刪除來源檔案
	minioClient.RemoveObject(srcBucketName, srcName)

	log.Println("rename Successfully.")
}

// UploadFile, Upload file
func uploadFile(bucketName, objectName, filePath, contentType string) {
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}

//location default is "us-east-1"
func createBucket(bucketName string, location string) {
	err := minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check bucket exist
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)
}
