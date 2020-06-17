package pb

import (
	"context"
)

func (m *HealthDefaultServer) CustomListHealthMenstruationDailyEntry(ctx context.Context, req *ListHealthMenstruationDailyEntryRequest) (*ListHealthMenstruationDailyEntryResponse, error) {
	res, err := DefaultListHealthMenstruationDailyEntry(ctx, m.DB, req.Filter, req.OrderBy, req.Paging, req.Fields)
	if err != nil {
		return nil, err
	}
	return &ListHealthMenstruationDailyEntryResponse{Results: res}, nil
}
