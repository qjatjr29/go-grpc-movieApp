// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package movie

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MovieClient is the client API for Movie service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieClient interface {
	GetMovie(ctx context.Context, in *GetMovieRequest, opts ...grpc.CallOption) (*GetMovieResponse, error)
	ListPopularMovies(ctx context.Context, in *ListPopularMovieRequest, opts ...grpc.CallOption) (*ListPopularMovieResponse, error)
	ListSearchMovies(ctx context.Context, in *ListSearchMovieRequest, opts ...grpc.CallOption) (*ListSearchMovieResponse, error)
}

type movieClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieClient(cc grpc.ClientConnInterface) MovieClient {
	return &movieClient{cc}
}

func (c *movieClient) GetMovie(ctx context.Context, in *GetMovieRequest, opts ...grpc.CallOption) (*GetMovieResponse, error) {
	out := new(GetMovieResponse)
	err := c.cc.Invoke(ctx, "/movie.Movie/GetMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieClient) ListPopularMovies(ctx context.Context, in *ListPopularMovieRequest, opts ...grpc.CallOption) (*ListPopularMovieResponse, error) {
	out := new(ListPopularMovieResponse)
	err := c.cc.Invoke(ctx, "/movie.Movie/ListPopularMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieClient) ListSearchMovies(ctx context.Context, in *ListSearchMovieRequest, opts ...grpc.CallOption) (*ListSearchMovieResponse, error) {
	out := new(ListSearchMovieResponse)
	err := c.cc.Invoke(ctx, "/movie.Movie/ListSearchMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServer is the server API for Movie service.
// All implementations must embed UnimplementedMovieServer
// for forward compatibility
type MovieServer interface {
	GetMovie(context.Context, *GetMovieRequest) (*GetMovieResponse, error)
	ListPopularMovies(context.Context, *ListPopularMovieRequest) (*ListPopularMovieResponse, error)
	ListSearchMovies(context.Context, *ListSearchMovieRequest) (*ListSearchMovieResponse, error)
	mustEmbedUnimplementedMovieServer()
}

// UnimplementedMovieServer must be embedded to have forward compatible implementations.
type UnimplementedMovieServer struct {
}

func (UnimplementedMovieServer) GetMovie(context.Context, *GetMovieRequest) (*GetMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}
func (UnimplementedMovieServer) ListPopularMovies(context.Context, *ListPopularMovieRequest) (*ListPopularMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPopularMovies not implemented")
}
func (UnimplementedMovieServer) ListSearchMovies(context.Context, *ListSearchMovieRequest) (*ListSearchMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSearchMovies not implemented")
}
func (UnimplementedMovieServer) mustEmbedUnimplementedMovieServer() {}

// UnsafeMovieServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieServer will
// result in compilation errors.
type UnsafeMovieServer interface {
	mustEmbedUnimplementedMovieServer()
}

func RegisterMovieServer(s grpc.ServiceRegistrar, srv MovieServer) {
	s.RegisterService(&Movie_ServiceDesc, srv)
}

func _Movie_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movie.Movie/GetMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServer).GetMovie(ctx, req.(*GetMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Movie_ListPopularMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPopularMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServer).ListPopularMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movie.Movie/ListPopularMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServer).ListPopularMovies(ctx, req.(*ListPopularMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Movie_ListSearchMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSearchMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServer).ListSearchMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movie.Movie/ListSearchMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServer).ListSearchMovies(ctx, req.(*ListSearchMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Movie_ServiceDesc is the grpc.ServiceDesc for Movie service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Movie_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "movie.Movie",
	HandlerType: (*MovieServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovie",
			Handler:    _Movie_GetMovie_Handler,
		},
		{
			MethodName: "ListPopularMovies",
			Handler:    _Movie_ListPopularMovies_Handler,
		},
		{
			MethodName: "ListSearchMovies",
			Handler:    _Movie_ListSearchMovies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/movie/movie.proto",
}
