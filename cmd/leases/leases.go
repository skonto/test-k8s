package main

import (
	"context"
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/klog/v2"
	"knative.dev/pkg/signals"
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	id := os.Getenv("POD_NAME")
	config, err := rest.InClusterConfig()
	if err != nil {
		klog.Fatal(err)
	}
	client := clientset.NewForConfigOrDie(config)
	ctx := signals.NewContext()
	eg, egCtx := errgroup.WithContext(ctx)
	for i := range 10 {
		lock := &resourcelock.LeaseLock{
			LeaseMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("lease-%d", i),
				Namespace: "default",
			},
			Client: client.CoordinationV1(),
			LockConfig: resourcelock.ResourceLockConfig{
				Identity: id,
			},
		}
		eg.Go(func() error {
			leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
				Lock:            lock,
				ReleaseOnCancel: true,
				LeaseDuration:   60 * time.Second,
				RenewDeadline:   40 * time.Second,
				RetryPeriod:     10 * time.Second,
				Callbacks: leaderelection.LeaderCallbacks{
					OnStartedLeading: func(ctx context.Context) {
						klog.Infof("leader elected: %s", id)
					},
					OnStoppedLeading: func() {
						klog.Infof("leader lost: %s", id)
					},
				},
			})
			return nil
		})
	}

	<-egCtx.Done()
	time.Sleep(5 * time.Second) // without this nothing is cleaned up, comment out and check

	//select {} // works and leases are cleaned up
}
