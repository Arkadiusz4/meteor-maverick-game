FROM ubuntu:latest
LABEL authors="amika"

RUN apt-get update
RUN apt-get install -y nginx

RUN echo "deamon off;" >> /etc/nginx/nginx.conf

EXPOSE 80

CMD ["/usr/bin/nginx", "-c", "/etc/nginx/nginx.conf"]
