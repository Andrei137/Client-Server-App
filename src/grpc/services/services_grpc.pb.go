// protoc --go_out=. --go-grpc_out=. services.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: services.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Services_MixWords_FullMethodName                = "/services.Services/MixWords"
	Services_CountPerfectSquares_FullMethodName     = "/services.Services/CountPerfectSquares"
	Services_SumReversed_FullMethodName             = "/services.Services/SumReversed"
	Services_AverageWithinRange_FullMethodName      = "/services.Services/AverageWithinRange"
	Services_ConvertBinary_FullMethodName           = "/services.Services/ConvertBinary"
	Services_CaesarCipher_FullMethodName            = "/services.Services/CaesarCipher"
	Services_DecodeText_FullMethodName              = "/services.Services/DecodeText"
	Services_CountDigitsPrime_FullMethodName        = "/services.Services/CountDigitsPrime"
	Services_CountEvenVowels_FullMethodName         = "/services.Services/CountEvenVowels"
	Services_GCD_FullMethodName                     = "/services.Services/GCD"
	Services_SumPermuted_FullMethodName             = "/services.Services/SumPermuted"
	Services_SumDuplicatedFirst_FullMethodName      = "/services.Services/SumDuplicatedFirst"
	Services_FilterComplexNumbers_FullMethodName    = "/services.Services/FilterComplexNumbers"
	Services_FilterValidPasswords_FullMethodName    = "/services.Services/FilterValidPasswords"
	Services_GenerateRandomPasswords_FullMethodName = "/services.Services/GenerateRandomPasswords"
)

// ServicesClient is the client API for Services service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesClient interface {
	MixWords(ctx context.Context, in *StringSliceReq, opts ...grpc.CallOption) (*StringSliceRes, error)
	CountPerfectSquares(ctx context.Context, in *StringSliceReq, opts ...grpc.CallOption) (*Int32Res, error)
	SumReversed(ctx context.Context, in *Int32SliceReq, opts ...grpc.CallOption) (*Int32Res, error)
	AverageWithinRange(ctx context.Context, in *Int32SliceReq, opts ...grpc.CallOption) (*Int32Res, error)
	ConvertBinary(ctx context.Context, in *StringSliceReq, opts ...grpc.CallOption) (*Int32SliceRes, error)
	CaesarCipher(ctx context.Context, in *CaesarCipherReq, opts ...grpc.CallOption) (*StringRes, error)
	DecodeText(ctx context.Context, in *StringReq, opts ...grpc.CallOption) (*StringRes, error)
	CountDigitsPrime(ctx context.Context, in *Int32SliceReq, opts ...grpc.CallOption) (*Int32Res, error)
	CountEvenVowels(ctx context.Context, in *StringSliceReq, opts ...grpc.CallOption) (*Int32Res, error)
	GCD(ctx context.Context, in *StringSliceReq, opts ...grpc.CallOption) (*Int32Res, error)
	SumPermuted(ctx context.Context, in *SumPermutedReq, opts ...grpc.CallOption) (*Int32Res, error)
	SumDuplicatedFirst(ctx context.Context, in *Int32SliceReq, opts ...grpc.CallOption) (*Int32Res, error)
	FilterComplexNumbers(ctx context.Context, in *Int32SliceReq, opts ...grpc.CallOption) (*DoubleSliceRes, error)
	FilterValidPasswords(ctx context.Context, in *StringSliceReq, opts ...grpc.CallOption) (*StringSliceRes, error)
	GenerateRandomPasswords(ctx context.Context, in *StringSliceReq, opts ...grpc.CallOption) (*StringSliceRes, error)
}

type servicesClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesClient(cc grpc.ClientConnInterface) ServicesClient {
	return &servicesClient{cc}
}

