"""
create an S3 bucket on AWS using boto3 python library
pass the name of the bucket as an argument variable when executed
"""

import sys
import boto3

try:
    def main():
        """This is the main function that will call other functions"""
        create_s3bucket(bucket_name)

except Exception as err:
    print(err)


def create_s3bucket(bucket_name):
    """s3_bucket variable calls boto3 library.
    client defines which AWS resource to use - in this case s3.
    bucket variable calls s3_bucket which invokes boto3 client
    pass bucketName and ACL parameters and Location. Location should match
    region in ~/.aws/config"""
    s3_bucket = boto3.client(
        's3',
    )
    bucket = s3_bucket.create_bucket(
        Bucket=bucket_name,
        ACL='private',
        CreateBucketConfiguration={'LocationConstraint': 'eu-west-2'}
    )
    print(bucket)


# bucketName variable defined by passing 1 argument when executed
bucket_name = sys.argv[1]

# when run standalone, not as an imported function, call the 'main' function
if __name__ == '__main__':
    main()
