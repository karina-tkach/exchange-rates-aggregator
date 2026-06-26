<?php

declare(strict_types=1);

namespace App\Modules\Pairs\Application\Services;

use App\Modules\Pairs\Infrastructure\Repositories\PairRepository;

readonly class PairService
{
    public function __construct(
        private PairRepository $pairRepository
    ) {}

    public function getPairs(): array {
        return $this->pairRepository->getPairs();
    }
}
