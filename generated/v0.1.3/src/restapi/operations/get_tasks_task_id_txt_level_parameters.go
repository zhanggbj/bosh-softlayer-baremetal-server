package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetTasksTaskIDTxtLevelParams creates a new GetTasksTaskIDTxtLevelParams object
// with the default values initialized.
func NewGetTasksTaskIDTxtLevelParams() GetTasksTaskIDTxtLevelParams {
	var ()
	return GetTasksTaskIDTxtLevelParams{}
}

// GetTasksTaskIDTxtLevelParams contains all the bound params for the get tasks task ID txt level operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetTasksTaskIDTxtLevel
type GetTasksTaskIDTxtLevelParams struct {
	/*the log level to retrieve
	  Required: true
	  In: path
	*/
	Level int32
	/*the ID for the task
	  Required: true
	  In: path
	*/
	TaskID int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetTasksTaskIDTxtLevelParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	rLevel, rhkLevel, _ := route.Params.GetOK("level")
	if err := o.bindLevel(rLevel, rhkLevel, route.Formats); err != nil {
		res = append(res, err)
	}

	rTaskID, rhkTaskID, _ := route.Params.GetOK("taskId")
	if err := o.bindTaskID(rTaskID, rhkTaskID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTasksTaskIDTxtLevelParams) bindLevel(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("level", "path", "int32", raw)
	}
	o.Level = value

	return nil
}

func (o *GetTasksTaskIDTxtLevelParams) bindTaskID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("taskId", "path", "int32", raw)
	}
	o.TaskID = value

	return nil
}
