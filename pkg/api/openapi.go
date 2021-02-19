// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Annotations defines model for Annotations.
type Annotations map[string]interface{}

// Cluster defines model for Cluster.
type Cluster struct {
	// Embedded struct due to allOf(#/components/schemas/ClusterId)
	ClusterId
	// Embedded struct due to allOf(#/components/schemas/ClusterTenant)
	ClusterTenant
	// Embedded struct due to allOf(#/components/schemas/ClusterProperties)
	ClusterProperties
}

// ClusterFacts defines model for ClusterFacts.
type ClusterFacts map[string]interface{}

// ClusterId defines model for ClusterId.
type ClusterId struct {

	// A unique object identifier string. Automatically generated by the API on creation (in the form "<letter>-<adjective>-<noun>-<digits>" where all letters are lowercase, max 63 characters in total).
	Id Id `json:"id"`
}

// ClusterProperties defines model for ClusterProperties.
type ClusterProperties struct {

	// Unstructured key value map containing arbitrary metadata
	Annotations *Annotations `json:"annotations,omitempty"`

	// Display Name of the cluster
	DisplayName *string `json:"displayName,omitempty"`

	// Facts about a cluster object. Statically configured key/value pairs.
	Facts *ClusterFacts `json:"facts,omitempty"`

	// Configuration Git repository, usually generated by the API
	GitRepo *GitRepo `json:"gitRepo,omitempty"`

	// Git revision to use with the global configruation git repository.
	// This takes precedence over the revision configured on the Tenant.
	GlobalGitRepoRevision *string `json:"globalGitRepoRevision,omitempty"`

	// URL to fetch install manifests for Steward cluster agent. This will only be set if the cluster's token is still valid.
	InstallURL *string `json:"installURL,omitempty"`

	// Git revision to use with the tenant configruation git repository.
	// This takes precedence over the revision configured on the Tenant.
	TenantGitRepoRevision *string `json:"tenantGitRepoRevision,omitempty"`
}

// ClusterTenant defines model for ClusterTenant.
type ClusterTenant struct {

	// Id of the tenant this cluster belongs to
	Tenant string `json:"tenant"`
}

// GitRepo defines model for GitRepo.
type GitRepo struct {

	// SSH public key / deploy key for clusterconfiguration catalog Git repository. This property is managed by Steward.
	DeployKey *string `json:"deployKey,omitempty"`

	// SSH known hosts of the git server (multiline possible for multiple keys)
	HostKeys *string `json:"hostKeys,omitempty"`

	// Specifies if a repo should be managed by the git controller. A value of 'unmanaged' means it's not manged by the controller
	Type *string `json:"type,omitempty"`

	// Full URL of the git repo
	Url *string `json:"url,omitempty"`
}

// Id defines model for Id.
type Id string

// Inventory defines model for Inventory.
type Inventory struct {
	Cluster   string                  `json:"cluster"`
	Inventory *map[string]interface{} `json:"inventory,omitempty"`
}

// Reason defines model for Reason.
type Reason struct {

	// The reason message
	Reason string `json:"reason"`
}

// Revision defines model for Revision.
type Revision struct {

	// Revision to use with a git repository.
	Revision *string `json:"revision,omitempty"`
}

// RevisionedGitRepo defines model for RevisionedGitRepo.
type RevisionedGitRepo struct {
	// Embedded struct due to allOf(#/components/schemas/GitRepo)
	GitRepo
	// Embedded struct due to allOf(#/components/schemas/Revision)
	Revision
}

// Tenant defines model for Tenant.
type Tenant struct {
	// Embedded struct due to allOf(#/components/schemas/TenantId)
	TenantId
	// Embedded struct due to allOf(#/components/schemas/TenantProperties)
	TenantProperties
}

// TenantId defines model for TenantId.
type TenantId struct {

	// A unique object identifier string. Automatically generated by the API on creation (in the form "<letter>-<adjective>-<noun>-<digits>" where all letters are lowercase, max 63 characters in total).
	Id Id `json:"id"`
}

// TenantProperties defines model for TenantProperties.
type TenantProperties struct {

	// Unstructured key value map containing arbitrary metadata
	Annotations *Annotations `json:"annotations,omitempty"`

	// Display name of the tenant
	DisplayName *string            `json:"displayName,omitempty"`
	GitRepo     *RevisionedGitRepo `json:"gitRepo,omitempty"`

	// Git revision to use with the global configruation git repository.
	GlobalGitRepoRevision *string `json:"globalGitRepoRevision,omitempty"`

	// Full URL of the global configuration git repo
	GlobalGitRepoURL *string `json:"globalGitRepoURL,omitempty"`
}

