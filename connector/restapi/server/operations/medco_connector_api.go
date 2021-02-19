// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ldsec/medco/connector/restapi/models"
	"github.com/ldsec/medco/connector/restapi/server/operations/explore_statistics"
	"github.com/ldsec/medco/connector/restapi/server/operations/genomic_annotations"
	"github.com/ldsec/medco/connector/restapi/server/operations/medco_network"
	"github.com/ldsec/medco/connector/restapi/server/operations/medco_node"
	"github.com/ldsec/medco/connector/restapi/server/operations/survival_analysis"
)

// NewMedcoConnectorAPI creates a new MedcoConnector instance
func NewMedcoConnectorAPI(spec *loads.Document) *MedcoConnectorAPI {
	return &MedcoConnectorAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		MedcoNodeDeleteCohortsHandler: medco_node.DeleteCohortsHandlerFunc(func(params medco_node.DeleteCohortsParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation medco_node.DeleteCohorts has not yet been implemented")
		}),
		MedcoNodeExploreQueryHandler: medco_node.ExploreQueryHandlerFunc(func(params medco_node.ExploreQueryParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation medco_node.ExploreQuery has not yet been implemented")
		}),
		MedcoNodeExploreSearchConceptHandler: medco_node.ExploreSearchConceptHandlerFunc(func(params medco_node.ExploreSearchConceptParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation medco_node.ExploreSearchConcept has not yet been implemented")
		}),
		MedcoNodeExploreSearchModifierHandler: medco_node.ExploreSearchModifierHandlerFunc(func(params medco_node.ExploreSearchModifierParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation medco_node.ExploreSearchModifier has not yet been implemented")
		}),
		ExploreStatisticsExploreStatisticsHandler: explore_statistics.ExploreStatisticsHandlerFunc(func(params explore_statistics.ExploreStatisticsParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation explore_statistics.ExploreStatistics has not yet been implemented")
		}),
		MedcoNodeGetCohortsHandler: medco_node.GetCohortsHandlerFunc(func(params medco_node.GetCohortsParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation medco_node.GetCohorts has not yet been implemented")
		}),
		MedcoNodeGetExploreQueryHandler: medco_node.GetExploreQueryHandlerFunc(func(params medco_node.GetExploreQueryParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation medco_node.GetExploreQuery has not yet been implemented")
		}),
		MedcoNetworkGetMetadataHandler: medco_network.GetMetadataHandlerFunc(func(params medco_network.GetMetadataParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation medco_network.GetMetadata has not yet been implemented")
		}),
		GenomicAnnotationsGetValuesHandler: genomic_annotations.GetValuesHandlerFunc(func(params genomic_annotations.GetValuesParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation genomic_annotations.GetValues has not yet been implemented")
		}),
		GenomicAnnotationsGetVariantsHandler: genomic_annotations.GetVariantsHandlerFunc(func(params genomic_annotations.GetVariantsParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation genomic_annotations.GetVariants has not yet been implemented")
		}),
		MedcoNodePostCohortsHandler: medco_node.PostCohortsHandlerFunc(func(params medco_node.PostCohortsParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation medco_node.PostCohorts has not yet been implemented")
		}),
		MedcoNodePutCohortsHandler: medco_node.PutCohortsHandlerFunc(func(params medco_node.PutCohortsParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation medco_node.PutCohorts has not yet been implemented")
		}),
		SurvivalAnalysisSurvivalAnalysisHandler: survival_analysis.SurvivalAnalysisHandlerFunc(func(params survival_analysis.SurvivalAnalysisParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation survival_analysis.SurvivalAnalysis has not yet been implemented")
		}),

		MedcoJwtAuth: func(token string, scopes []string) (*models.User, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (medco-jwt) has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*MedcoConnectorAPI API of the MedCo connector, that orchestrates the query at the MedCo node and provides information about the MedCo network. */
type MedcoConnectorAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// MedcoJwtAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	MedcoJwtAuth func(string, []string) (*models.User, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// MedcoNodeDeleteCohortsHandler sets the operation handler for the delete cohorts operation
	MedcoNodeDeleteCohortsHandler medco_node.DeleteCohortsHandler
	// MedcoNodeExploreQueryHandler sets the operation handler for the explore query operation
	MedcoNodeExploreQueryHandler medco_node.ExploreQueryHandler
	// MedcoNodeExploreSearchConceptHandler sets the operation handler for the explore search concept operation
	MedcoNodeExploreSearchConceptHandler medco_node.ExploreSearchConceptHandler
	// MedcoNodeExploreSearchModifierHandler sets the operation handler for the explore search modifier operation
	MedcoNodeExploreSearchModifierHandler medco_node.ExploreSearchModifierHandler
	// ExploreStatisticsExploreStatisticsHandler sets the operation handler for the explore statistics operation
	ExploreStatisticsExploreStatisticsHandler explore_statistics.ExploreStatisticsHandler
	// MedcoNodeGetCohortsHandler sets the operation handler for the get cohorts operation
	MedcoNodeGetCohortsHandler medco_node.GetCohortsHandler
	// MedcoNodeGetExploreQueryHandler sets the operation handler for the get explore query operation
	MedcoNodeGetExploreQueryHandler medco_node.GetExploreQueryHandler
	// MedcoNetworkGetMetadataHandler sets the operation handler for the get metadata operation
	MedcoNetworkGetMetadataHandler medco_network.GetMetadataHandler
	// GenomicAnnotationsGetValuesHandler sets the operation handler for the get values operation
	GenomicAnnotationsGetValuesHandler genomic_annotations.GetValuesHandler
	// GenomicAnnotationsGetVariantsHandler sets the operation handler for the get variants operation
	GenomicAnnotationsGetVariantsHandler genomic_annotations.GetVariantsHandler
	// MedcoNodePostCohortsHandler sets the operation handler for the post cohorts operation
	MedcoNodePostCohortsHandler medco_node.PostCohortsHandler
	// MedcoNodePutCohortsHandler sets the operation handler for the put cohorts operation
	MedcoNodePutCohortsHandler medco_node.PutCohortsHandler
	// SurvivalAnalysisSurvivalAnalysisHandler sets the operation handler for the survival analysis operation
	SurvivalAnalysisSurvivalAnalysisHandler survival_analysis.SurvivalAnalysisHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *MedcoConnectorAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *MedcoConnectorAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *MedcoConnectorAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *MedcoConnectorAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *MedcoConnectorAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *MedcoConnectorAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *MedcoConnectorAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *MedcoConnectorAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *MedcoConnectorAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the MedcoConnectorAPI
func (o *MedcoConnectorAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.MedcoJwtAuth == nil {
		unregistered = append(unregistered, "MedcoJwtAuth")
	}

	if o.MedcoNodeDeleteCohortsHandler == nil {
		unregistered = append(unregistered, "medco_node.DeleteCohortsHandler")
	}
	if o.MedcoNodeExploreQueryHandler == nil {
		unregistered = append(unregistered, "medco_node.ExploreQueryHandler")
	}
	if o.MedcoNodeExploreSearchConceptHandler == nil {
		unregistered = append(unregistered, "medco_node.ExploreSearchConceptHandler")
	}
	if o.MedcoNodeExploreSearchModifierHandler == nil {
		unregistered = append(unregistered, "medco_node.ExploreSearchModifierHandler")
	}
	if o.ExploreStatisticsExploreStatisticsHandler == nil {
		unregistered = append(unregistered, "explore_statistics.ExploreStatisticsHandler")
	}
	if o.MedcoNodeGetCohortsHandler == nil {
		unregistered = append(unregistered, "medco_node.GetCohortsHandler")
	}
	if o.MedcoNodeGetExploreQueryHandler == nil {
		unregistered = append(unregistered, "medco_node.GetExploreQueryHandler")
	}
	if o.MedcoNetworkGetMetadataHandler == nil {
		unregistered = append(unregistered, "medco_network.GetMetadataHandler")
	}
	if o.GenomicAnnotationsGetValuesHandler == nil {
		unregistered = append(unregistered, "genomic_annotations.GetValuesHandler")
	}
	if o.GenomicAnnotationsGetVariantsHandler == nil {
		unregistered = append(unregistered, "genomic_annotations.GetVariantsHandler")
	}
	if o.MedcoNodePostCohortsHandler == nil {
		unregistered = append(unregistered, "medco_node.PostCohortsHandler")
	}
	if o.MedcoNodePutCohortsHandler == nil {
		unregistered = append(unregistered, "medco_node.PutCohortsHandler")
	}
	if o.SurvivalAnalysisSurvivalAnalysisHandler == nil {
		unregistered = append(unregistered, "survival_analysis.SurvivalAnalysisHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *MedcoConnectorAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *MedcoConnectorAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "medco-jwt":
			result[name] = o.BearerAuthenticator(name, func(token string, scopes []string) (interface{}, error) {
				return o.MedcoJwtAuth(token, scopes)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *MedcoConnectorAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *MedcoConnectorAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *MedcoConnectorAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *MedcoConnectorAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the medco connector API
func (o *MedcoConnectorAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *MedcoConnectorAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/node/explore/cohorts/{name}"] = medco_node.NewDeleteCohorts(o.context, o.MedcoNodeDeleteCohortsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/node/explore/query"] = medco_node.NewExploreQuery(o.context, o.MedcoNodeExploreQueryHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/node/explore/search/concept"] = medco_node.NewExploreSearchConcept(o.context, o.MedcoNodeExploreSearchConceptHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/node/explore/search/modifier"] = medco_node.NewExploreSearchModifier(o.context, o.MedcoNodeExploreSearchModifierHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/node/explore-statistics/query"] = explore_statistics.NewExploreStatistics(o.context, o.ExploreStatisticsExploreStatisticsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/node/explore/cohorts"] = medco_node.NewGetCohorts(o.context, o.MedcoNodeGetCohortsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/node/explore/query/{queryId}"] = medco_node.NewGetExploreQuery(o.context, o.MedcoNodeGetExploreQueryHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/network"] = medco_network.NewGetMetadata(o.context, o.MedcoNetworkGetMetadataHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/genomic-annotations/{annotation}"] = genomic_annotations.NewGetValues(o.context, o.GenomicAnnotationsGetValuesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/genomic-annotations/{annotation}/{value}"] = genomic_annotations.NewGetVariants(o.context, o.GenomicAnnotationsGetVariantsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/node/explore/cohorts/{name}"] = medco_node.NewPostCohorts(o.context, o.MedcoNodePostCohortsHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/node/explore/cohorts/{name}"] = medco_node.NewPutCohorts(o.context, o.MedcoNodePutCohortsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/node/analysis/survival/query"] = survival_analysis.NewSurvivalAnalysis(o.context, o.SurvivalAnalysisSurvivalAnalysisHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *MedcoConnectorAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *MedcoConnectorAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *MedcoConnectorAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *MedcoConnectorAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *MedcoConnectorAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
