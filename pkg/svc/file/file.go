package controller

import (
	"context"
	"errors"
	"log"

	pb_svc "github.com/aglide100/go-simple-file-manager/pb/svc"
	pb_unit_file "github.com/aglide100/go-simple-file-manager/pb/unit/file"
)

type FileServer struct {
	pb_svc.FileServer
}

func NewFileServiceController() *FileServer {
	return &FileServer{}
}

func (s *FileServer) GetFile(ctx context.Context, in *pb_svc.GetFileReq) (*pb_svc.GetFileRes, error) {
	log.Println("received GetFileReq", in)

	if (in == nil) {
		return nil, errors.New("GetFileReq is nil! please check request!")
	}

	file := &pb_unit_file.File{Spec: &pb_unit_file.Spec{

	}}

	return &pb_svc.GetFileRes{File: file}, nil

}

func (s *FileServer) DeleteFile(ctx context.Context, in *pb_svc.DeleteFileReq) (*pb_svc.DeleteFileRes, error) {
	log.Println("received DeleteFileReq", in)

	if (in == nil) {
		return nil, errors.New("DeleteFileReq is nil! please check request")
	}

	return &pb_svc.DeleteFileRes{}, nil
}
