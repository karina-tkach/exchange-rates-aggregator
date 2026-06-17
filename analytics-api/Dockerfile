FROM php:8.2-cli
# Встановлення системних залежностей
RUN apt-get update && apt-get install -y \
git \
curl \
zip \
unzip \
libzip-dev \
libpq-dev \
nodejs npm \
&& rm -rf /var/lib/apt/lists/*
# Встановлення розширень PHP для роботи з базами даних
RUN docker-php-ext-install pdo pdo_pgsql
# Встановлення Composer з офіційного образу
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer
# Встановлення розширення Redis
RUN pecl install redis && docker-php-ext-enable redis
# Робоча директорія
WORKDIR /var/www/html
# Встановлення прав доступу
RUN chown -R www-data:www-data /var/www/html
