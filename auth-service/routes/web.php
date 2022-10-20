<?php

use Illuminate\Support\Facades\Route;
use Junges\Kafka\Facades\Kafka;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', function () {
    $message = new Message(
        headers: ['header-key' => 'header-value'],
        body: ['key' => 'value'],
        key: 'kafka key here'  
    );
    Kafka::publishOn('topic')->withMessage($message);
    return view('welcome');
});
