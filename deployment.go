package skube

import "fmt"

func (s *Skube) ListDeployments(namespace string) {
	if namespace == "" {
		namespace = "default"
	}
	url := fmt.Sprintf("%s/apis/extensions/v1beta1/namespaces/%s/deployments",
		s.server, namespace)

	out, err := s.getRequest(url, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

}