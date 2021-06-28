# Golang Playground Scraper

First ever project to learn Golang.

This 'scraper' allows to download (at this time) JPG, GIF & WEBM.

## How to use

```bash
go run . --url https://wordpress.org/plugins/image-wall/ -e jpg

// Output:
// 6 files are going to be downloaded
// Downloaded a file output/screenshot-2.jpg with size 93657
// Downloaded a file output/banner-772x250.jpg with size 29869
// Downloaded a file output/icon-256x256.jpg with size 51168
// Downloaded a file output/screenshot-1.jpg with size 181728
// Downloaded a file output/banner-1544x500.jpg with size 186488
// Downloaded a file output/icon-128x128.jpg with size 297585
// Process done!
```

Use `--dry-run` option to estimate number of available files without downloading them.

