package kintohub

import (
	"encoding/json"
	"github.com/fasthttp/websocket"
	"github.com/magiconair/properties/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var upgrader = websocket.Upgrader{} // use default options

func TestWebsocket(t *testing.T) {

	testContainer := TestContainer{
		T: t,
	}

	gatewayServer := httptest.NewServer(http.HandlerFunc(testContainer.gatewayHandler))
	proxyServer := httptest.NewServer(http.HandlerFunc(testContainer.proxyHandler))

	defer gatewayServer.Close()
	defer proxyServer.Close()

	createWebsocketProxy(strings.TrimPrefix(gatewayServer.URL, "http://"),
		"abcd1234", "auth", strings.Split(proxyServer.URL, ":")[2])
}

type TestContainer struct {
	T *testing.T
}

func (c *TestContainer) gatewayHandler(w http.ResponseWriter, r *http.Request) {
	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer con.Close()

	//TODO: Check request URI

	reqMsg := ProxyRequestMessage{
		RequestId: "1",
		Method:    http.MethodPost,
		Data:      []byte("test-message"),
	}

	err = con.WriteJSON(reqMsg)

	if err != nil {
		c.T.Fatal("Error sending proxy message", err)
	}

	resp := ProxyResponseMessage{}
	err = con.ReadJSON(&resp)

	if err != nil {
		c.T.Fatal("Could not read response json", err)
	}

	assert.Equal(c.T, resp.Data, reqMsg.Data)
	assert.Equal(c.T, resp.RequestId, reqMsg.RequestId)
}

func (c *TestContainer) proxyHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		c.T.Fatal("Proxied Service could not read message", err)
	}

	msg := ProxyRequestMessage{}
	err = json.Unmarshal(b, &msg)

	if err != nil {
		c.T.Fatal("Proxied service cannot parse json message", err)
	}

	w.Write(msg.Data)
}
