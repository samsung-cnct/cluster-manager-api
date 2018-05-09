package k8sutil

import (
	"fmt"
	"log"
	"time"

	"github.com/samsung-cnct/cluster-manager-api/pkg/util/retryutil"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateCRD creates the objects in kubernetes
func CreateCRD(clientset apiextensionsclient.Interface, crd apiextensionsv1beta1.CustomResourceDefinition) error {
	_, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(&crd)
	if err != nil && !IsResourceAlreadyExistsError(err) {
		log.Printf("Error %v\n", err)
		return err
	} else if IsResourceAlreadyExistsError(err) {
		log.Printf("CRD %s already exists, need to work on a patch/update\n", crd.Name)
	}
	return nil
}

// WaitCRDReady waits until proper condition is obtained.
func WaitCRDReady(clientset apiextensionsclient.Interface, crdName string) error {
	err := retryutil.Retry(5*time.Second, 20, func() (bool, error) {
		crd, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(crdName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		for _, cond := range crd.Status.Conditions {
			switch cond.Type {
			case apiextensionsv1beta1.Established:
				if cond.Status == apiextensionsv1beta1.ConditionTrue {
					return true, nil
				}
			case apiextensionsv1beta1.NamesAccepted:
				if cond.Status == apiextensionsv1beta1.ConditionFalse {
					return false, fmt.Errorf("Name conflict: %v", cond.Reason)
				}
			}
		}
		return false, nil
	})
	if err != nil {
		return fmt.Errorf("wait CRD created failed: %v", err)
	}
	return nil
}
