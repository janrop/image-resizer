# Image resizer CLI-Util

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

zum compilieren für Linux

```bash
env GOOS=linux GOARCH=amd64 go build -v
```