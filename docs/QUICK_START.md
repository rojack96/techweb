# TechWeb - Quick Start (Minimalist)

Build e deploy locale in 3 minuti.

## 🚀 Start

```bash
# 1. Build everything 
./build.sh

# Output: immagini Docker create e versione incrementata

# 2. Start services
./deploy.sh up

# 3. Access
# Frontend:   http://localhost:3000
# Backend:    http://localhost:8080

# 4. View logs
./deploy.sh logs

# 5. Stop
./deploy.sh down
```

## 📊 Versione

La versione si trova in `.version` e viene **auto-incrementata** ad ogni build.

Esempio flusso:
```
Build #1: versione 1.0.0
Build #2: versione 1.0.1  (auto-incrementata)
Build #3: versione 1.0.2  (auto-incrementata)
```

## 🔧 Comandi Disponibili

### build.sh
```bash
./build.sh           # Build everything (backend + frontend + images)
./build.sh backend   # Solo backend
./build.sh frontend  # Solo frontend
./build.sh images    # Solo immagini Docker
./build.sh clean     # Pulisci artifact
./build.sh help      # Help
```

### deploy.sh
```bash
./deploy.sh up       # Start services
./deploy.sh down     # Stop services
./deploy.sh logs     # View all logs
./deploy.sh logs-api    # View backend logs
./deploy.sh logs-ui     # View frontend logs
./deploy.sh status   # Service status
./deploy.sh help     # Help
```

## 🐳 Immagini Docker

Build **locali** (no registry remoti):
- `streetcats-api:latest`  (backend)
- `streetcats-ui:latest`   (frontend)

Versionate da `.version`: `streetcats-api:1.0.1`, etc.

## 🆘 Troubleshooting

**Port in use:**
```bash
./deploy.sh down
```

**Docker not running:**
```bash
docker ps  # Verifica
# Avvia Docker
```

**Build failed:**
```bash
./build.sh help  # Guida
```

---

**Fatto!** Minimale e semplice. 🎉

