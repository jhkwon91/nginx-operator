package assets

import (
	"embed"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	// go:embed manifests/*
	manifests  embed.FS
	appsScheme = runtime.NewScheme()
	appsCodecs = serializer.NewCodecFactory(appsScheme)
)

func init() {
	if err := appsv1.AddToScheme(appsScheme); err != nil {
		panic(err)
	}
}

func GetDeploymentFromFile(name string) *appsv1.Deployment {

	deploymentSample := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: "nginx-deployment"
  namespace: "nginx-operator-ns"
  labels:
    app: "nginx"
spec:
  replicas: 1
  selector:
    matchLabels:
      app: "nginx"
  template:
    metadata:
      labels:
        app: "nginx"
    spec:
      containers:
        - name: "nginx"
          image: "nginx:latest"
          ports:
          - containerPort: 80`

	deploymentBytes, err := manifests.ReadFile(name)
	if err != nil {
		// panic(err)
		fmt.Printf("[error] %#v\n", err)
		deploymentBytes = []byte(deploymentSample)
	}

	deploymentObject, err := runtime.Decode(
		appsCodecs.UniversalDecoder(appsv1.SchemeGroupVersion),
		deploymentBytes,
	)
	if err != nil {
		panic(err)
	}

	return deploymentObject.(*appsv1.Deployment)
}
