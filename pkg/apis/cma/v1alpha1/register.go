package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	SDSClusterResourceKind          = "SDSCluster"
	SDSClusterResourcePlural        = "sdsclusters"
	SDSApplicationResourceKind      = "SDSApplication"
	SDSApplicationResourcePlural    = "sdsapplications"
	SDSPackageManagerResourceKind   = "SDSPackageManager"
	SDSPackageManagerResourcePlural = "sdspackagemanagers"
	groupName                       = "cma.sds.samsung.com"
)

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme

	SchemeGroupVersion       = schema.GroupVersion{Group: groupName, Version: "v1alpha1"}
	SDSClusterCRDName        = SDSClusterResourcePlural + "." + groupName
	SDSApplicationCRDName    = SDSApplicationResourcePlural + "." + groupName
	SDSPackageManagerCRDName = SDSPackageManagerResourcePlural + "." + groupName
)

// addKnownTypes adds the set of types defined in this package to the supplied scheme.
func addKnownTypes(s *runtime.Scheme) error {
	s.AddKnownTypes(SchemeGroupVersion,
		&SDSCluster{},
		&SDSClusterList{},
		&SDSApplication{},
		&SDSApplicationList{},
		&SDSPackageManager{},
		&SDSPackageManagerList{},
	)
	metav1.AddToGroupVersion(s, SchemeGroupVersion)
	return nil
}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return VersionKind(kind).GroupKind()
}

// VersionKind takes an unqualified kind and returns back a Group qualified GroupVersionKind
func VersionKind(kind string) schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind(kind)
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
