<?php

declare(strict_types=1);

namespace App\Modules\Exchanges\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class UpdateExchangeRequest extends FormRequest
{
    public function rules(): array
    {
        return [
            'is_enabled' => ['required', 'boolean'],
        ];
    }
}
