package bitget

import (
	"net/http"
	"strings"

	"github.com/outtoin/bitget-golang-sdk-api/internal/common"
	"github.com/outtoin/bitget-golang-sdk-api/internal/model"
	"github.com/outtoin/bitget-golang-sdk-api/pkg/client/broker"
	"github.com/outtoin/bitget-golang-sdk-api/pkg/client/mix"
	"github.com/outtoin/bitget-golang-sdk-api/pkg/client/spot"
	"github.com/outtoin/bitget-golang-sdk-api/pkg/client/ws"
)

// client
type Client struct {
	client *common.BitgetRestClient

	brokerService *broker.BrokerAccountClient

	mixAccountService  *mix.MixAccountClient
	mixMarketService   *mix.MixMarketClient
	mixOrderService    *mix.MixOrderClient
	mixPlanService     *mix.MixPlanClient
	mixPositionService *mix.MixPositionClient
	mixTraceService    *mix.MixTraceClient

	spotAccountService *spot.SpotAccountClient
	spotMarketService  *spot.SpotMarketClient
	spotOrderService   *spot.SpotOrderClient
	spotPublicService  *spot.SpotPublicClient
	//spotWalletService	*spot.SpotWalletClient			// @todo
}

func NewClient() *Client {
	bc := new(common.BitgetRestClient).Init()

	return &Client{
		client: bc,

		brokerService:      &broker.BrokerAccountClient{bc},
		mixAccountService:  &mix.MixAccountClient{bc},
		mixMarketService:   &mix.MixMarketClient{bc},
		mixOrderService:    &mix.MixOrderClient{bc},
		mixPlanService:     &mix.MixPlanClient{bc},
		mixPositionService: &mix.MixPositionClient{bc},
		mixTraceService:    &mix.MixTraceClient{bc},
		spotAccountService: &spot.SpotAccountClient{bc},
		spotMarketService:  &spot.SpotMarketClient{bc},
		spotOrderService:   &spot.SpotOrderClient{bc},
		spotPublicService:  &spot.SpotPublicClient{bc},
	}
}

func (c *Client) SetHttpClient(client *http.Client) *Client {
	c.client.HttpClient = client
	return c
}

// broker
func (c *Client) GetBrokerService() *broker.BrokerAccountClient {
	return c.brokerService
}

// mix
func (c *Client) GetMixAccountService() *mix.MixAccountClient {
	return c.mixAccountService
}
func (c *Client) GetMixMarketService() *mix.MixMarketClient {
	return c.mixMarketService
}
func (c *Client) GetMixOrderService() *mix.MixOrderClient {
	return c.mixOrderService
}
func (c *Client) GetMixPlanService() *mix.MixPlanClient {
	return c.mixPlanService
}
func (c *Client) GetMixPositionService() *mix.MixPositionClient {
	return c.mixPositionService
}
func (c *Client) GetMixTraceService() *mix.MixTraceClient {
	return c.mixTraceService
}

// spot
func (c *Client) GetSpotAccountService() *spot.SpotAccountClient {
	return c.spotAccountService
}
func (c *Client) GetSpotMarketService() *spot.SpotMarketClient {
	return c.spotMarketService
}
func (c *Client) GetSpotOrderService() *spot.SpotOrderClient {
	return c.spotOrderService
}
func (c *Client) GetSpotPublicService() *spot.SpotPublicClient {
	return c.spotPublicService
}

/* @todo
func (c *Client) NewSpotWalletService() *spot.SpotWalletClient {
	return c.spotWalletService
}
//*/

// ws
type WsClient struct {
	bws *ws.BitgetWsClient
}

func NewWsClient() *WsClient {
	return &WsClient{
		new(ws.BitgetWsClient),
	}
}

func (ws *WsClient) Init(listener common.OnReceive, errorListener common.OnReceive, secure bool) *ws.BitgetWsClient {
	return ws.bws.Init(secure, listener, errorListener)
}

func (ws *WsClient) Close() {
	ws.bws.Close()
}

type UnscribeFunc func()

// spot
func (ws *WsClient) SubscribeSpot(channel string, symbols ...string) UnscribeFunc {
	subs := make([]model.SubscribeReq, len(symbols))
	for i, s := range symbols {
		subs[i] = model.SubscribeReq{
			InstType: "SP",
			Channel:  channel,
			InstId:   s,
		}
	}
	ws.bws.SubscribeDef(subs)

	return func() { ws.bws.UnSubscribe(subs) }
}

func (ws *WsClient) SubscribeSpotAccount() UnscribeFunc {
	sub := []model.SubscribeReq{
		{
			InstType: "spbl",
			Channel:  "account",
			InstId:   "default",
		},
	}
	ws.bws.SubscribeDef(sub)

	return func() { ws.bws.UnSubscribe(sub) }
}
func (ws *WsClient) SubscribeSpotOrder(symbols ...string) UnscribeFunc {
	subs := make([]model.SubscribeReq, len(symbols))
	for i, s := range symbols {
		subs[i] = model.SubscribeReq{
			InstType: "spbl",
			Channel:  "orders",
			InstId:   strings.ToUpper(s) + "_SPBL",
		}
	}
	ws.bws.SubscribeDef(subs)

	return func() { ws.bws.UnSubscribe(subs) }
}

// futures
func (ws *WsClient) SubscribeFutures(channel string, symbols ...string) UnscribeFunc {
	subs := make([]model.SubscribeReq, len(symbols))
	for i, s := range symbols {
		subs[i] = model.SubscribeReq{
			InstType: "MC",
			Channel:  channel,
			InstId:   s,
		}
	}
	ws.bws.SubscribeDef(subs)

	return func() { ws.bws.UnSubscribe(subs) }
}

func (ws *WsClient) SubscribeForContracts(channel string, contracts ...string) UnscribeFunc {
	subs := make([]model.SubscribeReq, len(contracts))
	for i, cid := range contracts {
		subs[i] = model.SubscribeReq{
			InstType: cid,
			Channel:  channel,
			InstId:   "default",
		}
	}
	ws.bws.SubscribeDef(subs)

	return func() { ws.bws.UnSubscribe(subs) }
}
