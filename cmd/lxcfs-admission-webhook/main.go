package main

import (
	"context"
	"crypto/tls"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Meoop/lxcfs-admission-webhook/pkg/admission"
	"github.com/Meoop/lxcfs-admission-webhook/pkg/cert"
	"github.com/Meoop/lxcfs-admission-webhook/pkg/lxcfs"
	"github.com/Meoop/lxcfs-admission-webhook/pkg/version"
	"github.com/golang/glog"
)

func main() {
	// get kubeconfig file path
	var kubeconfig string
	flag.StringVar(&kubeconfig, "kubeconfig", "", "kubernetes config file path.")
	flag.Parse()

	glog.Infof("Version:%s, Commit:%s, RepoRoot:%s", version.VERSION, version.COMMIT, version.REPOROOT)

	// generate cert
	cert, key, err := cert.GenCert("kube-system", "lxcfs-admission-webhook")
	if err != nil {
		glog.Fatalf("generate cert error: %v", err)
	}
	sCert, err := tls.X509KeyPair(cert, key)
	if err != nil {
		glog.Fatalf("parses a public/private key error: %v", err)
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{sCert},
		ClientAuth:   tls.NoClientCert,
	}

	err = admission.CreateMutatingWebhookConfiguration(kubeconfig, cert)
	if err != nil {
		glog.Fatalf("Create MutatingWebhookConfiguration error: %v", err)
	}

	whsvr := lxcfs.NewServer(443, tlsConfig)

	// define http server and server handler
	mux := http.NewServeMux()
	mux.HandleFunc("/mutate", whsvr.Serve)
	whsvr.Server.Handler = mux

	// start webhook server in new rountine
	go func() {
		glog.Infof("Start Webhook Werver...")
		if err := whsvr.Server.ListenAndServeTLS("", ""); err != nil {
			glog.Errorf("Filed to listen and serve webhook server: %v", err)
		}
	}()

	// listening OS shutdown singal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	glog.Infof("Got OS shutdown signal, shutting down wenhook server gracefully...")
	err = whsvr.Server.Shutdown(context.Background())
	if err != nil {
		glog.Errorf("Webhook Serve shutdown error: %v", err)
	}
}
