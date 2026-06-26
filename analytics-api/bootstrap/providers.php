<?php

return [
    App\Modules\Auth\AuthServiceProvider::class,
    App\Modules\Quotes\QuotesServiceProvider::class,
    App\Modules\Exchanges\ExchangesServiceProvider::class,
    App\Modules\Pairs\PairsServiceProvider::class,

    App\Providers\AppServiceProvider::class,
    App\Providers\TelescopeServiceProvider::class,
];
