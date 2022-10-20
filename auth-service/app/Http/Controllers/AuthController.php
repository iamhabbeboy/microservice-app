<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Junges\Kafka\Facades\Kafka;
use Junges\Kafka\Message\Message;

class AuthController extends Controller
{
    public function __invoke(Request $request)
    {
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
        $this->logger($email);
    
        return response()->json(['message' => 'successful', 'error' => false, 'data' => $user ], 202);
    }

    private function logger(string $data)
    {
        $log = $data . " is logged in";
        // log data here and publish to queue
        $message = new Message(
            headers: ['header-key' => 'header-value'],
            body: ['name' => 'log', 'data' => $log],
            key: 'auth'  
        );

        Kafka::publishOn('logger')->withMessage($message);

    }
}
