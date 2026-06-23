<?php

declare(strict_types=1);

namespace App\Modules\Quotes\Infrastructure\Repositories;

use App\Modules\Quotes\Domain\Quote;
use Illuminate\Support\Collection;

class EloquentQuoteRepository implements QuoteRepository
{

    public function getRates(string $pair): Collection
    {
        return Quote::query()
            ->select('source', 'price', 'bid', 'ask', 'time')
            ->distinct(['source'])
            ->join(
                'pairs',
                'pairs.id',
                '=',
                'quotes.pair_id'
            )
            ->whereRaw(
                "CONCAT(pairs.base, '-', pairs.quote) = ?",
                [$pair]
            )
            ->orderBy('source')
            ->orderByDesc('time')
            ->get();
    }
}
