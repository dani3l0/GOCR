# GOCR

A tiny API that converts text from images into raw text


# Using curl

Let's get raw text from image called `myimage.png`:

```
curl -F 'image=@meme.jpg' http://localhost:8080
```


# How to run

Just download binary file from releases and run it

-----

**CGO HAS TO BE ENABLED!**

(but just during compilation)

and, tesseract dev libs as well