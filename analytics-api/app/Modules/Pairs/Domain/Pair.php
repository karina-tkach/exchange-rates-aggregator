<?php

declare(strict_types=1);

namespace App\Modules\Pairs\Domain;

use Illuminate\Database\Eloquent\Model;

class Pair extends Model
{
    protected $fillable = ['base', 'quote'];
}
