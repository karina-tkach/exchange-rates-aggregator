<?php

declare(strict_types=1);

namespace App\Modules\Exchanges\Http\Controllers;

use App\Http\Controllers\Controller;
use App\Modules\Exchanges\Application\Services\ExchangeService;
use App\Modules\Exchanges\Http\Requests\UpdateExchangeRequest;
use App\Modules\Exchanges\Http\Resources\ExchangeResource;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;

class ExchangeController extends Controller
{
    public function __construct(
        private readonly ExchangeService $exchangeService,
    ) {}

    public function index(Request $request)
    {
        $perPage = min($request->integer('per_page', 10), 100);
        $page = $request->integer('page', 1);

        return ExchangeResource::collection(
            $this->exchangeService->paginate($page, $perPage)
        );
    }

    public function update(UpdateExchangeRequest $request, int $id): JsonResponse {
        return (new ExchangeResource($this->exchangeService->update($id, $request->validated())))
            ->response()->setStatusCode(200);
    }

    public function active(): JsonResponse
    {
        return response()->json([
            'data' => $this->exchangeService->getEnabledExchangesNames(),
        ]);
    }

}
