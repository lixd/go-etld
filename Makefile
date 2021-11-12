renew:
	curl https://publicsuffix.org/list/effective_tld_names.dat > eTLDList.dat

test:
	go test
