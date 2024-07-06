package main

import (
	"flag"
	apixclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
	"path/filepath"

	"knative.dev/pkg/signals"
)

func main() {
	klog.InitFlags(nil)
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		klog.Fatal(err)
	}
	ctx := signals.NewContext()
	aClient := apixclient.NewForConfigOrDie(config)
	crdClient := aClient.ApiextensionsV1().CustomResourceDefinitions()
	//patch := `{"status":{"storedVersions":["` + fmt.Sprintf("v1alpha1\", \"v1beta1") + `"]}}`
	patch := `{"status":{"storedVersions":["v1beta1"]}}`
	crd, err := crdClient.Get(ctx, "domainmappings.serving.knative.dev", metav1.GetOptions{})
	if err != nil {
		klog.Fatal(err)
	}
	_, err = crdClient.Patch(ctx, crd.Name, types.StrategicMergePatchType, []byte(patch), metav1.PatchOptions{}, "status")
	if err != nil {
		klog.Fatal(err)
	}
}
