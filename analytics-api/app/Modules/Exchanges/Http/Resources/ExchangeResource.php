<?php

declare(strict_types=1);

namespace App\Modules\Exchanges\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;

class ExchangeResource extends JsonResource
{
    public function toArray($request): array
    {
        return [
            'id' => $this->id,
            'name' => $this->name,
            'enabled' => $this->is_enabled,
        ];
    }
}
