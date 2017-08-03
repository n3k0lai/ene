
.PHONY: test_server lint test_db

test_server:
	ngircd -n --config test/ngircd.conf

lint:
	tup
	git ls-files | grep '\.moon$$' | grep -v config.moon | grep -v stats_server | xargs -n 100 moonc -l

test_db:
	tup
	-dropdb -U postgres faye_test
	createdb -U postgres faye_test
	lapis migrate test

init_db: 
	tup
	-dropdb -U postgres faye
	createdb -U postgres faye
	lapis migrate
