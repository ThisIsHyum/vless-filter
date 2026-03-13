# vless-filter
VLESS Filtering Server

## Installation and running
### binary
1. download release
2. run server
```bash
vless-<os>-<arch>
```
### compile
1. clone repository
```bash
git clone https://github.com/thisishyum/vless-filter.git
cd vless-filter
```
2. compile or run server without compiling
```bash
go build .
# or
go run .
```
3. if you compiled, run server
```bash
./vless-filter
```
### docker
1. clone repository
```bash
git clone github.com/thisishyum/vless-filter
cd vless-filter
```
2. build image
```bash
docker build . -t vless-fiter
```
3. run container
```bash
docker run -p 2505:80 -v ./sub_urls.txt:/app/sub_urls.txt vless-filter
```

## Configuration

You can configure vless-filter using CLI flags or environment variables.

1. **host**  
   IP address or domain for the server to bind.  
   Default: `127.0.0.1`

1. **port**  
   Server port.  
   Default: `80`

1. **interval**  
   Subscription update interval.  
   Default: `30m`

1. **timeout**  
   Timeout per node check.  
   Default: `3s`

1. **workers**  
   Number of concurrent workers.  
   Default: `200`

1. **subs**  
   Path to `sub_urls.txt`.  
   Default: `sub_urls.txt`

### sub_urls.txt

`sub_urls.txt` contains subscription URLs.  
Each URL must be on a new line.

Comments are supported using `#`.

## Usage

Servers can be retrieved from the `/subs` endpoint.
You can also use query parameters for filtering:

1. `limit` — limit number of returned servers
2. `max_latency` — filter servers by maximum latencyalso you can use queries for filtering:

### example
```bash
curl "http://localhost:80/subs?limit=10&max_latency=20ms"
```
```
vless://UUID@IP:PORT?flow=xtls-rprx-vision&encryption=none&type=tcp&security=reality&fp=chrome&sni=SNI_URL&pbk=PBK_UUID
vless://UUID@IP:PORT?flow=xtls-rprx-vision&encryption=none&type=tcp&security=reality&fp=chrome&sni=SNI_URL&pbk=PBK_UUID
vless://UUID@IP:PORT?flow=xtls-rprx-vision&encryption=none&type=tcp&security=reality&fp=chrome&sni=SNI_URL&pbk=PBK_UUID
vless://UUID@IP:PORT?flow=xtls-rprx-vision&encryption=none&type=tcp&security=reality&fp=chrome&sni=SNI_URL&pbk=PBK_UUID
vless://UUID@IP:PORT?flow=xtls-rprx-vision&encryption=none&type=tcp&security=reality&fp=chrome&sni=SNI_URL&pbk=PBK_UUID
vless://UUID@IP:PORT?flow=xtls-rprx-vision&encryption=none&type=tcp&security=reality&fp=chrome&sni=SNI_URL&pbk=PBK_UUID
vless://UUID@IP:PORT?flow=xtls-rprx-vision&encryption=none&type=tcp&security=reality&fp=chrome&sni=SNI_URL&pbk=PBK_UUID
vless://UUID@IP:PORT?flow=xtls-rprx-vision&encryption=none&type=tcp&security=reality&fp=chrome&sni=SNI_URL&pbk=PBK_UUID
vless://UUID@IP:PORT?flow=xtls-rprx-vision&encryption=none&type=tcp&security=reality&fp=chrome&sni=SNI_URL&pbk=PBK_UUID
vless://UUID@IP:PORT?flow=xtls-rprx-vision&encryption=none&type=tcp&security=reality&fp=chrome&sni=SNI_URL&pbk=PBK_UUID
```
