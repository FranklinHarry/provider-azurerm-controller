/*
Copyright AppsCode Inc. and Contributors

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

package scheme

import (
	domainv1alpha1 "kubeform.dev/provider-azurerm-api/apis/domain/v1alpha1"
	firewallv1alpha1 "kubeform.dev/provider-azurerm-api/apis/firewall/v1alpha1"
	imagev1alpha1 "kubeform.dev/provider-azurerm-api/apis/image/v1alpha1"
	instancev1alpha1 "kubeform.dev/provider-azurerm-api/apis/instance/v1alpha1"
	lkev1alpha1 "kubeform.dev/provider-azurerm-api/apis/lke/v1alpha1"
	nodebalancerv1alpha1 "kubeform.dev/provider-azurerm-api/apis/nodebalancer/v1alpha1"
	objectv1alpha1 "kubeform.dev/provider-azurerm-api/apis/object/v1alpha1"
	rdnsv1alpha1 "kubeform.dev/provider-azurerm-api/apis/rdns/v1alpha1"
	sshkeyv1alpha1 "kubeform.dev/provider-azurerm-api/apis/sshkey/v1alpha1"
	stackscriptv1alpha1 "kubeform.dev/provider-azurerm-api/apis/stackscript/v1alpha1"
	tokenv1alpha1 "kubeform.dev/provider-azurerm-api/apis/token/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-azurerm-api/apis/user/v1alpha1"
	vlanv1alpha1 "kubeform.dev/provider-azurerm-api/apis/vlan/v1alpha1"
	volumev1alpha1 "kubeform.dev/provider-azurerm-api/apis/volume/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)
var localSchemeBuilder = runtime.SchemeBuilder{
	domainv1alpha1.AddToScheme,
	firewallv1alpha1.AddToScheme,
	imagev1alpha1.AddToScheme,
	instancev1alpha1.AddToScheme,
	lkev1alpha1.AddToScheme,
	nodebalancerv1alpha1.AddToScheme,
	objectv1alpha1.AddToScheme,
	rdnsv1alpha1.AddToScheme,
	sshkeyv1alpha1.AddToScheme,
	stackscriptv1alpha1.AddToScheme,
	tokenv1alpha1.AddToScheme,
	userv1alpha1.AddToScheme,
	vlanv1alpha1.AddToScheme,
	volumev1alpha1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   _ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(Scheme))
}
