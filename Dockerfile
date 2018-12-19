FROM ubuntu:18.04

RUN apt-get update && apt-get install git -y
RUN cd /root; git clone https://github.com/hakimel/reveal.js.git
COPY go-present /usr/bin/go-present
RUN chmod +x /usr/bin/go-present
EXPOSE 8080
CMD ["/bin/bash", "-l", "-c"]
ENTRYPOINT ["go-present"]
