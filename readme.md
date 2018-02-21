# Image resizer CLI-Util

```bash
Usage: 
  resize
          -w [int] (required) Resize all pictures to a max width of [int]
          -t [int] Only resize files older than [int] seconds
          -v Verbose output
          [target_directory] to resize files in (default: ./)
```

Erwartet folgende Ordner-Struktur

```
- Target-Directory/
  - [xyz]/
    - originals/
      - [Quelldatei].jpg ...
      - [Quelldateien].png ...
    - [dimensionen]/
    - .../
  - [abc]/
```
