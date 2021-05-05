package deploymentManager

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

type VendorSolutions []*VendorSolution

type Retry struct {
	RetryCount    int    `json:"retry_count"`
	JobAssignedAt string `json:"job_assigned_at"`
}

type ReleaseDownloadRequestLock struct {
	NodeId        string `json:"node_id"`
	LockAquiredAt string `json:"lock_aquired_at"`
}

func (v VendorSolutions) IndexOf(solution *VendorSolution) int {
	for index, s := range v {
		if s.SolutionId == solution.SolutionId {
			return index
		}
	}
	return -1
}

func (s VendorSolution) GetVersionsToDownload(downloaded *VendorSolution) []*SolutionVersion {
	var versionsToBeDownloaded []*SolutionVersion

	for _, v := range s.Versions {
		found := false
		for _, u := range downloaded.Versions {
			if v.Version == u.Version && u.State == AppDownloadState_Success {
				found = true
				break
			}
		}
		if !found {
			versionsToBeDownloaded = append(versionsToBeDownloaded, &SolutionVersion{State: AppDownloadState_AppDownloadStateUnknown, Version: v.Version})
		}
	}
	return versionsToBeDownloaded
}

func (s VendorSolution) GetVersion(version string) *SolutionVersion {
	for _, v := range s.Versions {
		if v.Version == version {
			return v
		}
	}
	return nil
}

func (s VendorSolution) IndexOf(version string) int {
	for i, v := range s.Versions {
		if v.Version == version {
			return i
		}
	}
	return -1
}

func (old *VendorSolution) Intersection(new *VendorSolution) *VendorSolution {
	var soln = &VendorSolution{
		SolutionId:   new.SolutionId,
		SolutionName: new.SolutionName,
		SolutionLogo: new.SolutionLogo,
	}
	var versions []*SolutionVersion
	for _, v := range new.Versions {
		oldVersion := old.GetVersion(v.Version)
		if oldVersion == nil {
			versions = append(versions, v)
		} else {
			versions = append(versions, &SolutionVersion{
				Version: v.Version,
				State:   oldVersion.State,
			})
		}
	}
	soln.Versions = versions
	return soln
}

func ExtractHeader(ctx context.Context) (string, string) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logrus.Info("could not get metadata")
	}
	logrus.Info("Headers are: ", headers)
	traceIDs := headers["x-traceid"]
	eventTypes := headers["x-eventtype"]
	traceID := ""
	eventType := ""
	for _, v := range traceIDs {
		traceID = v
	}

	for _, v := range eventTypes {
		eventType = v
	}

	return traceID, eventType

}
