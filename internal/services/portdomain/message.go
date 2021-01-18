package portdomain

import (
	"context"
	"database/sql"
	"grpc-services/internal/proto/messagepb"
	"log"
)

type Service struct {
	DB *sql.DB
}

func (s *Service) Add(ctx context.Context, input *messagepb.DataInput) (*messagepb.DataResponse, error) {
	log.Println("Start to add a new data: ", input)

	if err := InsertData(s.DB, *input); err != nil {
		return &messagepb.DataResponse{
			Message: "Failed",
			Error:   err.Error(),
		}, err
	}
	return &messagepb.DataResponse{
		Message: "SUCESSO",
		Error:   "",
	}, nil
}

func (s *Service) Get(ctx context.Context, input *messagepb.GetInput) (*messagepb.GetResponse, error) {
	log.Println("Start to return list of data")

	res, err := GetData(s.DB, *input)
	if err != nil {
		return &messagepb.GetResponse{}, err
	}
	return &messagepb.GetResponse{
		Data: res,
	}, nil
}
