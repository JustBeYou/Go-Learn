go run fancy_draw.go | sed 's/^......//' | base64 -d > img.png; feh img.png
