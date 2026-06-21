<?php

declare(strict_types=1);

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration {
    public function up(): void
    {
        Schema::table('quotes', function (Blueprint $table) {
            $table->decimal('price', 20, 8)->change();

            $table->decimal('bid', 20, 8)->change();

            $table->decimal('ask', 20, 8)->change();
        });
    }


    public function down(): void
    {
        Schema::table('quotes', function(Blueprint $table) {
            $table->double('price')->change();
            $table->double('bid')->change();
            $table->double('ask')->change();

        });
    }
};
