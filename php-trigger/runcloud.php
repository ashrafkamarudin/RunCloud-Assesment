#!/usr/bin/env php

<?php

require_once __DIR__ . '/vendor/autoload.php';

use Symfony\Component\Console\Application;
use RunCloud\TimeCommand;

$app = new Application();
$app -> add(new TimeCommand());
$app -> run();