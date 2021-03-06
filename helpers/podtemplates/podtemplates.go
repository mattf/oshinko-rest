package podtemplates

import (
	"github.com/redhatanalytics/oshinko-rest/helpers/containers"
	kapi "k8s.io/kubernetes/pkg/api"
)

type OPodTemplateSpec struct {
	kapi.PodTemplateSpec
}

// we might care about volumes
// we might care about terminationGracePeriodSeconds
// we might care about serviceaccountname
// we might care about security context
// we might care about image pull secrets

func PodTemplateSpec() *OPodTemplateSpec {
	// Note, name and namespace can be set on a PodTemplateSpec but
	// I assume that openshift takes care of that based on the DeploymentConfig
	p := OPodTemplateSpec{}
	p.Spec.DNSPolicy = kapi.DNSClusterFirst
	p.Spec.RestartPolicy = kapi.RestartPolicyAlways
	return &p
}

func (pt *OPodTemplateSpec) SetLabels(selectors map[string]string) *OPodTemplateSpec {
	pt.PodTemplateSpec.SetLabels(selectors)
	return pt
}

func (pt *OPodTemplateSpec) Label(name, value string) *OPodTemplateSpec {
	if pt.Labels == nil {
		pt.Labels = map[string]string{}
	}
	pt.Labels[name] = value
	return pt
}

func (pt *OPodTemplateSpec) Containers(cntnrs ...*containers.OContainer) *OPodTemplateSpec {
	kcntnrs := make([]kapi.Container, len(cntnrs))
	for idx, c := range cntnrs {
		kcntnrs[idx] = c.Container
	}
	pt.Spec.Containers = kcntnrs
	return pt
}
