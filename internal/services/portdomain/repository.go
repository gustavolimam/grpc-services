package portdomain

import (
	"database/sql"
	"fmt"
	"grpc-services/internal/models"
	"grpc-services/internal/proto/messagepb"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func InsertData(db *sql.DB, data messagepb.DataInput) error {
	var coordinates []string
	for _, x := range data.Coordinates {
		coordinates = append(coordinates, fmt.Sprint(x))
	}
	query := `INSERT INTO user_data(key_id, name, city, country, alias, regions, coordinates, province, timezone, unlocs, code)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`
	args := []interface{}{data.Key, data.Name, data.City, data.Country, strings.Join(data.Alias, ","), strings.Join(data.Regions, ","), strings.Join(coordinates, ","), data.Province, data.Timezone, strings.Join(data.Unlocs, ","), data.Code}

	if _, err := db.Exec(query, args...); err != nil {
		return err
	}

	return nil
}

func GetData(db *sql.DB, data messagepb.GetInput) ([]*messagepb.Data, error) {
	query := `SELECT 
	id,
	key_id, 
	name, 
	city, 
	country, 
	alias, 
	regions, 
	coordinates, 
	province, 
	timezone, 
	unlocs, 
	code 
	FROM user_data`
	args := []interface{}{}
	if data.Id != "" {
		args = append(args, data.Id)
		query += ` WHERE key_id = ?`
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	res := []*messagepb.Data{}
	for rows.Next() {
		d := models.MessageDB{}
		err = rows.Scan(
			&d.ID,
			&d.Key,
			&d.Name,
			&d.City,
			&d.Country,
			&d.Alias,
			&d.Regions,
			&d.Coordinates,
			&d.Province,
			&d.Timezone,
			&d.Unlocs,
			&d.Code,
		)

		s := strings.Split(d.Coordinates, ",")
		var n []float32
		for _, x := range s {
			nf, err := strconv.ParseFloat(x, 32)
			if err != nil {
				return nil, err
			}
			n = append(n, float32(nf))
		}

		dd := &messagepb.Data{
			Id:          d.ID,
			Key:         d.Key,
			Name:        d.Name,
			City:        d.City,
			Country:     d.Country,
			Alias:       []string{d.Alias},
			Regions:     []string{d.Regions},
			Coordinates: n,
			Province:    d.Province,
			Timezone:    d.Timezone,
			Unlocs:      []string{d.Unlocs},
			Code:        d.Code,
		}
		res = append(res, dd)
	}
	return res, nil
}
