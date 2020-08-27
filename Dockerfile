FROM ubuntu:18.04

RUN apt-get update
RUN apt-get install -y software-properties-common
RUN add-apt-repository ppa:libreoffice/ppa
RUN apt-get update
RUN apt-get install -y --force-yes libreoffice
RUN apt-get clean

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY ./bin/pdf-converter /usr/src/app

EXPOSE 3000

CMD [ "/usr/src/app/pdf-converter" ]