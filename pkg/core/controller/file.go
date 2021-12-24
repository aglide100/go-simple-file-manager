package controller

import (
	"context"
	"errors"
	"log"

	pb_svc "github.com/aglide100/go-simple-file-manager/pb/svc"
)

type FileServer struct {
	pb_svc.FileServer
}

func NewFileServiceController() *FileServer {
	return &FileServer{}
}

func (s *FileServer) DeleteFile(ctx context.Context, in *pb_svc.DeleteFileReq) (*pb_svc.DeleteFileRes, error) {
	log.Println("received DeleteFileReq", in)

	if (in == nil) {
		return nil, errors.New("DeleteFileReq is nil! please check request")
	}

	return &pb_svc.DeleteFileRes{}, nil
}