package domain

import (
	"time"
)

// Article is representing the Article data struct
type BMI struct {
	ID        int64     `json:"id"`
	Weight    string    `json:"weight" validate:"required"`
	Height    string    `json:"height" validate:"required"`
	ResultBMI string    `json:"result_bmi" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}
