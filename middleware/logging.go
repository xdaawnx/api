package middleware

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/labstack/echo/v4"
)

// Define a struct for the content
type LogContent struct {
	ID             snowflake.ID `json:"id"`
	Time           string       `json:"time"`
	RequestURI     string       `json:"request_uri"`
	RequestHeaders string       `json:"request_headers"`
	QueryParams    string       `json:"query_params"`
	RequestBody    string       `json:"request_body"`
	ResponseBody   string       `json:"response_body"`
}

// Custom function for logging request and response data
func LogRequestResponse(c echo.Context, reqBody, resBody []byte) {
	// Log the headers
	headers, _ := json.Marshal(c.Request().Header)
	node, err := snowflake.NewNode(1)
	if err != nil {
		return
	}
	id := node.Generate()

	// Log the query parameters
	query, _ := json.Marshal(c.QueryParams())
	uri := c.Request().RequestURI
	content := LogContent{
		ID:             id,
		Time:           time.Now().Format(time.RFC3339Nano),
		RequestURI:     uri,
		RequestHeaders: string(headers),
		QueryParams:    string(query),
		RequestBody:    string(reqBody),
		ResponseBody:   string(resBody),
	}
	// Log request and response body, headers, and query params
	k, _ := json.Marshal(content)
	log.Println(string(k))
	return
}

// Custom skipper function to skip certain routes
func LogSkipper(c echo.Context) bool {
	// Skip logging for the "/health" route
	if strings.Contains(c.Path(), "/swagger") {
		return true
	}
	// Log all other routes
	return false
}
