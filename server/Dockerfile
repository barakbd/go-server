# Reference
# https://fabianlee.org/2020/01/26/golang-using-multi-stage-builds-to-create-clean-docker-images/

# builder image
FROM golang:1.13-alpine3.11 as builder
RUN mkdir /build
ADD *.go /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o golang-memtest .


# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /build/golang-memtest .

# executable
ENTRYPOINT [ "./golang-memtest" ]
# arguments that can be overridden
CMD [ "3", "300" ]
