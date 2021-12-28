package dir

import (
	pb_svc "github.com/aglide100/go-simple-file-manager/pb/svc"
)

type DirServer struct {
	pb_svc.DirServer
}

func NewDirServiceController() *DirServer {
	return &DirServer{}
}
