package config

import "os"

func GetDeviceVerificationCodeUri() string {
	return "https://github.com/login/device/code"
}

func GetAccessCodeUri() string {
	return "https://github.com/login/oauth/access_token"
}

func GetClientID() string {
	clienteID, _ := os.LookupEnv("GITHUB_CLIENT_ID")
	return clienteID
}

func GetScopes() []string {
	return []string{"repo", "user"}
}

func GetGrantType() string {
	return "urn:ietf:params:oauth:grant-type:device_code"
}