// ClusterIdParameter defines model for ClusterIdParameter.
type ClusterIdParameter Id

// TenantIdParameter defines model for TenantIdParameter.
type TenantIdParameter Id

// Default defines model for Default.
type Default Reason

// ListClustersParams defines parameters for ListClusters.
type ListClustersParams struct {

	// Filter clusters by tenant id
	Tenant *string `json:"tenant,omitempty"`
}

// CreateClusterJSONBody defines parameters for CreateCluster.
type CreateClusterJSONBody Cluster

// InstallStewardParams defines parameters for InstallSteward.
type InstallStewardParams struct {

	// Initial bootstrap token
	Token *string `json:"token,omitempty"`
}

// QueryInventoryParams defines parameters for QueryInventory.
type QueryInventoryParams struct {

	// InfluxQL query string
	Q *string `json:"q,omitempty"`
}

// UpdateInventoryJSONBody defines parameters for UpdateInventory.
type UpdateInventoryJSONBody Inventory

// CreateTenantJSONBody defines parameters for CreateTenant.
type CreateTenantJSONBody Tenant

// CreateClusterRequestBody defines body for CreateCluster for application/json ContentType.
type CreateClusterJSONRequestBody CreateClusterJSONBody

// UpdateInventoryRequestBody defines body for UpdateInventory for application/json ContentType.
type UpdateInventoryJSONRequestBody UpdateInventoryJSONBody

// CreateTenantRequestBody defines body for CreateTenant for application/json ContentType.
type CreateTenantJSONRequestBody CreateTenantJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns a list of clusters
	// (GET /clusters)
	ListClusters(ctx echo.Context, params ListClustersParams) error
	// Creates a new cluster
	// (POST /clusters)
	CreateCluster(ctx echo.Context) error
	// Deletes a cluster
	// (DELETE /clusters/{clusterId})
	DeleteCluster(ctx echo.Context, clusterId ClusterIdParameter) error
	// Returns all values of a cluster
	// (GET /clusters/{clusterId})
	GetCluster(ctx echo.Context, clusterId ClusterIdParameter) error
	// Updates a cluster
	// (PATCH /clusters/{clusterId})
	UpdateCluster(ctx echo.Context, clusterId ClusterIdParameter) error
	// API documentation
	// (GET /docs)
	Docs(ctx echo.Context) error
	// API health check
	// (GET /healthz)
	Healthz(ctx echo.Context) error
	// Returns the Steward JSON installation manifest
	// (GET /install/steward.json)
	InstallSteward(ctx echo.Context, params InstallStewardParams) error
	// Returns inventory data according to query
	// (GET /inventory)
	QueryInventory(ctx echo.Context, params QueryInventoryParams) error
	// Write inventory data
	// (POST /inventory)
	UpdateInventory(ctx echo.Context) error
	// OpenAPI JSON spec
	// (GET /openapi.json)
	Openapi(ctx echo.Context) error
	// Returns a list of tenants
	// (GET /tenants)
	ListTenants(ctx echo.Context) error
	// Creates a new tenant
	// (POST /tenants)
	CreateTenant(ctx echo.Context) error
	// Deletes a tenant
	// (DELETE /tenants/{tenantId})
	DeleteTenant(ctx echo.Context, tenantId TenantIdParameter) error
	// Returns all values of a tenant
	// (GET /tenants/{tenantId})
	GetTenant(ctx echo.Context, tenantId TenantIdParameter) error
	// Updates a tenant
	// (PATCH /tenants/{tenantId})
	UpdateTenant(ctx echo.Context, tenantId TenantIdParameter) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListClusters converts echo context to params.
func (w *ServerInterfaceWrapper) ListClusters(ctx echo.Context) error {
	var err error

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListClustersParams
	// ------------- Optional query parameter "tenant" -------------

	err = runtime.BindQueryParameter("form", true, false, "tenant", ctx.QueryParams(), &params.Tenant)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenant: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListClusters(ctx, params)
	return err
}

// CreateCluster converts echo context to params.
func (w *ServerInterfaceWrapper) CreateCluster(ctx echo.Context) error {
	var err error

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateCluster(ctx)
	return err
}

// DeleteCluster converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteCluster(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clusterId" -------------
	var clusterId ClusterIdParameter

	err = runtime.BindStyledParameter("simple", false, "clusterId", ctx.Param("clusterId"), &clusterId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clusterId: %s", err))
	}

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteCluster(ctx, clusterId)
	return err
}

