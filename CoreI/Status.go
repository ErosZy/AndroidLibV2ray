package CoreI

import (
	"v2ray.com/core"
)

type Status struct {
	IsRunning       bool
	VpnSupportnodup bool
	PackageName     string
	DataDir 		string
	Vpoint core.Server
}

func (v *Status) SetDataDir(path string) {
	v.DataDir = path
}

func (v *Status) GetDataDir() string {
	if len(v.DataDir) == 0 {
		return v.getDataDir()
	}

	return v.DataDir;
}

func (v *Status) getDataDir() string {
	var datadir = "/data/data/org.kkdev.v2raygo/"
	if v.PackageName != "" {
		datadir = "/data/data/" + v.PackageName + "/"
	}
	return datadir
}
