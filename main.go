package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {

	var concurrencyPerCPU int
	var downloadKey, downloadBucket, region string
	flag.IntVar(&concurrencyPerCPU, "cpc", 1, "concurrency per cpu")
	flag.StringVar(&downloadKey, "key", "", "key of download s3 file")
	flag.StringVar(&downloadBucket, "bucket", "", "bucket of download s3 file")
	flag.StringVar(&region, "region", "us-west-2", "aws s3 bucket region code, default: us-west-2")
	flag.Parse()

	if downloadKey == "" || downloadBucket == "" {
		fmt.Println("--key and --bucket must be specified")
		os.Exit(1)
	}

	downloadConcurrrency := runtime.NumCPU() * concurrencyPerCPU
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-west-2")}))
	downloader := s3manager.NewDownloader(sess,
		func(d *s3manager.Downloader) { d.PartSize = 64 * 1024 * 1024 },
		func(d *s3manager.Downloader) { d.Concurrency = downloadConcurrrency },
	)

	f, err := os.Create("/dev/null")
	if err != nil {
		fmt.Printf("failed to create file %s, %v\n", "/dev/null", err)
		return
	}

	start := time.Now()
	downloadBytes, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String("xjiaqing-tidb-br-test"),
		Key:    aws.String("test-0329-onefile/1tzero"),
	})

	if err != nil {
		fmt.Println("download file error")
		return
	}

	diffSeconds := time.Since(start).Seconds()
	speedInGbps := float64(downloadBytes) / float64(diffSeconds*1024*1024*1024) * 8
	fmt.Printf("file downloaded, %d bytes, concurrency: %d, cost time: %f seconds, speed: %.2f Gbps\n", downloadBytes, downloadConcurrrency, diffSeconds, speedInGbps)
}
