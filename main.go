package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// Load Kubernetes config from the local computer
	config, err := loadKubeConfig()
	if err != nil {
		fmt.Printf("Error loading Kubernetes config: %v\n", err)
		return
	}

	// Create Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating Kubernetes clientset: %v\n", err)
		return
	}

	// Start the HTTP server on the specified port
	port := "9000"
	fmt.Printf("HTTP server started on http://localhost:%s...\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting HTTP server: %v\n", err)
		return
	}
}

func loadKubeConfig() (*rest.Config, error) {
	home := homedir.HomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")

	// Use in-cluster config if it exists, otherwise, use the kubeconfig from the local computer
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}
