<?php

namespace Lib\Framework\Twig;

use Twig\TwigFunction;

class TwigExtension
{
    public function __construct(
        private readonly array $routes,
    ){}

    public function functions(): array
    {
        $path = new TwigFunction('path', function(string $route) {
            if(!isset($this->routes[$route])) {
                $message = sprintf('Route "%s" does not exists', $route);
                throw new \Exception($message);
            }

            return $this->routes[$route]->path;
        });

        return [
            $path,
        ];
    }
}
