FROM redis:alpine
COPY ./redis.conf /usr/local/etc/redis/redis.conf
EXPOSE 6379
VOLUME [ "/data" ]
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]