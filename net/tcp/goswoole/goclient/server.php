<?php
//创建Server对象，监听 127.0.0.1:9501 端口
$server = new Swoole\Server('127.0.0.1', 9501);
echo "start tcp server: 127.0.0.1:9501 \n";

$server->set([
  'open_length_check'     => true,
  'package_max_length'    => 1 * 1024 * 1024 ,
  'package_length_type'   => 'n', //see php pack()
  'package_length_offset' => 0,
  'package_body_offset'   => 2
]);

//监听连接进入事件
$server->on('Connect', function ($server, $fd) {
    echo "Client: Connect.\n";
});

//监听数据接收事件
$server->on('Receive', function ($server, $fd, $reactor_id, $data) {
    var_dump("接收到的数据: " . $data);
    var_dump("fd: " . $fd);
    $server->send($fd, "Server: {$data}");
});

//监听连接关闭事件
$server->on('Close', function ($server, $fd) {
    echo "Client: Close.\n";
});

//启动服务器
$server->start();

/** 客户端发送消息到该服务端，并发高的情况必现报错：
    [2021-05-22 10:09:36 #29865.8]  WARNING swSignal_callback (ERRNO 707): Unable to find callback function for signal Broken pipe: 13

    原因：是后端服务关闭了，你还在用这个socket就会触发这个
    复现：模拟并发
    监听信号：
        1. https://www.php.net/manual/zh/function.pcntl-signal.php
        2. swoole 也有 Process::signal可以使用
        3. 或者改用框架吧 框架把这些情况都处理了https://github.com/lizhichao/one
*/