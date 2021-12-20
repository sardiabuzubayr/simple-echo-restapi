package model

import "mime/multipart"

type Users struct {
	Email       string          `json:"email" form:"email"`
	Nama        string          `json:"nama" form:"nama"`
	NoHandphone string          `json:"no_handphone" form:"no_handphone"`
	ALamat      string          `json:"alamat" form:"alamat"`
	Ktp         *multipart.File `json:"ktp" form:"ktp"`
}
