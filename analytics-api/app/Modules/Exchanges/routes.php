<?php

declare(strict_types=1);

use App\Modules\Exchanges\Http\Controllers\ExchangeController;

Route::middleware('throttle:60,1')
    ->group(static function(): void {
        Route::get('/exchanges/active', [ExchangeController::class, 'index'])
            ->name('exchanges.active.index');
    });
