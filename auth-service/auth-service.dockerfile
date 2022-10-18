FROM php:8.1.1-fpm-alpine

RUN docker-php-ext-install pdo pdo_mysql

RUN apk add shadow && usermod -u 1000 www-data && groupmod -g 1000 www-data

RUN set -eux; apk add libzip-dev; docker-php-ext-install zip
## Setup supervisor
#RUN apk add supervisor curl

#COPY docker/worker/supervisord.conf /etc/supervisord.conf
#
#COPY docker/worker/laravel.conf /etc/supervisor/conf.d/laravel.conf
#
#COPY docker/worker/phpfpm.conf /etc/supervisor/conf.d/phpfpm.conf
#
#CMD sh docker/worker/worker.sh

#COPY deploy/php/php.ini /etc/php/8.1/fpm/pool.d
#

# RUN apk add zlib-dev
# RUN docker-php-ext-configure zip --with-zlib-dir=/usr
# RUN apk add docker-php-ext-install zip
# RUN apk add librdkafka-dev
RUN docker-php-ext-install pcntl
RUN docker-php-ext-configure pcntl --enable-pcntl
## Setup Redis
# RUN docker-php-ext-install librdkafka
## Setup Xdebug
# RUN pecl install xdebug \
#     && docker-php-ext-enable xdebug
ENV LIBRDKAFKA_VERSION 1.10.2
ENV PHP_RDKAFKA_VERSION 5.0.0

# RUN apk add --no-cache --virtual .build-deps \
#     $PHPIZE_DEPS \
#     autoconf \
#     gcc \
#     g++ \
#     make \
#     libtool \
#     libxml2-dev \
#     oniguruma-dev \
#     && apk add --no-cache \
#     git
# COPY docker/php/xdebug.ini /usr/local/etc/php/conf.d/lands_xdebug.ini
# RUN git clone --depth 1 --branch v$LIBRDKAFKA_VERSION https://github.com/edenhill/librdkafka.git \
#     && cd librdkafka \
#     && ./configure \
#     && make \
#     && make install

# RUN pecl channel-update pecl.php.net \
#     && pecl install rdkafka-$PHP_RDKAFKA_VERSION \
#     && docker-php-ext-enable rdkafka

# RUN rm -rf /librdkafka \
#     && apk del -f .build-deps

RUN apk update \
    && php -r "copy('https://getcomposer.org/installer', '/tmp/composer-setup.php');" \
    && php /tmp/composer-setup.php --no-ansi --install-dir=/usr/local/bin --filename=composer \
    && rm -rf /tmp/composer-setup.php

# RUN composer require mateusjunges/laravel-kafka


# Add user for laravel application
# RUN groupadd -g 1000 www
# RUN useradd -u 1000 -ms /bin/bash -g www www

WORKDIR /var/www
# Copy existing application directory contents
COPY . .

# Expose port 9000 and start php-fpm server
EXPOSE 9000
CMD ["php-fpm"]