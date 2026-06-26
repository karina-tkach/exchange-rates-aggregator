<?php

declare(strict_types=1);

namespace App\Modules\Exchanges\Infrastructure\Repositories;

use App\Modules\Exchanges\Domain\Exchange;
use Illuminate\Contracts\Pagination\LengthAwarePaginator;

interface ExchangeRepository
{
    public function getEnabledExchangesNames(): array;
    public function paginate(int $page, int $perPage = 10): LengthAwarePaginator;
    public function update(int $id, array $data): Exchange;

}
