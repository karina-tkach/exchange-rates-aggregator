<?php

declare(strict_types=1);

namespace App\Modules\Pairs\Infrastructure\Repositories;

use App\Modules\Pairs\Domain\Pair;

class EloquentPairRepository implements PairRepository
{
    public function getPairs(): array
    {
        return Pair::query()
            ->select('base', 'quote')
            ->get()
            ->map(fn ($p) => "{$p->base}-{$p->quote}")
            ->toArray();
    }
}
