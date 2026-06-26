<?php

declare(strict_types=1);

namespace Database\Seeders;

use App\Modules\Exchanges\Domain\Exchange;
use Illuminate\Database\Seeder;

class ExchangeSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        Exchange::query()->create(['name' => 'Binance', 'is_enabled' => true]);
        Exchange::query()->create(['name' => 'Coinbase', 'is_enabled' => true]);
        Exchange::query()->create(['name' => 'Kraken', 'is_enabled' => true]);
    }
}
