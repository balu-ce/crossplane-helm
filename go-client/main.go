package main

import (
	"fmt"
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
)

func main() {
	chartPath := "../helm-chart-sources/eks"
	namespace := "default"
	releaseName := "crossplane-eks"

	settings := cli.New()

	actionConfig := new(action.Configuration)
	// You can pass an empty string instead of settings.Namespace() to list
	// all namespaces
	if err := actionConfig.Init(settings.RESTClientGetter(), namespace,
		os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		fmt.Printf("coming in")
		log.Printf("%+v", err)
		os.Exit(1)
	}

	// define values
	vals := map[string]interface{}{
		"clusterName": "helm-poc-test",
		"secGroupIds": [1]string{"sg-0af816ae1b1ffccad"},
		"region":      "us-east-1",
		"version":     "1.21",
		"roleArn":     "arn:aws:iam::092463844305:role/fwks-staging-fwks-poc-master-role",
	}

	// load chart from the path
	chart, err := loader.Load(chartPath)
	if err != nil {
		panic(err)
	}

	client := action.NewInstall(actionConfig)
	client.Namespace = namespace
	client.ReleaseName = releaseName
	// client.DryRun = true - very handy!

	// install the chart here
	rel, err := client.Run(chart, vals)
	if err != nil {
		panic(err)
	}

	log.Printf("Installed Chart from path: %s in namespace: %s\n", rel.Name, rel.Namespace)
	// this will confirm the values set during installation
	log.Println(rel.Config)
}
