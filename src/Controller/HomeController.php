<?php

namespace App\Controller;

use Lib\Framework\Controller;
use Lib\Http\Request;
use Lib\Http\Response;

class HomeController extends Controller
{
    public function index(Request $request): Response
    {
        $content = $this->render('home/index.html.twig', []);
        return Response::html($content);
    }
}