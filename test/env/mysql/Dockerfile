# Copyright (c) 2022 xxx xxx Team. All rights reserved.

ARG REGISTRY_HOST=xxx.xxx.com.cn

FROM $REGISTRY_HOST/xxx-base/mysql:5.7

COPY . /docker-entrypoint-initdb.d

RUN chmod -R 0755 /docker-entrypoint-initdb.d
