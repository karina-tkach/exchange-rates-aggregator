<?php

declare(strict_types=1);

namespace App\Modules\Quotes\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class OhlcResource extends JsonResource
{
    public function toArray(Request $request): array
    {
        return [
            'time' => $this->time,
            'open' => $this->open,
            'high' => $this->high,
            'low' => $this->low,
            'close' => $this->close,
        ];
    }
}
