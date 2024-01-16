from golang:1.21-alpine3.19 as build

workdir /app
copy . /app/
run go build -o afbapp

from alpine as runner
copy --from=build /app/afbapp /
CMD ./afbapp



