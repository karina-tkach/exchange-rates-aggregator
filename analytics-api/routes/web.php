<?php

use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('welcome');
});

Route::get('/go-test', function () {
    //$response = Http::get('http://collector:8080/test');

    return [
        'from_go' => "sth"
    ];
});
