/*
Copyright 2023.

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

package support

import (
	"github.com/onsi/gomega"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Ingress(t Test, namespace, name string) func(g gomega.Gomega) *networkingv1.Ingress {
	return func(g gomega.Gomega) *networkingv1.Ingress {
		ingress, err := t.Client().Core().NetworkingV1().Ingresses(namespace).Get(t.Ctx(), name, metav1.GetOptions{})
		g.Expect(err).NotTo(gomega.HaveOccurred())
		return ingress
	}
}

func GetIngress(t Test, namespace, name string) *networkingv1.Ingress {
	t.T().Helper()
	return Ingress(t, namespace, name)(t)
}

func LoadBalancerIngresses(ingress *networkingv1.Ingress) []networkingv1.IngressLoadBalancerIngress {
	return ingress.Status.LoadBalancer.Ingress
}
