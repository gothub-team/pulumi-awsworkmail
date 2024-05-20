// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package types

import (
	"context"
	"reflect"

	"github.com/gothub-team/pulumi-awsworkmail/sdk/go/awsworkmail/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

type Domain struct {
}

type DomainArgs struct {
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil)).Elem()
}

func (i DomainArgs) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i DomainArgs) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

func (i *DomainArgs) ToOutput(ctx context.Context) pulumix.Output[*DomainArgs] {
	return pulumix.Val(i)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func (o DomainOutput) ToOutput(ctx context.Context) pulumix.Output[Domain] {
	return pulumix.Output[Domain]{
		OutputState: o.OutputState,
	}
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
}
