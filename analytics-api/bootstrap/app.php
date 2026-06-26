<?php

use App\Http\Middleware\AddCorrelationId;
use App\Http\Middleware\RoleMiddleware;
use App\Http\Middleware\SecurityHeaders;
use Illuminate\Auth\AuthenticationException;
use Illuminate\Database\Eloquent\ModelNotFoundException;
use Illuminate\Foundation\Application;
use Illuminate\Foundation\Configuration\Exceptions;
use Illuminate\Foundation\Configuration\Middleware;
use Illuminate\Http\Exceptions\ThrottleRequestsException;
use Illuminate\Validation\ValidationException;
use Sentry\Laravel\Integration;
use Symfony\Component\HttpKernel\Exception\AccessDeniedHttpException;
use Illuminate\Http\Request;

return Application::configure(basePath: dirname(__DIR__))
    ->withRouting(
        commands: __DIR__.'/../routes/console.php',
        health: '/up',
    )
    ->withMiddleware(function (Middleware $middleware): void {
        $middleware->api([
            AddCorrelationId::class,
        ]);
        $middleware->append(SecurityHeaders::class);
        $middleware->alias([
            'role' => RoleMiddleware::class,
        ]);
    })
    ->withExceptions(function (Exceptions $exceptions): void {
        $exceptions->dontFlash([
            'current_password',
            'password',
            'password_confirmation',
            'token',
            'access_token',
            'refresh_token',
            'tax_id',
            'iban',
            'cvv',
            'card_number',
        ]);

        Integration::handles($exceptions);

        $exceptions->render(function (ValidationException $e, Request $request) {
            if ($request->expectsJson() || $request->is('api/*')) {
                return response()->json([
                    'message' => 'Validation error.',
                    'errors' => $e->errors(),
                ], 422);
            }
            return null;
        });

        $exceptions->render(function (ThrottleRequestsException $e, Request $request) {
            if ($request->expectsJson() || $request->is('api/*')) {

                return response()->json([
                    'message' => 'Too many requests. Try later.',
                    'code' => 'RATE_LIMIT_EXCEEDED',
                ], 429)
                    ->withHeaders([
                        'X-RateLimit-Limit' => $e->getHeaders()['X-RateLimit-Limit'] ?? '',
                        'X-RateLimit-Remaining' => $e->getHeaders()['X-RateLimit-Remaining'] ?? '',
                        'Retry-After' => $e->getHeaders()['Retry-After'] ?? '',
                    ]);
            }

            return null;
        });

        $exceptions->render(function (AuthenticationException $e, Request $request) {
            if ($request->expectsJson() || $request->is('api/*')) {
                return response()->json([
                    'message' => 'Authentication error.',
                ], 401);
            }

            return null;
        });

        $exceptions->render(function (AccessDeniedHttpException $e, Request $request) {
            if ($request->expectsJson() || $request->is('api/*')) {
                $user = $request->user();

                Log::warning('authorization.denied', [
                    'user_id' => $user?->id,
                    'role' => $user?->role,
                    'method' => $request->method(),
                    'path' => $request->path(),
                    'correlation_id' => $request->attributes->get('correlation_id'),
                ]);

                return response()->json([
                    'message' => 'Access denied.'
                ], 403);
            }

            return null;
        });

        $exceptions->render(function (ModelNotFoundException $e, Request $request) {
            if ($request->expectsJson() || $request->is('api/*')) {
                return response()->json([
                    'message' => 'Not found.'
                ], 404);
            }
            return null;
        });

        $exceptions->render(function (Throwable $e, Request $request) {
            if ($request->expectsJson() || $request->is('api/*')) {
                Log::error('Unexpected error during request', [
                    'error_message' => $e->getMessage(),
                    'trace' => $e->getTraceAsString(),
                ]);

                return response()->json([
                    'message' => 'Internal server error. Please try again later.'
                ], 500);
            }

            return null;
        });
    })->create();
