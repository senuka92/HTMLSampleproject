package paymentsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const PaymentsService_CreatePayment_FullMethodName = "/payments.v1.PaymentsService/CreatePayment"

type CreatePaymentRequest struct {
	CardId      string
	StatementId string
	Amount      float64
	PaidAt      string
	Note        string
}

type Payment struct {
	Id          string
	CardId      string
	StatementId string
	Amount      float64
	PaidAt      string
	Note        string
}

func (m *CreatePaymentRequest) GetCardId() string      { return m.CardId }
func (m *CreatePaymentRequest) GetStatementId() string { return m.StatementId }
func (m *CreatePaymentRequest) GetAmount() float64     { return m.Amount }
func (m *CreatePaymentRequest) GetPaidAt() string      { return m.PaidAt }
func (m *CreatePaymentRequest) GetNote() string        { return m.Note }

// PaymentsServiceClient is the client API for PaymentsService service.
type PaymentsServiceClient interface {
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*Payment, error)
}

type paymentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentsServiceClient(cc grpc.ClientConnInterface) PaymentsServiceClient {
	return &paymentsServiceClient{cc}
}

func (c *paymentsServiceClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*Payment, error) {
	out := new(Payment)
	err := c.cc.Invoke(ctx, PaymentsService_CreatePayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentsServiceServer is the server API for PaymentsService service.
type PaymentsServiceServer interface {
	CreatePayment(context.Context, *CreatePaymentRequest) (*Payment, error)
	mustEmbedUnimplementedPaymentsServiceServer()
}

type UnimplementedPaymentsServiceServer struct{}

func (UnimplementedPaymentsServiceServer) CreatePayment(context.Context, *CreatePaymentRequest) (*Payment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedPaymentsServiceServer) mustEmbedUnimplementedPaymentsServiceServer() {}

func RegisterPaymentsServiceServer(s grpc.ServiceRegistrar, srv PaymentsServiceServer) {
	s.RegisterService(&PaymentsService_ServiceDesc, srv)
}

func _PaymentsService_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentsService_CreatePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServiceServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var PaymentsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payments.v1.PaymentsService",
	HandlerType: (*PaymentsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayment",
			Handler:    _PaymentsService_CreatePayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/payments.proto",
}
