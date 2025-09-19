<?php

namespace Config;

function location(string $directory): string
{
    $root = __DIR__ . '/..';

    return match($directory) {
        'root' => realpath($root),
        'src' => realpath($root . '/src'),
        'templates' => realpath($root . '/templates'),
    };
}