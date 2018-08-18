package k8sutil

import (
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
)

// IsResourceAlreadyExistsError determines if error is an already exist type
func IsResourceAlreadyExistsError(err error) bool {
	return k8serrors.IsAlreadyExists(err)
}

// IsResourceNotFoundError determines if error is of type not found
func IsResourceNotFoundError(err error) bool {
	return k8serrors.IsNotFound(err)
}
