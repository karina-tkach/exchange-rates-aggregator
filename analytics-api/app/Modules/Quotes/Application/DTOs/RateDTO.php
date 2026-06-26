<?php

declare(strict_types=1);

namespace App\Modules\Quotes\Application\DTOs;

readonly class RateDTO
{
    public function __construct(
        public string $source,
        public string $price,
        public string $bid,
        public string $ask,
        public string $updatedAt,
    ) {}
}
