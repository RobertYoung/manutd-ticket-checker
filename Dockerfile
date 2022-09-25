FROM debian:bullseye-slim
# slim?
RUN apt update && apt install -y chromium

COPY manutd-ticket-checker /

ENTRYPOINT ["/manutd-ticket-checker"]