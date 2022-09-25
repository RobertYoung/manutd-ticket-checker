FROM debian:bullseye 
# slim?
RUN apt update && apt install -y chromium

COPY ./dist/manutd-ticket-checker_linux_amd64_v1/manutd-ticket-checker /

ENTRYPOINT ["/manutd-ticket-checker"]