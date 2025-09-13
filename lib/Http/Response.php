<?php

namespace Lib\Http;

class Response
{
    public function __construct(
        public readonly string $content,
        public readonly int $statusCode,
        public readonly array $headers,
    ){}

    public static function html(string $content, int $statusCode = 200): self
    {
        return new self(
            content: $content,
            statusCode: $statusCode,
            headers: []
        );
    }

    public static function json(mixed $data): self
    {
        $content = json_encode($data);

        $headers = [
            'Content-Type' => 'application/json',
        ];

        return new self(
            content: $content,
            statusCode: 200,
            headers: $headers,
        );
    }
}