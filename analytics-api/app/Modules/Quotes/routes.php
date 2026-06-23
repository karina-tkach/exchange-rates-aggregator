<?php

declare(strict_types=1);

use App\Modules\Quotes\Http\Controllers\OhlcController;
use App\Modules\Quotes\Http\Controllers\RateController;
use Illuminate\Support\Facades\Route;

Route::middleware('throttle:60,1')
    ->group(static function(): void {
        Route::get('/rates', [RateController::class, 'index'])
            ->name('rates.index');
        Route::get('/spread', [RateController::class, 'spread'])
            ->name('rates.spread');
        Route::get('/ohlc', [OhlcController::class, 'index'])
        ->name('ohlc.index');
    });
