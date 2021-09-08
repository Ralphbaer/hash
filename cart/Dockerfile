FROM golang:1.8

WORKDIR /app

ADD ./.bin/cart-service .

COPY Dockerfile ./gen/swagger.yaml* ./gen/
COPY Dockerfile ./repository/products.json* ./repository/

RUN rm ./gen/Dockerfile

RUN groupadd -r appusers && useradd --no-log-init -r -g appusers appuser

EXPOSE 9091

ENTRYPOINT ["./cart-service", "0.0.0.0"]