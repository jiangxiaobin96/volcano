/*
 Copyright 2023 The Volcano Authors.

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

package api

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	"volcano.sh/volcano/pkg/scheduler/api/devices/nvidia/gpushare"
)

const (
	GPUSharingDevice = "GpuShare"
)

type Devices interface {
	//following two functions used in node_info
	//AddResource is to add the corresponding device resource of this 'pod' into current scheduler cache
	AddResource(pod *v1.Pod)
	//SubResoure is to substract the corresponding device resource of this 'pod' from current scheduler cache
	SubResource(pod *v1.Pod)

	//following four functions used in predicate
	//HasDeviceRequest checks if the 'pod' request this device
	HasDeviceRequest(pod *v1.Pod) bool
	//FiltreNode checks if the 'pod' fit in current node
	FilterNode(pod *v1.Pod) (bool, error)
	//Allocate action in predicate
	Allocate(kubeClient kubernetes.Interface, pod *v1.Pod) error
	//Release action in predicate
	Release(kubeClient kubernetes.Interface, pod *v1.Pod) error

	//used for debug and monitor
	GetStatus() string
}

// make sure GPUDevices implements Devices interface
var _ Devices = new(gpushare.GPUDevices)
