<?php

namespace Lib\Routing;

class Route
{
    public function __construct(
        public readonly string $path,
        public readonly array $methods,
        public readonly string $controller,
        public readonly string $action,
    ){}
}