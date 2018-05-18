package cma

import (
	api "github.com/samsung-cnct/cluster-manager-api/pkg/apis/cma/v1alpha1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GenerateSDSClusterCRD creates the SDSCluster CRD object
func GenerateSDSClusterCRD() apiextensionsv1beta1.CustomResourceDefinition {
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

// GenerateSDSPackageManagerCRD creates the SDSPackageManager CRD object
func GenerateSDSPackageManagerCRD() apiextensionsv1beta1.CustomResourceDefinition {
	crd := apiextensionsv1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: api.SDSPackageManagerCRDName,
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   api.SchemeGroupVersion.Group,
			Version: api.SchemeGroupVersion.Version,
			Scope:   apiextensionsv1beta1.NamespaceScoped,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Plural: api.SDSPackageManagerResourcePlural,
				Kind:   api.SDSPackageManagerResourceKind,
			},
		},
	}

	return crd
}

// GenerateSDSApplicationCRD creates the SDSApplication CRD object
func GenerateSDSApplicationCRD() apiextensionsv1beta1.CustomResourceDefinition {
	crd := apiextensionsv1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: api.SDSApplicationCRDName,
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   api.SchemeGroupVersion.Group,
			Version: api.SchemeGroupVersion.Version,
			Scope:   apiextensionsv1beta1.NamespaceScoped,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Plural: api.SDSApplicationResourcePlural,
				Kind:   api.SDSApplicationResourceKind,
			},
		},
	}

	return crd
}
