package cardsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const (
	CardsService_CreateBank_FullMethodName      = "/cards.v1.CardsService/CreateBank"
	CardsService_CreateCard_FullMethodName      = "/cards.v1.CardsService/CreateCard"
	CardsService_CreateStatement_FullMethodName = "/cards.v1.CardsService/CreateStatement"
)

type CreateBankRequest struct {
	Name string
}

type CreateCardRequest struct {
	BankId   string
	Name     string
	LastFour string
	DueDay   string
}

type CreateStatementRequest struct {
	CardId         string
	StatementMonth string
	DueDate        string
	DueAmount      float64
}

type Bank struct {
	Id   string
	Name string
}

type Card struct {
	Id       string
	BankId   string
	Name     string
	LastFour string
	DueDay   string
}

type Statement struct {
	Id             string
	CardId         string
	StatementMonth string
	DueDate        string
	DueAmount      float64
}

func (m *CreateBankRequest) GetName() string                { return m.Name }
func (m *CreateCardRequest) GetBankId() string              { return m.BankId }
func (m *CreateCardRequest) GetName() string                { return m.Name }
func (m *CreateCardRequest) GetLastFour() string            { return m.LastFour }
func (m *CreateCardRequest) GetDueDay() string              { return m.DueDay }
func (m *CreateStatementRequest) GetCardId() string         { return m.CardId }
func (m *CreateStatementRequest) GetStatementMonth() string { return m.StatementMonth }
func (m *CreateStatementRequest) GetDueDate() string        { return m.DueDate }
func (m *CreateStatementRequest) GetDueAmount() float64     { return m.DueAmount }

// CardsServiceClient is the client API for CardsService service.
type CardsServiceClient interface {
	CreateBank(ctx context.Context, in *CreateBankRequest, opts ...grpc.CallOption) (*Bank, error)
	CreateCard(ctx context.Context, in *CreateCardRequest, opts ...grpc.CallOption) (*Card, error)
	CreateStatement(ctx context.Context, in *CreateStatementRequest, opts ...grpc.CallOption) (*Statement, error)
}

type cardsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCardsServiceClient(cc grpc.ClientConnInterface) CardsServiceClient {
	return &cardsServiceClient{cc}
}

func (c *cardsServiceClient) CreateBank(ctx context.Context, in *CreateBankRequest, opts ...grpc.CallOption) (*Bank, error) {
	out := new(Bank)
	err := c.cc.Invoke(ctx, CardsService_CreateBank_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsServiceClient) CreateCard(ctx context.Context, in *CreateCardRequest, opts ...grpc.CallOption) (*Card, error) {
	out := new(Card)
	err := c.cc.Invoke(ctx, CardsService_CreateCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsServiceClient) CreateStatement(ctx context.Context, in *CreateStatementRequest, opts ...grpc.CallOption) (*Statement, error) {
	out := new(Statement)
	err := c.cc.Invoke(ctx, CardsService_CreateStatement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardsServiceServer is the server API for CardsService service.
type CardsServiceServer interface {
	CreateBank(context.Context, *CreateBankRequest) (*Bank, error)
	CreateCard(context.Context, *CreateCardRequest) (*Card, error)
	CreateStatement(context.Context, *CreateStatementRequest) (*Statement, error)
	mustEmbedUnimplementedCardsServiceServer()
}

type UnimplementedCardsServiceServer struct{}

func (UnimplementedCardsServiceServer) CreateBank(context.Context, *CreateBankRequest) (*Bank, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBank not implemented")
}
func (UnimplementedCardsServiceServer) CreateCard(context.Context, *CreateCardRequest) (*Card, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCard not implemented")
}
func (UnimplementedCardsServiceServer) CreateStatement(context.Context, *CreateStatementRequest) (*Statement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStatement not implemented")
}
func (UnimplementedCardsServiceServer) mustEmbedUnimplementedCardsServiceServer() {}

func RegisterCardsServiceServer(s grpc.ServiceRegistrar, srv CardsServiceServer) {
	s.RegisterService(&CardsService_ServiceDesc, srv)
}

func _CardsService_CreateBank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServiceServer).CreateBank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CardsService_CreateBank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServiceServer).CreateBank(ctx, req.(*CreateBankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardsService_CreateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServiceServer).CreateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CardsService_CreateCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServiceServer).CreateCard(ctx, req.(*CreateCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardsService_CreateStatement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStatementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServiceServer).CreateStatement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CardsService_CreateStatement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServiceServer).CreateStatement(ctx, req.(*CreateStatementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var CardsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cards.v1.CardsService",
	HandlerType: (*CardsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBank",
			Handler:    _CardsService_CreateBank_Handler,
		},
		{
			MethodName: "CreateCard",
			Handler:    _CardsService_CreateCard_Handler,
		},
		{
			MethodName: "CreateStatement",
			Handler:    _CardsService_CreateStatement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cards.proto",
}
