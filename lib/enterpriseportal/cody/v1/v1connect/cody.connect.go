// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: cody.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/sourcegraph/sourcegraph/lib/enterpriseportal/cody/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// EnterprisePortalCodyServiceName is the fully-qualified name of the EnterprisePortalCodyService
	// service.
	EnterprisePortalCodyServiceName = "sourcegraph.enterpriseportal.cody.v1.EnterprisePortalCodyService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// EnterprisePortalCodyServiceGetCodyGatewayAccessProcedure is the fully-qualified name of the
	// EnterprisePortalCodyService's GetCodyGatewayAccess RPC.
	EnterprisePortalCodyServiceGetCodyGatewayAccessProcedure = "/sourcegraph.enterpriseportal.cody.v1.EnterprisePortalCodyService/GetCodyGatewayAccess"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	enterprisePortalCodyServiceServiceDescriptor                    = v1.File_cody_proto.Services().ByName("EnterprisePortalCodyService")
	enterprisePortalCodyServiceGetCodyGatewayAccessMethodDescriptor = enterprisePortalCodyServiceServiceDescriptor.Methods().ByName("GetCodyGatewayAccess")
)

// EnterprisePortalCodyServiceClient is a client for the
// sourcegraph.enterpriseportal.cody.v1.EnterprisePortalCodyService service.
type EnterprisePortalCodyServiceClient interface {
	// Retrieve Cody Gateway access granted to an Enterprise subscription.
	// Properties may be inferred from the active license, or be defined in
	// overrides.
	GetCodyGatewayAccess(context.Context, *connect.Request[v1.GetCodyGatewayAccessRequest]) (*connect.Response[v1.GetCodyGatewayAccessResponse], error)
}

// NewEnterprisePortalCodyServiceClient constructs a client for the
// sourcegraph.enterpriseportal.cody.v1.EnterprisePortalCodyService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEnterprisePortalCodyServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) EnterprisePortalCodyServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &enterprisePortalCodyServiceClient{
		getCodyGatewayAccess: connect.NewClient[v1.GetCodyGatewayAccessRequest, v1.GetCodyGatewayAccessResponse](
			httpClient,
			baseURL+EnterprisePortalCodyServiceGetCodyGatewayAccessProcedure,
			connect.WithSchema(enterprisePortalCodyServiceGetCodyGatewayAccessMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// enterprisePortalCodyServiceClient implements EnterprisePortalCodyServiceClient.
type enterprisePortalCodyServiceClient struct {
	getCodyGatewayAccess *connect.Client[v1.GetCodyGatewayAccessRequest, v1.GetCodyGatewayAccessResponse]
}

// GetCodyGatewayAccess calls
// sourcegraph.enterpriseportal.cody.v1.EnterprisePortalCodyService.GetCodyGatewayAccess.
func (c *enterprisePortalCodyServiceClient) GetCodyGatewayAccess(ctx context.Context, req *connect.Request[v1.GetCodyGatewayAccessRequest]) (*connect.Response[v1.GetCodyGatewayAccessResponse], error) {
	return c.getCodyGatewayAccess.CallUnary(ctx, req)
}

// EnterprisePortalCodyServiceHandler is an implementation of the
// sourcegraph.enterpriseportal.cody.v1.EnterprisePortalCodyService service.
type EnterprisePortalCodyServiceHandler interface {
	// Retrieve Cody Gateway access granted to an Enterprise subscription.
	// Properties may be inferred from the active license, or be defined in
	// overrides.
	GetCodyGatewayAccess(context.Context, *connect.Request[v1.GetCodyGatewayAccessRequest]) (*connect.Response[v1.GetCodyGatewayAccessResponse], error)
}

// NewEnterprisePortalCodyServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEnterprisePortalCodyServiceHandler(svc EnterprisePortalCodyServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	enterprisePortalCodyServiceGetCodyGatewayAccessHandler := connect.NewUnaryHandler(
		EnterprisePortalCodyServiceGetCodyGatewayAccessProcedure,
		svc.GetCodyGatewayAccess,
		connect.WithSchema(enterprisePortalCodyServiceGetCodyGatewayAccessMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/sourcegraph.enterpriseportal.cody.v1.EnterprisePortalCodyService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case EnterprisePortalCodyServiceGetCodyGatewayAccessProcedure:
			enterprisePortalCodyServiceGetCodyGatewayAccessHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedEnterprisePortalCodyServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEnterprisePortalCodyServiceHandler struct{}

func (UnimplementedEnterprisePortalCodyServiceHandler) GetCodyGatewayAccess(context.Context, *connect.Request[v1.GetCodyGatewayAccessRequest]) (*connect.Response[v1.GetCodyGatewayAccessResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("sourcegraph.enterpriseportal.cody.v1.EnterprisePortalCodyService.GetCodyGatewayAccess is not implemented"))
}
