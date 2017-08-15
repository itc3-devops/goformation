package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DMS::EventSubscription AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-eventsubscription.html
type AWSDMSEventSubscription struct {

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-eventsubscription.html#cfn-dms-eventsubscription-enabled
	Enabled bool `json:"Enabled"`

	// EventCategories AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-eventsubscription.html#cfn-dms-eventsubscription-eventcategories
	EventCategories []string `json:"EventCategories"`

	// SnsTopicArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-eventsubscription.html#cfn-dms-eventsubscription-snstopicarn
	SnsTopicArn string `json:"SnsTopicArn"`

	// SourceIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-eventsubscription.html#cfn-dms-eventsubscription-sourceids
	SourceIds []string `json:"SourceIds"`

	// SourceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-eventsubscription.html#cfn-dms-eventsubscription-sourcetype
	SourceType string `json:"SourceType"`

	// SubscriptionName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-eventsubscription.html#cfn-dms-eventsubscription-subscriptionname
	SubscriptionName string `json:"SubscriptionName"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-eventsubscription.html#cfn-dms-eventsubscription-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDMSEventSubscription) AWSCloudFormationType() string {
	return "AWS::DMS::EventSubscription"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDMSEventSubscription) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDMSEventSubscriptionResources retrieves all AWSDMSEventSubscription items from a CloudFormation template
func GetAllAWSDMSEventSubscription(template *Template) map[string]*AWSDMSEventSubscription {

	results := map[string]*AWSDMSEventSubscription{}
	for name, resource := range template.Resources {
		result := &AWSDMSEventSubscription{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDMSEventSubscriptionWithName retrieves all AWSDMSEventSubscription items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDMSEventSubscription(name string, template *Template) (*AWSDMSEventSubscription, error) {

	result := &AWSDMSEventSubscription{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDMSEventSubscription{}, errors.New("resource not found")

}
