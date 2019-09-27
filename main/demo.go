package main

import (
	"flag"
	"path/filepath"

	"github.com/mlycore/log"
	corev1 "k8s.io/api/core/v1"
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
        "k8s.io/client-go/util/homedir"
        "k8s.io/client-go/tools/clientcmd"
)

func main() {
	c := KubernetesClientset()
	list := c.ListPod("kube-system").Items
	for _, p := range list {
		log.Infof("%s", p.Name)
	}
	log.Infof("this is a log")
}

// Client is built upon the real Kubernetes client-go
type Client struct {
	Config *rest.Config
	*kubernetes.Clientset
}

// KubernetesClientset create kubernetes clients
func KubernetesClientset() *Client {
	var kubeconfig *string
	var err error
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("read config from flags error: %s", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("init client with config error: %s", err)
	}
	return &Client{
		Config:    config,
		Clientset: clientset,
	}
}

func (c *Client) ListPod(namespace string) *corev1.PodList {
	podlist, err := c.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Errorf("list pod error: %s", err)
		return nil
	}
	return podlist
}

func (c *Client) GetPod(namespace, podName string) *corev1.Pod {
	pod, err := c.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("get pod error: %s", err)
		return nil
	}
	return pod
}
