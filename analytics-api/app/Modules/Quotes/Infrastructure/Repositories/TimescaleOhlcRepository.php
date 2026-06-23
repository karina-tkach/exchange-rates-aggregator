<?php

declare(strict_types=1);

namespace App\Modules\Quotes\Infrastructure\Repositories;

use App\Modules\Quotes\Application\DTOs\OhlcDTO;
use Illuminate\Support\Collection;
use Illuminate\Support\Facades\DB;

class TimescaleOhlcRepository implements OhlcRepository
{
    public function getCandles(string $pair, string $period, string $timeframe): Collection
    {

        $bucket = match ($timeframe) {
            '1m' => '1 minute',
            '5m' => '5 minutes',
            '1h' => '1 hour',
            '4h' => '4 hours',
            '1d' => '1 day',
        };

        $from = match ($period) {
            '1h' => now()->subHour(),
            '24h' => now()->subDay(),
            '7d' => now()->subDays(7),
        };

        $sql = "
            SELECT
                time_bucket('$bucket', q.time) as bucket,
                first(q.price, q.time) as open,
                max(q.price) as high,
                min(q.price) as low,
                last(q.price, q.time) as close
            FROM quotes q
            JOIN pairs p
                ON p.id = q.pair_id
            WHERE CONCAT(
                p.base,
                '-',
                p.quote
            ) = ?
            AND q.time >= ?
            GROUP BY bucket
            ORDER BY bucket
        ";

        $rows = DB::select(
            $sql,
            [$pair, $from]
        );

        return collect($rows)
            ->map(
                fn ($row) => new OhlcDto(
                    time: $row->bucket,
                    open: (string)$row->open,
                    high: (string)$row->high,
                    low: (string)$row->low,
                    close: (string)$row->close,
                )
            );
    }
}
