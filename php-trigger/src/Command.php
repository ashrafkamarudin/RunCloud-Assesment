<?php namespace RunCloud;

use Symfony\Component\Console\Command\Command as SymfonyCommand;
use Symfony\Component\Console\Input\InputArgument;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;
use GuzzleHttp\Client;

/**
 * Author: Ashraf Kamarudin <ashrafkamarudin1995@gmail.com>
 */
class Command extends SymfonyCommand
{
    
    public function __construct()
    {
        parent::__construct();
    }

    protected function sendRequest(InputInterface $input, OutputInterface $output)
    {
        $client = new Client([
            // Base URI is used with relative requests
            'base_uri' => 'http://192.168.0.6:1323',
            // You can set any number of default request options.
            'timeout'  => 2.0,
        ]);

        $response = $client->request('GET', '/');

        //$response = $client->post('http://httpbin.org/post');
        
    }
}