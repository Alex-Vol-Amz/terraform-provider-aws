// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package s3control

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceAccountPublicAccessBlock,
			TypeName: "aws_s3_account_public_access_block",
		},
		{
			Factory:  dataSourceMultiRegionAccessPoint,
			TypeName: "aws_s3control_multi_region_access_point",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccessPoint,
			TypeName: "aws_s3_access_point",
		},
		{
			Factory:  resourceAccountPublicAccessBlock,
			TypeName: "aws_s3_account_public_access_block",
		},
		{
			Factory:  resourceAccessPointPolicy,
			TypeName: "aws_s3control_access_point_policy",
		},
		{
			Factory:  resourceBucket,
			TypeName: "aws_s3control_bucket",
		},
		{
			Factory:  resourceBucketLifecycleConfiguration,
			TypeName: "aws_s3control_bucket_lifecycle_configuration",
		},
		{
			Factory:  resourceBucketPolicy,
			TypeName: "aws_s3control_bucket_policy",
		},
		{
			Factory:  resourceMultiRegionAccessPoint,
			TypeName: "aws_s3control_multi_region_access_point",
		},
		{
			Factory:  resourceMultiRegionAccessPointPolicy,
			TypeName: "aws_s3control_multi_region_access_point_policy",
		},
		{
			Factory:  resourceObjectLambdaAccessPoint,
			TypeName: "aws_s3control_object_lambda_access_point",
		},
		{
			Factory:  resourceObjectLambdaAccessPointPolicy,
			TypeName: "aws_s3control_object_lambda_access_point_policy",
		},
		{
			Factory:  resourceStorageLensConfiguration,
			TypeName: "aws_s3control_storage_lens_configuration",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.S3Control
}

var ServicePackage = &servicePackage{}
