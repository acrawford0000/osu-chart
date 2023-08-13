package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"project/backend/api/client/opts"
	"project/backend/api/enum"
	"project/backend/api/model"
	"reflect"
	"strconv"
	"strings"
)

// OsuClient is the client used to access all the communication with the osu! API.
// Requires osu! client ID and client secret to work.
type OsuClient struct {
	httpClient      *http.Client
	userCredentials model.UserCredentials
	loginExpiration int64
	token           model.Token
}

// NewClient creates new OsuClient.
// Client ID and client secret are needed to access the API.
func NewClient(clientID int, clientSecret string) (*OsuClient, error) {
	client := OsuClient{
		httpClient: http.DefaultClient,
		userCredentials: model.UserCredentials{
			ClientID:     model.SensitiveInt(clientID),
			ClientSecret: model.SensitiveString(clientSecret),
			GrantType:    enum.GrantTypeClientCredentials,
			Scope:        enum.ScopePublic,
		},
	}

	err := client.login()
	if err != nil {
		return nil, err
	}

	return &client, nil
}

// doReq makes a request to the specified osu! API endpoint with a specified method and opts.
func (c *OsuClient) doReq(method string, endpoint enum.Endpoint, opts opts.Opts) ([]byte, error) {
	if err := c.loginIfTokenExpired(); err != nil {
		return nil, err
	}

	req, err := buildRequest(endpoint, &c.token.AccessToken, method, opts)
	if err != nil {
		return nil, fmt.Errorf("unable to create a GET request to %s: %v", endpoint, err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	} else if res.StatusCode < 200 || res.StatusCode > 299 {
		log.Println(res.Request.URL.String())
		return nil, fmt.Errorf("response status code outside of 200-299 boundaries: %v", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %v", err)
	}

	return body, nil
}

// getReq is a GET request wrapper to the doReq function.
func (c *OsuClient) getReq(endpoint enum.Endpoint, opts opts.Opts) ([]byte, error) {
	return c.doReq("GET", endpoint, opts)
}

// postReq is a POST request wrapper to the doReq function.
func (c *OsuClient) postReq(endpoint enum.Endpoint, opts opts.Opts) ([]byte, error) {
	return c.doReq("POST", endpoint, opts)
}

// handleResponse handles the response of the OsuClient.doReq method.
// Unmarshalls the response body to the specified struct.
func handleResponse[T any](body []byte, err error) (*T, error) {
	if err != nil {
		return nil, err
	}

	var obj T
	err = json.Unmarshal(body, &obj)
	if err != nil {
		log.Println(string(body))
		return nil, err
	}

	return &obj, nil
}

// addQueryParams adds query parameters to the URL values.
func addQueryParams(req *http.Request, params map[string]interface{}) {
	query := req.URL.Query()
	for k, v := range params {
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Slice {
			for i := 0; i < val.Len(); i++ {
				query.Add(k, fmt.Sprintf("%v", val.Index(i).Interface()))
			}
		} else {
			query.Add(k, fmt.Sprintf("%v", v))
		}
	}

	req.URL.RawQuery = query.Encode()
}

// addHeaders adds required headers to the request.
func addHeaders(req *http.Request, accessToken *model.SensitiveString) {
	req.Header.Add("Authorization", "Bearer "+accessToken.GetActualValue())
	req.Header.Add("Content-Type", "application/json")
}

// buildRequest builds a request with the supplied opts that implement the opts.Opts interface.
func buildRequest(endpoint enum.Endpoint, accessToken *model.SensitiveString, method string, opts opts.Opts) (*http.Request, error) {
	var reqURL string
	pathParams := make(map[int]interface{})
	queryParams := make(map[string]interface{})
	bodyParams := make(map[string]interface{})

	if opts != nil {
		if err := handleOpts(pathParams, queryParams, bodyParams, opts); err != nil {
			return nil, err
		}
	}

	sortedPathParams := sortPathParams(pathParams)
	reqURL = fmt.Sprintf(endpoint.String(), sortedPathParams...)
	marshalledBody, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, enum.BaseEndpointOsu.String()+reqURL, bytes.NewBuffer(marshalledBody))
	if err != nil {
		return nil, err
	}

	addHeaders(req, accessToken)
	addQueryParams(req, queryParams)

	return req, nil
}

// sortPathParams sorts path parameters by key. A slice of path parameters is returned.
func sortPathParams(pathParams map[int]interface{}) (sorted []interface{}) {
	for i := 0; i < len(pathParams); i++ {
		sorted = append(sorted, pathParams[i])
	}

	return
}

// handleOpts handles opts provided by the user.
// reflect package is used to access pre-defined struct tags and parse them. Then, the value is routed into the correct
// map based on previously mentioned struct tags.
func handleOpts(pathParams map[int]interface{}, queryParams map[string]interface{}, bodyParams map[string]interface{}, opt opts.Opts) error {
	if err := opt.Valid(); err != nil {
		return err
	}

	reflectValue := reflect.ValueOf(opt)
	reflectType := reflect.TypeOf(opt)

	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		param := field.Tag.Get("param")
		index, err := strconv.ParseInt(field.Tag.Get("index"), 10, 0)
		if err != nil {
			// Index is not set on the field.
			index, err = 0, nil
		}

		var value interface{}
		paramSplit := strings.Split(param, ",")
		vField := reflectValue.Field(i)
		if !vField.IsZero() {
			if vField.Kind() == reflect.Ptr && !vField.IsNil() {
				value = vField.Elem().Interface()
			} else {
				value = vField.Interface()
			}
		} else if _, exists := pathParams[int(index)]; paramSplit[0] == "path" && !exists {
			value = ""
		}

		if value != nil {
			switch paramSplit[0] {
			case "query":
				switch v := value.(type) {
				case opts.Cursor:
					for x, y := range v {
						queryParams[fmt.Sprintf("cursor[%v]", x)] = y
					}
					continue
				default:
					queryParams[paramSplit[1]] = value
					continue
				}
			case "body":
				bodyParams[paramSplit[1]] = value
				continue
			case "path":
				pathParams[int(index)] = value
				continue
			default:
				return fmt.Errorf("unable to correctly parse struct tags")
			}
		}
	}

	return nil
}
