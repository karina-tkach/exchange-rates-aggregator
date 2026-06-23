<?php

declare(strict_types=1);

namespace App\Modules\Quotes\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class OhlcRequest extends FormRequest
{
    public function authorize(): bool
    {
        return true;
    }

    public function rules(): array
    {
        return [
            'pair' => ['required', 'string','regex:/^[A-Z0-9]+-[A-Z0-9]+$/'],
            'period' => ['required', 'string', 'in:1h,24h,7d'],
            'timeframe' => ['required', 'string', 'in:1m,5m,1h,4h,1d'],
        ];
    }

    public function after(): array
    {
        return [
            function () {
                $period = $this->input('period');
                $timeframe = $this->input('timeframe');

                $allowed = [
                    '1h' => ['1m', '5m'],
                    '24h' => ['1m', '5m', '1h', '4h'],
                    '7d' => ['1m', '5m', '1h', '4h', '1d'],
                ];

                if (!in_array($timeframe, $allowed[$period] ?? [], true)
                ) {
                    $this->validator->errors()->add(
                        'timeframe',
                        'Invalid timeframe for selected period.'
                    );
                }
            }
        ];
    }

    public function messages(): array
    {
        return [
            'pair.required' => 'Pair is required.',
            'pair.regex' => 'Pair must be in format BASE-QUOTE (BTC-USDT).',

            'period.required' => 'Period is required.',
            'period.in' => 'Period must be 1h, 24h or 7d.',
            'timeframe.required' => 'Timeframe is required.',
            'timeframe.in' => 'Timeframe must be 1m, 5m, 1h, 4h or 1d.',
        ];
    }
}
