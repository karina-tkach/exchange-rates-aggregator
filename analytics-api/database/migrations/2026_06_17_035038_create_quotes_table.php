<?php

declare(strict_types=1);

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    public $withinTransaction = false;
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        DB::statement("CREATE EXTENSION IF NOT EXISTS timescaledb");

        Schema::create('quotes', function (Blueprint $table) {
            $table->timestampTz('time');
            $table->unsignedBigInteger('pair_id');
            $table->double('price');
            $table->double('bid');
            $table->double('ask');
            $table->string('source', 50);

            $table->index(['pair_id', 'time', 'source'], 'quotes_pair_time_source_idx');
        });

        DB::statement("
            SELECT create_hypertable('quotes', 'time', if_not_exists => TRUE);
        ");
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('quotes');
    }
};
