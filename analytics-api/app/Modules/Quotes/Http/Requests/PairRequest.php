<?php

declare(strict_types=1);

namespace App\Modules\Quotes\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class PairRequest extends FormRequest
{
    public function authorize(): bool
    {
        return true;
    }

    public function rules(): array
    {
        return [
            'pair' => [
                'required',
                'regex:/^[A-Z0-9]+-[A-Z0-9]+$/'
            ]
        ];
    }

    public function messages(): array
    {
        return [
            'pair.required' => 'Pair is required.',
            'pair.regex' => 'Pair must be in format BASE-QUOTE (BTC-USDT).',
        ];
    }
}
