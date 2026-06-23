<?php

declare(strict_types=1);

namespace App\Modules\Quotes\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class RateResource extends JsonResource
{
    public function toArray(Request $request): array
    {
        return [
            'exchange' => $this->source,
            'price' => $this->price,
        ];
    }
}
