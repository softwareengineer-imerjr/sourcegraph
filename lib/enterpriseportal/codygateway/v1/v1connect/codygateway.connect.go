// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: codygateway.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/sourcegraph/sourcegraph/lib/enterpriseportal/codygateway/v1"
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
	// CodyGatewayServiceName is the fully-qualified name of the CodyGatewayService service.
	CodyGatewayServiceName = "enterpriseportal.codygateway.v1.CodyGatewayService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CodyGatewayServiceGetCodyGatewayAccessProcedure is the fully-qualified name of the
	// CodyGatewayService's GetCodyGatewayAccess RPC.
	CodyGatewayServiceGetCodyGatewayAccessProcedure = "/enterpriseportal.codygateway.v1.CodyGatewayService/GetCodyGatewayAccess"
	// CodyGatewayServiceListCodyGatewayAccessesProcedure is the fully-qualified name of the
	// CodyGatewayService's ListCodyGatewayAccesses RPC.
	CodyGatewayServiceListCodyGatewayAccessesProcedure = "/enterpriseportal.codygateway.v1.CodyGatewayService/ListCodyGatewayAccesses"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	codyGatewayServiceServiceDescriptor                       = v1.File_codygateway_proto.Services().ByName("CodyGatewayService")
	codyGatewayServiceGetCodyGatewayAccessMethodDescriptor    = codyGatewayServiceServiceDescriptor.Methods().ByName("GetCodyGatewayAccess")
	codyGatewayServiceListCodyGatewayAccessesMethodDescriptor = codyGatewayServiceServiceDescriptor.Methods().ByName("ListCodyGatewayAccesses")
)

// CodyGatewayServiceClient is a client for the enterpriseportal.codygateway.v1.CodyGatewayService
// service.
type CodyGatewayServiceClient interface {
	// Retrieve Cody Gateway access granted to an Enterprise subscription.
	GetCodyGatewayAccess(context.Context, *connect.Request[v1.GetCodyGatewayAccessRequest]) (*connect.Response[v1.GetCodyGatewayAccessResponse], error)
	// List all Cody Gateway accesses granted to any Enterprise subscription.
	ListCodyGatewayAccesses(context.Context, *connect.Request[v1.ListCodyGatewayAccessesRequest]) (*connect.Response[v1.ListCodyGatewayAccessesResponse], error)
}

// NewCodyGatewayServiceClient constructs a client for the
// enterpriseportal.codygateway.v1.CodyGatewayService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCodyGatewayServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) CodyGatewayServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &codyGatewayServiceClient{
		getCodyGatewayAccess: connect.NewClient[v1.GetCodyGatewayAccessRequest, v1.GetCodyGatewayAccessResponse](
			httpClient,
			baseURL+CodyGatewayServiceGetCodyGatewayAccessProcedure,
			connect.WithSchema(codyGatewayServiceGetCodyGatewayAccessMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		listCodyGatewayAccesses: connect.NewClient[v1.ListCodyGatewayAccessesRequest, v1.ListCodyGatewayAccessesResponse](
			httpClient,
			baseURL+CodyGatewayServiceListCodyGatewayAccessesProcedure,
			connect.WithSchema(codyGatewayServiceListCodyGatewayAccessesMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// codyGatewayServiceClient implements CodyGatewayServiceClient.
type codyGatewayServiceClient struct {
	getCodyGatewayAccess    *connect.Client[v1.GetCodyGatewayAccessRequest, v1.GetCodyGatewayAccessResponse]
	listCodyGatewayAccesses *connect.Client[v1.ListCodyGatewayAccessesRequest, v1.ListCodyGatewayAccessesResponse]
}

// GetCodyGatewayAccess calls
// enterpriseportal.codygateway.v1.CodyGatewayService.GetCodyGatewayAccess.
func (c *codyGatewayServiceClient) GetCodyGatewayAccess(ctx context.Context, req *connect.Request[v1.GetCodyGatewayAccessRequest]) (*connect.Response[v1.GetCodyGatewayAccessResponse], error) {
	return c.getCodyGatewayAccess.CallUnary(ctx, req)
}

// ListCodyGatewayAccesses calls
// enterpriseportal.codygateway.v1.CodyGatewayService.ListCodyGatewayAccesses.
func (c *codyGatewayServiceClient) ListCodyGatewayAccesses(ctx context.Context, req *connect.Request[v1.ListCodyGatewayAccessesRequest]) (*connect.Response[v1.ListCodyGatewayAccessesResponse], error) {
	return c.listCodyGatewayAccesses.CallUnary(ctx, req)
}

// CodyGatewayServiceHandler is an implementation of the
// enterpriseportal.codygateway.v1.CodyGatewayService service.
type CodyGatewayServiceHandler interface {
	// Retrieve Cody Gateway access granted to an Enterprise subscription.
	GetCodyGatewayAccess(context.Context, *connect.Request[v1.GetCodyGatewayAccessRequest]) (*connect.Response[v1.GetCodyGatewayAccessResponse], error)
	// List all Cody Gateway accesses granted to any Enterprise subscription.
	ListCodyGatewayAccesses(context.Context, *connect.Request[v1.ListCodyGatewayAccessesRequest]) (*connect.Response[v1.ListCodyGatewayAccessesResponse], error)
}

// NewCodyGatewayServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCodyGatewayServiceHandler(svc CodyGatewayServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	codyGatewayServiceGetCodyGatewayAccessHandler := connect.NewUnaryHandler(
		CodyGatewayServiceGetCodyGatewayAccessProcedure,
		svc.GetCodyGatewayAccess,
		connect.WithSchema(codyGatewayServiceGetCodyGatewayAccessMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	codyGatewayServiceListCodyGatewayAccessesHandler := connect.NewUnaryHandler(
		CodyGatewayServiceListCodyGatewayAccessesProcedure,
		svc.ListCodyGatewayAccesses,
		connect.WithSchema(codyGatewayServiceListCodyGatewayAccessesMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/enterpriseportal.codygateway.v1.CodyGatewayService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CodyGatewayServiceGetCodyGatewayAccessProcedure:
			codyGatewayServiceGetCodyGatewayAccessHandler.ServeHTTP(w, r)
		case CodyGatewayServiceListCodyGatewayAccessesProcedure:
			codyGatewayServiceListCodyGatewayAccessesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCodyGatewayServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCodyGatewayServiceHandler struct{}

func (UnimplementedCodyGatewayServiceHandler) GetCodyGatewayAccess(context.Context, *connect.Request[v1.GetCodyGatewayAccessRequest]) (*connect.Response[v1.GetCodyGatewayAccessResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("enterpriseportal.codygateway.v1.CodyGatewayService.GetCodyGatewayAccess is not implemented"))
}

func (UnimplementedCodyGatewayServiceHandler) ListCodyGatewayAccesses(context.Context, *connect.Request[v1.ListCodyGatewayAccessesRequest]) (*connect.Response[v1.ListCodyGatewayAccessesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("enterpriseportal.codygateway.v1.CodyGatewayService.ListCodyGatewayAccesses is not implemented"))
}
