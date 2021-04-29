package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/sarpt/gamedbv/pkg/rpc/dl"
	"github.com/sarpt/gamedbv/pkg/rpc/idx"
)

type operation string

const (
	startOp operation = "start"
	closeOp operation = "close"
)

type operationHandler = func(interface{}, io.Writer) error

func (s *Server) getOperationHandlers() map[operation]operationHandler {
	return map[operation]operationHandler{
		startOp: s.handleStartOperation,
	}
}

func (s *Server) handleOperationMessage(msg clientOpertionMessage, w io.Writer) error {
	handler, ok := s.operationHandlers[msg.Op]
	if !ok {
		return fmt.Errorf("no handler for the '%s' operation", msg.Op)
	}

	err := handler(msg.Payload, w)
	return err
}

func (s *Server) handleStartOperation(payload interface{}, w io.Writer) error {
	startPayload, ok := payload.(startPayload)
	if !ok || len(startPayload.Platforms) < 1 {
		return fmt.Errorf("incorrect payload for start operation")
	}

	wg := sync.WaitGroup{}
	for _, platform := range startPayload.Platforms {
		wg.Add(1)

		go func(platform string) {
			defer wg.Done()

			err := s.updatePlatform(platform, w)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Update for platform %s failed: %v", platform, err) // tbd: tee writer
				fmt.Fprintf(w, "Update for platform %s failed", platform)                  // tbd: error writer
				return
			}

			status, err := json.Marshal(PlatformUpdateEndStatus(platform))
			if err != nil {
				fmt.Fprintf(w, "Error writing done status for platform %s", platform) // tbd: error writer
				return
			}

			fmt.Fprintf(w, string(status))
		}(platform)
	}

	wg.Wait()

	return nil
}

func (s Server) updatePlatform(platform string, w io.Writer) error {
	dlReq := dl.PlatformsDownloadReq{
		Platforms: []string{platform},
	}

	dlCtx, dlCancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer dlCancel()

	dlStream, err := s.dlService.DownloadPlatforms(dlCtx, &dlReq)
	if err != nil {
		return fmt.Errorf("could not download platforms through grpc: %w", err)
	}

	for {
		platDlStatus, err := dlStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("error while receiving response through grpc: %w", err)
		}

		res, err := json.Marshal(&platDlStatus)
		if err != nil {
			return fmt.Errorf("could not marshal grpc json response for writer: %w", err)
		}

		w.Write(res)
	}

	idxReq := idx.PreparePlatformsReq{
		Platforms: []string{platform},
	}

	idxCtx, idxCancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer idxCancel()

	stream, err := s.idxService.PreparePlatforms(idxCtx, &idxReq)
	if err != nil {
		return fmt.Errorf("could not prepare platforms through grpc: %w", err)
	}

	for {
		idxPlatStatus, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("error while receiving response through grpc: %w", err)
		}

		res, err := json.Marshal(&idxPlatStatus)
		if err != nil {
			return fmt.Errorf("could not marshal grpc json response for writer: %w", err)
		}

		w.Write(res)
	}

	return nil
}
