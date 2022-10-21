<?php
    return [
        /*
        | Your kafka brokers url.
        */
        'brokers' => env('KAFKA_BROKERS', 'kafka:9092'),

        /*
        | Kafka consumers belonging to the same consumer group share a group id.
        | The consumers in a group then divides the topic partitions as fairly amongst themselves as possible by
        | establishing that each partition is only consumed by a single consumer from the group.
        | This config defines the consumer group id you want to use for your project.
        */
        'consumer_group_id' => env('KAFKA_CONSUMER_GROUP_ID', 'logger-group'),

        /*
        | Choose if debug is enabled or not.
        */
        'debug' => env('KAFKA_DEBUG', true),
    ];