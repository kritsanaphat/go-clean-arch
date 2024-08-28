package mysql

import (
	"context"
	"database/sql"

	"github.com/bxcodec/go-clean-arch/domain"
)

type BMIRepository struct {
	Conn *sql.DB
}

// NewArticleRepository will create an object that represent the article.Repository interface
func NewBMIRepository(conn *sql.DB) *BMIRepository {
	return &BMIRepository{conn}
}

func (b *BMIRepository) StoreBMILog(ctx context.Context, a *domain.BMI) (err error) {
	query := `INSERT  bmi SET weight=? , height=? , result_bmi=? , created_at=?`
	stmt, err := b.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, a.Weight, a.Height, a.ResultBMI, a.CreatedAt)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	a.ID = lastID
	return
}
