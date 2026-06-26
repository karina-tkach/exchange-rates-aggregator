<?php

namespace App\Modules\Quotes\Infrastructure\Repositories;

use Illuminate\Support\Collection;

interface OhlcRepository
{
    public function getCandles(string $pair, string $period, string $timeframe, array $exchanges): Collection;
}
