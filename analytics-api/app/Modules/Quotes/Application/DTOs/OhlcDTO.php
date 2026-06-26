<?php

declare(strict_types=1);

namespace App\Modules\Quotes\Application\DTOs;

final readonly class OhlcDTO
{
    public function __construct(
        public string $time,
        public string $open,
        public string $high,
        public string $low,
        public string $close,
    ) {}
}
