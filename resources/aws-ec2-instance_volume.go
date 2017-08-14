package resources

// AWS::EC2::Instance.Volume AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-mount-point.html
type AWSEC2InstanceVolume struct {

	// Device AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-mount-point.html#cfn-ec2-mountpoint-device
	Device string `json:"Device"`

	// VolumeId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-mount-point.html#cfn-ec2-mountpoint-volumeid
	VolumeId string `json:"VolumeId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2InstanceVolume) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.Volume"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2InstanceVolume) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}