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

package helloworld

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
        "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
        "k8s.io/kubernetes/pkg/scheduler/framework"
	"sigs.k8s.io/scheduler-plugins/pkg/apis/config"
)

// Helloworld is a plugin that schedules pods in a group.
type Helloworld struct {
	handle framework.Handle
}

var _ framework.PreBindPlugin = &Helloworld{}

const (
	// Name is the name of the plugin used in Registry and configurations.
	Name = "HelloWorld"
)

func New(obj runtime.Object, h framework.Handle) (framework.Plugin, error) {
	args, ok := obj.(*config.HelloWorldArgs)
	if !ok {
		return nil, fmt.Errorf("[Helloworld] want args to be of type HelloworldArgs, got %T", obj)
	}

	klog.Infof("[Helloworld] args received. String: %s", args.Names)

	return &Helloworld{
		handle: h,
	}, nil
}

func (n *Helloworld) Name() string {
	return Name
}

func (n *Helloworld) PreBind(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) (*framework.Status) {
	klog.Infof("[Helloworld] tupt18 reach this %s node", nodeName)
	return nil
}
