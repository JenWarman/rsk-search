// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResultList, error)
	GetSearchMetadata(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SearchMetadata, error)
	ListFieldValues(ctx context.Context, in *ListFieldValuesRequest, opts ...grpc.CallOption) (*FieldValueList, error)
	GetEpisode(ctx context.Context, in *GetEpisodeRequest, opts ...grpc.CallOption) (*Episode, error)
	ListEpisodes(ctx context.Context, in *ListEpisodesRequest, opts ...grpc.CallOption) (*EpisodeList, error)
	// tscript is an incomplete transcription
	// chunks are ~2 min sections of the transcription
	GetTscriptChunkStats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ChunkStats, error)
	GetTscriptChunk(ctx context.Context, in *GetTscriptChunkRequest, opts ...grpc.CallOption) (*TscriptChunk, error)
	ListTscriptChunkSubmissions(ctx context.Context, in *ListTscriptChunkSubmissionsRequest, opts ...grpc.CallOption) (*ChunkSubmissionList, error)
	SubmitTscriptChunk(ctx context.Context, in *TscriptChunkSubmissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SubmitDialogCorrection(ctx context.Context, in *SubmitDialogCorrectionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResultList, error) {
	out := new(SearchResultList)
	err := c.cc.Invoke(ctx, "/rsksearch.SearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetSearchMetadata(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SearchMetadata, error) {
	out := new(SearchMetadata)
	err := c.cc.Invoke(ctx, "/rsksearch.SearchService/GetSearchMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) ListFieldValues(ctx context.Context, in *ListFieldValuesRequest, opts ...grpc.CallOption) (*FieldValueList, error) {
	out := new(FieldValueList)
	err := c.cc.Invoke(ctx, "/rsksearch.SearchService/ListFieldValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetEpisode(ctx context.Context, in *GetEpisodeRequest, opts ...grpc.CallOption) (*Episode, error) {
	out := new(Episode)
	err := c.cc.Invoke(ctx, "/rsksearch.SearchService/GetEpisode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) ListEpisodes(ctx context.Context, in *ListEpisodesRequest, opts ...grpc.CallOption) (*EpisodeList, error) {
	out := new(EpisodeList)
	err := c.cc.Invoke(ctx, "/rsksearch.SearchService/ListEpisodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetTscriptChunkStats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ChunkStats, error) {
	out := new(ChunkStats)
	err := c.cc.Invoke(ctx, "/rsksearch.SearchService/GetTscriptChunkStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) GetTscriptChunk(ctx context.Context, in *GetTscriptChunkRequest, opts ...grpc.CallOption) (*TscriptChunk, error) {
	out := new(TscriptChunk)
	err := c.cc.Invoke(ctx, "/rsksearch.SearchService/GetTscriptChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) ListTscriptChunkSubmissions(ctx context.Context, in *ListTscriptChunkSubmissionsRequest, opts ...grpc.CallOption) (*ChunkSubmissionList, error) {
	out := new(ChunkSubmissionList)
	err := c.cc.Invoke(ctx, "/rsksearch.SearchService/ListTscriptChunkSubmissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SubmitTscriptChunk(ctx context.Context, in *TscriptChunkSubmissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/rsksearch.SearchService/SubmitTscriptChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SubmitDialogCorrection(ctx context.Context, in *SubmitDialogCorrectionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/rsksearch.SearchService/SubmitDialogCorrection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations should embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	Search(context.Context, *SearchRequest) (*SearchResultList, error)
	GetSearchMetadata(context.Context, *emptypb.Empty) (*SearchMetadata, error)
	ListFieldValues(context.Context, *ListFieldValuesRequest) (*FieldValueList, error)
	GetEpisode(context.Context, *GetEpisodeRequest) (*Episode, error)
	ListEpisodes(context.Context, *ListEpisodesRequest) (*EpisodeList, error)
	// tscript is an incomplete transcription
	// chunks are ~2 min sections of the transcription
	GetTscriptChunkStats(context.Context, *emptypb.Empty) (*ChunkStats, error)
	GetTscriptChunk(context.Context, *GetTscriptChunkRequest) (*TscriptChunk, error)
	ListTscriptChunkSubmissions(context.Context, *ListTscriptChunkSubmissionsRequest) (*ChunkSubmissionList, error)
	SubmitTscriptChunk(context.Context, *TscriptChunkSubmissionRequest) (*emptypb.Empty, error)
	SubmitDialogCorrection(context.Context, *SubmitDialogCorrectionRequest) (*emptypb.Empty, error)
}

// UnimplementedSearchServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) Search(context.Context, *SearchRequest) (*SearchResultList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedSearchServiceServer) GetSearchMetadata(context.Context, *emptypb.Empty) (*SearchMetadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearchMetadata not implemented")
}
func (UnimplementedSearchServiceServer) ListFieldValues(context.Context, *ListFieldValuesRequest) (*FieldValueList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFieldValues not implemented")
}
func (UnimplementedSearchServiceServer) GetEpisode(context.Context, *GetEpisodeRequest) (*Episode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEpisode not implemented")
}
func (UnimplementedSearchServiceServer) ListEpisodes(context.Context, *ListEpisodesRequest) (*EpisodeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEpisodes not implemented")
}
func (UnimplementedSearchServiceServer) GetTscriptChunkStats(context.Context, *emptypb.Empty) (*ChunkStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTscriptChunkStats not implemented")
}
func (UnimplementedSearchServiceServer) GetTscriptChunk(context.Context, *GetTscriptChunkRequest) (*TscriptChunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTscriptChunk not implemented")
}
func (UnimplementedSearchServiceServer) ListTscriptChunkSubmissions(context.Context, *ListTscriptChunkSubmissionsRequest) (*ChunkSubmissionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTscriptChunkSubmissions not implemented")
}
func (UnimplementedSearchServiceServer) SubmitTscriptChunk(context.Context, *TscriptChunkSubmissionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTscriptChunk not implemented")
}
func (UnimplementedSearchServiceServer) SubmitDialogCorrection(context.Context, *SubmitDialogCorrectionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitDialogCorrection not implemented")
}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsksearch.SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetSearchMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetSearchMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsksearch.SearchService/GetSearchMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetSearchMetadata(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_ListFieldValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFieldValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).ListFieldValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsksearch.SearchService/ListFieldValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).ListFieldValues(ctx, req.(*ListFieldValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetEpisode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEpisodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetEpisode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsksearch.SearchService/GetEpisode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetEpisode(ctx, req.(*GetEpisodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_ListEpisodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEpisodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).ListEpisodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsksearch.SearchService/ListEpisodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).ListEpisodes(ctx, req.(*ListEpisodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetTscriptChunkStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetTscriptChunkStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsksearch.SearchService/GetTscriptChunkStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetTscriptChunkStats(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_GetTscriptChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTscriptChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).GetTscriptChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsksearch.SearchService/GetTscriptChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).GetTscriptChunk(ctx, req.(*GetTscriptChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_ListTscriptChunkSubmissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTscriptChunkSubmissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).ListTscriptChunkSubmissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsksearch.SearchService/ListTscriptChunkSubmissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).ListTscriptChunkSubmissions(ctx, req.(*ListTscriptChunkSubmissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SubmitTscriptChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TscriptChunkSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SubmitTscriptChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsksearch.SearchService/SubmitTscriptChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SubmitTscriptChunk(ctx, req.(*TscriptChunkSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SubmitDialogCorrection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitDialogCorrectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SubmitDialogCorrection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rsksearch.SearchService/SubmitDialogCorrection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SubmitDialogCorrection(ctx, req.(*SubmitDialogCorrectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rsksearch.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
		{
			MethodName: "GetSearchMetadata",
			Handler:    _SearchService_GetSearchMetadata_Handler,
		},
		{
			MethodName: "ListFieldValues",
			Handler:    _SearchService_ListFieldValues_Handler,
		},
		{
			MethodName: "GetEpisode",
			Handler:    _SearchService_GetEpisode_Handler,
		},
		{
			MethodName: "ListEpisodes",
			Handler:    _SearchService_ListEpisodes_Handler,
		},
		{
			MethodName: "GetTscriptChunkStats",
			Handler:    _SearchService_GetTscriptChunkStats_Handler,
		},
		{
			MethodName: "GetTscriptChunk",
			Handler:    _SearchService_GetTscriptChunk_Handler,
		},
		{
			MethodName: "ListTscriptChunkSubmissions",
			Handler:    _SearchService_ListTscriptChunkSubmissions_Handler,
		},
		{
			MethodName: "SubmitTscriptChunk",
			Handler:    _SearchService_SubmitTscriptChunk_Handler,
		},
		{
			MethodName: "SubmitDialogCorrection",
			Handler:    _SearchService_SubmitDialogCorrection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}
