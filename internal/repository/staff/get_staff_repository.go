package staffRepository

import (
	"context"
	"eniqlo/internal/entity"
	"fmt"
)

func (sr staffRepository) GetStaffByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.Staff, error) {
	var (
		resp entity.Staff
		err  error
	)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE "phone_number" = $1`, resp.TableName())

	err = sr.db.GetContext(ctx, &resp, query, phoneNumber)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
