<?php

declare(strict_types=1);

use Lib\Framework\Controller;
use Lib\Framework\Twig\TwigExtension;
use Lib\Http\Request;
use Lib\Http\Response;
use Lib\Routing\Route;
use Twig\Environment;
use Twig\Loader\FilesystemLoader;

require_once __DIR__ . '/../vendor/autoload.php';

$request = (function(): Request {
    $uri = $_SERVER['REQUEST_URI'];
    $method = $_SERVER['REQUEST_METHOD'];

    return new Request($uri, $method);
})();

$routes = Config\routes();

$route = (function() use ($routes, $request): ?Route {
    foreach($routes as $route) {
        if($route->path === $request->path && in_array($request->method, $route->methods)) {
            return $route;
        }
    }

    return null;
})();

$twig = (function() use ($routes): Environment {
    $loader = new FilesystemLoader(Config\location('templates'));
    
    $options = [
        'strict_variables' => true,
        'optimizations' => 0,
    ];

    $environment = new Environment($loader, $options);
    $extension = new TwigExtension($routes);

    foreach($extension->functions() as $function) {
        $environment->addFunction($function);
    }

    return $environment;
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

http_response_code($response->statusCode);
echo $response->content;