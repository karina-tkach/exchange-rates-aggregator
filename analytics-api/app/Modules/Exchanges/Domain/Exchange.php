<?php

namespace App\Modules\Exchanges\Domain;

use Illuminate\Database\Eloquent\Model;

class Exchange extends Model
{
    protected $fillable = ['name', 'is_enabled'];
}
