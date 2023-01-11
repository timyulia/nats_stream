package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"nats"
	"time"
	_ "time"
)

type CreatorPostgres struct {
	db *sqlx.DB
}

func NewCreatorPostgres(db *sqlx.DB) *CreatorPostgres {
	return &CreatorPostgres{db: db}
}

func (r *CreatorPostgres) Create(ord *nats.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("insert into %s (order_uid, track_number, entry, locale, internal_signature,"+
		" customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)"+
		"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", ordersTable)
	_, err = tx.ExecContext(ctx,
		query,
		ord.OrderUid,
		ord.TrackNumber,
		ord.Entry,
		ord.Locale,
		ord.InternalSignature,
		ord.CustomerId,
		ord.DeliveryService,
		ord.Shardkey,
		ord.SmId,
		ord.DateCreated,
		ord.OofShard,
	)
	if err != nil {
		tx.Rollback()
		return err

	}

	query = fmt.Sprintf("insert into %s (order_id, name, phone, zip, city, address, region, email) "+
		"values ($1, $2, $3, $4, $5, $6, $7, $8)", deliveryTable)
	_, err = tx.ExecContext(ctx,
		query,
		ord.OrderUid,
		ord.Delivery.Name,
		ord.Delivery.Phone,
		ord.Delivery.Zip,
		ord.Delivery.City,
		ord.Delivery.Address,
		ord.Delivery.Region,
		ord.Delivery.Email,
	)
	if err != nil {
		tx.Rollback()
		return err

	}

	query = fmt.Sprintf("insert into %s (order_id, transaction, request_id, currency, provider, "+
		"amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) values"+
		"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", paymentTable)
	_, err = tx.ExecContext(ctx,
		query,
		ord.OrderUid,
		ord.Payment.Transaction,
		ord.Payment.RequestId,
		ord.Payment.Currency,
		ord.Payment.Provider,
		ord.Payment.Amount,
		ord.Payment.PaymentDt,
		ord.Payment.Bank,
		ord.Payment.DeliveryCost,
		ord.Payment.GoodsTotal,
		ord.Payment.CustomFee,
	)
	if err != nil {
		tx.Rollback()
		return err

	}

	for _, item := range ord.Items {
		query = fmt.Sprintf("insert into %s  (order_id,chrt_id, track_number, price, rid, name, sale, size,"+
			" total_price, nm_id, brand, status) values"+
			"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)", itemsTable)
		_, err = tx.ExecContext(ctx,
			query,
			ord.OrderUid,
			item.ChrtId,
			item.TrackNumber,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmId,
			item.Brand,
			item.Status,
		)
		if err != nil {
			tx.Rollback()
			return err
		}

	}

	return tx.Commit()
}
