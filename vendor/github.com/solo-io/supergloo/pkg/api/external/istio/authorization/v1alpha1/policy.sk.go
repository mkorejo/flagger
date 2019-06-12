// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"sort"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/hashutils"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewPolicy(namespace, name string) *Policy {
	policy := &Policy{}
	policy.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return policy
}

func (r *Policy) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *Policy) SetStatus(status core.Status) {
	r.Status = status
}

func (r *Policy) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.Targets,
		r.Peers,
		r.PeerIsOptional,
		r.Origins,
		r.OriginIsOptional,
		r.PrincipalBinding,
	)
}

type PolicyList []*Policy
type PoliciesByNamespace map[string]PolicyList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list PolicyList) Find(namespace, name string) (*Policy, error) {
	for _, policy := range list {
		if policy.GetMetadata().Name == name {
			if namespace == "" || policy.GetMetadata().Namespace == namespace {
				return policy, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find policy %v.%v", namespace, name)
}

func (list PolicyList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, policy := range list {
		ress = append(ress, policy)
	}
	return ress
}

func (list PolicyList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, policy := range list {
		ress = append(ress, policy)
	}
	return ress
}

func (list PolicyList) Names() []string {
	var names []string
	for _, policy := range list {
		names = append(names, policy.GetMetadata().Name)
	}
	return names
}

func (list PolicyList) NamespacesDotNames() []string {
	var names []string
	for _, policy := range list {
		names = append(names, policy.GetMetadata().Namespace+"."+policy.GetMetadata().Name)
	}
	return names
}

func (list PolicyList) Sort() PolicyList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list PolicyList) Clone() PolicyList {
	var policyList PolicyList
	for _, policy := range list {
		policyList = append(policyList, resources.Clone(policy).(*Policy))
	}
	return policyList
}

func (list PolicyList) Each(f func(element *Policy)) {
	for _, policy := range list {
		f(policy)
	}
}

func (list PolicyList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Policy) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

func (byNamespace PoliciesByNamespace) Add(policy ...*Policy) {
	for _, item := range policy {
		byNamespace[item.GetMetadata().Namespace] = append(byNamespace[item.GetMetadata().Namespace], item)
	}
}

func (byNamespace PoliciesByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace PoliciesByNamespace) List() PolicyList {
	var list PolicyList
	for _, policyList := range byNamespace {
		list = append(list, policyList...)
	}
	return list.Sort()
}

func (byNamespace PoliciesByNamespace) Clone() PoliciesByNamespace {
	cloned := make(PoliciesByNamespace)
	for ns, list := range byNamespace {
		cloned[ns] = list.Clone()
	}
	return cloned
}

var _ resources.Resource = &Policy{}

// Kubernetes Adapter for Policy

func (o *Policy) GetObjectKind() schema.ObjectKind {
	t := PolicyCrd.TypeMeta()
	return &t
}

func (o *Policy) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Policy)
}

var PolicyCrd = crd.NewCrd("authentication.istio.io",
	"policies",
	"authentication.istio.io",
	"v1alpha1",
	"Policy",
	"policy",
	false,
	&Policy{})