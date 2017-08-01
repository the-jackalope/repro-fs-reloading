FROM debian

#ENV DEMETER_ENV $DEMETER_ENV
ENV INSTALL_DIR '/app'

RUN mkdir $INSTALL_DIR
WORKDIR $INSTALL_DIR
RUN mkdir mounted

COPY ./watch-events-in-mounted /usr/bin/watch-events-in-mounted
CMD watch-events-in-mounted
