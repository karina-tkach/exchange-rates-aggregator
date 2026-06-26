<?php

declare(strict_types=1);

namespace App\Modules\Pairs;

use App\Modules\Pairs\Application\Services\PairService;
use App\Modules\Pairs\Infrastructure\Repositories\EloquentPairRepository;
use App\Modules\Pairs\Infrastructure\Repositories\PairRepository;
use Illuminate\Support\Facades\Route;
use Illuminate\Support\ServiceProvider;

class PairsServiceProvider extends ServiceProvider
{
    public function register(): void
    {
        $this->app->singleton(PairService::class);

        $this->app->bind(
            PairRepository::class,
            EloquentPairRepository::class
        );
    }

    public function boot(): void
    {
        Route::middleware('api')
            ->prefix('api/v1')
            ->group(__DIR__ . '/routes.php');
    }
}
