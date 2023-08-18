FROM golang:latest as BUILDER

MAINTAINER TommyLike<tommylikehu@gmail.com>

# build binary
COPY . /go/src/github.com/opensourceways/app-cla-stat
RUN cd /go/src/github.com/opensourceways/app-cla-stat && GO111MODULE=on CGO_ENABLED=0 go build -o cla-stat

# copy binary config and utils
FROM golang:latest
RUN groupadd -g 1000 cla
RUN useradd -u 1000 -g cla -s /bin/bash -m cla
USER cla
WORKDIR /home/cla
COPY --chown=cla ./deploy/app.conf /home/cla/conf/app.conf
COPY --chown=cla --from=BUILDER /go/src/github.com/opensourceways/app-cla-stat/cla-stat /home/cla

ENTRYPOINT ["/home/cla/cla-stat"]
