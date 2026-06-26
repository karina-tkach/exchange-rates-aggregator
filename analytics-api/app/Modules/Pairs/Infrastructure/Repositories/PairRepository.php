<?php

declare(strict_types=1);

namespace App\Modules\Pairs\Infrastructure\Repositories;

interface PairRepository
{
    public function getPairs(): array;
}
