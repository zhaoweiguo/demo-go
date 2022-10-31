package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func main() {
	path := "demo4.json"
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}

	cfg := Spec{}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(cfg.Steps[1].Docker.Command)
}

type (
	Spec struct {
		Metadata Metadata  `json:"metadata,omitempty"`
		Platform Platform  `json:"platform,omitempty"`
		Secrets  []*Secret `json:"secrets,omitempty"`
		Steps    []*Step   `json:"steps,omitempty"`
		Files    []*File   `json:"files,omitempty"`

		// Docker-specific settings. These settings are
		// only used by the Docker and Kubernetes runtime
		// drivers.
		Docker *DockerConfig `json:"docker,omitempty"`

		// Qemu-specific settings. These settings are only
		// used by the qemu runtime driver.
		Qemu *QemuConfig `json:"qemu,omitempty"`

		// VMWare Fusion settings. These settings are only
		// used by the VMWare runtime driver.
		Fusion *FusionConfig `json:"fusion,omitempty"`
	}

	Metadata struct {
		UID       string            `json:"uid,omitempty"`
		Namespace string            `json:"namespace,omitempty"`
		Name      string            `json:"name,omitempty"`
		Labels    map[string]string `json:"labels,omitempty"`
	}

	// Platform defines the target platform.
	Platform struct {
		OS      string `json:"os,omitempty"`
		Arch    string `json:"arch,omitempty"`
		Variant string `json:"variant,omitempty"`
		Version string `json:"version,omitempty"`
	}

	// Secret represents a secret variable.
	Secret struct {
		Metadata Metadata `json:"metadata,omitempty"`
		Data     string   `json:"data,omitempty"`
	}

	// Step defines a pipeline step.
	Step struct {
		Metadata     Metadata          `json:"metadata,omitempty"`
		Detach       bool              `json:"detach,omitempty"`
		DependsOn    []string          `json:"depends_on,omitempty"`
		Devices      []*VolumeDevice   `json:"devices,omitempty"`
		Envs         map[string]string `json:"environment,omitempty"`
		Files        []*FileMount      `json:"files,omitempty"`
		IgnoreErr    bool              `json:"ignore_err,omitempty"`
		IgnoreStdout bool              `json:"ignore_stderr,omitempty"`
		IgnoreStderr bool              `json:"ignore_stdout,omitempty"`
		Resources    *Resources        `json:"resources,omitempty"`
		RunPolicy    RunPolicy         `json:"run_policy,omitempty"`
		Secrets      []*SecretVar      `json:"secrets,omitempty"`
		Volumes      []*VolumeMount    `json:"volumes,omitempty"`
		WorkingDir   string            `json:"working_dir,omitempty"`

		// Docker-specific settings. These settings are
		// only used by the Docker and Kubernetes runtime
		// drivers.
		Docker *DockerStep `json:"docker,omitempty"`
	}

	// File defines a file that should be uploaded or
	// mounted somewhere in the step container or virtual
	// machine prior to command execution.
	File struct {
		Metadata Metadata `json:"metadata,omitempty"`
		Data     []byte   `json:"data,omitempty"`
	}

	// DockerConfig configures a Docker-based pipeline.
	DockerConfig struct {
		Auths   []*DockerAuth `json:"auths,omitempty"`
		Volumes []*Volume     `json:"volumes,omitempty"`
	}

	// QemuConfig configures a Qemu-based pipeline.
	QemuConfig struct {
		Image string `json:"image,omitempty"`
	}

	// FusionConfig configures a VMWare Fusion-based pipeline.
	FusionConfig struct {
		Image string `json:"image,omitempty"`
	}

	// 3级

	// VolumeDevice describes a mapping of a raw block
	// device within a container.
	VolumeDevice struct {
		Name       string `json:"name,omitempty"`
		DevicePath string `json:"path,omitempty"`
	}

	// FileMount defines how a file resource should be
	// mounted or included in the runtime environment.
	FileMount struct {
		Name string `json:"name,omitempty"`
		Path string `json:"path,omitempty"`
		Mode int64  `json:"mode,omitempty"`

		// Base string `json:"base,omitempty"`
	}

	// VolumeMount describes a mounting of a Volume
	// within a container.
	VolumeMount struct {
		Name string `json:"name,omitempty"`
		Path string `json:"path,omitempty"`
	}

	// Resources describes the compute resource
	// requirements.
	Resources struct {
		// Limits describes the maximum amount of compute
		// resources allowed.
		Limits *ResourceObject `json:"limits,omitempty"`

		// Requests describes the minimum amount of
		// compute resources required.
		Requests *ResourceObject `json:"requests,omitempty"`
	}

	// SecretVar represents an environment variable
	// sources from a secret.
	SecretVar struct {
		Name string `json:"name,omitempty"`
		Env  string `json:"env,omitempty"`
	}

	// DockerStep configures a docker step.
	DockerStep struct {
		Args       []string   `json:"args,omitempty"`
		Command    []string   `json:"command,omitempty"`
		DNS        []string   `json:"dns,omitempty"`
		DNSSearch  []string   `json:"dns_search,omitempty"`
		ExtraHosts []string   `json:"extra_hosts,omitempty"`
		Image      string     `json:"image,omitempty"`
		Network    string     `json:"network,omitempty"`
		Networks   []string   `json:"networks,omitempty"`
		Ports      []*Port    `json:"ports,omitempty"`
		Privileged bool       `json:"privileged,omitempty"`
		PullPolicy PullPolicy `json:"pull_policy,omitempty"`
		User       string     `json:"user"`
	}

	// 4级

	// DockerAuth defines dockerhub authentication credentials.
	DockerAuth struct {
		Address  string `json:"address,omitempty"`
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}

	// Volume that can be mounted by containers.
	Volume struct {
		Metadata Metadata        `json:"metadata,omitempty"`
		EmptyDir *VolumeEmptyDir `json:"temp,omitempty"`
		HostPath *VolumeHostPath `json:"host,omitempty"`
	}

	// ResourceObject describes compute resource
	// requirements.
	ResourceObject struct {
		CPU    int64 `json:"cpu,omitempty"`
		Memory int64 `json:"memory,omitempty"`
	}

	// Port represents a network port in a single container.
	Port struct {
		Port     int    `json:"port,omitempty"`
		Host     int    `json:"host,omitempty"`
		Protocol string `json:"protocol,omitempty"`
	}

	// VolumeEmptyDir mounts a temporary directory from the
	// host node's filesystem into the container. This can
	// be used as a shared scratch space.
	VolumeEmptyDir struct {
		Medium    string `json:"medium,omitempty"`
		SizeLimit int64  `json:"size_limit,omitempty"`
	}

	// VolumeHostPath mounts a file or directory from the
	// host node's filesystem into your container.
	VolumeHostPath struct {
		Path string `json:"path,omitempty"`
	}
)

// RunPolicy defines the policy for starting containers
// based on the point-in-time pass or fail state of
// the pipeline.
type RunPolicy int

// RunPolicy enumeration.
const (
	RunOnSuccess RunPolicy = iota
	RunOnFailure
	RunAlways
	RunNever
)

func (r RunPolicy) String() string {
	return runPolicyID[r]
}

var runPolicyID = map[RunPolicy]string{
	RunOnSuccess: "on-success",
	RunOnFailure: "on-failure",
	RunAlways:    "always",
	RunNever:     "never",
}
var runPolicyName = map[string]RunPolicy{
	"":           RunOnSuccess,
	"on-success": RunOnSuccess,
	"on-failure": RunOnFailure,
	"always":     RunAlways,
	"never":      RunNever,
}

// MarshalJSON marshals the string representation of the
// run type to JSON.
func (r *RunPolicy) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(runPolicyID[*r])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmarshals the json representation of the
// run type from a string value.
func (r *RunPolicy) UnmarshalJSON(b []byte) error {
	// unmarshal as string
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	// lookup value
	*r = runPolicyName[s]
	return nil
}

// PullPolicy defines the container image pull policy.
type PullPolicy int
