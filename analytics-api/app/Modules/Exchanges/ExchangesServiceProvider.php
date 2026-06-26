<?php

declare(strict_types=1);

namespace App\Modules\Exchanges;

use App\Modules\Exchanges\Application\Services\ExchangeService;
use App\Modules\Exchanges\Infrastructure\Repositories\EloquentExchangeRepository;
use App\Modules\Exchanges\Infrastructure\Repositories\ExchangeRepository;
use Illuminate\Support\Facades\Route;
use Illuminate\Support\ServiceProvider;

class ExchangesServiceProvider extends ServiceProvider
{
    public function register(): void
    {
        $this->app->singleton(ExchangeService::class);

        $this->app->bind(
            ExchangeRepository::class,
            EloquentExchangeRepository::class
        );
    }

    public function boot(): void
    {
        Route::middleware('api')
            ->prefix('api/v1')
            ->group(__DIR__ . '/routes.php');
    }
}
