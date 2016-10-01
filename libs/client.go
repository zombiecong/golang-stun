package libs

type Client struct {
	UID string
	Session string
	IPAddr string
	ReturnPort int
	Info map[string]interface{}
}

func MakeClient() *Client {
	return new(Client)
	
}
