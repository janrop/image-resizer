# Image resizer CLI-Util

```bash
Usage: 
  resize [int] # To resize all pictures with a max width and height of 400")
          -v Verbose output
```

Erwartet folgende Ordner-Struktur

```
- Zielordner/
  - [xyz]/
    - originals/
      - [Quelldatei].jpg ...
      - [Quelldateien].png ...
    - [dimensionen]/
    - .../
  - [abc]/
```

zum compilieren für Linux

```bash
env GOOS=linux GOARCH=amd64 go build -v
```