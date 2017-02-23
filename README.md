# ConformityGopher

A tool for generating reports on AWS conformity.

ConformityGopher will look for specific tags or naming conventions in AWS resources and generate a report of resources that are not compliant.

Could be expanded to include other conformity rules for keeping AWS resources uniform. 

## Config

Location: `./conformitygopher.json`
Example:

    {
        "resources": ["ec2", "s3" ],
        "profiles": ["sandbox", "dev"],
        "db": {
            "type": "in-memory",
            "location": "./conformitygopher.db"
        },
        "rules": {
            ...
        }
    }

## Rules

- `"required_tags": ["tag1", "tag2"]` - Validates that taggable resources have all required tags. Reports on resources missing tags.
- `"empty_elb": true` - repots on any ELB that has no registered instances
- Many more to come, open to any ideas

## Running

`go run conformitygopher/main.go`

## Database

Currently uses Bolt in-memory database: https://github.com/boltdb/bolt
This provides an in-memory key-value store for generating the reports.

Would like to add redis support in the future

## Dependencies

- AWS SDK: https://github.com/aws/aws-sdk-go
- Bolt: https://github.com/boltdb/bolt
