package kintohub

import (
	"encoding/json"
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
)

type ProxyRequestMessage struct {
	RequestId string            `json:"requestId"`
	Method    string            `json:"method"`
	Uri       string            `json:"uri"`
	Headers   map[string][]byte `json:"headers"`
	Data      []byte            `json:"data"`
}

type ProxyResponseMessage struct {
	RequestId string            `json:"requestId"`
	Headers   map[string][]byte `json:"headers"`
	Data      []byte            `json:"data"`
}

func createWebsocketProxy(gatewayHostname string, authToken string, blockName string, portToForward string) {
	proxyClient := fasthttp.HostClient{
		Addr: "localhost:" + portToForward,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	path := fmt.Sprintf("/proxy?authToken=%v&blockName=%v",
		authToken,
		blockName)

	u := url.URL{Scheme: "ws", Host: gatewayHostname, Path: path}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal("Failed to create proxy:", err)
	}

	defer c.Close()
	done := make(chan struct{})
	defer close(done)

	for {
		proxyMsg := ProxyRequestMessage{}
		err := c.ReadJSON(&proxyMsg)

		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				log.Print("Websocket has been closed!")
			} else {
				log.Println("Error occurred reading proxy msg", err)
			}
			return
		}

		prettyReqBytes, _ := json.MarshalIndent(proxyMsg, "", "\t")
		log.Printf(">> Proxying Request to localhost:%v: %v", portToForward, string(prettyReqBytes))
		response := proxyMessage(proxyClient, proxyMsg)
		prettyRespBytes, _ := json.MarshalIndent(response, "", "\t")
		log.Printf("<< Proxying Response to %v: %v", u.String(), string(prettyRespBytes))
		c.WriteJSON(response)
	}
}

func proxyMessage(client fasthttp.HostClient, msg ProxyRequestMessage) ProxyResponseMessage {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	req.Header.SetHost(client.Addr)
	req.Header.SetMethod(msg.Method)
	req.SetRequestURI(msg.Uri)
	msgBytes, err := json.Marshal(msg)
	req.SetBody(msgBytes)

	err = client.Do(req, resp)

	if err != nil {
		log.Fatal("Error proxying request to client", err)
	}

	response := ProxyResponseMessage{
		RequestId: msg.RequestId,
		Data:      resp.Body(),
		Headers:   map[string][]byte{},
	}

	resp.Header.VisitAll(func(key, value []byte) {
		response.Headers[string(key)] = value
	})

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(resp)

	return response
}
