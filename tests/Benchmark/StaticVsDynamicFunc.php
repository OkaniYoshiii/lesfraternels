<?php

namespace App\Tests\Benchmark;

use PhpBench\Attributes;

const revolutions = 100;
const iterations = 10;

class StaticVsDynamicFunc
{
    #[Attributes\Revs(revolutions)]
    #[Attributes\Iterations(iterations)]
    public function benchStatic()
    {
        staticFunc();
    }

    #[Attributes\Revs(revolutions)]
    #[Attributes\Iterations(iterations)]
    public function benchDynamic()
    {
        dynamicFunc();
    }
}

function dynamicFunc(): string
{
    return __DIR__ . '/../';
}

function staticFunc(): string
{
    static $result = __DIR__ . '/../';

    return $result;
}
