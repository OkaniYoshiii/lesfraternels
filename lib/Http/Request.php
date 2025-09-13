<?php

namespace Lib\Http;

class Request
{
    public string $path {
        get => parse_url($this->uri, PHP_URL_PATH);
    }

    public function __construct(
        public readonly string $uri,
        public readonly string $method,
    ){}
}