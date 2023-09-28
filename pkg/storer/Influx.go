package storer

import (
	"3sigmas-monitorVisualization/pkg/data"
	"context"
	"errors"
	"github.com/getsentry/sentry-go"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"log"
	"sync"
)

type InfluxStorer struct {
	client       influxdb2.Client
	organization *domain.Organization
	bucketApi    api.BucketsAPI
	bucketPrefix string
	mu           sync.Mutex
}

func NewInfluxStorer(env data.Env) *InfluxStorer {
	client := influxdb2.NewClient(env.InfluxUrl, env.InfluxToken)
	org, err := client.OrganizationsAPI().FindOrganizationByName(context.Background(), env.InfluxOrg)
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}
	return &InfluxStorer{
		client:       client,
		organization: org,
		bucketApi:    client.BucketsAPI(),
		bucketPrefix: env.InfluxPrefix,
	}
}

func (s *InfluxStorer) setBucket(bucketName string) *domain.Bucket {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.client == nil {
		return nil
	}
	bucketApi := s.client.BucketsAPI()
	bucket, err := bucketApi.FindBucketByName(context.Background(), bucketName)
	if bucket == nil {
		log.Printf("Bucket %s not found, creating it\n", bucketName)
		bucket, err = bucketApi.CreateBucketWithName(context.Background(), s.organization, bucketName)
		if err != nil {
			sentry.CaptureException(err)
		}
	}
	return bucket
}

func (s *InfluxStorer) Store(project string, source string, measures []data.Measure) error {

	bucket := s.setBucket(s.bucketPrefix + "." + project + "." + source)
	if bucket == nil {
		return errors.New("bucket not found")
	}

	writeAPI := s.client.WriteAPIBlocking(s.organization.Name, bucket.Name)

	for _, measure := range measures {
		tags := measure.Tags()
		fields := measure.Fields()
		point := write.NewPoint(measure.Measurement(), tags, fields, measure.Date())
		s.mu.Lock()
		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			sentry.CaptureException(err)
			panic(err)
		}
		s.mu.Unlock()
		// log.Printf("Stored measure %s\n", measure)
	}
	log.Printf("Stored %d measures in %s\n", len(measures), bucket.Name)
	return nil
}
