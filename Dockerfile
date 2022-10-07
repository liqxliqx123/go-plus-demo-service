# Copyright (c) 2022 xxx xxx Team. All rights reserved.

ARG REGISTRY_HOST=xxx.xxx.com.cn
FROM $REGISTRY_HOST/xxx-base/ci-base-go:1.15.1 as build_base

ENV APP_DIR=/service/
WORKDIR $APP_DIR

ENV CMD_DIR=$APP_DIR/cmd
ENV DAILY_xxx_CRONJOB_DIR=$CMD_DIR/daily_xxx_cron
ENV xxx_CONSUMER=$CMD_DIR/xxx_consumer_group

COPY . .
RUN go build -mod=vendor -o /bin/app . \
    && echo "Building executable cmd..." \
    && go build -mod=vendor -o /bin/daily_xxx_cron $DAILY_xxx_CRONJOB_DIR \
    && go build -mod=vendor -o /bin/xxx_consumer $xxx_CONSUMER

# Start fresh from a smaller image
FROMxxx/xxx-base/go-base:1.15

WORKDIR /service/

COPY --from=build_base /bin/app ./app
COPY --from=build_base /bin/daily_xxx_cron /bin/daily_xxx_cron
COPY --from=build_base /bin/xxx_consumer /bin/xxx_consumer
ADD settings.yaml .
COPY assets ./assets

EXPOSE 9042
HEALTHCHECK --interval=5s --timeout=1s \
  CMD curl -fs http://localhost:8080/status || exit 1

CMD ["/service/app"]
