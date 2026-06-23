<?php

declare(strict_types=1);

namespace App\Modules\Quotes\Domain;

use App\Modules\Pairs\Domain\Pair;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class Quote extends Model
{
    public $incrementing = false;
    public $timestamps = false;
    public $primaryKey = null;

    protected $fillable = [
        'time',
        'pair_id',
        'price',
        'bid',
        'ask',
        'source'
    ];

    protected $casts = [
        'time' => 'datetime',
        'price' => 'decimal:8',
        'bid'   => 'decimal:8',
        'ask'   => 'decimal:8',
    ];

    public function pair(): BelongsTo
    {
        return $this->belongsTo(Pair::class);
    }
}
