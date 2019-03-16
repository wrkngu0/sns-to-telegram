# SNS -> Lambda -> Telegram

**Description**

This is sns lambda subscriber that sends all messages to Telegram channel/group.

Written on Golang as application component for AWS Serverless Application Repository.

Designed as part of AWS ChatOps stack.

**Usage**

CloudFormation resource example:

```yaml

  PipelineSucceededEventRule:
    Type: AWS::Events::Rule
    Properties:
      Description: !Sub Pipeline Succeeded Event Rule for service ${ServiceName}
      EventPattern:
        source:
        - aws.codepipeline
        detail-type:
        - CodePipeline Stage Execution State Change
        detail:
          state:
          - SUCCEEDED
          pipeline:
          - !Sub ${AWS::StackName}
      State: "ENABLED"
      Targets:
      - Arn: !Sub 'arn:aws:sns:${AWS::Region}:${AWS::AccountId}:${NotificationTopicName}'
        Id: "SucceededTopic"
        InputTransformer:
          InputTemplate:
            Fn::Sub: >
              "Pipeline <pipeline> has succeeded in stage <stage>."
          InputPathsMap:
            pipeline: "$.detail.pipeline"
            stage: "$.detail.stage"

```


> **See [Serverless Application Model (SAM) HOWTO Guide](https://github.com/awslabs/serverless-application-model/blob/master/HOWTO.md) for more details in how to get started.**


