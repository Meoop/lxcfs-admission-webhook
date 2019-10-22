package admission

import (
	"github.com/golang/glog"

	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	mutatingWebhookConfigurationName = "lxcfs-admission-webhook"
	mutatingWebhookName              = "lxcfs-admission-webhook.caicloud.io"
)

// CreateMutatingWebhookConfiguration create mutating webhook configure
func CreateMutatingWebhookConfiguration(kubeconfig string, caBundle []byte) error {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		glog.Errorf("build config from flags error: %v", err)
		return err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		glog.Errorf("create clientset from config error: %v", err)
		return err
	}

	path := "/mutate"
	failurePolicy := admissionregistrationv1beta1.Fail
	webhookConfiguration := &admissionregistrationv1beta1.MutatingWebhookConfiguration{
		ObjectMeta: metav1.ObjectMeta{
			Name: mutatingWebhookConfigurationName,
			Labels: map[string]string{
				"app": "lxcfs-admission-webhook",
			},
		},
		Webhooks: []admissionregistrationv1beta1.Webhook{
			{
				Name: mutatingWebhookName,
				ClientConfig: admissionregistrationv1beta1.WebhookClientConfig{
					Service: &admissionregistrationv1beta1.ServiceReference{
						Name:      "lxcfs-admission-webhook",
						Namespace: "kube-system",
						Path:      &path,
					},
					CABundle: caBundle,
				},
				FailurePolicy: &failurePolicy,
				Rules: []admissionregistrationv1beta1.RuleWithOperations{
					{
						Operations: []admissionregistrationv1beta1.OperationType{
							admissionregistrationv1beta1.Create,
						},
						Rule: admissionregistrationv1beta1.Rule{
							APIGroups:   []string{"core", ""},
							APIVersions: []string{"v1"},
							Resources:   []string{"pods"},
						},
					},
				},
			},
		},
	}

	_, err = clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().Create(webhookConfiguration)
	if k8serrors.IsAlreadyExists(err) {
		for i := 0; i < 5; i++ {
			webhook, err := clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().Get(webhookConfiguration.Name, metav1.GetOptions{})
			if err != nil {
				return err
			}
			webhook.Webhooks = webhookConfiguration.Webhooks
			_, err = clientset.AdmissionregistrationV1beta1().MutatingWebhookConfigurations().Update(webhook)
			if err != nil {
				if !k8serrors.IsConflict(err) {
					return err
				}
				continue
			}
			return nil
		}
	}
	return nil
}
