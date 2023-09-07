package storer

import (
	"3sigmas-monitorVisualization/pkg"
	"3sigmas-monitorVisualization/pkg/reader/data"
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"log"
)

type InfluxStorer struct {
	client       influxdb2.Client
	organization *domain.Organization
	bucketApi    api.BucketsAPI
}

func NewInfluxStorer() *InfluxStorer {
	client := influxdb2.NewClient(pkg.Url, pkg.TOKEN)
	org, err := client.OrganizationsAPI().FindOrganizationByName(context.Background(), pkg.OrganizationName)
	if err != nil {
		log.Fatal(err)
	}
	return &InfluxStorer{
		client:       client,
		organization: org,
		bucketApi:    client.BucketsAPI(),
	}
}

func (s *InfluxStorer) setBucket(bucketName string) *domain.Bucket {
	bucketApi := s.client.BucketsAPI()
	bucket, err := bucketApi.FindBucketByName(context.Background(), bucketName)
	if bucket == nil {
		bucket, err = bucketApi.CreateBucketWithName(context.Background(), s.organization, bucketName)
		if err != nil {
			panic(err)
		}
	}
	return bucket
}

func (s *InfluxStorer) Store(project string, measures []data.Measure) {

	bucket := s.setBucket(project)

	writeAPI := s.client.WriteAPIBlocking(s.organization.Name, bucket.Name)

	for _, measure := range measures {
		tags := map[string]string{
			"type": measure.Sensor,
		}
		fields := map[string]interface{}{
			"value":       measure.Value,
			"temperature": measure.Temperature,
		}
		point := write.NewPoint(measure.Captor, tags, fields, measure.Date)
		fmt.Println("Writing point: ", measure)
		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			panic(err)
		}
	}
}
