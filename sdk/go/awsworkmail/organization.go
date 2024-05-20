// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsworkmail

import (
	"context"
	"reflect"

	"github.com/gothub-team/pulumi-awsworkmail/sdk/go/awsworkmail/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type Organization struct {
	pulumi.CustomResourceState
}

// NewOrganization registers a new resource with the given unique name, arguments, and options.
func NewOrganization(ctx *pulumi.Context,
	name string, args *OrganizationArgs, opts ...pulumi.ResourceOption) (*Organization, error) {
	if args == nil {
		args = &OrganizationArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Organization
	err := ctx.RegisterResource("awsworkmail:index:Organization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganization gets an existing Organization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationState, opts ...pulumi.ResourceOption) (*Organization, error) {
	var resource Organization
	err := ctx.ReadResource("awsworkmail:index:Organization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Organization resources.
type organizationState struct {
}

type OrganizationState struct {
}

func (OrganizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationState)(nil)).Elem()
}

type organizationArgs struct {
}

// The set of arguments for constructing a Organization resource.
type OrganizationArgs struct {
}

func (OrganizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationArgs)(nil)).Elem()
}

type OrganizationOutput struct{ *pulumi.OutputState }

func (OrganizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Organization)(nil)).Elem()
}

func (o OrganizationOutput) ToOrganizationOutput() OrganizationOutput {
	return o
}

func (o OrganizationOutput) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return o
}

func (o OrganizationOutput) ToOutput(ctx context.Context) pulumix.Output[Organization] {
	return pulumix.Output[Organization]{
		OutputState: o.OutputState,
	}
}

func init() {
	pulumi.RegisterOutputType(OrganizationOutput{})
}
