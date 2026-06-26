<?php

declare(strict_types=1);

namespace App\Modules\Quotes\Application\DTOs;

readonly class SpreadDTO
{
    public function __construct(
        public string $buyExchange,
        public string $sellExchange,
        public string $buyPrice,
        public string $sellPrice,
        public string $spreadPercent,
    ) {}
}
