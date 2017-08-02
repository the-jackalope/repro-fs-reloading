FROM debian

WORKDIR /app

COPY ./watch-events-in-mounted /usr/bin/watch-events-in-mounted
CMD watch-events-in-mounted
