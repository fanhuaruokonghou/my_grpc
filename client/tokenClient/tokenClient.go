package tokenClient

import "context"

type TokenAuthentication struct {
	AppKey    string
	AppSecret string
}

func (i TokenAuthentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appKey":    i.AppKey,
		"appSecret": i.AppSecret,
	}, nil
}

func (i TokenAuthentication) RequireTransportSecurity() bool {
	return true
}
