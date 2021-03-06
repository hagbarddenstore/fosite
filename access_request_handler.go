package fosite

import (
	"github.com/go-errors/errors"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

// Implements
// * https://tools.ietf.org/html/rfc6749#section-2.3.1
//   Clients in possession of a client password MAY use the HTTP Basic
//   authentication scheme as defined in [RFC2617] to authenticate with
//   the authorization server.  The client identifier is encoded using the
//   "application/x-www-form-urlencoded" encoding algorithm per
//   Appendix B, and the encoded value is used as the username; the client
//   password is encoded using the same algorithm and used as the
//   password.  The authorization server MUST support the HTTP Basic
//   authentication scheme for authenticating clients that were issued a
//   client password.
//   Including the client credentials in the request-body using the two
//   parameters is NOT RECOMMENDED and SHOULD be limited to clients unable
//   to directly utilize the HTTP Basic authentication scheme (or other
//   password-based HTTP authentication schemes).  The parameters can only
//   be transmitted in the request-body and MUST NOT be included in the
//   request URI.
//   * https://tools.ietf.org/html/rfc6749#section-3.2.1
//   - Confidential clients or other clients issued client credentials MUST
//   authenticate with the authorization server as described in
//   Section 2.3 when making requests to the token endpoint.
//   - If the client type is confidential or the client was issued client
//   credentials (or assigned other authentication requirements), the
//   client MUST authenticate with the authorization server as described
//   in Section 3.2.1.
func (f *Fosite) NewAccessRequest(ctx context.Context, r *http.Request, session interface{}) (AccessRequester, error) {
	ar := NewAccessRequest()
	if r.Method != "POST" {
		return ar, errors.New(ErrInvalidRequest)
	}

	if f.RequiredScope == "" {
		f.RequiredScope = DefaultRequiredScopeName
	}

	if err := r.ParseForm(); err != nil {
		return ar, errors.New(ErrInvalidRequest)
	}

	if session == nil {
		return ar, errors.New("Session must not be nil")
	}

	ar.Scopes = removeEmpty(strings.Split(r.Form.Get("scope"), " "))
	ar.GrantType = r.Form.Get("grant_type")
	if ar.GrantType == "" {
		return ar, errors.New(ErrInvalidRequest)
	}

	clientID, clientSecret, ok := r.BasicAuth()
	if !ok {
		return ar, errors.New(ErrInvalidRequest)
	}

	client, err := f.Store.GetClient(clientID)
	if err != nil {
		return ar, errors.New(ErrInvalidClient)
	}

	// Enforce client authentication
	if err := f.Hasher.Compare(client.GetHashedSecret(), []byte(clientSecret)); err != nil {
		return ar, errors.New(ErrInvalidClient)
	}
	ar.Client = client

	for _, loader := range f.TokenEndpointHandlers {
		if err := loader.ValidateTokenEndpointRequest(ctx, r, ar, session); err != nil {
			return ar, err
		}
	}

	if !ar.DidHandleGrantType() {
		return ar, errors.New(ErrUnsupportedGrantType)
	}

	if !ar.GetScopes().Has(f.RequiredScope) {
		return ar, errors.New(ErrInvalidScope)
	}

	ar.GrantScope(f.RequiredScope)
	return ar, nil
}
