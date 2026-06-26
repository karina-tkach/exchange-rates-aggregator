<?php

declare(strict_types=1);

namespace App\Modules\Pairs\Http\Controllers;

use App\Http\Controllers\Controller;
use App\Modules\Pairs\Application\Services\PairService;
use Illuminate\Http\JsonResponse;

class PairController extends Controller
{
    public function __construct(
        private readonly PairService $pairService,
    ) {}

    public function index(): JsonResponse
    {
        return response()->json([
            'data' => $this->pairService->getPairs(),
        ]);
    }
}
