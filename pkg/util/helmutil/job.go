package helmutil

import(
	batchv1 "k8s.io/api/batch/v1"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	corev1 "k8s.io/api/core/v1"
)

type TillerInitOptions struct {
	Name string
	Namespace string
	ServiceAccount string
	Version string
	BackoffLimit int32
}

func GenerateTillerInitJob(options TillerInitOptions) batchv1.Job {
	jobSpec := batchv1.JobSpec{
		Template: corev1.PodTemplateSpec{
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:  "helm",
						Image: "quay.io/venezia/helm:" + options.Version,
						Command: []string{
							"/helm",
							"init",
							"--force-upgrade",
							"--tiller-namespace",
							options.Namespace,
							"--service-account",
							options.ServiceAccount,
							"--skip-refresh",
						},
					},
				},
				RestartPolicy: corev1.RestartPolicyOnFailure,
				ServiceAccountName: options.ServiceAccount,
			},
		},
		BackoffLimit: &options.BackoffLimit,
	}
	return k8sutil.GenerateJob(options.Name, jobSpec)
}
