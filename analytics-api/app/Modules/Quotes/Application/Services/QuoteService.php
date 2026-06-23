<?php

declare(strict_types=1);

namespace App\Modules\Quotes\Application\Services;

use App\Modules\Quotes\Application\DTOs\RateDTO;
use App\Modules\Quotes\Application\DTOs\SpreadDTO;
use App\Modules\Quotes\Infrastructure\Repositories\OhlcRepository;
use App\Modules\Quotes\Infrastructure\Repositories\QuoteRepository;
use Illuminate\Support\Collection;
use Illuminate\Support\Facades\Redis;

readonly class QuoteService
{
    public function __construct(
        private QuoteRepository $quoteRepository,
        private OhlcRepository $ohlcRepository,
    ) {}

    public function getRates(string $pair): Collection
    {
        $cached = Redis::get("rates:$pair");

        if ($cached !== null) {
            return $this->fromRedis($cached);
        }

        return $this->quoteRepository
            ->getRates($pair)
            ->map(
                fn ($quote) => new RateDto(
                    source: $quote->source,
                    price: (string) $quote->price,
                    bid: (string) $quote->bid,
                    ask: (string) $quote->ask,
                    updatedAt: (string) $quote->time,
                )
            );
    }

    public function calculate(Collection $rates): ?SpreadDto {
        if ($rates->count() < 2) {
            return null;
        }

        $buy = null;
        $sell = null;

        foreach ($rates as $rate) {
            if ($buy === null || bccomp($rate->ask, $buy->ask, 8) < 0)
            {
                $buy = $rate;
            }

            if ($sell === null || bccomp($rate->bid, $sell->bid, 8) > 0) {
                $sell = $rate;
            }
        }

        if ($buy->source === $sell->source)
        {
            return null;
        }


        $spread = bcmul(
            bcdiv(bcsub($sell->bid, $buy->ask, 8), $buy->ask, 8),
            '100',
            4
        );


        return new SpreadDto(
            buyExchange: $buy->source,
            sellExchange: $sell->source,
            buyPrice: $buy->ask,
            sellPrice: $sell->bid,
            spreadPercent: $spread
        );
    }

    private function fromRedis(string $json): Collection
    {
        $data = json_decode($json, true);

        return collect($data)
            ->map(
                fn (array $rate, string $exchange) =>
                new RateDto(
                    source: $exchange,
                    price: $rate['price'],
                    bid: $rate['bid'],
                    ask: $rate['ask'],
                    updatedAt: $rate['updated_at'],
                )
            )->values();
    }

    public function getCandles(string $pair, string $period, string $timeframe): Collection {
        return $this->ohlcRepository->getCandles($pair, $period, $timeframe);
    }
}
