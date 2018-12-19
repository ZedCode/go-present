# Go-Present
The tool `go-present` is a simple Golang wrapper around `reveal.js` that takes a Markdown file, injects it into the template, and starts up a docker container/http server to serve the presentation.

## Usage
To use this tool you only need Docker installed on your system. To build the tool and the docker container run `make linux` from this repo directory. This will output something like the following:
```
➜ make linux
docker run --rm -v [path/to/current/dir]/go-present:/usr/src/go-present -w /usr/src/go-present golang:1.8 go build -v
_/usr/src/go-present
# We should also rebuild the main container.
docker build -t="go-present:latest" .
Sending build context to Docker daemon  8.049MB
Step 1/8 : FROM ubuntu:18.04
 ---> 93fd78260bd1
Step 2/8 : RUN apt-get update && apt-get install git -y
 ---> Using cache
 ---> d35f6f882c2f
Step 3/8 : RUN cd /root; git clone https://github.com/hakimel/reveal.js.git
 ---> Using cache
 ---> c855dbf72ab1
Step 4/8 : COPY go-present /usr/bin/go-present
 ---> e496ee3435a3
Step 5/8 : RUN chmod +x /usr/bin/go-present
 ---> Running in 4b22954492d6
Removing intermediate container 4b22954492d6
 ---> c006a577cdf4
Step 6/8 : EXPOSE 8080
 ---> Running in 6f41fb58fcc2
Removing intermediate container 6f41fb58fcc2
 ---> c83a495e818c
Step 7/8 : CMD ["/bin/bash", "-l", "-c"]
 ---> Running in c29e2f2b44cd
Removing intermediate container c29e2f2b44cd
 ---> 9239e5eb31c8
Step 8/8 : ENTRYPOINT ["go-present"]
 ---> Running in b2b1c5a39ae6
Removing intermediate container b2b1c5a39ae6
 ---> 8139a4b2a240
Successfully built 8139a4b2a240
Successfully tagged go-present:latest
```
Now, to start creating presentations, start adding them to the `md` directory. For example, if you had a presentation called "sprockets.md" you would start your container as:
```
docker run -d -p 8080:8080 -v `pwd`/md:/root/md go-present -p sprockets.md
```
Optionally, you can override the `reveal.js` theme by passing the `-t` flag, for example:
```
docker run -d -p 8080:8080 -v `pwd`/md:/root/md go-present -p sprockets.md -t black
```
A list of available templates can be found in the `reveal.js` repo [HERE](https://github.com/hakimel/reveal.js/tree/master/css/theme) -- note, don't include the `.css` in your command.

To see this working with the example markdown and default template, simply run the container without args:
```
docker run -d -p 8080:8080 -v `pwd`/md:/root/md go-present
```
