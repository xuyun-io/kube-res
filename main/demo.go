package main

import (
	"fmt"
	"flag"
	"path/filepath"
	"github.com/xuyun-io/kube-res/main/resources"

	"github.com/mlycore/log"
	corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
    "k8s.io/client-go/util/homedir"
    "k8s.io/client-go/tools/clientcmd"
)

const (
	testns = "observability"
)

func main() {
	c := KubernetesClientset()
	log.Infof("this is a log")
	if list, err := resources.Pod(testns); err == nil {
		for _, p := range list.Items{
			fmt.Printf("pod/%s/%s\n", p.Namespace, p.Name)		
		}
	}
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
