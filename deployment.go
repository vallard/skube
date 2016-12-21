package skube

import (
	"encoding/json"
	"fmt"
	"time"
)

type (
	Metadata struct {
		CreatationTimestamp time.Time
		Name                string `json:"name,omitempty"`
		Namespace           string `json:"namespace,omitempty"`
	}
	Spec struct {
		Replicas int `json:"replicas,omitempty"`
	}

	Deployment struct {
		Metadata Metadata
		Spec     Spec
	}
	Deployments struct {
		ApiVersion string
		Items      []Deployment
		Kind       string
	}
)

func (s *Skube) ListDeployments(namespace string) ([]Deployment, error) {
	var d Deployments
	if namespace == "" {
		namespace = "default"
	}
	url := fmt.Sprintf("%s/apis/extensions/v1beta1/namespaces/%s/deployments",
		s.server, namespace)

	bytes, err := s.getRequest(url, nil)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(bytes, &d)
	fmt.Println(string(bytes))
	return d.Items, err
}

// http://kubernetes.io/docs/api-reference/extensions/v1beta1/operations/
func (s *Skube) UpdateDeployment(namespace string, deployment string, updates []byte) {
	url := fmt.Sprintf("%s/apis/extensions/v1beta1/namespaces/%s/deployments/%s",
		s.server,
		namespace,
		deployment)
	bytes, err := s.patchRequest(url, body)
}
