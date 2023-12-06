aws ec2 run-instances \
--image-id="ami-0230bd60aa48260c6" \
--instance-type="t2.micro" \
--security-groups="launch-wizard-2" \
--key-name="mymachinekey"