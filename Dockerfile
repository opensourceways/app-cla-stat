FROM golang:latest as BUILDER

MAINTAINER TommyLike<tommylikehu@gmail.com>

# build binary
COPY . /go/src/github.com/opensourceways/app-cla-stat
RUN cd /go/src/github.com/opensourceways/app-cla-stat && GO111MODULE=on CGO_ENABLED=0 go build -o cla-stat

# copy binary config and utils
FROM golang:latest
RUN  mkdir -p /opt/app/
COPY ./deploy/app.conf /opt/app/conf/app.conf
COPY  --from=BUILDER /go/src/github.com/opensourceways/app-cla-stat/cla-stat /opt/app

WORKDIR /opt/app/
ENTRYPOINT ["/opt/app/cla-stat"]
