package api

import (
	"context"
	"fmt"
	"sync"

	"google.golang.org/grpc"
)

func (s *Server) closeGRPCConnections() error {
	err := s.dlService.Close()
	if err != nil {
		s.errLog.Printf("could not close Dl service: %v", err)

		return err
	}

	s.idxService.Close()
	if err != nil {
		s.errLog.Printf("could not close Idx service: %v", err)

		return err
	}

	s.outLog.Printf("GRPC connections closed")

	return nil
}

func (s *Server) dialGRPCServices() error {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go s.estabilishGRPC(s.dlService, wg) // TODO: establishGRPC returns error that is ignored here. Implement proper error propagation (through channel?)

	wg.Add(1)
	go s.estabilishGRPC(s.idxService, wg)

	wg.Wait()

	return nil
}

func (s *Server) estabilishGRPC(srvc Service, wg *sync.WaitGroup) error {
	defer wg.Done()

	if s.cfg.StartServices {
		err := srvc.Start()
		if err != nil {
			return fmt.Errorf("could not start GRPC service at %s: %w", srvc.Address(), err)
		}
	}
	s.outLog.Printf("started GRPC service at %s", srvc.Address())

	err := dialGrpc(srvc)
	if err != nil {
		return fmt.Errorf("could not dial GRPC at %s: %w", srvc.Address(), err)
	}
	s.outLog.Printf("connected to GRPC service at %s", srvc.Address())

	return nil
}

func dialGrpc(srvc Service) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())

	timeoutContext, cancel := context.WithTimeout(context.Background(), srvc.Timeout())
	defer cancel()

	conn, err := grpc.DialContext(timeoutContext, srvc.Address(), opts...)
	if err != nil {
		return fmt.Errorf("GRPC dial failure: %w", err)
	}

	srvc.SetClient(conn)

	return nil
}
