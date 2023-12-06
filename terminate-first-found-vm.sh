INSTANCEID=$(aws ec2 describe-instances --query "Reservations[0].Instances[0].InstanceId" --output text)

aws ec2 terminate-instances --instance-ids $INSTANCEID --debug