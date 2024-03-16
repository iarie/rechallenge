# RE challenge

http://rechallenge-1287392037.eu-west-2.elb.amazonaws.com


## Project structure

- `/app` - contains the main http application
- `/cmd` - boostrap entrypoint
- `/data` - contains the data structures
- `/internal` - contains the business logic
- `/internal/packer.go` - algorithm
- `/deploy` - terraform scripts
- `/static` - static assets
- `/templates` - html templates

## Tests

To run dockerized tests:
```bash
make testd
```