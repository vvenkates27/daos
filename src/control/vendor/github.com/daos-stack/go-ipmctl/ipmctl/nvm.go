//
// (C) Copyright 2018 Intel Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// GOVERNMENT LICENSE RIGHTS-OPEN SOURCE SOFTWARE
// The Government's rights to use, modify, reproduce, release, perform, display,
// or disclose this software are subject to the terms of the Apache License as
// provided in Contract No. 8F-30005.
// Any reproduction of computer software, computer software documentation, or
// portions thereof marked with this legend must also reproduce the markings.
//

// Package ipmctl provides Go bindings for libipmctl Native Management API
package ipmctl

import "fmt"

// NvmMgmt is an implementation of the IpmCtl interface which exercises
// libipmctl's NVM API.
type NvmMgmt struct{}

// Discover queries number of SCM modules and retrieves device_discovery structs
// for each.
func (n *NvmMgmt) Discover() (devices []DeviceDiscovery, err error) {
	err = fmt.Errorf("ipmctl lib not present\n")
	return
}

// GetStatuses return status for each device in devices
func (n *NvmMgmt) GetStatuses(devices []DeviceDiscovery) (
	statuses []DeviceStatus, err error) {

	err = fmt.Errorf("ipmctl lib not present\n")
	return
}
