package consts

import "github.com/kubeflow/katib/pkg/util/v1alpha3/env"

const (
	// ConfigExperimentSuggestionName is the config name of the
	// suggestion client implementation in experiment controller.
	ConfigExperimentSuggestionName = "experiment-suggestion-name"
	// ConfigCertLocalFS is the config name which indicates if we
	// should store the cert in file system.
	ConfigCertLocalFS = "cert-local-filesystem"

	// LabelExperimentName is the label of experiment name.
	LabelExperimentName = "experiment"
	// LabelSuggestionName is the label of suggestion name.
	LabelSuggestionName = "suggestion"

	// ContainerSuggestion is the container name in Suggestion.
	ContainerSuggestion = "suggestion"

	// DefaultSuggestionPort is the default port of suggestion service.
	DefaultSuggestionPort = 6789
	// DefaultSuggestionPortName is the default port name of suggestion service.
	DefaultSuggestionPortName = "katib-api"
	// DefaultGRPCService is the default service name in Suggestion,
	// which is used to run healthz check using grpc probe.
	DefaultGRPCService = "manager.v1alpha3.Suggestion"

	// DefaultKatibNamespaceEnvName is the default env name of katib namespace
	DefaultKatibNamespaceEnvName = "KATIB_CORE_NAMESPACE"

	// KatibConfigMapName is the config map constants
	// Configmap name which includes Katib's configuration
	KatibConfigMapName = "katib-config"
	// LabelSuggestionTag is the name of suggestion config in configmap.
	LabelSuggestionTag = "suggestion"
	// LabelSuggestionImageTag is the name of suggestion image config in configmap.
	LabelSuggestionImageTag = "image"
	// LabelSuggestionCPULimitTag is the name of suggestion CPU Limit config in configmap.
	LabelSuggestionCPULimitTag = "cpuLimit"
	// DefaultCPULimit is the default value for CPU Limit
	DefaultCPULimit = "500m"
	// LabelSuggestionCPURequestTag is the name of suggestion CPU Request config in configmap.
	LabelSuggestionCPURequestTag = "cpuRequest"
	// DefaultCPURequest is the default value for CPU Request
	DefaultCPURequest = "50m"
	// LabelSuggestionMemLimitTag is the name of suggestion Mem Limit config in configmap.
	LabelSuggestionMemLimitTag = "memLimit"
	// DefaultMemLimit is the default value for mem Limit
	DefaultMemLimit = "100Mi"
	// LabelSuggestionMemRequestTag is the name of suggestion Mem Request config in configmap.
	LabelSuggestionMemRequestTag = "memRequest"
	// DefaultMemRequest is the default value for mem Request
	DefaultMemRequest = "10Mi"
	// LabelMetricsCollectorSidecar is the name of metrics collector config in configmap.
	LabelMetricsCollectorSidecar = "metrics-collector-sidecar"
	// LabelMetricsCollectorSidecarImage is the name of metrics collector image config in configmap.
	LabelMetricsCollectorSidecarImage = "image"

	// ReconcileErrorReason is the reason when there is a reconcile error.
	ReconcileErrorReason = "ReconcileError"

	// JobKindJob is the kind of the Kubernetes Job.
	JobKindJob = "Job"
	// JobKindTF is the kind of TFJob.
	JobKindTF = "TFJob"
	// JobKindPyTorch is the kind of PyTorchJob.
	JobKindPyTorch = "PyTorchJob"

	// JobVersionJob is the api version of Kubernetes Job.
	JobVersionJob = "v1"
	// JobVersionTF is the api version of TFJob.
	JobVersionTF = "v1"
	// JobVersionPyTorch is the api version of PyTorchJob.
	JobVersionPyTorch = "v1"

	// JobGroupJob is the group name of Kubernetes Job.
	JobGroupJob = "batch"
	// JobGroupKubeflow is the group name of Kubeflow.
	JobGroupKubeflow = "kubeflow.org"
)

var (
	DefaultKatibNamespace = env.GetEnvOrDefault(DefaultKatibNamespaceEnvName, "kubeflow")
)
