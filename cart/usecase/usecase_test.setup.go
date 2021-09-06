package usecase

import (
	"testing"

	"github.com/Ralphbaer/hash/cart/gen/mock"
	"github.com/golang/mock/gomock"
)

func setupCartUseCaseMocks(mockCtrl *gomock.Controller, t *testing.T) CartUseCase {
	return CartUseCase{
		Repository: setupCartRepo(mockCtrl),
		Gateway:    setupDiscountServiceGateway(mockCtrl),
	}
}

func setupCartRepo(mockCtrl *gomock.Controller) *mock.MockProductRepository {
	m := mock.NewMockProductRepository(mockCtrl)
	m.
		EXPECT().
		List(gomock.Any(), giftProductFilter).
		Return(giftScenario, nil).
		AnyTimes()
	m.
		EXPECT().
		List(gomock.Any(), gomock.Any()).
		Return(happyScenario, nil).
		AnyTimes()
	m.
		EXPECT().
		FindRandom(gomock.Any()).
		Return(products[6], nil).
		AnyTimes()

	return m
}

func setupDiscountServiceGateway(mockCtrl *gomock.Controller) *mock.MockDiscountServiceGateway {
	m := mock.NewMockDiscountServiceGateway(mockCtrl)

	m.
		EXPECT().
		GetDiscount(gomock.Any(), gomock.Eq(int32(1))).
		Return(float32(0.05), nil).
		AnyTimes()
	m.
		EXPECT().
		GetDiscount(gomock.Any(), gomock.Eq(int32(2))).
		Return(float32(0.1), nil).
		AnyTimes()
	m.
		EXPECT().
		GetDiscount(gomock.Any(), gomock.Any()).
		Return(float32(0.0), nil).
		AnyTimes()

	return m
}
