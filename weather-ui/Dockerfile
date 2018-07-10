FROM node:8.1-alpine

COPY . /app
WORKDIR /app

RUN chmod +x run.sh && \
    ./run.sh

CMD ["serve", "-s", "build"]
