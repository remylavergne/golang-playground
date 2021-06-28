# Golang Playground Scraper

First ever project to learn Golang.

This 'scraper' allows to download (at this time) JPG, GIF & WEBM.

## How to use

```bash
go run . --url https://wordpress.org/plugins/image-wall/ -e jpg

// Output:
// 6 files are going to be downloaded
// File output/ab398fea5005cbd986afe0b88f6717dc/screenshot-2.jpg downloaded (size: 29869)
// File output/ab398fea5005cbd986afe0b88f6717dc/banner-772x250.jpg downloaded (size: 51168)
// File output/ab398fea5005cbd986afe0b88f6717dc/icon-256x256.jpg downloaded (size: 93657)
// File output/ab398fea5005cbd986afe0b88f6717dc/banner-1544x500.jpg downloaded (size: 181728)
// File output/ab398fea5005cbd986afe0b88f6717dc/icon-128x128.jpg downloaded (size: 186488)
// File output/ab398fea5005cbd986afe0b88f6717dc/screenshot-1.jpg downloaded (size: 297585)
// Process done!
```

Use `--dry-run` option to estimate number of available files without downloading them.

