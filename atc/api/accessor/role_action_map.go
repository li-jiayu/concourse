package accessor

import (
	"io/ioutil"

	"code.cloudfoundry.org/lager"
	"gopkg.in/yaml.v2"

	"github.com/concourse/concourse/atc"
)

// requiredRoles should be a const, never be updated.
const System = "system"
const (
	MemberRole   = "member"
	OwnerRole    = "owner"
	OperatorRole = "pipeline-operator"
	ViewerRole   = "viewer"
)

var requiredRoles = map[string]string{
	atc.SaveConfig:                    MemberRole,
	atc.GetConfig:                     ViewerRole,
	atc.GetCC:                         ViewerRole,
	atc.GetBuild:                      ViewerRole,
	atc.GetCheck:                      ViewerRole,
	atc.GetBuildPlan:                  ViewerRole,
	atc.CreateBuild:                   MemberRole,
	atc.ListBuilds:                    ViewerRole,
	atc.BuildEvents:                   ViewerRole,
	atc.BuildResources:                ViewerRole,
	atc.AbortBuild:                    OperatorRole,
	atc.GetBuildPreparation:           ViewerRole,
	atc.GetJob:                        ViewerRole,
	atc.CreateJobBuild:                OperatorRole,
	atc.RerunJobBuild:                 OperatorRole,
	atc.ListAllJobs:                   ViewerRole,
	atc.ListJobs:                      ViewerRole,
	atc.ListJobBuilds:                 ViewerRole,
	atc.ListJobInputs:                 ViewerRole,
	atc.GetJobBuild:                   ViewerRole,
	atc.PauseJob:                      OperatorRole,
	atc.UnpauseJob:                    OperatorRole,
	atc.ScheduleJob:                   OperatorRole,
	atc.GetVersionsDB:                 ViewerRole,
	atc.JobBadge:                      ViewerRole,
	atc.MainJobBadge:                  ViewerRole,
	atc.ClearTaskCache:                OperatorRole,
	atc.ListAllResources:              ViewerRole,
	atc.ListResources:                 ViewerRole,
	atc.ListResourceTypes:             ViewerRole,
	atc.GetResource:                   ViewerRole,
	atc.UnpinResource:                 OperatorRole,
	atc.SetPinCommentOnResource:       OperatorRole,
	atc.CheckResource:                 OperatorRole,
	atc.CheckResourceWebHook:          OperatorRole,
	atc.CheckResourceType:             OperatorRole,
	atc.ListResourceVersions:          ViewerRole,
	atc.GetResourceVersion:            ViewerRole,
	atc.EnableResourceVersion:         OperatorRole,
	atc.DisableResourceVersion:        OperatorRole,
	atc.PinResourceVersion:            OperatorRole,
	atc.ListBuildsWithVersionAsInput:  ViewerRole,
	atc.ListBuildsWithVersionAsOutput: ViewerRole,
	atc.GetResourceCausality:          ViewerRole,
	atc.ListAllPipelines:              ViewerRole,
	atc.ListPipelines:                 ViewerRole,
	atc.GetPipeline:                   ViewerRole,
	atc.DeletePipeline:                MemberRole,
	atc.OrderPipelines:                MemberRole,
	atc.PausePipeline:                 OperatorRole,
	atc.UnpausePipeline:               OperatorRole,
	atc.ExposePipeline:                MemberRole,
	atc.HidePipeline:                  MemberRole,
	atc.RenamePipeline:                MemberRole,
	atc.ListPipelineBuilds:            ViewerRole,
	atc.CreatePipelineBuild:           MemberRole,
	atc.PipelineBadge:                 ViewerRole,
	atc.RegisterWorker:                MemberRole,
	atc.LandWorker:                    MemberRole,
	atc.RetireWorker:                  MemberRole,
	atc.PruneWorker:                   MemberRole,
	atc.HeartbeatWorker:               MemberRole,
	atc.ListWorkers:                   ViewerRole,
	atc.DeleteWorker:                  MemberRole,
	atc.SetLogLevel:                   MemberRole,
	atc.GetLogLevel:                   ViewerRole,
	atc.DownloadCLI:                   ViewerRole,
	atc.GetInfo:                       ViewerRole,
	atc.GetInfoCreds:                  ViewerRole,
	atc.ListContainers:                ViewerRole,
	atc.GetContainer:                  ViewerRole,
	atc.HijackContainer:               MemberRole,
	atc.ListDestroyingContainers:      ViewerRole,
	atc.ReportWorkerContainers:        MemberRole,
	atc.ListVolumes:                   ViewerRole,
	atc.ListDestroyingVolumes:         ViewerRole,
	atc.ReportWorkerVolumes:           MemberRole,
	atc.ListTeams:                     ViewerRole,
	atc.GetTeam:                       ViewerRole,
	atc.SetTeam:                       OwnerRole,
	atc.RenameTeam:                    OwnerRole,
	atc.DestroyTeam:                   OwnerRole,
	atc.ListTeamBuilds:                ViewerRole,
	atc.CreateArtifact:                MemberRole,
	atc.GetArtifact:                   MemberRole,
	atc.ListBuildArtifacts:            ViewerRole,
	atc.GetWall:                       ViewerRole,
}

type CustomActionRoleMap map[string][]string

//go:generate counterfeiter . ActionRoleMap

type ActionRoleMap interface {
	RoleOfAction(string) string
}

//go:generate counterfeiter . ActionRoleMapModifier

type ActionRoleMapModifier interface {
	CustomizeActionRoleMap(lager.Logger, CustomActionRoleMap) error
}

func ParseCustomActionRoleMap(filename string, mapping *CustomActionRoleMap) error {
	if filename == "" {
		return nil
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, mapping)
	if err != nil {
		return err
	}

	return nil
}
