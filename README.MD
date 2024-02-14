# ascii-art-web

This will host a local website containing the ascii-art project.

Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII, this project was originally made for kood / Johvi.

## Usage

Syntax:

```
go build .
./asciiwebkood
```

In case of the file not properly executing, be aware that the file should have the right permissions, you check that with

```
l
```

And if needed, add the permissions.

```
chmod +x audit.sh
```
You can also just do

```
go run .
```

Endpoints should be managed perfectly with the error handlers in 52 and 59 in main.go, you can see the current status in your navigator inspecting the website and then going to Network and reloading the page again for scan for HTTP requests.


## Author

Federico Alvarez