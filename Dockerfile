FROM golang

RUN go get github.com/gorilla/mux

RUN go get golang.org/x/oauth2

RUN	go get golang.org/x/oauth2/google

RUN	go get google.golang.org/api/analytics/v3

WORKDIR /go/src/github.com/heaptracetechnology/google-analytics

ADD . /go/src/github.com/heaptracetechnology/google-analytics

RUN go install github.com/heaptracetechnology/google-analytics

ENTRYPOINT google-analytics

EXPOSE 3000