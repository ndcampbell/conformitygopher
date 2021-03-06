# ConformityGopher

A tool for generating reports on AWS conformity.

ConformityGopher will look for specific tags or naming conventions in AWS resources and generate a report of resources that are not compliant.

Could be expanded to include other conformity rules for keeping AWS resources uniform. 

## Config

Location: `~/conformitygopher.json`

Example: [conformitygopher.json](conformitygopher.json.example)

## Rules

- `"required_tags": ["tag1", "tag2"]` - Validates that taggable resources have all required tags. Reports on resources missing tags.
- `"empty_elb": true` - repots on any ELB that has no registered instances
- Many more to come, open to any ideas

## Running
`go get github.com/ndcampbell/conformitygopher`
`go run conformitygopher/main.go`

## Dependencies

Handled via Govendor. 

- AWS SDK: https://github.com/aws/aws-sdk-go
- Gomail: https://github.com/go-gomail/gomail
