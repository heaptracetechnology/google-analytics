package analytics

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	result "github.com/heaptracetechnology/google-analytics/result"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/analytics/v3"
	"net/http"
	"os"
)

//Arguments struct
type Arguments struct {
	AccountID     string `json:"accountId,omitempty"`
	WebPropertyID string `json:"webPropertyId,omitempty"`
	ProfileID     string `json:"profileId,omitempty"`
}

//AccountList Google-Analytics
func AccountList(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decodedJSON, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	conf, confErr := google.JWTConfigFromJSON(decodedJSON, analytics.AnalyticsReadonlyScope)
	if confErr != nil {
		result.WriteErrorResponseString(responseWriter, confErr.Error())
		return
	}

	httpClient := conf.Client(oauth2.NoContext)
	service, serviceErr := analytics.New(httpClient)
	if serviceErr != nil {
		result.WriteErrorResponseString(responseWriter, serviceErr.Error())
		return
	}

	accountResponse, accountErr := service.Management.Accounts.List().Do()
	if accountErr != nil {
		result.WriteErrorResponseString(responseWriter, accountErr.Error())
		return
	}

	bytes, _ := json.Marshal(accountResponse)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//WebPropertiesList Google-Analytics
func WebPropertiesList(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decodedJSON, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	decoder := json.NewDecoder(request.Body)
	var param Arguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	conf, confErr := google.JWTConfigFromJSON(decodedJSON, analytics.AnalyticsReadonlyScope)
	if confErr != nil {
		result.WriteErrorResponseString(responseWriter, confErr.Error())
		return
	}

	httpClient := conf.Client(oauth2.NoContext)

	service, serviceErr := analytics.New(httpClient)
	if serviceErr != nil {
		result.WriteErrorResponseString(responseWriter, serviceErr.Error())
		return
	}

	webProps, webPropsErr := service.Management.Webproperties.List(param.AccountID).Do()
	if webPropsErr != nil {
		result.WriteErrorResponseString(responseWriter, webPropsErr.Error())
		return
	}

	bytes, _ := json.Marshal(webProps)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//ProfilesList Google-Analytics
func ProfilesList(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decodedJSON, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	decoder := json.NewDecoder(request.Body)
	var param Arguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	conf, confErr := google.JWTConfigFromJSON(decodedJSON, analytics.AnalyticsReadonlyScope)
	if confErr != nil {
		result.WriteErrorResponseString(responseWriter, confErr.Error())
		return
	}

	httpClient := conf.Client(oauth2.NoContext)

	service, serviceErr := analytics.New(httpClient)
	if serviceErr != nil {
		result.WriteErrorResponseString(responseWriter, serviceErr.Error())
		return
	}

	profiles, profilesErr := service.Management.Profiles.List(param.AccountID, param.WebPropertyID).Do()
	if profilesErr != nil {
		result.WriteErrorResponseString(responseWriter, profilesErr.Error())
		return
	}

	bytes, _ := json.Marshal(profiles)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//RealtimeData Google-Analytics
func RealtimeData(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decodedJSON, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	decoder := json.NewDecoder(request.Body)
	var param Arguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	conf, confErr := google.JWTConfigFromJSON(decodedJSON, analytics.AnalyticsReadonlyScope)
	if confErr != nil {
		result.WriteErrorResponseString(responseWriter, confErr.Error())
		return
	}

	httpClient := conf.Client(oauth2.NoContext)

	service, serviceErr := analytics.New(httpClient)
	if serviceErr != nil {
		result.WriteErrorResponseString(responseWriter, serviceErr.Error())
		return
	}

	profileID := "ga:" + param.ProfileID
	fmt.Println("profileID ::", profileID)
	realtime, realtimeErr := service.Data.Realtime.Get(profileID, "rt:activeUsers").Do()
	if realtimeErr != nil {
		result.WriteErrorResponseString(responseWriter, realtimeErr.Error())
		return
	}

	bytes, _ := json.Marshal(realtime)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
