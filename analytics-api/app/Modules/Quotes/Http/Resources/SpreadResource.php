<?php

declare(strict_types=1);

namespace App\Modules\Quotes\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class SpreadResource extends JsonResource
{
    public function toArray(Request $request): array
    {
        return [
            'buy_exchange' => $this->buyExchange,
            'sell_exchange' => $this->sellExchange,
            'buy_price' => $this->buyPrice,
            'sell_price' => $this->sellPrice,
            'spread_percent' => $this->spreadPercent,
        ];
    }
}
