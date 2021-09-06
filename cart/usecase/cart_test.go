package usecase

import (
	"context"
	"reflect"
	"testing"

	e "github.com/Ralphbaer/hash/cart/entity"
	"github.com/golang/mock/gomock"
)

func TestCartUseCase_CreateCart(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	uc := setupCartUseCaseMocks(mockCtrl, t)
	got, err := uc.Create(context.TODO(), createCartInput)
	if err != nil {
		t.Error(err)
		return
	}

	m := struct {
		want *e.Cart
		got  *e.Cart
	}{
		want: &e.Cart{
			TotalAmount:             308000,
			TotalAmountWithDiscount: 286400,
			TotalDiscount:           21600,
			Products: []e.ProductAmount{
				{
					ID:          1,
					Quantity:    2,
					UnitAmount:  60000,
					TotalAmount: 120000,
					Discount:    6000,
					IsGift:      false,
				},
				{
					ID:          2,
					Quantity:    3,
					UnitAmount:  52000,
					TotalAmount: 156000,
					Discount:    15600,
					IsGift:      false,
				},
				{
					ID:          3,
					Quantity:    1,
					UnitAmount:  32000,
					TotalAmount: 32000,
					Discount:    0,
					IsGift:      false,
				},
			},
		},
		got: &e.Cart{
			TotalAmount:             got.TotalAmount,
			TotalAmountWithDiscount: got.TotalAmountWithDiscount,
			TotalDiscount:           got.TotalDiscount,
			Products:                got.Products,
		},
	}

	if !reflect.DeepEqual(m.want, m.got) {
		t.Error("Got:", m.got, "Want:", m.want)
		return
	}
}

func TestCartUseCase_CreateCart_ErrCantBuyAGift(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	uc := setupCartUseCaseMocks(mockCtrl, t)
	_, got := uc.Create(context.TODO(), createCartWithGiftInput)
	if got == nil {
		t.Error("Got:", got, ",want: not nil")
		return
	}

	want := ErrCantBuyAGift
	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Error("Got:", got, ",want:", want)
		return
	}
}
