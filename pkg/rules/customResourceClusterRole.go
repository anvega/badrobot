// OPR-R20-RBAC - ClusterRole has full permissions over any custom resource definitions
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func CustomResourceClusterRole(json []byte) int {
	rbac := 0

	jqAPI := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules").
		Only("apiGroups")

	jqResources := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules").
		Only("resources")

	jqVerbs := gojsonq.New().Reader(bytes.NewReader(json)).
		From("rules").
		Only("verbs")

	if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[apiextensions.k8s.io]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[customresourcedefinitions]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[apiextensions.k8s.io]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[customresourcedefinitions]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "create")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[apiextensions.k8s.io]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[customresourcedefinitions]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "patch")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[apiextensions.k8s.io]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[customresourcedefinitions]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "update")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[apiextensions.k8s.io]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[customresourcedefinitions]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "delete")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[apiextensions.k8s.io]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[customresourcedefinitions]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "deletecollection")) {
		rbac++
	}

	return rbac

}
