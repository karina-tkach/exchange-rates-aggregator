<?php

declare(strict_types=1);

namespace App\Modules\Exchanges\Http\Controllers;

use App\Http\Controllers\Controller;
use App\Modules\Exchanges\Application\Services\ExchangeService;
use Illuminate\Http\JsonResponse;

class ExchangeController extends Controller
{
    public function __construct(
        private readonly ExchangeService $exchangeService,
    ) {}

    public function index(): JsonResponse
    {
        return response()->json([
            'data' => $this->exchangeService->getEnabledExchangesNames(),
        ]);
    }

}
