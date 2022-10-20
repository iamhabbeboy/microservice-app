FROM php:8.1.1-fpm-alpine

RUN apk add shadow && usermod -u 1000 www-data && groupmod -g 1000 www-data
RUN set -eux; apk add libzip-dev; docker-php-ext-install zip
RUN apk add --no-cache --update --virtual buildDeps autoconf
RUN docker-php-ext-install pcntl
RUN docker-php-ext-configure pcntl --enable-pcntl
ADD https://github.com/mlocati/docker-php-extension-installer/releases/latest/download/install-php-extensions /usr/local/bin/
RUN chmod +x /usr/local/bin/install-php-extensions && \
    install-php-extensions rdkafka

RUN apk update \
    && php -r "copy('https://getcomposer.org/installer', '/tmp/composer-setup.php');" \
    && php /tmp/composer-setup.php --no-ansi --install-dir=/usr/local/bin --filename=composer \
    && rm -rf /tmp/composer-setup.php

RUN composer require mateusjunges/laravel-kafka

WORKDIR /var/www
COPY . .

EXPOSE 9000
CMD ["php-fpm"]