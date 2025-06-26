FROM golang:1.23 AS builder

WORKDIR /opt/app
COPY ./ /opt/app/

RUN cd cmd \
  && go mod edit -replace=github.com/bnjns/rich-chat-statuses=../ \
  && go mod edit -replace=github.com/bnjns/rich-chat-statuses/calendars/google=../calendars/google \
  && go mod edit -replace=github.com/bnjns/rich-chat-statuses/clients/slack=../clients/slack \
  && go mod tidy

RUN go build -C cmd -tags lambda.norpc -o dist/rich-chat-statuses

FROM public.ecr.aws/lambda/provided:al2023

COPY --from=builder /opt/app/cmd/dist/rich-chat-statuses /usr/local/bin/rich-chat-statuses

ENTRYPOINT ["/usr/local/bin/rich-chat-statuses"]
