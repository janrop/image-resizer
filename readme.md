# Image resizer CLI-Util für Prämienshop

Erwartet folgende Struktur für die Bilder

```
- Zielordner/
  - [xyz]/
    - originals/
      - Quelldateien ...
      - Quelldateien ...
    - [dimensionen]/
    - .../
  - [abc]/
```

zum compilieren für TOB-Server

```bash
env GOOS=linux GOARCH=amd64 go build -v
```