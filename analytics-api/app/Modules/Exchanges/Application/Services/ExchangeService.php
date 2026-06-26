<?php

declare(strict_types=1);

namespace App\Modules\Exchanges\Application\Services;

use App\Modules\Exchanges\Infrastructure\Repositories\ExchangeRepository;
use Illuminate\Support\Facades\Redis;

readonly class ExchangeService
{
    public function __construct(
        private ExchangeRepository $exchangeRepository
    ) {}

    public function getEnabledExchangesNames(): array
    {
        $value = Redis::get('enabled_exchanges');

        if ($value !== null) {
            return json_decode($value, true);
        }

        $value = $this->exchangeRepository->getEnabledExchangesNames();

        Redis::setex('enabled_exchanges', 30, json_encode($value));

        return $value;
    }
}
