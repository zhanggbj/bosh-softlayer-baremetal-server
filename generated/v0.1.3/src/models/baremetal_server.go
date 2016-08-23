package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*BaremetalServer baremetal server

swagger:model BaremetalServer
*/
type BaremetalServer struct {

	/* cpu
	 */
	CPU *int32 `json:"cpu,omitempty"`

	/* a string status of the hardware
	 */
	HardwareStatus *string `json:"hardware_status,omitempty"`

	/* the hostname for the baremetal server
	 */
	Hostname *string `json:"hostname,omitempty"`

	/* id
	 */
	ID *int32 `json:"id,omitempty"`

	/* memory
	 */
	Memory *int32 `json:"memory,omitempty"`

	/* the private IP address for the baremetal server
	 */
	PrivateIPAddress *string `json:"private_ip_address,omitempty"`

	/* the date for the provisioning, e.g., 2016-07-05T19:38:08+08:00
	 */
	ProvisionDate *string `json:"provision_date,omitempty"`

	/* the public IP address for the baremetal server
	 */
	PublicIPAddress *string `json:"public_ip_address,omitempty"`

	/* tags
	 */
	Tags []string `json:"tags,omitempty"`
}

// Validate validates this baremetal server
func (m *BaremetalServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaremetalServer) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.RequiredString("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i])); err != nil {
			return err
		}

	}

	return nil
}
