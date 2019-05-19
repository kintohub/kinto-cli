package kintohub

type Api struct {
	authToken string
}

func InitApi(authToken string) Api {
	return Api{
		authToken: authToken,
	}
}

func (a *Api) CreateProxy(gatewayHost string, blockName string, forwardPort string) {
	createWebsocketProxy(gatewayHost, a.authToken, blockName, forwardPort)
}
