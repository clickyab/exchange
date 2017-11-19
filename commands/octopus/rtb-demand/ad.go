package main

import (
	"context"
	"net/http"
)

var ad = `<!doctype html>
<html lang="en">
<head>
<meta charset="UTF-8">
             <meta name="viewport" content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
                         <meta http-equiv="X-UA-Compatible" content="ie=edge">
             <title>Test done</title>
</head>
<body>
  <h1>Congratulation!!</h1>
  <p></p>
</body>
</html>`

func adHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.Write([]byte(ad))
}