// GetCluster converts echo context to params.
func (w *ServerInterfaceWrapper) GetCluster(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clusterId" -------------
	var clusterId ClusterIdParameter

	err = runtime.BindStyledParameter("simple", false, "clusterId", ctx.Param("clusterId"), &clusterId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clusterId: %s", err))
	}

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCluster(ctx, clusterId)
	return err
}

// UpdateCluster converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateCluster(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clusterId" -------------
	var clusterId ClusterIdParameter

	err = runtime.BindStyledParameter("simple", false, "clusterId", ctx.Param("clusterId"), &clusterId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clusterId: %s", err))
	}

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateCluster(ctx, clusterId)
	return err
}

// Docs converts echo context to params.
func (w *ServerInterfaceWrapper) Docs(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Docs(ctx)
	return err
}

// Healthz converts echo context to params.
func (w *ServerInterfaceWrapper) Healthz(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Healthz(ctx)
	return err
}

// InstallSteward converts echo context to params.
func (w *ServerInterfaceWrapper) InstallSteward(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params InstallStewardParams
	// ------------- Optional query parameter "token" -------------

	err = runtime.BindQueryParameter("form", true, false, "token", ctx.QueryParams(), &params.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter token: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.InstallSteward(ctx, params)
	return err
}

// QueryInventory converts echo context to params.
func (w *ServerInterfaceWrapper) QueryInventory(ctx echo.Context) error {
	var err error

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params QueryInventoryParams
	// ------------- Optional query parameter "q" -------------

	err = runtime.BindQueryParameter("form", true, false, "q", ctx.QueryParams(), &params.Q)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter q: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.QueryInventory(ctx, params)
	return err
}

// UpdateInventory converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateInventory(ctx echo.Context) error {
	var err error

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateInventory(ctx)
	return err
}

// Openapi converts echo context to params.
func (w *ServerInterfaceWrapper) Openapi(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Openapi(ctx)
	return err
}

// ListTenants converts echo context to params.
func (w *ServerInterfaceWrapper) ListTenants(ctx echo.Context) error {
	var err error

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListTenants(ctx)
	return err
}

// CreateTenant converts echo context to params.
func (w *ServerInterfaceWrapper) CreateTenant(ctx echo.Context) error {
	var err error

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateTenant(ctx)
	return err
}

// DeleteTenant converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTenant(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenantId" -------------
	var tenantId TenantIdParameter

	err = runtime.BindStyledParameter("simple", false, "tenantId", ctx.Param("tenantId"), &tenantId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenantId: %s", err))
	}

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteTenant(ctx, tenantId)
	return err
}

// GetTenant converts echo context to params.
func (w *ServerInterfaceWrapper) GetTenant(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenantId" -------------
	var tenantId TenantIdParameter

	err = runtime.BindStyledParameter("simple", false, "tenantId", ctx.Param("tenantId"), &tenantId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenantId: %s", err))
	}

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTenant(ctx, tenantId)
	return err
}

// UpdateTenant converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateTenant(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenantId" -------------
	var tenantId TenantIdParameter

	err = runtime.BindStyledParameter("simple", false, "tenantId", ctx.Param("tenantId"), &tenantId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenantId: %s", err))
	}

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateTenant(ctx, tenantId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/clusters", wrapper.ListClusters)
	router.POST(baseURL+"/clusters", wrapper.CreateCluster)
	router.DELETE(baseURL+"/clusters/:clusterId", wrapper.DeleteCluster)
	router.GET(baseURL+"/clusters/:clusterId", wrapper.GetCluster)
	router.PATCH(baseURL+"/clusters/:clusterId", wrapper.UpdateCluster)
	router.GET(baseURL+"/docs", wrapper.Docs)
	router.GET(baseURL+"/healthz", wrapper.Healthz)
	router.GET(baseURL+"/install/steward.json", wrapper.InstallSteward)
	router.GET(baseURL+"/inventory", wrapper.QueryInventory)
	router.POST(baseURL+"/inventory", wrapper.UpdateInventory)
	router.GET(baseURL+"/openapi.json", wrapper.Openapi)
	router.GET(baseURL+"/tenants", wrapper.ListTenants)
	router.POST(baseURL+"/tenants", wrapper.CreateTenant)
	router.DELETE(baseURL+"/tenants/:tenantId", wrapper.DeleteTenant)
	router.GET(baseURL+"/tenants/:tenantId", wrapper.GetTenant)
	router.PATCH(baseURL+"/tenants/:tenantId", wrapper.UpdateTenant)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
