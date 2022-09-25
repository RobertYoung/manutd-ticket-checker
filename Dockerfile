FROM debian:bullseye-slim

RUN apt update && apt install -y chromium

COPY manutd-ticket-checker /usr/bin/

ENTRYPOINT ["/usr/bin/manutd-ticket-checker"]