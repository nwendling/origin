package v1beta3

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/conversion"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"

	newer "github.com/openshift/origin/pkg/deploy/api"
	imageapi "github.com/openshift/origin/pkg/image/api"
)

func init() {
	err := api.Scheme.AddDefaultingFuncs(
		func(obj *DeploymentConfigSpec) {
			if len(obj.Strategy.Type) == 0 && util.AllPtrFieldsNil(obj.Strategy) {
				obj.Strategy.Type = DeploymentStrategyTypeRecreate
			}
		},
		func(obj *DeploymentTriggerImageChangeParams) {
			if len(obj.From.Kind) == 0 {
				obj.From.Kind = "ImageStreamTag"
			}
		},
	)
	if err != nil {
		panic(err)
	}

	err = api.Scheme.AddConversionFuncs(
		func(in *DeploymentConfig, out *newer.DeploymentConfig, s conversion.Scope) error {
			if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields); err != nil {
				return err
			}
			if err := s.Convert(&in.Spec, &out.Template, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Spec.Triggers, &out.Triggers, 0); err != nil {
				return err
			}
			out.LatestVersion = in.Status.LatestVersion
			if err := s.Convert(&in.Status.Details, &out.Details, 0); err != nil {
				return err
			}
			return nil
		},
		func(in *newer.DeploymentConfig, out *DeploymentConfig, s conversion.Scope) error {
			if err := s.DefaultConvert(in, out, conversion.IgnoreMissingFields); err != nil {
				return err
			}
			if err := s.Convert(&in.Template, &out.Spec, 0); err != nil {
				return err
			}
			if err := s.Convert(&in.Triggers, &out.Spec.Triggers, 0); err != nil {
				return err
			}
			out.Status.LatestVersion = in.LatestVersion
			if err := s.Convert(&in.Details, &out.Status.Details, 0); err != nil {
				return err
			}
			return nil
		},

		func(in *DeploymentConfigSpec, out *newer.DeploymentTemplate, s conversion.Scope) error {
			out.ControllerTemplate.Replicas = in.Replicas
			if in.Selector != nil {
				out.ControllerTemplate.Selector = make(map[string]string)
				for k, v := range in.Selector {
					out.ControllerTemplate.Selector[k] = v
				}
			}
			if in.Template != nil {
				if err := s.Convert(&in.Template, &out.ControllerTemplate.Template, 0); err != nil {
					return err
				}
			}
			if in.TemplateRef != nil {
				if err := s.Convert(&in.TemplateRef, &out.ControllerTemplate.TemplateRef, 0); err != nil {
					return err
				}
			}
			if err := s.Convert(&in.Strategy, &out.Strategy, 0); err != nil {
				return err
			}
			return nil
		},
		func(in *newer.DeploymentTemplate, out *DeploymentConfigSpec, s conversion.Scope) error {
			out.Replicas = in.ControllerTemplate.Replicas
			if in.ControllerTemplate.Selector != nil {
				out.Selector = make(map[string]string)
				for k, v := range in.ControllerTemplate.Selector {
					out.Selector[k] = v
				}
			}
			if in.ControllerTemplate.Template != nil {
				if err := s.Convert(&in.ControllerTemplate.Template, &out.Template, 0); err != nil {
					return err
				}
			}
			if in.ControllerTemplate.TemplateRef != nil {
				if err := s.Convert(&in.ControllerTemplate.TemplateRef, &out.TemplateRef, 0); err != nil {
					return err
				}
			}
			if err := s.Convert(&in.Strategy, &out.Strategy, 0); err != nil {
				return err
			}
			return nil
		},

		func(in *DeploymentTriggerImageChangeParams, out *newer.DeploymentTriggerImageChangeParams, s conversion.Scope) error {
			out.Automatic = in.Automatic
			out.ContainerNames = make([]string, len(in.ContainerNames))
			copy(out.ContainerNames, in.ContainerNames)
			out.LastTriggeredImage = in.LastTriggeredImage
			if err := s.Convert(&in.From, &out.From, 0); err != nil {
				return err
			}
			switch in.From.Kind {
			case "DockerImage":
				ref, err := imageapi.ParseDockerImageReference(in.From.Name)
				if err != nil {
					return err
				}
				out.Tag = ref.Tag
				ref.Tag, ref.ID = "", ""
				out.RepositoryName = ref.String()
			}
			return nil
		},
		func(in *newer.DeploymentTriggerImageChangeParams, out *DeploymentTriggerImageChangeParams, s conversion.Scope) error {
			out.Automatic = in.Automatic
			out.ContainerNames = make([]string, len(in.ContainerNames))
			copy(out.ContainerNames, in.ContainerNames)
			out.LastTriggeredImage = in.LastTriggeredImage
			if err := s.Convert(&in.From, &out.From, 0); err != nil {
				return err
			}
			return nil
		},

		func(in *DeploymentCauseImageTrigger, out *newer.DeploymentCauseImageTrigger, s conversion.Scope) error {
			switch in.From.Kind {
			case "DockerImage":
				ref, err := imageapi.ParseDockerImageReference(in.From.Name)
				if err != nil {
					return err
				}
				out.Tag = ref.Tag
				ref.Tag, ref.ID = "", ""
				out.RepositoryName = ref.String()
			}
			return nil
		},
		func(in *newer.DeploymentCauseImageTrigger, out *DeploymentCauseImageTrigger, s conversion.Scope) error {
			if len(in.RepositoryName) != 0 {
				ref, err := imageapi.ParseDockerImageReference(in.RepositoryName)
				if err != nil {
					return err
				}
				ref.Tag = in.Tag
				out.From.Kind = "DockerImage"
				out.From.Name = ref.Minimal().String()
			}
			return nil
		},
	)
	if err != nil {
		panic(err)
	}
}
