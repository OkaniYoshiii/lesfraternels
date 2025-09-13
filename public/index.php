<?php

declare(strict_types=1);

use Lib\Framework\Controller;
use Lib\Http\Request;
use Lib\Http\Response;
use Lib\Routing\Route;
use Lib\Templating\Renderer;
use Twig\Environment;
use Twig\Loader\FilesystemLoader;

define('ROOT_DIR', realpath(__DIR__ . '/..'));
define('TEMPLATE_DIR', realpath(ROOT_DIR . '/templates'));
define('SRC_DIR', realpath(ROOT_DIR . '/src'));

require_once ROOT_DIR . '/vendor/autoload.php';

$twig = (function(): Environment {
    $loader = new FilesystemLoader(TEMPLATE_DIR);
    
    return new Environment($loader, []);
})();

$request = (function(): Request {
    $uri = $_SERVER['REQUEST_URI'];
    $method = $_SERVER['REQUEST_METHOD'];

    return new Request($uri, $method);
})();

$routes = require_once ROOT_DIR . '/config/routes/routes.php';

$route = (function() use ($routes, $request): ?Route {
    foreach($routes as $route) {
        if($route->path === $request->path && in_array($request->method, $route->methods)) {
            return $route;
        }
    }

    return null;
})();

$controller = (function() use ($twig, $route): ?object {
    if(!($route instanceof Route)) {
        return null;
    }

    $fqcn = $route->controller;

    if(!is_a($fqcn, Controller::class, true)) {
        return null;
    }

    return new $fqcn($twig);
})();

$action = (function() use ($controller, $route): ?string {
    if($controller === null || $route === null) {
        return null;
    }

    if(!method_exists($controller, $route->action)) {
        return null;
    }

    return $route->action;
})();

$response = (function() use ($controller, $action, $request): Response {
    if($controller === null || $action === null) {
        return Response::html(
            content: "",
            statusCode: 404,
        );
    }

    $response = $controller->$action($request);

    if(!($response instanceof Response)) {
        $message = sprintf('Each controller action must accept a %s object and return a %s object', Request::class, Response::class);

        throw new Exception($message);
    }

    return $response;
})();

echo $response->content;