func (c *servicesClient) MixWords(ctx context.Context, in *StringSliceReq, opts ...grpc.CallOption) (*StringSliceRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StringSliceRes)
	err := c.cc.Invoke(ctx, Services_MixWords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) CountPerfectSquares(ctx context.Context, in *StringSliceReq, opts ...grpc.CallOption) (*Int32Res, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Int32Res)
	err := c.cc.Invoke(ctx, Services_CountPerfectSquares_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) SumReversed(ctx context.Context, in *Int32SliceReq, opts ...grpc.CallOption) (*Int32Res, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Int32Res)
	err := c.cc.Invoke(ctx, Services_SumReversed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) AverageWithinRange(ctx context.Context, in *Int32SliceReq, opts ...grpc.CallOption) (*Int32Res, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Int32Res)
	err := c.cc.Invoke(ctx, Services_AverageWithinRange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) ConvertBinary(ctx context.Context, in *StringSliceReq, opts ...grpc.CallOption) (*Int32SliceRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Int32SliceRes)
	err := c.cc.Invoke(ctx, Services_ConvertBinary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) CaesarCipher(ctx context.Context, in *CaesarCipherReq, opts ...grpc.CallOption) (*StringRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StringRes)
	err := c.cc.Invoke(ctx, Services_CaesarCipher_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) DecodeText(ctx context.Context, in *StringReq, opts ...grpc.CallOption) (*StringRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StringRes)
	err := c.cc.Invoke(ctx, Services_DecodeText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) CountDigitsPrime(ctx context.Context, in *Int32SliceReq, opts ...grpc.CallOption) (*Int32Res, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Int32Res)
	err := c.cc.Invoke(ctx, Services_CountDigitsPrime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) CountEvenVowels(ctx context.Context, in *StringSliceReq, opts ...grpc.CallOption) (*Int32Res, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Int32Res)
	err := c.cc.Invoke(ctx, Services_CountEvenVowels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) GCD(ctx context.Context, in *StringSliceReq, opts ...grpc.CallOption) (*Int32Res, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Int32Res)
	err := c.cc.Invoke(ctx, Services_GCD_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) SumPermuted(ctx context.Context, in *SumPermutedReq, opts ...grpc.CallOption) (*Int32Res, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Int32Res)
	err := c.cc.Invoke(ctx, Services_SumPermuted_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) SumDuplicatedFirst(ctx context.Context, in *Int32SliceReq, opts ...grpc.CallOption) (*Int32Res, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Int32Res)
	err := c.cc.Invoke(ctx, Services_SumDuplicatedFirst_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) FilterComplexNumbers(ctx context.Context, in *Int32SliceReq, opts ...grpc.CallOption) (*DoubleSliceRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DoubleSliceRes)
	err := c.cc.Invoke(ctx, Services_FilterComplexNumbers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) FilterValidPasswords(ctx context.Context, in *StringSliceReq, opts ...grpc.CallOption) (*StringSliceRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StringSliceRes)
	err := c.cc.Invoke(ctx, Services_FilterValidPasswords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) GenerateRandomPasswords(ctx context.Context, in *StringSliceReq, opts ...grpc.CallOption) (*StringSliceRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StringSliceRes)
	err := c.cc.Invoke(ctx, Services_GenerateRandomPasswords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicesServer is the server API for Services service.
// All implementations must embed UnimplementedServicesServer
// for forward compatibility.
type ServicesServer interface {
	MixWords(context.Context, *StringSliceReq) (*StringSliceRes, error)
	CountPerfectSquares(context.Context, *StringSliceReq) (*Int32Res, error)
	SumReversed(context.Context, *Int32SliceReq) (*Int32Res, error)
	AverageWithinRange(context.Context, *Int32SliceReq) (*Int32Res, error)
	ConvertBinary(context.Context, *StringSliceReq) (*Int32SliceRes, error)
	CaesarCipher(context.Context, *CaesarCipherReq) (*StringRes, error)
	DecodeText(context.Context, *StringReq) (*StringRes, error)
	CountDigitsPrime(context.Context, *Int32SliceReq) (*Int32Res, error)
	CountEvenVowels(context.Context, *StringSliceReq) (*Int32Res, error)
	GCD(context.Context, *StringSliceReq) (*Int32Res, error)
	SumPermuted(context.Context, *SumPermutedReq) (*Int32Res, error)
	SumDuplicatedFirst(context.Context, *Int32SliceReq) (*Int32Res, error)
	FilterComplexNumbers(context.Context, *Int32SliceReq) (*DoubleSliceRes, error)
	FilterValidPasswords(context.Context, *StringSliceReq) (*StringSliceRes, error)
	GenerateRandomPasswords(context.Context, *StringSliceReq) (*StringSliceRes, error)
	mustEmbedUnimplementedServicesServer()
}

// UnimplementedServicesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServicesServer struct{}

func (UnimplementedServicesServer) MixWords(context.Context, *StringSliceReq) (*StringSliceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MixWords not implemented")
}
func (UnimplementedServicesServer) CountPerfectSquares(context.Context, *StringSliceReq) (*Int32Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountPerfectSquares not implemented")
}
func (UnimplementedServicesServer) SumReversed(context.Context, *Int32SliceReq) (*Int32Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumReversed not implemented")
}
func (UnimplementedServicesServer) AverageWithinRange(context.Context, *Int32SliceReq) (*Int32Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AverageWithinRange not implemented")
}
func (UnimplementedServicesServer) ConvertBinary(context.Context, *StringSliceReq) (*Int32SliceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertBinary not implemented")
}
func (UnimplementedServicesServer) CaesarCipher(context.Context, *CaesarCipherReq) (*StringRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CaesarCipher not implemented")
}
func (UnimplementedServicesServer) DecodeText(context.Context, *StringReq) (*StringRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeText not implemented")
}
func (UnimplementedServicesServer) CountDigitsPrime(context.Context, *Int32SliceReq) (*Int32Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountDigitsPrime not implemented")
}
func (UnimplementedServicesServer) CountEvenVowels(context.Context, *StringSliceReq) (*Int32Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountEvenVowels not implemented")
}
func (UnimplementedServicesServer) GCD(context.Context, *StringSliceReq) (*Int32Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GCD not implemented")
}
func (UnimplementedServicesServer) SumPermuted(context.Context, *SumPermutedReq) (*Int32Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumPermuted not implemented")
}
func (UnimplementedServicesServer) SumDuplicatedFirst(context.Context, *Int32SliceReq) (*Int32Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumDuplicatedFirst not implemented")
}
func (UnimplementedServicesServer) FilterComplexNumbers(context.Context, *Int32SliceReq) (*DoubleSliceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterComplexNumbers not implemented")
}
func (UnimplementedServicesServer) FilterValidPasswords(context.Context, *StringSliceReq) (*StringSliceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterValidPasswords not implemented")
}
func (UnimplementedServicesServer) GenerateRandomPasswords(context.Context, *StringSliceReq) (*StringSliceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateRandomPasswords not implemented")
}
func (UnimplementedServicesServer) mustEmbedUnimplementedServicesServer() {}
func (UnimplementedServicesServer) testEmbeddedByValue()                  {}

// UnsafeServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesServer will
// result in compilation errors.
type UnsafeServicesServer interface {
	mustEmbedUnimplementedServicesServer()
}

func RegisterServicesServer(s grpc.ServiceRegistrar, srv ServicesServer) {
	// If the following call pancis, it indicates UnimplementedServicesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Services_ServiceDesc, srv)
}

func _Services_MixWords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringSliceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).MixWords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_MixWords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).MixWords(ctx, req.(*StringSliceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_CountPerfectSquares_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringSliceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).CountPerfectSquares(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_CountPerfectSquares_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).CountPerfectSquares(ctx, req.(*StringSliceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_SumReversed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32SliceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).SumReversed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_SumReversed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).SumReversed(ctx, req.(*Int32SliceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_AverageWithinRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32SliceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).AverageWithinRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_AverageWithinRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).AverageWithinRange(ctx, req.(*Int32SliceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_ConvertBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringSliceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).ConvertBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_ConvertBinary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).ConvertBinary(ctx, req.(*StringSliceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_CaesarCipher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaesarCipherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).CaesarCipher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_CaesarCipher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).CaesarCipher(ctx, req.(*CaesarCipherReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_DecodeText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).DecodeText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_DecodeText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).DecodeText(ctx, req.(*StringReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_CountDigitsPrime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32SliceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).CountDigitsPrime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_CountDigitsPrime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).CountDigitsPrime(ctx, req.(*Int32SliceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_CountEvenVowels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringSliceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).CountEvenVowels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_CountEvenVowels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).CountEvenVowels(ctx, req.(*StringSliceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_GCD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringSliceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).GCD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_GCD_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).GCD(ctx, req.(*StringSliceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_SumPermuted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumPermutedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).SumPermuted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_SumPermuted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).SumPermuted(ctx, req.(*SumPermutedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_SumDuplicatedFirst_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32SliceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).SumDuplicatedFirst(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_SumDuplicatedFirst_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).SumDuplicatedFirst(ctx, req.(*Int32SliceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_FilterComplexNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32SliceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).FilterComplexNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_FilterComplexNumbers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).FilterComplexNumbers(ctx, req.(*Int32SliceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_FilterValidPasswords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringSliceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).FilterValidPasswords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_FilterValidPasswords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).FilterValidPasswords(ctx, req.(*StringSliceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_GenerateRandomPasswords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringSliceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).GenerateRandomPasswords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Services_GenerateRandomPasswords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).GenerateRandomPasswords(ctx, req.(*StringSliceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Services_ServiceDesc is the grpc.ServiceDesc for Services service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Services_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.Services",
	HandlerType: (*ServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MixWords",
			Handler:    _Services_MixWords_Handler,
		},
		{
			MethodName: "CountPerfectSquares",
			Handler:    _Services_CountPerfectSquares_Handler,
		},
		{
			MethodName: "SumReversed",
			Handler:    _Services_SumReversed_Handler,
		},
		{
			MethodName: "AverageWithinRange",
			Handler:    _Services_AverageWithinRange_Handler,
		},
		{
			MethodName: "ConvertBinary",
			Handler:    _Services_ConvertBinary_Handler,
		},
		{
			MethodName: "CaesarCipher",
			Handler:    _Services_CaesarCipher_Handler,
		},
		{
			MethodName: "DecodeText",
			Handler:    _Services_DecodeText_Handler,
		},
		{
			MethodName: "CountDigitsPrime",
			Handler:    _Services_CountDigitsPrime_Handler,
		},
		{
			MethodName: "CountEvenVowels",
			Handler:    _Services_CountEvenVowels_Handler,
		},
		{
			MethodName: "GCD",
			Handler:    _Services_GCD_Handler,
		},
		{
			MethodName: "SumPermuted",
			Handler:    _Services_SumPermuted_Handler,
		},
		{
			MethodName: "SumDuplicatedFirst",
			Handler:    _Services_SumDuplicatedFirst_Handler,
		},
		{
			MethodName: "FilterComplexNumbers",
			Handler:    _Services_FilterComplexNumbers_Handler,
		},
		{
			MethodName: "FilterValidPasswords",
			Handler:    _Services_FilterValidPasswords_Handler,
		},
		{
			MethodName: "GenerateRandomPasswords",
			Handler:    _Services_GenerateRandomPasswords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
