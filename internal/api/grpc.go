package api

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sarpt/gamedbv/internal/cmds"
	pbDl "github.com/sarpt/gamedbv/pkg/rpc/dl"
	pbIdx "github.com/sarpt/gamedbv/pkg/rpc/idx"
	"google.golang.org/grpc"
)

type service struct {
	address     string
	cmd         cmds.Cmd
	connHandler func(*grpc.ClientConn)
	timeout     time.Duration
}

func (s *Server) closeRPCConnections() error {
	for _, closer := range s.rpcConnClosers {
		_ = closer()
	}

	return nil
}

func (s *Server) dialGRPCServices() error {
	wg := &sync.WaitGroup{}
	s.rpcConnClosers = [](func() error){}
	mtx := &sync.Mutex{}

	// TODO: services below should be stored and initialized on a Server instance.
	// Server instance should manage lifetime/watch the services it constructs/listens to.
	dlCfg := cmds.DlCfg{}
	dlArgs := cmds.DlArguments{
		GRPC: true,
	}
	dlCmd := cmds.NewDl(dlCfg, dlArgs)
	dlService := service{
		address:     dlRPCAddress(s.cfg),
		cmd:         dlCmd,
		connHandler: s.dlConnectionHandler,
		timeout:     s.cfg.RPCDialTimeout,
	}
	wg.Add(1)
	go s.estabilishRPC(dlService, mtx, wg)

	idxCfg := cmds.IdxCfg{}
	idxArgs := cmds.IdxArguments{
		GRPC: true,
	}
	idxCmd := cmds.NewIdx(idxCfg, idxArgs)
	idxService := service{
		address:     idxRPCAddress(s.cfg),
		cmd:         idxCmd,
		connHandler: s.idxConnectionHandler,
		timeout:     s.cfg.RPCDialTimeout,
	}
	wg.Add(1)
	go s.estabilishRPC(idxService, mtx, wg)

	wg.Wait()

	return nil
}

func (s *Server) dlConnectionHandler(conn *grpc.ClientConn) {
	s.dlServiceClient = pbDl.NewDlClient(conn)
}

func (s *Server) estabilishRPC(srvc service, mtx *sync.Mutex, wg *sync.WaitGroup) error {
	defer wg.Done()

	if s.cfg.StartServices {
		err := s.startService(srvc)
		if err != nil {
			return fmt.Errorf("could not start RPC service at %s: %w", srvc.address, err)
		}
	}

	dlConnCloser, err := dialGrpc(srvc)
	if err != nil {
		return fmt.Errorf("could not dial RPC at %s: %w", srvc.address, err)
	}

	mtx.Lock()
	defer mtx.Unlock()

	s.rpcConnClosers = append(s.rpcConnClosers, dlConnCloser)

	return nil
}

func (s *Server) idxConnectionHandler(conn *grpc.ClientConn) {
	s.idxServiceClient = pbIdx.NewIdxClient(conn)
}

func (s *Server) startService(srvc service) error {
	err := srvc.cmd.Start()
	return err
}

func dialGrpc(srvc service) (func() error, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())

	timeoutContext, cancel := context.WithTimeout(context.Background(), srvc.timeout)
	defer cancel()

	conn, err := grpc.DialContext(timeoutContext, srvc.address, opts...)
	if err != nil {
		return nil, fmt.Errorf("grpc dial failure: %w", err)
	}

	srvc.connHandler(conn)

	return conn.Close, nil
}

func dlRPCAddress(cfg Config) string {
	return fmt.Sprintf("%s:%s", cfg.IdxRPCAddress, cfg.IdxRPCPort)
}

func idxRPCAddress(cfg Config) string {
	return fmt.Sprintf("%s:%s", cfg.IdxRPCAddress, cfg.IdxRPCPort)
}
