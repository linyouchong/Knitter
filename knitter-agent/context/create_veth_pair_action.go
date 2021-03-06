/*
Copyright 2018 ZTE Corporation. All rights reserved.
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

package context

import (
	"github.com/ZTE/Knitter/knitter-agent/domain/object/os-obj"
	"github.com/ZTE/Knitter/pkg/klog"
	"github.com/ZTE/Knitter/pkg/trans-dsl"
)

type CreateVethPairAction struct {
}

func (this *CreateVethPairAction) Exec(transInfo *transdsl.TransInfo) (err error) {
	klog.Infof("***CreateVethPairAction:Exec begin***")
	defer func() {
		if p := recover(); p != nil {
			RecoverErr(p, &err, "CreateVethPairAction")
		}
		AppendActionName(&err, "CreateVethPairAction")
	}()

	osObj := osobj.GetOsObjSingleton()
	vethPair, err := osObj.VethPairRole.Create()
	if err != nil {
		klog.Errorf("CreateVethPairAction.Exec: osObj.VethPairRole.Create error: %v", err)
		return err
	}
	transInfo.AppInfo.(*KnitterInfo).vethPair = &vethPair
	klog.Infof("***CreateVethPairAction:Exec end***")
	return nil
}

func (this *CreateVethPairAction) RollBack(transInfo *transdsl.TransInfo) {
	klog.Infof("***CreateVethPairAction:RollBack begin***")
	osObj := osobj.GetOsObjSingleton()
	osObj.VethPairRole.Destroy(transInfo.AppInfo.(*KnitterInfo).vethPair.VethNameOfBridge)
	klog.Infof("***CreateVethPairAction:RollBack end***")
}
