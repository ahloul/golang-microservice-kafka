package models

type Purchase struct {
	PurchasedAt      string `json:"purchased_at" bson:"purchased_at" jsonapi:"attr,purchased_at"`
	VoucherId        string `json:"voucher_id" bson:"voucher_id" jsonapi:"attr,voucher_id"`
	VoucherPrice     string `json:"voucher_price" bson:"voucher_price" jsonapi:"attr,voucher_price"`
	VoucherValidTo   string `json:"voucher_valid_to" bson:"voucher_valid_to" jsonapi:"attr,voucher_valid_to"`
	VoucherValidFrom string `json:"voucher_price_valid_from" bson:"voucher_valid_from" jsonapi:"attr,voucher_valid_from"`
	VoucherPriceCode string `json:"voucher_price_code" bson:"voucher_price_code" jsonapi:"attr,voucher_price_code"`
	UserId           string `json:"user_id" bson:"user_id" jsonapi:"attr,user_id"`
	UserEmail        string `json:"user_email" bson:"user_email" jsonapi:"attr,user_email"`
	UserTier         string `json:"user_tier" bson:"user_tier" jsonapi:"attr,user_tier"`
}
