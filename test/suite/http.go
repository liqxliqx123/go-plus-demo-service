package suite

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zgs225/httpassert"
)

// HTTPResponseAssert assert http body
type HTTPResponseAssert struct {
	t        *testing.T
	response *http.Response
}

// ExpectCode assert http code
func (a *HTTPResponseAssert) ExpectCode(code int) *HTTPResponseAssert {
	assert.Equal(a.t, code, a.response.StatusCode, "unexpected http response code")
	return a
}

// ExpectHeader assert response has header key, and its value equals vals
func (a *HTTPResponseAssert) ExpectHeader(key string, vals ...string) *HTTPResponseAssert {
	assert.Equalf(a.t, vals, a.response.Header.Values(key), "unexpected header %s", textproto.CanonicalMIMEHeaderKey(key))
	return a
}

// ExpectString assert response body equals string s
func (a *HTTPResponseAssert) ExpectString(s string) *HTTPResponseAssert {
	b, _ := ioutil.ReadAll(a.response.Body)
	assert.Equal(a.t, s, string(b), "unexpected response body")
	return a
}

// ExpectJSON expect http body
func (a *HTTPResponseAssert) ExpectJSON(i interface{}) *HTTPResponseAssert {
	httpassert.EqualJSON(a.t, i, a.response.Body)
	return a
}

// ExpectThat run the custom assert function
func (a *HTTPResponseAssert) ExpectThat(fn func(*testing.T, *http.Response)) *HTTPResponseAssert {
	fn(a.t, a.response)
	return a
}

// Get do get request and assert response
func (suite *APISuite) Get(path string) *HTTPResponseAssert {
	return suite.DoRequest("GET", path, nil, nil)
}

// PostJSON do post request with json data and assert response
func (suite *APISuite) PostJSON(path string, data interface{}) *HTTPResponseAssert {
	b := new(bytes.Buffer)

	err := json.NewEncoder(b).Encode(data)

	suite.NoError(err)

	return suite.DoRequest("POST", path, nil, b)
}

// DoRequest do http request and assert response
func (suite *APISuite) DoRequest(method, path string, headers map[string]string, body io.Reader) *HTTPResponseAssert {
	req, err := http.NewRequest(method, suite.buildURL(path), body)

	suite.NoError(err)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := http.DefaultClient.Do(req)

	suite.NoError(err)

	return &HTTPResponseAssert{
		t:        suite.T(),
		response: res,
	}
}

func (suite *APISuite) buildURL(path string) string {
	return fmt.Sprintf("http://%s%s", suite.Listener.Addr().String(), path)
}
