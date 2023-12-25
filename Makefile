generate-in-container: ## resync client with current graph endpoint
	hack/gen-api-client.sh

update-schema: ## download schema from plural
	curl -L https://raw.githubusercontent.com/pluralsh/console/master/schema/schema.graphql --output schema/schema.graphql

generate: update-schema
	go run github.com/Yamashou/gqlgenc

release-vsn:
	@read -p "Version: " tag; \
	git checkout main; \
	git pull --rebase; \
	git tag -a $$tag -m "new release"; \
	git push origin $$tag;