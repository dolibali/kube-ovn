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

package fake

import (
	v1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	kubeovnv1 "github.com/kubeovn/kube-ovn/pkg/client/clientset/versioned/typed/kubeovn/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeIptablesFIPRules implements IptablesFIPRuleInterface
type fakeIptablesFIPRules struct {
	*gentype.FakeClientWithList[*v1.IptablesFIPRule, *v1.IptablesFIPRuleList]
	Fake *FakeKubeovnV1
}

func newFakeIptablesFIPRules(fake *FakeKubeovnV1) kubeovnv1.IptablesFIPRuleInterface {
	return &fakeIptablesFIPRules{
		gentype.NewFakeClientWithList[*v1.IptablesFIPRule, *v1.IptablesFIPRuleList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("iptables-fip-rules"),
			v1.SchemeGroupVersion.WithKind("IptablesFIPRule"),
			func() *v1.IptablesFIPRule { return &v1.IptablesFIPRule{} },
			func() *v1.IptablesFIPRuleList { return &v1.IptablesFIPRuleList{} },
			func(dst, src *v1.IptablesFIPRuleList) { dst.ListMeta = src.ListMeta },
			func(list *v1.IptablesFIPRuleList) []*v1.IptablesFIPRule { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.IptablesFIPRuleList, items []*v1.IptablesFIPRule) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
