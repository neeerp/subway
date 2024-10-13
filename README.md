# Random Subway Station Exploration

This is a simple webpage hosted on AWS Lambda using go templ + HTMX. It suggests a random subway station to explore in Toronto!

See this hosted at https://subway.davidg.xyz

## Build and Deploy Instructions
Run `./bundle.sh` to regenerate templates, bundle lambda handlers, and compile
CSS with tailwind. 

Deploy by `cd`-ing into the `cdk/` directory and running `cdk deploy`. Note
that the infrastructure as code currently only creates the Lambda functions and
static asset S3 bucket, however it does not create all necessary
infrastructure. I'll try to get to that eventually.

## TODO List
- [X] Automate deployment to AWS so I don't have to manually upload this stuff every time I make a change
- [X] Use Tailwind CSS to make things ✨ p r e t t y ✨
- [ ] "Fix" the existing infrastructure definitions that are probably missing fields (I got things to work by importing my existing infrastructure into my CDK stack)
- [ ] Move static assets behind CloudFront
- [ ] Define the rest of the infrastructure using CDK
