<?php

declare(strict_types=1);


namespace App\Modules\Quotes\Http\Controllers;

use App\Http\Controllers\Controller;
use App\Modules\Quotes\Application\Services\QuoteService;
use App\Modules\Quotes\Http\Requests\PairRequest;
use App\Modules\Quotes\Http\Resources\RateResource;
use App\Modules\Quotes\Http\Resources\SpreadResource;
use Illuminate\Http\JsonResponse;

class RateController extends Controller
{
    public function __construct(
        private readonly QuoteService $quoteService,
    ) {}
    public function index(PairRequest $request): JsonResponse
    {
        return RateResource::collection(
            $this->quoteService->getRates($request->validated('pair'))
        )->response();
    }


    public function spread(PairRequest $request): JsonResponse
    {
        $spread = $this->quoteService->calculateSpread($request->validated('pair'));

        if ($spread === null) {
            return response()->json([
                'message' => 'Not enough exchanges or no arbitrage opportunity',
            ], 422);
        }

        return (new SpreadResource($spread))->response()->setStatusCode(200);
    }
}
