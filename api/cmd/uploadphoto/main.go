package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"golang.org/x/term"
)

const (
	colorReset  = "\033[0m"
	colorBold   = "\033[1m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Printf("%s%s Upload Photo (test) %s\n", colorBold, colorCyan, colorReset)
	fmt.Println(strings.Repeat("─", 40))
	fmt.Println()

	baseURL := prompt(reader, "API base URL [http://localhost:8080]")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	baseURL = strings.TrimRight(baseURL, "/")

	email := prompt(reader, "Email")
	if email == "" {
		fatal("Email is required.")
	}
	password := promptPassword("Password")
	if password == "" {
		fatal("Password is required.")
	}

	path := prompt(reader, "Image file")
	if path == "" {
		fatal("Image file is required.")
	}
	info, err := os.Stat(path)
	if err != nil {
		fatal(fmt.Sprintf("Cannot read file: %v", err))
	}
	if info.IsDir() {
		fatal("That path is a directory, not a file.")
	}

	fmt.Println()
	fmt.Println(strings.Repeat("─", 40))
	fmt.Printf("%s%sSummary%s\n", colorBold, colorYellow, colorReset)
	fmt.Printf("  %-16s %s\n", "Endpoint:", baseURL+"/photos")
	fmt.Printf("  %-16s %s\n", "User:", email)
	fmt.Printf("  %-16s %s (%d bytes)\n", "File:", filepath.Base(path), info.Size())
	fmt.Println(strings.Repeat("─", 40))
	fmt.Println()

	if !promptBool(reader, "Upload this file?") {
		fmt.Printf("%sAborted.%s\n\n", colorYellow, colorReset)
		return
	}

	token, err := login(baseURL, email, password)
	if err != nil {
		fatal(fmt.Sprintf("Login failed: %v", err))
	}
	fmt.Printf("%sLogged in.%s\n", colorGreen, colorReset)

	status, body, err := upload(baseURL, token, path)
	if err != nil {
		fatal(fmt.Sprintf("Upload request failed: %v", err))
	}

	fmt.Println()
	color := colorGreen
	if status < 200 || status >= 300 {
		color = colorRed
	}
	fmt.Printf("%s%sResponse: %d %s%s\n", colorBold, color, status, http.StatusText(status), colorReset)
	if len(body) > 0 {
		fmt.Println(string(body))
	}
	fmt.Println()

	if status < 200 || status >= 300 {
		os.Exit(1)
	}
}

// login posts credentials to /login and returns the access token.
func login(baseURL, email, password string) (string, error) {
	resp, err := http.PostForm(baseURL+"/login", url.Values{
		"email":    {email},
		"password": {password},
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var out struct {
		Token string `json:"token"`
		Error string `json:"error"`
	}
	if err := json.Unmarshal(data, &out); err != nil {
		return "", fmt.Errorf("unexpected response (%d): %s", resp.StatusCode, strings.TrimSpace(string(data)))
	}
	if resp.StatusCode != http.StatusOK || out.Token == "" {
		msg := out.Error
		if msg == "" {
			msg = resp.Status
		}
		return "", fmt.Errorf("%s", msg)
	}
	return out.Token, nil
}

// upload sends the file as multipart form field "photo" to POST /photos.
func upload(baseURL, token, path string) (int, []byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, nil, err
	}
	defer f.Close()

	// Sniff the content type so the part is labeled correctly.
	head := make([]byte, 512)
	n, _ := f.Read(head)
	contentType := http.DetectContentType(head[:n])
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return 0, nil, err
	}

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="photo"; filename=%q`, filepath.Base(path)))
	h.Set("Content-Type", contentType)
	part, err := w.CreatePart(h)
	if err != nil {
		return 0, nil, err
	}
	if _, err := io.Copy(part, f); err != nil {
		return 0, nil, err
	}
	if err := w.Close(); err != nil {
		return 0, nil, err
	}

	req, err := http.NewRequest(http.MethodPost, baseURL+"/photos", &buf)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return resp.StatusCode, body, err
}

func prompt(r *bufio.Reader, label string) string {
	fmt.Printf("%s%s:%s ", colorBold, label, colorReset)
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}

func promptPassword(label string) string {
	fmt.Printf("%s%s:%s ", colorBold, label, colorReset)
	b, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		fatal(fmt.Sprintf("Failed to read password: %v", err))
	}
	return strings.TrimSpace(string(b))
}

func promptBool(r *bufio.Reader, label string) bool {
	for {
		fmt.Printf("%s%s%s [y/n]: ", colorBold, label, colorReset)
		input, _ := r.ReadString('\n')
		switch strings.ToLower(strings.TrimSpace(input)) {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Printf("%sPlease enter y or n.%s\n", colorYellow, colorReset)
		}
	}
}

func fatal(msg string) {
	fmt.Printf("\n%s%sError:%s %s\n\n", colorBold, colorRed, colorReset, msg)
	os.Exit(1)
}
