# Shell script to run the project on unix based systems

set -e

(
	cd "$(dirname "$0")"

	go build -o /build/InMemDB app/*.go
)

exec /build/InMemDB "$@"
