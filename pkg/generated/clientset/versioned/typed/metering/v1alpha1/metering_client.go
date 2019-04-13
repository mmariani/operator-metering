// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/operator-framework/operator-metering/pkg/apis/metering/v1alpha1"
	"github.com/operator-framework/operator-metering/pkg/generated/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type MeteringV1alpha1Interface interface {
	RESTClient() rest.Interface
	HiveTablesGetter
	PrestoTablesGetter
	ReportsGetter
	ReportDataSourcesGetter
	ReportGenerationQueriesGetter
	ReportPrometheusQueriesGetter
	StorageLocationsGetter
}

// MeteringV1alpha1Client is used to interact with features provided by the metering.openshift.io group.
type MeteringV1alpha1Client struct {
	restClient rest.Interface
}

func (c *MeteringV1alpha1Client) HiveTables(namespace string) HiveTableInterface {
	return newHiveTables(c, namespace)
}

func (c *MeteringV1alpha1Client) PrestoTables(namespace string) PrestoTableInterface {
	return newPrestoTables(c, namespace)
}

func (c *MeteringV1alpha1Client) Reports(namespace string) ReportInterface {
	return newReports(c, namespace)
}

func (c *MeteringV1alpha1Client) ReportDataSources(namespace string) ReportDataSourceInterface {
	return newReportDataSources(c, namespace)
}

func (c *MeteringV1alpha1Client) ReportGenerationQueries(namespace string) ReportGenerationQueryInterface {
	return newReportGenerationQueries(c, namespace)
}

func (c *MeteringV1alpha1Client) ReportPrometheusQueries(namespace string) ReportPrometheusQueryInterface {
	return newReportPrometheusQueries(c, namespace)
}

func (c *MeteringV1alpha1Client) StorageLocations(namespace string) StorageLocationInterface {
	return newStorageLocations(c, namespace)
}

// NewForConfig creates a new MeteringV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*MeteringV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MeteringV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new MeteringV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MeteringV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MeteringV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *MeteringV1alpha1Client {
	return &MeteringV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MeteringV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
