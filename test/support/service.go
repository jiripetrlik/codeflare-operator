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

import "net/url"

func ExposeService(t Test, name string, namespace string, serviceName string, servicePort string) url.URL {
	if IsOpenShift(t) {
		route := ExposeServiceRoute(t, name, namespace, serviceName, servicePort)
		route = GetRoute(t, route.Namespace, route.Name)

		serviceURL := url.URL{
			Scheme: "http",
			Host:   route.Status.Ingress[0].Host,
		}

		return serviceURL
	} else {
		ingress := ExposeServiceIngress(t, name, namespace, serviceName, servicePort)
		serviceURL := url.URL{
			Scheme: "http",
			Host:   ingress.Status.LoadBalancer.Ingress[0].IP,
			Path:   name,
		}
		return serviceURL
	}
}
