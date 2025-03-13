# OAuth2 Package for Go

This package provides a simple and flexible OAuth2 client for Go applications. It supports authentication via various OAuth providers such as Google, Discord, and Slack.

## 🚀 Installation

```sh
go get github.com/gokhanaltun/oauth2
```

## 📌 Features
- Supports multiple OAuth providers.
- Easily configurable authentication flow.
- Handles extra parameters dynamically.
- Provides both raw and formatted API responses.
- Supports custom state generation.

## 🛠️ Usage

### 1️⃣ Initialize OAuth Client

```go
import (
    "fmt"
    "github.com/gokhanaltun/oauth2"
)

auth := oauth2.OAuth{
    Provider: oauth2.Providers.Google, // Choose a provider
    Config: oauth2.Config{
        ClientID:     "YOUR_CLIENT_ID",
        ClientSecret: "YOUR_CLIENT_SECRET",
        RedirectURL:  "YOUR_REDIRECT_URL",
        Scopes:       []string{"openid", "email", "profile"},
        ExtraParams:  map[string]string{"access_type": "offline"},
        StateFunc:    func() string { return "random-state-value" },
    },
}
```

### 2️⃣ Get Authorization URL

```go
authURL, err := auth.AuthCodeURL()
if err != nil {
    fmt.Println("Error generating auth URL:", err)
    return
}
fmt.Println("Visit this URL to authenticate:", authURL)
```

### 3️⃣ Exchange Authorization Code for Access Token

```go
exchangeResponse, err := auth.Exchange("authorization_code")
if err != nil {
    fmt.Println("Error exchanging token:", err)
    return
}

fmt.Println("Access Token:", exchangeResponse.AccessToken)
```

## 🔍 Handling API Responses: Raw vs Formatted

The `Exchange` function returns both **raw** and **formatted** API responses. These are useful in different scenarios.

### 1️⃣ Using `FormattedResponseBody`

Most of the time, you can use `FormattedResponseBody`, which contains the parsed JSON response.

```go
// Get additional response data
expiresIn, ok := exchangeResponse.FormattedResponseBody["expires_in"].(float64)
if ok {
    fmt.Println("Token expires in:", expiresIn, "seconds")
}
```

✅ **Use case:** When you expect a **valid JSON response** from the provider.

### 2️⃣ Debugging with `RawResponseBody`

If the API returns an **unexpected response** (e.g., an HTML error page), you can inspect `RawResponseBody` for debugging.

```go
if err != nil {
    fmt.Println("Error exchanging token:", err)
    fmt.Println("Raw response body:", string(exchangeResponse.RawResponseBody)) // Debugging
}
```

🔍 **Use case:** When the provider sends **non-JSON data**, causing `FormattedResponseBody` to be `nil`.

## ⚠️ Example Output for a Failed Request
```sh
Error exchanging token: failed to parse JSON: invalid character '<' looking for beginning of value
Raw response body: <html><body><h1>500 Internal Server Error</h1></body></html>
```

## ✅ Supported OAuth Providers

The package comes with built-in support for:
- **Google**
- **Discord**
- **Slack**

You can also define a custom provider:

```go
customProvider := oauth2.Provider{
    AuthURL:  "https://example.com/oauth/authorize",
    TokenURL: "https://example.com/oauth/token",
}
```

## 🔄 Extra Parameters

You can pass additional parameters using `ExtraParams`. Example:

```go
Config{
    ExtraParams: map[string]string{
        "prompt": "consent",
        "access_type": "offline",
    },
}
```

## 🛡️ Custom State Handling

If you need dynamic state generation:

```go
Config{
    StateFunc: func() string {
        return "secure-random-state"
    },
}
```

