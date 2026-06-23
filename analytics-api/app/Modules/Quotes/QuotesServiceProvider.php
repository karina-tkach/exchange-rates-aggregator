<?php

declare(strict_types=1);

namespace App\Modules\Quotes;

use App\Modules\Quotes\Application\Services\QuoteService;
use App\Modules\Quotes\Infrastructure\Repositories\EloquentQuoteRepository;
use App\Modules\Quotes\Infrastructure\Repositories\OhlcRepository;
use App\Modules\Quotes\Infrastructure\Repositories\QuoteRepository;
use App\Modules\Quotes\Infrastructure\Repositories\TimescaleOhlcRepository;
use Illuminate\Support\Facades\Route;
use Illuminate\Support\ServiceProvider;

class QuotesServiceProvider extends ServiceProvider
{
    public function register(): void
    {
        $this->app->singleton(QuoteService::class);

        $this->app->bind(
            QuoteRepository::class,
            EloquentQuoteRepository::class
        );
        $this->app->bind(
            OhlcRepository::class,
            TimescaleOhlcRepository::class
        );
    }

    public function boot(): void
    {
        Route::middleware('api')
            ->prefix('api/v1')
            ->group(__DIR__ . '/routes.php');
    }
}
