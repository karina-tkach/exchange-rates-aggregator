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
        $pair = $request->validated('pair');

        $rates = $this->quoteService->getRates($pair);
        $spread = $this->quoteService->calculate($rates);

        return (new SpreadResource($spread))->response()->setStatusCode(200);
    }
}
