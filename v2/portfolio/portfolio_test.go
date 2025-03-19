package portfolio

import (
	"testing"
)

const (
	ApiKey    = "pvokofVECRKOcKdnE7hnfj8ZNf4QMk2j4nfUFNAsWXThEP1mAmZANCyGKMmnQ4yR"
	SecretKey = "Cb6FkMJoGlvQiJ9EF9Hv7dlBmzhHmXs9oFup9pU14biusorBDmZsNB9gwZywhTmb"
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
	res, err := client.NewGetPositionModeService().Do(newContext())
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
