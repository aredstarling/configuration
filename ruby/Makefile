-include .$(env).sh

.PHONY: spec

spec: test-spec staging-spec production-spec

test-spec:
	APP_ENV=test MARKETPLACES=US,DE,UK FIND_CATEGORY='DE|Küche, Haushalt & Wohnen' bundle exec rspec

staging-spec:
	APP_ENV=staging MARKETPLACES=US FIND_CATEGORY='US|Beauty & Personal Care' bundle exec rspec

production-spec:
	APP_ENV=production MARKETPLACES=US,DE,UK FIND_CATEGORY='DE|Küche, Haushalt & Wohnen' bundle exec rspec

analysis:
	bundle exec rubocop

format:
	bundle exec rubocop -a

console:
	./bin/console
