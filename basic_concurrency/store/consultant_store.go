package store

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
)

type Consultant struct {
	Id              uuid.UUID `db:"id" json:"id"`
	Slug            string    `db:"slug" json:"slug"`
	ConsultantFName string    `db:"consultant_f_name" json:"consultant_f_name"`
	ConsultantLName string    `db:"consultant_l_name" json:"consultant_l_name"`
	ImgPath         string    `db:"img_path" json:"img_path"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}
