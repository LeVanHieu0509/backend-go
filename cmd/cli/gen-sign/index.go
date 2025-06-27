package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime"
	"net/http"
	"sort"
)

// CalSign calculates the signature based on the request and app secret
func CalSign(req *http.Request, secret string) string {
	queries := req.URL.Query()

	// Extract all query parameters excluding 'sign' and 'access_token'
	keys := make([]string, 0, len(queries))
	for k := range queries {
		// Skip 'sign' and 'access_token'
		if k != "sign" && k != "access_token" {
			keys = append(keys, k)
		}
	}

	// Reorder the parameters' keys alphabetically
	sort.Strings(keys)

	// Concatenate all the parameters in the format of {key}{value}
	var input string
	for _, key := range keys {
		input += key + queries.Get(key)
	}

	// Append the request path
	input = req.URL.Path + input

	// If the request header Content-Type is not multipart/form-data, append body to the end
	mediaType, _, _ := mime.ParseMediaType(req.Header.Get("Content-Type"))
	if mediaType != "multipart/form-data" && req.Body != nil {
		// Only read the body if it's not nil
		body, _ := io.ReadAll(req.Body)
		input += string(body)

		req.Body.Close()
		// Reset the body after reading it, so the request body can still be used later
		req.Body = io.NopCloser(bytes.NewReader(body))
	}

	// Wrap the string with the app secret at both ends
	input = secret + input + secret

	// Generate the HMAC SHA-256 signature
	return generateSHA256(input, secret)
}

// generateSHA256 generates the HMAC SHA-256 signature
func generateSHA256(input, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := h.Write([]byte(input)); err != nil {
		// Error logging can be handled here
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"your/tiktok-sdk-path/apis" // Import the TikTok API SDK package (adjust the import path)
)

func authorization202309GetAuthorizedShopsGet() {
	// Create configuration with your appKey and appSecret
	configuration := apis.NewConfiguration()
	configuration.AddAppInfo(appKey, appSecret)

	// Initialize the API client with the configuration
	apiClient := apis.NewAPIClient(configuration)

	// Prepare the request to fetch authorized shops
	request := apiClient.AuthorizationV202309API.Authorization202309ShopsGet(context.Background())

	// Set the access token and content type for the request
	request = request.XTtsAccessToken("your access token") // Replace with actual access token
	request = request.ContentType("application/json")

	// Execute the request and handle the response
	resp, httpResp, err := request.Execute()
	if err != nil {
		// Handle the error if the request failed
		log.Printf("Error executing request: %v", err)
		return
	}
	if httpResp.StatusCode != 200 {
		// Handle non-200 status code responses
		body := string(httpResp.Body)
		log.Printf("Request failed with status code %d and response body: %s", httpResp.StatusCode, body)
		return
	}

	if resp == nil {
		// Check if the response is nil
		log.Println("Received nil response")
		return
	}

	// Check for business logic errors (e.g., error codes from TikTok API)
	if resp.GetCode() != 0 {
		log.Printf("Error from TikTok API: ErrorCode: %d, ErrorMessage: %s", resp.GetCode(), resp.GetMessage())
		return
	}

	// If the response is successful, print the data
	respDataJson, err := json.MarshalIndent(resp.GetData(), "", "  ")
	if err != nil {
		log.Printf("Error marshalling response data: %v", err)
		return
	}

	// Print the formatted response data
	fmt.Println("Response data:", string(respDataJson))
}




func main() {
	// Example parameters for generating the signature
	req, _ := http.NewRequest("POST", "https://open-api.tiktokglobalshop.com/product/202502/products/search?app_key=6g1geqaov0amu&limit=20&timestamp=1745908667&access_token=ROW_ea9megAAAADbBdjhlnxOLBa53GSRWyq8CFFpawl5fZ-W0-xhFl3m1QoZVJplJmFf7LeAStnw1GquDljr0D_6edcerj4CbD1LTKIdZVTWyINPiHOMvh_70sFUcbm4SsN6fD4Za7UFyfSfktswFekdC6Wpj7g_HDEUUQdYd21MLHq9qufo7pJVng", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Generate the signature using the provided secret ("777")
	sign := CalSign(req, "777")

	// Output the generated signature
	fmt.Println("Generated Signature:", sign)

	authorization202309GetAuthorizedShopsGet()

}
