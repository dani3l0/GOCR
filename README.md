# GOCR

A tiny API that converts text from images into raw text


# Using curl

Let's get raw text from image called `meme.png`:

```
curl -F 'image=@meme.jpg' http://localhost:8080
```


# How to run

Just download binary file from releases and run it


# Configuration

**`config.yaml` file will be created on first run:**

- `listenaddr` - IP address and port HTTP server is listening on
- `imagedir` - path where images will be saved
- `uploadlimit` - upload limit in megabytes


# Development

**CGO HAS TO BE ENABLED!**

(but just during compilation)

and, tesseract dev libs as well