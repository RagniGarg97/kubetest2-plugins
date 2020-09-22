package tfvars

type TFVars struct {
	ReleaseMarker  string `json:"release_marker"`
	BuildVersion   string `json:"build_version"`
	ClusterName    string `json:"cluster_name"`
	ApiServerPort  int    `json:"apiserver_port"`
	WorkersCount   int    `json:"workers_count"`
	BootstrapToken string `json:"bootstrap_token"`
	KubeconfigPath string `json:"kubeconfig_path"`
	SSHPrivateKey  string `json:"ssh_private_key"`
	ExtraCerts     string `json:"extra_cert,omitempty"`
	IgnoreDestroy  bool   `json:"-"`
}
