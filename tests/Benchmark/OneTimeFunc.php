<?php

namespace App\Tests\Benchmark;

use Exception;
use PhpBench\Attributes;

const revolutions = 1;
const iterations = 100;

class OneTimeFunc
{
    #[Attributes\Revs(revolutions)]
    #[Attributes\Iterations(iterations)]
    public function benchIIFE()
    {
        $test = (function(): string {
            return 'test';
        })();

        $test === '';
    }

    #[Attributes\Revs(revolutions)]
    #[Attributes\Iterations(iterations)]
    public function benchDeclared()
    {
        $test = test();

        $test === '';
    }
    
    #[Attributes\Revs(revolutions)]
    #[Attributes\Iterations(iterations)]
    public function benchAnonymous()
    {
        $func = function() {
            static $count = 0;

            if($count > 0) {
                throw new \Exception('Patate');
            }

            $count++;
            return 'test';
        };

        $test = $func();

        $test === '';
    }
}

function test() {
    static $count = 0;

    if($count > 0) {
        throw new \Exception('Patate');
    }

    $count++;
    return 'test';
};
