infinititest:
	 while :; do gradle  --parallel test  | tail -n +7 | grep -B 100 ':test FAILED' ; sleep 1; done
