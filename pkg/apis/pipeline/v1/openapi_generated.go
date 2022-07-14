//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Tekton Authors

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/tektoncd/pipeline/pkg/apis/pipeline/pod.AffinityAssistantTemplate": schema_pkg_apis_pipeline_pod_AffinityAssistantTemplate(ref),
		"github.com/tektoncd/pipeline/pkg/apis/pipeline/pod.Template":                  schema_pkg_apis_pipeline_pod_Template(ref),
		"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1.Task":                       schema_pkg_apis_pipeline_v1_Task(ref),
		"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1.TaskList":                   schema_pkg_apis_pipeline_v1_TaskList(ref),
		"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1.TaskSpec":                   schema_pkg_apis_pipeline_v1_TaskSpec(ref),
	}
}

func schema_pkg_apis_pipeline_pod_AffinityAssistantTemplate(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AffinityAssistantTemplate holds pod specific configuration and is a subset of the generic pod Template",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nodeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"tolerations": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "If specified, the pod's tolerations.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.Toleration"),
									},
								},
							},
						},
					},
					"imagePullSecrets": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "ImagePullSecrets gives the name of the secret used by the pod to pull the image if specified",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.LocalObjectReference"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.LocalObjectReference", "k8s.io/api/core/v1.Toleration"},
	}
}

func schema_pkg_apis_pipeline_pod_Template(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Template holds pod specific configuration",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nodeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"tolerations": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "If specified, the pod's tolerations.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.Toleration"),
									},
								},
							},
						},
					},
					"affinity": {
						SchemaProps: spec.SchemaProps{
							Description: "If specified, the pod's scheduling constraints",
							Ref:         ref("k8s.io/api/core/v1.Affinity"),
						},
					},
					"securityContext": {
						SchemaProps: spec.SchemaProps{
							Description: "SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty.  See type description for default values of each field.",
							Ref:         ref("k8s.io/api/core/v1.PodSecurityContext"),
						},
					},
					"volumes": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type":       "atomic",
								"x-kubernetes-patch-merge-key": "name",
								"x-kubernetes-patch-strategy":  "merge,retainKeys",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.Volume"),
									},
								},
							},
						},
					},
					"runtimeClassName": {
						SchemaProps: spec.SchemaProps{
							Description: "RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod. If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty, the \"legacy\" RuntimeClass will be used, which is an implicit class with an empty definition that uses the default runtime handler. More info: https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md This is a beta feature as of Kubernetes v1.14.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"automountServiceAccountToken": {
						SchemaProps: spec.SchemaProps{
							Description: "AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"dnsPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "Set DNS policy for the pod. Defaults to \"ClusterFirst\". Valid values are 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dnsConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",
							Ref:         ref("k8s.io/api/core/v1.PodDNSConfig"),
						},
					},
					"enableServiceLinks": {
						SchemaProps: spec.SchemaProps{
							Description: "EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"priorityClassName": {
						SchemaProps: spec.SchemaProps{
							Description: "If specified, indicates the pod's priority. \"system-node-critical\" and \"system-cluster-critical\" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"schedulerName": {
						SchemaProps: spec.SchemaProps{
							Description: "SchedulerName specifies the scheduler to be used to dispatch the Pod",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imagePullSecrets": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "ImagePullSecrets gives the name of the secret used by the pod to pull the image if specified",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.LocalObjectReference"),
									},
								},
							},
						},
					},
					"hostAliases": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file if specified. This is only valid for non-hostNetwork pods.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.HostAlias"),
									},
								},
							},
						},
					},
					"hostNetwork": {
						SchemaProps: spec.SchemaProps{
							Description: "HostNetwork specifies whether the pod may use the node network namespace",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.Affinity", "k8s.io/api/core/v1.HostAlias", "k8s.io/api/core/v1.LocalObjectReference", "k8s.io/api/core/v1.PodDNSConfig", "k8s.io/api/core/v1.PodSecurityContext", "k8s.io/api/core/v1.Toleration", "k8s.io/api/core/v1.Volume"},
	}
}

func schema_pkg_apis_pipeline_v1_Task(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Task represents a collection of sequential steps that are run as part of a Pipeline using a set of inputs and producing a set of outputs. Tasks execute when TaskRuns are created that provide the input parameters and resources and output resources the Task requires.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Spec holds the desired state of the Task from the client",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/tektoncd/pipeline/pkg/apis/pipeline/v1.TaskSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1.TaskSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_pipeline_v1_TaskList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TaskList contains a list of Task",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/tektoncd/pipeline/pkg/apis/pipeline/v1.Task"),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1.Task", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_pipeline_v1_TaskSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TaskSpec defines the desired state of Task.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description is a user-facing description of the task that may be used to populate a UI.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}