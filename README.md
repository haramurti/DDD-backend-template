# hackathon-be

Template backend Go + Gin + Gorm, struktur DDD per modul. Isinya kosongan: plumbing udah jalan, logic bisnis tinggal diisi di bagian `TODO`.

## Setup pertama kali

```bash
go mod tidy                      # bikin go.sum (WAJIB sebelum docker build)
cp .env.example .env
docker compose up -d postgres redis
go run ./cmd/api                 # cek: curl localhost:8080/health
```

## Status isi template

| Bagian | Status |
|---|---|
| config, koneksi DB, AutoMigrate, Redis connect | jalan |
| `pkg/response` (format JSON + mapping error → HTTP) | jalan |
| `pkg/jwt`, `pkg/utils/password`, middleware `Auth`, `RequireRole`, `CORS` | jalan |
| Wiring modul (repo → service → handler → routes), `/health`, graceful shutdown | jalan |
| Handler semua modul (bind, validasi, panggil service, response) | jalan |
| Repository & service semua modul | **TODO** (balas 501 Not Implemented) |
| `internal/realtime` (WebSocket hub) | **TODO** |
| `pkg/storage` (MinIO), middleware `Metrics`, request_id di `Logger`, `database.Seed` | **TODO** |
| Test | placeholder (`t.Skip`) |

## Struktur

```
cmd/
  api/main.go              entrypoint
  bootstrap/bootstrap.go   wiring infra + modul + server
config/                    env config
internal/
  app/
    user/                  auth (register, login, me)
    feature_a/             template fitur, ganti nama sesuai PRD
    feature_b/
      contract/            interface Repository & Service
      dto/                 request (binding tag) & response
      entity/              struct Gorm
      handler/             handler + routes
      repository/          akses DB
      service/             aturan bisnis + hak akses (+ test)
  infra/database/          connection, migration, seed
  infra/cache/             redis
  middleware/              jwt, role, cors, logger, metrics
  realtime/                websocket hub
pkg/
  jwt/ response/ storage/ utils/
deploy/prometheus.yml
.github/workflows/ci.yml   test -> deploy via SSH
```

Alur request: `routes → handler (bind) → service (aturan bisnis, akses, transaksi, cache, broadcast) → repository (query) → entity`.

## Ganti nama fitur

Contoh `feature_a` → `idea`:

```bash
git mv internal/app/feature_a internal/app/idea
git mv internal/app/idea/entity/feature_a.go internal/app/idea/entity/idea.go
grep -rl "feature_a\|FeatureA\|feature-a" --include=*.go . \
  | xargs sed -i 's/feature_a/idea/g; s/FeatureA/Idea/g; s/feature-a/ideas/g; s/featureA/idea/g'
go build ./...
```

(Di macOS pakai `sed -i ''`.) Modul baru: copy folder `feature_b`, rename dengan cara yang sama, lalu tambah 1 baris di `bootstrap.go` dan entity-nya di `migration.go`.

## Deploy ke VM

1. `ssh -i ./groupX -p 1111X dev@proxy.bccdev.id`
2. Install Docker kalau belum ada, clone repo ke `~/backend` (pakai deploy key read-only).
3. `cp .env.example .env`, isi secret, lalu `docker compose up -d --build`.
4. GitHub secrets: `SSH_PORT`, `SSH_PRIVATE_KEY`. Push ke `main` → test → deploy.

BE wajib di port VM `8080` → `https://groupX-be.bccdev.id`.
