package projects

import (
	"github.com/Flagsmith/flagsmith-go-client/flagengine/organisations"
	"github.com/Flagsmith/flagsmith-go-client/flagengine/segments"
)

type ProjectModel struct {
	ID                int                              `json:"id"`
	Name              string                           `json:"name"`
	HideDisabledFlags bool                             `json:"hide_disabled_flags"`
	Organization      *organisations.OrganisationModel `json:"organization"`
	Segments          []*segments.SegmentModel         `json:"segments"`
}
