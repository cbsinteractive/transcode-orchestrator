package common

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/log"
	"github.com/bitmovin/bitmovin-api-sdk-go/rest"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/serialization"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"errors"
)

type ApiClient struct {
	restClient  *rest.RestClient
	Logger      log.Logger
	BaseUrl     string
	ApiKey      string
	TenantOrgId string
}

type PathParams map[string]interface{}

const QueryParamTagName = "query"
const DefaultApiBaseUrl = "https://api.bitmovin.com/v1"
const ContentTypeJson = "application/json"
const ApiClientVersion = "1.30.0-alpha.0"
const ApiClientName = "bitmovin-api-sdk-go"
const NoApiKeyErrorMsg = "there was no api key provided"

type QueryParams interface {
	Params() map[string]string
}

type RequestParams struct {
	PathParams  PathParams
	QueryParams QueryParams
}

func NewApiClient(configs ...func(configuration *ApiClient)) (*ApiClient, error) {
	apiClient := ApiClient{
		BaseUrl: DefaultApiBaseUrl,
		Logger:  &log.StdOutLogger{},
	}

	for _, config := range configs {
		config(&apiClient)
	}

	if apiClient.ApiKey == "" {
		return nil, fmt.Errorf(NoApiKeyErrorMsg)
	}

	apiClient.restClient = rest.NewRestClient(&rest.Configuration{})
	return &apiClient, nil
}

func GetParamsMap(u interface{}) map[string]string {
	val := reflect.ValueOf(u)
	if val.Kind() != reflect.Ptr {
		panic("Parameter of GetParamsMap has to be a pointer")
	}

	val = reflect.ValueOf(u).Elem()
	t := val.Type()

	queryParamMap := make(map[string]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		typeField := t.Field(i)
		tag := typeField.Tag.Get(QueryParamTagName)

		valueField := val.Field(i)

		switch valueField.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value := valueField.Addr().Elem().Int()
			if value > 0 {
				queryParamMap[tag] = strconv.FormatInt(value, 10)
			}
		case reflect.String:
			value := valueField.Addr().Elem().String()
			if len(value) > 0 {
				queryParamMap[tag] = value
			}
			//ToDo check more cases e.g. date/time and format it correctly if that is needed at the moment?
		}

	}
	return queryParamMap
}

func (apiClient *ApiClient) prepareUrl(relUrl string, reqParams ...func(reqParams *RequestParams)) (string, error) {
	u, err := url.Parse(apiClient.BaseUrl)
	if err != nil {
		panic(err)
	}

	requestParams := &RequestParams{
		PathParams: make(map[string]interface{}),
	}
	for _, reqParam := range reqParams {
		reqParam(requestParams)
	}

	regex := regexp.MustCompile(`{(.*?)}`)
	found := regex.FindAllStringSubmatch(relUrl, -1)

	for k, v := range requestParams.PathParams {
		stringVal := ""
		if timeParam, isTime := v.(time.Time); isTime {
			stringVal = timeParam.Format(time.RFC3339)
		} else if stringParam, isString := v.(string); isString {
			stringVal = stringParam
		} else {
			return "", errors.New(fmt.Sprintf("path param can not be of type %s", reflect.TypeOf(v).String()))
		}

		for _, f := range found {
			if k == f[1] {
				relUrl = strings.Replace(relUrl, "{"+f[1]+"}", stringVal, -1)
			}
		}
	}

	found = regex.FindAllStringSubmatch(relUrl, -1)
	if len(found) > 0 {
		return "", fmt.Errorf("there are still some placeholders left in the URL. "+
			"Please make sure to pass the correct values to this function to replace all placeholders. url=%s", relUrl)
	}

	if requestParams.QueryParams != nil {
		queryParamMap := requestParams.QueryParams.Params()
		var queryParams []string
		queryParamString := ""
		for k, v := range queryParamMap {
			queryParam := fmt.Sprintf("%s=%s", k, v)
			queryParams = append(queryParams, queryParam)
		}
		queryParamString = strings.Join(queryParams, "&")
		if len(queryParamString) > 0 {
			u.RawQuery = queryParamString
		}
	}

	u.Path = path.Join(u.Path, relUrl)
	return u.String(), nil
}

func (apiClient *ApiClient) createHeaders() http.Header {
	headers := http.Header{}
	headers.Add("X-Api-Key", apiClient.ApiKey)
	if apiClient.TenantOrgId != "" {
		headers.Add(
			"X-Tenant-Org-Id", apiClient.TenantOrgId,
		)
	}
	headers.Add("X-Api-Client", ApiClientName)
	headers.Add("X-Api-Client-Version", ApiClientVersion)
	headers.Add("Content-Type", ContentTypeJson)

	return headers
}

