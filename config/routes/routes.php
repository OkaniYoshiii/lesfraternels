<?php

use Lib\Routing\Route;

return [
    'home' => new Route(path: '/', methods: ['GET', 'POST'], controller: 'App\\Controller\\HomeController', action: 'index'),
    'mods' => new Route(path: '/mods', methods: ['GET'], controller: 'App\\Controller\\ModsController', action: 'index'),
    'legal_mentions' => new Route(path: '/mentions-legales', methods: ['GET'], controller: 'App\\Controller\\LegalMentionsController', action: 'index'),
    'contact' => new Route(path: '/contact', methods: ['GET'], controller: 'App\\Controller\\ContactController', action: 'index'),
];