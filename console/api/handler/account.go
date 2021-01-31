package handler

import (
	"context"

	"github.com/hb-go/pkg/conv"
	mApi "github.com/stack-labs/stack-rpc/api"
	hApi "github.com/stack-labs/stack-rpc/api/handler/api"
	"github.com/stack-labs/stack-rpc/api/handler/rpc"
	api "github.com/stack-labs/stack-rpc/api/proto"
	"github.com/stack-labs/stack-rpc/server"
	"github.com/stack-labs/stack-rpc/util/errors"
	"github.com/stack-labs/stack-rpc/util/log"

	"github.com/stack-labs/starter-kit/console/api/client"
	pb "github.com/stack-labs/starter-kit/console/api/genproto/api"
	account "github.com/stack-labs/starter-kit/console/api/genproto/srv"
)

type Account struct{}

// Example.Call is called by the API as /example/call with post body {"name": "foo"}
func (*Account) Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.Response) error {
	log.Log("Received Example.Call request")

	if err := req.Validate(); err != nil {
		return errors.BadRequest("stack.rpc.api.example.example.call", err.Error())
	}

	// extract the client from the context
	ac, ok := client.AccountFromContext(ctx)
	if !ok {
		return errors.InternalServerError("stack.rpc.api.example.example.call", "example client not found")
	}

	// make request
	r := &account.LoginRequest{}
	conv.StructToStruct(req, r)
	response, err := ac.Login(ctx, r)
	if err != nil {
		return errors.InternalServerError("stack.rpc.api.example.example.call", err.Error())
	}

	rsp.Code = 20000
	rsp.Data = response

	return nil
}

// Example.Call is called by the API as /example/call with post body {"name": "foo"}
func (*Account) Logout(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Example.Call request")

	// extract the client from the context
	ac, ok := client.AccountFromContext(ctx)
	if !ok {
		return errors.InternalServerError("stack.rpc.api.example.example.call", "example client not found")
	}

	// make request
	response, err := ac.Logout(ctx, &account.Request{
		Id: 0,
	})
	if err != nil {
		return errors.InternalServerError("stack.rpc.api.example.example.call", err.Error())
	}

	b, err := ResponseBody(20000, response)
	if err != nil {
		return errors.InternalServerError("stack.rpc.api.example.example.call", err.Error())
	}

	rsp.StatusCode = 200
	rsp.Body = b

	return nil
}

// Example.Call is called by the API as /example/call with post body {"name": "foo"}
func (*Account) Info(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Example.Call request")

	// extract the client from the context
	ac, ok := client.AccountFromContext(ctx)
	if !ok {
		return errors.InternalServerError("stack.rpc.api.example.example.call", "example client not found")
	}

	// make request
	response, err := ac.Info(ctx, &account.Request{
		Id: 0,
	})
	if err != nil {
		return errors.InternalServerError("stack.rpc.api.example.example.call", err.Error())
	}

	b, err := ResponseBody(20000, response)
	if err != nil {
		return errors.InternalServerError("stack.rpc.api.example.example.call", err.Error())
	}

	rsp.StatusCode = 200
	rsp.Body = b

	return nil
}

func registerAccount(server server.Server) {
	pb.RegisterAccountHandler(server, new(Account),
		mApi.WithEndpoint(&mApi.Endpoint{
			// The RPC method
			Name: "Account.Login",
			// The HTTP paths. This can be a POSIX regex
			Path: []string{"^/account/login$"},
			// The HTTP Methods for this endpoint
			Method: []string{"POST"},
			// The API handler to use
			Handler: rpc.Handler,
		}),
		mApi.WithEndpoint(&mApi.Endpoint{
			// The RPC method
			Name: "Account.Logout",
			// The HTTP paths. This can be a POSIX regex
			Path: []string{"^/account/logout$"},
			// The HTTP Methods for this endpoint
			Method: []string{"POST"},
			// The API handler to use
			Handler: hApi.Handler,
		}),
		mApi.WithEndpoint(&mApi.Endpoint{
			// The RPC method
			Name: "Account.Info",
			// The HTTP paths. This can be a POSIX regex
			Path: []string{"^/account/info$"},
			// The HTTP Methods for this endpoint
			Method: []string{"GET"},
			// The API handler to use
			Handler: hApi.Handler,
		}),
	)
}
