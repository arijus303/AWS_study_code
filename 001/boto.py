import boto3

# Replace these values with your actual AWS region and SQS queue URL
region_name = 'us-east-1a'
queue_url = 'https://sqs.us-east-1.amazonaws.com/549565409924/testq'

# Create an SQS client using default Boto3 configuration
sqs = boto3.client('sqs', region_name=region_name)

# Receive messages from the queue without deleting
response = sqs.receive_message(
    QueueUrl=queue_url,
    AttributeNames=[
        'All'
    ],
    MessageAttributeNames=[
        'All'
    ],
    MaxNumberOfMessages=10,  # Adjust as needed
    VisibilityTimeout=1,
    WaitTimeSeconds=10
)

# Extract and print messages
if 'Messages' in response:
    for message in response['Messages']:
        receipt_handle = message['ReceiptHandle']
        body = message['Body']
        print(f"Received message: {body}")



# No need to delete messages in this example as they are not removed from the queue
