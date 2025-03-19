package portfolio

import (
	"github.com/stretchr/testify/assert"
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
	res, err := client.NewCreateOrderService().Symbol("BTCUSDT").Side(SideTypeSell).PositionSide(PositionSideTypeShort).Type(OrderTypeLimit).TimeInForce(TimeInForceTypeGTC).Quantity("0.003").Price("100000").Do(newContext())
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestDemo06(t *testing.T) {
	client := NewClient(ApiKey, SecretKey)
	res, err := client.NewCreateOrderService().TimeInForce(TimeInForceTypeGTC).Symbol("GMXUSDT").Side(SideTypeSell).Type(OrderTypeTakeProfitMarket).Quantity("1").StopPrice("30").Price("31").Do(ctx)
	assert.Empty(t, err)
}
