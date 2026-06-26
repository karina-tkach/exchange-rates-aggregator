<?php

declare(strict_types=1);

use Illuminate\Database\Migrations\Migration;
use Illuminate\Support\Facades\DB;

return new class extends Migration
{
    public function up(): void
    {
        DB::statement("
            SELECT add_retention_policy('quotes', INTERVAL '7 days');
        ");
    }

    public function down(): void
    {
        DB::statement("
            SELECT remove_retention_policy('quotes');
        ");
    }
};
