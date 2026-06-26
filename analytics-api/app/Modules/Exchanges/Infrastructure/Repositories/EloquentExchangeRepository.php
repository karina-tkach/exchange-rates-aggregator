<?php

declare(strict_types=1);

namespace App\Modules\Exchanges\Infrastructure\Repositories;

use App\Modules\Exchanges\Domain\Exchange;

class EloquentExchangeRepository implements ExchangeRepository
{
    public function getEnabledExchangesNames(): array
    {
        return Exchange::query()
            ->where('is_enabled', true)
            ->pluck('name')
            ->all();
    }
}
