package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// CreateConditionalOrderService create order
type CreateConditionalOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	positionSide            *PositionSideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	quantity                string
	reduceOnly              *string
	price                   *string
	newClientOrderID        *string
	stopPrice               *string
	workingType             *WorkingType
	activationPrice         *string
	callbackRate            *string
	priceProtect            *string
	selfTradePreventionMode *SelfTradePreventionMode
}

// Symbol set symbol
func (s *CreateConditionalOrderService) Symbol(symbol string) *CreateConditionalOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateConditionalOrderService) Side(side SideType) *CreateConditionalOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *CreateConditionalOrderService) PositionSide(positionSide PositionSideType) *CreateConditionalOrderService {
	s.positionSide = &positionSide
	return s
}

// Type set type
func (s *CreateConditionalOrderService) Type(orderType OrderType) *CreateConditionalOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateConditionalOrderService) TimeInForce(timeInForce TimeInForceType) *CreateConditionalOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CreateConditionalOrderService) Quantity(quantity string) *CreateConditionalOrderService {
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateConditionalOrderService) ReduceOnly(reduceOnly bool) *CreateConditionalOrderService {
	reduceOnlyStr := strconv.FormatBool(reduceOnly)
	s.reduceOnly = &reduceOnlyStr
	return s
}

// Price set price
func (s *CreateConditionalOrderService) Price(price string) *CreateConditionalOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CreateConditionalOrderService) NewClientOrderID(newClientOrderID string) *CreateConditionalOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// StopPrice set stopPrice
func (s *CreateConditionalOrderService) StopPrice(stopPrice string) *CreateConditionalOrderService {
	s.stopPrice = &stopPrice
	return s
}

// WorkingType set workingType
func (s *CreateConditionalOrderService) WorkingType(workingType WorkingType) *CreateConditionalOrderService {
	s.workingType = &workingType
	return s
}

// ActivationPrice set activationPrice
func (s *CreateConditionalOrderService) ActivationPrice(activationPrice string) *CreateConditionalOrderService {
	s.activationPrice = &activationPrice
	return s
}

// CallbackRate set callbackRate
func (s *CreateConditionalOrderService) CallbackRate(callbackRate string) *CreateConditionalOrderService {
	s.callbackRate = &callbackRate
	return s
}

// PriceProtect set priceProtect
func (s *CreateConditionalOrderService) PriceProtect(priceProtect bool) *CreateConditionalOrderService {
	priceProtectStr := strconv.FormatBool(priceProtect)
	s.priceProtect = &priceProtectStr
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *CreateConditionalOrderService) SelfTradePreventionMode(selfTradePreventionMode SelfTradePreventionMode) *CreateConditionalOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

func (s *CreateConditionalOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":       s.symbol,
		"side":         s.side,
		"strategyType": s.orderType,
	}
	if s.quantity != "" {
		m["quantity"] = s.quantity
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.newClientOrderID != nil {
		m["newClientStrategyId"] = *s.newClientOrderID
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
	}
	if s.workingType != nil {
		m["workingType"] = *s.workingType
	}
	if s.priceProtect != nil {
		m["priceProtect"] = *s.priceProtect
	}
	if s.activationPrice != nil {
		m["activationPrice"] = *s.activationPrice
	}
	if s.callbackRate != nil {
		m["callbackRate"] = *s.callbackRate
	}
	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}
	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *CreateConditionalOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateConditionalOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/papi/v1/um/conditional/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateConditionalOrderResponse)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateConditionalOrderResponse define create order response
type CreateConditionalOrderResponse struct {
	Symbol                  string           `json:"symbol"`                      //
	OrderID                 int64            `json:"orderId"`                     //
	ClientOrderID           string           `json:"clientOrderId"`               //
	Price                   string           `json:"price"`                       //
	OrigQuantity            string           `json:"origQty"`                     //
	ExecutedQuantity        string           `json:"executedQty"`                 //
	CumQuote                string           `json:"cumQuote"`                    //
	ReduceOnly              bool             `json:"reduceOnly"`                  //
	Status                  OrderStatusType  `json:"status"`                      //
	StopPrice               string           `json:"stopPrice"`                   // please ignore when order type is TRAILING_STOP_MARKET
	TimeInForce             TimeInForceType  `json:"timeInForce"`                 //
	Type                    OrderType        `json:"type"`                        //
	Side                    SideType         `json:"side"`                        //
	UpdateTime              int64            `json:"updateTime"`                  // update time
	WorkingType             WorkingType      `json:"workingType"`                 //
	ActivatePrice           string           `json:"activatePrice"`               // activation price, only return with TRAILING_STOP_MARKET order
	PriceRate               string           `json:"priceRate"`                   // callback rate, only return with TRAILING_STOP_MARKET order
	AvgPrice                string           `json:"avgPrice"`                    //
	PositionSide            PositionSideType `json:"positionSide"`                //
	ClosePosition           bool             `json:"closePosition"`               // if Close-All
	PriceProtect            bool             `json:"priceProtect"`                // if conditional order trigger is protected
	PriceMatch              string           `json:"priceMatch"`                  // price match mode
	SelfTradePreventionMode string           `json:"selfTradePreventionMode"`     // self trading prevention mode
	GoodTillDate            int64            `json:"goodTillDate"`                // order pre-set auto cancel time for TIF GTD order
	CumQty                  string           `json:"cumQty"`                      //
	OrigType                OrderType        `json:"origType"`                    //
	RateLimitOrder10s       string           `json:"rateLimitOrder10s,omitempty"` //
	RateLimitOrder1m        string           `json:"rateLimitOrder1m,omitempty"`  //
}

// ListConditionalOpenOrdersService list opened orders
type ListConditionalOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *ListConditionalOpenOrdersService) Symbol(symbol string) *ListConditionalOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *ListConditionalOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*ConditionalOrder, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/conditional/openOrders",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*ConditionalOrder{}, err
	}
	res = make([]*ConditionalOrder, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*ConditionalOrder{}, err
	}
	return res, nil
}

type ConditionalOrder struct {
	NewClientStrategyId     string `json:"newClientStrategyId"`
	StrategyId              int    `json:"strategyId"`
	StrategyStatus          string `json:"strategyStatus"`
	StrategyType            string `json:"strategyType"`
	OrigQty                 string `json:"origQty"`
	Price                   string `json:"price"`
	ReduceOnly              bool   `json:"reduceOnly"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	StopPrice               string `json:"stopPrice"`
	Symbol                  string `json:"symbol"`
	BookTime                int64  `json:"bookTime"`
	UpdateTime              int64  `json:"updateTime"`
	TimeInForce             string `json:"timeInForce"`
	ActivatePrice           string `json:"activatePrice"`
	PriceRate               string `json:"priceRate"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	GoodTillDate            int    `json:"goodTillDate"`
	PriceMatch              string `json:"priceMatch"`
}
