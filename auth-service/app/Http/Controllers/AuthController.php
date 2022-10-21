<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

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

        $payload = [
            'name' => 'log',
            'data' => $log
        ];

        $conf = new \RdKafka\Conf();
        $conf->set('log_level', config('kafka.debug'));
        $conf->set('debug', 'all');
        $rk = new \RdKafka\Producer($conf);
        $rk->addBrokers(config('kafka.brokers'));
        $topic = $rk->newTopic("logger");
        $res = $topic->produce(RD_KAFKA_PARTITION_UA, 0, json_encode($payload));
        $rk->flush(1000);
    }
}
