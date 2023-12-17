package store

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
)

type Consultant struct {
	Id              uuid.UUID `db:"id"`
	Slug            string    `db:"slug"`
	ConsultantFName string    `db:"consultant_f_name"`
	ConsultantLName string    `db:"consultant_l_name"`
	ImgPath         string    `db:"img_path"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
