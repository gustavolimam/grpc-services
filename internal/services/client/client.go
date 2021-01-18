package client

import (
	"context"
	"encoding/json"
	"grpc-services/internal/models"
	"grpc-services/internal/proto/messagepb"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Service struct {
	grpcService messagepb.MessageServiceClient
}

func NewService(grpc messagepb.MessageServiceClient) Service {
	return Service{
		grpcService: grpc,
	}
}

func (s Service) AddDataFromJson(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("./assets/ports.json")
	if err != nil {
		log.Println("Error - ", err)
		json.NewEncoder(w).Encode(http.StatusBadRequest)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	users := map[string]models.Message{}

	if err := json.Unmarshal(byteValue, &users); err != nil {
		log.Println("Error - ", err)
		json.NewEncoder(w).Encode(http.StatusBadRequest)
	}

	go func() {
		for key, value := range users {
			input := &messagepb.DataInput{
				Name:        value.Name,
				City:        value.City,
				Country:     value.Country,
				Alias:       value.Alias,
				Regions:     value.Regions,
				Coordinates: value.Coordinates,
				Province:    value.Province,
				Timezone:    value.Timezone,
				Unlocs:      value.Unlocs,
				Code:        value.Code,
				Key:         key,
			}

			_, err := s.grpcService.Add(context.Background(), input)
			if err != nil {
				log.Println("ERRO - ", err)
				json.NewEncoder(w).Encode(http.StatusBadRequest)
			}
		}
	}()

	json.NewEncoder(w).Encode(http.StatusOK)
}

func (s Service) GetAllData(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	id := ""
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
	} else {
		id = keys[0]
	}

	input := &messagepb.GetInput{
		Id: id,
	}
	res, err := s.grpcService.Get(context.Background(), input)
	if err != nil {
		json.NewEncoder(w).Encode(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(res)
}
