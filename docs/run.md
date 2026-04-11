# Run backend

## Start

From the repository root:

```bash
cd backend
go run ./cmd/api
```

By default the server listens on `:8080`.

To use another port:

```bash
cd backend
PORT=9090 go run ./cmd/api
```

## Check health

When the server is running, verify the health endpoint:

```bash
curl http://localhost:8080/health
```

Expected response:

```json
{"status":"ok"}
```

If you changed the port, replace `8080` with your `PORT` value.
