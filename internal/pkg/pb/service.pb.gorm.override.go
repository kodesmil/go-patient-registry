package pb

import (
	"context"
	"fmt"
	"github.com/infobloxopen/atlas-app-toolkit/query"
)

func (m *ServicesDefaultServer) CustomListServiceSession(ctx context.Context, req *ListServiceSessionRequest) (*ListServiceSessionResponse, error) {
	accountID := fmt.Sprintf("%v", ctx.Value("AccountID"))
	res, err := DefaultListServiceSession(ctx, m.DB, &query.Filtering{
		Root: &query.Filtering_StringCondition{
			StringCondition: &query.StringCondition{
				FieldPath: []string{"profile_id"},
				Value:     accountID,
				Type:      query.StringCondition_EQ,
			},
		},
	}, req.OrderBy, req.Paging, req.Fields)
	if err != nil {
		return nil, err
	}
	return &ListServiceSessionResponse{Results: res}, nil
}

func (m *ServicesDefaultServer) CustomListServiceOfferSession(ctx context.Context, req *ListServiceOfferSessionRequest) (*ListServiceOfferSessionResponse, error) {
	db := m.DB.
		Joins("left join service_offers so on service_sessions.service_offer_id::uuid = so.id").
		Joins("left join service_employments se on so.service_employment_id = se.id")
	accountID := fmt.Sprintf("%v", ctx.Value("AccountID"))
	res, err := DefaultListServiceSession(ctx, db, &query.Filtering{
		Root: &query.Filtering_StringCondition{
			StringCondition: &query.StringCondition{
				FieldPath: []string{"se", "profile_id"},
				Value:     accountID,
				Type:      query.StringCondition_EQ,
			},
		},
	}, req.OrderBy, req.Paging, req.Fields)
	if err != nil {
		return nil, err
	}
	return &ListServiceOfferSessionResponse{Results: res}, nil
}
