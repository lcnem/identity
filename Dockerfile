FROM centos:latest

RUN curl -O https://dl.google.com/go/go1.13.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && \
    chmod -R 777 "$GOPATH"

RUN yum update -y && \
    yum install git -y

RUN git clone https://21900d196b19ea1479417b70bd317c8449b66d98:x-oauth-basic@github.com/lcnem/identity.git && \
    git clone https://21900d196b19ea1479417b70bd317c8449b66d98:x-oauth-basic@github.com/lcnem/identity-public.git && \
    cd identity && \
    go install ./cmd/identitycli && \
    go install ./cmd/identityd && \
    cp identityd.service /etc/systemd/system/identityd.service && \
    cd ../ && \
    rm -rf identity && \
    systemctl enable identityd

EXPOSE 26656
EXPOSE 26657
