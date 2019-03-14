package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type k8s struct {
	clientset kubernetes.Interface
}

func newk8s() (*k8s, error) {
	path := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return nil, err
	}

	client := k8s{}
	client.clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (k *k8s) getVersion() (string, error) {
	v, err := k.clientset.Discovery().ServerVersion()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", v), err
}

func (k *k8s) isVersion(major string, minor string) (bool, error) {
	v, err := k.clientset.Discovery().ServerVersion()
	if err != nil {
		return false, err
	}

	if v.Major != major {
		return false, errors.New("Major version does not match")
	} else {
		fmt.Println("Major version is correct")
	}
	if v.Minor != minor {
		return false, errors.New("Minor version does not match")
	} else {
		fmt.Println("Minor version is correct")
	}
	return true, nil
}

func parseError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	k8s, err := newk8s()
	parseError(err)
	v, err := k8s.getVersion()
	parseError(err)
	fmt.Println("The  Kubernetes Cluster Verison is ", v)
	isVersionCorrect, err := k8s.isVersion("1", "11+")
	parseError(err)
	fmt.Println("The Kubernetes version in use is ", isVersionCorrect)
}
