FROM docker.jsjit.cn/golang:1.12 as appBuild
FROM scratch as prod

EXPOSE 8080

COPY --from=appBuild /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=appBuild /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ADD app /app

CMD ["/app"]
