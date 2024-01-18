package assetlayer

type App struct {
	AppID         string  `json:"appId"`
	HandcashAppID string  `json:"handcashAppId"`
	AppName       string  `json:"appName"`
	AppImage      string  `json:"appImage"`
	AppBanner     string  `json:"appBanner"`
	Description   string  `json:"description"`
	URL           string  `json:"url"`
	AutoGrantRead bool    `json:"autoGrantRead"`
	TeamID        string  `json:"teamId"`
	Status        string  `json:"status"`
	CreatedAt     int64   `json:"createdAt"`
	UpdatedAt     int64   `json:"updatedAt"`
	Slots         []*Slot `json:"slots"`
}

func (client *Client) NewAppWallet(handle string) (string, error) {
	data, err := client.Try(
		"POST",
		"/api/v1/app/newAppWallet",
		nil,
		map[string]interface{}{
			"appHandle": handle,
		},
	)
	if err != nil {
		return "", err
	}

	m, err := assertMapStringInterface(data)
	if err != nil {
		return "", err
	}
	s, err := assertString(m["userId"])
	if err != nil {
		return "", err
	}

	return s, nil
}
