package alertsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const (
	AlertsService_ScheduleAlert_FullMethodName = "/alerts.v1.AlertsService/ScheduleAlert"
	DashboardService_GetSummary_FullMethodName = "/alerts.v1.DashboardService/GetSummary"
)

type ScheduleAlertRequest struct {
	CardId      string
	StatementId string
	DueDate     string
	DueAmount   float64
}

type Alert struct {
	Id          string
	CardId      string
	StatementId string
	DueDate     string
	DueAmount   float64
	Status      string
}

type GetSummaryRequest struct{}

type DashboardSummary struct {
	TotalDue       float64
	TotalPaid      float64
	TotalRemaining float64
	OverdueCount   int32
}

func (m *ScheduleAlertRequest) GetCardId() string      { return m.CardId }
func (m *ScheduleAlertRequest) GetStatementId() string { return m.StatementId }
func (m *ScheduleAlertRequest) GetDueDate() string     { return m.DueDate }
func (m *ScheduleAlertRequest) GetDueAmount() float64  { return m.DueAmount }
func (m *Alert) GetDueAmount() float64                 { return m.DueAmount }
func (m *Alert) GetDueDate() string                    { return m.DueDate }

// AlertsServiceClient is the client API for AlertsService service.
type AlertsServiceClient interface {
	ScheduleAlert(ctx context.Context, in *ScheduleAlertRequest, opts ...grpc.CallOption) (*Alert, error)
}

type alertsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertsServiceClient(cc grpc.ClientConnInterface) AlertsServiceClient {
	return &alertsServiceClient{cc}
}

func (c *alertsServiceClient) ScheduleAlert(ctx context.Context, in *ScheduleAlertRequest, opts ...grpc.CallOption) (*Alert, error) {
	out := new(Alert)
	err := c.cc.Invoke(ctx, AlertsService_ScheduleAlert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DashboardServiceClient is the client API for DashboardService service.
type DashboardServiceClient interface {
	GetSummary(ctx context.Context, in *GetSummaryRequest, opts ...grpc.CallOption) (*DashboardSummary, error)
}

type dashboardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDashboardServiceClient(cc grpc.ClientConnInterface) DashboardServiceClient {
	return &dashboardServiceClient{cc}
}

func (c *dashboardServiceClient) GetSummary(ctx context.Context, in *GetSummaryRequest, opts ...grpc.CallOption) (*DashboardSummary, error) {
	out := new(DashboardSummary)
	err := c.cc.Invoke(ctx, DashboardService_GetSummary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertsServiceServer is the server API for AlertsService service.
type AlertsServiceServer interface {
	ScheduleAlert(context.Context, *ScheduleAlertRequest) (*Alert, error)
	mustEmbedUnimplementedAlertsServiceServer()
}

type UnimplementedAlertsServiceServer struct{}

func (UnimplementedAlertsServiceServer) ScheduleAlert(context.Context, *ScheduleAlertRequest) (*Alert, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleAlert not implemented")
}
func (UnimplementedAlertsServiceServer) mustEmbedUnimplementedAlertsServiceServer() {}

// DashboardServiceServer is the server API for DashboardService service.
type DashboardServiceServer interface {
	GetSummary(context.Context, *GetSummaryRequest) (*DashboardSummary, error)
	mustEmbedUnimplementedDashboardServiceServer()
}

type UnimplementedDashboardServiceServer struct{}

func (UnimplementedDashboardServiceServer) GetSummary(context.Context, *GetSummaryRequest) (*DashboardSummary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSummary not implemented")
}
func (UnimplementedDashboardServiceServer) mustEmbedUnimplementedDashboardServiceServer() {}

func RegisterAlertsServiceServer(s grpc.ServiceRegistrar, srv AlertsServiceServer) {
	s.RegisterService(&AlertsService_ServiceDesc, srv)
}

func RegisterDashboardServiceServer(s grpc.ServiceRegistrar, srv DashboardServiceServer) {
	s.RegisterService(&DashboardService_ServiceDesc, srv)
}

func _AlertsService_ScheduleAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServiceServer).ScheduleAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertsService_ScheduleAlert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServiceServer).ScheduleAlert(ctx, req.(*ScheduleAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardService_GetSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).GetSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardService_GetSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).GetSummary(ctx, req.(*GetSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var AlertsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alerts.v1.AlertsService",
	HandlerType: (*AlertsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScheduleAlert",
			Handler:    _AlertsService_ScheduleAlert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/alerts.proto",
}

var DashboardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alerts.v1.DashboardService",
	HandlerType: (*DashboardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSummary",
			Handler:    _DashboardService_GetSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/alerts.proto",
}
