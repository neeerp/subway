import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";

import * as s3 from "aws-cdk-lib/aws-s3";
import * as s3deploy from "aws-cdk-lib/aws-s3-deployment";
import * as lambda from "aws-cdk-lib/aws-lambda";

const BUILD_DIR = "../build";

export class CdkStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, {
      ...props,
      analyticsReporting: false, // Disable CDK Metadata
    });

    const subwayPageLambda = new lambda.Function(this, "SubwayPageLambda", {
      runtime: lambda.Runtime.PROVIDED_AL2023,
      code: lambda.Code.fromAsset(`${BUILD_DIR}/bundles/home.zip`),
      handler: "handler",
    });
    subwayPageLambda.applyRemovalPolicy(cdk.RemovalPolicy.RETAIN);
    subwayPageLambda.role?.applyRemovalPolicy(cdk.RemovalPolicy.RETAIN);

    const subwayStationLambda = new lambda.Function(
      this,
      "SubwayStationLambda",
      {
        runtime: lambda.Runtime.PROVIDED_AL2023,
        code: lambda.Code.fromAsset(`${BUILD_DIR}/bundles/station.zip`),
        handler: "handler",
      },
    );
    subwayStationLambda.applyRemovalPolicy(cdk.RemovalPolicy.RETAIN);
    subwayStationLambda.role?.applyRemovalPolicy(cdk.RemovalPolicy.RETAIN);

    const staticAssetBucket = new s3.Bucket(this, "StaticAssetBucket", {
      removalPolicy: cdk.RemovalPolicy.RETAIN,
    });

    new s3deploy.BucketDeployment(this, "DeployCSS", {
      sources: [s3deploy.Source.asset(`${BUILD_DIR}/static`)],
      destinationBucket: staticAssetBucket,
      destinationKeyPrefix: "",
    });
  }
}
