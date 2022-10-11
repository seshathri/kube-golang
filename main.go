package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	var version string
	var flagvar int
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		flag.StringVar(&version, "version", "1.14", "version number for ngix")
		flag.IntVar(&flagvar, "scale", 2, "scale value for the flag")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		flag.StringVar(&version, "version", "1.14", "version number for nginx")
		flag.IntVar(&flagvar, "scale", 2, "scale value for the flag")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	fmt.Println(reflect.TypeOf(deploymentsClient))
	nginxVersion := "nginx" + ":" + version
	nginxScale := int32(flagvar)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test1-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(nginxScale),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "test",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "test",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: nginxVersion,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 9080,
								},
							},
						},
					},
				},
			},
		},
	}
	// Create Deployment
	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

kubeOptions:
	kubeSelection := prompt()

	switch kubeSelection {
	case "get":
		fmt.Printf("Please enter Deployment Name:")
		deploymentScanner := bufio.NewScanner(os.Stdin)
		var deploymentName string
		for deploymentScanner.Scan() {
			deploymentName = deploymentScanner.Text()
			break
		}
		fmt.Println("Getting deployment...", deploymentName)
		deployment, err := deploymentsClient.Get(context.TODO(), deploymentName, metav1.GetOptions{})
		if err != nil {
			panic(err)
		}
		response, err := json.MarshalIndent(*deployment, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("\n", string(response))
		goto kubeOptions
	case "list":
		fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
		list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for _, d := range list.Items {
			fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
		}
		goto kubeOptions
	case "delete":
		fmt.Println("Deleting deployment...")
		deletePolicy := metav1.DeletePropagationForeground
		if err := deploymentsClient.Delete(context.TODO(), "test1-deployment", metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
		fmt.Println("Deleted deployment.")
		goto kubeOptions
	case "exit":
		break
	default:
		fmt.Println("No Specifc Option Selected , Please select 1,2,3")
		goto kubeOptions
	}

}

func prompt() string {
	fmt.Printf(`-> Please Select Options to continue
	get ->  To GET deployment,
	list ->  To List the Kube Deployments
	delete ->  To Delete the Current Kube Deployment
	exit -> To Exit the CLI `)
	fmt.Printf("\n")
	fmt.Printf("Enter Option Here ->")
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return input
}

func int32Ptr(i int32) *int32 { return &i }
