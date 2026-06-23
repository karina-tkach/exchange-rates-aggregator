<?php

namespace App\Modules\Quotes\Http\Controllers;

use App\Http\Controllers\Controller;
use App\Modules\Quotes\Application\Services\QuoteService;
use App\Modules\Quotes\Http\Requests\OhlcRequest;
use App\Modules\Quotes\Http\Resources\OhlcResource;

class OhlcController extends Controller
{
    public function __construct(
        private readonly QuoteService $quoteService
    ) {}

    public function index(OhlcRequest $request)
    {
        $validated = $request->validated();

        return OhlcResource::collection(
            $this->quoteService->getCandles(
                $validated['pair'],
                $validated['period'],
                $validated['timeframe'],
            )
        );
    }
}
