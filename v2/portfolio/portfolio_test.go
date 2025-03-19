package portfolio

import (
	"testing"
)

const (
	ApiKey    = ""
	SecretKey = ""
)

func TestGetBalanceService_Do(t *testing.T) {

	client := NewClient(ApiKey, SecretKey)
	res, err := client.NewGetBalanceService().Do(newContext())
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestDemo2(t *testing.T) {

	client := NewClient(ApiKey, SecretKey)
	res, err := client.NewGetPositionRiskService().Do(newContext())
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestDemo3(t *testing.T) {
	client := NewClient(ApiKey, SecretKey)
	res, err := client.NewListOpenOrdersService().Do(newContext())
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestDemo04(t *testing.T) {
	client := NewClient(ApiKey, SecretKey)
	res, err := client.NewCancelOrderService().Symbol("BTCUSDT").OrderID(628950371368).Do(newContext())
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestDemo05(t *testing.T) {
	client := NewClient(ApiKey, SecretKey)
	res, err := client.NewCreateConditionalOrderService().Side(SideTypeBuy).StopPrice("80000").Quantity("0.003").PositionSide(PositionSideTypeShort).Symbol("BTCUSDT").
		TimeInForce(TimeInForceTypeGTC).Type(OrderTypeTakeProfitMarket).
		Do(newContext())
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
