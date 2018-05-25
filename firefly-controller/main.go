package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/Sirupsen/logrus"

	fireflyclient "github.com/frodehus/firefly/pkg/client/clientset/versioned"
	fireflyinformer "github.com/frodehus/firefly/pkg/client/informers/externalversions/firefly/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
)

func main() {
	client, fireflyClient := getKubernetesClient()
	informer := fireflyinformer.NewFireflyApplicationInformer(
		fireflyClient,
		meta_v1.NamespaceAll,
		0,
		cache.Indexers{},
	)
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			log.Infof("Added firefly application: %s", key)
			if err == nil {
				queue.Add(key)
			}
		},
	})

	stop := make(chan struct{})
	defer close(stop)

	controller := Controller{
		logger:    log.NewEntry(log.New()),
		clientset: client,
		informer:  informer,
		queue:     queue,
		handler:   &FireFlyAppHandler{},
	}

	go controller.Run(stop)

	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGTERM)
	signal.Notify(sigTerm, syscall.SIGINT)
	<-sigTerm
}

func getKubernetesClient() (kubernetes.Interface, fireflyclient.Interface) {
	kubeConfigPath := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Fatalf("getClusterConfig: %v", err)
		}
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("getClusterConfig: %v", err)
	}

	fireflyClient, err := fireflyclient.NewForConfig(config)
	if err != nil {
		log.Fatalf("getClusterConfig: %v", err)
	}

	log.Print("Successfully constructed k8s client")
	return client, fireflyClient
}
