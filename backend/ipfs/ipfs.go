package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

// Standard V1/V2 Response (Most stable)
type PinataResponse struct {
	IpfsHash  string `json:"IpfsHash"`
	PinSize   int    `json:"PinSize"`
	Timestamp string `json:"Timestamp"`
	Error     string `json:"error,omitempty"`
}

func Upload(file io.Reader, filename string) (string, error) {
	jwt := os.Getenv("PINATA_JWT_TOKEN")
	if jwt == "" {
		return "", fmt.Errorf("PINATA_JWT_TOKEN is empty")
	}

	// Use the stable V1/V2 endpoint
	url := "https://api.pinata.cloud/pinning/pinFileToIPFS"
	log.Printf("🚀 Uploading %s to: %s", filename, url)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 1. Create File Part
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return "", fmt.Errorf("form create err: %v", err)
	}

	// Copy file into buffer
	if _, err := io.Copy(part, file); err != nil {
		return "", fmt.Errorf("file copy err: %v", err)
	}

	// 2. Add Pinata Metadata (Optional but good practice)
	// This helps you find the file in the Pinata UI
	metadata := fmt.Sprintf(`{"name": "%s"}`, filename)
	_ = writer.WriteField("pinataMetadata", metadata)

	// 3. Add Pinata Options (Optional)
	_ = writer.WriteField("pinataOptions", `{"cidVersion": 1}`)

	writer.Close()

	// 4. Create Request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", fmt.Errorf("req create err: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+jwt)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// CRITICAL FIX: Disable Keep-Alives to prevent "unexpected EOF" on reused connections
	req.Close = true

	// 5. Execute
	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		// Log the JWT length to debug if it's being cut off (Don't log the full key!)
		log.Printf("❌ Network Error. JWT Length: %d", len(jwt))
		return "", fmt.Errorf("network err: %v", err)
	}
	defer resp.Body.Close()

	// 6. Parse Response
	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		log.Printf("🔴 Pinata Error Body: %s", string(respBody))
		return "", fmt.Errorf("pinata API status %d", resp.StatusCode)
	}

	var pinataResp PinataResponse
	if err := json.Unmarshal(respBody, &pinataResp); err != nil {
		return "", fmt.Errorf("json parse err: %v", err)
	}

	log.Printf("✅ Upload Success! CID: %s", pinataResp.IpfsHash)
	return pinataResp.IpfsHash, nil
}
