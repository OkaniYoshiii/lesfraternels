<?php

namespace App\Controller;

use Lib\Framework\Controller;
use Lib\Http\Request;
use Lib\Http\Response;

class ModsController extends Controller
{
    public function index(Request $request): Response
    {
        return Response::html("");
    }
}