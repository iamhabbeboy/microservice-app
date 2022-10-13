<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/
Route::post('/', function(Request $request) {
    $users = collect([
        [
            'email' => 'admin@gmail.com',
            'password' => 'verify'
        ],
        [
            'email' => 'user@gmail.com',
            'password' => 'admin'
        ]
    ]);
    $payload = $request->json()->all();
    $email = $payload['email'];
    $pass = $payload['password'];
    
    $user = $users->where('email', $email)->where('password', $pass);

    if(count($user) === 0) {
        return response()->json(['message' => 'failed to authenticate', 'error' => true, 'data' => [] ], 401);
    }

    return response()->json(['message' => 'successful', 'error' => false, 'data' => $user ], 202);
});

Route::get('/health-check', function() {
    return response()->json(['status' => 'All is well here from auth']);
});

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});

