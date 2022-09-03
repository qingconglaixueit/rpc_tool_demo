/*
____  ___   _____ ___________
\   \/  /  /     \\__    ___/
 \     /  /  \ /  \ |    |
 /     \ /    Y    \|    |
/___/\  \\____|__  /|____|
      \_/        \/
createTime:2022/9/3
createUser:Administrator
*/
package services

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"rpc_tool_demo/protoc/tenant"
)

type RPCDemo struct{
	tenant.TenantServer
}

func StartGrpc() {
	listen, err := net.Listen("tcp", ":8899")
	if err != nil {
		log.Fatal(err)
	}

	// 实例化grpc servers
	s := grpc.NewServer()
	reflection.Register(s)
	// 注册Love服务
	tenant.RegisterTenantServer(s, new(RPCDemo))

	log.Println("Listen on 127.0.0.1:8899...")
	s.Serve(listen)
}

func (r *RPCDemo) GetTenantDetail(context.Context, *tenant.TenantDetailReq) (*tenant.TenantDetailRsp, error) {
	return &tenant.TenantDetailRsp{
	Id: "001",
	Name: "伟大的公司",
	Addr: "未来的路",
	}, nil
}
func (r *RPCDemo) GetTenantList(context.Context, *tenant.TenantListReq) (*tenant.TenantListRsp, error) {
	return &tenant.TenantListRsp{
		Name: []string{"伟大的公司","理发店","小猪佩奇宠物店"},
	}, nil
}
