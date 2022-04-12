// OPR-R17-RBAC - ClusterRole has impersonate permissions
package rules

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func ServiceAccountClusterRole(json []byte) int {
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

	if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[serviceaccounts]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[serviceaccounts]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "create")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[serviceaccounts/token]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "*")) {
		rbac++
	} else if (strings.Contains(fmt.Sprintf("%v", jqAPI), "[]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqResources), "[serviceaccounts/token]")) &&
		(strings.Contains(fmt.Sprintf("%v", jqVerbs), "create")) {
		rbac++
	}

	return rbac

}
