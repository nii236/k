package k8s

import (
	"io"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	v1 "k8s.io/api/core/v1"
)

type ClientSet interface {
	GetPods(namespace string) (*v1.PodList, error)
	GetNamespaces() (*v1.NamespaceList, error)
	GetPodContainers(podName string, namespace string) []string
	DeletePod(podName string, namespace string) error
	GetPodContainerLogs(podName string, containerName string, namespace string, o io.Writer) error
}
