package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	httpkit "github.com/go-swagger/go-swagger/httpkit"
	middleware "github.com/go-swagger/go-swagger/httpkit/middleware"
	spec "github.com/go-swagger/go-swagger/spec"
	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewSoftLayerBaremetalProvisioningAPI creates a new SoftLayerBaremetalProvisioning instance
func NewSoftLayerBaremetalProvisioningAPI(spec *spec.Document) *SoftLayerBaremetalProvisioningAPI {
	o := &SoftLayerBaremetalProvisioningAPI{
		spec:            spec,
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
	}

	return o
}

/*SoftLayerBaremetalProvisioningAPI BOSH SoftLayer baremetal provisioning (BMP) server APIs */
type SoftLayerBaremetalProvisioningAPI struct {
	spec            *spec.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer httpkit.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer httpkit.Producer

	// GetBmsDeploymentNameHandler sets the operation handler for the get bms deployment name operation
	GetBmsDeploymentNameHandler GetBmsDeploymentNameHandler
	// GetInfoHandler sets the operation handler for the get info operation
	GetInfoHandler GetInfoHandler
	// GetLoginUsernamePasswordHandler sets the operation handler for the get login username password operation
	GetLoginUsernamePasswordHandler GetLoginUsernamePasswordHandler
	// GetSlPackagePackageIDOptionsHandler sets the operation handler for the get sl package package ID options operation
	GetSlPackagePackageIDOptionsHandler GetSlPackagePackageIDOptionsHandler
	// GetSlPackagesHandler sets the operation handler for the get sl packages operation
	GetSlPackagesHandler GetSlPackagesHandler
	// GetStemcellsHandler sets the operation handler for the get stemcells operation
	GetStemcellsHandler GetStemcellsHandler
	// GetTasksHandler sets the operation handler for the get tasks operation
	GetTasksHandler GetTasksHandler
	// GetTasksTaskIDTxtLevelHandler sets the operation handler for the get tasks task ID txt level operation
	GetTasksTaskIDTxtLevelHandler GetTasksTaskIDTxtLevelHandler
	// PutBaremetalServerIDStateHandler sets the operation handler for the put baremetal server ID state operation
	PutBaremetalServerIDStateHandler PutBaremetalServerIDStateHandler
	// PutBaremetalsHandler sets the operation handler for the put baremetals operation
	PutBaremetalsHandler PutBaremetalsHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup
}

// SetDefaultProduces sets the default produces media type
func (o *SoftLayerBaremetalProvisioningAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SoftLayerBaremetalProvisioningAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// DefaultProduces returns the default produces media type
func (o *SoftLayerBaremetalProvisioningAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SoftLayerBaremetalProvisioningAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SoftLayerBaremetalProvisioningAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SoftLayerBaremetalProvisioningAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SoftLayerBaremetalProvisioningAPI
func (o *SoftLayerBaremetalProvisioningAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.GetBmsDeploymentNameHandler == nil {
		unregistered = append(unregistered, "GetBmsDeploymentNameHandler")
	}

	if o.GetInfoHandler == nil {
		unregistered = append(unregistered, "GetInfoHandler")
	}

	if o.GetLoginUsernamePasswordHandler == nil {
		unregistered = append(unregistered, "GetLoginUsernamePasswordHandler")
	}

	if o.GetSlPackagePackageIDOptionsHandler == nil {
		unregistered = append(unregistered, "GetSlPackagePackageIDOptionsHandler")
	}

	if o.GetSlPackagesHandler == nil {
		unregistered = append(unregistered, "GetSlPackagesHandler")
	}

	if o.GetStemcellsHandler == nil {
		unregistered = append(unregistered, "GetStemcellsHandler")
	}

	if o.GetTasksHandler == nil {
		unregistered = append(unregistered, "GetTasksHandler")
	}

	if o.GetTasksTaskIDTxtLevelHandler == nil {
		unregistered = append(unregistered, "GetTasksTaskIDTxtLevelHandler")
	}

	if o.PutBaremetalServerIDStateHandler == nil {
		unregistered = append(unregistered, "PutBaremetalServerIDStateHandler")
	}

	if o.PutBaremetalsHandler == nil {
		unregistered = append(unregistered, "PutBaremetalsHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SoftLayerBaremetalProvisioningAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SoftLayerBaremetalProvisioningAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]httpkit.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *SoftLayerBaremetalProvisioningAPI) ConsumersFor(mediaTypes []string) map[string]httpkit.Consumer {

	result := make(map[string]httpkit.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *SoftLayerBaremetalProvisioningAPI) ProducersFor(mediaTypes []string) map[string]httpkit.Producer {

	result := make(map[string]httpkit.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SoftLayerBaremetalProvisioningAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

func (o *SoftLayerBaremetalProvisioningAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/bms/{deploymentName}"] = NewGetBmsDeploymentName(o.context, o.GetBmsDeploymentNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/info"] = NewGetInfo(o.context, o.GetInfoHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/login/{username}/{password}"] = NewGetLoginUsernamePassword(o.context, o.GetLoginUsernamePasswordHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sl/package/{packageId}/options"] = NewGetSlPackagePackageIDOptions(o.context, o.GetSlPackagePackageIDOptionsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sl/packages"] = NewGetSlPackages(o.context, o.GetSlPackagesHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/stemcells"] = NewGetStemcells(o.context, o.GetStemcellsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tasks"] = NewGetTasks(o.context, o.GetTasksHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tasks/{taskId}/txt/{level}"] = NewGetTasksTaskIDTxtLevel(o.context, o.GetTasksTaskIDTxtLevelHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/baremetal/{serverId}/{state}"] = NewPutBaremetalServerIDState(o.context, o.PutBaremetalServerIDStateHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/baremetals"] = NewPutBaremetals(o.context, o.PutBaremetalsHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SoftLayerBaremetalProvisioningAPI) Serve(builder middleware.Builder) http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler(builder)
}
