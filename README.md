A high-performance, lightweight URL shortening service built with Go. This project uses cryptographic hashing to generate short, deterministic IDs with zero external dependencies.

## Features

* **Deterministic IDs**: Uses SHA-256 hashing so the same URL always generates the same short code.
* **Thread-Safe**: Implements `sync.RWMutex` to safely handle concurrent read/write requests.
* **Standard Library Only**: Built using only Go's built-in packages (`net/http`, `crypto/sha256`, etc.).
* **Fast Redirection**: Uses HTTP 302 status codes for instant browser redirection.

## Technical Overview

The application works by taking a long URL, hashing it into a unique fingerprint, and storing it in a thread-safe map.

| Component | Technology | Purpose |
| --- | --- | --- |
| **Server** | `net/http` | Handles routing and requests |
| **Hashing** | `crypto/sha256` | Generates unique IDs from URLs |
| **Encoding** | `base64` | Converts hash bytes to URL-safe strings |
| **Concurrency** | `sync.RWMutex` | Prevents data races during high traffic |

---

## How to Run

1. **Clone the project**
```bash
git clone https://github.com/Anchalrawat/url-shorter.git
cd url-shorter

```


2. **Start the server**
```bash
go run main.go

```


The server will start on `http://localhost:8090`.

---

## API Usage

### 1. Shorten a URL

**Endpoint:** `POST /shorten`

**Body:** `url=<your_url>`

```bash
curl -X POST -d "url=https://www.google.com" http://localhost:8090/shorten

```

**Response:** `Shortened URL: http://localhost:8090/q8_7zA21`

### 2. Redirect

**Endpoint:** `GET /{id}`

Simply paste the shortened URL into your browser, and you will be redirected to the original destination.

