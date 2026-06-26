<?php

declare(strict_types=1);

namespace Database\Seeders;

use App\Modules\Pairs\Domain\Pair;
use Illuminate\Database\Seeder;

class PairSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run(): void
    {
        Pair::query()->create(['base' => 'BTC', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'ETH', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'BNB', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'SOL', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'AVAX', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'MATIC', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'DOGE', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'XRP', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'LTC', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'ADA', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'DOT', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'SHIB', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'LINK', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'UNI', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'AAVE', 'quote' => 'USDT']);
        Pair::query()->create(['base' => 'SUSHI', 'quote' => 'USDT']);
    }
}
