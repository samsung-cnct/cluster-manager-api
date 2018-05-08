package sdscluster

import (
	api "github.com/samsung-cnct/cluster-manager-api/pkg/apis/sdscluster/v1alpha1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateCRD creates the objects in kubernetes
func GenerateCRD() apiextensionsv1beta1.CustomResourceDefinition {
	crd := apiextensionsv1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: api.SDSClusterCRDName,
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   api.SchemeGroupVersion.Group,
			Version: api.SchemeGroupVersion.Version,
			Scope:   apiextensionsv1beta1.NamespaceScoped,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Plural: api.SDSClusterResourcePlural,
				Kind:   api.SDSClusterResourceKind,
			},
		},
	}

	return crd
}