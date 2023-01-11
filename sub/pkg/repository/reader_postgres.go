package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"nats"
	"time"
)

type ReaderPostgres struct {
	db *sqlx.DB
}

func NewReaderPostgres(db *sqlx.DB) *ReaderPostgres {
	return &ReaderPostgres{db: db}
}

func (r *ReaderPostgres) Recover() ([]nats.Order, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var orders []nats.Order
	query :=
		fmt.Sprintf("select order_uid, track_number, entry, locale, internal_signature, "+
			"customer_id, delivery_service, shardkey,sm_id, date_created, oof_shard from %s", ordersTable)
	err := r.db.SelectContext(ctx, &orders, query)
	if err != nil {
		return orders, err
	}

	for _, order := range orders {

		delivery, err := r.getDel(order.OrderUid)
		if err != nil {
			return orders, err
		}

		payment, err := r.getPayment(order.OrderUid)
		if err != nil {
			return orders, err
		}

		items, err := r.getItems(order.OrderUid)
		if err != nil {
			return orders, err
		}

		order.Delivery = delivery
		order.Payment = payment
		order.Items = items

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *ReaderPostgres) getPayment(id string) (nats.Pay, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var payment nats.Pay
	query := fmt.Sprintf("select transaction, request_id, currency, provider, amount, payment_dt, "+
		"bank, delivery_cost, goods_total,custom_fee from %s where order_id = $1", paymentTable)

	err := r.db.GetContext(ctx, &payment, query, id)

	if err != nil && err != sql.ErrNoRows {
		return payment, err
	}

	return payment, nil
}

func (r *ReaderPostgres) getDel(id string) (nats.Del, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var delivery nats.Del
	query := fmt.Sprintf("select name, phone, zip, city, address, region, "+
		"email from %s where order_id = $1", deliveryTable)

	err := r.db.GetContext(ctx, &delivery, query, id)

	if err != nil && err != sql.ErrNoRows {
		return delivery, err
	}

	return delivery, nil
}

func (r *ReaderPostgres) getItems(id string) ([]nats.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var items []nats.Item
	query := fmt.Sprintf("select  chrt_id, track_number, price, rid, name, sale, size, total_price,"+
		" nm_id, brand, status from %s where order_id = $1",
		itemsTable)

	err := r.db.SelectContext(ctx, &items, query, id)
	if err != nil {
		return items, err
	}

	return items, nil
}
