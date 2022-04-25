package rules

import (
	"testing"

	"github.com/ghodss/yaml"
)

func Test_PVC_All_Permissions(t *testing.T) {
	var data = `
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: example-operator
rules:
- apiGroups:
  - ""
  resources:
  - persistentvolumeclaims
  verbs:
  - "*"
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac := PersistentVolumeClusterRole(json)
	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}

func Test_Incorrect_PVC_Permissions(t *testing.T) {
	var data = `
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: example-operator
rules:
- apiGroups:
  - authorization.k8s.io
  resources:
  - persistentvolumeclaims
  verbs:
  - "*"
`

	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac := PersistentVolumeClusterRole(json)
	if rbac != 0 {
		t.Errorf("Got %v permissions wanted %v", rbac, 0)
	}
}

func Test_PVC_Verbs_Permissions(t *testing.T) {
	var data = `
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: example-operator
rules:
- apiGroups:
  - ""
  resources:
  - persistentvolumeclaims
  verbs:
  - delete
  - deletecollection
  - create
  - patch
  - get
  - list
  - update
  - watch
`
	json, err := yaml.YAMLToJSON([]byte(data))
	if err != nil {
		t.Fatal(err.Error())
	}

	rbac := PersistentVolumeClusterRole(json)
	if rbac != 1 {
		t.Errorf("Got %v permissions wanted %v", rbac, 1)
	}
}
