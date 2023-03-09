package gangextender

import (
	"k8s.io/klog/v2"
	"volcano.sh/volcano/pkg/scheduler/api"
)

const (
	sparkPodLabelRoleKey         = "spark-role"
	sparkPodLabelRoleDriverValue = "driver"
)

func GangRejectExtender(preemptee *api.TaskInfo) bool {
	sparkReject := sparkGangExtender(preemptee)
	return sparkReject
}

func sparkGangExtender(preemptee *api.TaskInfo) bool {
	if value, ok := preemptee.Pod.Labels[sparkPodLabelRoleKey]; ok {
		if value == sparkPodLabelRoleDriverValue {
			klog.V(4).Infof("Failed to preempt task <%v/%v> because this task is a driver of spark job",
				preemptee.Namespace, preemptee.Name)
			return true
		} else {
			return false
		}
	}
	return false
}
