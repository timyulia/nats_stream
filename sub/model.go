package nats

import "time"

type Order struct {
	OrderUid          string    `db:"order_uid" json:"order_uid" validate:"required"`
	TrackNumber       string    `db:"track_number" json:"track_number" validate:"required"`
	Entry             string    `db:"entry" json:"entry" validate:"required"`
	Delivery          Del       `db:"delivery" json:"delivery" validate:"required"`
	Items             []Item    `db:"items" json:"items" validate:"required"`
	Payment           Pay       `db:"payment" json:"payment" validate:"required"`
	Locale            string    `db:"locale" json:"locale" validate:"required"`
	InternalSignature *string   `db:"internal_signature" json:"internal_signature" validate:"required"`
	CustomerId        string    `db:"customer_id" json:"customer_id" validate:"required"`
	DeliveryService   string    `db:"delivery_service" json:"delivery_service" validate:"required"`
	Shardkey          string    `db:"shardkey" json:"shardkey" validate:"required"`
	SmId              int       `db:"sm_id" json:"sm_id" validate:"min=0"`
	DateCreated       time.Time `db:"date_created" json:"date_created" validate:"required"`
	OofShard          string    `db:"oof_shard" json:"oof_shard" validate:"required"`
}

type Del struct {
	OrderId string `db:"order-id"`
	Name    string `db:"name" json:"name" validate:"required"`
	Phone   string `db:"phone" json:"phone" validate:"required"`
	Zip     string `db:"zip" json:"zip" validate:"required"`
	City    string `db:"city" json:"city" validate:"required"`
	Address string `db:"address" json:"address" validate:"required"`
	Region  string `db:"region" json:"region" validate:"required"`
	Email   string `db:"email" json:"email" validate:"required,email"`
}

type Item struct {
	Id          int     `db:"id"`
	OrderId     string  `db:"order-id"`
	ChrtId      int     `db:"chrt_id" json:"chrt_id" validate:"min=0"`
	TrackNumber string  `db:"track_number" json:"track_number" validate:"required"`
	Price       float64 `db:"price" json:"price" validate:"min=0"`
	Rid         string  `db:"rid" json:"rid" validate:"required"`
	Name        string  `db:"name" json:"name" validate:"required"`
	Sale        int     `db:"sale" json:"sale" validate:"min=0"`
	Size        string  `db:"size" json:"size" validate:"required"`
	TotalPrice  float64 `db:"total_price" json:"total_price" validate:"min=0"`
	NmId        int     `db:"nm_id" json:"nm_id" validate:"min=0"`
	Brand       string  `db:"brand" json:"brand" validate:"required"`
	Status      int     `db:"status" json:"status" validate:"min=0"`
}

type Pay struct {
	OrderId      string  `db:"order-id"`
	Transaction  string  `db:"transaction" json:"transaction" validate:"required"`
	RequestId    *string `db:"request_id" json:"request_id" validate:"required"`
	Currency     string  `db:"currency" json:"currency" validate:"required"`
	Provider     string  `db:"provider" json:"provider" validate:"required"`
	Amount       float64 `db:"amount" json:"amount" validate:"required"`
	PaymentDt    int     `db:"payment_dt" json:"payment_dt" validate:"required"`
	Bank         string  `db:"bank" json:"bank" validate:"required"`
	DeliveryCost float64 `db:"delivery_cost" json:"delivery_cost" validate:"required"`
	GoodsTotal   int     `db:"goods_total" json:"goods_total" validate:"min=0"`
	CustomFee    float64 `db:"custom_fee" json:"custom_fee" validate:"min=0"`
}

type PayDTO struct {
	Transaction  string  `db:"transaction" json:"transaction" validate:"required"`
	RequestId    *string `db:"request_id" json:"request_id" validate:"required"`
	Currency     string  `db:"currency" json:"currency" validate:"required"`
	Provider     string  `db:"provider" json:"provider" validate:"required"`
	Amount       float64 `db:"amount" json:"amount" validate:"required"`
	PaymentDt    int     `db:"payment_dt" json:"payment_dt" validate:"required"`
	Bank         string  `db:"bank" json:"bank" validate:"required"`
	DeliveryCost float64 `db:"delivery_cost" json:"delivery_cost" validate:"required"`
	GoodsTotal   int     `db:"goods_total" json:"goods_total" validate:"min=0"`
	CustomFee    float64 `db:"custom_fee" json:"custom_fee" validate:"min=0"`
}

type ItemDTO struct {
	ChrtId      int     `db:"chrt_id" json:"chrt_id" validate:"min=0"`
	TrackNumber string  `db:"track_number" json:"track_number" validate:"required"`
	Price       float64 `db:"price" json:"price" validate:"min=0"`
	Rid         string  `db:"rid" json:"rid" validate:"required"`
	Name        string  `db:"name" json:"name" validate:"required"`
	Sale        int     `db:"sale" json:"sale" validate:"min=0"`
	Size        string  `db:"size" json:"size" validate:"required"`
	TotalPrice  float64 `db:"total_price" json:"total_price" validate:"min=0"`
	NmId        int     `db:"nm_id" json:"nm_id" validate:"min=0"`
	Brand       string  `db:"brand" json:"brand" validate:"required"`
	Status      int     `db:"status" json:"status" validate:"min=0"`
}

type DelDTO struct {
	Name    string `db:"name" json:"name" validate:"required"`
	Phone   string `db:"phone" json:"phone" validate:"required"`
	Zip     string `db:"zip" json:"zip" validate:"required"`
	City    string `db:"city" json:"city" validate:"required"`
	Address string `db:"address" json:"address" validate:"required"`
	Region  string `db:"region" json:"region" validate:"required"`
	Email   string `db:"email" json:"email" validate:"required,email"`
}

type OrderDTO struct {
	OrderUid          string    `db:"order_uid" json:"order_uid" validate:"required"`
	TrackNumber       string    `db:"track_number" json:"track_number" validate:"required"`
	Entry             string    `db:"entry" json:"entry" validate:"required"`
	Delivery          DelDTO    `db:"delivery" json:"delivery" validate:"required"`
	Items             []ItemDTO `db:"items" json:"items" validate:"required"`
	Payment           PayDTO    `db:"payment" json:"payment" validate:"required"`
	Locale            string    `db:"locale" json:"locale" validate:"required"`
	InternalSignature *string   `db:"internal_signature" json:"internal_signature" validate:"required"`
	CustomerId        string    `db:"customer_id" json:"customer_id" validate:"required"`
	DeliveryService   string    `db:"delivery_service" json:"delivery_service" validate:"required"`
	Shardkey          string    `db:"shardkey" json:"shardkey" validate:"required"`
	SmId              int       `db:"sm_id" json:"sm_id" validate:"min=0"`
	DateCreated       time.Time `db:"date_created" json:"date_created" validate:"required"`
	OofShard          string    `db:"oof_shard" json:"oof_shard" validate:"required"`
}
