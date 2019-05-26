
<?php

$urls  = [];
$content = file_get_contents("url");

foreach (explode("\n", $content) as $value) {
    $urls[md5($value)] = $value;
}

$str = "";
foreach (array_values($urls) as $url) {
    
    $str .= $url . PHP_EOL;

}

file_put_contents("url_new", $str );
