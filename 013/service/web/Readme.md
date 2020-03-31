# web application that uses db

## description

This creates a web application which uses a db.
The webpod.yaml stitches it with a postgres container in the same pod.

commands supported
1. GET /read     returns all posted lines
2. POST /write   { "data":"any line you want" }
3. GET /quit     kills the web application

## setup

`make web`


