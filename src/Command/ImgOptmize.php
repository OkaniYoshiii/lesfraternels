<?php

namespace App\Command;

use Lib\Command\Command;

class ImgOptmize extends Command
{
    public function run(array $args): void
    {
        $this->validate($args);

        $inputDir = $args[0];
        $outputDir = $args[1];

        $path = realpath($inputDir);
        $files = glob($path . '/*', GLOB_ERR);

        if($files === false) {
            Command::error(sprintf("Input directory \"%s\" does not exists", $inputDir));
            exit;
        }

        if(!is_dir($outputDir)) {
            mkdir($outputDir);
        }

        $out = realpath($outputDir);

        $formats = ['avif', 'webp', 'png'];
        $sizes = ['xs' => 80, 'sm' => 160, 'md' => 320, 'lg' => 640, 'xl' => 960, '2xl' => 1280, '3xl' => 1920, '4xl' => 1560];

        foreach($files as $file) {
            $filename = pathinfo($file, PATHINFO_FILENAME);
            $extension = pathinfo($file, PATHINFO_EXTENSION);

            foreach($formats as $format) {
                foreach($sizes as $sizeName => $sizeVal) {
                    $image = new \Imagick($file);

                    $outputFile = sprintf('%s/%s-%s.%s', $out,$filename, $sizeName, $format);

                    // $image->setFormat($format);
                    // self::info(sprintf("Image %s format changed from %s to %s", $file, $extension, $format));

                    $image->resizeImage($sizeVal, 0, 0, 0);
                    self::info(sprintf("Image %s has been resized to %spx wide\n", $file, $sizeVal));

                    $image->writeImage($outputFile);
                    self::info(sprintf("Image %s has been save in %s\n", $file, $outputFile));
                }
            }
        }
    }

    public function validate(array $args): void
    {
        $expected = 2;
        if(count($args) < $expected) {
            self::invalidArgumentsCount(expected: $expected, received: count($args));
            exit;
        }
    }
}