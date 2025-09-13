FROM php:8.4-fpm-alpine as node

WORKDIR /var/www/html

COPY --from=composer /usr/bin/composer /usr/bin/composer

COPY ./composer* ./

RUN composer install

COPY ./ ./

CMD ["php-fpm"]
