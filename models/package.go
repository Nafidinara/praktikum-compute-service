package models

// Data paket terdiri dari ID dengan tipe angka, nama, nama pengirim, nama penerima, lokasi pengirim, lokasi penerima, biaya dan berat paket.

type Package struct {
	ID               int     `json:"id" form:"id"`
	Name             string  `json:"name" form:"name"`
	Sender           string  `json:"sender" form:"sender"`
	Receiver         string  `json:"receiver" form:"receiver"`
	SenderLocation   string  `json:"sender_location" form:"sender_location"`
	ReceiverLocation string  `json:"receiver_location" form:"receiver_location"`
	Fee              float64 `json:"fee" form:"fee"`
	Weight           float64 `json:"weight" form:"weight"`
}
