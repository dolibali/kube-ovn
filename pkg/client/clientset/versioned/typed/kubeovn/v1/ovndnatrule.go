/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	kubeovnv1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	scheme "github.com/kubeovn/kube-ovn/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// OvnDnatRulesGetter has a method to return a OvnDnatRuleInterface.
// A group's client should implement this interface.
type OvnDnatRulesGetter interface {
	OvnDnatRules() OvnDnatRuleInterface
}

// OvnDnatRuleInterface has methods to work with OvnDnatRule resources.
type OvnDnatRuleInterface interface {
	Create(ctx context.Context, ovnDnatRule *kubeovnv1.OvnDnatRule, opts metav1.CreateOptions) (*kubeovnv1.OvnDnatRule, error)
	Update(ctx context.Context, ovnDnatRule *kubeovnv1.OvnDnatRule, opts metav1.UpdateOptions) (*kubeovnv1.OvnDnatRule, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, ovnDnatRule *kubeovnv1.OvnDnatRule, opts metav1.UpdateOptions) (*kubeovnv1.OvnDnatRule, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubeovnv1.OvnDnatRule, error)
	List(ctx context.Context, opts metav1.ListOptions) (*kubeovnv1.OvnDnatRuleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *kubeovnv1.OvnDnatRule, err error)
	OvnDnatRuleExpansion
}

// ovnDnatRules implements OvnDnatRuleInterface
type ovnDnatRules struct {
	*gentype.ClientWithList[*kubeovnv1.OvnDnatRule, *kubeovnv1.OvnDnatRuleList]
}

// newOvnDnatRules returns a OvnDnatRules
func newOvnDnatRules(c *KubeovnV1Client) *ovnDnatRules {
	return &ovnDnatRules{
		gentype.NewClientWithList[*kubeovnv1.OvnDnatRule, *kubeovnv1.OvnDnatRuleList](
			"ovn-dnat-rules",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *kubeovnv1.OvnDnatRule { return &kubeovnv1.OvnDnatRule{} },
			func() *kubeovnv1.OvnDnatRuleList { return &kubeovnv1.OvnDnatRuleList{} },
		),
	}
}
