FROM golang:alpine
WORKDIR '/app'
COPY *.go ./
COPY *.mod ./
COPY ./public/* ./public/
CMD ["go", "run", "."]