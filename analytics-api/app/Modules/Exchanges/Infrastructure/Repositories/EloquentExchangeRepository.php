<?php

declare(strict_types=1);

namespace App\Modules\Exchanges\Infrastructure\Repositories;

use App\Modules\Exchanges\Domain\Exchange;
use Illuminate\Contracts\Pagination\LengthAwarePaginator;

class EloquentExchangeRepository implements ExchangeRepository
{
    public function getEnabledExchangesNames(): array
    {
        return Exchange::query()
            ->where('is_enabled', true)
            ->pluck('name')
            ->all();
    }

    public function paginate(int $page, int $perPage = 10): LengthAwarePaginator
    {
        return Exchange::query()
            ->orderBy('id')
            ->paginate(perPage: $perPage,
                page: $page);
    }

    public function update(int $id, array $data): Exchange
    {
        $exchange = Exchange::query()->findOrFail($id);
        $exchange->update($data);
        return $exchange;
    }
}
