package incomev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const IncomeService_UpsertIncome_FullMethodName = "/income.v1.IncomeService/UpsertIncome"

type UpsertIncomeRequest struct {
	Month    string
	Currency string
	Amount   float64
	Source   string
	Note     string
}

type Income struct {
	Id       string
	Month    string
	Currency string
	Amount   float64
	Source   string
	Note     string
}

func (m *UpsertIncomeRequest) GetMonth() string    { return m.Month }
func (m *UpsertIncomeRequest) GetCurrency() string { return m.Currency }
func (m *UpsertIncomeRequest) GetAmount() float64  { return m.Amount }
func (m *UpsertIncomeRequest) GetSource() string   { return m.Source }
func (m *UpsertIncomeRequest) GetNote() string     { return m.Note }

// IncomeServiceClient is the client API for IncomeService service.
type IncomeServiceClient interface {
	UpsertIncome(ctx context.Context, in *UpsertIncomeRequest, opts ...grpc.CallOption) (*Income, error)
}

type incomeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIncomeServiceClient(cc grpc.ClientConnInterface) IncomeServiceClient {
	return &incomeServiceClient{cc}
}

func (c *incomeServiceClient) UpsertIncome(ctx context.Context, in *UpsertIncomeRequest, opts ...grpc.CallOption) (*Income, error) {
	out := new(Income)
	err := c.cc.Invoke(ctx, IncomeService_UpsertIncome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IncomeServiceServer is the server API for IncomeService service.
type IncomeServiceServer interface {
	UpsertIncome(context.Context, *UpsertIncomeRequest) (*Income, error)
	mustEmbedUnimplementedIncomeServiceServer()
}

type UnimplementedIncomeServiceServer struct{}

func (UnimplementedIncomeServiceServer) UpsertIncome(context.Context, *UpsertIncomeRequest) (*Income, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertIncome not implemented")
}
func (UnimplementedIncomeServiceServer) mustEmbedUnimplementedIncomeServiceServer() {}

func RegisterIncomeServiceServer(s grpc.ServiceRegistrar, srv IncomeServiceServer) {
	s.RegisterService(&IncomeService_ServiceDesc, srv)
}

func _IncomeService_UpsertIncome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertIncomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncomeServiceServer).UpsertIncome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncomeService_UpsertIncome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncomeServiceServer).UpsertIncome(ctx, req.(*UpsertIncomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var IncomeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "income.v1.IncomeService",
	HandlerType: (*IncomeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertIncome",
			Handler:    _IncomeService_UpsertIncome_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/income.proto",
}
