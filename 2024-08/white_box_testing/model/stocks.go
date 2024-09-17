package model

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	dateFormat = "2006-01-02"
)

type PriceData struct {
	ID        int
	Timestamp time.Time
	Price     float64
}

type PriceProvider interface {
	Latest() (*PriceData, error)
	List(date time.Time) ([]*PriceData, error)
}

type priceProvider struct {
	db *sql.DB
}

func NewPriceProvider(db *sql.DB) PriceProvider {
	return &priceProvider{
		db: db,
	}
}

func (p *priceProvider) Latest() (*PriceData, error) {
	var priceData PriceData

	err := p.db.QueryRow("SELECT * FROM stockprices ORDER BY timestamp DESC limit 1").Scan(&priceData.Timestamp, &priceData.Price)
	if err != nil {
		return &priceData, fmt.Errorf("unable to query table. Error %s", err.Error())
	}

	return &priceData, nil

}

func (p *priceProvider) List(date time.Time) ([]*PriceData, error) {
	priceData := make([]*PriceData, 0)

	var rows *sql.Rows
	var err error

	// 时间戳倒序查找股票价格
	rows, err = p.db.Query("SELECT * FROM stockprices where DATE(timestamp) = ? ORDER BY timestamp DESC",
		date.Format(dateFormat))

	if err != nil {
		return priceData, fmt.Errorf("unable to prepare SELECT statement. Error %s", err.Error())
	}

	var id int
	var timestamp time.Time
	var price float64

	for rows.Next() {
		err = rows.Scan(&id, &timestamp, &price)
		if err != nil {
			return priceData, fmt.Errorf("unable to query table. Error %s", err.Error())
		}

		priceData = append(priceData, &PriceData{
			ID:        id,
			Timestamp: timestamp,
			Price:     price,
		})
	}

	return priceData, nil

}
