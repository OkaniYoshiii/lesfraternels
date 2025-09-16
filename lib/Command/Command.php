<?php

namespace Lib\Command;

abstract class Command
{
    public static function error(string $message): void
    {
        fprintf(STDERR, $message);
    }

    public static function info(string $message): void 
    {
        fprintf(STDOUT, $message);
    }

    public static function invalidArgumentsCount(int $expected, int $received): void
    {
        fprintf(STDERR, "Invalid number of arguments. Expected \"%s\", received \"%s\".\n", $expected, $received);
    }
}