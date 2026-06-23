<?php

declare(strict_types=1);

namespace App\Modules\Quotes\Infrastructure\Repositories;

use Illuminate\Support\Collection;

interface QuoteRepository
{
    public function getRates(string $pair): Collection;

}
