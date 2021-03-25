package api

import (
	"context"
	"fmt"

	pbDl "github.com/sarpt/gamedbv/pkg/rpc/dl"
	pbIdx "github.com/sarpt/gamedbv/pkg/rpc/idx"
	"google.golang.org/grpc"
)

func (s *Server) dialGrpcServices() (func(), error) {
	closeDialConn, err := s.dialDlGrpc()
	if err != nil {
		return nil, fmt.Errorf("could not dial Dl RPC: %w", err)
	}

	closeIdxConn, err := s.dialIdxGrpc()
	if err != nil {
		return nil, fmt.Errorf("could not dial Idx RPC: %w", err)
	}

	return func() {
		err := closeDialConn()
		if err != nil {
			s.errLog.Printf("could not close Dl GRPC connection: %v", err)
		}

		err = closeIdxConn()
		if err != nil {
			s.errLog.Printf("could not close Idx GRPC connection: %v", err)
		}
	}, nil
}

func (s *Server) dialDlGrpc() (func() error, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())

	timeoutContext, cancel := context.WithTimeout(context.Background(), s.cfg.RPCDialTimeout)
	defer cancel()

	dlcRPCTarget := fmt.Sprintf("%s:%s", s.cfg.DlRPCAddress, s.cfg.DlRPCPort)
	conn, err := grpc.DialContext(timeoutContext, dlcRPCTarget, opts...) // TODO: option to start own process from api
	if err != nil {
		return nil, fmt.Errorf("grpc dial failure: %w", err)
	}

	s.dlServiceClient = pbDl.NewDlClient(conn)

	return conn.Close, nil
}

func (s *Server) dialIdxGrpc() (func() error, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())

	timeoutContext, cancel := context.WithTimeout(context.Background(), s.cfg.RPCDialTimeout)
	defer cancel()

	idxRPCTarget := fmt.Sprintf("%s:%s", s.cfg.IdxRPCAddress, s.cfg.IdxRPCPort)
	conn, err := grpc.DialContext(timeoutContext, idxRPCTarget, opts...)
	if err != nil {
		return nil, fmt.Errorf("grpc dial failure: %w", err)
	}

	s.idxServiceClient = pbIdx.NewIdxClient(conn)

	return conn.Close, nil
}
