<?php

declare(strict_types=1);

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Log;
use Illuminate\Support\Str;
use Sentry\State\Scope;
use Symfony\Component\HttpFoundation\Response;

use function Sentry\configureScope;

class AddCorrelationId
{
    /**
     * Handle an incoming request.
     *
     * @param  Closure(Request): (Response)  $next
     */
    public function handle(Request $request, Closure $next): Response
    {
        $correlationId = $request->header('X-Correlation-Id');

        if (!is_string($correlationId) || empty($correlationId) || !Str::isUuid($correlationId)) {
            $correlationId = Str::uuid()->toString();
        }

        $request->attributes->set('correlation_id', $correlationId);

        Log::withContext([
            'correlation_id' => $correlationId,
            'endpoint' => $request->method() . ' ' . $request->path(),
            'module' => $this->resolveModule($request->path()),
        ]);

        if (app()->bound('sentry')) {
            configureScope(function (Scope $scope) use ($correlationId, $request): void {
                $scope->setTag('correlation_id', $correlationId);
                $scope->setExtra('correlation_id', $correlationId);

                $scope->setTag('endpoint', $request->method() . ' ' . $request->path());

                if ($request->user()) {
                    $scope->setUser([
                        'id' => (string) $request->user()->id,
                    ]);
                }
            });
        }

        $response = $next($request);

        $response->headers->set('X-Correlation-Id', $correlationId);

        return $response;
    }

    private function resolveModule(string $path): string
    {
        $trimmed = preg_replace('#^api/v1/#', '', $path);

        $parts = explode('/', $trimmed);

        return $parts[0];
    }
}
