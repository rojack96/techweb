# TechWeb - Build & Deploy (Minimalist)

Build e deploy locale senza registry remoti.

## 🚀 Quick Start

```bash
# 1. Build (auto-incrementa versione)
./build.sh

# 2. Start services
./deploy.sh up

# 3. Access
#   Frontend:   http://localhost:3000
#   Backend:    http://localhost:8080

# 4. View logs
./deploy.sh logs

# 5. Stop
./deploy.sh down
```

## 📝 Comandi

```bash
# Build
./build.sh           # Full build (backend + frontend + images)
./build.sh backend   # Backend only
./build.sh frontend  # Frontend only  
./build.sh images    # Docker images only
./build.sh clean     # Clean artifacts
./build.sh help      # Help

# Deploy
./deploy.sh up       # Start services
./deploy.sh down     # Stop services
./deploy.sh logs     # View logs
./deploy.sh help     # Help
```

## 📁 Structure

```
.version                        # Version file (auto-incremented)
build.sh                        # Build script
deploy.sh                       # Deploy script
docker-compose.yml              # Services config
Dockerfile                      # (backend & frontend)
```

## 🐳 Docker Images

**Local only (no registry):**
- `streetcats-api:latest`   (backend)
- `streetcats-ui:latest`    (frontend)

Versioned: `streetcats-api:1.0.1`, etc.

## 📊 Versioning

Version automatically incremented from `.version` file:
```
1.0.0 → 1.0.1 → 1.0.2 → ...
```

## 📚 Documentation

- [QUICK_START.md](QUICK_START.md) - 3-minute guide
- [README_BUILD.md](README_BUILD.md) - This file

---

Minimalist. Simple. Done. 🎉
