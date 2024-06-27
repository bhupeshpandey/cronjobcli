# cronjobcli

<p>To  run the code: </p>
<ol>
<li><p>Locally:</p>
<p>   go build -o cron-parser
  ./cron-parser &quot;*/15 0 1,15 * 1-5 /usr/bin/find&quot;</p>
</li>
<li><p>With docker:</p>
<p>   docker build -t cron-parser:latest .
   docker run --rm cron-parser:latest &quot;*/15 0 1,15 * 1-5 /usr/bin/find&quot;</p>
</li>
<li><p>Docker compose</p>
<p>  docker-compose up --build</p>
</li>
<li><p>Apply kubernetes cronjob:</p>
<p> kubectl apply -f cronjob.yaml</p>
</li>
</ol>


