<?php

declare(strict_types=1);

use App\Modules\Auth\Http\Controllers\AuthController;

Route::middleware('throttle:60,1')
    ->prefix('auth')
    ->group(static function(): void {
        Route::post('login', [AuthController::class, 'login']);
        Route::post('register', [AuthController::class, 'register']);

        Route::group(['middleware' => 'auth:api'], function () {
            Route::post('logout', [AuthController::class, 'logout']);
            Route::post('refresh', [AuthController::class, 'refresh']);
        });
    });
