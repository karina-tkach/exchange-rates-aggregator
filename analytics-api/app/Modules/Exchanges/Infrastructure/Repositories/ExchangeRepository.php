<?php

declare(strict_types=1);

namespace App\Modules\Exchanges\Infrastructure\Repositories;

interface ExchangeRepository
{
    public function getEnabledExchangesNames(): array;

}
