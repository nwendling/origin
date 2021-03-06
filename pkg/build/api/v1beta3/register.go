package v1beta3

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
)

func init() {
	api.Scheme.AddKnownTypes("v1beta3",
		&Build{},
		&BuildList{},
		&BuildConfig{},
		&BuildConfigList{},
		&BuildLog{},
		&BuildRequest{},
	)
}

func (*Build) IsAnAPIObject()           {}
func (*BuildList) IsAnAPIObject()       {}
func (*BuildConfig) IsAnAPIObject()     {}
func (*BuildConfigList) IsAnAPIObject() {}
func (*BuildLog) IsAnAPIObject()        {}
func (*BuildRequest) IsAnAPIObject()    {}
