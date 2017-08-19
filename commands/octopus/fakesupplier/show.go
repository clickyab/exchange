package main

const show = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Show</title>
</head>
<body>
{{ range $key, $value := . }}
<div style="border: 1px solid #ccc;margin: 10px;">
    {{ $value.Width }} * {{ $value.Height }}
    "{{ $value.Code }}
</div>
{{ end }}
</body>
</html>
`
