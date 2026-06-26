<?php

declare(strict_types=1);

use App\Modules\Exchanges\Http\Controllers\ExchangeController;

Route::middleware('throttle:60,1')
    ->group(static function(): void {
        Route::get('/exchanges/active', [ExchangeController::class, 'active'])
            ->name('exchanges.active');

        Route::middleware(['auth:api', 'role:admin'])
            ->prefix('exchanges')
            ->group(function () {
                Route::get('/', [ExchangeController::class, 'index']);
                Route::patch('/{id}', [ExchangeController::class, 'update']);
            });
    });
