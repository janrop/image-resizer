# Image resizer CLI-Util

```bash
Usage: 
  resize
          -w [int] (required) Resize all pictures to a max width of [int]
          -t [int] Only resize files older than [int] seconds
          -v Verbose output
          [target_directory] to resize files in (default: ./)
```

Expects the following directory structure and creates additional [dimension]-folders in every sub-directory of [target_directory]

```
- target_directory/
  - [xyz]/
    - originals/
      - [Quelldatei].jpg ...
      - [Quelldateien].png ...
    - [dimension]/
    - .../
  - [abc]/
```
