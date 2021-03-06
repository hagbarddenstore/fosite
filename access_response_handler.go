package fosite

import (
	"github.com/go-errors/errors"
	"golang.org/x/net/context"
	"net/http"
)

func (f *Fosite) NewAccessResponse(ctx context.Context, req *http.Request, requester AccessRequester, session interface{}) (AccessResponder, error) {
	var err error
	var tk TokenEndpointHandler

	response := NewAccessResponse()
	for _, tk = range f.TokenEndpointHandlers {
		if err = tk.HandleTokenEndpointRequest(ctx, req, requester, response, session); err != nil {
			return nil, errors.Wrap(err, 1)
		}
	}

	if response.GetAccessToken() == "" || response.GetTokenType() == "" {
		return nil, ErrUnsupportedGrantType
	}

	return response, nil
}
