<?php

use Lib\Routing\Route;

return [
    'home' => new Route(path: '/', methods: ['GET', 'POST'], controller: 'App\\Controller\\HomeController', action: 'index'),
];