FROM php:8.4-fpm-alpine as node

WORKDIR /var/www/html

COPY --from=composer /usr/bin/composer /usr/bin/composer

COPY ./composer* ./

RUN composer install

RUN apk update
RUN apk add --no-cache --virtual .build-deps $PHPIZE_DEPS imagemagick-dev
RUN apk add libwebp-tools
RUN pecl install imagick
RUN docker-php-ext-enable imagick

COPY ./ ./

CMD ["php-fpm"]
