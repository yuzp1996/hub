// Code generated by goa v3.14.5, DO NOT EDIT.
//
// catalog views
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// Job is the viewed result type that is projected based on a view.
type Job struct {
	// Type to project
	Projected *JobView
	// View to render
	View string
}

// JobView is a type that runs validations on a projected type.
type JobView struct {
	// id of the job
	ID *uint
	// Name of the catalog
	CatalogName *string
	// status of the job
	Status *string
}

var (
	// JobMap is a map indexing the attribute names of Job by view name.
	JobMap = map[string][]string{
		"default": {
			"id",
			"catalogName",
			"status",
		},
	}
)

// ValidateJob runs the validations defined on the viewed result type Job.
func ValidateJob(result *Job) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateJobView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateJobView runs the validations defined on JobView using the "default"
// view.
func ValidateJobView(result *JobView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.CatalogName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("catalogName", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	return
}
