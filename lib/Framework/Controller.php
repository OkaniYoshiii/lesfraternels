<?php

namespace Lib\Framework;

use Lib\Templating\Renderer;
use Twig\Environment;

class Controller
{
    public function __construct(
        private readonly Environment $twig,
    ){}

    public function render(string $template, array $data): string
    {
        return $this->twig->render($template, $data);
    }
}