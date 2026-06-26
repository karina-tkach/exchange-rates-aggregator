<?php

declare(strict_types=1);

use App\Modules\Pairs\Http\Controllers\PairController;

Route::middleware('throttle:60,1')
    ->group(static function(): void {
        Route::get('/pairs/names', [PairController::class, 'index'])
            ->name('pairs.names.index');
    });
