# ConformityGopher

A tool for generating reports on AWS conformity.

ConformityGopher will look for specific tags or naming conventions in AWS resources and generate a report of resources that are not compliant.

Could be expanded to include other conformity rules for keeping AWS resources uniform. 

# Config

Location: `./conformitygopher.json`
Example:
    ```{
        "resources": ["ec2", "s3" ],
        "profiles": ["sandbox", "dev"],
        "requiredTags": [ "Group", "App" ],
        "reportEmail": [ "example@example.com", "admin@example.com"]
       }```

# Running

`go run conformitygopher/main.go`

# To Do
- Add AWS resources (Start: EC2, ELBs, S3, RDS)
- Add notifications (Start: Email, Slack)
- Decide how this runs (does it watch a schedule or just a one time command?)
- Added in-memory database (Bolt)



