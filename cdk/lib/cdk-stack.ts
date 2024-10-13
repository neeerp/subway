import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";

import * as s3 from "aws-cdk-lib/aws-s3";
import * as s3deploy from "aws-cdk-lib/aws-s3-deployment";

export class CdkStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, {
      ...props,
      analyticsReporting: false, // Disable CDK Metadata
    });

    const staticAssetBucket = new s3.Bucket(this, "StaticAssetBucket", {
      removalPolicy: cdk.RemovalPolicy.RETAIN,
    });

    new s3deploy.BucketDeployment(this, "DeployCSS", {
      sources: [s3deploy.Source.asset("../build/static")],
      destinationBucket: staticAssetBucket,
      destinationKeyPrefix: "",
    });
  }
}
