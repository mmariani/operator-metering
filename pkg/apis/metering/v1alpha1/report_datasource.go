package v1alpha1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReportDataSourceList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`
	Items         []*ReportDataSource `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ReportDataSource struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReportDataSourceSpec   `json:"spec"`
	Status ReportDataSourceStatus `json:"status"`
}

type ReportDataSourceSpec struct {
	// Prommsum represents a datasource which holds Prometheus metrics
	Promsum *PrometheusMetricsDataSource `json:"promsum"`
	// AWSBilling represents a datasource which points to a pre-existing S3
	// bucket.
	AWSBilling *AWSBillingDataSource `json:"awsBilling"`
}

type AWSBillingDataSource struct {
	Source *S3Bucket `json:"source"`
}

type S3Bucket struct {
	Region string `json:"region"`
	Bucket string `json:"bucket"`
	Prefix string `json:"prefix"`
}

type PrometheusQueryConfig struct {
	QueryInterval *meta.Duration `json:"queryInterval,omitempty"`
	StepSize      *meta.Duration `json:"stepSize,omitempty"`
	ChunkSize     *meta.Duration `json:"chunkSize,omitempty"`
}

type PrometheusConnectionConfig struct {
	URL string `json:"url,omitempty"`
}

type PrometheusMetricsDataSource struct {
	Query            string                      `json:"query"`
	QueryConfig      *PrometheusQueryConfig      `json:"queryConfig,omitempty"`
	Storage          *StorageLocationRef         `json:"storage,omitempty"`
	PrometheusConfig *PrometheusConnectionConfig `json:"prometheusConfig,omitempty"`
}

type ReportDataSourceStatus struct {
	TableName                    string                        `json:"tableName,omitempty"`
	PrometheusMetricImportStatus *PrometheusMetricImportStatus `json:"prometheusMetricImportStatus,omitempty"`
}

type PrometheusMetricImportStatus struct {
	LastImportTime             *meta.Time `json:"lastImportTime,omitempty"`
	EarliestImportedMetricTime *meta.Time `json:"earliestImportedMetricTime,omitempty"`
	NewestImportedMetricTime   *meta.Time `json:"newestImportedMetricTime,omitempty"`
}
