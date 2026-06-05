.PHONY: tag print-tags go-get

TAG_VERSION := $(word 2,$(MAKECMDGOALS))
EXTRA_TAG_ARGS := $(wordlist 3,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

ifneq (,$(filter $(firstword $(MAKECMDGOALS)),tag print-tags go-get))
$(TAG_VERSION):
	@:
endif

tag:
	@if [ -z "$(TAG_VERSION)" ]; then \
		echo "usage: make tag v0.1.0"; \
		exit 1; \
	fi
	@if [ -n "$(EXTRA_TAG_ARGS)" ]; then \
		echo "too many arguments: $(EXTRA_TAG_ARGS)"; \
		echo "usage: make tag v0.1.0"; \
		exit 1; \
	fi
	@case "$(TAG_VERSION)" in \
		v[0-9]*.[0-9]*.[0-9]*) ;; \
		*) echo "tag version must look like v0.1.0"; exit 1 ;; \
	esac
	@if [ -n "$$(git status --porcelain)" ]; then \
		echo "working tree is not clean; commit or stash changes before tagging"; \
		exit 1; \
	fi
	@set -eu; \
	version="$(TAG_VERSION)"; \
	tags="$$version"; \
	for mod in $$(find . -name go.mod ! -path './go.mod' | sed 's#^\./##; s#/go.mod$$##' | sort); do \
		tags="$$tags $$mod/$$version"; \
	done; \
	for tag in $$tags; do \
		if git rev-parse -q --verify "refs/tags/$$tag" >/dev/null; then \
			echo "tag already exists locally: $$tag"; \
			exit 1; \
		fi; \
		if git ls-remote --exit-code --tags origin "refs/tags/$$tag" >/dev/null 2>&1; then \
			echo "tag already exists on origin: $$tag"; \
			exit 1; \
		fi; \
	done; \
	for tag in $$tags; do \
		echo "creating tag $$tag"; \
		git tag "$$tag"; \
	done; \
	echo "pushing tags: $$tags"; \
	git push origin $$tags

print-tags:
	@if [ -z "$(TAG_VERSION)" ]; then \
		echo "usage: make print-tags v0.1.0"; \
		exit 1; \
	fi
	@set -eu; \
	version="$(TAG_VERSION)"; \
	echo "$$version"; \
	for mod in $$(find . -name go.mod ! -path './go.mod' | sed 's#^\./##; s#/go.mod$$##' | sort); do \
		echo "$$mod/$$version"; \
	done

go-get:
	@if [ -z "$(TAG_VERSION)" ]; then \
		echo "usage: make go-get v0.1.0"; \
		exit 1; \
	fi
	@if [ -n "$(EXTRA_TAG_ARGS)" ]; then \
		echo "too many arguments: $(EXTRA_TAG_ARGS)"; \
		echo "usage: make go-get v0.1.0"; \
		exit 1; \
	fi
	@case "$(TAG_VERSION)" in \
		v[0-9]*.[0-9]*.[0-9]*) ;; \
		*) echo "tag version must look like v0.1.0"; exit 1 ;; \
	esac
	@set -eu; \
	version="$(TAG_VERSION)"; \
	tmp="$$(mktemp -d)"; \
	trap 'chmod -R u+w "$$tmp" 2>/dev/null || true; rm -rf "$$tmp"' EXIT; \
	cd "$$tmp"; \
	go mod init proxy-primer >/dev/null; \
	for gomod in "$(CURDIR)/go.mod" $$(find "$(CURDIR)" -name go.mod ! -path "$(CURDIR)/go.mod" | sort); do \
		module="$$(awk '/^module / { print $$2; exit }' "$$gomod")"; \
		echo "go get $$module@$$version"; \
		GOFLAGS="-modcacherw" GOMODCACHE="$$tmp/gomodcache" GOCACHE="$$tmp/gocache" GOPROXY="https://proxy.golang.org,direct" GOSUMDB="sum.golang.org" go get "$$module@$$version"; \
	done
