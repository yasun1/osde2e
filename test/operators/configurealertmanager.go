package operators

import (
	"github.com/onsi/ginkgo"
	"github.com/openshift/osde2e/pkg/helper"
)

var _ = ginkgo.Describe("[OSD] Configure AlertManager Operator", func() {
	var operatorName = "configure-alertmanager-operator"
	var operatorNamespace string = "openshift-monitoring"
	var operatorLockFile string = "configure-alertmanager-operator-lock"
	var defaultDesiredReplicas int32 = 1

	// NOTE: CAM clusterRoles have random-ish names like:
	// configure-alertmanager-operator.v0.1.80-03136c1-l589
	//
	// Test need to incorporate a regex-like test?
	//
	// var clusterRoles = []string{
	// 	"configure-alertmanager-operator",
	// }

	var clusterRoleBindings = []string{}

	var roleBindings = []string{
		"configure-alertmanager-operator",
	}

	var roles = []string{
		"configure-alertmanager-operator",
	}

	h := helper.New()
	checkClusterServiceVersion(h, operatorNamespace, operatorName)
	checkConfigMapLockfile(h, operatorNamespace, operatorLockFile)
	checkDeployment(h, operatorNamespace, operatorName, defaultDesiredReplicas)
	checkClusterRoleBindings(h, clusterRoleBindings)
	checkRole(h, operatorNamespace, roles)
	checkRoleBindings(h, operatorNamespace, roleBindings)
})
