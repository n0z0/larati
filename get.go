package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
)

// cookie: laravel_session=dasdsa0o
func RunGET(url string, cookie string) {
	// Membuat permintaan HTTP GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Menyertakan header yang dibutuhkan
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,id;q=0.8")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Cookie", cookie) // Memasukkan cookie yang diberikan
	req.Header.Set("Lang", "null")
	req.Header.Set("Host", "dev.pamongdesa.id")
	req.Header.Set("Origin", "https://dev.pamongdesa.id")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Sec-Ch-Ua", `"Not/A)Brand";v="8", "Chromium";v="126", "Google Chrome";v="126"`)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", `"Windows"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	// Mengirimkan permintaan HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Cek apakah respons menggunakan GZIP
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Println("Error creating GZIP reader:", err)
			return
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	// Membaca respons dari server
	body, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Menampilkan status dan respons
	fmt.Println("Response Status:", resp.Status)
	//fmt.Println("Response Body:", string(body))
	// Menulis respons ke file
	file, err := os.Create("response_body.html")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
