<?php

$client = new Swoole\Client(SWOOLE_SOCK_TCP);
if (!$client->connect('127.0.0.1', 8201, -1)) {
    exit("connect failed. Error: {$client->errCode}\n");
}

$context = "princelovechina";

// 粘包处理
$len = pack("n", strlen($context));
var_dump($len);
$data = $len . $context;
var_dump($data);

$client->send($data);

echo $client->recv();

$client->close();