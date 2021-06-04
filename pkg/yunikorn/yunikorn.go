/*
Copyright 2020 The Kubernetes Authors.

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

package yunikorn

import (
	"context"
	"fmt"
	"sync"

	"k8s.io/api/core/v1"
	"k8s.io/klog/v2"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type YuniKorn struct {
	sync.RWMutex
	frameworkHandle framework.Handle
}

var _ framework.PreFilterPlugin = &YuniKorn{}

const (
	Name              = "YuniKornSchedulerPlugin"
	preFilterStateKey = "PreFilter" + Name
)

// Name returns name of the plugin. It is used in logs, etc.
func (y *YuniKorn) Name() string {
	return Name
}

// New initializes a new plugin and returns it.
func New(obj runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	y := &YuniKorn{
		frameworkHandle: handle,
	}

	klog.Infof("YuniKornSchedulingPlugin started")
	return y, nil
}

func (y *YuniKorn) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) *framework.Status {
	return framework.NewStatus(framework.Unschedulable, fmt.Sprintf("Pod %v/%v is rejected in Prefilter - this is a test", pod.Namespace, pod.Name))
}

func (y *YuniKorn) PreFilterExtensions() framework.PreFilterExtensions {
	return nil
}
