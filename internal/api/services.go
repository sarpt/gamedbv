package api

import (
	"fmt"
	"time"

	"github.com/sarpt/gamedbv/internal/cmds"
	pbDl "github.com/sarpt/gamedbv/pkg/rpc/dl"
	pbIdx "github.com/sarpt/gamedbv/pkg/rpc/idx"
	"google.golang.org/grpc"
)

type Service interface {
	Address() string
	Close() error
	SetClient(*grpc.ClientConn)
	Start() error
	Timeout() time.Duration
}

type base struct {
	address string
	cmds.Cmd
	connection *grpc.ClientConn
	timeout    time.Duration
}

func (b *base) Address() string {
	return b.address
}

func (b *base) Timeout() time.Duration {
	return b.timeout
}

func (b *base) Close() error {
	return b.connection.Close()
}

type DlService struct {
	*base
	pbDl.DlClient
}

func NewDlService(cfg Config) *DlService {
	dlCfg := cmds.DlCfg{}
	dlArgs := cmds.DlArguments{
		GRPC: true,
	}
	dlCmd := cmds.NewDl(dlCfg, dlArgs)

	return &DlService{
		base: &base{
			address: fmt.Sprintf("%s:%s", cfg.DlRPCAddress, cfg.DlRPCPort),
			Cmd:     dlCmd,
			timeout: cfg.RPCDialTimeout,
		},
	}
}

func (dl *DlService) SetClient(connection *grpc.ClientConn) {
	dl.connection = connection
	dl.DlClient = pbDl.NewDlClient(connection)
}

type IdxService struct {
	*base
	pbIdx.IdxClient
}

func NewIdxService(cfg Config) *IdxService {
	idxCfg := cmds.IdxCfg{}
	idxArgs := cmds.IdxArguments{
		GRPC: true,
	}
	idxCmd := cmds.NewIdx(idxCfg, idxArgs)
	return &IdxService{
		base: &base{
			address: fmt.Sprintf("%s:%s", cfg.IdxRPCAddress, cfg.IdxRPCPort),
			Cmd:     idxCmd,
			timeout: cfg.RPCDialTimeout,
		},
	}
}

func (idx *IdxService) SetClient(connection *grpc.ClientConn) {
	idx.connection = connection
	idx.IdxClient = pbIdx.NewIdxClient(connection)
}
