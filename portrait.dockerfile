FROM golang:1.12.7 as go-builder
COPY go /go/nwn-do
RUN apt update \
    && cd /go/nwn-do \
    && go mod download \
    && go build -o ./bin/b \
    && mv bin/b /usr/local/bin/
FROM ubuntu:latest 
LABEL maintainer "urothis@gmail.com"
COPY --from=go-builder /usr/local/bin/b /usr/local/bin/upload
COPY tga /tga
RUN apt update \
    && apt-get install rename imagemagick ca-certificates -y
CMD cd /tga \
    && echo converting tga\
    && rename -f 'y/A-Z/a-z/' * \
    && for i in *_t.tga; do convert $i -crop 16x25+0+0 $i.png; echo $i converted; done \
    && for i in *_s.tga; do convert $i -crop 32x50+0+0 $i.png; echo $i converted; done \
    && for i in *_m.tga; do convert $i -crop 64x100+0+0 $i.png; echo $i converted; done \
    && for i in *_l.tga; do convert $i -crop 128x200+0+0 $i.png; echo $i converted; done \
    && for i in *_h.tga; do convert $i -crop 256x400+0+0 $i.png; echo $i converted; done \
    && rename s/tga.png/png/ *.png \
    && mkdir /png && mv *.png /png\
    && upload