func (apiClient *ApiClient) Get(relUrl string, requestModel interface{}, responseModel interface{}, requestParams ...func(params *RequestParams)) error {
	return apiClient.request(http.MethodGet, relUrl, nil, responseModel, requestParams...)
}

func (apiClient *ApiClient) Patch(relUrl string, requestModel interface{}, responseModel interface{}, requestParams ...func(params *RequestParams)) error {
	return apiClient.request(http.MethodPatch, relUrl, requestModel, responseModel, requestParams...)
}

func (apiClient *ApiClient) Post(relUrl string, requestModel interface{}, responseModel interface{}, requestParams ...func(params *RequestParams)) error {
	return apiClient.request(http.MethodPost, relUrl, requestModel, responseModel, requestParams...)
}

func (apiClient *ApiClient) Put(relUrl string, requestModel interface{}, responseModel interface{}, requestParams ...func(params *RequestParams)) error {
	return apiClient.request(http.MethodPut, relUrl, requestModel, responseModel, requestParams...)
}

func (apiClient *ApiClient) Delete(relUrl string, requestModel interface{}, responseModel interface{}, requestParams ...func(params *RequestParams)) error {
	return apiClient.request(http.MethodDelete, relUrl, nil, responseModel, requestParams...)
}

func (apiClient *ApiClient) logRequest(reqUrl string, method string, headers http.Header, rawBody []byte) {
	apiClient.Logger.LogRequest(log.RequestLogInfo{
		URL:     reqUrl,
		Method:  method,
		Headers: headers,
		RawBody: rawBody,
	})
}

func (apiClient *ApiClient) logResponse(reqUrl string, method string, headers http.Header, statusCode int, rawBody []byte) {
	apiClient.Logger.LogResponse(log.ResponseLogInfo{
		URL:        reqUrl,
		Method:     method,
		Headers:    headers,
		RawBody:    rawBody,
		StatusCode: statusCode,
	})
}

func createBitmovinError(envelope model.GenericResponseEnvelope) BitmovinError {
	var bitErr BitmovinError
	json.Unmarshal(envelope.Data, &bitErr)
	return bitErr
}

func (apiClient *ApiClient) request(reqMethod string, relUrl string, requestModel interface{}, responseModel interface{}, requestParams ...func(params *RequestParams)) error {

	reqUrl, err := apiClient.prepareUrl(relUrl, requestParams...)
	if err != nil {
		return err
	}

	headers := apiClient.createHeaders()
	if err != nil {
		return err
	}

	var resp *http.Response
	switch reqMethod {
	case http.MethodPatch:
		serialized := serialization.Serialize(requestModel)
		apiClient.logRequest(reqUrl, reqMethod, headers, serialized)
		resp, err = apiClient.restClient.Patch(reqUrl, serialized, headers)
	case http.MethodPost:
		serialized := serialization.Serialize(requestModel)
		apiClient.logRequest(reqUrl, reqMethod, headers, serialized)
		resp, err = apiClient.restClient.Post(reqUrl, serialized, headers)
	case http.MethodPut:
		serialized := serialization.Serialize(requestModel)
		apiClient.logRequest(reqUrl, reqMethod, headers, serialized)
		resp, err = apiClient.restClient.Put(reqUrl, serialized, headers)
	case http.MethodGet:
		apiClient.logRequest(reqUrl, reqMethod, headers, nil)
		resp, err = apiClient.restClient.Get(reqUrl, headers)
	case http.MethodDelete:
		apiClient.logRequest(reqUrl, reqMethod, headers, nil)
		resp, err = apiClient.restClient.Delete(reqUrl, headers)
	}

	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	apiClient.logResponse(reqUrl, reqMethod, resp.Header, resp.StatusCode, respBody)

	var envelope model.GenericResponseEnvelope
	err = json.Unmarshal(respBody, &envelope)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var bitmovinError BitmovinError
		if err == nil {
			bitmovinError = createBitmovinError(envelope)
		}
		bitmovinError.RawResponse = string(respBody)
		bitmovinError.HttpStatusCode = &resp.StatusCode
		return bitmovinError

	} else if err != nil {
		return err
	}

	if responseModel != nil {
		convertToModel(envelope, &responseModel)
	}
	return nil
}

func convertToModel(envelope model.GenericResponseEnvelope, out interface{}) {
	var successData model.ResultWrapper
	json.Unmarshal(envelope.Data, &successData)
	json.Unmarshal(successData.Result, out)
}

func headerContentTypeIsJson(respHeaders http.Header) bool {
	return respHeaders.Get("Content-Type") != ContentTypeJson
}
